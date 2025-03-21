package main

import "github.com/Giovani-RodriguesS/CI-CD/app/api/internal/initializers"
import "github.com/Giovani-RodriguesS/CI-CD/app/api/internal/model"

func init() {
	initializers.LoadEnvVars()
	initializers.ConnectDB()
}

func main() {
	initializers.DB.AutoMigrate(&model.Product{})
}
