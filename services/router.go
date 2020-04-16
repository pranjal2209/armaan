package services

import (
	"fmt"
	"log"
	"net/http"
	//"github.com/gorilla/mux"
)

func HandlerFunc() {
	//router := mux.NewRouter().StrictSlash(true)
	//router := NewRouter().StrictSlash(true)
	http.HandleFunc("/", HomePage)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}
