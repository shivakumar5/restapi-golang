package main

import (
	"encoding/json"
	"io/ioutil"

	"log"
	"net/http"
	"fmt"

	"github.com/gorilla/mux"
)

// Employee ...
type Employee struct{
	ID string `json:"id"`
	Name string `json :"name"`
	Age int `json :"age"`
}

// allEmployees ...
var allEmployees []Employee

// Employees records ...
var Employees = []Employee{
	Employee{"1","shiva", 24},
	Employee{"2","kumar", 24},
	Employee{"3","patil", 24},
}


// CreateEmployee ...
func CreateEmployee(w http.ResponseWriter, r *http.Request){
	reqbody, err := ioutil.ReadAll(r.Body)
	if err != nil{
		fmt.Println("please provide valid data to create", err)
	}

	var employee Employee
	json.Unmarshal(reqbody, &employee)

	// append created entry to slice
	Employees = append(Employees, employee)
	w.WriteHeader(http.StatusCreated)
}

// GetAllEmployees ...
func GetAllEmployees(w http.ResponseWriter, r *http.Request){
	json.NewEncoder(w).Encode(Employees)
}

// GetSingleEmployee ...
func GetSingleEmployee(w http.ResponseWriter, r *http.Request){

	key := mux.Vars(r)["id"]

	for _, emp := range Employees{
		if emp.ID == key{
			json.NewEncoder(w).Encode(emp)
		}
	}

}

// UpdateEmployee ...
func UpdateEmployee(w http.ResponseWriter, r *http.Request){

	key := mux.Vars(r)["id"]
	reqbody, err := ioutil.ReadAll(r.Body)
	if err != nil{
		fmt.Println("please provide valid data to update", err)
	}

	var updateEmployee Employee
	json.Unmarshal(reqbody, &updateEmployee)

	for i, singleEmployee := range Employees{
		if singleEmployee.ID == key{
			singleEmployee.Age = updateEmployee.Age
			singleEmployee.Name = updateEmployee.Name
			singleEmployee.ID = updateEmployee.ID
			
			//append updated entry to slice
			Employees = append(Employees[:i], singleEmployee)			
		}
	}
}


// DeleteEmployee ...
func DeleteEmployee(w http.ResponseWriter, r *http.Request){
	
 	key := mux.Vars(r)["id"]
	
	for index, employee := range Employees{
		if employee.ID == key{
			Employees = append(Employees[:index], Employees[index+1:]...)
			w.WriteHeader(http.StatusOK)
		}
	}
	
}


func main(){

	newRouter := mux.NewRouter()

	newRouter.HandleFunc("/CreateEmployee", CreateEmployee).Methods("POST")
	newRouter.HandleFunc("/GetAllEmployees", GetAllEmployees)
	newRouter.HandleFunc("/GetAllEmployees/{id}", GetSingleEmployee)
	newRouter.HandleFunc("/DeleteEmployee/{id}", DeleteEmployee).Methods("DELETE")
	newRouter.HandleFunc("/UpdateEmployee/{id}", UpdateEmployee).Methods("PUT")

	log.Fatal(http.ListenAndServe(":8009", newRouter))

}
