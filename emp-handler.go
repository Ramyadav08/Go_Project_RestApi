package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func CreateEmp(w http.ResponseWriter,r *http.Request){
	w.Header().Set("content-Type","application/json")
	var emp Employee
	json.NewDecoder(r.Body).Decode(&emp)
	DataBase.Create(&emp)
	json.NewEncoder(w).Encode(emp)
}

func GetEmployees(w http.ResponseWriter,r *http.Request){
	w.Header().Set("content-Type","application/json")
	var emp []Employee
	DataBase.Find(&emp)
	json.NewEncoder(w).Encode(emp)
}

func GetEmployeeById(w http.ResponseWriter,r *http.Request){
	w.Header().Set("content-Type","application/json")
	var emp Employee
	DataBase.Find(&emp,mux.Vars(r)["eid"])
	json.NewEncoder(w).Encode(emp)
}

func UpdateEmployee(w http.ResponseWriter,r *http.Request){
	w.Header().Set("content-Type","application/json")
	var emp Employee
	DataBase.Find(&emp,mux.Vars(r)["eid"])
	json.NewDecoder(r.Body).Decode(&emp)
	DataBase.Save(&emp)
	json.NewEncoder(w).Encode(emp)
}

func DeleteEmployee(w http.ResponseWriter,r *http.Request){
	w.Header().Set("content-Type","application/json")
	var emp Employee
	DataBase.Delete(&emp,mux.Vars(r)["eid"])
	
	json.NewEncoder(w).Encode("employee deleted now")
}