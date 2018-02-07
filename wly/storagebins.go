package wly
 

/* 
func generateStorageBins(rack Rack, layout LayoutParameters) []StorageBin {

	/*var points []geom.Point = edgePoints(polygon)
	var padding = meterToDeg(layout.Padding)
		
	var p1 = addPadding(points[0], padding, 1)
	var p2 = addPadding(points[1], padding, 2)
	var p3 = addPadding(points[2], padding, 3)
	var p4 = addPadding(points[3], padding, 4) 

	fmt.Println(toPolygonGeoJSON([]geom.Point{p1,p2,p3,p4,p1}))
	
	var racks = []Rack{}
	
	if(layout.Horizontal){
	
		var rackSpace = degToMeter(xy.Distance(flattenPoint(p1),flattenPoint(p2)))
		var numOfRacks = int(math.Floor(rackSpace / float64(layout.RackSize + layout.AisleSize)))
		
		//var totalSize = 2.0 * layout.Padding + float64(numOfRacks) * (layout.AisleSize + layout.RackSize)
		//var paddingCor = (distance - totalSize)/2
		var rackStartPart = 0.0
		var collection = geom.NewGeometryCollection()

		for i := 0; i < numOfRacks; i++{
			rackStartPart += layout.AisleSize / rackSpace
			var rackEndPart = rackStartPart + (layout.RackSize / rackSpace)

			var rackP1 = pointBetween(p1,p2,rackStartPart)
			var rackP2 = pointBetween(p1,p2,rackEndPart)
			
			var rackP3 = pointBetween(p4,p3,rackEndPart)
			var rackP4 = pointBetween(p4,p3,rackStartPart)

			rackStartPart = rackEndPart
			 
			
			var polyPoints = []geom.Point{*rackP1, *rackP2, *rackP3, *rackP4, *rackP1}
			var rackPolygon = geom.NewPolygon(geom.XY).MustSetCoords([][]geom.Coord{flattenPointsToCoords(polyPoints)})
			collection.Push(rackPolygon);

			var newRack =  Rack{ i, *rackPolygon, mustMarshallToGeoJSON(rackPolygon)}
			newRack.StorageBins = generateStorageBins(newRack, layout)

			racks = append(racks,rack)
			
		}

		result, err := geojson.Marshal(collection)
		check(err)
		_ = result
		fmt.Println(string(result))
		//calculatePointBetween();


	} else {

		panic("vertical not implemented")
		
	}

	return []StorageBin{}

}*/
	  
