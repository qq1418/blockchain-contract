// Code generated by protoc-gen-go.
// source: contract.proto
// DO NOT EDIT!

/*
Package protos is a generated protocol buffer package.

It is generated from these files:
	contract.proto
	response-data.proto

It has these top-level messages:
	ContractSignature
	ContractAsset
	ExpressionResult
	ComponentsExpression
	ComponentData
	ContractComponent
	ContractBody
	ContractHead
	Contract
	ResponseData
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

type ContractSignature struct {
	OwnerPubkey   string `protobuf:"bytes,1,opt,name=OwnerPubkey" json:"OwnerPubkey"`
	Signature     string `protobuf:"bytes,2,opt,name=Signature" json:"Signature"`
	SignTimestamp string `protobuf:"bytes,3,opt,name=SignTimestamp" json:"SignTimestamp"`
}

func (m *ContractSignature) Reset()                    { *m = ContractSignature{} }
func (m *ContractSignature) String() string            { return proto.CompactTextString(m) }
func (*ContractSignature) ProtoMessage()               {}
func (*ContractSignature) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ContractSignature) GetOwnerPubkey() string {
	if m != nil {
		return m.OwnerPubkey
	}
	return ""
}

func (m *ContractSignature) GetSignature() string {
	if m != nil {
		return m.Signature
	}
	return ""
}

func (m *ContractSignature) GetSignTimestamp() string {
	if m != nil {
		return m.SignTimestamp
	}
	return ""
}

type ContractAsset struct {
	AssetId     string  `protobuf:"bytes,1,opt,name=AssetId" json:"AssetId"`
	Name        string  `protobuf:"bytes,2,opt,name=Name" json:"Name"`
	Caption     string  `protobuf:"bytes,3,opt,name=Caption" json:"Caption"`
	Description string  `protobuf:"bytes,4,opt,name=Description" json:"Description"`
	Unit        string  `protobuf:"bytes,5,opt,name=Unit" json:"Unit"`
	Amount      float32 `protobuf:"fixed32,6,opt,name=Amount" json:"Amount"`
	MetaData    []byte  `protobuf:"bytes,7,opt,name=MetaData,proto3" json:"MetaData"`
}

func (m *ContractAsset) Reset()                    { *m = ContractAsset{} }
func (m *ContractAsset) String() string            { return proto.CompactTextString(m) }
func (*ContractAsset) ProtoMessage()               {}
func (*ContractAsset) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ContractAsset) GetAssetId() string {
	if m != nil {
		return m.AssetId
	}
	return ""
}

func (m *ContractAsset) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ContractAsset) GetCaption() string {
	if m != nil {
		return m.Caption
	}
	return ""
}

func (m *ContractAsset) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *ContractAsset) GetUnit() string {
	if m != nil {
		return m.Unit
	}
	return ""
}

func (m *ContractAsset) GetAmount() float32 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *ContractAsset) GetMetaData() []byte {
	if m != nil {
		return m.MetaData
	}
	return nil
}

type ExpressionResult struct {
	Messsage string `protobuf:"bytes,1,opt,name=Messsage" json:"Messsage"`
	Code     string `protobuf:"bytes,2,opt,name=Code" json:"Code"`
}

func (m *ExpressionResult) Reset()                    { *m = ExpressionResult{} }
func (m *ExpressionResult) String() string            { return proto.CompactTextString(m) }
func (*ExpressionResult) ProtoMessage()               {}
func (*ExpressionResult) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ExpressionResult) GetMesssage() string {
	if m != nil {
		return m.Messsage
	}
	return ""
}

func (m *ExpressionResult) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

type ComponentsExpression struct {
	Cname            string            `protobuf:"bytes,1,opt,name=Cname" json:"Cname"`
	Ctype            string            `protobuf:"bytes,2,opt,name=Ctype" json:"Ctype"`
	Caption          string            `protobuf:"bytes,3,opt,name=Caption" json:"Caption"`
	Description      string            `protobuf:"bytes,4,opt,name=Description" json:"Description"`
	ExpressionStr    string            `protobuf:"bytes,5,opt,name=ExpressionStr" json:"ExpressionStr"`
	ExpressionResult *ExpressionResult `protobuf:"bytes,6,opt,name=ExpressionResult" json:"ExpressionResult"`
	LogicValue       string            `protobuf:"bytes,7,opt,name=LogicValue" json:"LogicValue"`
}

func (m *ComponentsExpression) Reset()                    { *m = ComponentsExpression{} }
func (m *ComponentsExpression) String() string            { return proto.CompactTextString(m) }
func (*ComponentsExpression) ProtoMessage()               {}
func (*ComponentsExpression) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ComponentsExpression) GetCname() string {
	if m != nil {
		return m.Cname
	}
	return ""
}

func (m *ComponentsExpression) GetCtype() string {
	if m != nil {
		return m.Ctype
	}
	return ""
}

func (m *ComponentsExpression) GetCaption() string {
	if m != nil {
		return m.Caption
	}
	return ""
}

func (m *ComponentsExpression) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *ComponentsExpression) GetExpressionStr() string {
	if m != nil {
		return m.ExpressionStr
	}
	return ""
}

func (m *ComponentsExpression) GetExpressionResult() *ExpressionResult {
	if m != nil {
		return m.ExpressionResult
	}
	return nil
}

func (m *ComponentsExpression) GetLogicValue() string {
	if m != nil {
		return m.LogicValue
	}
	return ""
}

type ComponentData struct {
	Cname        string `protobuf:"bytes,1,opt,name=Cname" json:"Cname"`
	Ctype        string `protobuf:"bytes,2,opt,name=Ctype" json:"Ctype"`
	Caption      string `protobuf:"bytes,3,opt,name=Caption" json:"Caption"`
	Description  string `protobuf:"bytes,4,opt,name=Description" json:"Description"`
	ModifyDate   string `protobuf:"bytes,5,opt,name=ModifyDate" json:"ModifyDate"`
	HardConvType string `protobuf:"bytes,6,opt,name=HardConvType" json:"HardConvType"`
	//    map<string,?> Category = 7; // map[string]interface{} Category
	Parent    *ComponentData `protobuf:"bytes,8,opt,name=Parent" json:"Parent"`
	Mandatory bool           `protobuf:"varint,9,opt,name=Mandatory" json:"Mandatory"`
	//    google.protobuf.Any DefaultValue = 10; // interface{} DefaultValue
	Unit string `protobuf:"bytes,11,opt,name=Unit" json:"Unit"`
	//    google.protobuf.Any Value = 12; // interface{} Value
	Options map[string]int32 `protobuf:"bytes,13,rep,name=Options" json:"Options" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`
	//    repeated google.protobuf.Any DataRange = 14; // []interface{} DataRange
	Format string `protobuf:"bytes,15,opt,name=Format" json:"Format"`
}

func (m *ComponentData) Reset()                    { *m = ComponentData{} }
func (m *ComponentData) String() string            { return proto.CompactTextString(m) }
func (*ComponentData) ProtoMessage()               {}
func (*ComponentData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *ComponentData) GetCname() string {
	if m != nil {
		return m.Cname
	}
	return ""
}

func (m *ComponentData) GetCtype() string {
	if m != nil {
		return m.Ctype
	}
	return ""
}

func (m *ComponentData) GetCaption() string {
	if m != nil {
		return m.Caption
	}
	return ""
}

func (m *ComponentData) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *ComponentData) GetModifyDate() string {
	if m != nil {
		return m.ModifyDate
	}
	return ""
}

func (m *ComponentData) GetHardConvType() string {
	if m != nil {
		return m.HardConvType
	}
	return ""
}

func (m *ComponentData) GetParent() *ComponentData {
	if m != nil {
		return m.Parent
	}
	return nil
}

func (m *ComponentData) GetMandatory() bool {
	if m != nil {
		return m.Mandatory
	}
	return false
}

func (m *ComponentData) GetUnit() string {
	if m != nil {
		return m.Unit
	}
	return ""
}

func (m *ComponentData) GetOptions() map[string]int32 {
	if m != nil {
		return m.Options
	}
	return nil
}

func (m *ComponentData) GetFormat() string {
	if m != nil {
		return m.Format
	}
	return ""
}

type ContractComponent struct {
	Cname                         string                  `protobuf:"bytes,1,opt,name=Cname" json:"Cname"`
	Ctype                         string                  `protobuf:"bytes,2,opt,name=Ctype" json:"Ctype"`
	Caption                       string                  `protobuf:"bytes,3,opt,name=Caption" json:"Caption"`
	Description                   string                  `protobuf:"bytes,4,opt,name=Description" json:"Description"`
	State                         string                  `protobuf:"bytes,5,opt,name=State" json:"State"`
	PreCondition                  []*ComponentsExpression `protobuf:"bytes,6,rep,name=PreCondition" json:"PreCondition"`
	CompleteCondition             []*ComponentsExpression `protobuf:"bytes,7,rep,name=CompleteCondition" json:"CompleteCondition"`
	DisgardCondition              []*ComponentsExpression `protobuf:"bytes,8,rep,name=DisgardCondition" json:"DisgardCondition"`
	NextTasks                     []string                `protobuf:"bytes,9,rep,name=NextTasks" json:"NextTasks"`
	DataList                      []*ComponentData        `protobuf:"bytes,10,rep,name=DataList" json:"DataList"`
	DataValueSetterExpressionList []*ComponentsExpression `protobuf:"bytes,11,rep,name=DataValueSetterExpressionList" json:"DataValueSetterExpressionList"`
	CandidateList                 *ContractComponent      `protobuf:"bytes,12,opt,name=CandidateList" json:"CandidateList"`
	DecisionResult                *ContractComponent      `protobuf:"bytes,13,opt,name=DecisionResult" json:"DecisionResult"`
	TaskList                      []string                `protobuf:"bytes,14,rep,name=TaskList" json:"TaskList"`
	SupportArguments              []string                `protobuf:"bytes,15,rep,name=SupportArguments" json:"SupportArguments"`
	AgainstArguments              []string                `protobuf:"bytes,16,rep,name=AgainstArguments" json:"AgainstArguments"`
	Support                       int32                   `protobuf:"varint,17,opt,name=Support" json:"Support"`
	Text                          []string                `protobuf:"bytes,18,rep,name=Text" json:"Text"`
}

func (m *ContractComponent) Reset()                    { *m = ContractComponent{} }
func (m *ContractComponent) String() string            { return proto.CompactTextString(m) }
func (*ContractComponent) ProtoMessage()               {}
func (*ContractComponent) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *ContractComponent) GetCname() string {
	if m != nil {
		return m.Cname
	}
	return ""
}

func (m *ContractComponent) GetCtype() string {
	if m != nil {
		return m.Ctype
	}
	return ""
}

func (m *ContractComponent) GetCaption() string {
	if m != nil {
		return m.Caption
	}
	return ""
}

func (m *ContractComponent) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *ContractComponent) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func (m *ContractComponent) GetPreCondition() []*ComponentsExpression {
	if m != nil {
		return m.PreCondition
	}
	return nil
}

func (m *ContractComponent) GetCompleteCondition() []*ComponentsExpression {
	if m != nil {
		return m.CompleteCondition
	}
	return nil
}

func (m *ContractComponent) GetDisgardCondition() []*ComponentsExpression {
	if m != nil {
		return m.DisgardCondition
	}
	return nil
}

func (m *ContractComponent) GetNextTasks() []string {
	if m != nil {
		return m.NextTasks
	}
	return nil
}

func (m *ContractComponent) GetDataList() []*ComponentData {
	if m != nil {
		return m.DataList
	}
	return nil
}

func (m *ContractComponent) GetDataValueSetterExpressionList() []*ComponentsExpression {
	if m != nil {
		return m.DataValueSetterExpressionList
	}
	return nil
}

func (m *ContractComponent) GetCandidateList() *ContractComponent {
	if m != nil {
		return m.CandidateList
	}
	return nil
}

func (m *ContractComponent) GetDecisionResult() *ContractComponent {
	if m != nil {
		return m.DecisionResult
	}
	return nil
}

func (m *ContractComponent) GetTaskList() []string {
	if m != nil {
		return m.TaskList
	}
	return nil
}

func (m *ContractComponent) GetSupportArguments() []string {
	if m != nil {
		return m.SupportArguments
	}
	return nil
}

func (m *ContractComponent) GetAgainstArguments() []string {
	if m != nil {
		return m.AgainstArguments
	}
	return nil
}

func (m *ContractComponent) GetSupport() int32 {
	if m != nil {
		return m.Support
	}
	return 0
}

func (m *ContractComponent) GetText() []string {
	if m != nil {
		return m.Text
	}
	return nil
}

type ContractBody struct {
	ContractId         string               `protobuf:"bytes,1,opt,name=ContractId" json:"ContractId"`
	Cname              string               `protobuf:"bytes,2,opt,name=Cname" json:"Cname"`
	Ctype              string               `protobuf:"bytes,3,opt,name=Ctype" json:"Ctype"`
	Caption            string               `protobuf:"bytes,4,opt,name=Caption" json:"Caption"`
	Description        string               `protobuf:"bytes,5,opt,name=Description" json:"Description"`
	ContractState      string               `protobuf:"bytes,6,opt,name=ContractState" json:"ContractState"`
	Creator            string               `protobuf:"bytes,7,opt,name=Creator" json:"Creator"`
	CreatorTime        string               `protobuf:"bytes,8,opt,name=CreatorTime" json:"CreatorTime"`
	StartTime          string               `protobuf:"bytes,9,opt,name=StartTime" json:"StartTime"`
	EndTime            string               `protobuf:"bytes,10,opt,name=EndTime" json:"EndTime"`
	ContractOwners     []string             `protobuf:"bytes,11,rep,name=ContractOwners" json:"ContractOwners"`
	ContractAssets     []*ContractAsset     `protobuf:"bytes,12,rep,name=ContractAssets" json:"ContractAssets"`
	ContractSignatures []*ContractSignature `protobuf:"bytes,13,rep,name=ContractSignatures" json:"ContractSignatures"`
	ContractComponents []*ContractComponent `protobuf:"bytes,14,rep,name=ContractComponents" json:"ContractComponents"`
}

func (m *ContractBody) Reset()                    { *m = ContractBody{} }
func (m *ContractBody) String() string            { return proto.CompactTextString(m) }
func (*ContractBody) ProtoMessage()               {}
func (*ContractBody) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *ContractBody) GetContractId() string {
	if m != nil {
		return m.ContractId
	}
	return ""
}

func (m *ContractBody) GetCname() string {
	if m != nil {
		return m.Cname
	}
	return ""
}

func (m *ContractBody) GetCtype() string {
	if m != nil {
		return m.Ctype
	}
	return ""
}

func (m *ContractBody) GetCaption() string {
	if m != nil {
		return m.Caption
	}
	return ""
}

func (m *ContractBody) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *ContractBody) GetContractState() string {
	if m != nil {
		return m.ContractState
	}
	return ""
}

func (m *ContractBody) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *ContractBody) GetCreatorTime() string {
	if m != nil {
		return m.CreatorTime
	}
	return ""
}

func (m *ContractBody) GetStartTime() string {
	if m != nil {
		return m.StartTime
	}
	return ""
}

func (m *ContractBody) GetEndTime() string {
	if m != nil {
		return m.EndTime
	}
	return ""
}

func (m *ContractBody) GetContractOwners() []string {
	if m != nil {
		return m.ContractOwners
	}
	return nil
}

func (m *ContractBody) GetContractAssets() []*ContractAsset {
	if m != nil {
		return m.ContractAssets
	}
	return nil
}

func (m *ContractBody) GetContractSignatures() []*ContractSignature {
	if m != nil {
		return m.ContractSignatures
	}
	return nil
}

func (m *ContractBody) GetContractComponents() []*ContractComponent {
	if m != nil {
		return m.ContractComponents
	}
	return nil
}

type ContractHead struct {
	MainPubkey string `protobuf:"bytes,1,opt,name=MainPubkey" json:"MainPubkey"`
	Version    int32  `protobuf:"varint,2,opt,name=Version" json:"Version"`
}

func (m *ContractHead) Reset()                    { *m = ContractHead{} }
func (m *ContractHead) String() string            { return proto.CompactTextString(m) }
func (*ContractHead) ProtoMessage()               {}
func (*ContractHead) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *ContractHead) GetMainPubkey() string {
	if m != nil {
		return m.MainPubkey
	}
	return ""
}

func (m *ContractHead) GetVersion() int32 {
	if m != nil {
		return m.Version
	}
	return 0
}

type Contract struct {
	Id           string        `protobuf:"bytes,1,opt,name=id" json:"id"`
	ContractHead *ContractHead `protobuf:"bytes,2,opt,name=ContractHead" json:"ContractHead"`
	ContractBody *ContractBody `protobuf:"bytes,3,opt,name=ContractBody" json:"ContractBody"`
}

func (m *Contract) Reset()                    { *m = Contract{} }
func (m *Contract) String() string            { return proto.CompactTextString(m) }
func (*Contract) ProtoMessage()               {}
func (*Contract) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *Contract) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Contract) GetContractHead() *ContractHead {
	if m != nil {
		return m.ContractHead
	}
	return nil
}

func (m *Contract) GetContractBody() *ContractBody {
	if m != nil {
		return m.ContractBody
	}
	return nil
}

func init() {
	proto.RegisterType((*ContractSignature)(nil), "protos.ContractSignature")
	proto.RegisterType((*ContractAsset)(nil), "protos.ContractAsset")
	proto.RegisterType((*ExpressionResult)(nil), "protos.ExpressionResult")
	proto.RegisterType((*ComponentsExpression)(nil), "protos.ComponentsExpression")
	proto.RegisterType((*ComponentData)(nil), "protos.ComponentData")
	proto.RegisterType((*ContractComponent)(nil), "protos.ContractComponent")
	proto.RegisterType((*ContractBody)(nil), "protos.ContractBody")
	proto.RegisterType((*ContractHead)(nil), "protos.ContractHead")
	proto.RegisterType((*Contract)(nil), "protos.Contract")
}

func init() { proto.RegisterFile("contract.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 967 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x56, 0x6d, 0x6b, 0x1b, 0xc7,
	0x13, 0x47, 0x92, 0x25, 0x4b, 0xa3, 0x87, 0x28, 0x8b, 0xff, 0x7f, 0xae, 0x21, 0x0d, 0xea, 0x11,
	0x8a, 0x28, 0xc4, 0x50, 0xf7, 0x8d, 0x09, 0x2d, 0xad, 0x2c, 0xb9, 0x38, 0x25, 0x4e, 0xcc, 0xc9,
	0xcd, 0xfb, 0xb5, 0x6e, 0x2b, 0x8e, 0x58, 0xbb, 0xc7, 0xee, 0x2a, 0xb5, 0xbe, 0x42, 0x5f, 0xf4,
	0x9b, 0xb4, 0xdf, 0xa1, 0xd0, 0x0f, 0x56, 0x66, 0x6e, 0xf7, 0x1e, 0x24, 0xc5, 0x35, 0x14, 0xf2,
	0xea, 0x76, 0x7e, 0x3b, 0xf3, 0x9b, 0xdd, 0xb9, 0x99, 0xd9, 0x81, 0xc1, 0x42, 0x49, 0xab, 0xf9,
	0xc2, 0x1e, 0xa7, 0x5a, 0x59, 0xc5, 0x5a, 0xf4, 0x31, 0xe1, 0x06, 0x1e, 0x4f, 0xdd, 0xce, 0x3c,
	0x59, 0x4a, 0x6e, 0xd7, 0x5a, 0xb0, 0x11, 0x74, 0xdf, 0xfe, 0x2a, 0x85, 0xbe, 0x5a, 0xdf, 0xbc,
	0x17, 0x9b, 0xa0, 0x36, 0xaa, 0x8d, 0x3b, 0x51, 0x19, 0x62, 0x4f, 0xa1, 0x93, 0xab, 0x07, 0x75,
	0xda, 0x2f, 0x00, 0xf6, 0x1c, 0xfa, 0x28, 0x5c, 0x27, 0x2b, 0x61, 0x2c, 0x5f, 0xa5, 0x41, 0x83,
	0x34, 0xaa, 0x60, 0xf8, 0x57, 0x0d, 0xfa, 0xde, 0xf7, 0xc4, 0x18, 0x61, 0x59, 0x00, 0x87, 0xb4,
	0x78, 0x15, 0x3b, 0x9f, 0x5e, 0x64, 0x0c, 0x0e, 0xde, 0xf0, 0x95, 0x77, 0x45, 0x6b, 0xd4, 0x9e,
	0xf2, 0xd4, 0x26, 0x4a, 0x3a, 0x7e, 0x2f, 0xe2, 0xf9, 0x67, 0xc2, 0x2c, 0x74, 0x92, 0xed, 0x1e,
	0x64, 0xe7, 0x2f, 0x41, 0xc8, 0xf7, 0xb3, 0x4c, 0x6c, 0xd0, 0xcc, 0xf8, 0x70, 0xcd, 0xfe, 0x0f,
	0xad, 0xc9, 0x4a, 0xad, 0xa5, 0x0d, 0x5a, 0xa3, 0xda, 0xb8, 0x1e, 0x39, 0x89, 0x3d, 0x81, 0xf6,
	0xa5, 0xb0, 0x7c, 0xc6, 0x2d, 0x0f, 0x0e, 0x47, 0xb5, 0x71, 0x2f, 0xca, 0xe5, 0xf0, 0x0c, 0x86,
	0xe7, 0x77, 0xa9, 0x16, 0xc6, 0x24, 0x4a, 0x46, 0xc2, 0xac, 0x6f, 0x9d, 0xbe, 0x31, 0x86, 0x2f,
	0x85, 0xbb, 0x46, 0x2e, 0xa3, 0xdf, 0xa9, 0x8a, 0xf3, 0x7b, 0xe0, 0x3a, 0xfc, 0xad, 0x0e, 0x47,
	0x53, 0xb5, 0x4a, 0x95, 0x14, 0xd2, 0x9a, 0x82, 0x8e, 0x1d, 0x41, 0x73, 0x2a, 0xf1, 0xd6, 0x19,
	0x4b, 0x26, 0x10, 0x6a, 0x37, 0xa9, 0xe7, 0xc8, 0x84, 0xff, 0x14, 0x8c, 0xe7, 0xd0, 0x2f, 0xbc,
	0xce, 0xad, 0x76, 0x51, 0xa9, 0x82, 0x6c, 0xb6, 0x7b, 0x55, 0x0a, 0x54, 0xf7, 0x24, 0xc8, 0x72,
	0xca, 0x1c, 0x6f, 0xef, 0x47, 0xbb, 0xc1, 0x79, 0x06, 0xf0, 0x5a, 0x2d, 0x93, 0xc5, 0x3b, 0x7e,
	0xbb, 0x16, 0x14, 0xce, 0x4e, 0x54, 0x42, 0xc2, 0x3f, 0x1b, 0x98, 0x14, 0x2e, 0x18, 0x18, 0xe2,
	0x4f, 0x16, 0x85, 0x67, 0x00, 0x97, 0x2a, 0x4e, 0x7e, 0xd9, 0xcc, 0xb8, 0x15, 0x2e, 0x04, 0x25,
	0x84, 0x85, 0xd0, 0xbb, 0xe0, 0x3a, 0x9e, 0x2a, 0xf9, 0xe1, 0x1a, 0x1d, 0xb7, 0x48, 0xa3, 0x82,
	0xb1, 0x17, 0xd0, 0xba, 0xe2, 0x5a, 0x48, 0x1b, 0xb4, 0x29, 0x32, 0xff, 0xf3, 0x91, 0xa9, 0x5c,
	0x29, 0x72, 0x4a, 0x58, 0x45, 0x97, 0x5c, 0xc6, 0xdc, 0x2a, 0xbd, 0x09, 0x3a, 0xa3, 0xda, 0xb8,
	0x1d, 0x15, 0x40, 0x9e, 0xa3, 0xdd, 0x52, 0x8e, 0x7e, 0x0b, 0x87, 0x6f, 0xe9, 0xb8, 0x26, 0xe8,
	0x8f, 0x1a, 0xe3, 0xee, 0x49, 0xb8, 0xd7, 0xc3, 0xb1, 0x53, 0x3a, 0x97, 0x56, 0x6f, 0x22, 0x6f,
	0x82, 0x19, 0xfe, 0xa3, 0xd2, 0x2b, 0x6e, 0x83, 0x47, 0xc4, 0xe9, 0xa4, 0x27, 0x2f, 0xa1, 0x57,
	0x36, 0x60, 0x43, 0x68, 0x14, 0x75, 0x8f, 0x4b, 0x0c, 0xf7, 0x07, 0xfa, 0x63, 0x18, 0xee, 0x66,
	0x94, 0x09, 0x2f, 0xeb, 0xa7, 0xb5, 0xf0, 0xef, 0x56, 0xd1, 0x41, 0xf2, 0x33, 0x7c, 0xb2, 0x9f,
	0x76, 0x04, 0xcd, 0xb9, 0x2d, 0xfe, 0x57, 0x26, 0xb0, 0x1f, 0xa0, 0x77, 0xa5, 0xc5, 0x54, 0xc9,
	0x38, 0x21, 0xc3, 0x16, 0x85, 0xea, 0xe9, 0x4e, 0xa8, 0x4a, 0xc5, 0x16, 0x55, 0x2c, 0xd8, 0x4f,
	0x78, 0xa9, 0x55, 0x7a, 0x2b, 0x6c, 0x89, 0xe6, 0xf0, 0x01, 0x34, 0xbb, 0x66, 0xec, 0x02, 0x86,
	0xb3, 0xc4, 0x2c, 0xb3, 0x3c, 0x71, 0x54, 0xed, 0x07, 0x50, 0xed, 0x58, 0x61, 0xbe, 0xbc, 0x11,
	0x77, 0xf6, 0x9a, 0x9b, 0xf7, 0x26, 0xe8, 0x8c, 0x1a, 0xd8, 0x75, 0x73, 0x80, 0x7d, 0x0d, 0x6d,
	0xfc, 0xf7, 0xaf, 0x13, 0x63, 0x03, 0x20, 0xfe, 0x8f, 0xa4, 0x5f, 0xae, 0xc6, 0x6e, 0xe0, 0x73,
	0x5c, 0x53, 0xe9, 0xcd, 0x85, 0xb5, 0x42, 0x17, 0xfe, 0x89, 0xa7, 0xfb, 0x80, 0x73, 0xde, 0x4f,
	0xc1, 0xbe, 0x87, 0xfe, 0x94, 0xcb, 0x38, 0x89, 0xb9, 0x15, 0xc4, 0xd9, 0xa3, 0xd2, 0xf8, 0xac,
	0xe0, 0xdc, 0x4a, 0x9e, 0xa8, 0xaa, 0xcf, 0x26, 0x30, 0x98, 0x89, 0x45, 0x52, 0x6a, 0x3b, 0xfd,
	0x7f, 0x63, 0xd8, 0x32, 0xc0, 0x96, 0x8c, 0x31, 0x22, 0xf7, 0x03, 0x8a, 0x5b, 0x2e, 0xb3, 0xaf,
	0x60, 0x38, 0x5f, 0xa7, 0xa9, 0xd2, 0x76, 0xa2, 0x97, 0xeb, 0x15, 0x5e, 0x2e, 0x78, 0x44, 0x3a,
	0x3b, 0x38, 0xea, 0x4e, 0x96, 0x3c, 0x91, 0xa6, 0xa4, 0x3b, 0xcc, 0x74, 0xb7, 0x71, 0x4c, 0x6b,
	0x67, 0x1f, 0x3c, 0xa6, 0xa2, 0xf1, 0x22, 0x16, 0xf6, 0xb5, 0xb8, 0xb3, 0x01, 0x23, 0x4b, 0x5a,
	0x87, 0x7f, 0x1c, 0x40, 0xcf, 0xdf, 0xe3, 0x4c, 0xc5, 0x1b, 0x6c, 0x47, 0x5e, 0xce, 0x9f, 0xc3,
	0x12, 0x52, 0x54, 0x58, 0x7d, 0x6f, 0x85, 0x35, 0x3e, 0x52, 0x61, 0x07, 0xf7, 0x56, 0x58, 0x73,
	0xef, 0xe3, 0x90, 0x0f, 0x08, 0x54, 0x69, 0x59, 0xdf, 0xab, 0x82, 0xe4, 0x41, 0x0b, 0xec, 0x5b,
	0xae, 0xa7, 0x7b, 0x11, 0x3d, 0xb8, 0x25, 0xbe, 0xfc, 0xd4, 0x17, 0x3b, 0x51, 0x19, 0xa2, 0x59,
	0xc2, 0x72, 0x6d, 0x69, 0xbf, 0xe3, 0x66, 0x09, 0x0f, 0x20, 0xf3, 0xb9, 0x8c, 0x69, 0x0f, 0x32,
	0x66, 0x27, 0xb2, 0x2f, 0x61, 0xe0, 0x0f, 0x41, 0xa3, 0x89, 0xa1, 0x6c, 0xed, 0x44, 0x5b, 0x28,
	0xfb, 0xae, 0xd0, 0xa3, 0x71, 0xc2, 0x04, 0xbd, 0xed, 0xea, 0x28, 0xed, 0x46, 0x5b, 0xca, 0xec,
	0x15, 0xb0, 0x9d, 0x09, 0xc9, 0x77, 0xdf, 0x9d, 0x14, 0xcc, 0x35, 0xa2, 0x3d, 0x46, 0x65, 0xaa,
	0xa2, 0x92, 0x28, 0x21, 0xef, 0xcd, 0xe6, 0x3d, 0x46, 0xe1, 0x45, 0x91, 0x2e, 0x17, 0x82, 0xc7,
	0xf4, 0x7a, 0xf1, 0x44, 0x56, 0x26, 0xb6, 0x12, 0x82, 0x61, 0x7c, 0x27, 0x34, 0x96, 0x84, 0x6b,
	0xe1, 0x5e, 0x0c, 0x7f, 0xaf, 0x41, 0xdb, 0x53, 0xb1, 0x01, 0xd4, 0x13, 0x9f, 0x6d, 0xf5, 0x24,
	0x66, 0xa7, 0x55, 0x37, 0x64, 0xdb, 0x3d, 0x39, 0xda, 0x3e, 0x2b, 0xee, 0x45, 0xd5, 0x03, 0x9d,
	0x56, 0xf3, 0x99, 0x12, 0x72, 0x8f, 0x25, 0xee, 0x45, 0x15, 0xcd, 0xb3, 0x13, 0xf8, 0x62, 0xa1,
	0x56, 0xc7, 0x6b, 0x99, 0xbc, 0xb8, 0x15, 0xf1, 0x52, 0x68, 0x5c, 0x56, 0xc7, 0x57, 0x73, 0xd6,
	0xbf, 0xc2, 0xaf, 0xb7, 0xbb, 0xc9, 0xc6, 0xd9, 0x6f, 0xfe, 0x09, 0x00, 0x00, 0xff, 0xff, 0x21,
	0xa9, 0x90, 0x78, 0xe7, 0x0a, 0x00, 0x00,
}
