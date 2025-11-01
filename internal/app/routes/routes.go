package routes

import (
	"qrhiar/internal/handlers"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App) {
	app.Get("/health", handlers.HealthHandler)

	// Companies
	app.Post("/companies", handlers.CreateCompany)
	app.Get("/companies", handlers.ListCompanies)
	app.Get("/companies/:id", handlers.GetCompany)
	app.Put("/companies/:id", handlers.UpdateCompany)
	app.Delete("/companies/:id", handlers.DeleteCompany)

	// Third-Party Users
	app.Post("/third-parties", handlers.CreateThirdPartyUser)
	app.Get("/third-parties", handlers.ListThirdPartyUsers)
	app.Get("/third-parties/:id", handlers.GetThirdPartyUser)
	app.Put("/third-parties/:id", handlers.UpdateThirdPartyUser)
	app.Delete("/third-parties/:id", handlers.DeleteThirdPartyUser)
	app.Get("/scim/v2/Users", handlers.ListSCIMUsers)
	app.Post("/scim/v2/Users", handlers.CreateSCIMUser)
	app.Patch("/scim/v2/Users/:id", handlers.PatchSCIMUser)
	app.Patch("/third-parties/:id/status", handlers.UpdateUserStatus)
	app.Get("/third-parties/:id/logs", handlers.GetUserLogsHandler)

	// Users (JML demo)
	app.Get("/users", handlers.ListUsers)
	app.Post("/users", handlers.CreateUser)
	app.Patch("/users/:id/status", handlers.UpdateUserStatus)

}
