package routes

import (
	"github.com/gin-gonic/gin"
	controller "restaurantApp/go_server/controllers"
	"restaurantApp/go_server/middlewares"
)

func ItemRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middlewares.Authenticate())
	incomingRoutes.GET("/items", controller.GetItems())
	incomingRoutes.GET("/items/:id", controller.GetItem())
	incomingRoutes.POST("/items", controller.CreateItem())
	incomingRoutes.PUT("/items/:id", controller.UpdateItem())
	incomingRoutes.DELETE("/items/:id", controller.DeleteItem())
}
