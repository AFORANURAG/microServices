package routes

import (
	validations "github.com/AFORANURAG/microServices/backend/apiGateway/pkg/validations/user"

	"github.com/AFORANURAG/microServices/backend/apiGateway/pkg/middlewares"

	controller "github.com/AFORANURAG/microServices/backend/apiGateway/app/controllers"

	"github.com/gofiber/fiber/v2"
)

func Authentication(routerPath string) func(a *fiber.App, basepath string, version string) {
	return func(a *fiber.App, basepath string, version string) {
		route := a.Group(basepath).Group(version).Group(routerPath)
		route.Post("/signup", middlewares.ValidateRequest(middlewares.Body, &validations.User{}), controller.AuthController.SignUp)
		route.Get("/verify", controller.AuthController.VerificationHandler)
	}
}
