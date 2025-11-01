package handlers

import (
	"qrhiar/internal/models"
	"qrhiar/internal/services"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func CreateThirdPartyUser(c *fiber.Ctx) error {
	var user models.ThirdPartyUser
	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "JSON inválido"})
	}

	if user.TenantID == "" {
		user.TenantID = "default-tenant"
	}

	created, err := services.CreateThirdPartyUser(&user)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(201).JSON(created)
}

func ListThirdPartyUsers(c *fiber.Ctx) error {
	users, err := services.ListThirdPartyUsers()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	var enriched []fiber.Map
	for _, u := range users {
		managerName := ""
		if u.ManagerID != nil {
			if manager, err := services.GetThirdPartyUser(*u.ManagerID); err == nil {
				managerName = manager.Name
			}
		}

		enriched = append(enriched, fiber.Map{
			"id":       u.ID,
			"name":     u.Name,
			"email":    u.Email,
			"position": u.Position,
			"status":   u.Status,
			"company": fiber.Map{
				"name": u.Company.Name,
			},
			"manager": fiber.Map{
				"displayName": managerName,
			},
		})
	}

	return c.JSON(enriched)
}

func GetThirdPartyUser(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "UUID inválido"})
	}

	user, err := services.GetThirdPartyUser(id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Usuário não encontrado"})
	}
	return c.JSON(user)
}

func UpdateThirdPartyUser(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "UUID inválido"})
	}

	var user models.ThirdPartyUser
	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "JSON inválido"})
	}

	user.ID = id

	updated, err := services.UpdateThirdPartyUser(&user)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(updated)
}

func DeleteThirdPartyUser(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "UUID inválido"})
	}

	if err := services.DeleteThirdPartyUser(id); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.SendStatus(204)
}
