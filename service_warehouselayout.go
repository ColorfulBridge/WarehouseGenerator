package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/md-golibs/whlayout"
)

func WarehouseLayoutService(w http.ResponseWriter, r *http.Request) {

	type input struct {
		Return         string
		Layout         whlayout.LayoutParameters
		OutlineGeoJSON map[string]interface{}
	}

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()

	//Read body
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	checkErrors(w, err)

	//Parse inputmessage
	inputMessage := input{}
	err = json.Unmarshal(body, &inputMessage)
	checkErrors(w, err)

	//Get MakeGeoJSON back to string
	polygonGeoJSON, err := json.Marshal(inputMessage.OutlineGeoJSON)
	checkErrors(w, err)

	//Get layout from input
	layout := inputMessage.Layout

	//Generate warehouse layout
	wh := whlayout.GenerateLayout(polygonGeoJSON, layout)

	//Prepare output
	var resultString string
	var resultByte []byte

	if inputMessage.Return == "Racks" {
		resultByte, err = json.Marshal(wh.Racks)
	} else if inputMessage.Return == "Docks" {
		resultByte, err = json.Marshal(wh.Docks)
	} else if inputMessage.Return == "WarehouseGeoJSON" {
		resultString = wh.AsGeoJSON
	} else if inputMessage.Return == "RacksGeoJSON" {
		resultString = wh.RacksAsGeoJSON
	} else if inputMessage.Return == "DocksGeoJSON" {
		resultString = wh.DocksAsGeoJSON
	} else {
		resultByte, err = json.Marshal(wh)
	}
	checkErrors(w, err)

	if resultByte != nil {
		fmt.Fprint(w, string(resultByte))
	} else {
		fmt.Fprint(w, resultString)
	}

}
