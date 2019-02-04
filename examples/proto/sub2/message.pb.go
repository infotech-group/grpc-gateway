// Code generated by protoc-gen-go. DO NOT EDIT.
// source: examples/proto/sub2/message.proto

package sub2 // import "github.com/infotech-group/grpc-gateway/examples/proto/sub2"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type IdMessage struct {
	Uuid                 string   `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IdMessage) Reset()         { *m = IdMessage{} }
func (m *IdMessage) String() string { return proto.CompactTextString(m) }
func (*IdMessage) ProtoMessage()    {}
func (*IdMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_message_1fa155de06adc960, []int{0}
}
func (m *IdMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IdMessage.Unmarshal(m, b)
}
func (m *IdMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IdMessage.Marshal(b, m, deterministic)
}
func (dst *IdMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IdMessage.Merge(dst, src)
}
func (m *IdMessage) XXX_Size() int {
	return xxx_messageInfo_IdMessage.Size(m)
}
func (m *IdMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_IdMessage.DiscardUnknown(m)
}

var xxx_messageInfo_IdMessage proto.InternalMessageInfo

func (m *IdMessage) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func init() {
	proto.RegisterType((*IdMessage)(nil), "sub2.IdMessage")
}

func init() {
	proto.RegisterFile("examples/proto/sub2/message.proto", fileDescriptor_message_1fa155de06adc960)
}

var fileDescriptor_message_1fa155de06adc960 = []byte{
	// 130 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x4c, 0xad, 0x48, 0xcc,
	0x2d, 0xc8, 0x49, 0x2d, 0xd6, 0x2f, 0x28, 0xca, 0x2f, 0xc9, 0xd7, 0x2f, 0x2e, 0x4d, 0x32, 0xd2,
	0xcf, 0x4d, 0x2d, 0x2e, 0x4e, 0x4c, 0x4f, 0xd5, 0x03, 0x0b, 0x09, 0xb1, 0x80, 0xc4, 0x94, 0xe4,
	0xb9, 0x38, 0x3d, 0x53, 0x7c, 0x21, 0x12, 0x42, 0x42, 0x5c, 0x2c, 0xa5, 0xa5, 0x99, 0x29, 0x12,
	0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x60, 0xb6, 0x93, 0x4d, 0x94, 0x55, 0x7a, 0x66, 0x49, 0x46,
	0x69, 0x92, 0x5e, 0x72, 0x7e, 0xae, 0x7e, 0x7a, 0x51, 0x41, 0xb2, 0x6e, 0x6a, 0x72, 0x7e, 0x71,
	0x65, 0x71, 0x49, 0x2a, 0x94, 0x9b, 0x9e, 0x58, 0x92, 0x5a, 0x9e, 0x58, 0xa9, 0x8f, 0xc5, 0xca,
	0x24, 0x36, 0x30, 0xdb, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xfd, 0x49, 0xe7, 0x2f, 0x90, 0x00,
	0x00, 0x00,
}
