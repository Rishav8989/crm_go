# UDAcity Customer Relationship Management Backend

## Available Endpoints

- localhost:3000/customers via "GET" method to **retrieve a list** of existing customers
- localhost:3000/customers/{id} via "GET" method to **retrieve data** about a specific customer
- localhost:3000/customers/{id} via "PUT" method to **update** a specific customer
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
