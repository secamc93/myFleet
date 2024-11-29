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
	dbConn := gorm.NewDBConnection()
	dbConn.GetDB()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	<-sigChan
	if err := dbConn.CloseDB(); err != nil {
		log.Error("Error closing database connection: %v", err)
	} else {
		log.Info("Database connection closed successfully")
	}
}
