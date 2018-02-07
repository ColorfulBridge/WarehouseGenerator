package main

import (
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

	layout := wly.LayoutParameters{Horizontal : true, AisleSize : 10, RackSize : 8, Padding: 20, BinsPerRack : 15 }
	wly.GenerateLayout(polygon, layout)

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
