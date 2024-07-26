//go:build wireinject
// +build wireinject

package userService

import (
	"github.com/AFORANURAG/microServices/authenticationService/userRepository"

	dbservice "github.com/AFORANURAG/microServices/authenticationService/services/dbService"

	"github.com/google/wire"
)

func InitializeUserService(phrase string) UserService {
	wire.Build(NewUserServiceProvider, userRepository.NewUserRepositoryProvider, dbservice.NewDBServiceClientProvider)
	return &UserServiceImpl{}
	
}
