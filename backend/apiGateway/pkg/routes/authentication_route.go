package routes

import (
	"backend/apiGateway/pkg/middlewares"
	validations "backend/apiGateway/pkg/validations/user"

	controller "backend/apiGateway/app/controllers"

	"github.com/gofiber/fiber/v2"
)

func Authentication(routerPath string) func(a *fiber.App, basepath string, version string) {
	return func(a *fiber.App, basepath string, version string) {
		route := a.Group(basepath).Group(version).Group(routerPath)
		route.Post("/signup", middlewares.ValidateRequest(middlewares.Body, &validations.User{}), controller.AuthController.SignUp)
	}
}
