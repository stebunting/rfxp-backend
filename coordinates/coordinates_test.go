package coordinates_test

import (
	"testing"

	"github.com/stebunting/rfxp-backend/coordinates"
)

func TestGbCoordinates(t *testing.T) {
	EastingsLowThreshold := 0.6  // metres
	EastingsHighThreshold := 2.8 // metres
	NorthingsThreshold := 0.1    // metres

	type TestCases struct {
		Name      string
		Lat       float64
		Lng       float64
		Easting   float64
		Northing  float64
		Code      string
		ShortCode string
	}
	testCases := []TestCases{
		{
			Name:      "The Lizard",
			Lat:       49.97454006765309,
			Lng:       -5.212325001930045,
			Easting:   169770,
			Northing:  13209,
			Code:      "SW6977013209",
			ShortCode: "SW697132",
		}, {
			Name:      "Wimbledon",
			Lat:       51.42761719993095,
			Lng:       -0.1908007959012176,
			Easting:   525876,
			Northing:  171398,
			Code:      "TQ2587571398", // TQ2587671398
			ShortCode: "TQ258713",
		}, {
			Name:      "Shetland Islands",
			Lat:       60.419662930284225,
			Lng:       -1.3939871714693122,
			Easting:   433474,
			Northing:  1170799,
			Code:      "HU3347370799", // HU3347470799
			ShortCode: "HU334707",
		}, {
			Name:      "Aberdeen",
			Lat:       57.152853710141585,
			Lng:       -2.1124450090652207,
			Easting:   393296,
			Northing:  806892,
			Code:      "NJ9329506892", // NJ9329606892
			ShortCode: "NJ932068",
		}, {
			Name:      "Isle Of Man",
			Lat:       54.317958131191475,
			Lng:       -4.384776130138725,
			Easting:   244980,
			Northing:  493998,
			Code:      "SC4498093998",
			ShortCode: "SC449939",
		}, {
			Name:      "Pembroke",
			Lat:       51.696604453740996,
			Lng:       -4.940054429431384,
			Easting:   196913,
			Northing:  203855,
			Code:      "SM9691303855",
			ShortCode: "SM969038",
		}, {
			Name:      "Folkestone",
			Lat:       51.09322953304597,
			Lng:       1.1018199180798527,
			Easting:   617298,
			Northing:  137234,
			Code:      "TR1729737234", // TR1729837234
			ShortCode: "TR172372",
		}, {
			Name:      "Isle Of Wight",
			Lat:       50.65911346559466,
			Lng:       -1.2544484482073506,
			Easting:   452795,
			Northing:  84647,
			Code:      "SZ5279484647", // SZ5279584647
			ShortCode: "SZ527846",
		}, {
			Name:      "Isles Of Scilly",
			Lat:       49.918089161149695,
			Lng:       -6.298469138464261,
			Easting:   91552,
			Northing:  10849,
			Code:      "SV9155410849", // SV9155210849
			ShortCode: "SV915108",
		}, {
			Name:      "Western Isles Of Scilly", // Maximum Error on Eastings
			Lat:       49.949269237524284,
			Lng:       -6.355019079981123,
			Easting:   87696,
			Northing:  14549,
			Code:      "SV8769814549", // SV8769614549
			ShortCode: "SV876145",
		},
	}

	for _, test := range testCases {
		lookup := coordinates.New(test.Lat, test.Lng)
		gridReference, err := lookup.GetGridReference("GB")
		if err != nil {
			t.Fatalf("Coordinates unexpectedly errored: %s", err.Error())
		}
		if gridReference.GetEasting() < test.Easting-EastingsLowThreshold || gridReference.GetEasting() > test.Easting+EastingsHighThreshold {
			t.Fatalf("\n--- Incorrect Easting ---\n    NAME: %s\n     GOT: %f\nEXPECTED: %f", test.Name, gridReference.GetEasting(), test.Easting)
		}
		if gridReference.GetNorthing() < test.Northing-NorthingsThreshold || gridReference.GetNorthing() > test.Northing+NorthingsThreshold {
			t.Fatalf("\n--- Incorrect Northing ---\n    NAME: %s\n     GOT: %f\nEXPECTED: %f", test.Name, gridReference.GetNorthing(), test.Northing)
		}
		if gridReference.GetCode() != test.Code {
			t.Fatalf("\n--- Incorrect Code ---\n    NAME: %s\n     GOT: %s\nEXPECTED: %s", test.Name, gridReference.GetCode(), test.Code)
		}
		if gridReference.GetShortCode() != test.ShortCode {
			t.Fatalf("\n--- Incorrect ShortCode ---\n    NAME: %s\n     GOT: %s\nEXPECTED: %s", test.Name, gridReference.GetShortCode(), test.ShortCode)
		}
	}
}

