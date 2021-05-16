package router

import (
	"context"
	"errors"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/joho/godotenv"
	"github.com/stebunting/rfxp-backend/channel"
	"github.com/stebunting/rfxp-backend/external/dk"
	"github.com/stebunting/rfxp-backend/external/gb"
	"github.com/stebunting/rfxp-backend/external/nl"
	"github.com/stebunting/rfxp-backend/external/se"
)

type LambdaRequest struct {
	Country   string `json:"country"`
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
}

type Response struct {
	Status   string            `json:"status"`
	Location Location          `json:"location"`
	Channels []channel.Channel `json:"channels"`
}

type Location struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type Api interface {
	Call() *[]channel.Channel
}

func init() {
	godotenv.Load()
}

func HandleLambdaEvent(ctx context.Context, r LambdaRequest) (Response, error) {
	err := sentry.Init(sentry.ClientOptions{
		Dsn:         os.Getenv("SENTRY_DSN"),
		Environment: os.Getenv("SENTRY_ENV"),
	})
	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}
	defer sentry.Flush(2 * time.Second)

	latitude, err := strconv.ParseFloat(r.Latitude, 64)
	if err != nil {
		return Response{}, errors.New("latitude must be a number")
	}
	if latitude < -180 || latitude > 180 {
		return Response{}, errors.New("latitude must be between -180 and 180 degrees")
	}

	longitude, err := strconv.ParseFloat(r.Longitude, 64)
	if err != nil {
		return Response{}, errors.New("longitude must be a number")
	}
	if longitude < -60 || longitude > 80 {
		return Response{}, errors.New("longitude must be between -60 and 80 degrees")
	}

	country := r.Country

	var api Api
	switch country {
	case "se":
		api = &se.Sweden{Latitude: latitude, Longitude: longitude}
	case "dk":
		api = &dk.Denmark{Latitude: latitude, Longitude: longitude}
	case "nl":
		api = &nl.Netherlands{Latitude: latitude, Longitude: longitude}
	case "gb", "im":
		api = &gb.GB{Latitude: latitude, Longitude: longitude, Code: "GB"}
	case "je", "gg":
		api = &gb.GB{Latitude: latitude, Longitude: longitude, Code: "UTM"}
	default:
		return Response{}, errors.New("invalid country code or country not implemented")
	}

	channelInfo := api.Call()
	return Response{
		Status: "OK",
		Location: Location{
			Latitude:  latitude,
			Longitude: longitude,
		},
		Channels: *channelInfo,
	}, nil
}
