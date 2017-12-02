// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protoc-gen-grpc-gateway/options/auth.proto

/*
Package options is a generated protocol buffer package.

It is generated from these files:
	protoc-gen-grpc-gateway/options/auth.proto

It has these top-level messages:
*/
package options

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/protoc-gen-go/descriptor"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

var E_IsAuthRequired = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.MethodOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         50001,
	Name:          "grpc.gateway.is_auth_required",
	Tag:           "varint,50001,opt,name=is_auth_required,json=isAuthRequired",
	Filename:      "protoc-gen-grpc-gateway/options/auth.proto",
}

func init() {
	proto.RegisterExtension(E_IsAuthRequired)
}

func init() { proto.RegisterFile("protoc-gen-grpc-gateway/options/auth.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 194 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x2a, 0x28, 0xca, 0x2f,
	0xc9, 0x4f, 0xd6, 0x4d, 0x4f, 0xcd, 0xd3, 0x4d, 0x2f, 0x2a, 0x48, 0xd6, 0x4d, 0x4f, 0x2c, 0x49,
	0x2d, 0x4f, 0xac, 0xd4, 0xcf, 0x2f, 0x28, 0xc9, 0xcc, 0xcf, 0x2b, 0xd6, 0x4f, 0x2c, 0x2d, 0xc9,
	0xd0, 0x03, 0x2b, 0x12, 0xe2, 0x01, 0x29, 0xd0, 0x83, 0x2a, 0x90, 0x52, 0x48, 0xcf, 0xcf, 0x4f,
	0xcf, 0x49, 0xd5, 0x07, 0xcb, 0x25, 0x95, 0xa6, 0xe9, 0xa7, 0xa4, 0x16, 0x27, 0x17, 0x65, 0x16,
	0x94, 0xe4, 0x17, 0x41, 0xd4, 0x5b, 0x79, 0x71, 0x09, 0x64, 0x16, 0xc7, 0x83, 0x0c, 0x88, 0x2f,
	0x4a, 0x2d, 0x2c, 0xcd, 0x2c, 0x4a, 0x4d, 0x11, 0x92, 0xd3, 0x83, 0x68, 0xd3, 0x83, 0x69, 0xd3,
	0xf3, 0x4d, 0x2d, 0xc9, 0xc8, 0x4f, 0xf1, 0x87, 0x58, 0x27, 0x71, 0xb1, 0x8d, 0x59, 0x81, 0x51,
	0x83, 0x23, 0x88, 0x2f, 0xb3, 0xd8, 0xb1, 0xb4, 0x24, 0x23, 0x08, 0xaa, 0xcf, 0xc9, 0x23, 0xca,
	0x2d, 0x3d, 0xb3, 0x24, 0xa3, 0x34, 0x49, 0x2f, 0x39, 0x3f, 0x57, 0x3f, 0x33, 0x2f, 0x2d, 0xbf,
	0x24, 0x35, 0x39, 0x43, 0x37, 0xbd, 0x28, 0xbf, 0xb4, 0x40, 0x1f, 0xc5, 0xe1, 0x04, 0x3c, 0x94,
	0xc4, 0x06, 0x56, 0x60, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x85, 0x17, 0x57, 0x0b, 0xfa, 0x00,
	0x00, 0x00,
}
