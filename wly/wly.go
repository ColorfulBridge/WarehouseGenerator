package wly

import "fmt"
import "github.com/twpayne/go-geom"


type LayoutParameters struct{
	Horizontal bool
	RackSize float64
	AisleSize float64
	Padding float64
	BinsPerRack int
}

type Rack struct{
	Id string
	Outline geom.Polygon
	AsGeoJSON string
	StorageBins []StorageBin
	BinsAsGeoJSON string
}

type StorageBin struct{
	Id string
	Outline geom.Polygon
	AsGeoJSON string
	accessPoint geom.Point
}

type WarehouseLayout struct{
	Racks []Rack
	Outline geom.Polygon
	AsGeoJSON string
	RacksAsGeoJSON string
}


func GenerateLayout(polygon *geom.Polygon, layout LayoutParameters) WarehouseLayout {

	var wh WarehouseLayout
	var points []geom.Point = edgePoints(polygon)
	var padding = meterToDeg(layout.Padding)
		
	var p1 = addPadding(points[0], padding, 1)
	var p2 = addPadding(points[1], padding, 2)
	var p3 = addPadding(points[2], padding, 3)
	var p4 = addPadding(points[3], padding, 4) 

	fmt.Println(toPolygonGeoJSON([]geom.Point{p1,p2,p3,p4,p1}))
	
	wh.Racks = []Rack{}
	wh.Outline = *polygon
	wh.AsGeoJSON = toPolygonGeoJSON([]geom.Point{p1,p2,p3,p4,p1})

	wh.Racks, wh.RacksAsGeoJSON = generateRacks(polygon, layout)

	return wh
}


func check(e error) {
	if e != nil {
		panic(e)
	}
}
