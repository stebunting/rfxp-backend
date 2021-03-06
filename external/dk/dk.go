package dk

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/getsentry/sentry-go"
	"github.com/stebunting/rfxp-backend/channel"
)

type Denmark struct {
	Latitude  float64
	Longitude float64
}

type ApiResponse struct {
	Status            string    `json:"status"`
	LatLng            []float64 `json:"latlng"`
	NorthingsEastings []int     `json:"northingseastings"`
	Results           []Results `json:"results"`
}

type Results struct {
	FrequencyAreas        [][]int `json:"frequencyAreas"`
	GuardBand             int     `json:"guardBand"`
	LastChannel           int     `json:"lastChannel"`
	TvChannelsNoGuardBand [][]int `json:"tvChannelsNoGuardBand"`
}

func (s *Denmark) GetCountryName() string {
	return "Denmark"
}

func (s *Denmark) GetServiceName() string {
	return "Energistyrelsen"
}

func (s *Denmark) Call() (*[]channel.Channel, error) {
	result, err := s.makeApiCall()
	if err != nil {
		return nil, err
	}
	channels := s.channelsFromApiResponse(result)
	return channels, nil
}

func (s *Denmark) makeApiCall() (*[][]int, error) {
	url, err := url.Parse("https://frekvens.ens.dk/findKanalerAPI.php")
	if err != nil {
		sentry.CaptureException(err)
		panic(err)
	}

	q := url.Query()
	q.Set("output", "JSON")
	q.Set("language", "en")
	q.Set("lat", fmt.Sprintf("%f", s.Latitude))
	q.Set("lng", fmt.Sprintf("%f", s.Longitude))
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
		sentry.CaptureException(err)
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
		sentry.CaptureException(err)
		panic(err)
	}

	if response.Status != "OK" {
		sentry.CaptureMessage(response.Status)
		return nil, errors.New(response.Status)
	}

	return &response.Results[0].TvChannelsNoGuardBand, nil
}

func (s *Denmark) channelsFromApiResponse(result *[][]int) *[]channel.Channel {
	startFrequency := 470000
	startChannel := 21
	endChannel := 48
	chWidth := 8000

	channels := []channel.Channel{}
	freqCounter := startFrequency
	apiIndex := 0
	apiResult := (*result)[apiIndex]
	for ch := startChannel; ch <= endChannel; ch++ {
		if ch > apiResult[1] && apiIndex < len(*result)-1 {
			apiIndex++
			apiResult = (*result)[apiIndex]
		}

		startFrequency := freqCounter
		endFrequency := startFrequency + chWidth
		available := ch >= apiResult[0] && ch <= apiResult[1]

		channels = append(channels, channel.Channel{
			Number:    ch,
			FreqStart: startFrequency,
			FreqEnd:   endFrequency,
			Indoors:   available,
			Outdoors:  available,
		})

		freqCounter += chWidth
	}

	return &channels
}
