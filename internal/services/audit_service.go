package services

import (
	"qrhiar/internal/db"
	"qrhiar/internal/models"
	"time"

	"github.com/google/uuid"
)

// CreateAuditLog registra uma ação no histórico de auditoria
func CreateAuditLog(userID uuid.UUID, action, reason, oldValue, newValue string) error {
	dbConn := db.GetDB()

	log := models.AuditLog{
		ID:        uuid.New(),
		UserID:    userID, // Agora tipo uuid.UUID, compatível com model
		Action:    action,
		Reason:    reason,
		OldValue:  oldValue,
		NewValue:  newValue,
		CreatedAt: time.Now(),
	}

	return dbConn.Create(&log).Error
}

// GetAuditLogsByUser retorna o histórico de auditoria de um usuário
func GetAuditLogsByUser(userID uuid.UUID) ([]models.AuditLog, error) {
	dbConn := db.GetDB()

	var logs []models.AuditLog
	err := dbConn.
		Where("user_id = ?", userID).
		Order("created_at DESC").
		Find(&logs).Error

	return logs, err
}
