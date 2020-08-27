package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func teste(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Teste 456")
}

func version(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("v2.0")
}

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/", teste).Methods("GET")
	router.HandleFunc("/version", version).Methods("GET")

	log.Fatal(http.ListenAndServe("0.0.0.0:8000", router))
}
