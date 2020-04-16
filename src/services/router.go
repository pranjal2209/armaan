package services

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HandlerFunc() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", HomePage)
	log.Fatal(http.ListenAndServe(":8080", router))
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}
