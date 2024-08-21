// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ves.io/schema/marketplace/xc_saas/asb/message/types.proto

package message

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
	golang_proto "github.com/golang/protobuf/proto"
	schema "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema"
	signup "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/signup"
	io "io"
	math "math"
	math_bits "math/bits"
	reflect "reflect"
	strconv "strconv"
	strings "strings"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = golang_proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// ActionType
//
// x-displayName: "Action Type"
// Type of action to perform
type ActionType int32

const (
	// x-displayName: "ENABLE"
	// Provision the requested service for tenant creation
	ENABLE ActionType = 0
)

var ActionType_name = map[int32]string{
	0: "ENABLE",
}

var ActionType_value = map[string]int32{
	"ENABLE": 0,
}

func (ActionType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_fd0ad07b3438cde8, []int{0}
}

// GlobalSpecType
//
// x-displayName: "Specification"
// Shape of message in the storage backend.
type GlobalSpecType struct {
	// Raw JSON
	//
	// x-displayName: "Raw JSON"
	// x-example: "{"documentType":"teem-service-order","documentVersion":"1.0","serviceOrderType":"provisioning"}"
	// Raw JSON message string that holds metadata and payload related to provisioning
	RawJson *types.Struct `protobuf:"bytes,1,opt,name=raw_json,json=rawJson,proto3" json:"raw_json,omitempty"`
	// Parsed Type
	//
	// x-displayName: "Parsed Type"
	// Parsed type is the parsed message representation as per the use case.
	//
	// Types that are valid to be assigned to ParsedType:
	//	*GlobalSpecType_None
	//	*GlobalSpecType_SignupMessage
	ParsedType isGlobalSpecType_ParsedType `protobuf_oneof:"parsed_type"`
}

func (m *GlobalSpecType) Reset()      { *m = GlobalSpecType{} }
func (*GlobalSpecType) ProtoMessage() {}
func (*GlobalSpecType) Descriptor() ([]byte, []int) {
	return fileDescriptor_fd0ad07b3438cde8, []int{0}
}
func (m *GlobalSpecType) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GlobalSpecType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *GlobalSpecType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GlobalSpecType.Merge(m, src)
}
func (m *GlobalSpecType) XXX_Size() int {
	return m.Size()
}
func (m *GlobalSpecType) XXX_DiscardUnknown() {
	xxx_messageInfo_GlobalSpecType.DiscardUnknown(m)
}

var xxx_messageInfo_GlobalSpecType proto.InternalMessageInfo

type isGlobalSpecType_ParsedType interface {
	isGlobalSpecType_ParsedType()
	Equal(interface{}) bool
	MarshalTo([]byte) (int, error)
	Size() int
}

type GlobalSpecType_None struct {
	None *schema.Empty `protobuf:"bytes,3,opt,name=none,proto3,oneof" json:"none,omitempty"`
}
type GlobalSpecType_SignupMessage struct {
	SignupMessage *SignupMessage `protobuf:"bytes,4,opt,name=signup_message,json=signupMessage,proto3,oneof" json:"signup_message,omitempty"`
}

func (*GlobalSpecType_None) isGlobalSpecType_ParsedType()          {}
func (*GlobalSpecType_SignupMessage) isGlobalSpecType_ParsedType() {}

func (m *GlobalSpecType) GetParsedType() isGlobalSpecType_ParsedType {
	if m != nil {
		return m.ParsedType
	}
	return nil
}

func (m *GlobalSpecType) GetRawJson() *types.Struct {
	if m != nil {
		return m.RawJson
	}
	return nil
}

func (m *GlobalSpecType) GetNone() *schema.Empty {
	if x, ok := m.GetParsedType().(*GlobalSpecType_None); ok {
		return x.None
	}
	return nil
}

