package ops

import (
	"errors"
	"gogeo/gpkg"
	"math"

	"github.com/slawler/gdal"
)

// LineLength returns the distance along a straight line in euclidean space
func (gl gpkg.GoLine) LineLength() float64 {
	x0, y0 := gl[0][0], gl[0][1]
	x1, y1 := gl[1][0], gl[1][1]
	return math.Sqrt(math.Pow((x1-x0), 2) + math.Pow((y1-y0), 2))
}

// LineLength returns the distance along a straight line in euclidean space
func (gl gpkg.GoLineZ) LineLengthZ() float64 {
	x0, y0 := gl[0][0], gl[0][1]
	x1, y1 := gl[1][0], gl[1][1]
	return math.Sqrt(math.Pow((x1-x0), 2) + math.Pow((y1-y0), 2))
}

// Segments is a count of GoLines in a GoPolyLine
func (gpl GoPolyLine) Segments() (int, error) {
	lineSegments := len(gpl)
	if lineSegments < 1 {
		return 0, errors.New("Not a GoPolyLine? GoPolyLine must have a minimum of 1 line segments")
	}
	return lineSegments, nil
}

// SegmentsZ is a count of GoLineZs in a GoPolyLineZ
func (gplz GoPolyLineZ) SegmentsZ() (int, error) {
	lineSegments := len(gplz)
	if lineSegments < 1 {
		return 0, errors.New("Not a GoPolyLineZ? GoPolyLineZ must have a minimum of 1 line segments")
	}
	return lineSegments, nil
}

// TotalLineLength is the sum of the lengths of all
// GoLiners in a polyline
func (gpl GoPolyLine) TotalLineLength() (float64, error) {
	lineSegments := len(gpl)
	if lineSegments < 1 {
		return 0, errors.New("Not a GoPolyLine? GoPolyLine must have a minimum of 1 line segments")
	}

	totalLength := 0.0
	for _, line := range pl {
		totalLength += gpl.LineLength()
	}
	return totalLength, nil
}

// TotalLineLengthZ is the sum of the lengths of all
// GoLiners in a polyline
func (gplz GoPolyLineZ) TotalLineLengthZ() (float64, error) {
	lineSegments := len(gplz)
	if lineSegments < 1 {
		return 0, errors.New("Not a GoPolyLineZ? GoPolyLineZ must have a minimum of 1 line segments")
	}

	totalLength := 0.0
	for _, line := range pl {
		totalLength += gplz.LineLength()
	}
	return totalLength, nil
}

// Interpolate returns a new point along a straight line in euclidean space
// at a specified distance
func (gl gpkg.GoLine) Interpolate(d float64) (GoPoint, error) {
	if d > gl.LineLength() || d <= 0.0 {
		return GoPoint{nil, nil}, errors.New("Distance error: Point not on line? Requires extrapolation?")
	}

	distanceRatio := d / gl.LineLength()

	x0, y0 := gl[0][0], gl[0][1]
	x1, y1 := gl[1][0], gl[1][1]

	newX := (1-distanceRatio)*x0 + distanceRatio*x1
	newY := (1-distanceRatio)*y1 + distanceRatio*y1
	return GoPoint{newX, newY}, nil
}

// InterpolateZ returns a new point along a straight line in euclidean space
// at a specified distance
func (glz GoLineZ) InterpolateZ(d float64) (GoPointZ, error) {
	if d > glz.LineLengthZ() || d <= 0.0 {
		return GoPointZ{nil, nil, nil}, errors.New("Distance error: Point not on line? Requires extrapolation?")
	}

	distanceRatio := d / line.LineLengthZ()

	x0, y0, z0 := gl[0][0], gl[0][1], gl[0][2]
	x1, y1, z0 := gl[1][0], gl[1][1], gl[1][2]

	newX := (1-distanceRatio)*x0 + distanceRatio*x1
	newY := (1-distanceRatio)*y0 + distanceRatio*y1
	newZ := z0 + (d-glz.LineLengthZ())*((z1-z0)/glz.LineLengthZ())
	return gpz.GoPointZ{newX, newY, newZ}, nil
}

// InterpolatePolyLine returns a new point along a straight line in euclidean space
// at a specified distance
func (gpl GoPolyLine) InterpolatePolyLine(d float64) GoPoint {
	distanceAlongLine := 0.0
	lineSegments, _ := gpl.Segments()

	for i := 0; i < lineSegments; i++ {
		lineSegmentLength := gpl[i].LineLength()
		distanceAlongLine += lineSegmentLength
		switch {
		case distanceAlongLine > d:
			delta := lineSegmentLength - (distanceAlongLine - d)
			newPoint, _ := gpl[i].Interpolate(delta)
			return gpz.GoPoint{p[0], p[1]}

		default:
			continue
		}

	}
	return GoPoint{0, 0}
}

// ToLineString converts PolyLine to GDAL geometry type linestrings
func (gpl GoPolyLine) ToLineString() gdal.Geometry {
	lineString := gdal.Create(gdal.GT_LineString)
	nPoints, _ := gpl.Segments() + 1

	for idx, gl := range gpl {

		x0, y0 := gl[0][0], gl[0][1]
		x1, y1 := gl[1][0], gl[1][1]

		startPoint := x0, y0
		endPoint := x1, y1
		switch idx {
		// The last line segment appends both start and end points
		case nPoints:
			lineString.AddPoint2D(x0, y0)
			lineString.AddPoint2D(x1, y1)

		default:
			lineString.AddPoint2D(x1, y1)
		}

	}

	return lineString
}

// ToLineStringZ converts PolyLine to GDAL geometry type linestrings
func (gplz GoPolyLineZ) ToLineStringZ() gdal.Geometry {
	lineString := gdal.Create(gdal.GT_LineString25D)
	nPoints, _ := gpl.SegmentsZ() + 1

	x0, y0, z0 := gl[0][0], gl[0][1], gl[0][2]
	x1, y1, z1 := gl[1][0], gl[1][1], gl[1][2]

	for idx, line := range gpl {
		startPoint := x0, y0, z0
		endPoint := x1, y1, z1
		switch idx {
		// The last line segment appends both start and end points
		case nPoints:
			lineString.AddPoint(x0, y0, z0)
			lineString.AddPoint(x1, y1, z1)

		default:
			lineString.AddPoint(x1, y1, z1)
		}

	}

	return lineString
}
