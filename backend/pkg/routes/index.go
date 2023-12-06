package routes

import (
	"github.com/gofiber/fiber/v2"
)

type RouteFunc func(app *fiber.App, basepath string, version string)

func AllRoutes() []RouteFunc {
	return []RouteFunc{HealthCheckRoute("health-check"), Authentication("authentication"), NotFoundRoute}
}
