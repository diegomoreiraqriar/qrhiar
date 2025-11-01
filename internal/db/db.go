package db

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var database *gorm.DB

// InitDB conecta ao Postgres e armazena a instância global
func InitDB() {
	host := os.Getenv("POSTGRES_HOST")
	if host == "" {
		host = "127.0.0.1"
	}
	user := os.Getenv("POSTGRES_USER")
	if user == "" {
		user = "qrhiar"
	}
	password := os.Getenv("POSTGRES_PASSWORD")
	if password == "" {
		password = "qrhiar123"
	}
	dbname := os.Getenv("POSTGRES_DB")
	if dbname == "" {
		dbname = "qrhiar_db"
	}
	port := os.Getenv("POSTGRES_PORT")
	if port == "" {
		port = "5432"
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host, user, password, dbname, port,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("❌ Erro ao conectar ao banco: %v", err)
	}

	database = db
	log.Println("✅ Banco conectado com sucesso!")
}

// GetDB retorna a instância global do GORM
func GetDB() *gorm.DB {
	if database == nil {
		log.Fatal("Banco de dados não inicializado. Chame InitDB() antes.")
	}
	return database
}
