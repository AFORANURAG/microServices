package authService

import (
	context "context"
	"fmt"
	"log"
	"os"

	authUtilities "github.com/AFORANURAG/microServices/backend/authenticationService/utilityFunctions/authUtilites"
	emailService "github.com/AFORANURAG/microServices/backend/emailService/services"
	ser "github.com/AFORANURAG/microServices/backend/userService/services/userService"
	"github.com/joho/godotenv"
)

type AuthenticationServiceImpl struct {
	client      ser.UserServiceClient
	emailclient emailService.EmailServiceClient
}

func (a *AuthenticationServiceImpl) Signup(c context.Context, s *SignUpRequest) (*SignUpResponse, error) {
	// check whether a user already exists or not

	fmt.Printf("Signup Request : %v", s)
	_, err := a.client.GetUserByName(context.Background(), &ser.Request{Name: s.Name})
	if err != nil {
		// user does not exist
		// create the user here
		// We are gonna implement the magic link flow
		// so send him an email containing a magic link , and on clicking on that link , user will  be signedUp
		_, err := a.client.CreateUser(context.Background(), &ser.Request{Name: s.Name, Email: *s.Email})
		// token:=authUtilities.
		_, emailError := a.emailclient.SendEmail(context.Background(), &emailService.EmailServiceRequest{Email: *s.Email})
		if emailError != nil {
			log.Printf("Error While Sending Email : %v", emailError)
		}
		return &SignUpResponse{Status: 200, Success: true}, err
	}
	return &SignUpResponse{Status: 200, Success: true}, err
}

func (a *AuthenticationServiceImpl) VerifyUser(c context.Context, req *VerifyAccountRequest) (*VerifyAccountResponse, error) {
	// We Are gonna Verify the user here
	//First check whether the user's email exists in our db or not
	// first decode the token
	godotenv.Load()
	secretKey := os.Getenv("SECRET_KEY")
	isVerified, email := authUtilities.VerifyJWT(req.Token, secretKey)
	fmt.Printf("isVerified: %v , email: %v\n", isVerified, email)
	if isVerified && email != authUtilities.INVALID_TOKEN {
		// User is verified now call the login function with this email
		// mark the user as verified in the database
		_, markUserErr := a.client.MarkAsVerfied(context.Background(), &ser.MarkUserAsVerfiedRequest{IsVerified: true, Email: email})
		if markUserErr != nil {
			log.Printf("Error While marking user as verified: %v", markUserErr)
			return &VerifyAccountResponse{
				Status:     500,
				IsVerified: false,
			}, markUserErr
		}
		return &VerifyAccountResponse{
			Status:     200,
			IsVerified: true,
		}, nil
	}
	return &VerifyAccountResponse{
		Status:     400,
		IsVerified: false,
	}, nil
}
func (a *AuthenticationServiceImpl) mustEmbedUnimplementedAuthenticationServiceServer() {

}

func NewAuthenticationServiceProvider(client ser.UserServiceClient, emailServiceClient emailService.EmailServiceClient) *AuthenticationServiceImpl {
	return &AuthenticationServiceImpl{client: client, emailclient: emailServiceClient}
}
