package userRepository

import (
	userSchema "backend/userService/models"
	dbservice "backend/userService/services/dbService"
)

type IUserRepository interface {
	GetUserByName(name string) (*userSchema.Response, error)
	GetUserByRowId(rowID string) (*userSchema.Response, error)
}

type UserRepository struct {
	db *dbservice.MYSQLDBService
}

func (s *UserRepository) GetUserByName(name string) (*userSchema.Response, error) {
	return &userSchema.Response{
		Name: "Anurag",
	}, nil
}

func (s *UserRepository) GetUserByRowId(rowID string) (*userSchema.Response, error) {
	return &userSchema.Response{
		Name: "Anurag",
	}, nil
}

func NewUserRepositoryProvider(dbS *dbservice.MYSQLDBService) *UserRepository {
	return &UserRepository{db: dbS}
}