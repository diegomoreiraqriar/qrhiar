package handlers

import (
	"qrhiar/internal/models"
	"qrhiar/internal/services"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func CreateCompany(c *fiber.Ctx) error {
	var company models.Company
	if err := c.BodyParser(&company); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "JSON inválido"})
	}

	if company.TenantID == "" {
		company.TenantID = "default-tenant"
	}

	created, err := services.CreateCompany(&company)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(201).JSON(created)
}

func ListCompanies(c *fiber.Ctx) error {
	companies, err := services.ListCompanies()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(companies)
}

func GetCompany(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "UUID inválido"})
	}

	company, err := services.GetCompany(id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Empresa não encontrada"})
	}
	return c.JSON(company)
}

func UpdateCompany(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "UUID inválido"})
	}

	var company models.Company
	if err := c.BodyParser(&company); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "JSON inválido"})
	}

	company.ID = id

	updated, err := services.UpdateCompany(&company)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(updated)
}

func DeleteCompany(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "UUID inválido"})
	}

	if err := services.DeleteCompany(id); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.SendStatus(204)
}
