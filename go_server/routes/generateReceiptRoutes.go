package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	controller "restaurantApp/go_server/controllers"
	"restaurantApp/go_server/middlewares"
)

func GenerateReceiptRoutes(incomingRoutes *gin.Engine) {
	fmt.Println("Generate Receipt Routes")
	incomingRoutes.Use(middlewares.Authenticate())
	incomingRoutes.POST("/generate-receipt/:orderId", controller.GenerateReceipt)
}
