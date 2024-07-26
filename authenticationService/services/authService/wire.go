//go:build wireinject
// +build wireinject

package authService

import (
	emailServiceClient "github.com/AFORANURAG/microServices/authenticationService/clientProviders/emailServiceClient"
	userService "github.com/AFORANURAG/microServices/authenticationService/services/userService"

	"github.com/google/wire"
)

func InitializeAuthenticationService(phrase string) AuthenticationServiceServer {
	wire.Build(NewAuthenticationServiceProvider, userService.InitializeUserService, emailServiceClient.NewEmailServiceServiceClientProvider)
	
	return &AuthenticationServiceImpl{}
}
