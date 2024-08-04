package authenticationController

import (
	"fmt"
	"log"
	"strconv"

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

	return c.Status(fiber.StatusAccepted).JSON(signedUpResponse)

}

func (a *AuthenticationController) Login(c *fiber.Ctx) error {
	var user validations.User
	c.BodyParser(&user)
	LoginResponse, LoginErr := a.authService.Login(user.PhoneNumber)
	if LoginErr != nil {
		log.Printf("error while signing up : % v", LoginErr)
	}

	return c.Status(fiber.StatusAccepted).JSON(LoginResponse)

}
func (a *AuthenticationController) VerificationHandler(c *fiber.Ctx) error {
	phoneNumber := c.Query("phoneNumber")
  otp, err := strconv.Atoi(c.Query("otp"))
  isSigningIn:=false
  if c.Query("isSigningIn")=="true"{
	  fmt.Printf("signIn kar raha hai,%s",c.Query("isSigningIn"))
	isSigningIn=true
  }

    if err != nil {
        fmt.Println("Error:", err)
return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message":"Invalid Otp,Please input a valid otp."})
    } else {
        fmt.Println("Integer:", otp)
    }
		originURL := c.OriginalURL()
	fmt.Printf("originURL: %v", originURL)
	err = godotenv.Load(".env")
	if err != nil {
		fmt.Printf("Error loading .env file:%v", err)
	}

	// Here we are gonna verify whether the user is verified or not
	verifiedResponse, _ := a.authService.VerifyAccount(phoneNumber,otp,isSigningIn)
	fmt.Printf("Verified response in verificationHandler is: %s",verifiedResponse)

	// this will redirect the user to the website if verification is successfull
	// if verifiedResponse.Status == 200 {
	// 	formattedURL := fmt.Sprintf("%s?token=%s", originURL, token)
	// 	c.Redirect(formattedURL)
	// }
	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"status": verifiedResponse.Status,
		"isVerified":verifiedResponse.IsVerified,
		"token":verifiedResponse.AccessToken,
	})
}

func ProvideAuthenticationController(authService ser.IAuthenticationService) *AuthenticationController {
	return &AuthenticationController{authService: authService}
}
