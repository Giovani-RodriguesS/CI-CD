package initializers

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB

func ConnectDB() {
	var err error
	initConn := os.Getenv("DB_URL")

	// Conectando ao banco de dados para ver se há base de dados
	DB, err = gorm.Open(postgres.Open(initConn), &gorm.Config{})

	if err != nil {
		log.Fatal("Erro ao conectar ao banco de dados")
	}

	selectedDb := os.Getenv("DB_SELECTED")

	// Verificando existencia da base de dados
	initQuery := fmt.Sprintf("SELECT 1 FROM pg_database WHERE datname = '%s'", selectedDb)
	var exists bool
	DB.Raw(initQuery).Scan(&exists)

	if !exists {
		initQuery := fmt.Sprintf("CREATE DATABASE %s", selectedDb)
		DB.Exec(initQuery)
		fmt.Println("Banco de dados criado com sucesso!")
	} else {
		fmt.Println("Banco de dados já existe.")
	}

	psqlInfo := fmt.Sprintf("%s dbname=%s", initConn, selectedDb)

	// Estabelecendo conexão
	DB, err = gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})

	if err != nil {
		log.Fatal("Erro ao conectar ao banco de dados")
	}

	fmt.Println("Conectado com sucesso!")

}
