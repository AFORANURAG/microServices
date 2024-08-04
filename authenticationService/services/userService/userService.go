package userService

import (
	"context"
)

type User struct {
    Name  *string `protobuf:"bytes,1,opt,name=name,proto3,oneof" json:"name,omitempty"`
    Email *string `protobuf:"bytes,2,opt,name=email,proto3,oneof" json:"email,omitempty"`
    Id    *string `protobuf:"bytes,3,opt,name=id,proto3,oneof" json:"id,omitempty"`
	PhoneNumber  *string `protobuf:"bytes,1,opt,name=phoneNumber,proto3,oneof" json:"phoneNumber,omitempty"`

}

type Response struct {
    Name       *string `protobuf:"bytes,1,opt,name=name,proto3,oneof" json:"name,omitempty"`
    Id         int32   `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`              
    Email      string  `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
    IsVerified bool    `protobuf:"varint,4,opt,name=isVerified,proto3" json:"isVerified,omitempty"`
	PhoneNumber string    `protobuf:"varint,5,opt,name=phoneNumber,proto3" json:"phoneNumber,omitempty"`

}

type GetUserWithEmail struct {
	Email string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`

}

type CreateUserResponse struct{
	Status  int32     `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Success bool      `protobuf:"varint,2,opt,name=success,proto3" json:"success,omitempty"`
	Data    *Response `protobuf:"bytes,3,opt,name=data,proto3,oneof" json:"data,omitempty"`
}

type MarkUserAsVerfiedRequest struct {

	IsVerified bool   `protobuf:"varint,1,opt,name=isVerified,proto3" json:"isVerified,omitempty"`
	PhoneNumber      string `protobuf:"bytes,2,opt,name=phoneNumber,proto3" json:"phoneNumber,omitempty"`
	Otp      int `protobuf:"bytes,3,opt,name=otp,proto3" json:"otp,omitempty"`
	IsSigningIn bool `protobuf:"bytes,4,opt,name=isSigningIn,proto3" json:"isSigningIn,omitempty"`

}
type MarkUserAsVerfiedResponse struct{

	Status  int32     `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Success bool      `protobuf:"varint,2,opt,name=success,proto3" json:"success,omitempty"`
	Data    *Response `protobuf:"bytes,3,opt,name=data,proto3,oneof" json:"data,omitempty"`
}
type UserService interface {
	GetUserByName(context.Context, *User) (*Response, error)
	GetUserById(context.Context, *User) (*Response, error)
	GetUserByEmail(context.Context, *GetUserWithEmail) (*Response, error)
	GetUserWithPhoneNumber(context.Context,*User)(*Response,error)
	CreateUser(context.Context,  *User) (*CreateUserResponse, error)
	MarkAsVerfied(context.Context, *MarkUserAsVerfiedRequest) (*MarkUserAsVerfiedResponse, error)
	mustEmbedUnimplementedUserServiceServer()
}

