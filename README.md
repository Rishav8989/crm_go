# UDAcity Customer Relationship Management Backend



## TLDR;
Just run the command `go run main.go` to start an HTTP server at
localhost:3000.
Alternative: `go build && ./crm`
(Seem like `go run` without a file name [never worked](https://groups.google.com/g/golang-nuts/c/Uc6FurMDtQc?pli=1).)

## Introduction
This is a demo project that was created for a course at you udacity.com about the language [golang](https://go.dev/).
This application represents an HTTP server, that stores customer data. 
This customer data can be manipulated in a CRUD session, meaning:
- customer data can be **created**
- customer data can be **retrieved/read**
- customer data can be **updated**
- customer data can be **deleted**

A customer has the following fields in the database:
- id (A customer's database internal identifier, this identifier is handled internally and can't be changed)
- name (A customers first and last name - String)
- phone (Contact information about the customer - String)
- contacted (A yes/no - true/false entry about whether the customer has been contacted yet)
- email (Contact information about the customer - String)
- role (Some form of details about the customer - String)

## Available Endpoints

- localhost:3000/customers via "GET" method to **retrieve a list** of existing customers
- localhost:3000/customers/{id} via "GET" method to **retrieve data** about a specific customer
- localhost:3000/customers/{id} via "PATCH" method to **update** a specific customer
- localhost:3000/customers via "POST" method to **create** a new customer
- localhost:3000/customers/{id} via "DELETE" method to **delete** a specific customer


## Provided shell scripts for testing end points
This project contains a few shell scripts that run curl commands against a running server at
localhost:3000. This obviously requires that the server is running.
Each shell script is meant to run individually, i.e. run again a freshly started server and use
the existing database structure, which contains *three test entries*.
All shell scripts use `curl` commands that run against existing endpoints. Each individual `curl`
execution follows an echo statement that states against which HTTP method is used, and what the
test should achieve.
So:  
*Be aware: when running multiple shell scripts, without restarting the server in between, the results may differ from what the echo statement assumes.*

## Usage
The usage of the endpoints can be derived from the test shell scripts.

### Example 
#### Create new customer
curl --data "{ \"name\":\"Test User\", \
  \"phone\":\"555-456\", \
  \"contacted\":false, \
  \"role\":\"New Lead\", \
  \"email\":\"testuser@udacity.com\" }" \
  --header 'Content-Type: application/json' \
  localhost:3000/customers

