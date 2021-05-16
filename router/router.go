package router

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/stebunting/rfxp-backend/channel"
	"github.com/stebunting/rfxp-backend/external/dk"
	"github.com/stebunting/rfxp-backend/external/gb"
	"github.com/stebunting/rfxp-backend/external/nl"
	"github.com/stebunting/rfxp-backend/external/se"
)

type Request struct {
	Latitude  float64
	Longitude float64
}

type Response struct {
	Status   string
	Details  string
	Location Location
	Channels []channel.Channel
}

type Location struct {
	Latitude  float64
	Longitude float64
}

type Api interface {
	Call() *[]channel.Channel
}

func GetData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	country, err := getCountry(r)
	if err != nil {
		handleError(w, err)
		return
	}

	latitude, err := getLatitude(r)
	if err != nil {
		handleError(w, err)
		return
	}

	longitude, err := getLongitude(r)
	if err != nil {
		handleError(w, err)
		return
	}

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
		handleError(w, errors.New("invalid country code"))
		return
	}

	channelInfo := api.Call()
	ret, err := json.Marshal(Response{
		Status: "OK",
		Location: Location{
			Latitude:  latitude,
			Longitude: longitude,
		},
		Channels: *channelInfo,
	})
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(ret)
}

func handleError(w http.ResponseWriter, err error) {
	j, err := json.Marshal(Response{
		Status:   "Error",
		Details:  err.Error(),
		Channels: nil,
	})
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte(j))
}

func getCountry(r *http.Request) (string, error) {
	country := r.URL.Query()["country"]
	if len(country) == 0 {
		return "", errors.New("missing country parameter")
	}

	return strings.ToLower(country[0]), nil
}

func getLatitude(r *http.Request) (float64, error) {
	lat := r.URL.Query()["lat"]
	if len(lat) == 0 {
		return 0, errors.New("missing lat parameter")
	}

	latitude, err := strconv.ParseFloat(lat[0], 64)
	if err != nil {
		return 0, errors.New("invalid lat")
	}

	return latitude, nil
}

func getLongitude(r *http.Request) (float64, error) {
	lng := r.URL.Query()["lng"]
	if len(lng) == 0 {
		return 0, errors.New("missing lng parameter")
	}

	longitude, err := strconv.ParseFloat(lng[0], 64)
	if err != nil {
		return 0, errors.New("invalid lng")
	}

	return longitude, nil
}
