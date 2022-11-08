package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type Customer struct {
  ID        int 
  Name      string `json: "name,omitempty"`
	Role      string `json: "name,omitempty"`
	Email     string `json: "name,omitempty"`
	Phone     string `json: "name,omitempty"`
	Contacted bool   `json: "name,omitempty"`
}

var database = map[int]Customer{
	1: Customer{
		ID:        1,
		Name:      "Michael J",
		Role:      "Student",
		Email:     "Michael@udacity.com",
		Phone:     "555-Nase",
		Contacted: true,
	},
	2: Customer{
		ID:        2,
		Name:      "John Doe",
		Role:      "Engineer",
		Email:     "JohnD@udacity.com",
		Phone:     "555-123456789",
		Contacted: false,
	},
	3: Customer{
		ID:        3,
		Name:      "Jane Doe",
		Role:      "Data Scientist",
		Email:     "JaneD@udacity.com",
		Phone:     "555-97654321",
		Contacted: false,
	},
}

func findNextIndex() int {
	index := 1

	for {
		if val, ok := database[index]; ok {
			fmt.Printf("found index: %v with value: %v - ok: %v\n", index, val, ok)
		} else {
			fmt.Printf("Did not find find index: %v - ok: %v\n", index, val, ok)
			break
		}
		index++
	}
	return index
}

func addNewCustomer(customer *Customer) {

	new_index := findNextIndex()
	customer.ID = new_index
	database[new_index] = *customer
}

func getCustomer(resp http.ResponseWriter, req *http.Request) {

	resp.Header().Set("Content-type", "application/json")
	id := mux.Vars(req)["id"]
	id_int, _ := strconv.Atoi(id)

	if _, ok := database[id_int]; ok {
		resp.WriteHeader(http.StatusOK)
		enc := json.NewEncoder(resp)
		enc.Encode(database[id_int])
	} else {
		resp.WriteHeader(http.StatusNotFound)
		enc := json.NewEncoder(resp)
		enc.Encode(database)
	}
}

func getCustomers(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-type", "application/json")
	resp.WriteHeader(http.StatusOK)
	json.NewEncoder(resp).Encode(database)
}

func createCustomer(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-type", "application/json")

	bytes, _ := ioutil.ReadAll(req.Body)

	customer := &Customer{}
	if err := json.Unmarshal(bytes, customer); err != nil {
		fmt.Println("Can't unmarshal given customer:", bytes)
	}

	addNewCustomer(customer)

	json.NewEncoder(resp).Encode(database)
}

func updateCustomer(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-type", "application/json")
	id := mux.Vars(req)["id"]
	id_int, _ := strconv.Atoi(id)
	if _, ok := database[id_int]; ok {
		resp.WriteHeader(http.StatusOK)
		enc := json.NewEncoder(resp)
		enc.Encode(database[id_int])
	}
}

func deleteCustomer(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-type", "application/json")
	id := mux.Vars(req)["id"]
	id_int, _ := strconv.Atoi(id)

	if _, ok := database[id_int]; ok {
		resp.WriteHeader(http.StatusOK)
    delete(database, id_int)
	} else {
		resp.WriteHeader(http.StatusNotFound)
  }
  json.NewEncoder(resp).Encode(database)
}

func setupRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", getCustomers).Methods("GET")
	r.HandleFunc("/customer/{id}", getCustomer).Methods("GET")

	r.HandleFunc("/customer", updateCustomer).Methods("PUT")

	r.HandleFunc("/customer", createCustomer).Methods("POST")
	r.HandleFunc("/customer/{id}", deleteCustomer).Methods("DELETE")
	return r
}

func main() {

	port := 3000
	server_port := fmt.Sprintf(":%v", port)

	fmt.Println("Server is available at localhost:", server_port)
	log.Fatal(http.ListenAndServe(server_port, setupRouter()))
}
