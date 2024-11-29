package gorm

import (
	"fmt"
	"my-fleet-auth/pkg/env"
	"my-fleet-auth/pkg/logger"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBConnection interface {
	GetDB() *gorm.DB
	CloseDB() error
	Reconnect() error
	PingDB() error
}

type dbConnection struct {
	db *gorm.DB
}

var (
	instance *dbConnection
	once     sync.Once
	log      = logger.NewLogger()
)

func NewDBConnection() DBConnection {
	once.Do(func() {
		envVars := env.LoadEnv()
		dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
			envVars.DBHost, envVars.DBPort, envVars.DBUser, envVars.DBName, envVars.DBPassword)
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatal("failed to connect database: %v", err)
			panic("failed to connect database")
		}
		log.Info("Database connection established")
		instance = &dbConnection{db: db}
	})
	return instance
}

func (conn *dbConnection) GetDB() *gorm.DB {
	if err := conn.PingDB(); err != nil {
		log.Warn("Database connection lost, attempting to reconnect")
		if err := conn.Reconnect(); err != nil {
			log.Fatal("Failed to reconnect to the database: %v", err)
			panic("Failed to reconnect to the database")
		}
	}
	return conn.db
}

func (conn *dbConnection) CloseDB() error {
	sqlDB, err := conn.db.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}

func (conn *dbConnection) Reconnect() error {
	envVars := env.LoadEnv()
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		envVars.DBHost, envVars.DBPort, envVars.DBUser, envVars.DBName, envVars.DBPassword)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to reconnect database: %v", err)
		return err
	}
	conn.db = db
	log.Info("Database reconnected successfully")
	return nil
}

func (conn *dbConnection) PingDB() error {
	sqlDB, err := conn.db.DB()
	if err != nil {
		return err
	}
	return sqlDB.Ping()
}
