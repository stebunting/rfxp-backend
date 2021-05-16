package coordinates

import (
	"errors"
	"strings"
)

type coordinates struct {
	latitude  float64
	longitude float64
	height    float64
	ellipsoid ellipsoid
}

func New(latitude float64, longitude float64) coordinates {
	g := coordinates{
		latitude:  latitude,
		longitude: longitude,
		height:    0,
		ellipsoid: WGS84,
	}
	return g
}

func NewFromDegrees(
	latDegrees int,
	latMinutes int,
	latSeconds float64,
	latDirection string,
	lngDegrees int,
	lngMinutes int,
	lngSeconds float64,
	lngDirection string,
) (coordinates, error) {
	latitude, err := toDecimal(latDegrees, latMinutes, latSeconds, latDirection)
	if err != nil {
		return coordinates{}, err
	}
	longitude, err := toDecimal(lngDegrees, lngMinutes, lngSeconds, lngDirection)
	if err != nil {
		return coordinates{}, err
	}
	return New(latitude, longitude), nil
}

func (s *coordinates) GetGridReference(system string) (gridReference, error) {
	system = strings.ToUpper(system)
	switch system {
	case "GB":
		return s.transform(NationalGrid, "GB"), nil
	case "IE":
		return s.transform(IrishNationalGrid, "IE"), nil
	case "NL":
		return s.getUTMWithZone(32), nil
	case "UTM":
		return s.GetUTM(), nil
	default:
		return gridReference{}, errors.New("invalid system")
	}
}

func (s *coordinates) GetUTM() gridReference {
	return s.getUTMWithZone(1 + (int)((s.longitude+180)/6))
}

func (s *coordinates) getUTMWithZone(zone int) gridReference {
	var northernHemisphere bool
	var datum datum
	if s.latitude >= 0 {
		northernHemisphere = true
		datum = UtmNorth
	} else {
		northernHemisphere = false
		datum = UtmSouth
	}

	eastingNorthing := NewEastingsNorthings(s.latitude, s.longitude+float64((30-zone)*6), datum)
	gridReference := NewGridReference(s, eastingNorthing, "UTM", zone, northernHemisphere)
	return gridReference
}

func (s *coordinates) transform(datum datum, system string) gridReference {
	cartesian := NewCartesian(s)
	lat, lon, _ := cartesian.transform(datum)
	eastingNorthing := NewEastingsNorthings(lat, lon, datum)
	gridReference := NewGridReference(s, eastingNorthing, "GB", 1, true)
	return gridReference
}
