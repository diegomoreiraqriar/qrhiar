package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Company struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	Name      string    `gorm:"size:255;not null" json:"name"`
	CNPJ      string    `gorm:"size:20;uniqueIndex;not null" json:"cnpj"`
	TenantID  string    `gorm:"size:36;not null" json:"tenant_id"`
	Status    string    `gorm:"default:'active'" json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Antes de criar, gera automaticamente um UUID se vazio
func (c *Company) BeforeCreate(tx *gorm.DB) (err error) {
	if c.ID == uuid.Nil {
		c.ID = uuid.New()
	}
	return
}
