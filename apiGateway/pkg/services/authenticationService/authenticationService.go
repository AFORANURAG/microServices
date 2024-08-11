package authenticationService

import (
	"context"
	"fmt"
	"log"

	ser "github.com/AFORANURAG/microServices/authenticationService/services/authService"
)

type AuthenticationResponse struct {
	// Message string `validate:"required" json:"message"`
	Status  int32 `validate:"required" json:"status"`
	Success bool  `validate:"required" json:"success"`
}

type IAuthenticationService interface {
	Signup(name string, email string, originURL string,phoneNumber string) (*AuthenticationResponse, error)
	VerifyAccount(phoneNumber string,otp int,isSigningIn bool) (*ser.VerifyAccountResponse, error)
	Login(phoneNumber string) (*AuthenticationResponse, error)
}

type AuthenticationServiceImpl struct {
	authServiceClient ser.AuthenticationServiceClient
}

func (a *AuthenticationServiceImpl) Signup(name string, email string, originURL string,phoneNumber string) (*AuthenticationResponse, error) {
	response, err := a.authServiceClient.Signup(context.Background(), &ser.SignUpRequest{Name: name, Email: email, OriginURL: originURL,PhoneNumber: phoneNumber})
	if err != nil {
		log.Printf("Error While Signing up in authentication in api gatewate : %v", err)
		return &AuthenticationResponse{Status: 500, Success: false}, err
	}
	log.Printf("UserId : %v", response.UserId)

	return &AuthenticationResponse{Status: response.Status, Success: response.Success}, nil
}
func (a *AuthenticationServiceImpl) Login(phoneNumber string) (*AuthenticationResponse, error) {
	response, err := a.authServiceClient.Login(context.Background(), &ser.LoginRequest{PhoneNumber: phoneNumber})
	
	if err != nil {
		log.Printf("Error While Signing up in authentication in api gatewate : %v", err.Error())
		return &AuthenticationResponse{Status: 500, Success: false}, err
	}
	return &AuthenticationResponse{Status: response.Status, Success: response.Success}, nil
}

func (a *AuthenticationServiceImpl) VerifyAccount(phoneNumber string,otp int,isSigningIn bool) (*ser.VerifyAccountResponse, error) {
	response, err := a.authServiceClient.VerifyUser(context.Background(), &ser.VerifyAccountRequest{PhoneNumber: phoneNumber,Otp: int64(otp),IsSigningIn: isSigningIn})
	if err != nil {
		log.Printf("Error verifying user : %v\n", err)
		return &ser.VerifyAccountResponse{Status: 500, IsVerified: false}, err
	}
	fmt.Printf("\n\nresponse in VerifyAccount: %t\n\n",response.IsVerified)
	return response, nil
}

// AuthenticationServiceProvider now takes an IUserService interface instead of UserServiceImpl
func AuthenticationServiceProvider(authRPCClient ser.AuthenticationServiceClient) *AuthenticationServiceImpl {
	return &AuthenticationServiceImpl{authRPCClient}
}
