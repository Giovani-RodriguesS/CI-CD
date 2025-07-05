package router

import (
	"github.com/Giovani-RodriguesS/CI-CD/app/api/internal/controller"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Router(productController *controller.ProductController) *gin.Engine {

	router := gin.Default()
	router.Use(cors.Default())
	router.GET("/products", productController.GetProducts)
	router.GET("/products/:Id", productController.GetProductById)
	router.POST("/products", productController.CreateProduct)
	router.PUT("/products/:Id", productController.UpdateProduct)
	router.DELETE("/products/:Id", productController.DeleteProduct)

	return router
}
