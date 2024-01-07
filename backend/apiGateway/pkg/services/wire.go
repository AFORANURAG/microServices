//go:build wireinject
// +build wireinject

package service

import "github.com/google/wire"

func InitializeAuthenticationService(phrase string) *AuthenticationServiceImpl {
	wire.Build(AuthenticationServiceProvider)
	return &AuthenticationServiceImpl{}
}
func InitializeEmailService(phrase string) *EmailServiceImpl {
	wire.Build(EmailServiceProvider)
	return &EmailServiceImpl{}
}
