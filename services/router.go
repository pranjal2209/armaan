package services

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func HandlerFunc() {
	//router := mux.NewRouter().StrictSlash(true)
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", HomePage)
	log.Fatal(http.ListenAndServe(getPort(), nil))
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome Greetings!")
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
