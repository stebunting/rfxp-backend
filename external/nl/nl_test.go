package nl_test

import (
	"log"
	"testing"

	"github.com/stebunting/rfxp-backend/external/nl"
)

func TestValidNl(t *testing.T) {
	type TestChannel struct {
		Channel  int
		Indoors  bool
		Outdoors bool
	}

	type TestCase struct {
		PlaceName string
		Latitude  float64
		Longitude float64
		Channels  []TestChannel
	}

	testCases := []TestCase{
		{
			PlaceName: "Apeldoorn",
			Latitude:  52.206008,
			Longitude: 5.972186,
			Channels: []TestChannel{
				{Channel: 21, Indoors: true, Outdoors: false},
				{Channel: 22, Indoors: true, Outdoors: false},
				{Channel: 23, Indoors: true, Outdoors: false},
				{Channel: 24, Indoors: true, Outdoors: true},
				{Channel: 25, Indoors: true, Outdoors: true},
				{Channel: 26, Indoors: false, Outdoors: false},
				{Channel: 27, Indoors: true, Outdoors: true},
				{Channel: 28, Indoors: false, Outdoors: false},
				{Channel: 29, Indoors: true, Outdoors: true},
				{Channel: 30, Indoors: false, Outdoors: false},
				{Channel: 31, Indoors: false, Outdoors: false},
				{Channel: 32, Indoors: true, Outdoors: true},
				{Channel: 33, Indoors: true, Outdoors: true},
				{Channel: 34, Indoors: true, Outdoors: true},
				{Channel: 35, Indoors: true, Outdoors: true},
				{Channel: 36, Indoors: false, Outdoors: false},
				{Channel: 37, Indoors: true, Outdoors: true},
				{Channel: 38, Indoors: false, Outdoors: false},
				{Channel: 39, Indoors: true, Outdoors: false},
				{Channel: 40, Indoors: true, Outdoors: true},
				{Channel: 41, Indoors: true, Outdoors: true},
				{Channel: 42, Indoors: false, Outdoors: false},
				{Channel: 43, Indoors: true, Outdoors: true},
				{Channel: 44, Indoors: false, Outdoors: false},
				{Channel: 45, Indoors: true, Outdoors: false},
				{Channel: 46, Indoors: true, Outdoors: false},
				{Channel: 47, Indoors: false, Outdoors: false},
				{Channel: 48, Indoors: true, Outdoors: false},
				{Channel: 49, Indoors: false, Outdoors: false},
			},
		}, {
			PlaceName: "Bergen Op Zoom",
			Latitude:  51.484943,
			Longitude: 4.281107,
			Channels: []TestChannel{
				{Channel: 21, Indoors: false, Outdoors: false},
				{Channel: 22, Indoors: true, Outdoors: false},
				{Channel: 23, Indoors: true, Outdoors: true},
				{Channel: 24, Indoors: true, Outdoors: true},
				{Channel: 25, Indoors: true, Outdoors: false},
				{Channel: 26, Indoors: true, Outdoors: false},
				{Channel: 27, Indoors: true, Outdoors: false},
				{Channel: 28, Indoors: true, Outdoors: true},
				{Channel: 29, Indoors: false, Outdoors: false},
				{Channel: 30, Indoors: true, Outdoors: false},
				{Channel: 31, Indoors: true, Outdoors: false},
				{Channel: 32, Indoors: false, Outdoors: false},
				{Channel: 33, Indoors: false, Outdoors: false},
				{Channel: 34, Indoors: true, Outdoors: false},
				{Channel: 35, Indoors: false, Outdoors: false},
				{Channel: 36, Indoors: true, Outdoors: true},
				{Channel: 37, Indoors: false, Outdoors: false},
				{Channel: 38, Indoors: false, Outdoors: false},
				{Channel: 39, Indoors: true, Outdoors: true},
				{Channel: 40, Indoors: false, Outdoors: false},
				{Channel: 41, Indoors: false, Outdoors: false},
				{Channel: 42, Indoors: true, Outdoors: true},
				{Channel: 43, Indoors: false, Outdoors: false},
				{Channel: 44, Indoors: false, Outdoors: false},
				{Channel: 45, Indoors: true, Outdoors: false},
				{Channel: 46, Indoors: false, Outdoors: false},
				{Channel: 47, Indoors: false, Outdoors: false},
				{Channel: 48, Indoors: false, Outdoors: false},
				{Channel: 49, Indoors: false, Outdoors: false},
			},
		}, {
			PlaceName: "Emmen",
			Latitude:  52.754689,
			Longitude: 6.934979,
			Channels: []TestChannel{
				{Channel: 21, Indoors: true, Outdoors: false},
				{Channel: 22, Indoors: false, Outdoors: false},
				{Channel: 23, Indoors: false, Outdoors: false},
				{Channel: 24, Indoors: true, Outdoors: false},
				{Channel: 25, Indoors: true, Outdoors: false},
				{Channel: 26, Indoors: true, Outdoors: true},
				{Channel: 27, Indoors: true, Outdoors: false},
				{Channel: 28, Indoors: true, Outdoors: true},
				{Channel: 29, Indoors: true, Outdoors: true},
				{Channel: 30, Indoors: true, Outdoors: true},
				{Channel: 31, Indoors: true, Outdoors: true},
				{Channel: 32, Indoors: true, Outdoors: true},
				{Channel: 33, Indoors: true, Outdoors: false},
				{Channel: 34, Indoors: true, Outdoors: true},
				{Channel: 35, Indoors: true, Outdoors: false},
				{Channel: 36, Indoors: false, Outdoors: false},
				{Channel: 37, Indoors: false, Outdoors: false},
				{Channel: 38, Indoors: false, Outdoors: false},
				{Channel: 39, Indoors: true, Outdoors: true},
				{Channel: 40, Indoors: true, Outdoors: false},
				{Channel: 41, Indoors: false, Outdoors: false},
				{Channel: 42, Indoors: true, Outdoors: false},
				{Channel: 43, Indoors: true, Outdoors: true},
				{Channel: 44, Indoors: false, Outdoors: false},
				{Channel: 45, Indoors: false, Outdoors: false},
				{Channel: 46, Indoors: false, Outdoors: false},
				{Channel: 47, Indoors: false, Outdoors: false},
				{Channel: 48, Indoors: true, Outdoors: false},
				{Channel: 49, Indoors: false, Outdoors: false},
			},
		}, {
			PlaceName: "Alkmaar",
			Latitude:  52.645318,
			Longitude: 4.771771,
			Channels: []TestChannel{
				{Channel: 21, Indoors: true, Outdoors: false},
				{Channel: 22, Indoors: true, Outdoors: true},
				{Channel: 23, Indoors: false, Outdoors: false},
				{Channel: 24, Indoors: false, Outdoors: false},
				{Channel: 25, Indoors: true, Outdoors: false},
				{Channel: 26, Indoors: true, Outdoors: false},
				{Channel: 27, Indoors: false, Outdoors: false},
				{Channel: 28, Indoors: true, Outdoors: true},
				{Channel: 29, Indoors: true, Outdoors: true},
				{Channel: 30, Indoors: true, Outdoors: true},
				{Channel: 31, Indoors: true, Outdoors: true},
				{Channel: 32, Indoors: false, Outdoors: false},
				{Channel: 33, Indoors: true, Outdoors: false},
				{Channel: 34, Indoors: false, Outdoors: false},
				{Channel: 35, Indoors: true, Outdoors: true},
				{Channel: 36, Indoors: false, Outdoors: false},
				{Channel: 37, Indoors: true, Outdoors: true},
				{Channel: 38, Indoors: false, Outdoors: false},
				{Channel: 39, Indoors: true, Outdoors: false},
				{Channel: 40, Indoors: true, Outdoors: false},
				{Channel: 41, Indoors: true, Outdoors: true},
				{Channel: 42, Indoors: true, Outdoors: true},
				{Channel: 43, Indoors: true, Outdoors: true},
				{Channel: 44, Indoors: false, Outdoors: false},
				{Channel: 45, Indoors: false, Outdoors: false},
				{Channel: 46, Indoors: true, Outdoors: false},
				{Channel: 47, Indoors: false, Outdoors: false},
				{Channel: 48, Indoors: true, Outdoors: false},
				{Channel: 49, Indoors: false, Outdoors: false},
			},
		},
	}

	for _, test := range testCases {
		s := nl.Netherlands{Latitude: test.Latitude, Longitude: test.Longitude}
		c, err := s.Call()
		if err != nil {
			log.Fatalf("unexpected error making network call")
		}
		channels := *c

		for i := 0; i < len(channels); i++ {
			if channels[i].Number != test.Channels[i].Channel {
				log.Fatalf("invalid channel in %s... expected %d, got %d", test.PlaceName, test.Channels[i].Channel, channels[i].Number)
			}
			if channels[i].Indoors != test.Channels[i].Indoors {
				log.Fatalf("invalid indoors availability in %s channel %d... expected %v, got %v", test.PlaceName, test.Channels[i].Channel, test.Channels[i].Indoors, channels[i].Indoors)
			}
			if channels[i].Outdoors != test.Channels[i].Outdoors {
				log.Fatalf("invalid outdoors availability in %s channel %d... expected %v, got %v", test.PlaceName, test.Channels[i].Channel, test.Channels[i].Outdoors, channels[i].Outdoors)
			}
		}
	}
}

func TestInvalidNl(t *testing.T) {
	s := nl.Netherlands{
		Latitude:  57.043188,
		Longitude: 49.921598,
	}
	_, err := s.Call()
	if err == nil {
		log.Fatalf("expected error making network call")
	}
}
