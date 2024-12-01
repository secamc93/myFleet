package migrate

import (
	"my-fleet-auth/internal/infrastructure/secondary/postgres"
	"my-fleet-auth/internal/infrastructure/secondary/postgres/models"
	"my-fleet-auth/pkg/logger"
)

func Migrate() {
	log := logger.NewLogger()
	db := postgres.NewDBConnection().GetDB()
	err := db.AutoMigrate(
		&models.User{},
		&models.Company{})
	if err != nil {
		log.Fatal("Failed to migrate database: " + err.Error())
	}
	log.Success("Database migrated")
}
