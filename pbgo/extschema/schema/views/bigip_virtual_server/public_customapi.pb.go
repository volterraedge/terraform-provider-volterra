// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ves.io/schema/views/bigip_virtual_server/public_customapi.proto

// BIG-IP virtual server
//
// x-displayName: "BIG-IP virtual server"
// BIG-IP virtual server view repesents the internal virtual host corresponding to the virtual-servers discovered from BIG-IPs
// It exposes parameters to enable API discovery and other WAAP security features on the virtual server.

package bigip_virtual_server

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/googleapis/google/api"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	golang_proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	schema "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema"
	_ "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/dos_mitigation"
	_ "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/vesenv"
	_ "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/views"
	common_security "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/views/common_security"
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

// Get Security Config request
//
// x-displayName: "Get Security Config Request"
// Request of GET Security Config Spec API
type GetSecurityConfigReq struct {
	// Namespace
	//
	// x-displayName: "Namespace"
	// x-example: "shared"
	// x-required
	// Namespace of the BigIP Load Balancer for current request
	Namespace string `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	// Get loadbalancer
	//
	// x-required
	// x-displayName: "Load Balancer Choice"
	// Fetch Security Config of All Load Balancers or list of LBs
	//
	// Types that are valid to be assigned to LoadbalancerChoice:
	//	*GetSecurityConfigReq_AllBigipVirtualServers
	//	*GetSecurityConfigReq_BigipVirtualServersList
	LoadbalancerChoice isGetSecurityConfigReq_LoadbalancerChoice `protobuf_oneof:"loadbalancer_choice"`
}

func (m *GetSecurityConfigReq) Reset()      { *m = GetSecurityConfigReq{} }
func (*GetSecurityConfigReq) ProtoMessage() {}
func (*GetSecurityConfigReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_17a09d428162d9a5, []int{0}
}
func (m *GetSecurityConfigReq) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetSecurityConfigReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetSecurityConfigReq.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetSecurityConfigReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetSecurityConfigReq.Merge(m, src)
}
func (m *GetSecurityConfigReq) XXX_Size() int {
	return m.Size()
}
func (m *GetSecurityConfigReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetSecurityConfigReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetSecurityConfigReq proto.InternalMessageInfo

type isGetSecurityConfigReq_LoadbalancerChoice interface {
	isGetSecurityConfigReq_LoadbalancerChoice()
	Equal(interface{}) bool
	MarshalTo([]byte) (int, error)
	Size() int
}

type GetSecurityConfigReq_AllBigipVirtualServers struct {
	AllBigipVirtualServers *schema.Empty `protobuf:"bytes,3,opt,name=all_bigip_virtual_servers,json=allBigipVirtualServers,proto3,oneof" json:"all_bigip_virtual_servers,omitempty"`
}
type GetSecurityConfigReq_BigipVirtualServersList struct {
	BigipVirtualServersList *BigIPVirtualServerList `protobuf:"bytes,4,opt,name=bigip_virtual_servers_list,json=bigipVirtualServersList,proto3,oneof" json:"bigip_virtual_servers_list,omitempty"`
}

func (*GetSecurityConfigReq_AllBigipVirtualServers) isGetSecurityConfigReq_LoadbalancerChoice()  {}
func (*GetSecurityConfigReq_BigipVirtualServersList) isGetSecurityConfigReq_LoadbalancerChoice() {}

func (m *GetSecurityConfigReq) GetLoadbalancerChoice() isGetSecurityConfigReq_LoadbalancerChoice {
	if m != nil {
		return m.LoadbalancerChoice
	}
	return nil
}

func (m *GetSecurityConfigReq) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *GetSecurityConfigReq) GetAllBigipVirtualServers() *schema.Empty {
	if x, ok := m.GetLoadbalancerChoice().(*GetSecurityConfigReq_AllBigipVirtualServers); ok {
		return x.AllBigipVirtualServers
	}
	return nil
}

func (m *GetSecurityConfigReq) GetBigipVirtualServersList() *BigIPVirtualServerList {
	if x, ok := m.GetLoadbalancerChoice().(*GetSecurityConfigReq_BigipVirtualServersList); ok {
		return x.BigipVirtualServersList
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*GetSecurityConfigReq) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*GetSecurityConfigReq_AllBigipVirtualServers)(nil),
		(*GetSecurityConfigReq_BigipVirtualServersList)(nil),
	}
}

// List of BigIP Virtual Servers
//
// x-displayName: "List of BigIP Virtual Servers"
type BigIPVirtualServerList struct {
	// BigIP Virtual Servers
	//
	// x-displayName: "BigIP Virtual Server List"
	// x-example: "[blogging-app]"
	// x-required
	BigipVirtualServers []string `protobuf:"bytes,1,rep,name=bigip_virtual_servers,json=bigipVirtualServers,proto3" json:"bigip_virtual_servers,omitempty"`
}

func (m *BigIPVirtualServerList) Reset()      { *m = BigIPVirtualServerList{} }
func (*BigIPVirtualServerList) ProtoMessage() {}
func (*BigIPVirtualServerList) Descriptor() ([]byte, []int) {
	return fileDescriptor_17a09d428162d9a5, []int{1}
}
func (m *BigIPVirtualServerList) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BigIPVirtualServerList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BigIPVirtualServerList.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BigIPVirtualServerList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BigIPVirtualServerList.Merge(m, src)
}
func (m *BigIPVirtualServerList) XXX_Size() int {
	return m.Size()
}
func (m *BigIPVirtualServerList) XXX_DiscardUnknown() {
	xxx_messageInfo_BigIPVirtualServerList.DiscardUnknown(m)
}

var xxx_messageInfo_BigIPVirtualServerList proto.InternalMessageInfo

func (m *BigIPVirtualServerList) GetBigipVirtualServers() []string {
	if m != nil {
		return m.BigipVirtualServers
	}
	return nil
}

func init() {
	proto.RegisterType((*GetSecurityConfigReq)(nil), "ves.io.schema.views.bigip_virtual_server.GetSecurityConfigReq")
	golang_proto.RegisterType((*GetSecurityConfigReq)(nil), "ves.io.schema.views.bigip_virtual_server.GetSecurityConfigReq")
	proto.RegisterType((*BigIPVirtualServerList)(nil), "ves.io.schema.views.bigip_virtual_server.BigIPVirtualServerList")
	golang_proto.RegisterType((*BigIPVirtualServerList)(nil), "ves.io.schema.views.bigip_virtual_server.BigIPVirtualServerList")
}

func init() {
	proto.RegisterFile("ves.io/schema/views/bigip_virtual_server/public_customapi.proto", fileDescriptor_17a09d428162d9a5)
}
func init() {
	golang_proto.RegisterFile("ves.io/schema/views/bigip_virtual_server/public_customapi.proto", fileDescriptor_17a09d428162d9a5)
}

var fileDescriptor_17a09d428162d9a5 = []byte{
	// 710 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0xbf, 0x6f, 0xd3, 0x4e,
	0x14, 0xf7, 0x39, 0x6d, 0xbf, 0x8d, 0xbf, 0x4b, 0x70, 0x4b, 0x49, 0x43, 0xe5, 0x46, 0x99, 0x22,
	0x84, 0x7d, 0xa2, 0x88, 0x81, 0x4a, 0xfc, 0x4a, 0x85, 0x68, 0x11, 0x88, 0x92, 0x4a, 0x0c, 0x0c,
	0x58, 0x67, 0xe7, 0xe2, 0x1e, 0xd8, 0xbe, 0xe3, 0xee, 0x92, 0x36, 0x42, 0xa8, 0x55, 0x27, 0x46,
	0x54, 0x16, 0xf8, 0x0f, 0xd8, 0x58, 0x11, 0x5d, 0xba, 0xd1, 0x09, 0x45, 0xb0, 0x74, 0xa4, 0x0e,
	0x42, 0xb0, 0xf5, 0x4f, 0x40, 0x76, 0xd2, 0x92, 0xa4, 0x96, 0xe8, 0xf6, 0x9e, 0xdf, 0xfb, 0x7c,
	0xee, 0xf9, 0xf3, 0x79, 0x77, 0xda, 0x8d, 0x26, 0x16, 0x16, 0xa1, 0x50, 0xb8, 0xab, 0x38, 0x40,
	0xb0, 0x49, 0xf0, 0x9a, 0x80, 0x0e, 0xf1, 0x08, 0xb3, 0x9b, 0x84, 0xcb, 0x06, 0xf2, 0x6d, 0x81,
	0x79, 0x13, 0x73, 0xc8, 0x1a, 0x8e, 0x4f, 0x5c, 0xdb, 0x6d, 0x08, 0x49, 0x03, 0xc4, 0x88, 0xc5,
	0x38, 0x95, 0x54, 0x2f, 0x77, 0x09, 0xac, 0x2e, 0x81, 0x95, 0x10, 0x58, 0x69, 0x04, 0x05, 0xd3,
	0x23, 0x72, 0xb5, 0xe1, 0x58, 0x2e, 0x0d, 0xa0, 0x47, 0x3d, 0x0a, 0x13, 0x02, 0xa7, 0x51, 0x4f,
	0xb2, 0x24, 0x49, 0xa2, 0x2e, 0x71, 0x61, 0xc6, 0xa3, 0xd4, 0xf3, 0x31, 0x44, 0x8c, 0x40, 0x14,
	0x86, 0x54, 0x22, 0x49, 0x68, 0x28, 0x7a, 0xd5, 0xf2, 0xe0, 0xdc, 0x35, 0x2a, 0xec, 0x80, 0x48,
	0xe2, 0x25, 0x4d, 0x50, 0xb6, 0x18, 0x3e, 0xea, 0x3c, 0x3f, 0xd8, 0x49, 0x59, 0x3f, 0xcd, 0xf4,
	0x60, 0xb1, 0x1f, 0x37, 0x33, 0xa4, 0x0c, 0xf2, 0x49, 0x0d, 0x49, 0xdc, 0xab, 0x96, 0x86, 0xaa,
	0x58, 0xe0, 0xb0, 0x39, 0x44, 0x5e, 0x3c, 0xa9, 0xad, 0x3d, 0xd8, 0x31, 0x9f, 0xa6, 0xbe, 0x4b,
	0x83, 0x80, 0x86, 0xb6, 0xc0, 0x6e, 0x83, 0x13, 0xd9, 0x82, 0x5d, 0xc5, 0x6d, 0xc4, 0x88, 0xdd,
	0x3f, 0xdf, 0x6c, 0x1a, 0xb6, 0xaf, 0xa1, 0xf4, 0x41, 0xd5, 0x26, 0xef, 0x60, 0xb9, 0xd2, 0xe3,
	0x59, 0xa0, 0x61, 0x9d, 0x78, 0x55, 0xfc, 0x5c, 0x9f, 0xd1, 0xb2, 0x21, 0x0a, 0xb0, 0x60, 0xc8,
	0xc5, 0x79, 0x50, 0x04, 0xe5, 0x6c, 0xf5, 0xef, 0x07, 0xfd, 0xa1, 0x36, 0x8d, 0x7c, 0xdf, 0x4e,
	0xb3, 0x50, 0xe4, 0x33, 0x45, 0x50, 0xfe, 0x7f, 0x6e, 0xd2, 0x1a, 0x34, 0xfd, 0x76, 0xc0, 0x64,
	0x6b, 0x51, 0xa9, 0x4e, 0x21, 0xdf, 0xaf, 0xc4, 0xb8, 0x47, 0x5d, 0xd8, 0x4a, 0x17, 0xa5, 0x6f,
	0x68, 0x85, 0x54, 0x3a, 0xdb, 0x27, 0x42, 0xe6, 0x47, 0x12, 0xce, 0x9b, 0xd6, 0x69, 0x17, 0xc9,
	0xaa, 0x10, 0x6f, 0x69, 0x79, 0xe0, 0x88, 0x7b, 0x44, 0xc8, 0x45, 0xa5, 0x7a, 0xce, 0x39, 0x79,
	0x78, 0x5c, 0xaa, 0xcc, 0x6a, 0x13, 0x3e, 0x45, 0x35, 0x07, 0xf9, 0x28, 0x74, 0x31, 0xb7, 0xdd,
	0x55, 0x4a, 0x5c, 0xac, 0x8f, 0xef, 0xee, 0x80, 0x4c, 0x7b, 0x07, 0xa8, 0x77, 0x47, 0xc6, 0xd5,
	0x5c, 0xa6, 0xb4, 0xa1, 0x4d, 0xa5, 0x73, 0xeb, 0x58, 0x3b, 0x9b, 0x2e, 0x08, 0x28, 0x66, 0xca,
	0xd9, 0xca, 0xa5, 0x4f, 0xbf, 0x77, 0x33, 0xda, 0x36, 0xf8, 0xaf, 0x34, 0xca, 0x33, 0xf9, 0x4d,
	0x35, 0x4e, 0xb3, 0xdb, 0x60, 0xac, 0x34, 0xc2, 0xd5, 0x1c, 0x88, 0xb3, 0xd1, 0x6d, 0xa0, 0xe6,
	0x46, 0x8f, 0xa2, 0x71, 0x50, 0x9d, 0x48, 0x99, 0x75, 0xee, 0xa7, 0xaa, 0x65, 0x17, 0x12, 0xbb,
	0x6f, 0x2d, 0x2f, 0xe9, 0xef, 0x54, 0xed, 0xcc, 0x09, 0x03, 0xf5, 0xeb, 0xa7, 0x17, 0x2a, 0xcd,
	0xfd, 0xc2, 0xd5, 0x54, 0xfc, 0xd0, 0xd2, 0xa5, 0x40, 0x05, 0x2b, 0xbd, 0x02, 0x7b, 0x1f, 0x55,
	0x10, 0x7d, 0xce, 0x9b, 0xf5, 0x2b, 0xeb, 0xae, 0xb9, 0x86, 0x10, 0x33, 0x85, 0x44, 0x61, 0x0d,
	0xf1, 0x9a, 0x19, 0xd0, 0x90, 0x48, 0xca, 0x2f, 0xc6, 0x5b, 0x6f, 0x12, 0x6a, 0x32, 0x4e, 0xd7,
	0x5b, 0x26, 0xc7, 0xa8, 0xb6, 0xf5, 0xed, 0xc7, 0x1b, 0xf5, 0x7e, 0x69, 0xb1, 0xf7, 0x84, 0xc0,
	0xe3, 0x5d, 0x13, 0xf0, 0xc5, 0x71, 0xfc, 0x32, 0xf5, 0xd5, 0x11, 0xd0, 0xc3, 0xf2, 0x78, 0x2a,
	0xdb, 0x4d, 0x46, 0x99, 0x07, 0x17, 0x0a, 0xd7, 0x62, 0xeb, 0xbe, 0xee, 0x00, 0x78, 0x6a, 0x31,
	0x1e, 0x38, 0x4f, 0xb1, 0x2b, 0xb7, 0xbe, 0xe4, 0xd5, 0x1c, 0xa8, 0xbc, 0x05, 0xed, 0x03, 0x43,
	0xd9, 0x3f, 0x30, 0x94, 0xc3, 0x03, 0x03, 0x6c, 0x46, 0x06, 0x78, 0x1f, 0x19, 0x60, 0x2f, 0x32,
	0x40, 0x3b, 0x32, 0xc0, 0xf7, 0xc8, 0x00, 0xbf, 0x22, 0x43, 0x39, 0x8c, 0x0c, 0xf0, 0xba, 0x63,
	0x28, 0xbb, 0x1d, 0x03, 0xb4, 0x3b, 0x86, 0xb2, 0xdf, 0x31, 0x94, 0xc7, 0x4f, 0x3c, 0xca, 0x9e,
	0x79, 0x56, 0x93, 0xfa, 0x12, 0x73, 0x8e, 0xac, 0x86, 0x80, 0x49, 0x50, 0xa7, 0x3c, 0x88, 0x7f,
	0xba, 0x49, 0x6a, 0x98, 0x9b, 0x47, 0x65, 0xc8, 0x1c, 0x8f, 0x42, 0xbc, 0x2e, 0x7b, 0xd7, 0xf4,
	0x9f, 0xef, 0xac, 0x33, 0x96, 0xdc, 0xde, 0xcb, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x61, 0x7f,
	0x66, 0xd6, 0x9a, 0x05, 0x00, 0x00,
}

func (this *GetSecurityConfigReq) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*GetSecurityConfigReq)
	if !ok {
		that2, ok := that.(GetSecurityConfigReq)
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
	if this.Namespace != that1.Namespace {
		return false
	}
	if that1.LoadbalancerChoice == nil {
		if this.LoadbalancerChoice != nil {
			return false
		}
	} else if this.LoadbalancerChoice == nil {
		return false
	} else if !this.LoadbalancerChoice.Equal(that1.LoadbalancerChoice) {
		return false
	}
	return true
}
func (this *GetSecurityConfigReq_AllBigipVirtualServers) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*GetSecurityConfigReq_AllBigipVirtualServers)
	if !ok {
		that2, ok := that.(GetSecurityConfigReq_AllBigipVirtualServers)
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
	if !this.AllBigipVirtualServers.Equal(that1.AllBigipVirtualServers) {
		return false
	}
	return true
}
func (this *GetSecurityConfigReq_BigipVirtualServersList) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*GetSecurityConfigReq_BigipVirtualServersList)
	if !ok {
		that2, ok := that.(GetSecurityConfigReq_BigipVirtualServersList)
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
	if !this.BigipVirtualServersList.Equal(that1.BigipVirtualServersList) {
		return false
	}
	return true
}
func (this *BigIPVirtualServerList) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*BigIPVirtualServerList)
	if !ok {
		that2, ok := that.(BigIPVirtualServerList)
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
	if len(this.BigipVirtualServers) != len(that1.BigipVirtualServers) {
		return false
	}
	for i := range this.BigipVirtualServers {
		if this.BigipVirtualServers[i] != that1.BigipVirtualServers[i] {
			return false
		}
	}
	return true
}
func (this *GetSecurityConfigReq) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 7)
	s = append(s, "&bigip_virtual_server.GetSecurityConfigReq{")
	s = append(s, "Namespace: "+fmt.Sprintf("%#v", this.Namespace)+",\n")
	if this.LoadbalancerChoice != nil {
		s = append(s, "LoadbalancerChoice: "+fmt.Sprintf("%#v", this.LoadbalancerChoice)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *GetSecurityConfigReq_AllBigipVirtualServers) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&bigip_virtual_server.GetSecurityConfigReq_AllBigipVirtualServers{` +
		`AllBigipVirtualServers:` + fmt.Sprintf("%#v", this.AllBigipVirtualServers) + `}`}, ", ")
	return s
}
func (this *GetSecurityConfigReq_BigipVirtualServersList) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&bigip_virtual_server.GetSecurityConfigReq_BigipVirtualServersList{` +
		`BigipVirtualServersList:` + fmt.Sprintf("%#v", this.BigipVirtualServersList) + `}`}, ", ")
	return s
}
func (this *BigIPVirtualServerList) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&bigip_virtual_server.BigIPVirtualServerList{")
	s = append(s, "BigipVirtualServers: "+fmt.Sprintf("%#v", this.BigipVirtualServers)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringPublicCustomapi(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CustomAPIClient is the client API for CustomAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CustomAPIClient interface {
	// Get Security Config for BigIP Load Balancers
	//
	// x-displayName: "Get Security Config for BigIP Load Balancer"
	// Fetch the corresponding Security Config for the given BigIP load balancers
	GetSecurityConfig(ctx context.Context, in *GetSecurityConfigReq, opts ...grpc.CallOption) (*common_security.GetSecurityConfigRsp, error)
}

type customAPIClient struct {
	cc *grpc.ClientConn
}

func NewCustomAPIClient(cc *grpc.ClientConn) CustomAPIClient {
	return &customAPIClient{cc}
}

func (c *customAPIClient) GetSecurityConfig(ctx context.Context, in *GetSecurityConfigReq, opts ...grpc.CallOption) (*common_security.GetSecurityConfigRsp, error) {
	out := new(common_security.GetSecurityConfigRsp)
	err := c.cc.Invoke(ctx, "/ves.io.schema.views.bigip_virtual_server.CustomAPI/GetSecurityConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CustomAPIServer is the server API for CustomAPI service.
type CustomAPIServer interface {
	// Get Security Config for BigIP Load Balancers
	//
	// x-displayName: "Get Security Config for BigIP Load Balancer"
	// Fetch the corresponding Security Config for the given BigIP load balancers
	GetSecurityConfig(context.Context, *GetSecurityConfigReq) (*common_security.GetSecurityConfigRsp, error)
}

// UnimplementedCustomAPIServer can be embedded to have forward compatible implementations.
type UnimplementedCustomAPIServer struct {
}

func (*UnimplementedCustomAPIServer) GetSecurityConfig(ctx context.Context, req *GetSecurityConfigReq) (*common_security.GetSecurityConfigRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSecurityConfig not implemented")
}

func RegisterCustomAPIServer(s *grpc.Server, srv CustomAPIServer) {
	s.RegisterService(&_CustomAPI_serviceDesc, srv)
}

func _CustomAPI_GetSecurityConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSecurityConfigReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomAPIServer).GetSecurityConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ves.io.schema.views.bigip_virtual_server.CustomAPI/GetSecurityConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomAPIServer).GetSecurityConfig(ctx, req.(*GetSecurityConfigReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _CustomAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ves.io.schema.views.bigip_virtual_server.CustomAPI",
	HandlerType: (*CustomAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSecurityConfig",
			Handler:    _CustomAPI_GetSecurityConfig_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ves.io/schema/views/bigip_virtual_server/public_customapi.proto",
}

func (m *GetSecurityConfigReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetSecurityConfigReq) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetSecurityConfigReq) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.LoadbalancerChoice != nil {
		{
			size := m.LoadbalancerChoice.Size()
			i -= size
			if _, err := m.LoadbalancerChoice.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	if len(m.Namespace) > 0 {
		i -= len(m.Namespace)
		copy(dAtA[i:], m.Namespace)
		i = encodeVarintPublicCustomapi(dAtA, i, uint64(len(m.Namespace)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *GetSecurityConfigReq_AllBigipVirtualServers) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetSecurityConfigReq_AllBigipVirtualServers) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.AllBigipVirtualServers != nil {
		{
			size, err := m.AllBigipVirtualServers.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintPublicCustomapi(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	return len(dAtA) - i, nil
}
func (m *GetSecurityConfigReq_BigipVirtualServersList) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetSecurityConfigReq_BigipVirtualServersList) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.BigipVirtualServersList != nil {
		{
			size, err := m.BigipVirtualServersList.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintPublicCustomapi(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	return len(dAtA) - i, nil
}
func (m *BigIPVirtualServerList) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BigIPVirtualServerList) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BigIPVirtualServerList) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.BigipVirtualServers) > 0 {
		for iNdEx := len(m.BigipVirtualServers) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.BigipVirtualServers[iNdEx])
			copy(dAtA[i:], m.BigipVirtualServers[iNdEx])
			i = encodeVarintPublicCustomapi(dAtA, i, uint64(len(m.BigipVirtualServers[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintPublicCustomapi(dAtA []byte, offset int, v uint64) int {
	offset -= sovPublicCustomapi(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GetSecurityConfigReq) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Namespace)
	if l > 0 {
		n += 1 + l + sovPublicCustomapi(uint64(l))
	}
	if m.LoadbalancerChoice != nil {
		n += m.LoadbalancerChoice.Size()
	}
	return n
}

func (m *GetSecurityConfigReq_AllBigipVirtualServers) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.AllBigipVirtualServers != nil {
		l = m.AllBigipVirtualServers.Size()
		n += 1 + l + sovPublicCustomapi(uint64(l))
	}
	return n
}
func (m *GetSecurityConfigReq_BigipVirtualServersList) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.BigipVirtualServersList != nil {
		l = m.BigipVirtualServersList.Size()
		n += 1 + l + sovPublicCustomapi(uint64(l))
	}
	return n
}
func (m *BigIPVirtualServerList) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.BigipVirtualServers) > 0 {
		for _, s := range m.BigipVirtualServers {
			l = len(s)
			n += 1 + l + sovPublicCustomapi(uint64(l))
		}
	}
	return n
}

func sovPublicCustomapi(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPublicCustomapi(x uint64) (n int) {
	return sovPublicCustomapi(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *GetSecurityConfigReq) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&GetSecurityConfigReq{`,
		`Namespace:` + fmt.Sprintf("%v", this.Namespace) + `,`,
		`LoadbalancerChoice:` + fmt.Sprintf("%v", this.LoadbalancerChoice) + `,`,
		`}`,
	}, "")
	return s
}
func (this *GetSecurityConfigReq_AllBigipVirtualServers) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&GetSecurityConfigReq_AllBigipVirtualServers{`,
		`AllBigipVirtualServers:` + strings.Replace(fmt.Sprintf("%v", this.AllBigipVirtualServers), "Empty", "schema.Empty", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *GetSecurityConfigReq_BigipVirtualServersList) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&GetSecurityConfigReq_BigipVirtualServersList{`,
		`BigipVirtualServersList:` + strings.Replace(fmt.Sprintf("%v", this.BigipVirtualServersList), "BigIPVirtualServerList", "BigIPVirtualServerList", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *BigIPVirtualServerList) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&BigIPVirtualServerList{`,
		`BigipVirtualServers:` + fmt.Sprintf("%v", this.BigipVirtualServers) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringPublicCustomapi(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *GetSecurityConfigReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPublicCustomapi
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
			return fmt.Errorf("proto: GetSecurityConfigReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetSecurityConfigReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Namespace", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPublicCustomapi
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
				return ErrInvalidLengthPublicCustomapi
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPublicCustomapi
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Namespace = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AllBigipVirtualServers", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPublicCustomapi
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
				return ErrInvalidLengthPublicCustomapi
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPublicCustomapi
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &schema.Empty{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.LoadbalancerChoice = &GetSecurityConfigReq_AllBigipVirtualServers{v}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BigipVirtualServersList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPublicCustomapi
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
				return ErrInvalidLengthPublicCustomapi
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPublicCustomapi
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &BigIPVirtualServerList{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.LoadbalancerChoice = &GetSecurityConfigReq_BigipVirtualServersList{v}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPublicCustomapi(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPublicCustomapi
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthPublicCustomapi
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
func (m *BigIPVirtualServerList) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPublicCustomapi
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
			return fmt.Errorf("proto: BigIPVirtualServerList: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BigIPVirtualServerList: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BigipVirtualServers", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPublicCustomapi
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
				return ErrInvalidLengthPublicCustomapi
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPublicCustomapi
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BigipVirtualServers = append(m.BigipVirtualServers, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPublicCustomapi(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPublicCustomapi
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthPublicCustomapi
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
func skipPublicCustomapi(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPublicCustomapi
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
					return 0, ErrIntOverflowPublicCustomapi
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
					return 0, ErrIntOverflowPublicCustomapi
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
				return 0, ErrInvalidLengthPublicCustomapi
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPublicCustomapi
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPublicCustomapi
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPublicCustomapi        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPublicCustomapi          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPublicCustomapi = fmt.Errorf("proto: unexpected end of group")
)
