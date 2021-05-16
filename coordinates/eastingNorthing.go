package coordinates

import "math"

type eastingNorthing struct {
	easting  float64
	northing float64
}

func NewEastingsNorthings(
	latitude float64,
	longitude float64,
	datum datum,
) eastingNorthing {
	phi := degreesToRadians(latitude)
	lambda := degreesToRadians(longitude)

	ellipsoid := datum.ellipsoid
	a := ellipsoid.equatorialRadius
	b := ellipsoid.polarRadius
	eSq := ellipsoid.eccentricitySquared
	phi0 := degreesToRadians(datum.trueOriginPhi)
	lambda0 := degreesToRadians(float64(datum.trueOriginLambda))
	N0 := float64(datum.trueOriginNorthing)
	E0 := float64(datum.trueOriginEasting)
	f0 := datum.scaleFactor

	n := (a - b) / (a + b)
	v := (a * f0) * math.Pow(1.0-(eSq*math.Pow(math.Sin(phi), 2)), -0.5)
	p := (a * f0) * (1.0 - eSq) * math.Pow(1.0-(eSq*math.Pow(math.Sin(phi), 2.0)), -1.5)
	nSq := (v / p) - 1.0

	M := (1 + n + (5 / 4.0 * math.Pow(n, 2.0)) + (5 / 4.0 * math.Pow(n, 3.0))) * (phi - phi0)
	M = M - ((3*n + 3*math.Pow(n, 2.0) + 21/8.0*math.Pow(n, 3.0)) * math.Sin(phi-phi0) * math.Cos(phi+phi0))
	M = M + (((15 / 8.0 * math.Pow(n, 2.0)) + (15 / 8.0 * math.Pow(n, 3.0))) * math.Sin(2*(phi-phi0)) * math.Cos(2*(phi+phi0)))
	M = M - (35 / 24.0 * math.Pow(n, 3.0) * math.Sin(3*(phi-phi0)) * math.Cos(3*(phi+phi0)))
	M = b * f0 * M
	I := M + N0
	II := v / 2.0 * math.Sin(phi) * math.Cos(phi)
	III := v / 24.0 * math.Sin(phi) * math.Pow(math.Cos(phi), 3.0) * (5.0 - math.Pow(math.Tan(phi), 2.0) + 9.0*nSq)
	IIIA := v / 720.0 * math.Sin(phi) * math.Pow(math.Cos(phi), 5.0) * (61.0 - 58.0*math.Pow(math.Tan(phi), 2.0) + math.Pow(math.Tan(phi), 4.0))
	northing := I + II*math.Pow(lambda-lambda0, 2.0) + III*math.Pow(lambda-lambda0, 4.0) + IIIA*math.Pow(lambda-lambda0, 6.0)

	IV := v * math.Cos(phi)
	V := v / 6.0 * math.Pow(math.Cos(phi), 3.0) * (v/p - math.Pow(math.Tan(phi), 2.0))
	VI := v / 120.0 * math.Pow(math.Cos(phi), 5.0) * (5.0 - 18.0*math.Pow(math.Tan(phi), 2.0) + math.Pow(math.Tan(phi), 4.0) + 14.0*nSq - 58.0*(math.Pow(math.Tan(phi), 2.0)*nSq))
	easting := E0 + IV*(lambda-lambda0) + V*math.Pow(lambda-lambda0, 3.0) + VI*math.Pow(lambda-lambda0, 5.0)

	return eastingNorthing{
		easting:  easting,
		northing: northing,
	}
}
