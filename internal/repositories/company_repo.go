package repositories

import (
	"qrhiar/internal/db"
	"qrhiar/internal/models"

	"github.com/google/uuid"
)

func CreateCompany(company *models.Company) error {
	return db.GetDB().Create(company).Error
}

func GetAllCompanies() ([]models.Company, error) {
	var companies []models.Company
	err := db.GetDB().Find(&companies).Error
	return companies, err
}

func GetCompanyByID(id uuid.UUID) (*models.Company, error) {
	var company models.Company
	err := db.GetDB().First(&company, "id = ?", id).Error
	return &company, err
}

func UpdateCompany(company *models.Company) error {
	return db.GetDB().Save(company).Error
}

func DeleteCompany(id uuid.UUID) error {
	return db.GetDB().Delete(&models.Company{}, "id = ?", id).Error
}
