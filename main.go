package main

import (
	"fmt"
	"io/ioutil"
	"whs/wly"

	"github.com/twpayne/go-geom"
	"github.com/twpayne/go-geom/encoding/geojson"
)

func main() {

	data, err := ioutil.ReadFile("input/warehouse1.json")
	check(err)

	geometry := new(geom.T)
	err = geojson.Unmarshal(data, geometry)
	check(err)

	polygon := (*geometry).(*geom.Polygon)

	layout := wly.LayoutParameters{ }
	layout.Horizontal = false
	layout.Padding = 10 
	layout.BinsPerRack = 10
	layout.HLayout = []float64{1.075,0.05,1.05,0.05,1.05,0.05,1.05,0.05,1.05,0.05,1.05,0.05,1.05,0.05,1.05,0.05,1.05,0.05,1.075}
	layout.VLayout = []float64{1.1,0.35,1.1,0.35,1.1}
	layout.HDocks  = []float64{0.05,1.9,0.05}
	layout.VDocks  = []float64{1.1,0.2,1.1,0.2,1.1,0.2,1.1}
	var wh wly.Warehouse = wly.GenerateLayout(polygon, layout)

	fmt.Println(wh.Racks[0].BinsAsGeoJSON)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
