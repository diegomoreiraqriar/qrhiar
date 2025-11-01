package handlers

import (
	"log"
	"qrhiar/internal/services"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// StatusRequest define o payload da requisição
type StatusRequest struct {
	Action string `json:"action"`
	Reason string `json:"reason,omitempty"`
}

// UpdateUserStatus altera o status do usuário e registra a ação no log
func UpdateUserStatus(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "UUID inválido"})
	}

	var req StatusRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "JSON inválido"})
	}

	user, err := services.GetThirdPartyUser(id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Usuário não encontrado"})
	}

	oldStatus := user.Status

	// Define o novo status baseado na ação
	switch req.Action {
	case "block":
		user.Status = "blocked"
	case "activate":
		user.Status = "active"
	case "leave":
		user.Status = "leave"
	case "inactive":
		user.Status = "inactive"
	case "vacation":
		user.Status = "vacation"
	case "terminate":
		user.Status = "terminated"
	default:
		return c.Status(400).JSON(fiber.Map{"error": "Ação inválida"})
	}

	// Atualiza o usuário
	updatedUser, err := services.UpdateThirdPartyUser(user)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error":   "Erro ao atualizar usuário",
			"details": err.Error(),
		})
	}

	// Registra o log de auditoria, mas sem interromper a resposta em caso de erro
	if logErr := services.CreateAuditLog(
		user.ID,     // tipo uuid.UUID
		req.Action,  // ação executada
		req.Reason,  // motivo (opcional)
		oldStatus,   // valor anterior
		user.Status, // valor atual
	); logErr != nil {
		log.Printf("⚠️ Falha ao registrar log de auditoria: %v", logErr)
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "Status atualizado com sucesso",
		"user":    updatedUser,
	})
}