func TestIeCoordinates(t *testing.T) {
	EastingsLowThreshold := 0.6  // metres
	EastingsHighThreshold := 1.0 // metres
	NorthingsThreshold := 0.4    // metres

	type TestCases struct {
		Name      string
		Lat       float64
		Lng       float64
		Easting   float64
		Northing  float64
		Code      string
		ShortCode string
	}
	testCases := []TestCases{
		{
			Name:      "Belfast",
			Lat:       54.596048,
			Lng:       -5.930201,
			Easting:   333825,
			Northing:  373948,
			Code:      "IJ3382573948",
			ShortCode: "IJ338739",
		}, {
			Name:      "Londonderry",
			Lat:       55.007925,
			Lng:       -7.325037,
			Easting:   243234,
			Northing:  418038,
			Code:      "IC4323418037",
			ShortCode: "IC432180",
		}, {
			Name:      "Enniskillen",
			Lat:       54.138185,
			Lng:       -7.352331,
			Easting:   242381,
			Northing:  321204,
			Code:      "IH4238021204",
			ShortCode: "IH423212",
		}, {
			Name:      "Ballycastle",
			Lat:       55.202954,
			Lng:       -6.234729,
			Easting:   312442,
			Northing:  440964,
			Code:      "ID1244140964",
			ShortCode: "ID124409",
		},
	}

	for _, test := range testCases {
		lookup := coordinates.New(test.Lat, test.Lng)
		gridReference, err := lookup.GetGridReference("IE")
		if err != nil {
			t.Fatalf("Coordinates unexpectedly errored: %s", err.Error())
		}
		if gridReference.GetEasting() < test.Easting-EastingsLowThreshold || gridReference.GetEasting() > test.Easting+EastingsHighThreshold {
			t.Fatalf("\n--- Incorrect Easting ---\n    NAME: %s\n     GOT: %f\nEXPECTED: %f", test.Name, gridReference.GetEasting(), test.Easting)
		}
		if gridReference.GetNorthing() < test.Northing-NorthingsThreshold || gridReference.GetNorthing() > test.Northing+NorthingsThreshold {
			t.Fatalf("\n--- Incorrect Northing ---\n    NAME: %s\n     GOT: %f\nEXPECTED: %f", test.Name, gridReference.GetNorthing(), test.Northing)
		}
		if gridReference.GetCode() != test.Code {
			t.Fatalf("\n--- Incorrect Code ---\n    NAME: %s\n     GOT: %s\nEXPECTED: %s", test.Name, gridReference.GetCode(), test.Code)
		}
		if gridReference.GetShortCode() != test.ShortCode {
			t.Fatalf("\n--- Incorrect ShortCode ---\n    NAME: %s\n     GOT: %s\nEXPECTED: %s", test.Name, gridReference.GetShortCode(), test.ShortCode)
		}
	}
}

