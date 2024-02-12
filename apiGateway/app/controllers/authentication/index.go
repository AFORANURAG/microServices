package authenticationController

import (
	"fmt"
	"log"

	ser "github.com/AFORANURAG/microServices/backend/apiGateway/pkg/services/authenticationService"
	validations "github.com/AFORANURAG/microServices/backend/apiGateway/pkg/validations/user"
	"github.com/gofiber/fiber/v2"
)

type IAuthentication interface {
	SignUp(c *fiber.Ctx) error
	VerificationHandler(c *fiber.Ctx) error
}

type AuthenticationController struct {
	authService ser.IAuthenticationService
}

func (a *AuthenticationController) SignUp(c *fiber.Ctx) error {
	var user validations.User
	c.BodyParser(&user)
	signedUpResponse, signUpError := a.authService.Signup(user.Name, user.Email)
	if signUpError != nil {
		log.Printf("error while signing up : % v", signUpError)
	}

	fmt.Printf("signedUpResponse is : %v", signedUpResponse)
	return c.Status(fiber.StatusAccepted).JSON(signedUpResponse)

}
func (a *AuthenticationController) VerificationHandler(c *fiber.Ctx) error {
	token := c.Query("token")
	// Here we are gonna verify whether the user is verified or not
	verifiedResponse, _ := a.authService.VerifyAccount(token)
	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"data": verifiedResponse,
	})
}

func ProvideAuthenticationController(authService ser.IAuthenticationService) *AuthenticationController {
	return &AuthenticationController{authService: authService}
}
