// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ves.io/schema/infraprotect_firewall_ruleset/types.proto

package infraprotect_firewall_ruleset

import (
	fmt "fmt"
	_ "github.com/gogo/googleapis/google/api"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	golang_proto "github.com/golang/protobuf/proto"
	schema "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema"
	_ "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/views"
	io "io"
	math "math"
	math_bits "math/bits"
	reflect "reflect"
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

// GlobalSpecType
//
// x-displayName: "DDos Transit Firewall Ruleset"
// DDos Transit Firewall Ruleset spec
type GlobalSpecType struct {
	// Version
	//
	// x-displayName: "Version"
	// x-required
	// x-example: "IPV4"
	// IP Version
	//
	// Types that are valid to be assigned to Version:
	//	*GlobalSpecType_VersionIpv4
	//	*GlobalSpecType_VersionIpv6
	Version isGlobalSpecType_Version `protobuf_oneof:"version"`
	// Firewall Rules
	//
	// x-displayName: "Firewall Rules"
	// x-required
	// Firewall Rules
	FirewallRule []*schema.ObjectRefType `protobuf:"bytes,4,rep,name=firewall_rule,json=firewallRule,proto3" json:"firewall_rule,omitempty"`
}

func (m *GlobalSpecType) Reset()      { *m = GlobalSpecType{} }
func (*GlobalSpecType) ProtoMessage() {}
func (*GlobalSpecType) Descriptor() ([]byte, []int) {
	return fileDescriptor_5b02053b1410338a, []int{0}
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

type isGlobalSpecType_Version interface {
	isGlobalSpecType_Version()
	Equal(interface{}) bool
	MarshalTo([]byte) (int, error)
	Size() int
}

type GlobalSpecType_VersionIpv4 struct {
	VersionIpv4 *schema.Empty `protobuf:"bytes,2,opt,name=version_ipv4,json=versionIpv4,proto3,oneof" json:"version_ipv4,omitempty"`
}
type GlobalSpecType_VersionIpv6 struct {
	VersionIpv6 *schema.Empty `protobuf:"bytes,3,opt,name=version_ipv6,json=versionIpv6,proto3,oneof" json:"version_ipv6,omitempty"`
}

func (*GlobalSpecType_VersionIpv4) isGlobalSpecType_Version() {}
func (*GlobalSpecType_VersionIpv6) isGlobalSpecType_Version() {}

func (m *GlobalSpecType) GetVersion() isGlobalSpecType_Version {
	if m != nil {
		return m.Version
	}
	return nil
}

func (m *GlobalSpecType) GetVersionIpv4() *schema.Empty {
	if x, ok := m.GetVersion().(*GlobalSpecType_VersionIpv4); ok {
		return x.VersionIpv4
	}
	return nil
}

func (m *GlobalSpecType) GetVersionIpv6() *schema.Empty {
	if x, ok := m.GetVersion().(*GlobalSpecType_VersionIpv6); ok {
		return x.VersionIpv6
	}
	return nil
}

func (m *GlobalSpecType) GetFirewallRule() []*schema.ObjectRefType {
	if m != nil {
		return m.FirewallRule
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*GlobalSpecType) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*GlobalSpecType_VersionIpv4)(nil),
		(*GlobalSpecType_VersionIpv6)(nil),
	}
}

// DDoS transit Firewall Ruleset
//
// x-displayName: "Replace DDoS transit Firewall Ruleset"
// Amends a DDoS transit Firewall Ruleset
type ReplaceSpecType struct {
	FirewallRule []*schema.ObjectRefType `protobuf:"bytes,4,rep,name=firewall_rule,json=firewallRule,proto3" json:"firewall_rule,omitempty"`
}

