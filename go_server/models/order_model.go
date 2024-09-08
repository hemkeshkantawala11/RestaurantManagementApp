package models

type Order struct {
	CustomerName string  `json:"customer_name"`
	Items        []Item  `json:"items"`
	Total        float64 `json:"total"`
}

type Item struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}
