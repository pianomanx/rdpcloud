// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.4
// source: services/netmgmt/user.proto

package netmgmt

import (
	netmgmt "github.com/s77rt/rdpcloud/proto/go/models/netmgmt"
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

type AddUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *netmgmt.User_3 `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *AddUserRequest) Reset() {
	*x = AddUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_netmgmt_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddUserRequest) ProtoMessage() {}

func (x *AddUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_netmgmt_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddUserRequest.ProtoReflect.Descriptor instead.
func (*AddUserRequest) Descriptor() ([]byte, []int) {
	return file_services_netmgmt_user_proto_rawDescGZIP(), []int{0}
}

func (x *AddUserRequest) GetUser() *netmgmt.User_3 {
	if x != nil {
		return x.User
	}
	return nil
}

type AddUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AddUserResponse) Reset() {
	*x = AddUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_netmgmt_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddUserResponse) ProtoMessage() {}

func (x *AddUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_netmgmt_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddUserResponse.ProtoReflect.Descriptor instead.
func (*AddUserResponse) Descriptor() ([]byte, []int) {
	return file_services_netmgmt_user_proto_rawDescGZIP(), []int{1}
}

type DeleteUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *netmgmt.User_1 `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *DeleteUserRequest) Reset() {
	*x = DeleteUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_netmgmt_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteUserRequest) ProtoMessage() {}

func (x *DeleteUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_netmgmt_user_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteUserRequest.ProtoReflect.Descriptor instead.
func (*DeleteUserRequest) Descriptor() ([]byte, []int) {
	return file_services_netmgmt_user_proto_rawDescGZIP(), []int{2}
}

func (x *DeleteUserRequest) GetUser() *netmgmt.User_1 {
	if x != nil {
		return x.User
	}
	return nil
}

type DeleteUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteUserResponse) Reset() {
	*x = DeleteUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_netmgmt_user_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteUserResponse) ProtoMessage() {}

func (x *DeleteUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_netmgmt_user_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteUserResponse.ProtoReflect.Descriptor instead.
func (*DeleteUserResponse) Descriptor() ([]byte, []int) {
	return file_services_netmgmt_user_proto_rawDescGZIP(), []int{3}
}

type GetUsersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetUsersRequest) Reset() {
	*x = GetUsersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_netmgmt_user_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUsersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUsersRequest) ProtoMessage() {}

func (x *GetUsersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_netmgmt_user_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUsersRequest.ProtoReflect.Descriptor instead.
func (*GetUsersRequest) Descriptor() ([]byte, []int) {
	return file_services_netmgmt_user_proto_rawDescGZIP(), []int{4}
}

type GetUsersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Users []*netmgmt.User `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
}

func (x *GetUsersResponse) Reset() {
	*x = GetUsersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_netmgmt_user_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUsersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUsersResponse) ProtoMessage() {}

func (x *GetUsersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_netmgmt_user_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUsersResponse.ProtoReflect.Descriptor instead.
func (*GetUsersResponse) Descriptor() ([]byte, []int) {
	return file_services_netmgmt_user_proto_rawDescGZIP(), []int{5}
}

func (x *GetUsersResponse) GetUsers() []*netmgmt.User {
	if x != nil {
		return x.Users
	}
	return nil
}

type GetUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *netmgmt.User_1 `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *GetUserRequest) Reset() {
	*x = GetUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_netmgmt_user_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserRequest) ProtoMessage() {}

func (x *GetUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_netmgmt_user_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserRequest.ProtoReflect.Descriptor instead.
func (*GetUserRequest) Descriptor() ([]byte, []int) {
	return file_services_netmgmt_user_proto_rawDescGZIP(), []int{6}
}

func (x *GetUserRequest) GetUser() *netmgmt.User_1 {
	if x != nil {
		return x.User
	}
	return nil
}

type GetUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *netmgmt.User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *GetUserResponse) Reset() {
	*x = GetUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_netmgmt_user_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserResponse) ProtoMessage() {}

