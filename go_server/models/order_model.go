package models

type Order struct {
	Items []string `bson:"Items"`
	Total float64  `bson:"total"`
}
