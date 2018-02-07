package wly


import "github.com/twpayne/go-geom"

func SplitPolygonEqually(polygon *geom.Polygon, axis int, splits int) []geom.Polygon {

	var p []geom.Point = edgePoints(polygon)
	var polygons []geom.Polygon
	var factor = 1 / float32(splits)

	if(axis == 0){

		var pLU = p[0]
		var pLO = p[3]

		for i := 1; i <= splits; i++{
			var pm1 = pointBetween(p[0], p[1], float64(factor * float32(i)))
			var pm2 = pointBetween(p[3], p[2], float64(factor * float32(i)))

			var polyPoints = []geom.Point{pLU, *pm1, *pm2, pLO, pLU}
			var splitPolygon = geom.NewPolygon(geom.XY).MustSetCoords([][]geom.Coord{flattenPointsToCoords(polyPoints)})
			polygons = append(polygons, *splitPolygon)
			
			pLU = *pm1
			pLO = *pm2
		}
	
	} else if(axis == 1){

		var pLU = p[0]
		var pRU = p[1]

		for i := 1; i <= splits; i++{
			var pm1 = pointBetween(p[0], p[3], float64(factor * float32(i)))
			var pm2 = pointBetween(p[1], p[2], float64(factor * float32(i)))

			var polyPoints = []geom.Point{pLU, pRU, *pm2, *pm1, pLU}
			var splitPolygon = geom.NewPolygon(geom.XY).MustSetCoords([][]geom.Coord{flattenPointsToCoords(polyPoints)})
			polygons = append(polygons, *splitPolygon)

			pLU = *pm1
			pRU = *pm2
		}

	}
	
	return polygons
}

func SplitPolygonWithFactors(polygon *geom.Polygon, axis int, splitFactors []float64) []geom.Polygon {

	var p []geom.Point = edgePoints(polygon)
	var polygons []geom.Polygon


	if(axis == 0){

		var pLU = p[0]
		var pLO = p[3]
		var factorSum float64 = 0

		for i := 1; i <= len(splitFactors); i++{
			factorSum += splitFactors[i-1]

			var pm1 = pointBetween(p[0], p[1], factorSum)
			var pm2 = pointBetween(p[3], p[2], factorSum)

			var polyPoints = []geom.Point{pLU, *pm1, *pm2, pLO, pLU}
			var splitPolygon = geom.NewPolygon(geom.XY).MustSetCoords([][]geom.Coord{flattenPointsToCoords(polyPoints)})
			polygons = append(polygons, *splitPolygon)
			
			pLU = *pm1
			pLO = *pm2
		}
	
	} else if(axis == 1){

		var pLU = p[0]
		var pRU = p[1]
		var factorSum float64 = 0

		for i := 1; i <= len(splitFactors); i++{
			factorSum += splitFactors[i-1]

			var pm1 = pointBetween(p[0], p[3], factorSum)
			var pm2 = pointBetween(p[1], p[2], factorSum)

			var polyPoints = []geom.Point{pLU, pRU, *pm2, *pm1, pLU}
			var splitPolygon = geom.NewPolygon(geom.XY).MustSetCoords([][]geom.Coord{flattenPointsToCoords(polyPoints)})
			polygons = append(polygons, *splitPolygon)
			
			pLU = *pm1
			pRU = *pm2
		}

	}
	
	return polygons
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
