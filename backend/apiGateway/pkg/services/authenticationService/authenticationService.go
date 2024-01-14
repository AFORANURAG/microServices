package authenticationService

import (
	"fmt"

	authenticationServiceClient "github.com/AFORANURAG/microServices/backend/apiGateway/pkg/clientProviders/authenticationServiceClient"

	ser "github.com/AFORANURAG/microServices/backend/userService/services/userService"
)

type IAuthenticationService interface {
	Signup() string
}

type AuthenticationServiceImpl struct {
	authServiceUri    string
	authServiceClient ser.UserServiceClient
}

func (a *AuthenticationServiceImpl) Signup() string {
	fmt.Println("Fuck off")
	return "hello"
}

// AuthenticationServiceProvider now takes an IUserService interface instead of UserServiceImpl
func AuthenticationServiceProvider(u string) *AuthenticationServiceImpl {
	authServiceClient := authenticationServiceClient.AuthenticationServiceClientProvider()
	return &AuthenticationServiceImpl{u, authServiceClient}
}
