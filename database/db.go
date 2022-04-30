package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/mipimentel/gin-api-rest/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConectaComBancoDeDados() {
	err := godotenv.Load("development.env")
	if err != nil {
		log.Panic("Error loading .env file")
	}

	pg_user := os.Getenv("POSTGRES_USER")
	pg_db := os.Getenv("POSTGRES_DB")
	pg_pass := os.Getenv("POSTGRES_PASSWORD")
	pg_host := os.Getenv("POSTGRES_HOST")

	dsn := fmt.Sprintf("user=%s dbname=%s password=%s "+
		"host=%s port=5432 sslmode=disable", pg_user, pg_db, pg_pass, pg_host)

	DB, err = gorm.Open(postgres.Open(dsn))

	if err != nil {
		log.Panic("Erro ao conectar com banco de dados")
	}

	DB.AutoMigrate(&models.Aluno{})

}
