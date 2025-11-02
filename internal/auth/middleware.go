package auth

import (
	"strings"

	"github.com/gofiber/fiber/v2"
)

// AuthMiddleware protege rotas exigindo um JWT válido
func AuthMiddleware(authService *AuthService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Token ausente"})
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		token, err := authService.ValidateToken(tokenString)
		if err != nil || !token.Valid {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "Token inválido ou expirado"})
		}

		return c.Next()
	}
}
