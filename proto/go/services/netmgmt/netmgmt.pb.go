// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.4
// source: services/netmgmt/netmgmt.proto

package netmgmt

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_services_netmgmt_netmgmt_proto protoreflect.FileDescriptor

var file_services_netmgmt_netmgmt_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x6e, 0x65, 0x74, 0x6d, 0x67,
	0x6d, 0x74, 0x2f, 0x6e, 0x65, 0x74, 0x6d, 0x67, 0x6d, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x10, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x6e, 0x65, 0x74, 0x6d, 0x67,
	0x6d, 0x74, 0x1a, 0x1b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x6e, 0x65, 0x74,
	0x6d, 0x67, 0x6d, 0x74, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x22, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x6e, 0x65, 0x74, 0x6d, 0x67, 0x6d,
	0x74, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x32, 0xd8, 0x09, 0x0a, 0x07, 0x4e, 0x65, 0x74, 0x6d, 0x67, 0x6d, 0x74, 0x12,
	0x50, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x55, 0x73, 0x65, 0x72, 0x12, 0x20, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x6e, 0x65, 0x74, 0x6d, 0x67, 0x6d, 0x74, 0x2e, 0x41, 0x64,
	0x64, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x6e, 0x65, 0x74, 0x6d, 0x67, 0x6d, 0x74, 0x2e,
	0x41, 0x64, 0x64, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x59, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12,
	0x23, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x6e, 0x65, 0x74, 0x6d, 0x67,
	0x6d, 0x74, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x6e, 0x65, 0x74, 0x6d, 0x67, 0x6d, 0x74, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x53, 0x0a, 0x08,
	0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x12, 0x21, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x6e, 0x65, 0x74, 0x6d, 0x67, 0x6d, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x6e, 0x65, 0x74, 0x6d, 0x67, 0x6d, 0x74, 0x2e, 0x47,
	0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x50, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x12, 0x20, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x6e, 0x65, 0x74, 0x6d, 0x67, 0x6d, 0x74, 0x2e,
	0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x6e, 0x65, 0x74, 0x6d, 0x67, 0x6d,
	0x74, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x71, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x6f,
	0x63, 0x61, 0x6c, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x12, 0x2b, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x6e, 0x65, 0x74, 0x6d, 0x67, 0x6d, 0x74, 0x2e, 0x47, 0x65, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x6e, 0x65, 0x74, 0x6d, 0x67, 0x6d, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65,
	0x72, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x71, 0x0a, 0x12, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x55, 0x73, 0x65, 0x72, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x2b, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x6e, 0x65, 0x74, 0x6d, 0x67, 0x6d, 0x74, 0x2e,
	0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x55, 0x73, 0x65, 0x72, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x6e, 0x65, 0x74, 0x6d, 0x67, 0x6d, 0x74, 0x2e, 0x43, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x55, 0x73, 0x65, 0x72, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x59, 0x0a, 0x0a, 0x45, 0x6e, 0x61,
	0x62, 0x6c, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x23, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x6e, 0x65, 0x74, 0x6d, 0x67, 0x6d, 0x74, 0x2e, 0x45, 0x6e, 0x61, 0x62, 0x6c,
	0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x6e, 0x65, 0x74, 0x6d, 0x67, 0x6d, 0x74, 0x2e,
	0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x5c, 0x0a, 0x0b, 0x44, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x55,
	0x73, 0x65, 0x72, 0x12, 0x24, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x6e,
	0x65, 0x74, 0x6d, 0x67, 0x6d, 0x74, 0x2e, 0x44, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x6e, 0x65, 0x74, 0x6d, 0x67, 0x6d, 0x74, 0x2e, 0x44, 0x69, 0x73,
	0x61, 0x62, 0x6c, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x74, 0x0a, 0x13, 0x41, 0x64, 0x64, 0x55, 0x73, 0x65, 0x72, 0x54, 0x6f, 0x4c,
	0x6f, 0x63, 0x61, 0x6c, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x2c, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x6e, 0x65, 0x74, 0x6d, 0x67, 0x6d, 0x74, 0x2e, 0x41, 0x64, 0x64,
	0x55, 0x73, 0x65, 0x72, 0x54, 0x6f, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x6e, 0x65, 0x74, 0x6d, 0x67, 0x6d, 0x74, 0x2e, 0x41, 0x64, 0x64, 0x55, 0x73,
	0x65, 0x72, 0x54, 0x6f, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x83, 0x01, 0x0a, 0x18, 0x52, 0x65, 0x6d,
	0x6f, 0x76, 0x65, 0x55, 0x73, 0x65, 0x72, 0x46, 0x72, 0x6f, 0x6d, 0x4c, 0x6f, 0x63, 0x61, 0x6c,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x31, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x6e, 0x65, 0x74, 0x6d, 0x67, 0x6d, 0x74, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x55,
	0x73, 0x65, 0x72, 0x46, 0x72, 0x6f, 0x6d, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x32, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x6e, 0x65, 0x74, 0x6d, 0x67, 0x6d, 0x74, 0x2e, 0x52, 0x65, 0x6d, 0x6f,
	0x76, 0x65, 0x55, 0x73, 0x65, 0x72, 0x46, 0x72, 0x6f, 0x6d, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x65,
	0x0a, 0x0e, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73,
	0x12, 0x27, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x6e, 0x65, 0x74, 0x6d,
	0x67, 0x6d, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x6e, 0x65, 0x74, 0x6d, 0x67, 0x6d, 0x74, 0x2e, 0x47, 0x65, 0x74,
	0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x77, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x73, 0x49, 0x6e, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x2d, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x6e, 0x65, 0x74, 0x6d, 0x67, 0x6d, 0x74,
	0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x49, 0x6e, 0x4c, 0x6f, 0x63, 0x61, 0x6c,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2e, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x6e, 0x65, 0x74, 0x6d, 0x67, 0x6d, 0x74, 0x2e,
	0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x49, 0x6e, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x35,
	0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x37, 0x37,
	0x72, 0x74, 0x2f, 0x72, 0x64, 0x70, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x67, 0x6f, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x6e, 0x65,
	0x74, 0x6d, 0x67, 0x6d, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_services_netmgmt_netmgmt_proto_goTypes = []interface{}{
	(*AddUserRequest)(nil),                   // 0: services.netmgmt.AddUserRequest
	(*DeleteUserRequest)(nil),                // 1: services.netmgmt.DeleteUserRequest
	(*GetUsersRequest)(nil),                  // 2: services.netmgmt.GetUsersRequest
	(*GetUserRequest)(nil),                   // 3: services.netmgmt.GetUserRequest
	(*GetUserLocalGroupsRequest)(nil),        // 4: services.netmgmt.GetUserLocalGroupsRequest
	(*ChangeUserPasswordRequest)(nil),        // 5: services.netmgmt.ChangeUserPasswordRequest
	(*EnableUserRequest)(nil),                // 6: services.netmgmt.EnableUserRequest
	(*DisableUserRequest)(nil),               // 7: services.netmgmt.DisableUserRequest
	(*AddUserToLocalGroupRequest)(nil),       // 8: services.netmgmt.AddUserToLocalGroupRequest
	(*RemoveUserFromLocalGroupRequest)(nil),  // 9: services.netmgmt.RemoveUserFromLocalGroupRequest
	(*GetLocalGroupsRequest)(nil),            // 10: services.netmgmt.GetLocalGroupsRequest
	(*GetUsersInLocalGroupRequest)(nil),      // 11: services.netmgmt.GetUsersInLocalGroupRequest
	(*AddUserResponse)(nil),                  // 12: services.netmgmt.AddUserResponse
	(*DeleteUserResponse)(nil),               // 13: services.netmgmt.DeleteUserResponse
	(*GetUsersResponse)(nil),                 // 14: services.netmgmt.GetUsersResponse
	(*GetUserResponse)(nil),                  // 15: services.netmgmt.GetUserResponse
	(*GetUserLocalGroupsResponse)(nil),       // 16: services.netmgmt.GetUserLocalGroupsResponse
	(*ChangeUserPasswordResponse)(nil),       // 17: services.netmgmt.ChangeUserPasswordResponse
	(*EnableUserResponse)(nil),               // 18: services.netmgmt.EnableUserResponse
	(*DisableUserResponse)(nil),              // 19: services.netmgmt.DisableUserResponse
	(*AddUserToLocalGroupResponse)(nil),      // 20: services.netmgmt.AddUserToLocalGroupResponse
	(*RemoveUserFromLocalGroupResponse)(nil), // 21: services.netmgmt.RemoveUserFromLocalGroupResponse
	(*GetLocalGroupsResponse)(nil),           // 22: services.netmgmt.GetLocalGroupsResponse
	(*GetUsersInLocalGroupResponse)(nil),     // 23: services.netmgmt.GetUsersInLocalGroupResponse
}
var file_services_netmgmt_netmgmt_proto_depIdxs = []int32{
	0,  // 0: services.netmgmt.Netmgmt.AddUser:input_type -> services.netmgmt.AddUserRequest
	1,  // 1: services.netmgmt.Netmgmt.DeleteUser:input_type -> services.netmgmt.DeleteUserRequest
	2,  // 2: services.netmgmt.Netmgmt.GetUsers:input_type -> services.netmgmt.GetUsersRequest
	3,  // 3: services.netmgmt.Netmgmt.GetUser:input_type -> services.netmgmt.GetUserRequest
	4,  // 4: services.netmgmt.Netmgmt.GetUserLocalGroups:input_type -> services.netmgmt.GetUserLocalGroupsRequest
	5,  // 5: services.netmgmt.Netmgmt.ChangeUserPassword:input_type -> services.netmgmt.ChangeUserPasswordRequest
	6,  // 6: services.netmgmt.Netmgmt.EnableUser:input_type -> services.netmgmt.EnableUserRequest
	7,  // 7: services.netmgmt.Netmgmt.DisableUser:input_type -> services.netmgmt.DisableUserRequest
	8,  // 8: services.netmgmt.Netmgmt.AddUserToLocalGroup:input_type -> services.netmgmt.AddUserToLocalGroupRequest
	9,  // 9: services.netmgmt.Netmgmt.RemoveUserFromLocalGroup:input_type -> services.netmgmt.RemoveUserFromLocalGroupRequest
	10, // 10: services.netmgmt.Netmgmt.GetLocalGroups:input_type -> services.netmgmt.GetLocalGroupsRequest
	11, // 11: services.netmgmt.Netmgmt.GetUsersInLocalGroup:input_type -> services.netmgmt.GetUsersInLocalGroupRequest
	12, // 12: services.netmgmt.Netmgmt.AddUser:output_type -> services.netmgmt.AddUserResponse
	13, // 13: services.netmgmt.Netmgmt.DeleteUser:output_type -> services.netmgmt.DeleteUserResponse
	14, // 14: services.netmgmt.Netmgmt.GetUsers:output_type -> services.netmgmt.GetUsersResponse
	15, // 15: services.netmgmt.Netmgmt.GetUser:output_type -> services.netmgmt.GetUserResponse
	16, // 16: services.netmgmt.Netmgmt.GetUserLocalGroups:output_type -> services.netmgmt.GetUserLocalGroupsResponse
	17, // 17: services.netmgmt.Netmgmt.ChangeUserPassword:output_type -> services.netmgmt.ChangeUserPasswordResponse
	18, // 18: services.netmgmt.Netmgmt.EnableUser:output_type -> services.netmgmt.EnableUserResponse
	19, // 19: services.netmgmt.Netmgmt.DisableUser:output_type -> services.netmgmt.DisableUserResponse
	20, // 20: services.netmgmt.Netmgmt.AddUserToLocalGroup:output_type -> services.netmgmt.AddUserToLocalGroupResponse
	21, // 21: services.netmgmt.Netmgmt.RemoveUserFromLocalGroup:output_type -> services.netmgmt.RemoveUserFromLocalGroupResponse
	22, // 22: services.netmgmt.Netmgmt.GetLocalGroups:output_type -> services.netmgmt.GetLocalGroupsResponse
	23, // 23: services.netmgmt.Netmgmt.GetUsersInLocalGroup:output_type -> services.netmgmt.GetUsersInLocalGroupResponse
	12, // [12:24] is the sub-list for method output_type
	0,  // [0:12] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_services_netmgmt_netmgmt_proto_init() }
func file_services_netmgmt_netmgmt_proto_init() {
	if File_services_netmgmt_netmgmt_proto != nil {
		return
	}
	file_services_netmgmt_user_proto_init()
	file_services_netmgmt_local_group_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_services_netmgmt_netmgmt_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_services_netmgmt_netmgmt_proto_goTypes,
		DependencyIndexes: file_services_netmgmt_netmgmt_proto_depIdxs,
	}.Build()
	File_services_netmgmt_netmgmt_proto = out.File
	file_services_netmgmt_netmgmt_proto_rawDesc = nil
	file_services_netmgmt_netmgmt_proto_goTypes = nil
	file_services_netmgmt_netmgmt_proto_depIdxs = nil
}
