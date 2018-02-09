package main

import (
	"testing"
	"fmt"
	"io/ioutil"
	"github.com/md-golibs/whlayout" 
)

func TestProcessing(t *testing.T) {

	t.Log("running test")

	polygonGeoJSON, err := ioutil.ReadFile("input/warehouse1.json")
	check(err)

	layout := wly.LayoutParameters{ }
	layout.Horizontal = false
	layout.Padding = 10 
	layout.BinsPerRack = 10
	layout.HLayout = []float64{1.075,0.05,1.05,0.05,1.05,0.05,1.05,0.05,1.05,0.05,1.05,0.05,1.05,0.05,1.05,0.05,1.05,0.05,1.075}
	layout.VLayout = []float64{1.1,0.35,1.1,0.35,1.1}
	layout.HDocks  = []float64{0.05,1.9,0.05}
	layout.VDocks  = []float64{1.1,0.2,1.1,0.2,1.1,0.2,1.1}
	
	var wh wly.Warehouse = wly.GenerateLayout(polygonGeoJSON, layout)
	fmt.Println(wh.Racks[0].BinsAsGeoJSON)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
