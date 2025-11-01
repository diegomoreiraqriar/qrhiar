package main

import (
	"log"
	"qrhiar/internal/app"
	"qrhiar/internal/db"

	"github.com/google/uuid"
)

func main() {
	// db.Connect()
	db.InitDB()
	db.Migrate()

	app := app.NewServer()
	log.Println("🚀 QRHiar API rodando em http://localhost:4000")

	if err := app.Listen(":4000"); err != nil {
		log.Fatalf("❌ Erro ao iniciar o servidor: %v", err)
	}
}

func parseUUID(id string) uuid.UUID {
	u, err := uuid.Parse(id)
	if err != nil {
		return uuid.Nil
	}
	return u
}
