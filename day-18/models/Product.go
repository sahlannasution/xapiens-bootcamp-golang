package models

// Struct ProductModels
type ProductModels struct {
	ID    int64   `json:"id"`
	Name  string  `json:"name"`
	Info  string  `json:"info"`
	Price float64 `json:"price"`
}
