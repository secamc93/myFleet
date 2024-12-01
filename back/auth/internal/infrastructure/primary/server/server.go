package server

import (
	"fmt"
	"my-fleet-auth/internal/infrastructure/secondary/postgres"
	"my-fleet-auth/internal/infrastructure/secondary/postgres/migrate"
	"my-fleet-auth/pkg/env"
	"my-fleet-auth/pkg/logger"
	"net/http"
)

func RunSever() {
	log := logger.NewLogger()
	postgres.NewDBConnection().GetDB()
	defer postgres.NewDBConnection().CloseDB()
	migrate.Migrate()

	port := env.LoadEnv().ServerPort
	if port == "" {
		port = "8080"
	}

	address := fmt.Sprintf(":%s", port)
	log.Success("Server running on port %s", port)

	err := http.ListenAndServe(address, nil)
	if err != nil {
		log.Fatal("Failed to start server: %v", err)
	}

}
