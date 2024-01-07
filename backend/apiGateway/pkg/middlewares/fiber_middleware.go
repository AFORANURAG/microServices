package middlewares

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func FiberMiddleware(a *fiber.App) {

	a.Use(
		cors.New(cors.Config{ExposeHeaders: "*", AllowOrigins: "*"}),
		limiter.New(limiter.Config{
			Max:               30,
			Expiration:        time.Minute * 1,
			LimiterMiddleware: limiter.SlidingWindow{},
		}),
		recover.New(),
	)

}
