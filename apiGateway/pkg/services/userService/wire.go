package userService

import (
	"github.com/google/wire"
)

func InitializeUserService(phrase string) *UserServiceImpl {
	wire.Build(UserServiceProvider)
	return &UserServiceImpl{}
}
