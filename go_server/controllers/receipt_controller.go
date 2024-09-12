package controllers

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"restaurantApp/go_server/database"
	"restaurantApp/go_server/models"
	"restaurantApp/go_server/views"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var orderCollection = database.OpenCollection(database.Client, "orders")

func GenerateReceipt(c *gin.Context) {
	orderId := c.Param("orderId")

	objID, err := primitive.ObjectIDFromHex(orderId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order ID"})
		return
	}

	var order models.Order
	err = orderCollection.FindOne(c, bson.M{"_id": objID}).Decode(&order)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Order not found"})
		return
	}

	var items []models.Item
	for _, itemID := range order.Items {
		var item models.Item
		objID, _ := primitive.ObjectIDFromHex(itemID)
		err := itemCollection.FindOne(c, bson.M{"_id": objID}).Decode(&item)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Item not found: " + itemID})
			return
		}
		items = append(items, item)
	}

	receipt := models.Receipt{
		OrderID: orderId,
		Items:   items,
		Total:   order.Total,
	}

	receiptData, err := json.Marshal(receipt)
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
