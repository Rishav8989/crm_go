// pkg/services/customer_service.go
package services

import (
	"cmr_go/pkg/database"
	"cmr_go/pkg/models"
	"database/sql"
	"log"
)

func AddNewCustomer(customer *models.Customer) error {
	query := `INSERT INTO customers(name, role, email, phone, contacted) VALUES (?, ?, ?, ?, ?)`
	stmt, err := database.DB.Prepare(query)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(customer.Name, customer.Role, customer.Email, customer.Phone, customer.Contacted)
	return err
}

func GetCustomerByID(id int) (*models.Customer, error) {
	query := `SELECT id, name, role, email, phone, contacted FROM customers WHERE id = ?`
	row := database.DB.QueryRow(query, id)

	customer := &models.Customer{}
	err := row.Scan(&customer.ID, &customer.Name, &customer.Role, &customer.Email, &customer.Phone, &customer.Contacted)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return customer, nil
}

func GetAllCustomers() ([]models.Customer, error) {
	query := `SELECT id, name, role, email, phone, contacted FROM customers`
	rows, err := database.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	customers := []models.Customer{}
	for rows.Next() {
		var customer models.Customer
		err := rows.Scan(&customer.ID, &customer.Name, &customer.Role, &customer.Email, &customer.Phone, &customer.Contacted)
		if err != nil {
			log.Println(err)
			continue
		}
		customers = append(customers, customer)
	}
	return customers, nil
}

func UpdateCustomer(id int, customer *models.Customer) error {
	query := `UPDATE customers SET name = ?, role = ?, email = ?, phone = ?, contacted = ? WHERE id = ?`
	_, err := database.DB.Exec(query, customer.Name, customer.Role, customer.Email, customer.Phone, customer.Contacted, id)
	return err
}

func DeleteCustomer(id int) error {
	query := `DELETE FROM customers WHERE id = ?`
	_, err := database.DB.Exec(query, id)
	return err
}
