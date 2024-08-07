package userService

import (
	context "context"
	"fmt"
	"log"

	userSchema "github.com/AFORANURAG/microServices/userService/models"
	userRepository "github.com/AFORANURAG/microServices/userService/userRepository"
)

type UserServiceImpl struct {
	UserRepo userRepository.IUserRepository
}

func (u *UserServiceImpl) GetUserByName(c context.Context, in *Request) (*Response, error) {
	userProfileRow, err := u.UserRepo.GetUserByName(*in.Name)
	fmt.Printf("Profile : %v", in)
	if err != nil {
		log.Printf("Error while fetching users :%v ", err)
	}
	log.Printf("userProfile row %v", userProfileRow)
	// Row,err
	var response Response
	scanErr := userProfileRow.Scan(&response.Id, &response.Name, &response.Email, &response.IsVerified)
	if scanErr != nil {
		log.Printf("<--------------------------Error While fetching user from the database:%v---------------------------->", scanErr)
		return nil, scanErr
	}

	log.Printf("Response is %v", *response.Name)
	return &response, scanErr

}

func (u *UserServiceImpl) GetUserById(c context.Context, in *Request) (*Response, error) {
	userProfileRow, err := u.UserRepo.GetUserByName(*in.Id)
	if err != nil {
		log.Printf("Error while fetching users :%v ", err)
	}
	// Row,err
	var response Response
	scanErr := userProfileRow.Scan(&response.Id, &response.Name, &response.Email)

	return &response, scanErr
}

func (u *UserServiceImpl) GetUserByEmail(c context.Context, in *GetUserWithEmail) (*Response, error) {
	userProfile, err := u.UserRepo.GetUserWithEmail(in.Email)
	if err != nil {
		log.Printf("Error while fetching users :%v ", err)
		return &Response{}, nil
	}
	// Row,err
	response := &Response{
		Email:      userProfile.Email,
		Id:         userProfile.Id,
		IsVerified: userProfile.IsVerified,
		Name:       &userProfile.Name,
	}

	return response, nil
}

func (u *UserServiceImpl) CreateUser(c context.Context, in *Request) (*CreateUserResponse, error) {
	// Create the user here
	log.Printf("In request is %v", in)
	userProfile, userFetchError := u.GetUserByName(context.Background(), in)
	if userFetchError != nil {
		_, err := u.UserRepo.CreateUser(&userSchema.UserSchema{Name: *in.Name, Email: *in.Email})
		if err != nil {
			return nil, fmt.Errorf("Error creating user:%v", err)
		}
		return &CreateUserResponse{Status: 200, Success: true}, nil
	}

	return &CreateUserResponse{Status: 200, Success: true, Data: userProfile}, nil
}

func (u *UserServiceImpl) MarkAsVerfied(c context.Context, req *MarkUserAsVerfiedRequest) (*MarkUserAsVerfiedResponse, error) {
	// extract the email from the request
	email := req.Email
	userProfile, err := u.UserRepo.GetUserWithEmail(email)
	// check if it exists or not
	if err != nil {
		// No User Might be exist
		log.Printf("Error While fetching user In MarkAsVerified  : %v\n", err)
		return &MarkUserAsVerfiedResponse{Status: 400, Success: false}, err
	}
	if userProfile.IsVerified {
		return &MarkUserAsVerfiedResponse{Status: 200, Success: true}, nil
	}
	markedUser, markUserErr := u.UserRepo.UpdateVerificationStatus(req.IsVerified, email)
	if markedUser {
		return &MarkUserAsVerfiedResponse{Status: 200, Success: true}, nil
	}
	return &MarkUserAsVerfiedResponse{Status: 400, Success: false}, markUserErr

}


func (u *UserServiceImpl) mustEmbedUnimplementedUserServiceServer() {

}
func NewUserServiceProvider(urepo *userRepository.UserRepository) *UserServiceImpl {
	return &UserServiceImpl{UserRepo: urepo}
}
