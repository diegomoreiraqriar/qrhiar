package handlers

import (
	"qrhiar/internal/services"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// GetUserLogsHandler retorna o histórico de auditoria de um usuário
func GetUserLogsHandler(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "UUID inválido"})
	}

	logs, err := services.GetAuditLogsByUser(id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Erro ao buscar logs de auditoria", "details": err.Error()})
	}

	return c.Status(200).JSON(fiber.Map{
		"user_id": id,
		"logs":    logs,
	})
}
