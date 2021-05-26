package nl

import (
	"errors"
	"image"
	_ "image/png"
	"io/ioutil"
	"math"
	"net/http"

	"github.com/getsentry/sentry-go"
	"github.com/stebunting/rfxp-backend/channel"
	"github.com/stebunting/rfxp-backend/coordinates"
)

type Netherlands struct {
	Latitude  float64
	Longitude float64
}

func (s *Netherlands) GetCountryName() string {
	return "The Netherlands"
}

func (s *Netherlands) GetServiceName() string {
	return "Microfoonbanden.nl"
}

func (s *Netherlands) Call() (*[]channel.Channel, error) {
	lookup := coordinates.New(s.Latitude, s.Longitude)
	gridReference, _ := lookup.GetGridReference("NL")

	channels, err := s.makeApiCall(gridReference.GetEasting(), gridReference.GetNorthing())
	if err != nil {
		return nil, err
	}

	return channels, nil
}

func (s *Netherlands) makeApiCall(easting float64, northing float64) (*[]channel.Channel, error) {
	response, err := http.Get("https://www.microfoonbanden.nl/images/inputData.png")
	if err != nil {
		sentry.CaptureException(err)
		panic(err)
	}
	defer response.Body.Close()

	reader := ioutil.NopCloser(response.Body)
	img, _, err := image.Decode(reader)
	if err != nil {
		sentry.CaptureException(err)
		panic(err)
	}

	minX := img.Bounds().Min.X
	minY := img.Bounds().Min.Y

	numXPixels := 585
	numYPixels := 699

	coordLBLong := 97467.188 - 82.043
	coordLBLat := 5947931.25 - 206.161
	coordRBLong := 389962.5 - 82.043
	coordRBLat := 5947931.25 - 206.161
	coordLOLong := 97467.188 - 82.043
	coordLOLat := 5598673.438 - 206.161
	coordROLong := 389962.5 - 82.043
	coordROLat := 5598673.438 - 206.161

	if northing > coordLBLat || northing < coordLOLat || easting > coordRBLong || easting < coordLBLong {
		return nil, errors.New("coordinates outside NL")
	}

	lengthLBRB := math.Sqrt(math.Pow(coordRBLong-coordLBLong, 2) + math.Pow(coordRBLat-coordLBLat, 2))
	lengthLORO := math.Sqrt(math.Pow(coordROLong-coordLOLong, 2) + math.Pow(coordROLat-coordLOLat, 2))
	lengthLBLO := math.Sqrt(math.Pow(coordLOLong-coordLBLong, 2) + math.Pow(coordLOLat-coordLBLat, 2))
	lengthRBRO := math.Sqrt(math.Pow(coordRBLong-coordROLong, 2) + math.Pow(coordROLat-coordRBLat, 2))

	deltaLBTargetLong := easting - coordLBLong
	deltaLOTargetLong := easting - coordLOLong
	deltaRBTargetLong := coordRBLong - easting
	deltaROTargetLong := coordROLong - easting
	deltaLBTargetLat := coordLBLat - northing
	deltaLOTargetLat := northing - coordLOLat
	deltaRBTargetLat := coordRBLat - northing
	deltaROTargetLat := northing - coordROLat

	factorLong := ((easting-coordLBLong)/lengthLBRB + (easting-coordLOLong)/lengthLORO) / 2
	factorLat := ((coordLBLat-northing)/lengthLBLO + (coordRBLat-northing)/lengthRBRO) / 2

	xPixelPos := int(math.Round(((deltaLBTargetLong/lengthLBRB*(1-factorLat)+(deltaLOTargetLong/lengthLORO)*factorLat)*(1-factorLong) + ((1-deltaRBTargetLong/lengthLBRB)*(1-factorLat)+(1-deltaROTargetLong/lengthLORO)*factorLat)*factorLong) * float64(numXPixels)))
	yPixelPos := int(math.Round(((deltaLBTargetLat/lengthLBLO*(1-factorLong)+(deltaRBTargetLat/lengthRBRO)*factorLong)*(1-factorLat) + ((1-deltaLOTargetLat/lengthLBLO)*(1-factorLong)+(1-deltaROTargetLat/lengthRBRO)*factorLong)*factorLat) * float64(numYPixels)))

	availability := s.initChannels()

	for k := 0; k < 2; k++ {
		xPixelPosMap := minX + xPixelPos + numXPixels*k
		for i := 0; i < 4; i++ {
			yPixelPosMap := minY + yPixelPos + (numYPixels * i)

			r, _, _, _ := img.At(xPixelPosMap, yPixelPosMap).RGBA()
			r = r / 257

			for j := 0; j < 8; j++ {
				index := j + (i * 8)
				if index >= len(availability) {
					break
				}
				if k == 0 {
					availability[index].Outdoors = r%2 == 1
				} else {
					availability[index].Indoors = r%2 == 1
				}
				r = r >> 1
			}
		}
	}

	return &availability, nil
}

func (s *Netherlands) initChannels() []channel.Channel {
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
