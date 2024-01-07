package service

import "fmt"

type IAuthenticationService interface {
	Signup() string
}

type AuthenticationServiceImpl struct {
	authServiceUri string
}

func (a *AuthenticationServiceImpl) Signup() string {
	fmt.Println("Fuck off")
	return "hello"
}

// AuthenticationServiceProvider now takes an IUserService interface instead of UserServiceImpl
func AuthenticationServiceProvider(u string) *AuthenticationServiceImpl {
	return &AuthenticationServiceImpl{u}
}
