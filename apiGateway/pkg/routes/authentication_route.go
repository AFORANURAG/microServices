package routes

import (
	validations "github.com/AFORANURAG/microServices/apiGateway/pkg/validations/user"

	"github.com/AFORANURAG/microServices/apiGateway/pkg/middlewares"

	controller "github.com/AFORANURAG/microServices/apiGateway/app/controllers"

	"github.com/gofiber/fiber/v2"
)

func Authentication(routerPath string) func(a *fiber.App, basepath string, version string) {
	return func(a *fiber.App, basepath string, version string) {
		route := a.Group(basepath).Group(version).Group(routerPath)
		route.Post("/signup", middlewares.ValidateRequest(middlewares.Body, &validations.User{}), controller.AuthController.SignUp)
		route.Post("/login", middlewares.ValidateRequest(middlewares.Body, &validations.User{}), controller.AuthController.Login)
		route.Get("/verify", controller.AuthController.VerificationHandler)
	}
}
