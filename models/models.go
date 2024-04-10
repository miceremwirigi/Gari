package models

var Cars []Car

type Car struct {
	ID    int     `json:"id"`
	Model string  `json:"model"`
	Color string  `json:"color"`
	Price float64 `json:"price"`
}
