package routes

import (
	"backend/app/controllers"

	"github.com/gofiber/fiber/v2"
)

func HealthCheckRoute(a *fiber.App, basepath string, version string) {
	route := a.Group(basepath).Group(version)
	route.Get("/books", controllers.HealthCheck)
}
