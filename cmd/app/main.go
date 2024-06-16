// cmd/app/main.go
package main

import (
	"cmr_go/pkg/database"
	"cmr_go/pkg/routes"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3" // Import SQLite driver
)

func main() {
	// Initialize SQLite database
	database.InitDB()
	defer database.DB.Close()

	// Set Gin to release mode
	gin.SetMode(gin.ReleaseMode)

	// Create a Gin router
	router := gin.Default()

	// Setup routes
	router = routes.SetupRouter()

	// Define server port
	port := 3000

	// Start the server
	serverPort := fmt.Sprintf(":%v", port)
	fmt.Printf("Server is running at localhost%s\n", serverPort)
	log.Fatal(http.ListenAndServe(serverPort, router))
}
