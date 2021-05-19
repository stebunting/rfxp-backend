package no

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"

	"github.com/getsentry/sentry-go"
	"github.com/stebunting/rfxp-backend/channel"
)

type Norway struct {
	Latitude  float64
	Longitude float64
}

type Result struct {
	Id              int     `json:"id"`
	Name            string  `json:"name"`
	Channel         string  `json:"channel"`
	Type            string  `json:"type"`
	LoFrequency     float64 `json:"minfreq"`
	CenterFrequency float64 `json:"frequence"`
	HiFrequency     float64 `json:"maxfreq"`
	Warning         bool    `json:"warning"`
}

func (s *Norway) Call() (*[]channel.Channel, error) {
	result, err := s.makeApiCall()
	if err != nil {
		return nil, err
	}
	channels := s.channelsFromApiResponse(result)
	return channels, nil
}

func (s *Norway) makeApiCall() (*[]Result, error) {
	url, err := url.Parse("https://finnsenderen.no/finnsenderen_service/rest/ledigefrekvenser")
	if err != nil {
		sentry.CaptureException(err)
		panic(err)
	}

	q := url.Query()
	q.Set("type", "mikrofon")
	q.Set("latitude", fmt.Sprintf("%f", s.Latitude))
	q.Set("longitude", fmt.Sprintf("%f", s.Longitude))
	url.RawQuery = q.Encode()

	request, err := http.NewRequest(http.MethodGet, url.String(), nil)
	if err != nil {
		sentry.CaptureException(err)
		panic(err)
	}
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Referer", "https://finnsenderen.no/")

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

	var response []Result
	err = json.Unmarshal(body, &response)
	if err != nil {
		sentry.CaptureException(err)
		panic(err)
	}

	if len(response) == 0 {
		sentry.CaptureMessage("no results")
		return nil, errors.New("no results")
	}

	return &response, nil
}

func (s *Norway) channelsFromApiResponse(result *[]Result) *[]channel.Channel {
	startChannel := 21
	endChannel := 48
	channels := s.initChannels()

	for _, r := range *result {
		ch, err := strconv.Atoi(r.Channel)
		if err != nil {
			continue
		}

		if ch <= endChannel {
			index := ch - startChannel
			channels[index].Indoors = true
			channels[index].Outdoors = true
		}
	}

	return &channels
}

func (s *Norway) initChannels() []channel.Channel {
	return []channel.Channel{
		{Number: 21, FreqStart: 470000, FreqEnd: 478000, Indoors: false, Outdoors: false},
		{Number: 22, FreqStart: 478000, FreqEnd: 486000, Indoors: false, Outdoors: false},
		{Number: 23, FreqStart: 486000, FreqEnd: 494000, Indoors: false, Outdoors: false},
		{Number: 24, FreqStart: 494000, FreqEnd: 502000, Indoors: false, Outdoors: false},
		{Number: 25, FreqStart: 502000, FreqEnd: 510000, Indoors: false, Outdoors: false},
		{Number: 26, FreqStart: 510000, FreqEnd: 518000, Indoors: false, Outdoors: false},
		{Number: 27, FreqStart: 518000, FreqEnd: 526000, Indoors: false, Outdoors: false},
		{Number: 28, FreqStart: 526000, FreqEnd: 534000, Indoors: false, Outdoors: false},
		{Number: 29, FreqStart: 534000, FreqEnd: 542000, Indoors: false, Outdoors: false},
		{Number: 30, FreqStart: 542000, FreqEnd: 550000, Indoors: false, Outdoors: false},
		{Number: 31, FreqStart: 550000, FreqEnd: 558000, Indoors: false, Outdoors: false},
		{Number: 32, FreqStart: 558000, FreqEnd: 566000, Indoors: false, Outdoors: false},
		{Number: 33, FreqStart: 566000, FreqEnd: 574000, Indoors: false, Outdoors: false},
		{Number: 34, FreqStart: 574000, FreqEnd: 582000, Indoors: false, Outdoors: false},
		{Number: 35, FreqStart: 582000, FreqEnd: 590000, Indoors: false, Outdoors: false},
		{Number: 36, FreqStart: 590000, FreqEnd: 598000, Indoors: false, Outdoors: false},
		{Number: 37, FreqStart: 598000, FreqEnd: 606000, Indoors: false, Outdoors: false},
		{Number: 38, FreqStart: 606000, FreqEnd: 614000, Indoors: false, Outdoors: false},
		{Number: 39, FreqStart: 614000, FreqEnd: 622000, Indoors: false, Outdoors: false},
		{Number: 40, FreqStart: 622000, FreqEnd: 630000, Indoors: false, Outdoors: false},
		{Number: 41, FreqStart: 630000, FreqEnd: 638000, Indoors: false, Outdoors: false},
		{Number: 42, FreqStart: 638000, FreqEnd: 646000, Indoors: false, Outdoors: false},
		{Number: 43, FreqStart: 646000, FreqEnd: 654000, Indoors: false, Outdoors: false},
		{Number: 44, FreqStart: 654000, FreqEnd: 662000, Indoors: false, Outdoors: false},
		{Number: 45, FreqStart: 662000, FreqEnd: 670000, Indoors: false, Outdoors: false},
		{Number: 46, FreqStart: 670000, FreqEnd: 678000, Indoors: false, Outdoors: false},
		{Number: 47, FreqStart: 678000, FreqEnd: 686000, Indoors: false, Outdoors: false},
		{Number: 48, FreqStart: 686000, FreqEnd: 694000, Indoors: false, Outdoors: false},
	}
}
