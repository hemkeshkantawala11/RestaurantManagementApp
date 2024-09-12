package controllers

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"io/ioutil"
	"net/http"
	"restaurantApp/go_server/models"
	"restaurantApp/go_server/views"
)

func ProcessOrder(c *gin.Context) {
	var order models.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		views.ErrorResponse(c, "Invalid request data")
		return
	}

	var total float64
	for _, itemID := range order.Items {
		var foundItem models.Item
		err := itemCollection.FindOne(c, bson.M{"item_id": itemID}).Decode(&foundItem)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Item unavailable: " + itemID})
			return
		}
		total += foundItem.Price
	}

	order.Total = total

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
