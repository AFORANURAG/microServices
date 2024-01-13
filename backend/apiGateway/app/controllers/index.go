package controller

import (
	authenticationController "github.com/AFORANURAG/microServices/backend/apiGateway/app/controllers/authentication"
)

type Controller struct {
	authController authenticationController.IAuthentication
}

func RegisterController() *Controller {
	return &Controller{
		authController: authenticationController.ProvideAuthenticationController(),
	}
}

var P = authenticationController.ProvideAuthenticationController()
var AuthController = RegisterController().authController
