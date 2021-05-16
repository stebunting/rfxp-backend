package coordinates

import "fmt"

type gridReference struct {
	latitude           float64
	longitude          float64
	easting            float64
	northing           float64
	gridSystem         string
	zone               int
	northernHemisphere bool
	code               string
	shortCode          string
}

func NewGridReference(
	coordinates *coordinates,
	eastingNorthing eastingNorthing,
	gridSystem string,
	zone int,
	northernHemisphere bool,
) gridReference {
	g := gridReference{
		latitude:           coordinates.latitude,
		longitude:          coordinates.longitude,
		easting:            eastingNorthing.easting,
		northing:           eastingNorthing.northing,
		zone:               zone,
		northernHemisphere: northernHemisphere,
		gridSystem:         gridSystem,
	}

	g.setGridReference()

	return g
}

func (s *gridReference) setGridReference() {
	eastingStr := fmt.Sprintf("%08d", int(s.easting))
	northingStr := fmt.Sprintf("%08d", int(s.northing))
	code := fmt.Sprintf("%s%s", eastingStr[0:3], northingStr[0:3])

	var gridSquare string
	if s.gridSystem == "GB" {
		gridSquare = s.getGbCode(code)
	} else if s.gridSystem == "IE" {
		gridSquare = s.getIeCode(code)
	} else if s.gridSystem == "UTM" && s.northernHemisphere && s.zone == 30 {
		gridSquare = s.getChannelIsleCode(code)
	}

	if gridSquare == "" {
		return
	}

	s.code = fmt.Sprintf("%s%s%s", gridSquare, eastingStr[3:], northingStr[3:])
	s.shortCode = fmt.Sprintf("%s%s", s.code[:5], s.code[7:10])
}

func (s *gridReference) getGbCode(code string) string {
	switch code {
	case "000012":
		return "HL"
	case "001012":
		return "HM"
	case "002012":
		return "HN"
	case "003012":
		return "HO"
	case "004012":
		return "HP"
	case "005012":
		return "JL"
	case "006012":
		return "JM"
	case "007012":
		return "JN"

	case "000011":
		return "HQ"
	case "001011":
		return "HR"
	case "002011":
		return "HS"
	case "003011":
		return "HT"
	case "004011":
		return "HU"
	case "005011":
		return "JQ"
	case "006011":
		return "JR"
	case "007011":
		return "JS"

	case "000010":
		return "HV"
	case "001010":
		return "HW"
	case "002010":
		return "HX"
	case "003010":
		return "HY"
	case "004010":
		return "HZ"
	case "005010":
		return "JV"
	case "006010":
		return "JW"
	case "007010":
		return "JX"

	case "000009":
		return "NA"
	case "001009":
		return "NB"
	case "002009":
		return "NC"
	case "003009":
		return "ND"
	case "004009":
		return "NE"
	case "005009":
		return "OA"
	case "006009":
		return "OB"
	case "007009":
		return "OC"

	case "000008":
		return "NF"
	case "001008":
		return "NG"
	case "002008":
		return "NH"
	case "003008":
		return "NJ"
	case "004008":
		return "NK"
	case "005008":
		return "OF"
	case "006008":
		return "OG"
	case "007008":
		return "OH"

	case "000007":
		return "NL"
	case "001007":
		return "NM"
	case "002007":
		return "NN"
	case "003007":
		return "NO"
	case "004007":
		return "NP"
	case "005007":
		return "OL"
	case "006007":
		return "OM"
	case "007007":
		return "ON"

	case "000006":
		return "NQ"
	case "001006":
		return "NR"
	case "002006":
		return "NS"
	case "003006":
		return "NT"
	case "004006":
		return "NU"
	case "005006":
		return "OQ"
	case "006006":
		return "OR"
	case "007006":
		return "OS"

	case "000005":
		return "NV"
	case "001005":
		return "NW"
	case "002005":
		return "NX"
	case "003005":
		return "NY"
	case "004005":
		return "NZ"
	case "005005":
		return "OV"
	case "006005":
		return "OW"
	case "007005":
		return "OX"

	case "000004":
		return "SA"
	case "001004":
		return "SB"
	case "002004":
		return "SC"
	case "003004":
		return "SD"
	case "004004":
		return "SE"
	case "005004":
		return "TA"
	case "006004":
		return "TB"
	case "007004":
		return "TC"

	case "000003":
		return "SF"
	case "001003":
		return "SG"
	case "002003":
		return "SH"
	case "003003":
		return "SJ"
	case "004003":
		return "SK"
	case "005003":
		return "TF"
	case "006003":
		return "TG"
	case "007003":
		return "TH"

	case "000002":
		return "SL"
	case "001002":
		return "SM"
	case "002002":
		return "SN"
	case "003002":
		return "SO"
	case "004002":
		return "SP"
	case "005002":
		return "TL"
	case "006002":
		return "TM"
	case "007002":
		return "TN"

	case "000001":
		return "SQ"
	case "001001":
		return "SR"
	case "002001":
		return "SS"
	case "003001":
		return "ST"
	case "004001":
		return "SU"
	case "005001":
		return "TQ"
	case "006001":
		return "TR"
	case "007001":
		return "TS"

	case "000000":
		return "SV"
	case "001000":
		return "SW"
	case "002000":
		return "SX"
	case "003000":
		return "SY"
	case "004000":
		return "SZ"
	case "005000":
		return "TV"
	case "006000":
		return "TW"
	case "007000":
		return "TX"

	default:
		return ""
	}
}

func (s *gridReference) getIeCode(code string) string {
	switch code {
	case "000004":
		return "A"
	case "001004":
		return "B"
	case "002004":
		return "C"
	case "003004":
		return "D"
	case "004004":
		return "E"
	case "000003":
		return "F"
	case "001003":
		return "G"
	case "002003":
		return "H"
	case "003003":
		return "J"
	case "004003":
		return "K"
	case "000002":
		return "L"
	case "001002":
		return "M"
	case "002002":
		return "N"
	case "003002":
		return "O"
	case "004002":
		return "P"
	case "000001":
		return "Q"
	case "001001":
		return "R"
	case "002001":
		return "S"
	case "003001":
		return "T"
	case "004001":
		return "U"
	case "000000":
		return "V"
	case "001000":
		return "W"
	case "002000":
		return "X"
	case "003000":
		return "Y"
	case "004000":
		return "Z"
	default:
		return ""
	}
}

func (s *gridReference) getChannelIsleCode(code string) string {
	switch code {
	case "005054":
		return "CJ"
	case "005055":
		return "CA"
	default:
		return ""
	}
}

func (s *gridReference) GetCode() string {
	return s.code
}

func (s *gridReference) GetShortCode() string {
	return s.shortCode
}

func (s *gridReference) GetEasting() float64 {
	return s.easting
}

func (s *gridReference) GetNorthing() float64 {
	return s.northing
}

func (s *gridReference) GetZone() int {
	return s.zone
}
