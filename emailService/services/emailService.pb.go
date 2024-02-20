// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.1
// source: emailService.proto

package emailService

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type EmailServiceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email     string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	OriginURL string `protobuf:"bytes,2,opt,name=originURL,proto3" json:"originURL,omitempty"`
}

func (x *EmailServiceRequest) Reset() {
	*x = EmailServiceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_emailService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmailServiceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmailServiceRequest) ProtoMessage() {}

func (x *EmailServiceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_emailService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmailServiceRequest.ProtoReflect.Descriptor instead.
func (*EmailServiceRequest) Descriptor() ([]byte, []int) {
	return file_emailService_proto_rawDescGZIP(), []int{0}
}

func (x *EmailServiceRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *EmailServiceRequest) GetOriginURL() string {
	if x != nil {
		return x.OriginURL
	}
	return ""
}

type EmailServiceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status  int32  `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Success bool   `protobuf:"varint,2,opt,name=success,proto3" json:"success,omitempty"`
	Data    string `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *EmailServiceResponse) Reset() {
	*x = EmailServiceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_emailService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmailServiceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmailServiceResponse) ProtoMessage() {}

func (x *EmailServiceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_emailService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmailServiceResponse.ProtoReflect.Descriptor instead.
func (*EmailServiceResponse) Descriptor() ([]byte, []int) {
	return file_emailService_proto_rawDescGZIP(), []int{1}
}

func (x *EmailServiceResponse) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *EmailServiceResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *EmailServiceResponse) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

var File_emailService_proto protoreflect.FileDescriptor

var file_emailService_proto_rawDesc = []byte{
	0x0a, 0x12, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x22, 0x49, 0x0a, 0x13, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12,
	0x1c, 0x0a, 0x09, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x55, 0x52, 0x4c, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x55, 0x52, 0x4c, 0x22, 0x5c, 0x0a,
	0x14, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a,
	0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07,
	0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x32, 0x64, 0x0a, 0x0c, 0x45,
	0x6d, 0x61, 0x69, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x54, 0x0a, 0x09, 0x73,
	0x65, 0x6e, 0x64, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x21, 0x2e, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x45, 0x6d, 0x61, 0x69, 0x6c,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x42, 0x04, 0x5a, 0x02, 0x2e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_emailService_proto_rawDescOnce sync.Once
	file_emailService_proto_rawDescData = file_emailService_proto_rawDesc
)

func file_emailService_proto_rawDescGZIP() []byte {
	file_emailService_proto_rawDescOnce.Do(func() {
		file_emailService_proto_rawDescData = protoimpl.X.CompressGZIP(file_emailService_proto_rawDescData)
	})
	return file_emailService_proto_rawDescData
}

var file_emailService_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_emailService_proto_goTypes = []interface{}{
	(*EmailServiceRequest)(nil),  // 0: emailService.EmailServiceRequest
	(*EmailServiceResponse)(nil), // 1: emailService.EmailServiceResponse
}
var file_emailService_proto_depIdxs = []int32{
	0, // 0: emailService.EmailService.sendEmail:input_type -> emailService.EmailServiceRequest
	1, // 1: emailService.EmailService.sendEmail:output_type -> emailService.EmailServiceResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_emailService_proto_init() }
func file_emailService_proto_init() {
	if File_emailService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_emailService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmailServiceRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_emailService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmailServiceResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_emailService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_emailService_proto_goTypes,
		DependencyIndexes: file_emailService_proto_depIdxs,
		MessageInfos:      file_emailService_proto_msgTypes,
	}.Build()
	File_emailService_proto = out.File
	file_emailService_proto_rawDesc = nil
	file_emailService_proto_goTypes = nil
	file_emailService_proto_depIdxs = nil
}
