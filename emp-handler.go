package main

import (
	"encoding/json"
	"net/http"
)

func CreateEmp(w http.ResponseWriter,r *http.Request){
	w.Header().Set("content-Type","application/json")
	var emp Employee
	json.NewDecoder(r.Body).Decode(&emp)
	DataBase.Create(&emp)
	json.NewEncoder(w).Encode(emp)
}