func (x *GetUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_netmgmt_user_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserResponse.ProtoReflect.Descriptor instead.
func (*GetUserResponse) Descriptor() ([]byte, []int) {
	return file_services_netmgmt_user_proto_rawDescGZIP(), []int{7}
}

func (x *GetUserResponse) GetUser() *netmgmt.User {
	if x != nil {
		return x.User
	}
	return nil
}

type GetUserLocalGroupsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *netmgmt.User_1 `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *GetUserLocalGroupsRequest) Reset() {
	*x = GetUserLocalGroupsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_netmgmt_user_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserLocalGroupsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserLocalGroupsRequest) ProtoMessage() {}

func (x *GetUserLocalGroupsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_netmgmt_user_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserLocalGroupsRequest.ProtoReflect.Descriptor instead.
func (*GetUserLocalGroupsRequest) Descriptor() ([]byte, []int) {
	return file_services_netmgmt_user_proto_rawDescGZIP(), []int{8}
}

func (x *GetUserLocalGroupsRequest) GetUser() *netmgmt.User_1 {
	if x != nil {
		return x.User
	}
	return nil
}

type GetUserLocalGroupsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LocalGroups []*netmgmt.LocalGroup_1 `protobuf:"bytes,1,rep,name=local_groups,json=localGroups,proto3" json:"local_groups,omitempty"`
}

func (x *GetUserLocalGroupsResponse) Reset() {
	*x = GetUserLocalGroupsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_netmgmt_user_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserLocalGroupsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserLocalGroupsResponse) ProtoMessage() {}

func (x *GetUserLocalGroupsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_netmgmt_user_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserLocalGroupsResponse.ProtoReflect.Descriptor instead.
func (*GetUserLocalGroupsResponse) Descriptor() ([]byte, []int) {
	return file_services_netmgmt_user_proto_rawDescGZIP(), []int{9}
}

func (x *GetUserLocalGroupsResponse) GetLocalGroups() []*netmgmt.LocalGroup_1 {
	if x != nil {
		return x.LocalGroups
	}
	return nil
}

type ChangeUserPasswordRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *netmgmt.User_3 `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *ChangeUserPasswordRequest) Reset() {
	*x = ChangeUserPasswordRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_netmgmt_user_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChangeUserPasswordRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangeUserPasswordRequest) ProtoMessage() {}

func (x *ChangeUserPasswordRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_netmgmt_user_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangeUserPasswordRequest.ProtoReflect.Descriptor instead.
func (*ChangeUserPasswordRequest) Descriptor() ([]byte, []int) {
	return file_services_netmgmt_user_proto_rawDescGZIP(), []int{10}
}

func (x *ChangeUserPasswordRequest) GetUser() *netmgmt.User_3 {
	if x != nil {
		return x.User
	}
	return nil
}

type ChangeUserPasswordResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ChangeUserPasswordResponse) Reset() {
	*x = ChangeUserPasswordResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_netmgmt_user_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChangeUserPasswordResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangeUserPasswordResponse) ProtoMessage() {}

func (x *ChangeUserPasswordResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_netmgmt_user_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangeUserPasswordResponse.ProtoReflect.Descriptor instead.
func (*ChangeUserPasswordResponse) Descriptor() ([]byte, []int) {
	return file_services_netmgmt_user_proto_rawDescGZIP(), []int{11}
}

type EnableUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *netmgmt.User_1 `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *EnableUserRequest) Reset() {
	*x = EnableUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_netmgmt_user_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EnableUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnableUserRequest) ProtoMessage() {}

func (x *EnableUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_netmgmt_user_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnableUserRequest.ProtoReflect.Descriptor instead.
func (*EnableUserRequest) Descriptor() ([]byte, []int) {
	return file_services_netmgmt_user_proto_rawDescGZIP(), []int{12}
}

func (x *EnableUserRequest) GetUser() *netmgmt.User_1 {
	if x != nil {
		return x.User
	}
	return nil
}

type EnableUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EnableUserResponse) Reset() {
	*x = EnableUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_netmgmt_user_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EnableUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnableUserResponse) ProtoMessage() {}

func (x *EnableUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_netmgmt_user_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnableUserResponse.ProtoReflect.Descriptor instead.
func (*EnableUserResponse) Descriptor() ([]byte, []int) {
	return file_services_netmgmt_user_proto_rawDescGZIP(), []int{13}
}

type DisableUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *netmgmt.User_1 `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *DisableUserRequest) Reset() {
	*x = DisableUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_netmgmt_user_proto_msgTypes[14]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DisableUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DisableUserRequest) ProtoMessage() {}

func (x *DisableUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_netmgmt_user_proto_msgTypes[14]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DisableUserRequest.ProtoReflect.Descriptor instead.
func (*DisableUserRequest) Descriptor() ([]byte, []int) {
	return file_services_netmgmt_user_proto_rawDescGZIP(), []int{14}
}

func (x *DisableUserRequest) GetUser() *netmgmt.User_1 {
	if x != nil {
		return x.User
	}
	return nil
}

type DisableUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DisableUserResponse) Reset() {
	*x = DisableUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_netmgmt_user_proto_msgTypes[15]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DisableUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DisableUserResponse) ProtoMessage() {}

