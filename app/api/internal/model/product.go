package model

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name     string  `json:"Name"`
	Price    float64 `json:"Price"`
	Quantity int     `json:"Quantity"`
}
