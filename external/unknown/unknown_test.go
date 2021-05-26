package unknown_test

import (
	"log"
	"testing"

	"github.com/stebunting/rfxp-backend/external/unknown"
)

func TestUnknown(t *testing.T) {
	s := &unknown.Unknown{}

	c, err := s.Call()
	if err != nil {
		log.Fatalf("unexpected error making network call")
	}

	channels := *c
	if len(channels) > 0 {
		log.Fatalf("unexpected entry in channels slice")
	}
}

func TestName(t *testing.T) {
	s := &unknown.Unknown{}
	name := s.GetCountryName()
	if name != "Unknown" {
		log.Fatalf("got wrong country name")
	}
}

func TestService(t *testing.T) {
	s := &unknown.Unknown{}
	name := s.GetServiceName()
	if name != "Unknown" {
		log.Fatalf("got wrong service name")
	}
}