func (x *DisableUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_netmgmt_user_proto_msgTypes[15]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DisableUserResponse.ProtoReflect.Descriptor instead.
func (*DisableUserResponse) Descriptor() ([]byte, []int) {
	return file_services_netmgmt_user_proto_rawDescGZIP(), []int{15}
}

var File_services_netmgmt_user_proto protoreflect.FileDescriptor

var file_services_netmgmt_user_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x6e, 0x65, 0x74, 0x6d, 0x67,
	0x6d, 0x74, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x6e, 0x65, 0x74, 0x6d, 0x67, 0x6d, 0x74, 0x1a,
	0x19, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x6e, 0x65, 0x74, 0x6d, 0x67, 0x6d, 0x74, 0x2f,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x73, 0x2f, 0x6e, 0x65, 0x74, 0x6d, 0x67, 0x6d, 0x74, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x6c,
	0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3c, 0x0a, 0x0e,
	0x41, 0x64, 0x64, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2a,
	0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x6e, 0x65, 0x74, 0x6d, 0x67, 0x6d, 0x74, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x5f, 0x33, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x11, 0x0a, 0x0f, 0x41, 0x64,
	0x64, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x3f, 0x0a,
	0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x2a, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x16, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x6e, 0x65, 0x74, 0x6d, 0x67, 0x6d,
	0x74, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x5f, 0x31, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x14,
	0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x11, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x3e, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x55, 0x73,
	0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x05, 0x75,
	0x73, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x73, 0x2e, 0x6e, 0x65, 0x74, 0x6d, 0x67, 0x6d, 0x74, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x22, 0x3c, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2a, 0x0a, 0x04, 0x75, 0x73, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73,
	0x2e, 0x6e, 0x65, 0x74, 0x6d, 0x67, 0x6d, 0x74, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x5f, 0x31, 0x52,
	0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x3b, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e,
	0x6e, 0x65, 0x74, 0x6d, 0x67, 0x6d, 0x74, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73,
	0x65, 0x72, 0x22, 0x47, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x6f, 0x63,
	0x61, 0x6c, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x2a, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x6e, 0x65, 0x74, 0x6d, 0x67, 0x6d, 0x74, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x5f, 0x31, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x5d, 0x0a, 0x1a, 0x47,
	0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a, 0x0c, 0x6c, 0x6f, 0x63,
	0x61, 0x6c, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1c, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x6e, 0x65, 0x74, 0x6d, 0x67, 0x6d, 0x74,
	0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x31, 0x52, 0x0b, 0x6c,
	0x6f, 0x63, 0x61, 0x6c, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x22, 0x47, 0x0a, 0x19, 0x43, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x55, 0x73, 0x65, 0x72, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2a, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x6e,
	0x65, 0x74, 0x6d, 0x67, 0x6d, 0x74, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x5f, 0x33, 0x52, 0x04, 0x75,
	0x73, 0x65, 0x72, 0x22, 0x1c, 0x0a, 0x1a, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x55, 0x73, 0x65,
	0x72, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x3f, 0x0a, 0x11, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2a, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x6e, 0x65,
	0x74, 0x6d, 0x67, 0x6d, 0x74, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x5f, 0x31, 0x52, 0x04, 0x75, 0x73,
	0x65, 0x72, 0x22, 0x14, 0x0a, 0x12, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x40, 0x0a, 0x12, 0x44, 0x69, 0x73, 0x61,
	0x62, 0x6c, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2a,
	0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x6e, 0x65, 0x74, 0x6d, 0x67, 0x6d, 0x74, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x5f, 0x31, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x15, 0x0a, 0x13, 0x44, 0x69,
	0x73, 0x61, 0x62, 0x6c, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x42, 0x35, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x73, 0x37, 0x37, 0x72, 0x74, 0x2f, 0x72, 0x64, 0x70, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2f, 0x6e, 0x65, 0x74, 0x6d, 0x67, 0x6d, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_services_netmgmt_user_proto_rawDescOnce sync.Once
	file_services_netmgmt_user_proto_rawDescData = file_services_netmgmt_user_proto_rawDesc
)

func file_services_netmgmt_user_proto_rawDescGZIP() []byte {
	file_services_netmgmt_user_proto_rawDescOnce.Do(func() {
		file_services_netmgmt_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_services_netmgmt_user_proto_rawDescData)
	})
	return file_services_netmgmt_user_proto_rawDescData
}

