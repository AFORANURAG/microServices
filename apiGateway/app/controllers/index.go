package controller

import (
	authenticationController "github.com/AFORANURAG/microServices/apiGateway/app/controllers/authentication"
	authenticationService "github.com/AFORANURAG/microServices/apiGateway/pkg/services/authenticationService"
)

type Controller struct {
	authController authenticationController.IAuthentication
}

func RegisterController() *Controller {
	return &Controller{
		authController: authenticationController.ProvideAuthenticationController(authenticationService.InitializeAuthenticationService()),
	}
}

var P = authenticationController.ProvideAuthenticationController(authenticationService.InitializeAuthenticationService())
var AuthController = RegisterController().authController
