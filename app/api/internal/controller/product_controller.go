package controller

import (
	"github.com/Giovani-RodriguesS/CI-CD/app/api/internal/model"
	"github.com/Giovani-RodriguesS/CI-CD/app/api/internal/usecase"
	"github.com/Giovani-RodriguesS/CI-CD/app/api/internal/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ProductController struct {
	productUsecase usecase.ProductUsecase
}

func NewProductController(usecase usecase.ProductUsecase) ProductController {
	return ProductController{
		productUsecase: usecase,
	}
}

func (p *ProductController) GetProducts(ctx *gin.Context) {

	products, err := p.productUsecase.GetProducts()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, products)
}

func (p *ProductController) GetProductById(ctx *gin.Context) {
	id, err := utils.IdValidate(ctx)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	product, err := p.productUsecase.GetProductById(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	empty := model.Product{}
	if product == empty {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "Produto não encontrado",
		})
		return
	}

	ctx.JSON(http.StatusOK, product)

}

func (p *ProductController) CreateProduct(ctx *gin.Context) {
	var productDTO model.ProductDTO

	if err := ctx.ShouldBindJSON(&productDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	id, err := p.productUsecase.CreateProduct(&productDTO)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"Id": id,
	})
}

func (p *ProductController) UpdateProduct(ctx *gin.Context) {
	var productDTO model.ProductDTO
	id, err := utils.IdValidate(ctx)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err = ctx.ShouldBindJSON(&productDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = p.productUsecase.UpdateProduct(id, &productDTO)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusAccepted, gin.H{"message": "Produto atualizado"})

}

func (p *ProductController) DeleteProduct(ctx *gin.Context) {
	id, err := utils.IdValidate(ctx)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = p.productUsecase.DeleteProduct(id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Produto deletado",
	})
}
