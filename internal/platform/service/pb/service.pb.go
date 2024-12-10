// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.28.3
// source: service.proto

package api

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type UserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the user.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// First name of the user.
	FirstName *string `protobuf:"bytes,2,opt,name=first_name,json=firstName,proto3,oneof" json:"first_name,omitempty"`
	// Last name of the user.
	LastName *string `protobuf:"bytes,3,opt,name=last_name,json=lastName,proto3,oneof" json:"last_name,omitempty"`
	// Nickname of the user.
	Nickname *string `protobuf:"bytes,4,opt,name=nickname,proto3,oneof" json:"nickname,omitempty"`
	// Password hash of the user.
	PasswordHash *string `protobuf:"bytes,5,opt,name=password_hash,json=passwordHash,proto3,oneof" json:"password_hash,omitempty"`
	// Email of the user.
	Email *string `protobuf:"bytes,6,opt,name=email,proto3,oneof" json:"email,omitempty"`
	// Country of the user.
	Country *string `protobuf:"bytes,7,opt,name=country,proto3,oneof" json:"country,omitempty"`
}

func (x *UserRequest) Reset() {
	*x = UserRequest{}
	mi := &file_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserRequest) ProtoMessage() {}

func (x *UserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserRequest.ProtoReflect.Descriptor instead.
func (*UserRequest) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{0}
}

func (x *UserRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UserRequest) GetFirstName() string {
	if x != nil && x.FirstName != nil {
		return *x.FirstName
	}
	return ""
}

func (x *UserRequest) GetLastName() string {
	if x != nil && x.LastName != nil {
		return *x.LastName
	}
	return ""
}

func (x *UserRequest) GetNickname() string {
	if x != nil && x.Nickname != nil {
		return *x.Nickname
	}
	return ""
}

func (x *UserRequest) GetPasswordHash() string {
	if x != nil && x.PasswordHash != nil {
		return *x.PasswordHash
	}
	return ""
}

func (x *UserRequest) GetEmail() string {
	if x != nil && x.Email != nil {
		return *x.Email
	}
	return ""
}

func (x *UserRequest) GetCountry() string {
	if x != nil && x.Country != nil {
		return *x.Country
	}
	return ""
}

var File_service_proto protoreflect.FileDescriptor

