package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `bson: "_id"`
	FirstName *string            `json: "firstName" validate: "required, min = 2, max = 100"`
	LastName  *string            `json: "lastName" validate: "required, min = 2, max = 100"`
	Email     *string            `json: "email" validate: " email, required"`
	Password  *string            `json: "password" validate: "required, min = 6"`
	User_type *string            `json: "user_type" validate: "required"` // USER, admin

	Token   *string `json: "token"`
	User_id string  `json: "user_id"`
}
