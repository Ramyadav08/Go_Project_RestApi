package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HandleRouting() {
	r := mux.NewRouter()
	r.HandleFunc("/createemp",CreateEmp).Methods("POST")
	log.Fatal(http.ListenAndServe(":6000",r))
}