var file_service_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0a, 0x61, 0x70, 0x69, 0x2e, 0x66, 0x61, 0x63, 0x65, 0x69, 0x74, 0x1a, 0x1c, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f,
	0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x62, 0x75, 0x66, 0x2f, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xf2, 0x05, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x3a, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x2a, 0xba, 0x48, 0x27, 0xba, 0x01, 0x1f, 0x12, 0x11, 0x6d, 0x75, 0x73, 0x74, 0x20, 0x6e, 0x6f,
	0x74, 0x20, 0x62, 0x65, 0x20, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0a, 0x74, 0x68, 0x69, 0x73,
	0x20, 0x21, 0x3d, 0x20, 0x27, 0x27, 0x72, 0x03, 0xb0, 0x01, 0x01, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x49, 0x0a, 0x0a, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x25, 0xba, 0x48, 0x22, 0xba, 0x01, 0x1f, 0x12, 0x11, 0x6d, 0x75, 0x73,
	0x74, 0x20, 0x6e, 0x6f, 0x74, 0x20, 0x62, 0x65, 0x20, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0a,
	0x74, 0x68, 0x69, 0x73, 0x20, 0x21, 0x3d, 0x20, 0x27, 0x27, 0x48, 0x00, 0x52, 0x09, 0x66, 0x69,
	0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x47, 0x0a, 0x09, 0x6c, 0x61,
	0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x25, 0xba,
	0x48, 0x22, 0xba, 0x01, 0x1f, 0x12, 0x11, 0x6d, 0x75, 0x73, 0x74, 0x20, 0x6e, 0x6f, 0x74, 0x20,
	0x62, 0x65, 0x20, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0a, 0x74, 0x68, 0x69, 0x73, 0x20, 0x21,
	0x3d, 0x20, 0x27, 0x27, 0x48, 0x01, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65,
	0x88, 0x01, 0x01, 0x12, 0x1f, 0x0a, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d,
	0x65, 0x88, 0x01, 0x01, 0x12, 0x86, 0x01, 0x0a, 0x0d, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x5c, 0xba, 0x48,
	0x59, 0xba, 0x01, 0x1f, 0x12, 0x11, 0x6d, 0x75, 0x73, 0x74, 0x20, 0x6e, 0x6f, 0x74, 0x20, 0x62,
	0x65, 0x20, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0a, 0x74, 0x68, 0x69, 0x73, 0x20, 0x21, 0x3d,
	0x20, 0x27, 0x27, 0xba, 0x01, 0x34, 0x12, 0x1e, 0x6d, 0x75, 0x73, 0x74, 0x20, 0x6e, 0x6f, 0x74,
	0x20, 0x65, 0x78, 0x63, 0x65, 0x65, 0x64, 0x20, 0x31, 0x32, 0x38, 0x20, 0x63, 0x68, 0x61, 0x72,
	0x61, 0x63, 0x74, 0x65, 0x72, 0x73, 0x1a, 0x12, 0x74, 0x68, 0x69, 0x73, 0x2e, 0x73, 0x69, 0x7a,
	0x65, 0x28, 0x29, 0x20, 0x3c, 0x3d, 0x20, 0x31, 0x32, 0x38, 0x48, 0x03, 0x52, 0x0c, 0x70, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x48, 0x61, 0x73, 0x68, 0x88, 0x01, 0x01, 0x12, 0x6a, 0x0a,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x42, 0x4f, 0xba, 0x48,
	0x4c, 0xba, 0x01, 0x1f, 0x12, 0x11, 0x6d, 0x75, 0x73, 0x74, 0x20, 0x6e, 0x6f, 0x74, 0x20, 0x62,
	0x65, 0x20, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0a, 0x74, 0x68, 0x69, 0x73, 0x20, 0x21, 0x3d,
	0x20, 0x27, 0x27, 0xba, 0x01, 0x27, 0x12, 0x15, 0x6d, 0x75, 0x73, 0x74, 0x20, 0x62, 0x65, 0x20,
	0x61, 0x20, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x20, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x1a, 0x0e, 0x74,
	0x68, 0x69, 0x73, 0x2e, 0x69, 0x73, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x28, 0x29, 0x48, 0x04, 0x52,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x88, 0x01, 0x01, 0x12, 0x71, 0x0a, 0x07, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x72, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x42, 0x52, 0xba, 0x48, 0x4f, 0xba,
	0x01, 0x1f, 0x12, 0x11, 0x6d, 0x75, 0x73, 0x74, 0x20, 0x6e, 0x6f, 0x74, 0x20, 0x62, 0x65, 0x20,
	0x65, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0a, 0x74, 0x68, 0x69, 0x73, 0x20, 0x21, 0x3d, 0x20, 0x27,
	0x27, 0xba, 0x01, 0x2a, 0x12, 0x16, 0x6d, 0x75, 0x73, 0x74, 0x20, 0x68, 0x61, 0x76, 0x65, 0x20,
	0x32, 0x20, 0x63, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x73, 0x1a, 0x10, 0x74, 0x68,
	0x69, 0x73, 0x2e, 0x73, 0x69, 0x7a, 0x65, 0x28, 0x29, 0x20, 0x3d, 0x3d, 0x20, 0x32, 0x48, 0x05,
	0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x88, 0x01, 0x01, 0x3a, 0x38, 0x92, 0x41,
	0x35, 0x0a, 0x33, 0x2a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x32, 0x1f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x20, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x20, 0x72, 0x65, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x20, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0xd2, 0x01, 0x02, 0x69, 0x64, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x66, 0x69, 0x72, 0x73, 0x74,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65,
	0x42, 0x10, 0x0a, 0x0e, 0x5f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x5f, 0x68, 0x61,
	0x73, 0x68, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x42, 0x0a, 0x0a, 0x08,
	0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x32, 0x9a, 0x03, 0x0a, 0x0d, 0x46, 0x61, 0x63,
	0x65, 0x69, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xae, 0x01, 0x0a, 0x07, 0x41,
	0x64, 0x64, 0x55, 0x73, 0x65, 0x72, 0x12, 0x17, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x66, 0x61, 0x63,
	0x65, 0x69, 0x74, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x72, 0x92, 0x41, 0x5b, 0x4a, 0x59, 0x0a, 0x03,
	0x32, 0x30, 0x34, 0x12, 0x52, 0x0a, 0x1c, 0x55, 0x73, 0x65, 0x72, 0x20, 0x77, 0x61, 0x73, 0x20,
	0x61, 0x64, 0x64, 0x65, 0x64, 0x20, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x66, 0x75, 0x6c,
	0x6c, 0x79, 0x2e, 0x12, 0x1a, 0x0a, 0x18, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22,
	0x16, 0x0a, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a,
	0x73, 0x6f, 0x6e, 0x12, 0x02, 0x7b, 0x7d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0e, 0x3a, 0x01, 0x2a,
	0x22, 0x09, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x12, 0xd7, 0x01, 0x0a, 0x0a,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x17, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x66, 0x61, 0x63, 0x65, 0x69, 0x74, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x97, 0x01, 0x92, 0x41,
	0x7b, 0x4a, 0x43, 0x0a, 0x03, 0x32, 0x30, 0x34, 0x12, 0x3c, 0x0a, 0x1e, 0x55, 0x73, 0x65, 0x72,
	0x20, 0x77, 0x61, 0x73, 0x20, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x20, 0x73, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x66, 0x75, 0x6c, 0x6c, 0x79, 0x2e, 0x12, 0x1a, 0x0a, 0x18, 0x1a, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x4a, 0x34, 0x0a, 0x03, 0x34, 0x30, 0x34, 0x12, 0x2d, 0x0a,
	0x0f, 0x55, 0x73, 0x65, 0x72, 0x20, 0x6e, 0x6f, 0x74, 0x20, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x2e,
	0x12, 0x1a, 0x0a, 0x18, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x13, 0x3a, 0x01, 0x2a, 0x32, 0x0e, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73,
	0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x42, 0xef, 0x03, 0x92, 0x41, 0xab, 0x03, 0x12, 0x3c, 0x0a, 0x06,
	0x66, 0x61, 0x63, 0x65, 0x69, 0x74, 0x12, 0x2d, 0x66, 0x61, 0x63, 0x65, 0x69, 0x74, 0x20, 0x69,
	0x73, 0x20, 0x73, 0x6d, 0x61, 0x6c, 0x6c, 0x20, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x20, 0x74, 0x6f, 0x20, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x20, 0x55,
	0x73, 0x65, 0x72, 0x73, 0x2e, 0x32, 0x03, 0x31, 0x2e, 0x30, 0x2a, 0x01, 0x01, 0x32, 0x10, 0x61,
	0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x3a,
	0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f,
	0x6e, 0x52, 0xbe, 0x01, 0x0a, 0x03, 0x34, 0x30, 0x30, 0x12, 0xb6, 0x01, 0x0a, 0x0c, 0x42, 0x61,
	0x64, 0x20, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x12, 0x16, 0x0a, 0x14, 0x1a, 0x12,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x22, 0x8d, 0x01, 0x0a, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x12, 0x79, 0x7b, 0x22, 0x63, 0x6f, 0x64, 0x65, 0x22,
	0x3a, 0x20, 0x34, 0x30, 0x30, 0x2c, 0x22, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x3a,
	0x20, 0x22, 0x42, 0x61, 0x64, 0x20, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x2c, 0x22,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x3a, 0x20, 0x22, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x69,
	0x64, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x22, 0x2c, 0x22, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73,
	0x22, 0x3a, 0x20, 0x5b, 0x7b, 0x22, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x22, 0x3a, 0x20, 0x22, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x22, 0x2c, 0x20, 0x22, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x22, 0x3a, 0x20, 0x22, 0x69, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x22, 0x7d,
	0x5d, 0x7d, 0x52, 0x82, 0x01, 0x0a, 0x03, 0x35, 0x30, 0x30, 0x12, 0x7b, 0x0a, 0x0f, 0x49, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x20, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x12, 0x16, 0x0a,
	0x14, 0x1a, 0x12, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x50, 0x0a, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x12, 0x3c, 0x7b, 0x22, 0x63, 0x6f, 0x64,
	0x65, 0x22, 0x3a, 0x20, 0x35, 0x30, 0x30, 0x2c, 0x20, 0x22, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x22, 0x3a, 0x20, 0x22, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x2c, 0x22, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x22, 0x3a, 0x20, 0x22, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x69, 0x64,
	0x5f, 0x75, 0x75, 0x69, 0x64, 0x22, 0x7d, 0x5a, 0x3e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x6f, 0x68, 0x65, 0x72, 0x6e, 0x61, 0x6e, 0x64, 0x65, 0x7a, 0x2f,
	0x66, 0x61, 0x63, 0x65, 0x69, 0x74, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f,
	0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2f, 0x70, 0x62, 0x2f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_service_proto_rawDescOnce sync.Once
	file_service_proto_rawDescData = file_service_proto_rawDesc
)

