package coordinates

type datum struct {
	name               string
	scaleFactor        float64
	trueOriginPhi      float64
	trueOriginLambda   int
	trueOriginEasting  int
	trueOriginNorthing int
	helmertTransform   [3]float64
	helmertScale       float64
	helmertRotation    [3]float64
	ellipsoid          ellipsoid
}

var (
	NationalGrid      datum
	IrishNationalGrid datum
	UtmNorth          datum
	UtmSouth          datum
)

func init() {
	NationalGrid = newDatum(
		"Ordnance Survey National Grid", 0.9996012717,
		49, -2, 400000, -100000,
		-446.448, 125.157, -542.06, 20.4894,
		-0.1502, -0.247, -0.8421, Airy1830)
	IrishNationalGrid = newDatum(
		"Irish National Grid", 1.000035,
		53.5, -8, 200000, 250000,
		-482.53, 130.596, -564.557, -8.15,
		1.042, 0.214, 0.631, Airy1830Modified)
	UtmNorth = newDatum(
		"UTM Northern Hemisphere", 0.9996,
		0, -3, 500000, 0, 0, 0, 0, 0, 0, 0, 0, WGS84)
	UtmSouth = newDatum(
		"UTM Southern Hemisphere", 0.9996,
		0, -3, 500000, 10000000, 0, 0, 0, 0, 0, 0, 0, WGS84)
}

func newDatum(
	name string,
	scaleFactor float64,
	trueOriginPhi float64,
	trueOriginLambda int,
	trueOriginEasting int,
	trueOriginNorthing int,
	helmertTransformX float64,
	helmertTransformY float64,
	helmertTransformZ float64,
	helmertScale float64,
	helmertRotationX float64,
	helmertRotationY float64,
	helmertRotationZ float64,
	ellipsoid ellipsoid,
) datum {
	return datum{
		name:               name,
		scaleFactor:        scaleFactor,
		trueOriginPhi:      trueOriginPhi,
		trueOriginLambda:   trueOriginLambda,
		trueOriginEasting:  trueOriginEasting,
		trueOriginNorthing: trueOriginNorthing,
		helmertTransform:   [3]float64{helmertTransformX, helmertTransformY, helmertTransformZ},
		helmertScale:       helmertScale,
		helmertRotation:    [3]float64{helmertRotationX, helmertRotationY, helmertRotationZ},
		ellipsoid:          ellipsoid,
	}
}