var file_services_netmgmt_user_proto_msgTypes = make([]protoimpl.MessageInfo, 16)
var file_services_netmgmt_user_proto_goTypes = []interface{}{
	(*AddUserRequest)(nil),             // 0: services.netmgmt.AddUserRequest
	(*AddUserResponse)(nil),            // 1: services.netmgmt.AddUserResponse
	(*DeleteUserRequest)(nil),          // 2: services.netmgmt.DeleteUserRequest
	(*DeleteUserResponse)(nil),         // 3: services.netmgmt.DeleteUserResponse
	(*GetUsersRequest)(nil),            // 4: services.netmgmt.GetUsersRequest
	(*GetUsersResponse)(nil),           // 5: services.netmgmt.GetUsersResponse
	(*GetUserRequest)(nil),             // 6: services.netmgmt.GetUserRequest
	(*GetUserResponse)(nil),            // 7: services.netmgmt.GetUserResponse
	(*GetUserLocalGroupsRequest)(nil),  // 8: services.netmgmt.GetUserLocalGroupsRequest
	(*GetUserLocalGroupsResponse)(nil), // 9: services.netmgmt.GetUserLocalGroupsResponse
	(*ChangeUserPasswordRequest)(nil),  // 10: services.netmgmt.ChangeUserPasswordRequest
	(*ChangeUserPasswordResponse)(nil), // 11: services.netmgmt.ChangeUserPasswordResponse
	(*EnableUserRequest)(nil),          // 12: services.netmgmt.EnableUserRequest
	(*EnableUserResponse)(nil),         // 13: services.netmgmt.EnableUserResponse
	(*DisableUserRequest)(nil),         // 14: services.netmgmt.DisableUserRequest
	(*DisableUserResponse)(nil),        // 15: services.netmgmt.DisableUserResponse
	(*netmgmt.User_3)(nil),             // 16: models.netmgmt.User_3
	(*netmgmt.User_1)(nil),             // 17: models.netmgmt.User_1
	(*netmgmt.User)(nil),               // 18: models.netmgmt.User
	(*netmgmt.LocalGroup_1)(nil),       // 19: models.netmgmt.LocalGroup_1
}
var file_services_netmgmt_user_proto_depIdxs = []int32{
	16, // 0: services.netmgmt.AddUserRequest.user:type_name -> models.netmgmt.User_3
	17, // 1: services.netmgmt.DeleteUserRequest.user:type_name -> models.netmgmt.User_1
	18, // 2: services.netmgmt.GetUsersResponse.users:type_name -> models.netmgmt.User
	17, // 3: services.netmgmt.GetUserRequest.user:type_name -> models.netmgmt.User_1
	18, // 4: services.netmgmt.GetUserResponse.user:type_name -> models.netmgmt.User
	17, // 5: services.netmgmt.GetUserLocalGroupsRequest.user:type_name -> models.netmgmt.User_1
	19, // 6: services.netmgmt.GetUserLocalGroupsResponse.local_groups:type_name -> models.netmgmt.LocalGroup_1
	16, // 7: services.netmgmt.ChangeUserPasswordRequest.user:type_name -> models.netmgmt.User_3
	17, // 8: services.netmgmt.EnableUserRequest.user:type_name -> models.netmgmt.User_1
	17, // 9: services.netmgmt.DisableUserRequest.user:type_name -> models.netmgmt.User_1
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_services_netmgmt_user_proto_init() }
func file_services_netmgmt_user_proto_init() {
	if File_services_netmgmt_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_services_netmgmt_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddUserRequest); i {
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
		file_services_netmgmt_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddUserResponse); i {
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
		file_services_netmgmt_user_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteUserRequest); i {
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
		file_services_netmgmt_user_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteUserResponse); i {
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
		file_services_netmgmt_user_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUsersRequest); i {
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
		file_services_netmgmt_user_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUsersResponse); i {
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
		file_services_netmgmt_user_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserRequest); i {
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
		file_services_netmgmt_user_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserResponse); i {
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
		file_services_netmgmt_user_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserLocalGroupsRequest); i {
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
		file_services_netmgmt_user_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserLocalGroupsResponse); i {
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
		file_services_netmgmt_user_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChangeUserPasswordRequest); i {
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
		file_services_netmgmt_user_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChangeUserPasswordResponse); i {
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
		file_services_netmgmt_user_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EnableUserRequest); i {
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
		file_services_netmgmt_user_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EnableUserResponse); i {
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
		file_services_netmgmt_user_proto_msgTypes[14].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DisableUserRequest); i {
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
		file_services_netmgmt_user_proto_msgTypes[15].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DisableUserResponse); i {
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
			RawDescriptor: file_services_netmgmt_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   16,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_services_netmgmt_user_proto_goTypes,
		DependencyIndexes: file_services_netmgmt_user_proto_depIdxs,
		MessageInfos:      file_services_netmgmt_user_proto_msgTypes,
	}.Build()
	File_services_netmgmt_user_proto = out.File
	file_services_netmgmt_user_proto_rawDesc = nil
	file_services_netmgmt_user_proto_goTypes = nil
	file_services_netmgmt_user_proto_depIdxs = nil
}
