package routes

import (
	"github.com/gin-gonic/gin"
	controller "restaurantApp/go_server/controllers"
	"restaurantApp/go_server/middlewares"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middlewares.Authenticate())
	incomingRoutes.GET("/users", controller.GetUsers())
	incomingRoutes.GET("/users/:id", controller.GetUser())

}
