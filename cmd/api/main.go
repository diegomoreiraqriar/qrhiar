package main

import (
	"log"
	"os"

	"qrhiar/internal/app/routes"
	"qrhiar/internal/db"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()
	db.InitDB()

	app := fiber.New()

	// ✅ Middleware CORS
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:5173", // endereço do seu front
		AllowMethods:     "GET,POST,PUT,PATCH,DELETE,OPTIONS",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowCredentials: true,
	}))

	// Rotas
	routes.RegisterRoutes(app)

	port := os.Getenv("PORT")
	if port == "" {
		port = "4000"
	}

	log.Printf("🚀 QRhiar API rodando na porta %s", port)
	if err := app.Listen(":" + port); err != nil {
		log.Fatalf("Erro ao iniciar servidor: %v", err)
	}
}
