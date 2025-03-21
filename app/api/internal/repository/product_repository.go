package repository

import (
	"errors"
	"github.com/Giovani-RodriguesS/CI-CD/app/api/internal/model"
	"gorm.io/gorm"
)

type ProductRepository struct {
	connection *gorm.DB
}

func NewProductRepository(connection *gorm.DB) ProductRepository {
	return ProductRepository{
		connection: connection,
	}
}

func (pr *ProductRepository) GetProducts() ([]model.Product, error) {
	var products []model.Product

	// Passando o ponteiro para o slice para o GORM preencher com os resultados
	result := pr.connection.Find(&products)

	if result.Error != nil {
		return []model.Product{}, result.Error
	}

	return products, nil
}

func (pr *ProductRepository) GetProductById(id uint) (model.Product, error) {
	var product model.Product
	result := pr.connection.First(&product, id)

	if result.Error != nil {
		return model.Product{}, result.Error
	}

	return product, nil
}

func (pr *ProductRepository) CreateProduct(product *model.Product) (uint, error) {
	result := pr.connection.Create(product)

	if result.Error != nil {
		return 0, result.Error
	}

	// Sucesso ao criar
	return product.ID, nil

}

func (pr *ProductRepository) UpdateProduct(id uint, product model.Product) error {

	result := pr.connection.Model(&model.Product{}).Where("id = ?", id).Updates(product)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected < 1 {
		return errors.New("produto não encontrado")
	}

	return nil
}

func (pr *ProductRepository) DeleteProduct(id uint) error {

	result := pr.connection.Unscoped().Delete(&model.Product{}, id)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil

}
