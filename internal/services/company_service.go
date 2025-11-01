package services

import (
	"errors"
	"qrhiar/internal/models"
	"qrhiar/internal/repositories"

	"github.com/google/uuid"
)

func CreateCompany(company *models.Company) (*models.Company, error) {
	if company.Name == "" || company.CNPJ == "" {
		return nil, errors.New("nome e CNPJ são obrigatórios")
	}
	err := repositories.CreateCompany(company)
	return company, err
}

func ListCompanies() ([]models.Company, error) {
	return repositories.GetAllCompanies()
}

func GetCompany(id uuid.UUID) (*models.Company, error) {
	return repositories.GetCompanyByID(id)
}

func UpdateCompany(company *models.Company) (*models.Company, error) {
	err := repositories.UpdateCompany(company)
	return company, err
}

func DeleteCompany(id uuid.UUID) error {
	return repositories.DeleteCompany(id)
}
