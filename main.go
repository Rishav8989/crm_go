package main

import (
  "fmt"
  "net/http"
  "github.com/gorilla/mux"
)

type Customer struct {
  ID int
  Name string
  Role string
  Email string
  Phone string
  Contacted bool
}


var database = map[int]Customer {
   1: Customer{
      ID: 1,
      Name: "Michael J",
      Role: "Student",
      Email: "Michael@udacity.com",
      Phone: "555-Nase",
      Contacted: true,
  },
   2: Customer{
      ID: 2,
      Name: "John Doe",
      Role: "Engineer",
      Email: "JohnD@udacity.com",
      Phone: "555-123456789",
      Contacted: false,
  },
   3: Customer{
      ID: 3,
      Name: "Jane Doe",
      Role: "Data Scientist",
      Email: "JaneD@udacity.com",
      Phone: "555-97654321",
      Contacted: false,
  },
}

func getCustomer(resp http.ResponseWriter, req *http.Request) {
  
}


func getCustomers(resp http.ResponseWriter, req *http.Request) {
}


func createCustomer(resp http.ResponseWriter, req *http.Request) {
}


func updateCustomer(resp http.ResponseWriter, req *http.Request) {
}


func deleteCustomer(resp http.ResponseWriter, req *http.Request) {
}

func setupRouter(r *mux.Router) {
  r.HandleFunc("/", getCustomers).Methods("GET")
  r.HandleFunc("/customer/{id}", getCustomer).Methods("GET")

  r.HandleFunc("/customer", updateCustomer).Methods("PUT")

  r.HandleFunc("/customer", createCustomer).Methods("POST")
  r.HandleFunc("/customer/{id}", deleteCustomer).Methods("DELETE")
}

func main() {
  
  port := 3000
  router := mux.NewRouter()
  fmt.Printf("type: %T\n", router)

  setupRouter(router)

  fmt.Println(database)
  server_port := fmt.Sprintf(":%v", port)
  fmt.Println("Server is available at localhost:", server_port)
  http.ListenAndServe(server_port, router)
}
