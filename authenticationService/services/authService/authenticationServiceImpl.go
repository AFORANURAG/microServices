package authService

import (
	context "context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	queueservice "github.com/AFORANURAG/microServices/authenticationService/services/queueService"
	userService "github.com/AFORANURAG/microServices/authenticationService/services/userService"
	authUtilities "github.com/AFORANURAG/microServices/authenticationService/utilityFunctions/authUtilites"
	emailService "github.com/AFORANURAG/microServices/emailService/services"
	"github.com/joho/godotenv"
)

type ConsumerMessageType struct {
	PhoneNumber string `validate:"required,len=10"`
	UserId int32`validate:"required"`
}
type AuthenticationServiceImpl struct {
	// interfaces can't be pointer
	client   userService.UserService
	emailclient emailService.EmailServiceClient
	p *queueservice.Producer 
}

func (a *AuthenticationServiceImpl) Signup(c context.Context, s *SignUpRequest) (*SignUpResponse, error) {
	// check whether a user already exists or not
	fmt.Printf("originURL is in authgrpc: %v", s.OriginURL)
	fmt.Printf("Signup Request : %v", s)
	_, err := a.client.GetUserByName(context.Background(), &userService.User{Name: &s.Name})
	if err != nil {
		// user does not exist
		// create the user here
		// We are gonna implement the magic link flow
		// so send him an email containing a magic link , and on clicking on that link , user will  be signedUp
		_, err := a.client.CreateUser(context.Background(), &userService.User{Name: &s.Name, Email: &s.Email,PhoneNumber: &s.PhoneNumber})

		if err!=nil {
			log.Printf("\n Error while creating user\n Error is %s",err)
			return &SignUpResponse{Status: 500,Success: false},err
		}

		// I don't understand why Name and Email are referenced differently.
		user,err:=a.client.GetUserByName(context.Background(), &userService.User{Name: &s.Name})
		if err!=nil{
			log.Fatalf("Error while fetching user with name in Signup for user fetching inside first erro block")
		}

fmt.Printf("user is:%v",user.Id)
		
		consumerMessage:=ConsumerMessageType{UserId:user.Id,PhoneNumber: s.PhoneNumber}
		    jsonString, err := json.Marshal(consumerMessage)
			fmt.Printf("json is %s",jsonString)
    if err != nil {
        log.Fatalf("Error marshalling struct to JSON: %v", err)
        
    }

	log.Printf("<----------------------------otp is send ==============================>")
		_, emailError := a.p.SendOtp(jsonString)
		if emailError != nil {
			log.Printf("Error While Sending OTP : %v", emailError)
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
	
	url := os.Getenv("2FA_AUTHENTICATION_URI")
	isVerified := authUtilities.Verify2factorOTP(url,req.PhoneNumber,int(req.Otp))
	fmt.Printf("isVerified: %v",isVerified)
	if isVerified{
		fmt.Printf("\nIsSigning\n:%t",req.IsSigningIn)
		// User is verified now call the login function with this email
		// mark the user as verified in the database
		_, markUserErr := a.client.MarkAsVerfied(context.Background(), &userService.MarkUserAsVerfiedRequest{IsVerified: true, PhoneNumber: req.PhoneNumber,Otp: int(req.Otp),IsSigningIn:req.IsSigningIn })
		if markUserErr != nil {
			log.Printf("Error While marking user as verified: %v", markUserErr)
			return &VerifyAccountResponse{
				Status:     500,
				IsVerified: false,
			}, markUserErr
		}
		// generate a jwt and give it user.
		// additionally this jwt should be saved to user profile each time he verifies in.
		// and based on this jwt, and based on this jwt, user auth state is determined, means
		// whether he is loggedin or signedUp
		token,err:=authUtilities.GenerateToken(req.PhoneNumber,os.Getenv("SECRET_KEY"))
		if err!=nil{
			log.Printf("\nError while Generating token\n Error")
			return &VerifyAccountResponse{Status: 500,IsVerified: false},err
		}

		return &VerifyAccountResponse{
			Status:     200,
			IsVerified: true,
			AccessToken: &token,
		}, nil
	}
	return &VerifyAccountResponse{
		Status:     400,
		IsVerified: false,
	}, nil
}

func (a *AuthenticationServiceImpl) Login(c context.Context, r *LoginRequest) (*LoginResponse, error) {
	// so we have a phoneNumber
	phoneNumber:=r.PhoneNumber
	// check whether this email exists in the database or not
	user, err := a.client.GetUserWithPhoneNumber(context.Background(), &userService.User{PhoneNumber: &phoneNumber})

	if err != nil {
		// there is an error
		log.Printf("Error While fetching user WithEmail: %v", err)
		return &LoginResponse{Status: 400, Success: false}, err
	}
	    // user exists , then sent him an email with a token
		// obviosly there a lot can be done here

		consumerMessage:=ConsumerMessageType{UserId:user.Id,PhoneNumber: phoneNumber}
		    jsonString, err := json.Marshal(consumerMessage)
			fmt.Printf("json is %s",jsonString)
    if err != nil {
        log.Fatalf("Error marshalling struct to JSON: %v", err)
        
    }

		_, emailError := a.p.SendOtp(jsonString)
		if emailError != nil {
			log.Printf("Error While Sending OTP : %v", emailError)
		return &LoginResponse{Status: 500, Success: false}, emailError
		}
		fmt.Print("<----------------OTP is send--------------->")
    // token,err:=authUtilities.GenerateToken(phoneNumber,os.Getenv("SECRET_KEY"))
	// if err!=nil{
	// 	return &LoginResponse{Status: 500,Success: false},err
	// }
	return &LoginResponse{Status: 200, Success: true}, nil
}

func (a *AuthenticationServiceImpl) mustEmbedUnimplementedAuthenticationServiceServer() {

}



func NewAuthenticationServiceProvider(client userService.UserService, emailServiceClient emailService.EmailServiceClient,p *queueservice.Producer ) AuthenticationServiceServer {
	return &AuthenticationServiceImpl{client: client, emailclient: emailServiceClient,p:p}
}
