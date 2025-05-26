package initializers

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnvVars() {
	value := os.Getenv("ENV")

	// Não é local
	if value != "local" {
		return
	}

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Erro ao carregar variaveis de ambiente")
	}
}
