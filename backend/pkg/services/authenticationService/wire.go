//go:build wireinject
// +build wireinject

package authenticationService

import "github.com/google/wire"

func InitializeAuthenticationService() *AuthenticationServiceImpl {
	wire.Build(AuthenticationServiceProvider, UserServiceProvider)
	return &AuthenticationServiceImpl{}
}
