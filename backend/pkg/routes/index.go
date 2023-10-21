package routes

import (
	"github.com/gofiber/fiber/v2"
)

// export all routes functions
type RouteFunc func(app *fiber.App, basepath string, version string)

func AllRoutes() []RouteFunc { return []RouteFunc{HealthCheckRoute, NotFoundRoute} }
