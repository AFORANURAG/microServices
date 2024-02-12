package userSchema

import "google.golang.org/protobuf/runtime/protoimpl"

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  *string `protobuf:"bytes,1,opt,name=name,proto3,oneof" json:"name,omitempty"`
	Id    int32   `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"` // Unique ID  for this person.
	Email string  `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
}

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Email string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Id    string `protobuf:"bytes,3,opt,name=id,proto3" json:"id,omitempty"`
}

type CreateUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status  int32     `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Success bool      `protobuf:"varint,2,opt,name=success,proto3" json:"success,omitempty"`
	Data    *Response `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

type UserSchema struct {
	Name       string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Email      string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Id         string `protobuf:"bytes,3,opt,name=id,proto3" json:"id,omitempty"`
	IsVerified bool   `protobuf:"bytes,4,opt,name=isVerified,proto3" json:"isVerified,omitempty"`
}
