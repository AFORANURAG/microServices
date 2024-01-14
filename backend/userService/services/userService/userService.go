package userService

import (
	context "context"
	"fmt"
	"log"

	userSchema "github.com/AFORANURAG/microServices/backend/userService/models"
	userRepository "github.com/AFORANURAG/microServices/backend/userService/userRepository"
)

type UserServiceImpl struct {
	UserRepo userRepository.IUserRepository
}

func (u *UserServiceImpl) GetUserByName(c context.Context, in *Request) (*Response, error) {
	userProfileRow, err := u.UserRepo.GetUserByName(in.Name)
	if err != nil {
		log.Fatalf("Error while fetching users :%v ", err)
	}
	// Row,err
	var response Response
	scanErr := userProfileRow.Scan(&response.Id, &response.Name, &response.Email)

	return &response, scanErr

}

func (u *UserServiceImpl) GetUserById(c context.Context, in *Request) (*Response, error) {
	userProfileRow, err := u.UserRepo.GetUserByName(in.Id)
	if err != nil {
		log.Fatalf("Error while fetching users :%v ", err)
	}
	// Row,err
	var response Response
	scanErr := userProfileRow.Scan(&response.Id, &response.Name, &response.Email)

	return &response, scanErr
}

func (u *UserServiceImpl) CreateUser(c context.Context, in *Request) (*CreateUserResponse, error) {
	// Create the user here

	userProfile, userFetchError := u.GetUserByName(context.Background(), in)
	if userFetchError != nil {
		_, err := u.UserRepo.CreateUser(&userSchema.UserSchema{Name: in.Name})
		if err != nil {
			return nil, fmt.Errorf("Error creating user:%v", err)
		}
		return &CreateUserResponse{Status: 200, Success: true}, nil
	}

	return &CreateUserResponse{Status: 200, Success: true, Data: userProfile}, nil
}

func (u *UserServiceImpl) mustEmbedUnimplementedUserServiceServer() {

}
func NewUserServiceProvider(urepo *userRepository.UserRepository) *UserServiceImpl {
	return &UserServiceImpl{UserRepo: urepo}
}