func TestChannelIslesGridReference(t *testing.T) {
	EastingsLowThreshold := 1.0  // metres
	EastingsHighThreshold := 1.0 // metres
	NorthingsThreshold := 1.0    // metres

	type TestCases struct {
		Name      string
		Lat       float64
		Lng       float64
		Easting   float64
		Northing  float64
		Code      string
		ShortCode string
	}
	testCases := []TestCases{
		{
			Name:      "St. Helier, Jersey",
			Lat:       49.179435,
			Lng:       -2.087105,
			Easting:   566530,
			Northing:  5447804,
			Code:      "CJ6653047804",
			ShortCode: "CJ665478",
		}, {
			Name:      "Saint Ouen, Jersey",
			Lat:       49.252992,
			Lng:       -2.239540,
			Easting:   555339,
			Northing:  5455858,
			Code:      "CJ5533955858",
			ShortCode: "CJ553558",
		}, {
			Name:      "Torteval, Guernsey",
			Lat:       49.433080,
			Lng:       -2.658784,
			Easting:   524740,
			Northing:  5475657,
			Code:      "CJ2474075657",
			ShortCode: "CJ247756",
		}, {
			Name:      "St Anne, Alderney",
			Lat:       49.714151,
			Lng:       -2.197707,
			Easting:   557837,
			Northing:  5507158,
			Code:      "CA5783707158",
			ShortCode: "CA578071",
		},
	}

	for _, test := range testCases {
		lookup := coordinates.New(test.Lat, test.Lng)
		gridReference, err := lookup.GetGridReference("UTM")
		if err != nil {
			t.Fatalf("Coordinates unexpectedly errored: %s", err.Error())
		}
		if gridReference.GetEasting() < test.Easting-EastingsLowThreshold || gridReference.GetEasting() > test.Easting+EastingsHighThreshold {
			t.Fatalf("\n--- Incorrect Easting ---\n    NAME: %s\n     GOT: %f\nEXPECTED: %f", test.Name, gridReference.GetEasting(), test.Easting)
		}
		if gridReference.GetNorthing() < test.Northing-NorthingsThreshold || gridReference.GetNorthing() > test.Northing+NorthingsThreshold {
			t.Fatalf("\n--- Incorrect Northing ---\n    NAME: %s\n     GOT: %f\nEXPECTED: %f", test.Name, gridReference.GetNorthing(), test.Northing)
		}
		if gridReference.GetCode() != test.Code {
			t.Fatalf("\n--- Incorrect Code ---\n    NAME: %s\n     GOT: %s\nEXPECTED: %s", test.Name, gridReference.GetCode(), test.Code)
		}
		if gridReference.GetShortCode() != test.ShortCode {
			t.Fatalf("\n--- Incorrect ShortCode ---\n    NAME: %s\n     GOT: %s\nEXPECTED: %s", test.Name, gridReference.GetShortCode(), test.ShortCode)
		}
	}
}

func TestUtmGridReference(t *testing.T) {
	EastingsLowThreshold := 1.0  // metres
	EastingsHighThreshold := 1.0 // metres
	NorthingsThreshold := 1.0    // metres

	type TestCases struct {
		Name               string
		Lat                float64
		Lng                float64
		Easting            float64
		Northing           float64
		Zone               int
		NorthernHemisphere bool
	}
	testCases := []TestCases{
		{
			Name:               "Groningen",
			Lat:                53.21484,
			Lng:                6.569683,
			Easting:            337725,
			Northing:           5898927,
			Zone:               32,
			NorthernHemisphere: true,
		}, {
			Name:               "Rotterdam",
			Lat:                51.920239,
			Lng:                4.450462,
			Easting:            599749,
			Northing:           5753160,
			Zone:               31,
			NorthernHemisphere: true,
		}, {
			Name:               "Venlo",
			Lat:                51.393114,
			Lng:                6.179330,
			Easting:            303769,
			Northing:           5697318,
			Zone:               32,
			NorthernHemisphere: true,
		},
	}

	for _, test := range testCases {
		lookup := coordinates.New(test.Lat, test.Lng)
		gridReference := lookup.GetUTM()
		if gridReference.GetEasting() < test.Easting-EastingsLowThreshold || gridReference.GetEasting() > test.Easting+EastingsHighThreshold {
			t.Fatalf("\n--- Incorrect Easting ---\n    NAME: %s\n     GOT: %f\nEXPECTED: %f", test.Name, gridReference.GetEasting(), test.Easting)
		}
		if gridReference.GetNorthing() < test.Northing-NorthingsThreshold || gridReference.GetNorthing() > test.Northing+NorthingsThreshold {
			t.Fatalf("\n--- Incorrect Northing ---\n    NAME: %s\n     GOT: %f\nEXPECTED: %f", test.Name, gridReference.GetNorthing(), test.Northing)
		}
		if gridReference.GetZone() != test.Zone {
			t.Fatalf("\n--- Incorrect Zone ---\n    NAME: %s\n     GOT: %f\nEXPECTED: %f", test.Name, gridReference.GetNorthing(), test.Northing)
		}
	}
}

