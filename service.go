package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"google.golang.org/appengine"
)

func main() {
	router := mux.NewRouter()

	//Register api paths
	router.Path("/api/whlayout").Methods("POST").HandlerFunc(WarehouseLayoutService)
	router.Path("/_ah/health").Methods("GET").HandlerFunc(healthCheckHandler)

	//Startup
	http.Handle("/", router)
	fmt.Println("server is starting up")
	appengine.Main()
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "ok")
}

func checkErrors(w http.ResponseWriter, err error) {
	if err != nil {
		fmt.Fprint(w, err.Error())
		w.WriteHeader(400)
		panic(err.Error())
	}
}
