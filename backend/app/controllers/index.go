package controller

import (
	authenticationController "backend/app/controllers/authentication"
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
