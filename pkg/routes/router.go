// pkg/routes/router.go
package routes

import (
	"cmr_go/pkg/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/", handlers.ListEndpoints)
	r.GET("/customers", handlers.GetCustomers)
	r.GET("/customers/:id", handlers.GetCustomer)
	r.POST("/customers", handlers.CreateCustomer)
	r.PATCH("/customers/:id", handlers.UpdateCustomer)
	r.DELETE("/customers/:id", handlers.DeleteCustomer)

	return r
}
