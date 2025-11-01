package app

import (
	"github.com/gofiber/fiber/v2"

	"qrhiar/internal/app/routes"
	"qrhiar/internal/logger"
	"qrhiar/internal/tenant"

	"github.com/gofiber/fiber/v2/middleware/cors"
)

func NewServer() *fiber.App {
	// Inicializa Fiber
	app := fiber.New()

	// Inicializa o logger global (Zap)
	log := logger.NewLogger()
	log.Info("🚀 Iniciando QRHiar API com middleware multi-tenant")

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173", // ou "*" se quiser abrir geral
		AllowMethods: "GET,POST,PUT,PATCH,DELETE,OPTIONS",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))

	// Aplica middleware de Tenant (lê X-Tenant-ID ou usa default)
	app.Use(tenant.Middleware(log))

	// Registra rotas
	routes.RegisterRoutes(app)

	// Mensagem de inicialização
	log.Info("✅ Rotas registradas e middleware ativo")

	return app
}
