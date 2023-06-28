package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HandleRouting() {
	r := mux.NewRouter()
	r.HandleFunc("/employees", GetEmployees).Methods("GET")
	r.HandleFunc("/employee/{eid}", GetEmployeeById).Methods("GET")
	r.HandleFunc("/employee/{eid}", UpdateEmployee).Methods("PUT")
	r.HandleFunc("/employee/{eid}", DeleteEmployee).Methods("DELETE")

	r.HandleFunc("/createemp",CreateEmp).Methods("POST")
	
	log.Fatal(http.ListenAndServe(":6000",r))
}