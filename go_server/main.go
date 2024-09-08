// main.go
package main

import (
	"restaurantApp/go_server/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/generate-receipt", controllers.GenerateReceipt)
	r.POST("/process-order", controllers.ProcessOrder)

	r.Run(":8080") // Listen and serve on 0.0.0.0:8080
}
