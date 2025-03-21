package utils

import "github.com/Giovani-RodriguesS/CI-CD/app/api/internal/model"

func ConvertProductDTO(pdto *model.ProductDTO) *model.Product {
	p := model.Product{
		Name:     pdto.Name,
		Price:    pdto.Price,
		Quantity: pdto.Quantity,
	}

	return &p
}
