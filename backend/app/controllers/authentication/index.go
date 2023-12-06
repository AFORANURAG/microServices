package authenticationController

import (
	"backend/pkg/configs"
	"backend/pkg/services/authenticationService"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type IAuthentication interface {
	SignUp(c *fiber.Ctx) error
	// SignIn(c *fiber.Ctx) error
}

type AuthenticationController struct {
}

var ser *authenticationService.AuthenticationServiceImpl

func (a *AuthenticationController) SignUp(c *fiber.Ctx) error {

	res := ""
	configs.Container.Invoke(func(authService *authenticationService.AuthenticationServiceImpl) {
		ser = authService
	})

	res = ser.Signup()
	// return res
	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{"message": fmt.Sprintf("%s", res)})

}

func ProvideAuthenticationController() *AuthenticationController {
	return &AuthenticationController{}
}
