// Code generated by protoc-gen-go.
// source: contract_execute_logs.proto
// DO NOT EDIT!

/*
Package protos is a generated protocol buffer package.

It is generated from these files:
	contract_execute_logs.proto
	contract_list.proto
	contract.proto
	response_pagination.proto
	response.proto

It has these top-level messages:
	ContractExecuteLog
	ContractExecuteLogList
	ContractList
	ContractSignature
	ContractAsset
	ExpressionResult
	ComponentsExpression
	ComponentDataSub
	ComponentData
	SelectBranchExpression
	ContractComponentSub
	ContractComponent
	ContractBody
	ContractHead
	Contract
	ResponsePagination
	Pagination
	Response
*/
package protos

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

type ContractExecuteLog struct {
	ContractHashId string `protobuf:"bytes,1,opt,name=ContractHashId" json:"ContractHashId"`
	// transaction.Relation.TaskId
	TaskId string `protobuf:"bytes,2,opt,name=TaskId" json:"TaskId"`
	// transaction timestamp
	Timestamp string `protobuf:"bytes,3,opt,name=Timestamp" json:"Timestamp"`
	// ContractComponents.Caption
	Caption string `protobuf:"bytes,4,opt,name=Caption" json:"Caption"`
	// ContractComponents.Cname
	Cname string `protobuf:"bytes,5,opt,name=Cname" json:"Cname"`
	// ContractComponents.Cname
	Ctype string `protobuf:"bytes,6,opt,name=Ctype" json:"Ctype"`
	// ContractComponents.Description
	Description string `protobuf:"bytes,7,opt,name=Description" json:"Description"`
	// ContractComponents.MetaAttribute
	MetaAttribute map[string]string `protobuf:"bytes,8,rep,name=MetaAttribute" json:"MetaAttribute" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// ContractComponents.State
	State string `protobuf:"bytes,9,opt,name=State" json:"State"`
	// ContractBody.CreateTime
	CreateTime string `protobuf:"bytes,10,opt,name=CreateTime" json:"CreateTime"`
}

func (m *ContractExecuteLog) Reset()                    { *m = ContractExecuteLog{} }
func (m *ContractExecuteLog) String() string            { return proto.CompactTextString(m) }
func (*ContractExecuteLog) ProtoMessage()               {}
func (*ContractExecuteLog) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ContractExecuteLog) GetContractHashId() string {
	if m != nil {
		return m.ContractHashId
	}
	return ""
}

func (m *ContractExecuteLog) GetTaskId() string {
	if m != nil {
		return m.TaskId
	}
	return ""
}

func (m *ContractExecuteLog) GetTimestamp() string {
	if m != nil {
		return m.Timestamp
	}
	return ""
}

func (m *ContractExecuteLog) GetCaption() string {
	if m != nil {
		return m.Caption
	}
	return ""
}

func (m *ContractExecuteLog) GetCname() string {
	if m != nil {
		return m.Cname
	}
	return ""
}

func (m *ContractExecuteLog) GetCtype() string {
	if m != nil {
		return m.Ctype
	}
	return ""
}

func (m *ContractExecuteLog) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *ContractExecuteLog) GetMetaAttribute() map[string]string {
	if m != nil {
		return m.MetaAttribute
	}
	return nil
}

func (m *ContractExecuteLog) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func (m *ContractExecuteLog) GetCreateTime() string {
	if m != nil {
		return m.CreateTime
	}
	return ""
}

type ContractExecuteLogList struct {
	ContractLogs []*ContractExecuteLog `protobuf:"bytes,1,rep,name=ContractLogs" json:"ContractLogs"`
}

func (m *ContractExecuteLogList) Reset()                    { *m = ContractExecuteLogList{} }
func (m *ContractExecuteLogList) String() string            { return proto.CompactTextString(m) }
func (*ContractExecuteLogList) ProtoMessage()               {}
func (*ContractExecuteLogList) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ContractExecuteLogList) GetContractLogs() []*ContractExecuteLog {
	if m != nil {
		return m.ContractLogs
	}
	return nil
}

func init() {
	proto.RegisterType((*ContractExecuteLog)(nil), "protos.ContractExecuteLog")
	proto.RegisterType((*ContractExecuteLogList)(nil), "protos.ContractExecuteLogList")
}

func init() { proto.RegisterFile("contract_execute_logs.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 343 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0xc1, 0x6b, 0xea, 0x40,
	0x10, 0xc6, 0x89, 0x79, 0xea, 0x73, 0x7c, 0x4f, 0x1e, 0x8b, 0xc8, 0xe2, 0x2b, 0x45, 0x3c, 0x14,
	0x2f, 0x0d, 0xa5, 0xbd, 0x94, 0x1e, 0x4a, 0x6b, 0x2a, 0x54, 0xb0, 0x50, 0xd4, 0x43, 0x6f, 0xb2,
	0xc6, 0x21, 0x0d, 0x9a, 0x6c, 0xd8, 0x9d, 0x94, 0xfa, 0x17, 0xf5, 0xdf, 0x2c, 0xd9, 0x4d, 0xa8,
	0x36, 0xf4, 0xb4, 0xf9, 0x7e, 0xdf, 0xcc, 0x30, 0xdf, 0x10, 0xf8, 0x1f, 0xc8, 0x84, 0x94, 0x08,
	0x68, 0x85, 0xef, 0x18, 0x64, 0x84, 0xab, 0x9d, 0x0c, 0xb5, 0x97, 0x2a, 0x49, 0x92, 0x35, 0xcc,
	0xa3, 0xfb, 0x9d, 0xb2, 0xc8, 0xf2, 0xe1, 0x87, 0x0b, 0xcc, 0x2f, 0xd0, 0xc4, 0xb6, 0xcd, 0x64,
	0xc8, 0xce, 0xa0, 0x53, 0xd2, 0x47, 0xa1, 0x5f, 0xa7, 0x1b, 0xee, 0x0c, 0x9c, 0x51, 0x6b, 0xfe,
	0x8d, 0xb2, 0x1e, 0x34, 0x96, 0x42, 0x6f, 0xa7, 0x1b, 0x5e, 0x33, 0x7e, 0xa1, 0xd8, 0x09, 0xb4,
	0x96, 0x51, 0x8c, 0x9a, 0x44, 0x9c, 0x72, 0xd7, 0x58, 0x5f, 0x80, 0x71, 0x68, 0xfa, 0x22, 0xa5,
	0x48, 0x26, 0xfc, 0x97, 0xf1, 0x4a, 0xc9, 0xba, 0x50, 0xf7, 0x13, 0x11, 0x23, 0xaf, 0x1b, 0x6e,
	0x85, 0xa1, 0xb4, 0x4f, 0x91, 0x37, 0x0a, 0x9a, 0x0b, 0x36, 0x80, 0xf6, 0x03, 0xea, 0x40, 0x45,
	0x76, 0x52, 0xd3, 0x78, 0x87, 0x88, 0x2d, 0xe0, 0xef, 0x13, 0x92, 0xb8, 0x27, 0x52, 0xd1, 0x3a,
	0x23, 0xe4, 0xbf, 0x07, 0xee, 0xa8, 0x7d, 0x79, 0x6e, 0xb3, 0x6b, 0xaf, 0x1a, 0xdc, 0x3b, 0xaa,
	0x9f, 0x24, 0xa4, 0xf6, 0xf3, 0xe3, 0x19, 0xf9, 0x32, 0x0b, 0x12, 0x84, 0xbc, 0x65, 0x97, 0x31,
	0x82, 0x9d, 0x02, 0xf8, 0x0a, 0x05, 0x61, 0x9e, 0x92, 0x83, 0xb1, 0x0e, 0x48, 0xff, 0x0e, 0x58,
	0x75, 0x34, 0xfb, 0x07, 0xee, 0x16, 0xf7, 0xc5, 0x6d, 0xf3, 0xcf, 0x7c, 0xfa, 0x9b, 0xd8, 0x65,
	0x58, 0xdc, 0xd3, 0x8a, 0x9b, 0xda, 0xb5, 0x33, 0x7c, 0x81, 0x5e, 0x75, 0xdf, 0x59, 0xa4, 0x89,
	0xdd, 0xc2, 0x9f, 0xd2, 0x99, 0xc9, 0x50, 0x73, 0xc7, 0xa4, 0xec, 0xff, 0x9c, 0x72, 0x7e, 0x54,
	0x3f, 0xbe, 0x80, 0x6e, 0x20, 0x63, 0x2f, 0x4b, 0xa2, 0x1d, 0x6e, 0x42, 0x54, 0x45, 0xe3, 0x98,
	0x3f, 0xe7, 0x6f, 0xb5, 0x5d, 0xaf, 0xed, 0xdf, 0x74, 0xf5, 0x19, 0x00, 0x00, 0xff, 0xff, 0xc1,
	0x54, 0xb8, 0xc6, 0x73, 0x02, 0x00, 0x00,
}
