package authenticationController

import (
	"fmt"
	"log"

	ser "github.com/AFORANURAG/microServices/apiGateway/pkg/services/authenticationService"
	validations "github.com/AFORANURAG/microServices/apiGateway/pkg/validations/user"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

type IAuthentication interface {
	SignUp(c *fiber.Ctx) error
	VerificationHandler(c *fiber.Ctx) error
	Login(c *fiber.Ctx) error
}

type AuthenticationController struct {
	authService ser.IAuthenticationService
}

func (a *AuthenticationController) SignUp(c *fiber.Ctx) error {
	originURL := c.OriginalURL()
	fmt.Printf("originURL is :%s", originURL)
	var user validations.User
	c.BodyParser(&user)
	signedUpResponse, signUpError := a.authService.Signup(user.Name, user.Email, originURL,user.PhoneNumber)
	if signUpError != nil {
		log.Printf("error while signing up : % v", signUpError)
	}

	fmt.Printf("signedUpResponse is : %v", signedUpResponse)
	return c.Status(fiber.StatusAccepted).JSON(signedUpResponse)

}

func (a *AuthenticationController) Login(c *fiber.Ctx) error {
	var user validations.User
	c.BodyParser(&user)
	LoginResponse, LoginErr := a.authService.Login(user.Email)
	if LoginErr != nil {
		log.Printf("error while signing up : % v", LoginErr)
	}

	fmt.Printf("signedUpResponse is : %v", LoginResponse)
	return c.Status(fiber.StatusAccepted).JSON(LoginResponse)

}
func (a *AuthenticationController) VerificationHandler(c *fiber.Ctx) error {
	token := c.Query("token")
	originURL := c.OriginalURL()
	fmt.Printf("originURL: %v", originURL)
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf("Error loading .env file:%v", err)
	}

	// Here we are gonna verify whether the user is verified or not
	verifiedResponse, _ := a.authService.VerifyAccount(token)

	// this will redirect the user to the website if verification is successfull
	if verifiedResponse.Status == 200 {
		formattedURL := fmt.Sprintf("%s?token=%s", originURL, token)
		c.Redirect(formattedURL)
	}
	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"data": verifiedResponse,
	})
}

func ProvideAuthenticationController(authService ser.IAuthenticationService) *AuthenticationController {
	return &AuthenticationController{authService: authService}
}
