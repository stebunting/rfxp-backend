package unknown

import (
	"github.com/stebunting/rfxp-backend/channel"
)

type Unknown struct {
}

func (s *Unknown) GetCountryName() string {
	return "Unknown"
}

func (s *Unknown) GetServiceName() string {
	return "Unknown"
}

func (s *Unknown) Call() (*[]channel.Channel, error) {
	channels := []channel.Channel{}
	return &channels, nil
}
