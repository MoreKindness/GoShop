package model

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Picture     string  `json:"picture"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}