func file_service_proto_rawDescGZIP() []byte {
	file_service_proto_rawDescOnce.Do(func() {
		file_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_proto_rawDescData)
	})
	return file_service_proto_rawDescData
}

var file_service_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_service_proto_goTypes = []any{
	(*UserRequest)(nil),   // 0: api.faceit.UserRequest
	(*emptypb.Empty)(nil), // 1: google.protobuf.Empty
}
var file_service_proto_depIdxs = []int32{
	0, // 0: api.faceit.FaceitService.AddUser:input_type -> api.faceit.UserRequest
	0, // 1: api.faceit.FaceitService.UpdateUser:input_type -> api.faceit.UserRequest
	1, // 2: api.faceit.FaceitService.AddUser:output_type -> google.protobuf.Empty
	1, // 3: api.faceit.FaceitService.UpdateUser:output_type -> google.protobuf.Empty
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_service_proto_init() }
func file_service_proto_init() {
	if File_service_proto != nil {
		return
	}
	file_service_proto_msgTypes[0].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_proto_goTypes,
		DependencyIndexes: file_service_proto_depIdxs,
		MessageInfos:      file_service_proto_msgTypes,
	}.Build()
	File_service_proto = out.File
	file_service_proto_rawDesc = nil
	file_service_proto_goTypes = nil
	file_service_proto_depIdxs = nil
}
