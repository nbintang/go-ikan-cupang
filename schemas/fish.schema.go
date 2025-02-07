package schemas

type FishSchema struct {
	Name        string  `json:"name" validate:"required"`
	Description *string `json:"description"`
	Price       float64 `json:"price" validate:"required"`
	Stock       int     `json:"stock" validate:"required"`
	Category    string  `json:"category" validate:"required"`
}
