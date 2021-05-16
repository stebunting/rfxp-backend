package coordinates

import "math"

type ellipsoid struct {
	name                string
	equatorialRadius    float64
	polarRadius         float64
	flattening          float64
	eccentricitySquared float64
}

var (
	WGS84            ellipsoid
	GRS80            ellipsoid
	Airy1830         ellipsoid
	Airy1830Modified ellipsoid
)

func newEllipsoid(name string, equatorialRadius float64, polarRadius float64) ellipsoid {
	return ellipsoid{
		name:                name,
		equatorialRadius:    equatorialRadius,
		polarRadius:         polarRadius,
		flattening:          1 - (polarRadius / equatorialRadius),
		eccentricitySquared: (math.Pow(equatorialRadius, 2.0) - math.Pow(polarRadius, 2.0)) / math.Pow(equatorialRadius, 2.0),
	}
}
