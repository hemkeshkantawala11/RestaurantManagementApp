// main.go
package main

import (
	"github.com/gin-gonic/gin"
	"restaurantApp/go_server/routes"
)

func main() {
	router := gin.Default()
	router.Use(gin.Logger())

	routes.AuthRoutes(router)
	routes.UserRoutes(router)
	routes.GenerateReceiptRoutes(router)
	routes.OrderProcessRoutes(router)
	routes.ItemRoutes(router)

	router.Run(":8080") // Listen and serve on 0.0.0.0:8080
}
