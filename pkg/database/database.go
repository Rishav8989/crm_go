// pkg/database/database.go
package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "./customers.db")
	if err != nil {
		log.Fatal(err)
	}

	createTable()
}

func createTable() {
	createCustomerTableSQL := `CREATE TABLE IF NOT EXISTS customers (
        id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
        name TEXT,
        role TEXT,
        email TEXT,
        phone TEXT,
        contacted BOOLEAN
    );`

	_, err := DB.Exec(createCustomerTableSQL)
	if err != nil {
		log.Fatal(err)
	}
}
