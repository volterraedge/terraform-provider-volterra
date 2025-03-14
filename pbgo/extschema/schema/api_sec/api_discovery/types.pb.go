// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ves.io/schema/api_sec/api_discovery/types.proto

package api_discovery

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	golang_proto "github.com/golang/protobuf/proto"
	_ "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema"
	_ "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/views"
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

// AuthParameterType
//
// x-displayName: "Authentication Parameter Type"
// Enumeration for authentication parameter types
type AuthParameterType int32

const (
	// x-displayName: "Query Parameter"
	QUERY_PARAMETER AuthParameterType = 0
	// x-displayName: "Header"
	HEADER AuthParameterType = 1
	// x-displayName: "Cookie"
	COOKIE AuthParameterType = 2
)

var AuthParameterType_name = map[int32]string{
	0: "QUERY_PARAMETER",
	1: "HEADER",
	2: "COOKIE",
}

var AuthParameterType_value = map[string]int32{
	"QUERY_PARAMETER": 0,
	"HEADER":          1,
	"COOKIE":          2,
}

func (AuthParameterType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_3ccd12d9c64560a6, []int{0}
}

// CustomAuthType
//
// x-displayName: "Custom Auth Type"
// Customize Authentication Type
type CustomAuthType struct {
	// parameter_type
	//
	// x-required
	// x-displayName: "Parameter Type"
	ParameterType AuthParameterType `protobuf:"varint,1,opt,name=parameter_type,json=parameterType,proto3,enum=ves.io.schema.api_sec.api_discovery.AuthParameterType" json:"parameter_type,omitempty"`
	// parameter name
	//
	// x-required
	// x-displayName: "Parameter Name"
	// x-example: "x-auth-header"
	// The authentication parameter name.
	ParameterName string `protobuf:"bytes,2,opt,name=parameter_name,json=parameterName,proto3" json:"parameter_name,omitempty"`
}

func (m *CustomAuthType) Reset()      { *m = CustomAuthType{} }
func (*CustomAuthType) ProtoMessage() {}
func (*CustomAuthType) Descriptor() ([]byte, []int) {
	return fileDescriptor_3ccd12d9c64560a6, []int{0}
}
func (m *CustomAuthType) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CustomAuthType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *CustomAuthType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CustomAuthType.Merge(m, src)
}
func (m *CustomAuthType) XXX_Size() int {
	return m.Size()
}
func (m *CustomAuthType) XXX_DiscardUnknown() {
	xxx_messageInfo_CustomAuthType.DiscardUnknown(m)
}

var xxx_messageInfo_CustomAuthType proto.InternalMessageInfo

func (m *CustomAuthType) GetParameterType() AuthParameterType {
	if m != nil {
		return m.ParameterType
	}
	return QUERY_PARAMETER
}

func (m *CustomAuthType) GetParameterName() string {
	if m != nil {
		return m.ParameterName
	}
	return ""
}

// GlobalSpecType
//
// x-displayName: "Specification"
// Shape of api_discovery in the storage backend.
type GlobalSpecType struct {
	// custom_auth_types
	//
	// x-displayName: "Custom Authentication Types"
	// Select your custom authentication types to be detected in the API discovery
	CustomAuthTypes []*CustomAuthType `protobuf:"bytes,1,rep,name=custom_auth_types,json=customAuthTypes,proto3" json:"custom_auth_types,omitempty"`
}

func (m *GlobalSpecType) Reset()      { *m = GlobalSpecType{} }
func (*GlobalSpecType) ProtoMessage() {}
func (*GlobalSpecType) Descriptor() ([]byte, []int) {
	return fileDescriptor_3ccd12d9c64560a6, []int{1}
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

func (m *GlobalSpecType) GetCustomAuthTypes() []*CustomAuthType {
	if m != nil {
		return m.CustomAuthTypes
	}
	return nil
}

// Create api discovery
//
// x-displayName: "Create Api Discovery"
// Create api discovery creates a new object in the storage backend for metadata.namespace.
type CreateSpecType struct {
	CustomAuthTypes []*CustomAuthType `protobuf:"bytes,1,rep,name=custom_auth_types,json=customAuthTypes,proto3" json:"custom_auth_types,omitempty"`
}

func (m *CreateSpecType) Reset()      { *m = CreateSpecType{} }
func (*CreateSpecType) ProtoMessage() {}
func (*CreateSpecType) Descriptor() ([]byte, []int) {
	return fileDescriptor_3ccd12d9c64560a6, []int{2}
}
func (m *CreateSpecType) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CreateSpecType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *CreateSpecType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateSpecType.Merge(m, src)
}
func (m *CreateSpecType) XXX_Size() int {
	return m.Size()
}
func (m *CreateSpecType) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateSpecType.DiscardUnknown(m)
}

