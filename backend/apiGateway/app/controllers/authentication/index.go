package authenticationController

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type IAuthentication interface {
	SignUp(c *fiber.Ctx) error
	// SignIn(c *fiber.Ctx) error
}

type AuthenticationController struct {
}

func (a *AuthenticationController) SignUp(c *fiber.Ctx) error {

	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{"message": fmt.Sprintf("%s", "hello")})

}

func ProvideAuthenticationController() *AuthenticationController {
	return &AuthenticationController{}
}
