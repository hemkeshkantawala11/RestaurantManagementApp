// views/error_response.go
package views

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorResponse(c *gin.Context, message string) {
	c.JSON(http.StatusBadRequest, gin.H{"error": message})
}
func SuccessResponse(w http.ResponseWriter, data interface{}) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}
