package middlewares

import (
	"github.com/gofiber/fiber/v2"
)

// what this middleware do is validate the token in the request.

func Authenticate(roles []string) fiber.Handler {

    return func(c *fiber.Ctx) error {
        // Your authentication logic here
authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).SendString("Please Authenticate.")			
		}
		
	return nil
    }
}
