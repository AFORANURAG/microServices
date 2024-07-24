//go:build wireinject
// +build wireinject

package authService

import (
	emailServiceClient "github.com/AFORANURAG/microServices/authenticationService/clientProviders/emailServiceClient"
	userServiceClient "github.com/AFORANURAG/microServices/authenticationService/clientProviders/userServiceClient"

	"github.com/google/wire"
)

func InitializeAuthenticationService(phrase string) *AuthenticationServiceImpl {
	wire.Build(NewAuthenticationServiceProvider, userServiceClient.NewUserServiceServiceClientProvider, emailServiceClient.NewEmailServiceServiceClientProvider)
	return &AuthenticationServiceImpl{}
}
