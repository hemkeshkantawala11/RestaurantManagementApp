package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
	helper "restaurantApp/go_server/helpers"
)

func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientToken := c.Request.Header.Get("token")
		if clientToken == "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "No token found"})
			c.Abort()
			return
		}

		claims, err := helper.ValidateToken(clientToken)
		if err != "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Please enter valid Token"})
			c.Abort()
			return
		}

		c.Set("email", claims.Email)
		c.Set("uid", claims.Uid)
		c.Set("FirstName", claims.FirstName)
		c.Set("LastName", claims.LastName)
		c.Set("User_type", claims.User_type)
		c.Next()
	}

}
