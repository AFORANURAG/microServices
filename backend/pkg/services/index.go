package services

import (
	"backend/pkg/services/authenticationService"
)

type Services struct {
	AuthenticationService authenticationService.IAuthenticationService
}

func RegisterServices() *Services {
	return &Services{AuthenticationService: authenticationService.AuthenticationServiceProvider()}

}
