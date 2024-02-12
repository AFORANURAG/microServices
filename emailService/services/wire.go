//go:build wireinject
// +build wireinject

package emailService

import (
	"github.com/google/wire"
)

func InitializeAuthenticationService() *EmailServiceImpl {
	wire.Build(NewEmailServiceProvider)
	return &EmailServiceImpl{}
}
