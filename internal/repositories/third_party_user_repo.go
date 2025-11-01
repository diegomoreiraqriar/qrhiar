package repositories

import (
	"qrhiar/internal/db"
	"qrhiar/internal/models"

	"github.com/google/uuid"
)

// CreateThirdPartyUser cria um novo registro de usuário
func CreateThirdPartyUser(user *models.ThirdPartyUser) error {
	conn := db.GetDB()
	return conn.Create(user).Error
}

// GetAllThirdPartyUsers retorna todos os usuários com suas empresas e gestores
func GetAllThirdPartyUsers() ([]models.ThirdPartyUser, error) {
	var users []models.ThirdPartyUser
	err := db.GetDB().
		Preload("Company").
		Preload("Manager"). // 👈 carrega o manager
		Find(&users).Error
	return users, err
}

// GetThirdPartyUserByID retorna um usuário específico pelo UUID
func GetThirdPartyUserByID(id uuid.UUID) (*models.ThirdPartyUser, error) {
	conn := db.GetDB()
	var user models.ThirdPartyUser
	err := conn.
		Preload("Company").
		Preload("Manager").
		First(&user, "id = ?", id).Error
	return &user, err
}

// UpdateThirdPartyUser atualiza um registro existente
func UpdateThirdPartyUser(user *models.ThirdPartyUser) error {
	dbConn := db.GetDB()
	return dbConn.Save(user).Error
}


// DeleteThirdPartyUser remove um usuário pelo ID
func DeleteThirdPartyUser(id uuid.UUID) error {
	conn := db.GetDB()
	return conn.Delete(&models.ThirdPartyUser{}, "id = ?", id).Error
}