var xxx_messageInfo_CreateSpecType proto.InternalMessageInfo

func (m *CreateSpecType) GetCustomAuthTypes() []*CustomAuthType {
	if m != nil {
		return m.CustomAuthTypes
	}
	return nil
}

// Replace api discovery
//
// x-displayName: "Replace Api Discovery"
// Replace api_discovery replaces an existing object in the storage backend for metadata.namespace.
type ReplaceSpecType struct {
	CustomAuthTypes []*CustomAuthType `protobuf:"bytes,1,rep,name=custom_auth_types,json=customAuthTypes,proto3" json:"custom_auth_types,omitempty"`
}

func (m *ReplaceSpecType) Reset()      { *m = ReplaceSpecType{} }
func (*ReplaceSpecType) ProtoMessage() {}
func (*ReplaceSpecType) Descriptor() ([]byte, []int) {
	return fileDescriptor_3ccd12d9c64560a6, []int{3}
}
func (m *ReplaceSpecType) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ReplaceSpecType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *ReplaceSpecType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplaceSpecType.Merge(m, src)
}
func (m *ReplaceSpecType) XXX_Size() int {
	return m.Size()
}
func (m *ReplaceSpecType) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplaceSpecType.DiscardUnknown(m)
}

var xxx_messageInfo_ReplaceSpecType proto.InternalMessageInfo

func (m *ReplaceSpecType) GetCustomAuthTypes() []*CustomAuthType {
	if m != nil {
		return m.CustomAuthTypes
	}
	return nil
}

// Get api discovery
//
// x-displayName: "Get Api Discovery"
// Get api_discovery reads a given object from storage backend for metadata.namespace.
type GetSpecType struct {
	CustomAuthTypes []*CustomAuthType `protobuf:"bytes,1,rep,name=custom_auth_types,json=customAuthTypes,proto3" json:"custom_auth_types,omitempty"`
}

func (m *GetSpecType) Reset()      { *m = GetSpecType{} }
func (*GetSpecType) ProtoMessage() {}
func (*GetSpecType) Descriptor() ([]byte, []int) {
	return fileDescriptor_3ccd12d9c64560a6, []int{4}
}
func (m *GetSpecType) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetSpecType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *GetSpecType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetSpecType.Merge(m, src)
}
func (m *GetSpecType) XXX_Size() int {
	return m.Size()
}
func (m *GetSpecType) XXX_DiscardUnknown() {
	xxx_messageInfo_GetSpecType.DiscardUnknown(m)
}

var xxx_messageInfo_GetSpecType proto.InternalMessageInfo

func (m *GetSpecType) GetCustomAuthTypes() []*CustomAuthType {
	if m != nil {
		return m.CustomAuthTypes
	}
	return nil
}

func init() {
	proto.RegisterEnum("ves.io.schema.api_sec.api_discovery.AuthParameterType", AuthParameterType_name, AuthParameterType_value)
	golang_proto.RegisterEnum("ves.io.schema.api_sec.api_discovery.AuthParameterType", AuthParameterType_name, AuthParameterType_value)
	proto.RegisterType((*CustomAuthType)(nil), "ves.io.schema.api_sec.api_discovery.CustomAuthType")
	golang_proto.RegisterType((*CustomAuthType)(nil), "ves.io.schema.api_sec.api_discovery.CustomAuthType")
	proto.RegisterType((*GlobalSpecType)(nil), "ves.io.schema.api_sec.api_discovery.GlobalSpecType")
	golang_proto.RegisterType((*GlobalSpecType)(nil), "ves.io.schema.api_sec.api_discovery.GlobalSpecType")
	proto.RegisterType((*CreateSpecType)(nil), "ves.io.schema.api_sec.api_discovery.CreateSpecType")
	golang_proto.RegisterType((*CreateSpecType)(nil), "ves.io.schema.api_sec.api_discovery.CreateSpecType")
	proto.RegisterType((*ReplaceSpecType)(nil), "ves.io.schema.api_sec.api_discovery.ReplaceSpecType")
	golang_proto.RegisterType((*ReplaceSpecType)(nil), "ves.io.schema.api_sec.api_discovery.ReplaceSpecType")
	proto.RegisterType((*GetSpecType)(nil), "ves.io.schema.api_sec.api_discovery.GetSpecType")
	golang_proto.RegisterType((*GetSpecType)(nil), "ves.io.schema.api_sec.api_discovery.GetSpecType")
}

