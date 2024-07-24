//go:build wireinject
// +build wireinject

package authenticationService

import (
	authenticationServiceClient "github.com/AFORANURAG/microServices/apiGateway/pkg/clientProviders/authenticationServiceClient"
	"github.com/google/wire"
)

func InitializeAuthenticationService() *AuthenticationServiceImpl {
	wire.Build(AuthenticationServiceProvider, authenticationServiceClient.AuthenticationServiceClientProvider)
	return &AuthenticationServiceImpl{}
}
