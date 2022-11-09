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
		fmt.Fprintf(resp, "The customer for the given ID %v can't be found in the database", id_int)
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

	customer := unmarshalCustomerData(bytes)

	if customer != nil {
		addNewCustomer(customer)
		json.NewEncoder(resp).Encode(database)
	} else {
		resp.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(resp, "The given JSON data could not be parsed on the server.", id_int)
	}
}

func unmarshalCustomerData(sent_data []byte) *Customer {
	customer := &Customer{}
	if err := json.Unmarshal(sent_data, customer); err != nil {
		fmt.Println("Can't unmarshal given customer:", sent_data)
		return nil
	}

	return customer
}

func updateDatabaseForID(id int, new_customer_data *Customer) {
	fmt.Printf("Name: %v, type: %T\n", new_customer_data.Name)

	updated_name := &new_customer_data.Name
	updated_role := &new_customer_data.Role
	updated_email := &new_customer_data.Email
	updated_phone := &new_customer_data.Phone
	updated_contacted := &new_customer_data.Contacted

	fmt.Printf("updated_name %v\n", updated_name)
	fmt.Printf("updated_role  %v\n", updated_role)
	fmt.Printf("updated_email  %v\n", updated_email)
	fmt.Printf("updated_phone  %v\n", updated_phone)
	fmt.Printf("updated_contacted  %v\n", updated_contacted)

	existing_customer := database[id]
	fmt.Println("Existing:", existing_customer)
	if new_customer_data.Name != "" {
		existing_customer.Name = new_customer_data.Name
	} else {
		fmt.Printf("New Customer: %v has no name\n", new_customer_data)
	}

	if new_customer_data.Role != "" {
		existing_customer.Role = new_customer_data.Role
	}

	if new_customer_data.Email != "" {
		existing_customer.Email = new_customer_data.Email
	}

	if new_customer_data.Phone != "" {
		existing_customer.Phone = new_customer_data.Phone
	}

	if new_customer_data.Contacted != existing_customer.Contacted {
		existing_customer.Contacted = new_customer_data.Contacted
	}
	fmt.Println("db before updating:", database)
	database[id] = existing_customer
	fmt.Println("db after updating:", database)
}

func updateCustomer(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-type", "application/json")

	id := mux.Vars(req)["id"]

	id_int, _ := strconv.Atoi(id)
	if _, ok := database[id_int]; ok {

		bytes, _ := ioutil.ReadAll(req.Body)
		sent_customer_data := unmarshalCustomerData(bytes)

		if sent_customer_data != nil {
			resp.WriteHeader(http.StatusOK)
			fmt.Println("Update data: %v\n", sent_customer_data)
			updateDatabaseForID(id_int, sent_customer_data)

			fmt.Println("Updated database...\n")
			enc := json.NewEncoder(resp)
			enc.Encode(database)
		} else {
			resp.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(resp, "The given JSON data could not be parsed on the server.", id_int)
		}
	}
}

func deleteCustomer(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-type", "application/json")
	id := mux.Vars(req)["id"]
	id_int, _ := strconv.Atoi(id)

	if _, ok := database[id_int]; ok {
		resp.WriteHeader(http.StatusOK)
		delete(database, id_int)
		json.NewEncoder(resp).Encode(database)
	} else {
		resp.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(resp, "The customer for the given ID %v can't be found in the database", id_int)
	}
}

func listEndpoints(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-type", "text/html;charset=utf-8")
	fmt.Fprintf(resp, "<h1>List of existing endpoints</h1>\n")
	fmt.Fprintf(resp, "<ul>\n")
	fmt.Fprintf(resp, "<li>/ - to list all endpoints</li>")
	fmt.Fprintf(resp, "<li>/customers (GET) - to list all customers in the database</li>")
	fmt.Fprintf(resp, "<li>/customers/{id} (GET) - to list a specific customer in the database by its ID</li>")
	fmt.Fprintf(resp, "<li>/customers/{id} (PATCH) - to update specific customer's data in the database by its ID</li>")
	fmt.Fprintf(resp, "<li>/customers (POST) - to create a new customer in the database</li>")
	fmt.Fprintf(resp, "<li>/customers/{id} (DELETE) - to delete a specific customer from the database by its ID<</li>")
	fmt.Fprintf(resp, "</ul>")
}

func setupRouter() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/", listEndpoints).Methods("GET")
	r.HandleFunc("/customers", getCustomers).Methods("GET")
	r.HandleFunc("/customers/{id}", getCustomer).Methods("GET")

	r.HandleFunc("/customers/{id}", updateCustomer).Methods("PATCH")

	r.HandleFunc("/customers", createCustomer).Methods("POST")
	r.HandleFunc("/customers/{id}", deleteCustomer).Methods("DELETE")
	return r
}

func main() {

	port := 3000
	server_port := fmt.Sprintf(":%v", port)

	fmt.Println("Server is available at localhost:", server_port)
	log.Fatal(http.ListenAndServe(server_port, setupRouter()))
}
