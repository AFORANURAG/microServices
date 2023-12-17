package configs

import (
	"backend/pkg/services/authenticationService"
	userservice "backend/pkg/services/userService"

	"go.uber.org/dig"
)

var Container = dig.New()

func InjectServices() {
	Container.Provide(authenticationService.AuthenticationServiceProvider)
	Container.Provide(userservice.UserServiceProvider)
}
