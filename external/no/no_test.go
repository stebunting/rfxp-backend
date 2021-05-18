package no_test

import (
	"log"
	"testing"

	"github.com/stebunting/rfxp-backend/external/no"
)

func TestNo(t *testing.T) {
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
			PlaceName: "Oslo",
			Latitude:  59.92341,
			Longitude: 10.62288,
			Channels: []TestChannel{
				{Channel: 21, Indoors: false, Outdoors: false},
				{Channel: 22, Indoors: false, Outdoors: false},
				{Channel: 23, Indoors: false, Outdoors: false},
				{Channel: 24, Indoors: false, Outdoors: false},
				{Channel: 25, Indoors: false, Outdoors: false},
				{Channel: 26, Indoors: true, Outdoors: true},
				{Channel: 27, Indoors: true, Outdoors: true},
				{Channel: 28, Indoors: false, Outdoors: false},
				{Channel: 29, Indoors: false, Outdoors: false},
				{Channel: 30, Indoors: false, Outdoors: false},
				{Channel: 31, Indoors: false, Outdoors: false},
				{Channel: 32, Indoors: true, Outdoors: true},
				{Channel: 33, Indoors: false, Outdoors: false},
				{Channel: 34, Indoors: false, Outdoors: false},
				{Channel: 35, Indoors: true, Outdoors: true},
				{Channel: 36, Indoors: true, Outdoors: true},
				{Channel: 37, Indoors: false, Outdoors: false},
				{Channel: 38, Indoors: false, Outdoors: false},
				{Channel: 39, Indoors: true, Outdoors: true},
				{Channel: 40, Indoors: false, Outdoors: false},
				{Channel: 41, Indoors: false, Outdoors: false},
				{Channel: 42, Indoors: false, Outdoors: false},
				{Channel: 43, Indoors: false, Outdoors: false},
				{Channel: 44, Indoors: false, Outdoors: false},
				{Channel: 45, Indoors: true, Outdoors: true},
				{Channel: 46, Indoors: true, Outdoors: true},
				{Channel: 47, Indoors: true, Outdoors: true},
				{Channel: 48, Indoors: false, Outdoors: false},
				{Channel: 49, Indoors: false, Outdoors: false},
			},
		}, {
			PlaceName: "Tromso",
			Latitude:  69.66946,
			Longitude: 18.92116,
			Channels: []TestChannel{
				{Channel: 21, Indoors: false, Outdoors: false},
				{Channel: 22, Indoors: false, Outdoors: false},
				{Channel: 23, Indoors: false, Outdoors: false},
				{Channel: 24, Indoors: false, Outdoors: false},
				{Channel: 25, Indoors: false, Outdoors: false},
				{Channel: 26, Indoors: false, Outdoors: false},
				{Channel: 27, Indoors: false, Outdoors: false},
				{Channel: 28, Indoors: true, Outdoors: true},
				{Channel: 29, Indoors: false, Outdoors: false},
				{Channel: 30, Indoors: true, Outdoors: true},
				{Channel: 31, Indoors: true, Outdoors: true},
				{Channel: 32, Indoors: true, Outdoors: true},
				{Channel: 33, Indoors: true, Outdoors: true},
				{Channel: 34, Indoors: false, Outdoors: false},
				{Channel: 35, Indoors: true, Outdoors: true},
				{Channel: 36, Indoors: true, Outdoors: true},
				{Channel: 37, Indoors: false, Outdoors: false},
				{Channel: 38, Indoors: true, Outdoors: true},
				{Channel: 39, Indoors: false, Outdoors: false},
				{Channel: 40, Indoors: false, Outdoors: false},
				{Channel: 41, Indoors: true, Outdoors: true},
				{Channel: 42, Indoors: false, Outdoors: false},
				{Channel: 43, Indoors: false, Outdoors: false},
				{Channel: 44, Indoors: true, Outdoors: true},
				{Channel: 45, Indoors: true, Outdoors: true},
				{Channel: 46, Indoors: false, Outdoors: false},
				{Channel: 47, Indoors: false, Outdoors: false},
				{Channel: 48, Indoors: true, Outdoors: true},
				{Channel: 49, Indoors: false, Outdoors: false},
			},
		}, {
			PlaceName: "Alesund",
			Latitude:  62.47073,
			Longitude: 6.14165,
			Channels: []TestChannel{
				{Channel: 21, Indoors: false, Outdoors: false},
				{Channel: 22, Indoors: false, Outdoors: false},
				{Channel: 23, Indoors: false, Outdoors: false},
				{Channel: 24, Indoors: false, Outdoors: false},
				{Channel: 25, Indoors: false, Outdoors: false},
				{Channel: 26, Indoors: true, Outdoors: true},
				{Channel: 27, Indoors: true, Outdoors: true},
				{Channel: 28, Indoors: false, Outdoors: false},
				{Channel: 29, Indoors: false, Outdoors: false},
				{Channel: 30, Indoors: true, Outdoors: true},
				{Channel: 31, Indoors: false, Outdoors: false},
				{Channel: 32, Indoors: true, Outdoors: true},
				{Channel: 33, Indoors: true, Outdoors: true},
				{Channel: 34, Indoors: false, Outdoors: false},
				{Channel: 35, Indoors: false, Outdoors: false},
				{Channel: 36, Indoors: true, Outdoors: true},
				{Channel: 37, Indoors: false, Outdoors: false},
				{Channel: 38, Indoors: false, Outdoors: false},
				{Channel: 39, Indoors: false, Outdoors: false},
				{Channel: 40, Indoors: false, Outdoors: false},
				{Channel: 41, Indoors: true, Outdoors: true},
				{Channel: 42, Indoors: false, Outdoors: false},
				{Channel: 43, Indoors: true, Outdoors: true},
				{Channel: 44, Indoors: true, Outdoors: true},
				{Channel: 45, Indoors: true, Outdoors: true},
				{Channel: 46, Indoors: false, Outdoors: false},
				{Channel: 47, Indoors: true, Outdoors: true},
				{Channel: 48, Indoors: true, Outdoors: true},
				{Channel: 49, Indoors: false, Outdoors: false},
			},
		}, {
			PlaceName: "Hemsedal",
			Latitude:  60.87947,
			Longitude: 8.47565,
			Channels: []TestChannel{
				{Channel: 21, Indoors: false, Outdoors: false},
				{Channel: 22, Indoors: false, Outdoors: false},
				{Channel: 23, Indoors: false, Outdoors: false},
				{Channel: 24, Indoors: false, Outdoors: false},
				{Channel: 25, Indoors: false, Outdoors: false},
				{Channel: 26, Indoors: false, Outdoors: false},
				{Channel: 27, Indoors: true, Outdoors: true},
				{Channel: 28, Indoors: true, Outdoors: true},
				{Channel: 29, Indoors: true, Outdoors: true},
				{Channel: 30, Indoors: false, Outdoors: false},
				{Channel: 31, Indoors: true, Outdoors: true},
				{Channel: 32, Indoors: false, Outdoors: false},
				{Channel: 33, Indoors: false, Outdoors: false},
				{Channel: 34, Indoors: false, Outdoors: false},
				{Channel: 35, Indoors: true, Outdoors: true},
				{Channel: 36, Indoors: false, Outdoors: false},
				{Channel: 37, Indoors: true, Outdoors: true},
				{Channel: 38, Indoors: false, Outdoors: false},
				{Channel: 39, Indoors: false, Outdoors: false},
				{Channel: 40, Indoors: false, Outdoors: false},
				{Channel: 41, Indoors: true, Outdoors: true},
				{Channel: 42, Indoors: false, Outdoors: false},
				{Channel: 43, Indoors: true, Outdoors: true},
				{Channel: 44, Indoors: true, Outdoors: true},
				{Channel: 45, Indoors: false, Outdoors: false},
				{Channel: 46, Indoors: true, Outdoors: true},
				{Channel: 47, Indoors: true, Outdoors: true},
				{Channel: 48, Indoors: true, Outdoors: true},
				{Channel: 49, Indoors: false, Outdoors: false},
			},
		},
	}

	for _, test := range testCases {
		s := no.Norway{Latitude: test.Latitude, Longitude: test.Longitude}
		channels := *(s.Call())

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