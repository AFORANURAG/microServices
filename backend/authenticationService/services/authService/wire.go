package authService

import (
	"github.com/google/wire"
)

func InitializeUserService(phrase string) *AuthenticationServiceImpl {
	wire.Build(NewAuthenticationServiceProvider)
	return &AuthenticationServiceImpl{}
}