func TestWorldwideUtmGridReference(t *testing.T) {
	EastingsThreshold := 1.0  // metres
	NorthingsThreshold := 1.0 // metres

	type TestCases struct {
		Name               string
		Lat                float64
		Lng                float64
		Easting            float64
		Northing           float64
		Zone               int
		NorthernHemisphere bool
	}
	testCases := []TestCases{
		{
			Name:               "Kansas City, USA",
			Lat:                38.627346882,
			Lng:                -95.3723847,
			Easting:            293485,
			Northing:           4278094,
			Zone:               15,
			NorthernHemisphere: true,
		}, {
			Name:               "State of Rio Grande Do Sul, Brazil",
			Lat:                -28.0123802749,
			Lng:                -54.12837482,
			Easting:            782380,
			Northing:           6898101,
			Zone:               21,
			NorthernHemisphere: false,
		}, {
			Name:               "Pyongyang, North Korea",
			Lat:                39.1738927,
			Lng:                125.723542987,
			Easting:            735276,
			Northing:           4339608,
			Zone:               51,
			NorthernHemisphere: true,
		},
	}

	for _, test := range testCases {
		lookup := coordinates.New(test.Lat, test.Lng)
		gridReference := lookup.GetUTM()
		if gridReference.GetEasting() < test.Easting-EastingsThreshold || gridReference.GetEasting() > test.Easting+EastingsThreshold {
			t.Fatalf("\n--- Incorrect Easting ---\n    NAME: %s\n     GOT: %f\nEXPECTED: %f", test.Name, gridReference.GetEasting(), test.Easting)
		}
		if gridReference.GetNorthing() < test.Northing-NorthingsThreshold || gridReference.GetNorthing() > test.Northing+NorthingsThreshold {
			t.Fatalf("\n--- Incorrect Northing ---\n    NAME: %s\n     GOT: %f\nEXPECTED: %f", test.Name, gridReference.GetNorthing(), test.Northing)
		}
		if gridReference.GetZone() != test.Zone {
			t.Fatalf("\n--- Incorrect Zone ---\n    NAME: %s\n     GOT: %f\nEXPECTED: %f", test.Name, gridReference.GetNorthing(), test.Northing)
		}
	}
}

func TestWorldwideUtmGridReferenceFromDegrees(t *testing.T) {
	EastingsThreshold := 1.0  // metres
	NorthingsThreshold := 1.0 // metres

	type TestCases struct {
		Name               string
		LatDegrees         int
		LatMinutes         int
		LatSeconds         float64
		LatDirection       string
		LngDegrees         int
		LngMinutes         int
		LngSeconds         float64
		LngDirection       string
		Easting            float64
		Northing           float64
		Zone               int
		NorthernHemisphere bool
	}
	testCases := []TestCases{
		{
			Name:               "Sydney, Australia",
			LatDegrees:         32,
			LatMinutes:         39,
			LatSeconds:         54.1237,
			LatDirection:       "S",
			LngDegrees:         151,
			LngMinutes:         22,
			LngSeconds:         11.127897,
			LngDirection:       "E",
			Easting:            347128,
			Northing:           6384672,
			Zone:               56,
			NorthernHemisphere: false,
		}, {
			Name:               "Cote D'Ivoire",
			LatDegrees:         6,
			LatMinutes:         11,
			LatSeconds:         5.21378,
			LatDirection:       "n",
			LngDegrees:         4,
			LngMinutes:         25,
			LngSeconds:         27.237894,
			LngDirection:       "w",
			Easting:            342419,
			Northing:           683842,
			Zone:               30,
			NorthernHemisphere: true,
		},
	}

	for _, test := range testCases {
		lookup, err := coordinates.NewFromDegrees(test.LatDegrees, test.LatMinutes, test.LatSeconds, test.LatDirection, test.LngDegrees, test.LngMinutes, test.LngSeconds, test.LngDirection)
		if err != nil {
			t.Fatalf("Coordinates unexpectedly errored: %s", err.Error())
		}
		gridReference := lookup.GetUTM()
		if gridReference.GetEasting() < test.Easting-EastingsThreshold || gridReference.GetEasting() > test.Easting+EastingsThreshold {
			t.Fatalf("\n--- Incorrect Easting ---\n    NAME: %s\n     GOT: %f\nEXPECTED: %f", test.Name, gridReference.GetEasting(), test.Easting)
		}
		if gridReference.GetNorthing() < test.Northing-NorthingsThreshold || gridReference.GetNorthing() > test.Northing+NorthingsThreshold {
			t.Fatalf("\n--- Incorrect Northing ---\n    NAME: %s\n     GOT: %f\nEXPECTED: %f", test.Name, gridReference.GetNorthing(), test.Northing)
		}
		if gridReference.GetZone() != test.Zone {
			t.Fatalf("\n--- Incorrect Zone ---\n    NAME: %s\n     GOT: %f\nEXPECTED: %f", test.Name, gridReference.GetNorthing(), test.Northing)
		}
	}
}
