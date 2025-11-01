package handlers

import (
	"qrhiar/internal/models"
	"qrhiar/internal/services"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// SCIMUser representa o formato SCIM simplificado
type SCIMUser struct {
	ID       string                 `json:"id"`
	UserName string                 `json:"userName"`
	Name     map[string]string      `json:"name"`
	Active   bool                   `json:"active"`
	Emails   []map[string]string    `json:"emails"`
	Ext      map[string]interface{} `json:"urn:ietf:params:scim:schemas:extension:enterprise:2.0:User"`
}

// GET /scim/v2/Users
func ListSCIMUsers(c *fiber.Ctx) error {
	users, err := services.ListThirdPartyUsers()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	var scimUsers []SCIMUser
	for _, u := range users {
		manager := map[string]interface{}{}

		// ✅ Exibe UUID e nome do gestor, se existir
		if u.Manager != nil {
			manager["value"] = u.Manager.ID.String()
			manager["displayName"] = u.Manager.Name
		} else if u.ManagerID != nil {
			manager["value"] = u.ManagerID.String()
			manager["displayName"] = "Desconhecido"
		}

		scimUsers = append(scimUsers, SCIMUser{
			ID:       u.ID.String(),
			UserName: u.Email,
			Name: map[string]string{
				"formatted": u.Name,
			},
			Active: u.Status == "active",
			Emails: []map[string]string{
				{"value": u.Email, "primary": "true"},
			},
			Ext: map[string]interface{}{
				"manager":    manager, // 👈 inclui value + displayName
				"department": u.Position,
				"companyId":  u.CompanyID.String(),
			},
		})
	}

	return c.JSON(fiber.Map{
		"Resources":    scimUsers,
		"totalResults": len(scimUsers),
		"startIndex":   1,
		"itemsPerPage": len(scimUsers),
	})
}

// POST /scim/v2/Users
func CreateSCIMUser(c *fiber.Ctx) error {
	var su SCIMUser
	if err := c.BodyParser(&su); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "JSON inválido"})
	}

	// Extrair companyId
	var companyID uuid.UUID
	if cid, ok := su.Ext["companyId"].(string); ok {
		if parsed, err := uuid.Parse(cid); err == nil {
			companyID = parsed
		}
	}

	// Extrair manager.value (UUID do gestor)
	var managerID *uuid.UUID
	if mgr, ok := su.Ext["manager"].(map[string]interface{}); ok {
		if val, ok := mgr["value"].(string); ok {
			if parsed, err := uuid.Parse(val); err == nil {
				managerID = &parsed
			}
		}
	}

	if managerID == nil {
		return c.Status(400).JSON(fiber.Map{"error": "manager.value (UUID) é obrigatório"})
	}

	// Extrair department/position com segurança
	var position string
	if dept, ok := su.Ext["department"].(string); ok {
		position = dept
	}

	user := models.ThirdPartyUser{
		Name:      su.Name["formatted"],
		Email:     su.UserName,
		CPF:       "000.000.000-00", // Placeholder
		Position:  position,
		Status:    "active",
		TenantID:  "default-tenant",
		CompanyID: companyID,
		ManagerID: managerID, // ✅ referência ao gestor
	}

	created, err := services.CreateThirdPartyUser(&user)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	su.ID = created.ID.String()

	return c.Status(201).JSON(su)
}

// PATCH /scim/v2/Users/:id
func PatchSCIMUser(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "UUID inválido"})
	}

	var payload map[string]interface{}
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "JSON inválido"})
	}

	user, err := services.GetThirdPartyUser(id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Usuário não encontrado"})
	}

	// Atualiza apenas campos enviados
	updatedFields := []string{}

	if email, ok := payload["email"].(string); ok {
		user.Email = email
		updatedFields = append(updatedFields, "email")
	}

	if status, ok := payload["status"].(string); ok {
		user.Status = status
		updatedFields = append(updatedFields, "status")
	}

	if position, ok := payload["position"].(string); ok {
		user.Position = position
		updatedFields = append(updatedFields, "position")
	}

	if mgr, ok := payload["manager_id"].(string); ok {
		if parsed, err := uuid.Parse(mgr); err == nil {
			user.ManagerID = &parsed
			updatedFields = append(updatedFields, "manager_id")
		}
	}

	if company, ok := payload["company_id"].(string); ok {
		if parsed, err := uuid.Parse(company); err == nil {
			user.CompanyID = parsed
			updatedFields = append(updatedFields, "company_id")
		}
	}

	if _, err := services.UpdateThirdPartyUser(user); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	message := "Usuário atualizado com sucesso"
	if len(updatedFields) > 0 {
		message = "Campos atualizados: " + joinFields(updatedFields)
	}

	return c.Status(200).JSON(fiber.Map{
		"message": message,
		"id":      id.String(),
		"user":    user,
	})
}

// joinFields formata os campos atualizados
func joinFields(fields []string) string {
	result := ""
	for i, f := range fields {
		if i > 0 {
			result += ", "
		}
		result += f
	}
	return result
}
