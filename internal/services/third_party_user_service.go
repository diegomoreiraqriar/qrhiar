package services

import (
	"errors"
	"qrhiar/internal/models"
	"qrhiar/internal/repositories"

	"github.com/google/uuid"
)

// CreateThirdPartyUser cria um novo usuário de terceiros com validações
func CreateThirdPartyUser(user *models.ThirdPartyUser) (*models.ThirdPartyUser, error) {
	// Nome e CPF são obrigatórios
	if user.Name == "" || user.CPF == "" {
		return nil, errors.New("nome e CPF são obrigatórios")
	}

	// E-mail pode ser vazio — será preenchido posteriormente pelo ISC
	if user.Email == "" {
		user.Email = ""
	}

	// Manager é obrigatório (ref por UUID)
	if user.ManagerID == nil {
		return nil, errors.New("manager_id é obrigatório")
	}

	// Tenant padrão (enquanto não houver multi-tenant)
	if user.TenantID == "" {
		user.TenantID = "default-tenant"
	}

	// Verifica se o manager existe antes de criar
	if _, err := repositories.GetThirdPartyUserByID(*user.ManagerID); err != nil {
		return nil, errors.New("manager_id inválido — gestor não encontrado")
	}

	err := repositories.CreateThirdPartyUser(user)
	return user, err
}

// ListThirdPartyUsers retorna todos os usuários de terceiros
func ListThirdPartyUsers() ([]models.ThirdPartyUser, error) {
	return repositories.GetAllThirdPartyUsers()
}

// GetThirdPartyUser busca um usuário pelo ID
func GetThirdPartyUser(id uuid.UUID) (*models.ThirdPartyUser, error) {
	return repositories.GetThirdPartyUserByID(id)
}

// UpdateThirdPartyUser atualiza os dados de um usuário existente
func UpdateThirdPartyUser(user *models.ThirdPartyUser) (*models.ThirdPartyUser, error) {
	// Validação: o status pode estar vazio em alguns casos — forçamos manter o anterior
	if user.Status == "" {
		existing, err := repositories.GetThirdPartyUserByID(user.ID)
		if err == nil && existing != nil {
			user.Status = existing.Status
		}
	}

	// Executa o update
	err := repositories.UpdateThirdPartyUser(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// DeleteThirdPartyUser remove um usuário pelo ID
func DeleteThirdPartyUser(id uuid.UUID) error {
	return repositories.DeleteThirdPartyUser(id)
}

// UpdateUserStatus altera o status de um usuário (Joiner, Mover, Leaver etc.)
func UpdateUserStatus(id uuid.UUID, newStatus string) (*models.ThirdPartyUser, error) {
	user, err := repositories.GetThirdPartyUserByID(id)
	if err != nil {
		return nil, errors.New("usuário não encontrado")
	}

	validStatuses := map[string]bool{
		"active":     true,
		"blocked":    true,
		"on_leave":   true,
		"terminated": true,
		"rehired":    true,
	}

	if !validStatuses[newStatus] {
		return nil, errors.New("status inválido — use active, blocked, on_leave, terminated ou rehired")
	}

	user.Status = newStatus

	if err := repositories.UpdateThirdPartyUser(user); err != nil {
		return nil, err
	}

	return user, nil
}
