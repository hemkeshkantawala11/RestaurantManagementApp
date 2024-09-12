package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Item struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	ItemID      string             `bson:"item_id" json:"item_id" validate:"required"`
	Name        string             `bson:"name" json:"name" validate:"required"`
	Description string             `bson:"description" json:"description"`
	Price       float64            `bson:"price" json:"price" validate:"required"`
}
