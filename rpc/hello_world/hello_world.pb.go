// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rpc/hello_world/hello_world.proto

package hello_world

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type HelloReq struct {
	Subject              string   `protobuf:"bytes,1,opt,name=subject,proto3" json:"subject,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloReq) Reset()         { *m = HelloReq{} }
func (m *HelloReq) String() string { return proto.CompactTextString(m) }
func (*HelloReq) ProtoMessage()    {}
func (*HelloReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_39ce725870d9296c, []int{0}
}

func (m *HelloReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloReq.Unmarshal(m, b)
}
func (m *HelloReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloReq.Marshal(b, m, deterministic)
}
func (m *HelloReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloReq.Merge(m, src)
}
func (m *HelloReq) XXX_Size() int {
	return xxx_messageInfo_HelloReq.Size(m)
}
func (m *HelloReq) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloReq.DiscardUnknown(m)
}

var xxx_messageInfo_HelloReq proto.InternalMessageInfo

func (m *HelloReq) GetSubject() string {
	if m != nil {
		return m.Subject
	}
	return ""
}

type HelloResp struct {
	Text                 string   `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloResp) Reset()         { *m = HelloResp{} }
func (m *HelloResp) String() string { return proto.CompactTextString(m) }
func (*HelloResp) ProtoMessage()    {}
func (*HelloResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_39ce725870d9296c, []int{1}
}

func (m *HelloResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloResp.Unmarshal(m, b)
}
func (m *HelloResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloResp.Marshal(b, m, deterministic)
}
func (m *HelloResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloResp.Merge(m, src)
}
func (m *HelloResp) XXX_Size() int {
	return xxx_messageInfo_HelloResp.Size(m)
}
func (m *HelloResp) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloResp.DiscardUnknown(m)
}

var xxx_messageInfo_HelloResp proto.InternalMessageInfo

func (m *HelloResp) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func init() {
	proto.RegisterType((*HelloReq)(nil), "hello_world.HelloReq")
	proto.RegisterType((*HelloResp)(nil), "hello_world.HelloResp")
}

func init() { proto.RegisterFile("rpc/hello_world/hello_world.proto", fileDescriptor_39ce725870d9296c) }

var fileDescriptor_39ce725870d9296c = []byte{
	// 141 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x2c, 0x2a, 0x48, 0xd6,
	0xcf, 0x48, 0xcd, 0xc9, 0xc9, 0x8f, 0x2f, 0xcf, 0x2f, 0xca, 0x49, 0x41, 0x66, 0xeb, 0x15, 0x14,
	0xe5, 0x97, 0xe4, 0x0b, 0x71, 0x23, 0x09, 0x29, 0xa9, 0x70, 0x71, 0x78, 0x80, 0xb8, 0x41, 0xa9,
	0x85, 0x42, 0x12, 0x5c, 0xec, 0xc5, 0xa5, 0x49, 0x59, 0xa9, 0xc9, 0x25, 0x12, 0x8c, 0x0a, 0x8c,
	0x1a, 0x9c, 0x41, 0x30, 0xae, 0x92, 0x3c, 0x17, 0x27, 0x54, 0x55, 0x71, 0x81, 0x90, 0x10, 0x17,
	0x4b, 0x49, 0x6a, 0x05, 0x4c, 0x0d, 0x98, 0x6d, 0xe4, 0xc2, 0xc5, 0x05, 0x56, 0x10, 0x0e, 0x32,
	0x54, 0xc8, 0x8c, 0x8b, 0x15, 0xcc, 0x13, 0x12, 0xd5, 0x43, 0xb6, 0x1e, 0x66, 0x91, 0x94, 0x18,
	0x36, 0xe1, 0xe2, 0x02, 0x27, 0xc1, 0x28, 0x7e, 0x34, 0xe7, 0x27, 0xb1, 0x81, 0xdd, 0x6c, 0x0c,
	0x08, 0x00, 0x00, 0xff, 0xff, 0x71, 0x93, 0x1f, 0x0c, 0xd8, 0x00, 0x00, 0x00,
}