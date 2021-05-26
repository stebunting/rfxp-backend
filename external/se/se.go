package se

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"sync"
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/stebunting/rfxp-backend/channel"
)

type Sweden struct {
	Latitude  float64
	Longitude float64
}

type ApiResponse struct {
	Success            bool
	ErrorMessage       string
	ResultGeneratedAt  time.Time
	Frequencies        []FrequencyInfo
	FrequencyBundles   []FrequencyBundleInfo
	CoordinatesQueried string
	Wgs84Lat           float64
	Wgs84Lng           float64
}

type FrequencyInfo struct {
	IsFree          bool
	BlockCode       int
	BlockDate       time.Time
	CenterFrequency int
	LowFrequency    int
	HighFrequency   int
	Band            string
	Channel         int
}

type FrequencyBundleInfo struct {
	IsFree        bool
	BlockCode     int
	BlockDate     time.Time
	LowFrequency  int
	HighFrequency int
	Band          string
	FirstChannel  int
	LastChannel   int
}

func (s *Sweden) GetCountryName() string {
	return "Sweden"
}

func (s *Sweden) GetServiceName() string {
	return "PTS Trådlös ljudöverföring"
}

func (s *Sweden) Call() (*[]channel.Channel, error) {
	var indoors *[]FrequencyBundleInfo
	var outdoors *[]FrequencyBundleInfo
	sentry.CaptureMessage("Message from SE")

	var wg sync.WaitGroup
	wg.Add(2)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		indoors, _ = s.makeApiCall(true)
	}(&wg)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		outdoors, _ = s.makeApiCall(false)
	}(&wg)
	wg.Wait()

	channels := s.channelsFromApiResponse(indoors, outdoors)

	return channels, nil
}

func (s *Sweden) makeApiCall(indoors bool) (*[]FrequencyBundleInfo, error) {
	url, err := url.Parse("https://wirelessaudio.pts.se/api/WirelessAudioTransmission/CheckLocation")
	if err != nil {
		sentry.CaptureException(err)
		panic(err)
	}

	q := url.Query()
	q.Set("q", fmt.Sprintf("%f,%f", s.Latitude, s.Longitude))
	if indoors {
		q.Set("indoor", "true")
	} else {
		q.Set("indoor", "false")
	}
	url.RawQuery = q.Encode()

	request, err := http.NewRequest(http.MethodGet, url.String(), nil)
	if err != nil {
		sentry.CaptureException(err)
		panic(err)
	}
	request.Header.Set("Accept", "application/json")

	client := &http.Client{}
	rawResponse, err := client.Do(request)
	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(rawResponse.Body)
	if err != nil {
		sentry.CaptureException(err)
		panic(err)
	}
	defer rawResponse.Body.Close()

	var response ApiResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		panic(err)
	}

	if !response.Success {
		sentry.CaptureMessage(response.ErrorMessage)
		return nil, errors.New(response.ErrorMessage)
	}

	return &response.FrequencyBundles, nil
}

func (s *Sweden) channelsFromApiResponse(
	indoors *[]FrequencyBundleInfo,
	outdoors *[]FrequencyBundleInfo,
) *[]channel.Channel {
	startFrequency := 470000
	startChannel := 21
	endChannel := 48
	chWidth := 8000

	channels := []channel.Channel{}
	freqCounter := startFrequency
	inIdx := 0
	outIdx := 0
	apiIndoors := (*indoors)[inIdx]
	apiOutdoors := (*outdoors)[outIdx]
	for ch := startChannel; ch <= endChannel; ch++ {
		if ch > apiIndoors.LastChannel && inIdx < len(*indoors)-1 {
			inIdx++
			apiIndoors = (*indoors)[inIdx]
		}

		if ch > apiOutdoors.LastChannel && outIdx < len(*outdoors)-1 {
			outIdx++
			apiOutdoors = (*outdoors)[outIdx]
		}

		startFrequency := freqCounter
		endFrequency := startFrequency + chWidth
		indoors := false
		if ch >= apiIndoors.FirstChannel && ch <= apiIndoors.LastChannel {
			indoors = apiIndoors.IsFree
		}
		outdoors := false
		if ch >= apiOutdoors.FirstChannel && ch <= apiOutdoors.LastChannel {
			outdoors = apiOutdoors.IsFree
		}

		channels = append(channels, channel.Channel{
			Number:    ch,
			FreqStart: startFrequency,
			FreqEnd:   endFrequency,
			Indoors:   indoors,
			Outdoors:  outdoors,
		})

		freqCounter += chWidth
	}

	return &channels
}
