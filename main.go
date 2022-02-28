package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var api_version = ""

/*
func readFile() {
	version, err := os.ReadFile("VERSION")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}
	api_version = string(version)
}
*/

//home page
func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to ReaQta %s", api_version)
}

//api
func apiHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to ReaQta API %s", api_version)
}

//HealthCheckHandler is for readiness and liveness probes
func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "The service is healthy...")
}

func main() {
	// read version
	//readFile()
	// create a new router
	router := mux.NewRouter().StrictSlash(true)
	log.Print("the service is working...")
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/api", apiHandler)
	router.HandleFunc("/health", healthCheckHandler)

	// listen and serve
	log.Fatal(http.ListenAndServe(":8080", router))
}
