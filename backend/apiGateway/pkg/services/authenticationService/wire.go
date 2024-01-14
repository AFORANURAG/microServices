//go:build wireinject
// +build wireinject

package authenticationService

import "github.com/google/wire"

func InitializeAuthenticationService(phrase string) *AuthenticationServiceImpl {
	wire.Build(AuthenticationServiceProvider)
	return &AuthenticationServiceImpl{}
}
