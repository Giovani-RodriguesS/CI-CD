package main

import (
	"github.com/Giovani-RodriguesS/CI-CD/app/api/internal/controller"
	"github.com/Giovani-RodriguesS/CI-CD/app/api/internal/initializers"
	"github.com/Giovani-RodriguesS/CI-CD/app/api/internal/repository"
	"github.com/Giovani-RodriguesS/CI-CD/app/api/internal/router"
	"github.com/Giovani-RodriguesS/CI-CD/app/api/internal/usecase"
)

func init() {
	initializers.LoadEnvVars()
}

func main() {

	initializers.ConnectDB()
	// cria um repositório que interage com o banco de dados.
	ProductRepository := repository.NewProductRepository(initializers.DB)
	// implementa a lógica de negócio usando o repositório.
	ProductUsecase := usecase.NewProductUseCase(ProductRepository)
	// cria um controlador que expõe a lógica de negócio via endpoints HTTP.
	ProductController := controller.NewProductController(ProductUsecase)
	  server := router.Router(&ProductController)

	server.Run()
}