func (m *ReplaceSpecType) Reset()      { *m = ReplaceSpecType{} }
func (*ReplaceSpecType) ProtoMessage() {}
func (*ReplaceSpecType) Descriptor() ([]byte, []int) {
	return fileDescriptor_5b02053b1410338a, []int{1}
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

func (m *ReplaceSpecType) GetFirewallRule() []*schema.ObjectRefType {
	if m != nil {
		return m.FirewallRule
	}
	return nil
}

// Get DDoS transit Firewall Ruleset
//
// x-displayName: "Get Infraprotect Firewall Ruleset"
// Get DDoS transit Firewall Ruleset
type GetSpecType struct {
	// Types that are valid to be assigned to Version:
	//	*GetSpecType_VersionIpv4
	//	*GetSpecType_VersionIpv6
	Version      isGetSpecType_Version   `protobuf_oneof:"version"`
	FirewallRule []*schema.ObjectRefType `protobuf:"bytes,4,rep,name=firewall_rule,json=firewallRule,proto3" json:"firewall_rule,omitempty"`
}

func (m *GetSpecType) Reset()      { *m = GetSpecType{} }
func (*GetSpecType) ProtoMessage() {}
func (*GetSpecType) Descriptor() ([]byte, []int) {
	return fileDescriptor_5b02053b1410338a, []int{2}
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

type isGetSpecType_Version interface {
	isGetSpecType_Version()
	Equal(interface{}) bool
	MarshalTo([]byte) (int, error)
	Size() int
}

type GetSpecType_VersionIpv4 struct {
	VersionIpv4 *schema.Empty `protobuf:"bytes,2,opt,name=version_ipv4,json=versionIpv4,proto3,oneof" json:"version_ipv4,omitempty"`
}
type GetSpecType_VersionIpv6 struct {
	VersionIpv6 *schema.Empty `protobuf:"bytes,3,opt,name=version_ipv6,json=versionIpv6,proto3,oneof" json:"version_ipv6,omitempty"`
}

func (*GetSpecType_VersionIpv4) isGetSpecType_Version() {}
func (*GetSpecType_VersionIpv6) isGetSpecType_Version() {}

func (m *GetSpecType) GetVersion() isGetSpecType_Version {
	if m != nil {
		return m.Version
	}
	return nil
}

func (m *GetSpecType) GetVersionIpv4() *schema.Empty {
	if x, ok := m.GetVersion().(*GetSpecType_VersionIpv4); ok {
		return x.VersionIpv4
	}
	return nil
}

func (m *GetSpecType) GetVersionIpv6() *schema.Empty {
	if x, ok := m.GetVersion().(*GetSpecType_VersionIpv6); ok {
		return x.VersionIpv6
	}
	return nil
}

func (m *GetSpecType) GetFirewallRule() []*schema.ObjectRefType {
	if m != nil {
		return m.FirewallRule
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*GetSpecType) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*GetSpecType_VersionIpv4)(nil),
		(*GetSpecType_VersionIpv6)(nil),
	}
}

func init() {
	proto.RegisterType((*GlobalSpecType)(nil), "ves.io.schema.infraprotect_firewall_ruleset.GlobalSpecType")
	golang_proto.RegisterType((*GlobalSpecType)(nil), "ves.io.schema.infraprotect_firewall_ruleset.GlobalSpecType")
	proto.RegisterType((*ReplaceSpecType)(nil), "ves.io.schema.infraprotect_firewall_ruleset.ReplaceSpecType")
	golang_proto.RegisterType((*ReplaceSpecType)(nil), "ves.io.schema.infraprotect_firewall_ruleset.ReplaceSpecType")
	proto.RegisterType((*GetSpecType)(nil), "ves.io.schema.infraprotect_firewall_ruleset.GetSpecType")
	golang_proto.RegisterType((*GetSpecType)(nil), "ves.io.schema.infraprotect_firewall_ruleset.GetSpecType")
}

func init() {
	proto.RegisterFile("ves.io/schema/infraprotect_firewall_ruleset/types.proto", fileDescriptor_5b02053b1410338a)
}
func init() {
	golang_proto.RegisterFile("ves.io/schema/infraprotect_firewall_ruleset/types.proto", fileDescriptor_5b02053b1410338a)
}

