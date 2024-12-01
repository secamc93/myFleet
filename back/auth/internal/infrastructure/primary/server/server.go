package server

import (
	"fmt"
	"my-fleet-auth/internal/infrastructure/secondary/postgres"
	"my-fleet-auth/internal/infrastructure/secondary/postgres/migrate"
	"my-fleet-auth/pkg/env"
	"my-fleet-auth/pkg/logger"
	"net/http"
	"time"
)

func RunSever() {
	log := logger.NewLogger()
	dbConn := postgres.NewDBConnection()
	defer func() {
		if err := dbConn.CloseDB(); err != nil {
			log.Error("Error closing DB: %v", err)
		}
	}()
	migrate.Migrate()

	port := env.LoadEnv().ServerPort
	if port == "" {
		port = "8080"
	}

	address := fmt.Sprintf(":%s", port)
	log.Success("Server running on port %s", port)

	server := &http.Server{
		Addr:         address,
		Handler:      nil,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Error("Failed to start server: %v", err)
	}

}
