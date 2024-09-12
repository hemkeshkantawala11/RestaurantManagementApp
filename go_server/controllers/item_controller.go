package controllers

import (
	"net/http"
	"restaurantApp/go_server/database"
	helper "restaurantApp/go_server/helpers"
	"restaurantApp/go_server/models"
	"restaurantApp/go_server/views"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var itemCollection = database.OpenCollection(database.Client, "items")

func GetItems() gin.HandlerFunc {
	return func(c *gin.Context) {
		var items []models.Item
		cursor, err := itemCollection.Find(c, bson.M{})
		if err != nil {
			views.ErrorResponse(c, "Failed to fetch items")
			return
		}
		if err = cursor.All(c, &items); err != nil {
			views.ErrorResponse(c, "Failed to parse items")
			return
		}
		c.JSON(http.StatusOK, items)
	}
}

func GetItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		itemID := c.Param("id")
		objID, _ := primitive.ObjectIDFromHex(itemID)
		var item models.Item
		err := itemCollection.FindOne(c, bson.M{"_id": objID}).Decode(&item)
		if err != nil {
			views.ErrorResponse(c, "Item not found")
			return
		}
		c.JSON(http.StatusOK, item)
	}
}

func CreateItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := helper.CheckUserType(c, "admin")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		var item models.Item
		if err := c.ShouldBindJSON(&item); err != nil {
			views.ErrorResponse(c, "Invalid request data")
			return
		}
		item.ID = primitive.NewObjectID()
		item.ItemID = item.ID.Hex()
		_, err = itemCollection.InsertOne(c, item)
		if err != nil {
			views.ErrorResponse(c, "Failed to create item")
			return
		}
		c.JSON(http.StatusCreated, item)
	}
}

func UpdateItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		itemID := c.Param("id")
		objID, _ := primitive.ObjectIDFromHex(itemID)
		var item models.Item
		if err := c.ShouldBindJSON(&item); err != nil {
			views.ErrorResponse(c, "Invalid request data")
			return
		}
		_, err := itemCollection.UpdateOne(c, bson.M{"_id": objID}, bson.M{"$set": item})
		if err != nil {
			views.ErrorResponse(c, "Failed to update item")
			return
		}
		c.JSON(http.StatusOK, item)
	}
}

func DeleteItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		itemID := c.Param("id")
		objID, _ := primitive.ObjectIDFromHex(itemID)
		_, err := itemCollection.DeleteOne(c, bson.M{"_id": objID})
		if err != nil {
			views.ErrorResponse(c, "Failed to delete item")
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Item deleted successfully"})
	}
}
