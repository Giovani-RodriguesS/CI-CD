package model

type ProductDTO struct {
	Name     string  `json:"Name" binding:"required"`
	Price    float64 `json:"Price" binding:"required"`
	Quantity int     `json:"Quantity" binding:"required"`
}
