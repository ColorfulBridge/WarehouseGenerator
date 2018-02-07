package wly

import "fmt"
import "github.com/twpayne/go-geom"

func splitPolygon(polygon *geom.Polygon, axis int, splits int) []geom.Polygon {

	var p []geom.Point = edgePoints(polygon)
	var polygonsSplit = make([]geom.Polygon,2)

	if(axis == 0){

		var pm1 = pointBetween(p[0], p[1], 0.5)
		var pm2 = pointBetween(p[3], p[4], 0.5)

		var poly1Points = []geom.Point{p[0], *pm1, *pm2, p[4], p[0]}
		polygonsSplit[0] = *geom.NewPolygon(geom.XY).MustSetCoords([][]geom.Coord{flattenPointsToCoords(poly1Points)})
		
		var poly2Points = []geom.Point{*pm1, p[1], p[3], *pm2, *pm1}
		polygonsSplit[1] = *geom.NewPolygon(geom.XY).MustSetCoords([][]geom.Coord{flattenPointsToCoords(poly2Points)})
		
		fmt.Println()
	}

	return polygonsSplit
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
