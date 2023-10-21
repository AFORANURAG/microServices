package middleware

import (
	"backend/pkg/configs"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func FiberMiddleware(a *fiber.App) {
	a.Use(
		cors.New(cors.Config{ExposeHeaders: "*", AllowOrigins: "*"}),
		logger.New(configs.LoggerConfig()),
	)
}
