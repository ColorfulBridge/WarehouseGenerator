package wly

import "github.com/twpayne/go-geom"
import "github.com/twpayne/go-geom/encoding/geojson"
import "github.com/twpayne/go-geom/xy"

import "math"


const DEGREES_TO_RADIANS float64 = math.Pi / 180
const EARTH_MEAN_RADIUS_KM float64 = 6371.0087714
const DEG_TO_METER float64 = DEGREES_TO_RADIANS * EARTH_MEAN_RADIUS_KM * 1000

func unused(){
	_ = xy.Angle
}

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
func toLineGeoJSON(points []geom.Point) string{
	var c []float64 = flattenPoints(points)
	var lineString = geom.NewLineStringFlat(geom.XY, c)
	result, err := geojson.Marshal(lineString)
	check(err)
	return string(result)
}

func toPolygonGeoJSON(points []geom.Point) string{
	var c []geom.Coord = flattenPointsToCoords(points)
	var polygon = geom.NewPolygon(geom.XY).MustSetCoords([][]geom.Coord{c})
	result, err := geojson.Marshal(polygon)
	check(err)
	return string(result)
}

func pointBetweenWithPadding(p1, p2 geom.Point, part float64, padding [2]float64) *geom.Point{
	var diffX = p2.Coords()[0] - p1.Coords()[0]
	var diffY = p2.Coords()[1] - p1.Coords()[1]
	
	var newCoordX = p1.Coords()[0] + diffX * part
	var newCoordY = p1.Coords()[1] + diffY * part

	var point = geom.NewPointFlat(geom.XY, []float64{newCoordX + padding[0],newCoordY + padding[1]})
	return point
}

func pointBetween(p1, p2 geom.Point, part float64) *geom.Point{
	return pointBetweenWithPadding(p1,p2,part,[2]float64{0.0,0.0})
}

func degToMeter(deg float64) float64 {
	return deg * DEG_TO_METER
}

func meterToDeg(deg float64) float64 {
	return deg / DEG_TO_METER
}


func edgePoints(p *geom.Polygon) []geom.Point {

	arr := make([]geom.Point, 4)
	for i := 0; i < 4; i++ {
		arr[i] = *geom.NewPointFlat(geom.XY, p.Coords()[0][i])
	}
	return arr
}

func addPadding(p geom.Point, padding float64, pos int) geom.Point{
	if(pos == 1){
		return *geom.NewPointFlat(geom.XY, []float64{p.Coords()[0] + padding, p.Coords()[1] + padding})
	}else if(pos == 2){
		return *geom.NewPointFlat(geom.XY, []float64{p.Coords()[0] - padding, p.Coords()[1] + padding})
	}else if(pos == 3){
		return *geom.NewPointFlat(geom.XY, []float64{p.Coords()[0] - padding, p.Coords()[1] - padding})
	} else if(pos == 4){
		return *geom.NewPointFlat(geom.XY, []float64{p.Coords()[0] + padding, p.Coords()[1] - padding})
	} else {
		panic("no support pointnumber")
		return *geom.NewPointFlat(geom.XY, []float64{p.Coords()[0], p.Coords()[1]})
	}
}

func mustMarshallToGeoJSON(g geom.T) string{
	result, err := geojson.Marshal(g)
	check(err)
	return string(result)
}
