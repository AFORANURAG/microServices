//go:build wireinject
// +build wireinject

package userService

import (
	dbservice "backend/userService/services/dbService"
	"backend/userService/userRepository"

	"github.com/google/wire"
)

func InitializeUserService(phrase string) *UserServiceImpl {
	wire.Build(NewUserServiceProvider, userRepository.NewUserRepositoryProvider, dbservice.NewDBServiceClientProvider)
	return &UserServiceImpl{}
}
