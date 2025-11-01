package handlers

import (
	"qrhiar/internal/models"
	"qrhiar/internal/services"

	"github.com/gofiber/fiber/v2"
)

// GET /users
func ListUsers(c *fiber.Ctx) error {
	users, err := services.ListThirdPartyUsers()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(users)
}

// POST /users
func CreateUser(c *fiber.Ctx) error {
	var user models.ThirdPartyUser
	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "JSON inválido"})
	}

	created, err := services.CreateThirdPartyUser(&user)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(201).JSON(created)
}

// PATCH /users/:id/status