func init() {
	proto.RegisterFile("ves.io/schema/api_sec/api_discovery/types.proto", fileDescriptor_3ccd12d9c64560a6)
}
func init() {
	golang_proto.RegisterFile("ves.io/schema/api_sec/api_discovery/types.proto", fileDescriptor_3ccd12d9c64560a6)
}

var fileDescriptor_3ccd12d9c64560a6 = []byte{
	// 536 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x94, 0x4f, 0x6b, 0x13, 0x41,
	0x18, 0xc6, 0xf7, 0x4d, 0xb0, 0xe0, 0x04, 0xf3, 0x4f, 0x0f, 0x31, 0xca, 0x18, 0xe2, 0x25, 0x08,
	0xd9, 0x85, 0x16, 0x3c, 0x28, 0x88, 0x49, 0x5d, 0xaa, 0x88, 0xb6, 0xae, 0xf5, 0x50, 0x2f, 0xcb,
	0x64, 0x33, 0xdd, 0x2c, 0xee, 0x76, 0x86, 0x99, 0xd9, 0x68, 0x2e, 0x52, 0x45, 0xef, 0xe2, 0xc5,
	0xaf, 0xe0, 0x47, 0x10, 0xf1, 0xd0, 0x8b, 0x20, 0x9e, 0x72, 0xcc, 0xb1, 0xd9, 0x5c, 0xf4, 0xd6,
	0x8f, 0x20, 0x99, 0xa4, 0x92, 0x6d, 0x40, 0x7a, 0x10, 0x7a, 0xda, 0x77, 0x78, 0xe6, 0x79, 0xde,
	0xdf, 0xbc, 0xec, 0x0c, 0xb2, 0xfa, 0x54, 0x9a, 0x01, 0xb3, 0xa4, 0xd7, 0xa3, 0x11, 0xb1, 0x08,
	0x0f, 0x5c, 0x49, 0x3d, 0xfd, 0xed, 0x06, 0xd2, 0x63, 0x7d, 0x2a, 0x06, 0x96, 0x1a, 0x70, 0x2a,
	0x4d, 0x2e, 0x98, 0x62, 0xe5, 0xeb, 0x33, 0x83, 0x39, 0x33, 0x98, 0x73, 0x83, 0x99, 0x32, 0x54,
	0x9b, 0x7e, 0xa0, 0x7a, 0x71, 0xc7, 0xf4, 0x58, 0x64, 0xf9, 0xcc, 0x67, 0x96, 0xf6, 0x76, 0xe2,
	0x5d, 0xbd, 0xd2, 0x0b, 0x5d, 0xcd, 0x32, 0xab, 0x57, 0xd2, 0x10, 0x8c, 0xab, 0x80, 0xed, 0xcd,
	0x1b, 0x56, 0x2f, 0xa7, 0xc5, 0x05, 0x96, 0xea, 0xd5, 0xb4, 0xd4, 0x27, 0x61, 0xd0, 0x25, 0x8a,
	0xce, 0xd5, 0xda, 0x09, 0x35, 0xa0, 0x2f, 0xdd, 0x74, 0xf4, 0xb5, 0xe5, 0x1d, 0x72, 0xb1, 0x41,
	0xfd, 0x0b, 0xa0, 0xfc, 0x7a, 0x2c, 0x15, 0x8b, 0x5a, 0xb1, 0xea, 0x6d, 0x0f, 0x38, 0x2d, 0x07,
	0x28, 0xcf, 0x89, 0x20, 0x11, 0x55, 0x54, 0xb8, 0xd3, 0xbd, 0x15, 0xa8, 0x41, 0x23, 0xbf, 0x7a,
	0xd3, 0x3c, 0xc5, 0x60, 0xcc, 0x69, 0xcc, 0xd6, 0xb1, 0x7d, 0x9a, 0xd7, 0x46, 0x5f, 0x7f, 0x1f,
	0x64, 0xcf, 0xbd, 0x85, 0x4c, 0x11, 0x9c, 0x0b, 0x7c, 0x51, 0x2a, 0xdf, 0x5e, 0x6c, 0xb5, 0x47,
	0x22, 0x5a, 0xc9, 0xd4, 0xa0, 0x71, 0xbe, 0x7d, 0x49, 0x5b, 0x44, 0xb6, 0xb1, 0x9f, 0x99, 0x57,
	0x23, 0x58, 0x34, 0x3f, 0x26, 0x11, 0xad, 0xbf, 0x46, 0xf9, 0x8d, 0x90, 0x75, 0x48, 0xf8, 0x94,
	0x53, 0x4f, 0xc7, 0x85, 0xa8, 0xe4, 0xe9, 0xb3, 0xb8, 0x24, 0x56, 0x3d, 0xcd, 0x2e, 0x2b, 0x50,
	0xcb, 0x36, 0x72, 0xab, 0x6b, 0xa7, 0x82, 0x4f, 0x4f, 0xa2, 0x9d, 0x4b, 0x0e, 0xbf, 0x67, 0x57,
	0x3e, 0x7e, 0xd3, 0xe8, 0x05, 0x2f, 0x25, 0xca, 0xfa, 0xbb, 0xe9, 0xe8, 0x04, 0x25, 0x8a, 0xfe,
	0x05, 0x70, 0xff, 0x2f, 0xc0, 0x52, 0xcf, 0x5b, 0xa5, 0x9f, 0x77, 0x4e, 0x1c, 0xba, 0xfe, 0x1e,
	0x50, 0xc1, 0xa1, 0x3c, 0x24, 0xde, 0xd9, 0x72, 0xbc, 0x01, 0x94, 0xdb, 0xa0, 0xea, 0x2c, 0x19,
	0x6e, 0xdc, 0x45, 0xa5, 0xa5, 0xff, 0xaf, 0x7c, 0x11, 0x15, 0x9e, 0x3c, 0xb3, 0x9d, 0x1d, 0x77,
	0xab, 0xe5, 0xb4, 0x1e, 0xd9, 0xdb, 0xb6, 0x53, 0x34, 0xca, 0x08, 0xad, 0xdc, 0xb7, 0x5b, 0xf7,
	0x6c, 0xa7, 0x08, 0xd3, 0x7a, 0x7d, 0x73, 0xf3, 0xe1, 0x03, 0xbb, 0x98, 0x69, 0x7f, 0x82, 0xe1,
	0x18, 0x1b, 0xa3, 0x31, 0x36, 0x8e, 0xc6, 0x18, 0xf6, 0x13, 0x0c, 0x9f, 0x13, 0x0c, 0x3f, 0x12,
	0x0c, 0xc3, 0x04, 0xc3, 0x28, 0xc1, 0x70, 0x98, 0x60, 0xf8, 0x95, 0x60, 0xe3, 0x28, 0xc1, 0xf0,
	0x61, 0x82, 0x8d, 0x83, 0x09, 0x86, 0xe1, 0x04, 0x1b, 0xa3, 0x09, 0x36, 0x9e, 0xef, 0xf8, 0x8c,
	0xbf, 0xf0, 0xcd, 0x3e, 0x0b, 0x15, 0x15, 0x82, 0x98, 0xb1, 0xb4, 0x74, 0xb1, 0xcb, 0x44, 0xd4,
	0xe4, 0x82, 0xf5, 0x83, 0x2e, 0x15, 0xcd, 0x63, 0xd9, 0xe2, 0x1d, 0x9f, 0x59, 0xf4, 0x95, 0x9a,
	0x5f, 0xcf, 0x7f, 0x3d, 0x51, 0x9d, 0x15, 0x7d, 0x61, 0xd7, 0xfe, 0x04, 0x00, 0x00, 0xff, 0xff,
	0x51, 0xa9, 0x42, 0xfb, 0xd0, 0x04, 0x00, 0x00,
}

