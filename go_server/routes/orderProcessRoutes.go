package routes

import (
	"github.com/gin-gonic/gin"
	controller "restaurantApp/go_server/controllers"
	"restaurantApp/go_server/middlewares"
)

func OrderProcessRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middlewares.Authenticate())
	incomingRoutes.POST("/order-process", controller.ProcessOrder)
}
