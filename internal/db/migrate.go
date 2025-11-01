package db

import (
	"log"
	"qrhiar/internal/models"
)

func Migrate() {
	db := GetDB()

	err := db.AutoMigrate(
		&models.Company{},
		&models.ThirdPartyUser{},
		&models.AuditLog{}, // ✅ adicionamos aqui
	)

	if err != nil {
		log.Fatalf("❌ Erro ao migrar modelos: %v", err)
	} else {
		log.Println("✅ Migração concluída com sucesso!")
	}
}