func (x AuthParameterType) String() string {
	s, ok := AuthParameterType_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (this *CustomAuthType) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*CustomAuthType)
	if !ok {
		that2, ok := that.(CustomAuthType)
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
	if this.ParameterType != that1.ParameterType {
		return false
	}
	if this.ParameterName != that1.ParameterName {
		return false
	}
	return true
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
	if len(this.CustomAuthTypes) != len(that1.CustomAuthTypes) {
		return false
	}
	for i := range this.CustomAuthTypes {
		if !this.CustomAuthTypes[i].Equal(that1.CustomAuthTypes[i]) {
			return false
		}
	}
	return true
}
func (this *CreateSpecType) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*CreateSpecType)
	if !ok {
		that2, ok := that.(CreateSpecType)
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
	if len(this.CustomAuthTypes) != len(that1.CustomAuthTypes) {
		return false
	}
	for i := range this.CustomAuthTypes {
		if !this.CustomAuthTypes[i].Equal(that1.CustomAuthTypes[i]) {
			return false
		}
	}
	return true
}
func (this *ReplaceSpecType) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ReplaceSpecType)
	if !ok {
		that2, ok := that.(ReplaceSpecType)
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
	if len(this.CustomAuthTypes) != len(that1.CustomAuthTypes) {
		return false
	}
	for i := range this.CustomAuthTypes {
		if !this.CustomAuthTypes[i].Equal(that1.CustomAuthTypes[i]) {
			return false
		}
	}
	return true
}
func (this *GetSpecType) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*GetSpecType)
	if !ok {
		that2, ok := that.(GetSpecType)
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
	if len(this.CustomAuthTypes) != len(that1.CustomAuthTypes) {
		return false
	}
	for i := range this.CustomAuthTypes {
		if !this.CustomAuthTypes[i].Equal(that1.CustomAuthTypes[i]) {
			return false
		}
	}
	return true
}
func (this *CustomAuthType) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&api_discovery.CustomAuthType{")
	s = append(s, "ParameterType: "+fmt.Sprintf("%#v", this.ParameterType)+",\n")
	s = append(s, "ParameterName: "+fmt.Sprintf("%#v", this.ParameterName)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *GlobalSpecType) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&api_discovery.GlobalSpecType{")
	if this.CustomAuthTypes != nil {
		s = append(s, "CustomAuthTypes: "+fmt.Sprintf("%#v", this.CustomAuthTypes)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *CreateSpecType) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&api_discovery.CreateSpecType{")
	if this.CustomAuthTypes != nil {
		s = append(s, "CustomAuthTypes: "+fmt.Sprintf("%#v", this.CustomAuthTypes)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *ReplaceSpecType) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&api_discovery.ReplaceSpecType{")
	if this.CustomAuthTypes != nil {
		s = append(s, "CustomAuthTypes: "+fmt.Sprintf("%#v", this.CustomAuthTypes)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *GetSpecType) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&api_discovery.GetSpecType{")
	if this.CustomAuthTypes != nil {
		s = append(s, "CustomAuthTypes: "+fmt.Sprintf("%#v", this.CustomAuthTypes)+",\n")
	}
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
func (m *CustomAuthType) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CustomAuthType) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CustomAuthType) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ParameterName) > 0 {
		i -= len(m.ParameterName)
		copy(dAtA[i:], m.ParameterName)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.ParameterName)))
		i--
		dAtA[i] = 0x12
	}
	if m.ParameterType != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.ParameterType))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
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
	if len(m.CustomAuthTypes) > 0 {
		for iNdEx := len(m.CustomAuthTypes) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.CustomAuthTypes[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTypes(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *CreateSpecType) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CreateSpecType) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CreateSpecType) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.CustomAuthTypes) > 0 {
		for iNdEx := len(m.CustomAuthTypes) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.CustomAuthTypes[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTypes(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *ReplaceSpecType) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ReplaceSpecType) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ReplaceSpecType) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.CustomAuthTypes) > 0 {
		for iNdEx := len(m.CustomAuthTypes) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.CustomAuthTypes[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTypes(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *GetSpecType) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetSpecType) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetSpecType) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.CustomAuthTypes) > 0 {
		for iNdEx := len(m.CustomAuthTypes) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.CustomAuthTypes[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTypes(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
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
func (m *CustomAuthType) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ParameterType != 0 {
		n += 1 + sovTypes(uint64(m.ParameterType))
	}
	l = len(m.ParameterName)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}

func (m *GlobalSpecType) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.CustomAuthTypes) > 0 {
		for _, e := range m.CustomAuthTypes {
			l = e.Size()
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	return n
}

func (m *CreateSpecType) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.CustomAuthTypes) > 0 {
		for _, e := range m.CustomAuthTypes {
			l = e.Size()
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	return n
}

func (m *ReplaceSpecType) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.CustomAuthTypes) > 0 {
		for _, e := range m.CustomAuthTypes {
			l = e.Size()
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	return n
}

func (m *GetSpecType) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.CustomAuthTypes) > 0 {
		for _, e := range m.CustomAuthTypes {
			l = e.Size()
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	return n
}

func sovTypes(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTypes(x uint64) (n int) {
	return sovTypes(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *CustomAuthType) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&CustomAuthType{`,
		`ParameterType:` + fmt.Sprintf("%v", this.ParameterType) + `,`,
		`ParameterName:` + fmt.Sprintf("%v", this.ParameterName) + `,`,
		`}`,
	}, "")
	return s
}
func (this *GlobalSpecType) String() string {
	if this == nil {
		return "nil"
	}
	repeatedStringForCustomAuthTypes := "[]*CustomAuthType{"
	for _, f := range this.CustomAuthTypes {
		repeatedStringForCustomAuthTypes += strings.Replace(f.String(), "CustomAuthType", "CustomAuthType", 1) + ","
	}
	repeatedStringForCustomAuthTypes += "}"
	s := strings.Join([]string{`&GlobalSpecType{`,
		`CustomAuthTypes:` + repeatedStringForCustomAuthTypes + `,`,
		`}`,
	}, "")
	return s
}
func (this *CreateSpecType) String() string {
	if this == nil {
		return "nil"
	}
	repeatedStringForCustomAuthTypes := "[]*CustomAuthType{"
	for _, f := range this.CustomAuthTypes {
		repeatedStringForCustomAuthTypes += strings.Replace(f.String(), "CustomAuthType", "CustomAuthType", 1) + ","
	}
	repeatedStringForCustomAuthTypes += "}"
	s := strings.Join([]string{`&CreateSpecType{`,
		`CustomAuthTypes:` + repeatedStringForCustomAuthTypes + `,`,
		`}`,
	}, "")
	return s
}
func (this *ReplaceSpecType) String() string {
	if this == nil {
		return "nil"
	}
	repeatedStringForCustomAuthTypes := "[]*CustomAuthType{"
	for _, f := range this.CustomAuthTypes {
		repeatedStringForCustomAuthTypes += strings.Replace(f.String(), "CustomAuthType", "CustomAuthType", 1) + ","
	}
	repeatedStringForCustomAuthTypes += "}"
	s := strings.Join([]string{`&ReplaceSpecType{`,
		`CustomAuthTypes:` + repeatedStringForCustomAuthTypes + `,`,
		`}`,
	}, "")
	return s
}
func (this *GetSpecType) String() string {
	if this == nil {
		return "nil"
	}
	repeatedStringForCustomAuthTypes := "[]*CustomAuthType{"
	for _, f := range this.CustomAuthTypes {
		repeatedStringForCustomAuthTypes += strings.Replace(f.String(), "CustomAuthType", "CustomAuthType", 1) + ","
	}
	repeatedStringForCustomAuthTypes += "}"
	s := strings.Join([]string{`&GetSpecType{`,
		`CustomAuthTypes:` + repeatedStringForCustomAuthTypes + `,`,
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
func (m *CustomAuthType) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: CustomAuthType: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CustomAuthType: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ParameterType", wireType)
			}
			m.ParameterType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ParameterType |= AuthParameterType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ParameterName", wireType)
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
			m.ParameterName = string(dAtA[iNdEx:postIndex])
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
				return fmt.Errorf("proto: wrong wireType = %d for field CustomAuthTypes", wireType)
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
			m.CustomAuthTypes = append(m.CustomAuthTypes, &CustomAuthType{})
			if err := m.CustomAuthTypes[len(m.CustomAuthTypes)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
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
func (m *CreateSpecType) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: CreateSpecType: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CreateSpecType: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CustomAuthTypes", wireType)
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
			m.CustomAuthTypes = append(m.CustomAuthTypes, &CustomAuthType{})
			if err := m.CustomAuthTypes[len(m.CustomAuthTypes)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
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
func (m *ReplaceSpecType) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: ReplaceSpecType: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ReplaceSpecType: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CustomAuthTypes", wireType)
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
			m.CustomAuthTypes = append(m.CustomAuthTypes, &CustomAuthType{})
			if err := m.CustomAuthTypes[len(m.CustomAuthTypes)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
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
func (m *GetSpecType) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: GetSpecType: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetSpecType: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CustomAuthTypes", wireType)
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
			m.CustomAuthTypes = append(m.CustomAuthTypes, &CustomAuthType{})
			if err := m.CustomAuthTypes[len(m.CustomAuthTypes)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
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