func (m *GlobalSpecType) GetSignupMessage() *SignupMessage {
	if x, ok := m.GetParsedType().(*GlobalSpecType_SignupMessage); ok {
		return x.SignupMessage
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*GlobalSpecType) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*GlobalSpecType_None)(nil),
		(*GlobalSpecType_SignupMessage)(nil),
	}
}

// Signup Message
//
// x-displayName: "Signup Message"
// Signup message holds customer information and other related provisioning information needed to sign up the
// customer for F5XC services.
type SignupMessage struct {
	// CRM Details
	//
	// x-displayName: "CRM Details"
	// This field holds CRM information
	CrmDetails *schema.CRMInfo `protobuf:"bytes,1,opt,name=crm_details,json=crmDetails,proto3" json:"crm_details,omitempty"`
	// Company Details
	//
	// x-displayName: "Company Details"
	// x-required
	// Details of the company
	CompanyDetails *signup.CompanyMeta `protobuf:"bytes,2,opt,name=company_details,json=companyDetails,proto3" json:"company_details,omitempty"`
	// User Details
	//
	// x-displayName: "User Details"
	// x-required
	// Details of the user
	UserDetails *signup.UserMeta `protobuf:"bytes,3,opt,name=user_details,json=userDetails,proto3" json:"user_details,omitempty"`
	// Account Details
	//
	// x-displayName: "Account Details"
	// x-required
	// Details of the new F5XC account to be created
	AccountDetails *signup.AccountMeta `protobuf:"bytes,4,opt,name=account_details,json=accountDetails,proto3" json:"account_details,omitempty"`
	// Action Type
	//
	// x-displayName: "Action Type"
	// x-required
	// Type of action to perform
	ActionType ActionType `protobuf:"varint,5,opt,name=action_type,json=actionType,proto3,enum=ves.io.schema.marketplace.xc_saas.asb.message.ActionType" json:"action_type,omitempty"`
	// Tenant ID
	//
	// x-displayName: "Tenant ID"
	// Identifier for the corresponding tenant object
	TenantId string `protobuf:"bytes,6,opt,name=tenant_id,json=tenantId,proto3" json:"tenant_id,omitempty"`
	// Signup ID
	//
	// x-displayName: "Signup ID"
	// Identifier for the corresponding signup object
	SignupId string `protobuf:"bytes,7,opt,name=signup_id,json=signupId,proto3" json:"signup_id,omitempty"`
}

func (m *SignupMessage) Reset()      { *m = SignupMessage{} }
func (*SignupMessage) ProtoMessage() {}
func (*SignupMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_fd0ad07b3438cde8, []int{1}
}
func (m *SignupMessage) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SignupMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *SignupMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignupMessage.Merge(m, src)
}
func (m *SignupMessage) XXX_Size() int {
	return m.Size()
}
func (m *SignupMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_SignupMessage.DiscardUnknown(m)
}

var xxx_messageInfo_SignupMessage proto.InternalMessageInfo

func (m *SignupMessage) GetCrmDetails() *schema.CRMInfo {
	if m != nil {
		return m.CrmDetails
	}
	return nil
}

func (m *SignupMessage) GetCompanyDetails() *signup.CompanyMeta {
	if m != nil {
		return m.CompanyDetails
	}
	return nil
}

func (m *SignupMessage) GetUserDetails() *signup.UserMeta {
	if m != nil {
		return m.UserDetails
	}
	return nil
}

func (m *SignupMessage) GetAccountDetails() *signup.AccountMeta {
	if m != nil {
		return m.AccountDetails
	}
	return nil
}

func (m *SignupMessage) GetActionType() ActionType {
	if m != nil {
		return m.ActionType
	}
	return ENABLE
}

func (m *SignupMessage) GetTenantId() string {
	if m != nil {
		return m.TenantId
	}
	return ""
}

func (m *SignupMessage) GetSignupId() string {
	if m != nil {
		return m.SignupId
	}
	return ""
}

