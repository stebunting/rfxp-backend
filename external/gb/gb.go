package gb

import (
	"errors"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"sort"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/stebunting/rfxp-backend/channel"
	"github.com/stebunting/rfxp-backend/coordinates"
)

type GB struct {
	Latitude  float64
	Longitude float64
	Code      string
	client    *http.Client
	url       *url.URL
	form      url.Values
}

func (s *GB) Call() *[]channel.Channel {
	lookup := coordinates.New(s.Latitude, s.Longitude)
	gridReference, _ := lookup.GetGridReference(s.Code)
	s.setupClient(gridReference.GetShortCode())

	err := s.initSession()
	if err != nil {
		// HANDLE ERROR
	}
	err = s.getLocationList()
	if err != nil {
		// HANDLE ERROR
	}
	channels, err := s.getData()
	if err != nil {
		// HANDLE ERROR
	}

	return channels
}

func (s *GB) setupClient(code string) {
	jar, err := cookiejar.New(nil)
	if err != nil {
		panic(err)
	}
	s.client = &http.Client{
		Jar: jar,
	}

	s.url, err = url.Parse("https://pmse.ofcom.org.uk/Pmse/wireless/public/microphone700.aspx")
	if err != nil {
		panic(err)
	}

	s.form = url.Values{}
	s.form.Add("ctl00$mcph$Where$Square", code[:2])
	s.form.Add("ctl00$mcph$Where$Easting", code[2:5])
	s.form.Add("ctl00$mcph$Where$Northing", code[5:])
	s.form.Add("ctl00$mcph$Where$btnSearch", "Find >")
	s.form.Add("ctl00$mcph$Where$LocationGroup", "radLocation2")
	s.form.Add("ctl00$mcph$Where$LocationList", "0")
}

func (s *GB) initSession() error {
	document, err := s.getDocument()
	if err != nil {
		return err
	}

	var exists bool
	viewState, exists := document.Find("#__VIEWSTATE").Attr("value")
	if !exists {
		return errors.New("no viewstate element")
	}
	s.form.Add("__VIEWSTATE", viewState)

	return nil
}

func (s *GB) getLocationList() error {
	document, err := s.getDocument()
	if err != nil {
		return err
	}

	locationConfirmed := false
	document.Find("#ctl00_mcph_Where_LocationList option").Each(func(i int, s *goquery.Selection) {
		value, exists := s.Attr("value")
		if exists && value == "0" {
			locationConfirmed = true
		}
	})
	if !locationConfirmed {
		return errors.New("location not found")
	}

	return nil
}

func (s *GB) getData() (*[]channel.Channel, error) {
	document, err := s.getDocument()
	if err != nil {
		return nil, err
	}

	indoorsThreshold := 3
	startFrequency := 470000
	startChannel := 21
	chWidth := 8000
	channels := []channel.Channel{{
		Number:    38,
		FreqStart: 606000,
		FreqEnd:   614000,
		Indoors:   true,
		Outdoors:  true,
	}}
	document.Find("#ctl00_mcph_rptMicrophoneDSO tbody tr").Each(func(i int, sel *goquery.Selection) {
		ch, err := strconv.Atoi(sel.Find("td").Eq(0).Text())
		if err == nil {
			freqStart := ((ch - startChannel) * chWidth) + startFrequency
			freqEnd := freqStart + chWidth

			var indoors bool
			inImg, exists := sel.Find("td div img").Eq(0).Attr("src")
			if exists {
				inImgSplit := strings.Split(inImg, "/")
				inImg = inImgSplit[len(inImgSplit)-1]
				indoorsQuality := int(inImg[12]) - 48
				if indoorsQuality >= indoorsThreshold {
					indoors = true
				} else {
					indoors = false
				}
			} else {
				indoors = true
			}

			out := sel.Find("td div span").Eq(1).Text()
			var outdoors bool
			if out == "Not available" {
				outdoors = false
			} else {
				outdoors = true
			}

			channels = append(channels, channel.Channel{
				Number:    ch,
				FreqStart: freqStart,
				FreqEnd:   freqEnd,
				Indoors:   indoors,
				Outdoors:  outdoors,
			})
		}
	})

	sort.Slice(channels, func(i, j int) bool {
		return channels[i].FreqStart < channels[j].FreqStart
	})

	return &channels, nil
}

func (s *GB) getDocument() (*goquery.Document, error) {
	request, err := http.NewRequest(http.MethodPost, s.url.String(), strings.NewReader(s.form.Encode()))
	if err != nil {
		panic(err)
	}
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	rawResponse, err := s.client.Do(request)
	if err != nil {
		return nil, errors.New("error making network call")
	}

	document, err := goquery.NewDocumentFromReader(rawResponse.Body)
	if err != nil {
		panic(err)
	}
	defer rawResponse.Body.Close()

	return document, nil
}
