package middlewares

import (
	"fmt"
	"os"
	"strings"

	utils "github.com/AFORANURAG/microServices/apiGateway/pkg/utils"
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
		// bearer token, so split
		parts:=strings.Split(authHeader," ")
		if len(parts)!=2 || strings.ToLower(parts[0])!="bearer"{
			return c.Status(fiber.StatusUnauthorized).SendString("Invalid Authorization Header.")
		}
		token:=parts[1]
		fmt.Printf("Token is: %s",token)
		IsVerified,_:=utils.VerifyJWT(token,os.Getenv("SECRET_KEY"))
fmt.Printf("Is verified: %t",IsVerified)
		if IsVerified{
			return c.Next()
		}
	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error":"Please Authenticate."})
    }
	
}