var fileDescriptor_5b02053b1410338a = []byte{
	// 562 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x54, 0x41, 0x6b, 0x13, 0x5f,
	0x10, 0xdf, 0xd9, 0xe4, 0xdf, 0x96, 0x4d, 0xff, 0x5a, 0x63, 0x91, 0x18, 0xcb, 0xb3, 0xe4, 0x54,
	0xd4, 0xec, 0x42, 0x0c, 0x91, 0x5a, 0xa8, 0x18, 0x90, 0x5a, 0x11, 0x94, 0x55, 0x7b, 0xf0, 0x12,
	0xde, 0x6e, 0x27, 0xdb, 0xd5, 0xcd, 0xbe, 0xc7, 0xee, 0xcb, 0xd6, 0x1c, 0x14, 0xf1, 0x13, 0x48,
	0xbf, 0x83, 0x20, 0x7e, 0x04, 0xe3, 0x21, 0x17, 0x41, 0x72, 0xca, 0x31, 0xc7, 0x66, 0x73, 0xd1,
	0x5b, 0xbd, 0x7a, 0x92, 0x6e, 0x36, 0xb5, 0xbb, 0xa2, 0x14, 0x41, 0xbc, 0xcd, 0xcc, 0x6f, 0x7e,
	0xf3, 0x66, 0xe6, 0xc7, 0x3c, 0xe5, 0x5a, 0x80, 0xbe, 0x6a, 0x33, 0xcd, 0x37, 0x77, 0xb0, 0x45,
	0x35, 0xdb, 0x6d, 0x7a, 0x94, 0x7b, 0x4c, 0xa0, 0x29, 0x1a, 0x4d, 0xdb, 0xc3, 0x5d, 0xea, 0x38,
	0x0d, 0xaf, 0xed, 0xa0, 0x8f, 0x42, 0x13, 0x1d, 0x8e, 0xbe, 0x7a, 0x08, 0xb3, 0xfc, 0xe5, 0x09,
	0x51, 0x9d, 0x10, 0xd5, 0xdf, 0x12, 0x8b, 0x65, 0xcb, 0x16, 0x3b, 0x6d, 0x43, 0x35, 0x59, 0x4b,
	0xb3, 0x98, 0xc5, 0xb4, 0xa8, 0x86, 0xd1, 0x6e, 0x46, 0x5e, 0xe4, 0x44, 0xd6, 0xa4, 0x76, 0x71,
	0xc9, 0x62, 0xcc, 0x72, 0x50, 0xa3, 0xdc, 0xd6, 0xa8, 0xeb, 0x32, 0x41, 0x85, 0xcd, 0xdc, 0xf8,
	0xe5, 0xe2, 0xf9, 0x18, 0x3d, 0xaa, 0x41, 0xdd, 0x4e, 0x0c, 0x5d, 0x48, 0x4e, 0xc3, 0x78, 0x82,
	0x97, 0x04, 0x8f, 0x0d, 0x53, 0x5c, 0x4a, 0x42, 0x01, 0x75, 0xec, 0x6d, 0x2a, 0x30, 0x46, 0x97,
	0x53, 0xa8, 0x8d, 0xbb, 0x8d, 0x64, 0xe9, 0x8b, 0x3f, 0x67, 0xf8, 0xc7, 0x1f, 0x28, 0x85, 0xb2,
	0x72, 0x6a, 0xc3, 0x61, 0x06, 0x75, 0x1e, 0x70, 0x34, 0x1f, 0x76, 0x38, 0xe6, 0x57, 0x95, 0xf9,
	0x00, 0x3d, 0xdf, 0x66, 0x6e, 0xc3, 0xe6, 0x41, 0xb5, 0x20, 0x2f, 0xc3, 0x4a, 0xae, 0xb2, 0xa8,
	0x26, 0xf7, 0x7a, 0xab, 0xc5, 0x45, 0xe7, 0xb6, 0xa4, 0xe7, 0xe2, 0xdc, 0x4d, 0x1e, 0x54, 0x53,
	0xd4, 0x5a, 0x21, 0x73, 0x52, 0x6a, 0x2d, 0xff, 0x42, 0xf9, 0x3f, 0xa1, 0x4e, 0x21, 0xbb, 0x9c,
	0x59, 0xc9, 0x55, 0x96, 0x52, 0xdc, 0x7b, 0xc6, 0x13, 0x34, 0x85, 0x8e, 0xcd, 0xc3, 0x56, 0xeb,
	0x6b, 0xef, 0x9e, 0x17, 0x7f, 0x2d, 0xf0, 0xfb, 0x2f, 0xbd, 0xcc, 0x7f, 0x7b, 0x20, 0x17, 0x20,
	0xdc, 0xff, 0x98, 0xc9, 0xf6, 0xba, 0x10, 0x19, 0x33, 0x7b, 0x1f, 0x40, 0x5e, 0x00, 0x7d, 0x7e,
	0x9a, 0xac, 0xb7, 0x1d, 0xbc, 0x7e, 0xa3, 0xdf, 0x85, 0x35, 0x65, 0x55, 0x99, 0xdd, 0x9a, 0x34,
	0x75, 0x49, 0x55, 0xae, 0x28, 0x8b, 0xc9, 0x55, 0x54, 0xb2, 0x9b, 0xf7, 0xb7, 0xaa, 0xa9, 0x68,
	0x2d, 0x8a, 0xd6, 0xea, 0x67, 0x95, 0xd9, 0x38, 0x9a, 0x9f, 0xeb, 0x75, 0x41, 0x1e, 0x74, 0x01,
	0xee, 0x64, 0xe7, 0x60, 0x41, 0x2e, 0x59, 0xca, 0x69, 0x1d, 0xb9, 0x43, 0x4d, 0x3c, 0x5a, 0xf2,
	0xcd, 0x3f, 0x18, 0x37, 0xd5, 0xf1, 0x99, 0xfe, 0x7a, 0x4a, 0xba, 0xd2, 0x57, 0x50, 0x72, 0x1b,
	0x28, 0xfe, 0xb1, 0x94, 0x7f, 0x65, 0xb6, 0xfa, 0xb9, 0x1f, 0xfb, 0xcd, 0xbd, 0xfa, 0x06, 0x53,
	0xa7, 0xfe, 0x06, 0xfa, 0xeb, 0x19, 0xfd, 0xee, 0xa3, 0xc1, 0x88, 0x48, 0xc3, 0x11, 0x91, 0x0e,
	0x46, 0x04, 0x5e, 0x86, 0x04, 0xde, 0x86, 0x04, 0x3e, 0x85, 0x04, 0x06, 0x21, 0x81, 0x61, 0x48,
	0x60, 0x3f, 0x24, 0xf0, 0x39, 0x24, 0xd2, 0x41, 0x48, 0xe0, 0xf5, 0x98, 0x48, 0xbd, 0x31, 0x81,
	0xc1, 0x98, 0x48, 0xc3, 0x31, 0x91, 0x1e, 0x53, 0x8b, 0xf1, 0xa7, 0x96, 0x1a, 0x30, 0x47, 0xa0,
	0xe7, 0x51, 0xb5, 0xed, 0x6b, 0x91, 0xd1, 0x64, 0x5e, 0xab, 0xcc, 0x3d, 0x16, 0xd8, 0xdb, 0xe8,
	0x95, 0xa7, 0xb0, 0xc6, 0x0d, 0x8b, 0x69, 0xf8, 0x4c, 0xc4, 0x67, 0x75, 0x92, 0x3f, 0xca, 0x98,
	0x89, 0x0e, 0xee, 0xea, 0xf7, 0x00, 0x00, 0x00, 0xff, 0xff, 0x56, 0x1f, 0xcd, 0x87, 0xd9, 0x04,
	0x00, 0x00,
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
	if that1.Version == nil {
		if this.Version != nil {
			return false
		}
	} else if this.Version == nil {
		return false
	} else if !this.Version.Equal(that1.Version) {
		return false
	}
	if len(this.FirewallRule) != len(that1.FirewallRule) {
		return false
	}
	for i := range this.FirewallRule {
		if !this.FirewallRule[i].Equal(that1.FirewallRule[i]) {
			return false
		}
	}
	return true
}
func (this *GlobalSpecType_VersionIpv4) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*GlobalSpecType_VersionIpv4)
	if !ok {
		that2, ok := that.(GlobalSpecType_VersionIpv4)
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
	if !this.VersionIpv4.Equal(that1.VersionIpv4) {
		return false
	}
	return true
}
func (this *GlobalSpecType_VersionIpv6) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*GlobalSpecType_VersionIpv6)
	if !ok {
		that2, ok := that.(GlobalSpecType_VersionIpv6)
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
	if !this.VersionIpv6.Equal(that1.VersionIpv6) {
		return false
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
	if len(this.FirewallRule) != len(that1.FirewallRule) {
		return false
	}
	for i := range this.FirewallRule {
		if !this.FirewallRule[i].Equal(that1.FirewallRule[i]) {
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
	if that1.Version == nil {
		if this.Version != nil {
			return false
		}
	} else if this.Version == nil {
		return false
	} else if !this.Version.Equal(that1.Version) {
		return false
	}
	if len(this.FirewallRule) != len(that1.FirewallRule) {
		return false
	}
	for i := range this.FirewallRule {
		if !this.FirewallRule[i].Equal(that1.FirewallRule[i]) {
			return false
		}
	}
	return true
}
func (this *GetSpecType_VersionIpv4) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*GetSpecType_VersionIpv4)
	if !ok {
		that2, ok := that.(GetSpecType_VersionIpv4)
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
	if !this.VersionIpv4.Equal(that1.VersionIpv4) {
		return false
	}
	return true
}
func (this *GetSpecType_VersionIpv6) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*GetSpecType_VersionIpv6)
	if !ok {
		that2, ok := that.(GetSpecType_VersionIpv6)
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
	if !this.VersionIpv6.Equal(that1.VersionIpv6) {
		return false
	}
	return true
}
func (this *GlobalSpecType) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 7)
	s = append(s, "&infraprotect_firewall_ruleset.GlobalSpecType{")
	if this.Version != nil {
		s = append(s, "Version: "+fmt.Sprintf("%#v", this.Version)+",\n")
	}
	if this.FirewallRule != nil {
		s = append(s, "FirewallRule: "+fmt.Sprintf("%#v", this.FirewallRule)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *GlobalSpecType_VersionIpv4) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&infraprotect_firewall_ruleset.GlobalSpecType_VersionIpv4{` +
		`VersionIpv4:` + fmt.Sprintf("%#v", this.VersionIpv4) + `}`}, ", ")
	return s
}
func (this *GlobalSpecType_VersionIpv6) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&infraprotect_firewall_ruleset.GlobalSpecType_VersionIpv6{` +
		`VersionIpv6:` + fmt.Sprintf("%#v", this.VersionIpv6) + `}`}, ", ")
	return s
}
func (this *ReplaceSpecType) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&infraprotect_firewall_ruleset.ReplaceSpecType{")
	if this.FirewallRule != nil {
		s = append(s, "FirewallRule: "+fmt.Sprintf("%#v", this.FirewallRule)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *GetSpecType) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 7)
	s = append(s, "&infraprotect_firewall_ruleset.GetSpecType{")
	if this.Version != nil {
		s = append(s, "Version: "+fmt.Sprintf("%#v", this.Version)+",\n")
	}
	if this.FirewallRule != nil {
		s = append(s, "FirewallRule: "+fmt.Sprintf("%#v", this.FirewallRule)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *GetSpecType_VersionIpv4) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&infraprotect_firewall_ruleset.GetSpecType_VersionIpv4{` +
		`VersionIpv4:` + fmt.Sprintf("%#v", this.VersionIpv4) + `}`}, ", ")
	return s
}
func (this *GetSpecType_VersionIpv6) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&infraprotect_firewall_ruleset.GetSpecType_VersionIpv6{` +
		`VersionIpv6:` + fmt.Sprintf("%#v", this.VersionIpv6) + `}`}, ", ")
	return s
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
	if len(m.FirewallRule) > 0 {
		for iNdEx := len(m.FirewallRule) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.FirewallRule[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTypes(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if m.Version != nil {
		{
			size := m.Version.Size()
			i -= size
			if _, err := m.Version.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	return len(dAtA) - i, nil
}

func (m *GlobalSpecType_VersionIpv4) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GlobalSpecType_VersionIpv4) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.VersionIpv4 != nil {
		{
			size, err := m.VersionIpv4.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	return len(dAtA) - i, nil
}
func (m *GlobalSpecType_VersionIpv6) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GlobalSpecType_VersionIpv6) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.VersionIpv6 != nil {
		{
			size, err := m.VersionIpv6.MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.FirewallRule) > 0 {
		for iNdEx := len(m.FirewallRule) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.FirewallRule[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTypes(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
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
	if len(m.FirewallRule) > 0 {
		for iNdEx := len(m.FirewallRule) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.FirewallRule[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTypes(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if m.Version != nil {
		{
			size := m.Version.Size()
			i -= size
			if _, err := m.Version.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	return len(dAtA) - i, nil
}

func (m *GetSpecType_VersionIpv4) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetSpecType_VersionIpv4) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.VersionIpv4 != nil {
		{
			size, err := m.VersionIpv4.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	return len(dAtA) - i, nil
}
func (m *GetSpecType_VersionIpv6) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetSpecType_VersionIpv6) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.VersionIpv6 != nil {
		{
			size, err := m.VersionIpv6.MarshalToSizedBuffer(dAtA[:i])
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
	if m.Version != nil {
		n += m.Version.Size()
	}
	if len(m.FirewallRule) > 0 {
		for _, e := range m.FirewallRule {
			l = e.Size()
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	return n
}

func (m *GlobalSpecType_VersionIpv4) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.VersionIpv4 != nil {
		l = m.VersionIpv4.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}
func (m *GlobalSpecType_VersionIpv6) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.VersionIpv6 != nil {
		l = m.VersionIpv6.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}
func (m *ReplaceSpecType) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.FirewallRule) > 0 {
		for _, e := range m.FirewallRule {
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
	if m.Version != nil {
		n += m.Version.Size()
	}
	if len(m.FirewallRule) > 0 {
		for _, e := range m.FirewallRule {
			l = e.Size()
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	return n
}

func (m *GetSpecType_VersionIpv4) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.VersionIpv4 != nil {
		l = m.VersionIpv4.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}
func (m *GetSpecType_VersionIpv6) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.VersionIpv6 != nil {
		l = m.VersionIpv6.Size()
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
	repeatedStringForFirewallRule := "[]*ObjectRefType{"
	for _, f := range this.FirewallRule {
		repeatedStringForFirewallRule += strings.Replace(fmt.Sprintf("%v", f), "ObjectRefType", "schema.ObjectRefType", 1) + ","
	}
	repeatedStringForFirewallRule += "}"
	s := strings.Join([]string{`&GlobalSpecType{`,
		`Version:` + fmt.Sprintf("%v", this.Version) + `,`,
		`FirewallRule:` + repeatedStringForFirewallRule + `,`,
		`}`,
	}, "")
	return s
}
func (this *GlobalSpecType_VersionIpv4) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&GlobalSpecType_VersionIpv4{`,
		`VersionIpv4:` + strings.Replace(fmt.Sprintf("%v", this.VersionIpv4), "Empty", "schema.Empty", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *GlobalSpecType_VersionIpv6) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&GlobalSpecType_VersionIpv6{`,
		`VersionIpv6:` + strings.Replace(fmt.Sprintf("%v", this.VersionIpv6), "Empty", "schema.Empty", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *ReplaceSpecType) String() string {
	if this == nil {
		return "nil"
	}
	repeatedStringForFirewallRule := "[]*ObjectRefType{"
	for _, f := range this.FirewallRule {
		repeatedStringForFirewallRule += strings.Replace(fmt.Sprintf("%v", f), "ObjectRefType", "schema.ObjectRefType", 1) + ","
	}
	repeatedStringForFirewallRule += "}"
	s := strings.Join([]string{`&ReplaceSpecType{`,
		`FirewallRule:` + repeatedStringForFirewallRule + `,`,
		`}`,
	}, "")
	return s
}
func (this *GetSpecType) String() string {
	if this == nil {
		return "nil"
	}
	repeatedStringForFirewallRule := "[]*ObjectRefType{"
	for _, f := range this.FirewallRule {
		repeatedStringForFirewallRule += strings.Replace(fmt.Sprintf("%v", f), "ObjectRefType", "schema.ObjectRefType", 1) + ","
	}
	repeatedStringForFirewallRule += "}"
	s := strings.Join([]string{`&GetSpecType{`,
		`Version:` + fmt.Sprintf("%v", this.Version) + `,`,
		`FirewallRule:` + repeatedStringForFirewallRule + `,`,
		`}`,
	}, "")
	return s
}
func (this *GetSpecType_VersionIpv4) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&GetSpecType_VersionIpv4{`,
		`VersionIpv4:` + strings.Replace(fmt.Sprintf("%v", this.VersionIpv4), "Empty", "schema.Empty", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *GetSpecType_VersionIpv6) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&GetSpecType_VersionIpv6{`,
		`VersionIpv6:` + strings.Replace(fmt.Sprintf("%v", this.VersionIpv6), "Empty", "schema.Empty", 1) + `,`,
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
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VersionIpv4", wireType)
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
			m.Version = &GlobalSpecType_VersionIpv4{v}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VersionIpv6", wireType)
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
			m.Version = &GlobalSpecType_VersionIpv6{v}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FirewallRule", wireType)
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
			m.FirewallRule = append(m.FirewallRule, &schema.ObjectRefType{})
			if err := m.FirewallRule[len(m.FirewallRule)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FirewallRule", wireType)
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
			m.FirewallRule = append(m.FirewallRule, &schema.ObjectRefType{})
			if err := m.FirewallRule[len(m.FirewallRule)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VersionIpv4", wireType)
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
			m.Version = &GetSpecType_VersionIpv4{v}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VersionIpv6", wireType)
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
			m.Version = &GetSpecType_VersionIpv6{v}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FirewallRule", wireType)
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
			m.FirewallRule = append(m.FirewallRule, &schema.ObjectRefType{})
			if err := m.FirewallRule[len(m.FirewallRule)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
