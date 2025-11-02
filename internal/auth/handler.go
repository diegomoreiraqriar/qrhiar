package auth

import (
	"os"

	"github.com/gofiber/fiber/v2"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// LoginHandler gera o token JWT para autenticação
func LoginHandler(authService *AuthService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req LoginRequest
		if err := c.BodyParser(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Requisição inválida",
			})
		}

		if req.Email != os.Getenv("ADMIN_USER") || req.Password != os.Getenv("ADMIN_PASS") {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Credenciais inválidas",
			})
		}

		token, err := authService.GenerateToken(req.Email)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Erro ao gerar token",
			})
		}

		return c.JSON(fiber.Map{"token": token})
	}
}
