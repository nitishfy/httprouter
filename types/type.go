package types

import "time"

type Cart struct {
	ID       int       `json:"id"`
	UserID   int       `json:"userId"`
	Date     time.Time `json:"date"`
	Products []Product `json:"products"`
}

type Product struct {
	ProductID int `json:"productId"`
	Quantity  int `json:"quantity"`
}
