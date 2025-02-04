package schemas





type CreateFishSchema struct {
	Name        string  `json:"name" validate:"required"`
	Description *string `json:"description"`
	Price       float64 `json:"price" validate:"required"`
	Stock       int     `json:"stock" validate:"required"`
	Image       *string  
	Category    Category `json:"category" validate:"required"`
}


type Category struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}