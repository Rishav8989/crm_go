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

func main() {
  
  port := 3000
  router := mux.NewRouter()

  fmt.Println(database)
  server_port := fmt.Sprintf(":%v", port)
  fmt.Println("Server is available at localhost:",server_port)
  http.ListenAndServe(server_port, router)
}
