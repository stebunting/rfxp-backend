package coordinates

func init() {
	WGS84 = newEllipsoid("WGS84", 6378137, 6356752.314245)
	GRS80 = newEllipsoid("GRS80", 6378137, 6356752.3141)
	Airy1830 = newEllipsoid("Airy 1830", 6377563.396, 6356256.909)
	Airy1830Modified = newEllipsoid("Airy 1830 Modified", 6377340.189, 6356034.447)

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
