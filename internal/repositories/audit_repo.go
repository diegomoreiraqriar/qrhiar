package repositories

import (
	"qrhiar/internal/db"
	"qrhiar/internal/models"
)

// Cria um log de auditoria
func CreateAuditLog(log *models.AuditLog) error {
	database := db.GetDB()
	return database.Create(log).Error
}

// Busca logs de um usuário específico
func GetAuditLogsByUserID(userID string) ([]models.AuditLog, error) {
	database := db.GetDB()
	var logs []models.AuditLog
	err := database.
		Where("user_id = ?", userID).
		Order("created_at DESC").
		Find(&logs).Error
	return logs, err
}
