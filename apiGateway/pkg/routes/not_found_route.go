package routes

import "github.com/gofiber/fiber/v2"

// aka default route
func NotFoundRoute(a *fiber.App, basepath string, version string) {
	a.Use(
		func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": true,
				"msg":   "This Endpoint does not exists",
			})
		},
	)
}
