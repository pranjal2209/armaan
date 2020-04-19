package services

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/pranjal2209/armaan/services/armaan"
)

func HandlerFunc() {
	//router := mux.NewRouter().StrictSlash(true)
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", armaan.RegisterUser)
	router.HandleFunc("/getcountries/", armaan.GetCountries)
	log.Fatal(http.ListenAndServe(getPort(), router))
}

func getPort() string {
	var port = os.Getenv("PORT")
	// Set a default port if there is nothing in the environment
	if port == "" {
		port = "5000"
		fmt.Println("INFO: No PORT environment variable detected, defaulting to " + port)
	}
	return ":" + port
}