func init() {
	proto.RegisterEnum("ves.io.schema.marketplace.xc_saas.asb.message.ActionType", ActionType_name, ActionType_value)
	golang_proto.RegisterEnum("ves.io.schema.marketplace.xc_saas.asb.message.ActionType", ActionType_name, ActionType_value)
	proto.RegisterType((*GlobalSpecType)(nil), "ves.io.schema.marketplace.xc_saas.asb.message.GlobalSpecType")
	golang_proto.RegisterType((*GlobalSpecType)(nil), "ves.io.schema.marketplace.xc_saas.asb.message.GlobalSpecType")
	proto.RegisterType((*SignupMessage)(nil), "ves.io.schema.marketplace.xc_saas.asb.message.SignupMessage")
	golang_proto.RegisterType((*SignupMessage)(nil), "ves.io.schema.marketplace.xc_saas.asb.message.SignupMessage")
}

func init() {
	proto.RegisterFile("ves.io/schema/marketplace/xc_saas/asb/message/types.proto", fileDescriptor_fd0ad07b3438cde8)
}
func init() {
	golang_proto.RegisterFile("ves.io/schema/marketplace/xc_saas/asb/message/types.proto", fileDescriptor_fd0ad07b3438cde8)
}

var fileDescriptor_fd0ad07b3438cde8 = []byte{
	// 658 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0xbf, 0x4e, 0x1b, 0x31,
	0x18, 0x3f, 0x93, 0x14, 0x82, 0x53, 0x52, 0x74, 0xaa, 0x4a, 0x1a, 0x90, 0x9b, 0xd2, 0x05, 0x21,
	0xc5, 0x27, 0xa5, 0x43, 0x85, 0xd4, 0x25, 0x50, 0x54, 0x40, 0xa5, 0xc3, 0xd1, 0x2e, 0x2c, 0x27,
	0xe7, 0xce, 0x1c, 0x57, 0x72, 0x67, 0xcb, 0x76, 0x02, 0xd9, 0xfa, 0x08, 0x7d, 0x85, 0x6e, 0x7d,
	0x85, 0x8a, 0x85, 0xb1, 0x63, 0x46, 0xc6, 0x72, 0x59, 0x3a, 0x32, 0x75, 0xae, 0x62, 0x5f, 0x42,
	0x2e, 0xaa, 0xaa, 0xb2, 0xd9, 0xfe, 0xfd, 0xb9, 0xef, 0xf7, 0x7d, 0xf6, 0xc1, 0xad, 0x1e, 0x95,
	0x38, 0x62, 0x8e, 0xf4, 0x4f, 0x69, 0x4c, 0x9c, 0x98, 0x88, 0x33, 0xaa, 0x78, 0x87, 0xf8, 0xd4,
	0xb9, 0xf0, 0x3d, 0x49, 0x88, 0x74, 0x88, 0x6c, 0x3b, 0x31, 0x95, 0x92, 0x84, 0xd4, 0x51, 0x7d,
	0x4e, 0x25, 0xe6, 0x82, 0x29, 0x66, 0x37, 0x8c, 0x14, 0x1b, 0x29, 0x9e, 0x92, 0xe2, 0x4c, 0x8a,
	0x89, 0x6c, 0xe3, 0x4c, 0x5a, 0x6b, 0x84, 0x91, 0x3a, 0xed, 0xb6, 0xb1, 0xcf, 0x62, 0x27, 0x64,
	0x21, 0x73, 0xb4, 0x4b, 0xbb, 0x7b, 0xa2, 0x77, 0x7a, 0xa3, 0x57, 0xc6, 0xbd, 0xf6, 0x34, 0x64,
	0x2c, 0xec, 0xd0, 0x3b, 0x16, 0x49, 0xfa, 0x19, 0xb4, 0x36, 0x0b, 0x49, 0x25, 0xba, 0xbe, 0xca,
	0xd0, 0x67, 0xb3, 0xa8, 0x8a, 0x62, 0x2a, 0x15, 0x89, 0x79, 0x46, 0x58, 0xcd, 0x47, 0x66, 0x5c,
	0x45, 0x2c, 0xc9, 0x42, 0xd5, 0x5e, 0xe4, 0x41, 0x19, 0x85, 0x49, 0x97, 0x9b, 0xd8, 0x5e, 0xaf,
	0x39, 0xae, 0x2d, 0x4f, 0x9a, 0x6a, 0x4a, 0x6d, 0x2d, 0x0f, 0xf5, 0x48, 0x27, 0x0a, 0x88, 0xa2,
	0x19, 0x5a, 0x9f, 0x41, 0x23, 0x7a, 0xee, 0xe5, 0xbe, 0xbf, 0xfe, 0x1b, 0xc0, 0xca, 0xdb, 0x0e,
	0x6b, 0x93, 0xce, 0x11, 0xa7, 0xfe, 0x87, 0x3e, 0xa7, 0x76, 0x13, 0x96, 0x04, 0x39, 0xf7, 0x3e,
	0x49, 0x96, 0x54, 0x41, 0x1d, 0x6c, 0x94, 0x9b, 0x2b, 0xd8, 0x64, 0xc4, 0xe3, 0x8c, 0xf8, 0x48,
	0x77, 0xc0, 0x5d, 0x10, 0xe4, 0xfc, 0x40, 0xb2, 0xc4, 0xde, 0x84, 0xc5, 0x84, 0x25, 0xb4, 0x5a,
	0xd0, 0xfc, 0xc7, 0x38, 0x3f, 0xaa, 0xdd, 0x98, 0xab, 0xfe, 0x9e, 0xe5, 0x6a, 0x8e, 0x4d, 0x61,
	0xc5, 0xc4, 0xf4, 0xb2, 0x51, 0x55, 0x8b, 0x5a, 0xf5, 0x1a, 0xdf, 0x6b, 0xc0, 0xf8, 0x48, 0x9b,
	0x1c, 0x9a, 0xdd, 0x9e, 0xe5, 0x2e, 0xc9, 0xe9, 0x83, 0xed, 0x15, 0x58, 0xe6, 0x44, 0x48, 0x1a,
	0x78, 0xa3, 0x7e, 0xd9, 0xa5, 0xab, 0x4b, 0x50, 0x18, 0x5c, 0x82, 0xb9, 0x83, 0x62, 0x69, 0x6e,
	0xb9, 0xb0, 0xfe, 0xbd, 0x00, 0x97, 0x72, 0x0e, 0xf6, 0x2b, 0x58, 0xf6, 0x45, 0xec, 0x05, 0x54,
	0x91, 0xa8, 0x23, 0xb3, 0xe8, 0x4f, 0x66, 0x8a, 0xda, 0x71, 0x0f, 0xf7, 0x93, 0x13, 0xe6, 0x42,
	0x5f, 0xc4, 0x6f, 0x0c, 0xd3, 0x3e, 0x80, 0x8f, 0x7c, 0x16, 0x73, 0x92, 0xf4, 0x27, 0xe2, 0x39,
	0x2d, 0x7e, 0x3e, 0x23, 0x36, 0x05, 0xe2, 0x1d, 0x43, 0x3e, 0xa4, 0x8a, 0xb8, 0x95, 0x4c, 0x39,
	0xf6, 0x6a, 0xc1, 0x87, 0x5d, 0x49, 0xc5, 0xc4, 0xc8, 0x34, 0x14, 0xfd, 0xdd, 0xe8, 0xa3, 0xa4,
	0x42, 0xbb, 0x94, 0x47, 0x9a, 0xa9, 0x72, 0x88, 0xef, 0xb3, 0x6e, 0xa2, 0x26, 0x2e, 0xc5, 0x7f,
	0x95, 0xd3, 0x32, 0x64, 0x53, 0x4e, 0xa6, 0x1c, 0x7b, 0x1d, 0xc3, 0x32, 0xf1, 0x47, 0xf7, 0x45,
	0x37, 0xb1, 0xfa, 0xa0, 0x0e, 0x36, 0x2a, 0xcd, 0xad, 0x7b, 0x0e, 0xaa, 0xa5, 0x1d, 0x46, 0x77,
	0xcb, 0x85, 0x64, 0xb2, 0xb6, 0x57, 0xe1, 0xa2, 0xa2, 0x09, 0x49, 0x94, 0x17, 0x05, 0xd5, 0xf9,
	0x3a, 0xd8, 0x58, 0x74, 0x4b, 0xe6, 0x60, 0x3f, 0x18, 0x81, 0xd9, 0x25, 0x89, 0x82, 0xea, 0x82,
	0x01, 0xcd, 0xc1, 0x7e, 0xb0, 0x59, 0x85, 0xf0, 0xce, 0xd3, 0x86, 0x70, 0x7e, 0xf7, 0x7d, 0x6b,
	0xfb, 0xdd, 0xee, 0xb2, 0xb5, 0xfd, 0x15, 0x0c, 0x6e, 0x90, 0x75, 0x7d, 0x83, 0xac, 0xdb, 0x1b,
	0x04, 0x3e, 0xa7, 0x08, 0x7c, 0x4b, 0x11, 0xf8, 0x91, 0x22, 0x30, 0x48, 0x11, 0xb8, 0x4e, 0x11,
	0xf8, 0x99, 0x22, 0xf0, 0x2b, 0x45, 0xd6, 0x6d, 0x8a, 0xc0, 0x97, 0x21, 0xb2, 0xae, 0x86, 0x08,
	0x0c, 0x86, 0xc8, 0xba, 0x1e, 0x22, 0xeb, 0xd8, 0x0f, 0x19, 0x3f, 0x0b, 0x71, 0x8f, 0x75, 0x14,
	0x15, 0x82, 0xe0, 0xae, 0x74, 0xf4, 0xe2, 0x84, 0x89, 0xb8, 0xc1, 0x05, 0xeb, 0x45, 0x01, 0x15,
	0x8d, 0x31, 0xec, 0xf0, 0x76, 0xc8, 0x1c, 0x7a, 0xa1, 0xc6, 0x8f, 0xf7, 0xbf, 0xfe, 0x69, 0xed,
	0x79, 0xfd, 0x8a, 0x5e, 0xfe, 0x09, 0x00, 0x00, 0xff, 0xff, 0x0e, 0xe7, 0x0f, 0xff, 0x0b, 0x05,
	0x00, 0x00,
}

