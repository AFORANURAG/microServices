package authenticationService

import (
	userservice "backend/pkg/services/userService"
)

type IUserService interface {
	GetUser(id string, name string) string
}

type UserServiceImpl struct {
}

func (u *UserServiceImpl) GetUser(id string, name string) string {
	return "Hello World!"
}

func UserServiceProvider() IUserService {
	return &UserServiceImpl{}
}

type IAuthenticationService interface {
	Signup() string
}

type AuthenticationServiceImpl struct {
	u userservice.IUserService
}

func (a *AuthenticationServiceImpl) Signup() string {
	return a.u.GetUser("hello", "hello")
}

// AuthenticationServiceProvider now takes an IUserService interface instead of UserServiceImpl
func AuthenticationServiceProvider(u IUserService) *AuthenticationServiceImpl {
	return &AuthenticationServiceImpl{u}
}
