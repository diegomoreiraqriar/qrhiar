package routes

import (
	"qrhiar/internal/auth"
	"qrhiar/internal/handlers"

	"github.com/gofiber/fiber/v2"
)

// RegisterRoutes define todas as rotas da aplicação QRhiar
func RegisterRoutes(app *fiber.App) {
	// Inicializa o serviço de autenticação
	authService := auth.NewAuthService()

	// Rota pública de login
	app.Post("/auth/login", auth.LoginHandler(authService))

	// Rota pública de health check
	app.Get("/health", handlers.HealthHandler)

	// Agrupamento de rotas protegidas por JWT
	api := app.Group("/api", auth.AuthMiddleware(authService))

	// Companies
	api.Post("/companies", handlers.CreateCompany)
	api.Get("/companies", handlers.ListCompanies)
	api.Get("/companies/:id", handlers.GetCompany)
	api.Put("/companies/:id", handlers.UpdateCompany)
	api.Delete("/companies/:id", handlers.DeleteCompany)

	// Third-Party Users
	api.Post("/third-parties", handlers.CreateThirdPartyUser)
	api.Get("/third-parties", handlers.ListThirdPartyUsers)
	api.Get("/third-parties/:id", handlers.GetThirdPartyUser)
	api.Put("/third-parties/:id", handlers.UpdateThirdPartyUser)
	api.Delete("/third-parties/:id", handlers.DeleteThirdPartyUser)
	api.Get("/scim/v2/Users", handlers.ListSCIMUsers)
	api.Post("/scim/v2/Users", handlers.CreateSCIMUser)
	api.Patch("/scim/v2/Users/:id", handlers.PatchSCIMUser)
	api.Patch("/third-parties/:id/status", handlers.UpdateUserStatus)
	api.Get("/third-parties/:id/logs", handlers.GetUserLogsHandler)

	// Users (JML demo)
	api.Get("/users", handlers.ListUsers)
	api.Post("/users", handlers.CreateUser)
	api.Patch("/users/:id/status", handlers.UpdateUserStatus)
}
