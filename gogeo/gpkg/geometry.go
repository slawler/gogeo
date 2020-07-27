package gpkg

const (
	// OGCWKT4269 is a placehoder for projection info
	OGCWKT4269 string = `GEOGCS["NAD83",DATUM["North_American_Datum_1983",SPHEROID["GRS 1980",6378137,298.257222101,AUTHORITY["EPSG","7019"]],AUTHORITY["EPSG","6269"]],PRIMEM["Greenwich",0,AUTHORITY["EPSG","8901"]],UNIT["degree",0.01745329251994328,AUTHORITY["EPSG","9122"]],AUTHORITY["EPSG","4269"]]`
)

// GoPoint ...
type GoPoint [2]float64

// GoPointZ ...
type GoPointZ [3]float64

// GoMultiPoint ...
type GoMultiPoint []GoPoint

// GoMultiPointZ ...
type GoMultiPointZ []GoPointZ

// GoLine ...
type GoLine [2]GoPoint

// GoLineZ ..
type GoLineZ [2]GoPointZ

// GoPolyLine ...
type GoPolyLine []GoLine

// GoPolyLineZ ..
type GoPolyLineZ []GoLineZ
