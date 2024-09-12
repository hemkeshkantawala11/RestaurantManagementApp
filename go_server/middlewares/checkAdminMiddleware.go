// go_server/middlewares/auth.go
package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"restaurantApp/go_server/helpers"
)

func CheckAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := helpers.CheckUserType(c, "admin")
		if err != nil {
			c.JSON(http.StatusForbidden, gin.H{"error": "You are not authorized to perform this action"})
			c.Abort()
			return
		}
		c.Next()
	}
}
