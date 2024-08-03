//go:build wireinject
// +build wireinject

package userService

import (
	"github.com/AFORANURAG/microServices/authenticationService/userRepository"

	dbservice "github.com/AFORANURAG/microServices/authenticationService/services/dbService"

	userServiceTypes "github.com/AFORANURAG/microServices/authenticationService/types/userServiceTypes"
	"github.com/google/wire"
)

func InitializeUserService(phrase userServiceTypes.UserServicePhrase) UserService {
	wire.Build(NewUserServiceProvider, userRepository.NewUserRepositoryProvider, dbservice.NewDBServiceClientProvider)
	return &UserServiceImpl{}
	
}
