// controllers/receipt_controller.go
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

func GenerateReceipt(c *gin.Context) {
	var order models.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		views.ErrorResponse(c, "Invalid request data")
		return
	}

	// Forward the order data to the Node.js microservice
	receiptData, err := json.Marshal(order)
	if err != nil {
		views.ErrorResponse(c, "Failed to generate receipt")
		return
	}

	resp, err := http.Post("http://localhost:3000/api/receipt", "application/json", bytes.NewBuffer(receiptData))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error contacting the receipt service"})
		return
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	c.Data(http.StatusOK, "application/json", body)
}
