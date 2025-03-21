package usecase

import (
	"errors"
	"fmt"
	"github.com/Giovani-RodriguesS/CI-CD/app/api/internal/model"
	"github.com/Giovani-RodriguesS/CI-CD/app/api/internal/repository"
	"github.com/Giovani-RodriguesS/CI-CD/app/api/internal/utils"
	"gorm.io/gorm"
)

type ProductUsecase struct {
	repository repository.ProductRepository
}

func NewProductUseCase(repo repository.ProductRepository) ProductUsecase {
	return ProductUsecase{
		repository: repo,
	}
}

func (pu *ProductUsecase) GetProducts() ([]model.Product, error) {
	products, err := pu.repository.GetProducts()

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = fmt.Errorf("nenhum registro encontrado")
		}
		return nil, err
	}

	return products, nil
}

func (pu *ProductUsecase) GetProductById(id uint) (model.Product, error) {
	product, err := pu.repository.GetProductById(id)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = fmt.Errorf("nenhum registro encontrado")
		}
		return model.Product{}, err
	}

	return product, nil
}

func (pu *ProductUsecase) CreateProduct(productDTO *model.ProductDTO) (uint, error) {
	product := utils.ConvertProductDTO(productDTO)

	product_id, err := pu.repository.CreateProduct(product)

	if err != nil {
		err = fmt.Errorf("erro ao inserir no banco de dados")
		return 0, err
	}

	return product_id, nil
}

func (pu *ProductUsecase) UpdateProduct(id uint, productDTO *model.ProductDTO) error {
	uId := uint(id)
	product := utils.ConvertProductDTO(productDTO)

	err := pu.repository.UpdateProduct(uId, *product)
	if err != nil {
		return err
	}

	return nil
}

func (pu *ProductUsecase) DeleteProduct(id uint) error {
	uId := uint(id)
	err := pu.repository.DeleteProduct(uId)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = fmt.Errorf("produto não encontrado")
			return err
		}
		err = fmt.Errorf("erro ao deletar no banco de dados")
		return err
	}

	return nil
}
