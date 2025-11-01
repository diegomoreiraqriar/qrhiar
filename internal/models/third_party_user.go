package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// ThirdPartyUser representa um colaborador terceiro controlado pelo QRHiar
type ThirdPartyUser struct {
	ID       uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	Name     string    `gorm:"size:255;not null" json:"name"`
	Email    string    `gorm:"size:255;uniqueIndex" json:"email,omitempty"`
	CPF      string    `gorm:"size:14;uniqueIndex;not null" json:"cpf"`
	Position string    `gorm:"size:100" json:"position"`
	Status   string    `gorm:"default:'active'" json:"status"`

	StartDate *time.Time `json:"start_date,omitempty"`
	EndDate   *time.Time `json:"end_date,omitempty"`

	// ✅ Referência obrigatória ao gestor
	ManagerID *uuid.UUID      `gorm:"type:uuid;not null" json:"manager_id"`
	Manager   *ThirdPartyUser `gorm:"foreignKey:ManagerID" json:"manager,omitempty"`

	// ✅ Chaves de relacionamento organizacional
	TenantID  string    `gorm:"size:36;not null" json:"tenant_id"`
	CompanyID uuid.UUID `gorm:"type:uuid;not null" json:"company_id"`
	Company   Company   `gorm:"foreignKey:CompanyID" json:"company,omitempty"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// BeforeCreate garante que cada usuário tenha UUID único
func (u *ThirdPartyUser) BeforeCreate(tx *gorm.DB) (err error) {
	if u.ID == uuid.Nil {
		u.ID = uuid.New()
	}
	return
}
