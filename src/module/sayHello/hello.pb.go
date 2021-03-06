// Code generated by protoc-gen-go. DO NOT EDIT.
// source: hello.proto

/*
Package sayHello is a generated protocol buffer package.

It is generated from these files:
	hello.proto

It has these top-level messages:
	Hello_SayHello_Response
	Hello_SayHello_Request
*/
package sayHello

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

type Hello_SayHello_Response struct {
	Content string `protobuf:"bytes,1,opt,name=content" json:"content,omitempty"`
}

func (m *Hello_SayHello_Response) Reset()                    { *m = Hello_SayHello_Response{} }
func (m *Hello_SayHello_Response) String() string            { return proto.CompactTextString(m) }
func (*Hello_SayHello_Response) ProtoMessage()               {}
func (*Hello_SayHello_Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Hello_SayHello_Response) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

type Hello_SayHello_Request struct {
	Param string `protobuf:"bytes,1,opt,name=param" json:"param,omitempty"`
}

func (m *Hello_SayHello_Request) Reset()                    { *m = Hello_SayHello_Request{} }
func (m *Hello_SayHello_Request) String() string            { return proto.CompactTextString(m) }
func (*Hello_SayHello_Request) ProtoMessage()               {}
func (*Hello_SayHello_Request) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Hello_SayHello_Request) GetParam() string {
	if m != nil {
		return m.Param
	}
	return ""
}

func init() {
	proto.RegisterType((*Hello_SayHello_Response)(nil), "sayHello.Hello_SayHello_Response")
	proto.RegisterType((*Hello_SayHello_Request)(nil), "sayHello.Hello_SayHello_Request")
}

func init() { proto.RegisterFile("hello.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 127 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xce, 0x48, 0xcd, 0xc9,
	0xc9, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x28, 0x4e, 0xac, 0xf4, 0x00, 0xf1, 0x95,
	0x8c, 0xb9, 0xc4, 0xc1, 0x8c, 0xf8, 0x60, 0xa8, 0x48, 0x7c, 0x50, 0x6a, 0x71, 0x41, 0x7e, 0x5e,
	0x71, 0xaa, 0x90, 0x04, 0x17, 0x7b, 0x72, 0x7e, 0x5e, 0x49, 0x6a, 0x5e, 0x89, 0x04, 0xa3, 0x02,
	0xa3, 0x06, 0x67, 0x10, 0x8c, 0xab, 0xa4, 0xc7, 0x25, 0x86, 0xa1, 0xa9, 0xb0, 0x34, 0xb5, 0xb8,
	0x44, 0x48, 0x84, 0x8b, 0xb5, 0x20, 0xb1, 0x28, 0x31, 0x17, 0xaa, 0x03, 0xc2, 0x71, 0x12, 0x8e,
	0x12, 0xd4, 0xd3, 0xcf, 0xcd, 0x4f, 0x29, 0xcd, 0x49, 0xd5, 0x87, 0xd9, 0x9c, 0xc4, 0x06, 0x76,
	0x8a, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xc5, 0xc8, 0xc2, 0x33, 0x99, 0x00, 0x00, 0x00,
}
