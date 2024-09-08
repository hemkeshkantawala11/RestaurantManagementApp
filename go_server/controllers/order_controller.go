// controllers/order_controller.go
package controllers

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"restaurantApp/go_server/models"
	"restaurantApp/go_server/views"

	"github.com/gin-gonic/gin"
)

func ProcessOrder(c *gin.Context) {
	var order models.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		views.ErrorResponse(c, "Invalid request data")
		return
	}

	// Forward the order to the Python microservice
	orderData, err := json.Marshal(order)
	if err != nil {
		views.ErrorResponse(c, "Failed to process order")
		return
	}

	resp, err := http.Post("http://localhost:5000/api/order", "application/json", bytes.NewBuffer(orderData))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error contacting the order service"})
		return
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	c.Data(http.StatusOK, "application/json", body)
}
