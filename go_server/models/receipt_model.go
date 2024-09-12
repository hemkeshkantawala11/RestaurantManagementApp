package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Receipt struct {
	ID      primitive.ObjectID `bson:"_id,omitempty"`
	OrderID string             `json:"order_id"`
	Items   []Item             `json:"items"`
	Total   float64            `json:"total"`
}