func (x ActionType) String() string {
	s, ok := ActionType_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (this *GlobalSpecType) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*GlobalSpecType)
	if !ok {
		that2, ok := that.(GlobalSpecType)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.RawJson.Equal(that1.RawJson) {
		return false
	}
	if that1.ParsedType == nil {
		if this.ParsedType != nil {
			return false
		}
	} else if this.ParsedType == nil {
		return false
	} else if !this.ParsedType.Equal(that1.ParsedType) {
		return false
	}
	return true
}
func (this *GlobalSpecType_None) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*GlobalSpecType_None)
	if !ok {
		that2, ok := that.(GlobalSpecType_None)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.None.Equal(that1.None) {
		return false
	}
	return true
}
func (this *GlobalSpecType_SignupMessage) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*GlobalSpecType_SignupMessage)
	if !ok {
		that2, ok := that.(GlobalSpecType_SignupMessage)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.SignupMessage.Equal(that1.SignupMessage) {
		return false
	}
	return true
}
func (this *SignupMessage) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*SignupMessage)
	if !ok {
		that2, ok := that.(SignupMessage)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.CrmDetails.Equal(that1.CrmDetails) {
		return false
	}
	if !this.CompanyDetails.Equal(that1.CompanyDetails) {
		return false
	}
	if !this.UserDetails.Equal(that1.UserDetails) {
		return false
	}
	if !this.AccountDetails.Equal(that1.AccountDetails) {
		return false
	}
	if this.ActionType != that1.ActionType {
		return false
	}
	if this.TenantId != that1.TenantId {
		return false
	}
	if this.SignupId != that1.SignupId {
		return false
	}
	return true
}
func (this *GlobalSpecType) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 7)
	s = append(s, "&message.GlobalSpecType{")
	if this.RawJson != nil {
		s = append(s, "RawJson: "+fmt.Sprintf("%#v", this.RawJson)+",\n")
	}
	if this.ParsedType != nil {
		s = append(s, "ParsedType: "+fmt.Sprintf("%#v", this.ParsedType)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *GlobalSpecType_None) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&message.GlobalSpecType_None{` +
		`None:` + fmt.Sprintf("%#v", this.None) + `}`}, ", ")
	return s
}
func (this *GlobalSpecType_SignupMessage) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&message.GlobalSpecType_SignupMessage{` +
		`SignupMessage:` + fmt.Sprintf("%#v", this.SignupMessage) + `}`}, ", ")
	return s
}
func (this *SignupMessage) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 11)
	s = append(s, "&message.SignupMessage{")
	if this.CrmDetails != nil {
		s = append(s, "CrmDetails: "+fmt.Sprintf("%#v", this.CrmDetails)+",\n")
	}
	if this.CompanyDetails != nil {
		s = append(s, "CompanyDetails: "+fmt.Sprintf("%#v", this.CompanyDetails)+",\n")
	}
	if this.UserDetails != nil {
		s = append(s, "UserDetails: "+fmt.Sprintf("%#v", this.UserDetails)+",\n")
	}
	if this.AccountDetails != nil {
		s = append(s, "AccountDetails: "+fmt.Sprintf("%#v", this.AccountDetails)+",\n")
	}
	s = append(s, "ActionType: "+fmt.Sprintf("%#v", this.ActionType)+",\n")
	s = append(s, "TenantId: "+fmt.Sprintf("%#v", this.TenantId)+",\n")
	s = append(s, "SignupId: "+fmt.Sprintf("%#v", this.SignupId)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringTypes(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *GlobalSpecType) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GlobalSpecType) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GlobalSpecType) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ParsedType != nil {
		{
			size := m.ParsedType.Size()
			i -= size
			if _, err := m.ParsedType.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	if m.RawJson != nil {
		{
			size, err := m.RawJson.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *GlobalSpecType_None) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GlobalSpecType_None) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.None != nil {
		{
			size, err := m.None.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	return len(dAtA) - i, nil
}
func (m *GlobalSpecType_SignupMessage) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GlobalSpecType_SignupMessage) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.SignupMessage != nil {
		{
			size, err := m.SignupMessage.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	return len(dAtA) - i, nil
}
func (m *SignupMessage) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SignupMessage) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SignupMessage) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.SignupId) > 0 {
		i -= len(m.SignupId)
		copy(dAtA[i:], m.SignupId)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.SignupId)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.TenantId) > 0 {
		i -= len(m.TenantId)
		copy(dAtA[i:], m.TenantId)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.TenantId)))
		i--
		dAtA[i] = 0x32
	}
	if m.ActionType != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.ActionType))
		i--
		dAtA[i] = 0x28
	}
	if m.AccountDetails != nil {
		{
			size, err := m.AccountDetails.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if m.UserDetails != nil {
		{
			size, err := m.UserDetails.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.CompanyDetails != nil {
		{
			size, err := m.CompanyDetails.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.CrmDetails != nil {
		{
			size, err := m.CrmDetails.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintTypes(dAtA []byte, offset int, v uint64) int {
	offset -= sovTypes(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GlobalSpecType) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.RawJson != nil {
		l = m.RawJson.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.ParsedType != nil {
		n += m.ParsedType.Size()
	}
	return n
}

func (m *GlobalSpecType_None) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.None != nil {
		l = m.None.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}
func (m *GlobalSpecType_SignupMessage) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.SignupMessage != nil {
		l = m.SignupMessage.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}
func (m *SignupMessage) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.CrmDetails != nil {
		l = m.CrmDetails.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.CompanyDetails != nil {
		l = m.CompanyDetails.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.UserDetails != nil {
		l = m.UserDetails.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.AccountDetails != nil {
		l = m.AccountDetails.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.ActionType != 0 {
		n += 1 + sovTypes(uint64(m.ActionType))
	}
	l = len(m.TenantId)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.SignupId)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}

func sovTypes(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTypes(x uint64) (n int) {
	return sovTypes(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *GlobalSpecType) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&GlobalSpecType{`,
		`RawJson:` + strings.Replace(fmt.Sprintf("%v", this.RawJson), "Struct", "types.Struct", 1) + `,`,
		`ParsedType:` + fmt.Sprintf("%v", this.ParsedType) + `,`,
		`}`,
	}, "")
	return s
}
func (this *GlobalSpecType_None) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&GlobalSpecType_None{`,
		`None:` + strings.Replace(fmt.Sprintf("%v", this.None), "Empty", "schema.Empty", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *GlobalSpecType_SignupMessage) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&GlobalSpecType_SignupMessage{`,
		`SignupMessage:` + strings.Replace(fmt.Sprintf("%v", this.SignupMessage), "SignupMessage", "SignupMessage", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *SignupMessage) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&SignupMessage{`,
		`CrmDetails:` + strings.Replace(fmt.Sprintf("%v", this.CrmDetails), "CRMInfo", "schema.CRMInfo", 1) + `,`,
		`CompanyDetails:` + strings.Replace(fmt.Sprintf("%v", this.CompanyDetails), "CompanyMeta", "signup.CompanyMeta", 1) + `,`,
		`UserDetails:` + strings.Replace(fmt.Sprintf("%v", this.UserDetails), "UserMeta", "signup.UserMeta", 1) + `,`,
		`AccountDetails:` + strings.Replace(fmt.Sprintf("%v", this.AccountDetails), "AccountMeta", "signup.AccountMeta", 1) + `,`,
		`ActionType:` + fmt.Sprintf("%v", this.ActionType) + `,`,
		`TenantId:` + fmt.Sprintf("%v", this.TenantId) + `,`,
		`SignupId:` + fmt.Sprintf("%v", this.SignupId) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringTypes(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *GlobalSpecType) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GlobalSpecType: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GlobalSpecType: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RawJson", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.RawJson == nil {
				m.RawJson = &types.Struct{}
			}
			if err := m.RawJson.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field None", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &schema.Empty{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.ParsedType = &GlobalSpecType_None{v}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SignupMessage", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &SignupMessage{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.ParsedType = &GlobalSpecType_SignupMessage{v}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *SignupMessage) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: SignupMessage: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SignupMessage: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CrmDetails", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.CrmDetails == nil {
				m.CrmDetails = &schema.CRMInfo{}
			}
			if err := m.CrmDetails.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CompanyDetails", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.CompanyDetails == nil {
				m.CompanyDetails = &signup.CompanyMeta{}
			}
			if err := m.CompanyDetails.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserDetails", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.UserDetails == nil {
				m.UserDetails = &signup.UserMeta{}
			}
			if err := m.UserDetails.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AccountDetails", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.AccountDetails == nil {
				m.AccountDetails = &signup.AccountMeta{}
			}
			if err := m.AccountDetails.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ActionType", wireType)
			}
			m.ActionType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ActionType |= ActionType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TenantId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TenantId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SignupId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SignupId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTypes(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTypes
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthTypes
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTypes
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTypes
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTypes        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTypes          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTypes = fmt.Errorf("proto: unexpected end of group")
)