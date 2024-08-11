package userService

import (
	context "context"
	"fmt"
	"log"
	"time"

	userSchema "github.com/AFORANURAG/microServices/authenticationService/models"
	userRepository "github.com/AFORANURAG/microServices/authenticationService/userRepository"
)

type UserServiceImpl struct {
	UserRepo userRepository.IUserRepository
}

func (u *UserServiceImpl) GetUserByName(c context.Context, in *User) (*Response, error) {
	userProfileRow, err := u.UserRepo.GetUserByName(*in.Name)
	fmt.Printf("Profile : %v", in)
	if err != nil {
		log.Printf("Error while fetching users :%v ", err)
	}
	log.Printf("userProfile row %v", userProfileRow)
	// Row,err
	var response userSchema.NewUser
	var createdAtString string // Temporary variable to store the timestamp as a string
	scanErr := userProfileRow.Scan(&response.UserID, &response.Email, &response.Name, &response.PhoneNumber,&response.IsVerified,&createdAtString)
	if scanErr != nil {
		log.Printf("<--------------------------Error While fetching user from the database:%v---------------------------->", scanErr)
		return nil, scanErr
	}
        response.CreatedAt, err = time.Parse(time.RFC3339, createdAtString)
		if err!=nil{
			log.Printf("Error while parsing date in GetUserByName")
		}
	log.Printf("Response is %v", response.Name)
	return &Response{Name: &response.Name,Id: int32(response.UserID),Email: response.Email,IsVerified: response.IsVerified}, nil

}

func (u *UserServiceImpl) GetUserById(c context.Context, in *User) (*Response, error) {
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
		Id:         int32(userProfile.UserID),
		IsVerified: userProfile.IsVerified,
		Name:       &userProfile.Name,
	}

	return response, nil
}


func (u *UserServiceImpl) GetUserWithPhoneNumber(c context.Context, in *User) (*Response, error) {
	userProfile, err := u.UserRepo.GetUserWithPhoneNumber(*in.PhoneNumber)
	if err != nil {
		log.Printf("Error while fetching users :%v ", err)
		
		return &Response{}, err
	}
	// Row,err
	response := &Response{
		Email:      userProfile.Email,
		Id:         int32(userProfile.UserID),
		IsVerified: userProfile.IsVerified,
		Name:       &userProfile.Name,
	}

	return response, nil
}
func (u *UserServiceImpl) CreateUser(c context.Context, in *User) (*CreateUserResponse, error) {
	// Create the user here
	log.Printf("In request is %v", *in.PhoneNumber)
	userProfile, userFetchError := u.GetUserByName(context.Background(), in)
	
	if userFetchError != nil {
		_, err := u.UserRepo.CreateUser(&userSchema.UserSchema{Name: *in.Name, Email: *in.Email,PhoneNumber: *in.PhoneNumber})
		if err != nil {
			return nil, fmt.Errorf("Error creating user:%v", err)
		}
		return &CreateUserResponse{Status: 200, Success: true}, nil
	}

	return &CreateUserResponse{Status: 200, Success: true, Data: userProfile}, nil
}

func (u *UserServiceImpl) MarkAsVerfied(c context.Context, req *MarkUserAsVerfiedRequest) (*MarkUserAsVerfiedResponse, error) {
	// extract the email from the request
	phoneNumber := req.PhoneNumber
	userProfile, err := u.UserRepo.GetUserWithPhoneNumber(phoneNumber)
fmt.Printf("user Profile is %v",userProfile)
fmt.Printf("isUser SigningIn,%t",req.IsSigningIn)
	// check if it exists or not
	if err != nil {
		// No User Might be exist
		log.Printf("Error While fetching user In MarkAsVerified  : %v\n", err)
		return &MarkUserAsVerfiedResponse{Status: 400, Success: false}, err
	}
	if userProfile.IsVerified&&!req.IsSigningIn {
		return &MarkUserAsVerfiedResponse{Status: 200, Success: true}, nil
	}
	markedUser, markUserErr := u.UserRepo.UpdateVerificationStatus(req.IsVerified, nil,&phoneNumber,req.Otp)
	if markedUser {
		return &MarkUserAsVerfiedResponse{Status: 200, Success: true}, nil
	}
	return &MarkUserAsVerfiedResponse{Status: 400, Success: false}, markUserErr

}


func (u *UserServiceImpl) mustEmbedUnimplementedUserServiceServer() {

}
func NewUserServiceProvider(urepo userRepository.IUserRepository) UserService {
	return &UserServiceImpl{UserRepo: urepo}
}
