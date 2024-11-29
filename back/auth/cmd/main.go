package main

import (
	"my-fleet-auth/internal/infrastructure/secondary/gorm"
	"my-fleet-auth/pkg/env"
	"my-fleet-auth/pkg/logger"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	log := logger.NewLogger()

	env.LoadEnv()
	log.Info("Environment variables loaded successfully")

	dbConn := gorm.NewDBConnection()
	db := dbConn.GetDB()
	log.Info("Database connection established: %v", db)

	// Manejar señales del sistema para cerrar la conexión a la base de datos
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	<-sigChan
	log.Info("Shutting down gracefully...")
	if err := dbConn.CloseDB(); err != nil {
		log.Error("Error closing database connection: %v", err)
	} else {
		log.Info("Database connection closed successfully")
	}
}
