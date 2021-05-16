package coordinates

import "math"

type cartesian struct {
	x float64
	y float64
	z float64
}

func NewCartesian(c *coordinates) cartesian {
	phi := degreesToRadians(c.latitude)
	lambda := degreesToRadians(c.longitude)
	height := c.height

	a := c.ellipsoid.equatorialRadius
	b := c.ellipsoid.polarRadius
	eSq := c.ellipsoid.eccentricitySquared

	N := a / math.Sqrt(1.0-eSq*math.Pow(math.Sin(phi), 2.0))

	return cartesian{
		x: (N + height) * math.Cos(phi) * math.Cos(lambda),
		y: (N + height) * math.Cos(phi) * math.Sin(lambda),
		z: (N*math.Pow(b, 2.0)/math.Pow(a, 2.0) + height) * math.Sin(phi),
	}
}

func (s *cartesian) transform(datum datum) (float64, float64, float64) {
	s.helmertTransformation(datum)
	return s.toGeodetic(datum)
}

func (s *cartesian) helmertTransformation(datum datum) {
	cx := datum.helmertTransform[0]
	cy := datum.helmertTransform[1]
	cz := datum.helmertTransform[2]

	scaleFactor := datum.helmertScale / 1000000

	rx := secondsToRadians(datum.helmertRotation[0])
	ry := secondsToRadians(datum.helmertRotation[1])
	rz := secondsToRadians(datum.helmertRotation[2])

	s.x = cx + (1+scaleFactor)*(s.x-rz*s.y+ry*s.z)
	s.y = cy + (1+scaleFactor)*(s.y+rz*s.x-rx*s.z)
	s.z = cz + (1+scaleFactor)*(s.z-ry*s.x+rx*s.y)
}

func (s *cartesian) toGeodetic(datum datum) (float64, float64, float64) {
	ellipsoid := datum.ellipsoid
	a := ellipsoid.equatorialRadius
	eSq := ellipsoid.eccentricitySquared

	p := math.Sqrt(math.Pow(s.x, 2.0) + math.Pow(s.y, 2.0))

	var oldPhi float64
	var N float64
	phi := math.Atan(s.z / (p * (1.0 - eSq)))
	for {
		oldPhi = phi
		N := a / math.Sqrt(1.0-eSq*math.Pow(math.Sin(oldPhi), 2.0))
		phi = math.Atan((s.z + (eSq * N * math.Sin(oldPhi))) / p)
		if math.Abs(oldPhi-phi) < 0.000000000000000001 {
			break
		}
	}

	lambda := math.Atan(s.y / s.x)
	H := (p / math.Cos(phi)) - N

	return radiansToDegrees(phi), radiansToDegrees(lambda), H
}
