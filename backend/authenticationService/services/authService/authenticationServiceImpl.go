package authService

import (
	context "context"

	ser "github.com/AFORANURAG/microServices/backend/userService"
	userServiceClient "github.com/AFORANURAG/microServices/clientProviders/authenticationServiceClient"
)

type AuthenticationServiceImpl struct {
	client ser.UserServiceClient
}

func (a *AuthenticationServiceImpl) Signup(c context.Context, s *SignUpRequest) (*SignUpResponse, error) {
	// check whether a user already exists or not
	u, err := a.client.GetUserById(context.Background(), ser.Request{Id: 1})
	if err != nil {
		// user does not exist
		// create the user here
		// We are gonna implement the magic link flow
		// so send him an email containing a magic link , and on clicking on that link , user will  be signedUp
		userProfile, err := a.client.CreateUser(context.Background(), &ser.Request{Name: "randomAnurag", Email: "randomAnurag123@gmail.com"})
		return &SignUpResponse{Status: 200, Success: true, UserId: userProfile.Id}
	}
	return &SignUpResponse{Status: 200, Success: true, UserId: u.Id}, err
}

// func (a *AuthenticationServiceImpl) Login(context.Context, *LoginRequest) (*LoginResponse, error) {

// }

func (a *AuthenticationServiceImpl) mustEmbedUnimplementedUserServiceServer() {

}

func NewAuthenticationServiceProvider(port string, host string) *AuthenticationServiceImpl {
	client := userServiceClient.UserServiceServiceClientProvider(port, host)
	return &AuthenticationServiceImpl{client: client}
}
