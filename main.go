package main

import (
	"fmt"
	"log"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"whs/wly"
	"github.com/twpayne/go-geom"
	"github.com/twpayne/go-geom/encoding/geojson"
)

type input struct{
	Return string
	Layout wly.LayoutParameters
	OutlineGeoJSON map[string]interface{}
}

func main() {
	http.HandleFunc("/api/whl", handle)
	http.HandleFunc("/_ah/health", healthCheckHandler)
	log.Print("Listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "ok")
}

func checkErrors(w http.ResponseWriter, err error){
	if(err != nil){
		fmt.Fprint(w, err.Error())
		w.WriteHeader(400)
		panic(err.Error())
	}
}

func handle(w http.ResponseWriter, r *http.Request) {
	defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered in f", r)
        }
	}()
	
	if(r.Method != "POST"){
		fmt.Fprint(w, "Only POST supported")
		w.WriteHeader(400)
		return
	}

	defer r.Body.Close()
	
	body, err := ioutil.ReadAll(r.Body)
	checkErrors(w, err);
	
	inputMessage := input{}
	err = json.Unmarshal(body, &inputMessage)
	checkErrors(w, err);

	//Get MakeGeoJSON back to string
	geoJSON, err := json.Marshal(inputMessage.OutlineGeoJSON)
	checkErrors(w, err);

	//Male geoJSON to Polygon
	geometry := new(geom.T)
	err = geojson.Unmarshal(geoJSON, geometry)
	checkErrors(w,err)
	polygon := (*geometry).(*geom.Polygon)

	//Get layout
	layout := inputMessage.Layout

	var wh wly.Warehouse = wly.GenerateLayout(polygon, layout)
	var resultString string
	var resultByte []byte

	if inputMessage.Return == "Racks"{
		resultByte, err = json.Marshal(wh.Racks)
	} else if inputMessage.Return == "Docks" {
		resultByte, err = json.Marshal(wh.Docks)
	} else if(inputMessage.Return == "WarehouseGeoJSON"){
		resultString = wh.AsGeoJSON
	} else if(inputMessage.Return == "RacksGeoJSON"){
		resultString = wh.RacksAsGeoJSON
	} else if(inputMessage.Return == "DocksGeoJSON"){
		resultString = wh.DocksAsGeoJSON
	} else {
		resultByte, err = json.Marshal(wh)
	}
		checkErrors(w,err)

	if(resultByte != nil){
		fmt.Fprint(w, string(resultByte))
	} else {
		fmt.Fprint(w, resultString)
	}

}

