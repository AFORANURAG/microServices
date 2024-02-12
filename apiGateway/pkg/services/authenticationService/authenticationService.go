package authenticationService

import (
	"context"
	"log"

	ser "github.com/AFORANURAG/microServices/backend/authenticationService/services/authService"
)

type SignedUpResponse struct {
	// Message string `validate:"required" json:"message"`
	Status  int32 `validate:"required" json:"status"`
	Success bool  `validate:"required" json:"success"`
}

type IAuthenticationService interface {
	Signup(name string, email string) (*SignedUpResponse, error)
	VerifyAccount(token string) (*ser.VerifyAccountResponse, error)
}

type AuthenticationServiceImpl struct {
	authServiceClient ser.AuthenticationServiceClient
}

func (a *AuthenticationServiceImpl) Signup(name string, email string) (*SignedUpResponse, error) {
	response, err := a.authServiceClient.Signup(context.Background(), &ser.SignUpRequest{Name: name, Email: &email})
	if err != nil {
		log.Printf("Error While Signing up in authentication in api gatewate : %v", err)
		return &SignedUpResponse{Status: 500, Success: false}, err
	}
	log.Printf("UserId : %v", response.UserId)

	return &SignedUpResponse{Status: response.Status, Success: response.Success}, nil
}

func (a *AuthenticationServiceImpl) VerifyAccount(token string) (*ser.VerifyAccountResponse, error) {
	response, err := a.authServiceClient.VerifyUser(context.Background(), &ser.VerifyAccountRequest{Token: token})
	if err != nil {
		log.Printf("Error verifying user : %v\n", err)
		return &ser.VerifyAccountResponse{Status: 500, IsVerified: false}, err
	}
	return response, nil
}

// AuthenticationServiceProvider now takes an IUserService interface instead of UserServiceImpl
func AuthenticationServiceProvider(authRPCClient ser.AuthenticationServiceClient) *AuthenticationServiceImpl {
	return &AuthenticationServiceImpl{authRPCClient}
}
