package routes

import (
	controllers "github.com/AFORANURAG/microServices/apiGateway/app/controllers/healthCheck"
	middlewares "github.com/AFORANURAG/microServices/apiGateway/pkg/middlewares"

	"github.com/gofiber/fiber/v2"
)

func HealthCheckRoute(routerPath string) func(a *fiber.App, basepath string, version string) {
	return func(a *fiber.App, basepath string, version string) {
		route := a.Group(basepath).Group(version).Group(routerPath)
		route.Get("/",middlewares.Authenticate([]string{""}) ,controllers.HealthCheck)
	}
}
