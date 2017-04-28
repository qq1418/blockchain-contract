// Code generated by protoc-gen-go.
// source: response-data.proto
// DO NOT EDIT!

package protos

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ResponseData struct {
	Ok   bool   `protobuf:"varint,1,opt,name=ok" json:"ok"`
	Msg  string `protobuf:"bytes,2,opt,name=msg" json:"msg"`
	Data string `protobuf:"bytes,3,opt,name=data" json:"data"`
}

func (m *ResponseData) Reset()                    { *m = ResponseData{} }
func (m *ResponseData) String() string            { return proto.CompactTextString(m) }
func (*ResponseData) ProtoMessage()               {}
func (*ResponseData) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *ResponseData) GetOk() bool {
	if m != nil {
		return m.Ok
	}
	return false
}

func (m *ResponseData) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *ResponseData) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

func init() {
	proto.RegisterType((*ResponseData)(nil), "protos.ResponseData")
}

func init() { proto.RegisterFile("response-data.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 136 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2e, 0x4a, 0x2d, 0x2e,
	0xc8, 0xcf, 0x2b, 0x4e, 0xd5, 0x4d, 0x49, 0x2c, 0x49, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0x62, 0x03, 0x53, 0xc5, 0x4a, 0x2e, 0x5c, 0x3c, 0x41, 0x50, 0x69, 0x97, 0xc4, 0x92, 0x44, 0x21,
	0x3e, 0x2e, 0xa6, 0xfc, 0x6c, 0x09, 0x46, 0x05, 0x46, 0x0d, 0x8e, 0x20, 0xa6, 0xfc, 0x6c, 0x21,
	0x01, 0x2e, 0xe6, 0xdc, 0xe2, 0x74, 0x09, 0x26, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x10, 0x53, 0x48,
	0x88, 0x8b, 0x05, 0x64, 0x8e, 0x04, 0x33, 0x58, 0x08, 0xcc, 0x76, 0xd2, 0xe2, 0x12, 0x4d, 0xce,
	0xcf, 0xd5, 0x2b, 0xcd, 0xcb, 0xd4, 0xcd, 0x49, 0x4d, 0x49, 0x4f, 0x2d, 0x82, 0xd8, 0x52, 0xec,
	0x24, 0x18, 0x00, 0xa2, 0x91, 0x6d, 0x48, 0x82, 0xd8, 0x6c, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff,
	0x04, 0x28, 0x43, 0x09, 0x97, 0x00, 0x00, 0x00,
}
