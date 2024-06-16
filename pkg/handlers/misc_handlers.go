// pkg/handlers/misc_handlers.go
package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListEndpoints(c *gin.Context) {
	c.HTML(http.StatusOK, "text/html", `<h1>List of existing endpoints</h1>
    <ul>
        <li>/ - to list all endpoints</li>
        <li>/customers (GET) - to list all customers</li>
        <li>/customers/:id (GET) - to get a customer by ID</li>
        <li>/customers (POST) - to create a new customer</li>
        <li>/customers/:id (PATCH) - to update a customer by ID</li>
        <li>/customers/:id (DELETE) - to delete a customer by ID</li>
    </ul>`)
}
