package wly

import "github.com/twpayne/go-geom"
import "math"


const DEGREES_TO_RADIANS float64 = math.Pi / 180
const EARTH_MEAN_RADIUS_KM float64 = 6371.0087714
const DEG_TO_METER float64 = DEGREES_TO_RADIANS * EARTH_MEAN_RADIUS_KM * 1000


func flattenPoints(points []geom.Point) []float64 {
	var a = make([]float64, len(points)*2)
	for i := 0; i < len(points); i++ {
		a[i*2] = points[i].Coords()[0]
		a[i*2+1] = points[i].Coords()[1]
	}
	return a
}

func flattenPointsToCoords(points []geom.Point) []geom.Coord {
	var a = make([]geom.Coord, len(points))
	for i := 0; i < len(points); i++ {
		a[i] = geom.Coord{points[i].Coords()[0], points[i].Coords()[1]}
	}
	return a
}

func flattenPoint(point geom.Point) []float64 {
	var a = make([]float64, 2)
		a[0] = point.Coords()[0]
		a[1] = point.Coords()[1]
	return a
}

func degToMeter(deg float64) float64 {
	return deg * DEG_TO_METER
}

func meterToDeg(deg float64) float64 {
	return deg / DEG_TO_METER
}