package repository

import (
	"my-fleet-auth/internal/domain/users/dtos"
	"my-fleet-auth/internal/domain/users/entities"
	"my-fleet-auth/internal/infrastructure/secondary/postgres"
	"sync"
)

type UserRepository interface {
	GetUsers(CompanyID uint) ([]dtos.UserDTO, error)
	CreateUser(user entities.Users) (bool, error)
}

type repository struct {
	dbConnection postgres.DBConnection
}

var (
	instance *repository
	once     sync.Once
)

func NewRepository(db postgres.DBConnection) UserRepository {
	once.Do(func() {
		instance = &repository{
			dbConnection: db,
		}
	})
	return instance
}

func (r *repository) GetUsers(CompanyID uint) ([]dtos.UserDTO, error) {
	return nil, nil
}

func (r *repository) CreateUser(user entities.Users) (bool, error) {
	return false, nil
}
