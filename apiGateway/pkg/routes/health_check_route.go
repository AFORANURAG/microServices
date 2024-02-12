package routes

import (
	controllers "github.com/AFORANURAG/microServices/backend/apiGateway/app/controllers/healthCheck"

	"github.com/gofiber/fiber/v2"
)

func HealthCheckRoute(routerPath string) func(a *fiber.App, basepath string, version string) {
	return func(a *fiber.App, basepath string, version string) {
		route := a.Group(basepath).Group(version).Group(routerPath)
		route.Get("/", controllers.HealthCheck)
	}
}
