//go:build wireinject
// +build wireinject

package userService

import (
	"github.com/AFORANURAG/microServices/backend/userService/userRepository"

	dbservice "github.com/AFORANURAG/microServices/backend/userService/services/dbService"

	"github.com/google/wire"
)

func InitializeUserService(phrase string) *UserServiceImpl {
	wire.Build(NewUserServiceProvider, userRepository.NewUserRepositoryProvider, dbservice.NewDBServiceClientProvider)
	return &UserServiceImpl{}
}
