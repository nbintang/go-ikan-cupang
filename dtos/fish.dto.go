package dtos



type Fish struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
	Image       *string
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
}