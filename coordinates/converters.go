package coordinates

import (
	"errors"
	"math"
	"strings"
)

func degreesToRadians(degrees float64) float64 {
	return degrees * math.Pi / 180
}

func secondsToRadians(seconds float64) float64 {
	return degreesToRadians(seconds) / 3600
}

func radiansToDegrees(radians float64) float64 {
	return radians * 180 / math.Pi
}

func toDecimal(degrees int, minutes int, seconds float64, direction string) (float64, error) {
	direction = strings.ToUpper(direction)
	if direction != "N" && direction != "S" && direction != "W" && direction != "E" {
		return 0, errors.New("invalid direction")
	}

	var multiplier float64 = 1.0
	if direction == "W" || direction == "S" {
		multiplier = -1.0
	}

	return multiplier * (float64(degrees) + float64(minutes)/60.0 + seconds/3600.0), nil
}
