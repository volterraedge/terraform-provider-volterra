// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ves.io/schema/namespace/public_customapi_asterix.proto

// Namespace API
//
// x-displayName: "Namespace"
// APIs in this file are custom APIs on namespace object served by asterix

package namespace

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
	_ "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema"
	_ "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/vesenv"
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

// Api Endpoints stats request
//
// x-displayName: "Api Endpoints Stats Request"
// Request shape for GET Api Endpoints Stats
type ApiEndpointsStatsNSReq struct {
	// List Of Virtual Hosts Name
	//
	// x-displayName: "List Of Virtual Hosts Name"
	// x-example: "blogging-app, test-app"
	// List of Virtual Hosts for current request
	// If list empty or not sent, will sum of vhosts under requested namespace
	VhostsFilter []string `protobuf:"bytes,1,rep,name=vhosts_filter,json=vhostsFilter,proto3" json:"vhosts_filter,omitempty"`
	// Namespace
	//
	// x-displayName: "Namespace"
	// x-example: "shared"
	// Namespace of the App type for current request
	Namespace string `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
}

func (m *ApiEndpointsStatsNSReq) Reset()      { *m = ApiEndpointsStatsNSReq{} }
func (*ApiEndpointsStatsNSReq) ProtoMessage() {}
func (*ApiEndpointsStatsNSReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_0d03a36da3cd3d47, []int{0}
}
func (m *ApiEndpointsStatsNSReq) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ApiEndpointsStatsNSReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ApiEndpointsStatsNSReq.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ApiEndpointsStatsNSReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApiEndpointsStatsNSReq.Merge(m, src)
}
func (m *ApiEndpointsStatsNSReq) XXX_Size() int {
	return m.Size()
}
func (m *ApiEndpointsStatsNSReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ApiEndpointsStatsNSReq.DiscardUnknown(m)
}

var xxx_messageInfo_ApiEndpointsStatsNSReq proto.InternalMessageInfo

func (m *ApiEndpointsStatsNSReq) GetVhostsFilter() []string {
	if m != nil {
		return m.VhostsFilter
	}
	return nil
}

func (m *ApiEndpointsStatsNSReq) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

// Api Endpoints stats all namespaces request
//
// x-displayName: "Api Endpoints Stats All Namespaces Request"
// Request shape for GET Api Endpoints Stats All Namespaces
type ApiEndpointsStatsAllNSReq struct {
	// Namespace
	//
	// x-displayName: "Namespace"
	// x-example: "shared"
	// Namespace of the App type for current request
	Namespace string `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
}

func (m *ApiEndpointsStatsAllNSReq) Reset()      { *m = ApiEndpointsStatsAllNSReq{} }
func (*ApiEndpointsStatsAllNSReq) ProtoMessage() {}
func (*ApiEndpointsStatsAllNSReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_0d03a36da3cd3d47, []int{1}
}
func (m *ApiEndpointsStatsAllNSReq) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ApiEndpointsStatsAllNSReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ApiEndpointsStatsAllNSReq.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ApiEndpointsStatsAllNSReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApiEndpointsStatsAllNSReq.Merge(m, src)
}
func (m *ApiEndpointsStatsAllNSReq) XXX_Size() int {
	return m.Size()
}
func (m *ApiEndpointsStatsAllNSReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ApiEndpointsStatsAllNSReq.DiscardUnknown(m)
}

var xxx_messageInfo_ApiEndpointsStatsAllNSReq proto.InternalMessageInfo

func (m *ApiEndpointsStatsAllNSReq) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

// Api Endpoints stats Response
//
// x-displayName: "Api Endpoints Stats Response"
// Response shape for GET API endpoints Stats.
type ApiEndpointsStatsNSRsp struct {
	// number of endpoints
	//
	// x-displayName: "Total Endpoints"
	// total endpoints
	TotalEndpoints int32 `protobuf:"varint,1,opt,name=total_endpoints,json=totalEndpoints,proto3" json:"total_endpoints,omitempty"`
	// number of discovered endpoints
	//
	// x-displayName: "Discovered"
	// number of endpoints that categorized as discover
	Discovered int32 `protobuf:"varint,2,opt,name=discovered,proto3" json:"discovered,omitempty"`
	// number of inventory endpoints
	//
	// x-displayName: "Inventory"
	// number of endpoints that categorized as inventory
	Inventory int32 `protobuf:"varint,3,opt,name=inventory,proto3" json:"inventory,omitempty"`
	// number of shadow endpoints
	//
	// x-displayName: "Shadow"
	// number of endpoints that categorized as shadow
	Shadow int32 `protobuf:"varint,4,opt,name=shadow,proto3" json:"shadow,omitempty"`
	// number of pii endpoints
	//
	// x-displayName: "PII Detected"
	//number of endpoints that detected with pii
	PiiDetected int32 `protobuf:"varint,5,opt,name=pii_detected,json=piiDetected,proto3" json:"pii_detected,omitempty"`
}

func (m *ApiEndpointsStatsNSRsp) Reset()      { *m = ApiEndpointsStatsNSRsp{} }
func (*ApiEndpointsStatsNSRsp) ProtoMessage() {}
func (*ApiEndpointsStatsNSRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_0d03a36da3cd3d47, []int{2}
}
func (m *ApiEndpointsStatsNSRsp) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ApiEndpointsStatsNSRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ApiEndpointsStatsNSRsp.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ApiEndpointsStatsNSRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApiEndpointsStatsNSRsp.Merge(m, src)
}
func (m *ApiEndpointsStatsNSRsp) XXX_Size() int {
	return m.Size()
}
func (m *ApiEndpointsStatsNSRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_ApiEndpointsStatsNSRsp.DiscardUnknown(m)
}

var xxx_messageInfo_ApiEndpointsStatsNSRsp proto.InternalMessageInfo

func (m *ApiEndpointsStatsNSRsp) GetTotalEndpoints() int32 {
	if m != nil {
		return m.TotalEndpoints
	}
	return 0
}

func (m *ApiEndpointsStatsNSRsp) GetDiscovered() int32 {
	if m != nil {
		return m.Discovered
	}
	return 0
}

func (m *ApiEndpointsStatsNSRsp) GetInventory() int32 {
	if m != nil {
		return m.Inventory
	}
	return 0
}

func (m *ApiEndpointsStatsNSRsp) GetShadow() int32 {
	if m != nil {
		return m.Shadow
	}
	return 0
}

func (m *ApiEndpointsStatsNSRsp) GetPiiDetected() int32 {
	if m != nil {
		return m.PiiDetected
	}
	return 0
}

func init() {
	proto.RegisterType((*ApiEndpointsStatsNSReq)(nil), "ves.io.schema.namespace.ApiEndpointsStatsNSReq")
	golang_proto.RegisterType((*ApiEndpointsStatsNSReq)(nil), "ves.io.schema.namespace.ApiEndpointsStatsNSReq")
	proto.RegisterType((*ApiEndpointsStatsAllNSReq)(nil), "ves.io.schema.namespace.ApiEndpointsStatsAllNSReq")
	golang_proto.RegisterType((*ApiEndpointsStatsAllNSReq)(nil), "ves.io.schema.namespace.ApiEndpointsStatsAllNSReq")
	proto.RegisterType((*ApiEndpointsStatsNSRsp)(nil), "ves.io.schema.namespace.ApiEndpointsStatsNSRsp")
	golang_proto.RegisterType((*ApiEndpointsStatsNSRsp)(nil), "ves.io.schema.namespace.ApiEndpointsStatsNSRsp")
}

func init() {
	proto.RegisterFile("ves.io/schema/namespace/public_customapi_asterix.proto", fileDescriptor_0d03a36da3cd3d47)
}
func init() {
	golang_proto.RegisterFile("ves.io/schema/namespace/public_customapi_asterix.proto", fileDescriptor_0d03a36da3cd3d47)
}

var fileDescriptor_0d03a36da3cd3d47 = []byte{
	// 749 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x94, 0xcf, 0x4f, 0xe3, 0x46,
	0x14, 0xc7, 0x33, 0xa1, 0x41, 0xc2, 0x85, 0x56, 0xb2, 0x10, 0x75, 0x53, 0x34, 0x0d, 0xe6, 0x50,
	0x8a, 0xb0, 0x5d, 0x85, 0xfe, 0x10, 0xbd, 0xd1, 0x9f, 0xaa, 0x54, 0x68, 0x15, 0xa4, 0x1e, 0xda,
	0x83, 0x35, 0xb1, 0x5f, 0x9c, 0x69, 0x6d, 0xcf, 0xe0, 0x99, 0x98, 0xa0, 0xaa, 0x52, 0xc5, 0xa5,
	0xd7, 0x4a, 0xfd, 0x27, 0xfa, 0x27, 0xac, 0x96, 0x0b, 0xb7, 0xe5, 0xb4, 0x42, 0xbb, 0x87, 0xe5,
	0xb8, 0x38, 0x7b, 0xd8, 0xbd, 0xa1, 0xfd, 0x07, 0x76, 0x95, 0x89, 0xf3, 0x93, 0x44, 0x0b, 0xb7,
	0x99, 0xf7, 0x7d, 0xdf, 0x37, 0xef, 0x93, 0x3c, 0x3f, 0xed, 0xf3, 0x14, 0x84, 0x4d, 0x99, 0x23,
	0xbc, 0x26, 0x44, 0xc4, 0x89, 0x49, 0x04, 0x82, 0x13, 0x0f, 0x1c, 0xde, 0xaa, 0x87, 0xd4, 0x73,
	0xbd, 0x96, 0x90, 0x2c, 0x22, 0x9c, 0xba, 0x44, 0x48, 0x48, 0x68, 0xdb, 0xe6, 0x09, 0x93, 0x4c,
	0x7f, 0xaf, 0xe7, 0xb3, 0x7b, 0x3e, 0x7b, 0xe0, 0x2b, 0x5b, 0x01, 0x95, 0xcd, 0x56, 0xdd, 0xf6,
	0x58, 0xe4, 0x04, 0x2c, 0x60, 0x8e, 0xca, 0xaf, 0xb7, 0x1a, 0xea, 0xa6, 0x2e, 0xea, 0xd4, 0xab,
	0x53, 0x5e, 0x0d, 0x18, 0x0b, 0x42, 0x70, 0x08, 0xa7, 0x0e, 0x89, 0x63, 0x26, 0x89, 0xa4, 0x2c,
	0x16, 0xb9, 0xba, 0x3e, 0xab, 0x3b, 0x79, 0xcc, 0xa1, 0x9f, 0xf4, 0xc1, 0x78, 0x12, 0xe3, 0xa3,
	0x15, 0x56, 0xc7, 0xc5, 0x94, 0x84, 0xd4, 0x27, 0x12, 0x72, 0xd5, 0x9c, 0x50, 0x41, 0x40, 0x9c,
	0x4e, 0x54, 0xa8, 0x4c, 0xe4, 0x50, 0x38, 0x72, 0xc7, 0x33, 0x3e, 0xbc, 0x99, 0x21, 0x46, 0x3b,
	0x34, 0x7f, 0xd3, 0x56, 0x76, 0x39, 0xfd, 0x36, 0xf6, 0x39, 0xa3, 0xb1, 0x14, 0x07, 0x92, 0x48,
	0xb1, 0x7f, 0x50, 0x83, 0x43, 0x7d, 0x5d, 0x5b, 0x4a, 0x9b, 0x4c, 0x48, 0xe1, 0x36, 0x68, 0x28,
	0x21, 0x31, 0x50, 0x65, 0x6e, 0x63, 0xa1, 0xb6, 0xd8, 0x0b, 0x7e, 0xa7, 0x62, 0xfa, 0xaa, 0xb6,
	0x30, 0x20, 0x37, 0x8a, 0x15, 0xb4, 0xb1, 0x50, 0x1b, 0x06, 0xcc, 0x1d, 0xed, 0xfd, 0x1b, 0xc5,
	0x77, 0xc3, 0xb0, 0x57, 0x7f, 0xcc, 0x8a, 0x26, 0xad, 0xf7, 0xd1, 0xf4, 0xc6, 0x04, 0xd7, 0x3f,
	0xd2, 0xde, 0x95, 0x4c, 0x92, 0xd0, 0x85, 0xbe, 0xa8, 0xec, 0xa5, 0xda, 0x3b, 0x2a, 0x3c, 0xb0,
	0xe8, 0x58, 0xd3, 0x7c, 0x2a, 0x3c, 0x96, 0x42, 0x02, 0xbe, 0xea, 0xae, 0x54, 0x1b, 0x89, 0x74,
	0x3b, 0xa0, 0x71, 0x0a, 0xb1, 0x64, 0xc9, 0xb1, 0x31, 0xa7, 0xe4, 0x61, 0x40, 0x5f, 0xd1, 0xe6,
	0x45, 0x93, 0xf8, 0xec, 0xc8, 0x78, 0x4b, 0x49, 0xf9, 0x4d, 0x5f, 0xd3, 0x16, 0x39, 0xa5, 0xae,
	0x0f, 0x12, 0x3c, 0x09, 0xbe, 0x51, 0x52, 0xea, 0xdb, 0x9c, 0xd2, 0x6f, 0xf2, 0x50, 0xf5, 0x55,
	0x49, 0x5b, 0xde, 0xef, 0xa3, 0xec, 0xfd, 0xf8, 0xb5, 0x1a, 0xd4, 0xdd, 0x9f, 0x7f, 0xd0, 0x9f,
	0x20, 0x6d, 0xe9, 0xa0, 0x15, 0x04, 0x20, 0xe4, 0x2f, 0x24, 0x6c, 0x81, 0xd0, 0x3f, 0xb6, 0x67,
	0x4c, 0xab, 0x3d, 0x96, 0x57, 0x83, 0xc3, 0xf2, 0xe6, 0x6d, 0x53, 0x05, 0x37, 0xd9, 0xf9, 0xbd,
	0x22, 0xca, 0x1e, 0x18, 0xdd, 0xb9, 0xb1, 0x28, 0xb3, 0x02, 0x88, 0x21, 0x21, 0xa1, 0x95, 0x00,
	0xf1, 0xb7, 0x2a, 0x8d, 0xcf, 0xda, 0x9e, 0x75, 0x44, 0x08, 0xb7, 0x88, 0x9f, 0x92, 0xd8, 0x03,
	0xdf, 0x8a, 0x58, 0x4c, 0x25, 0x4b, 0x4e, 0x1e, 0x3f, 0xfb, 0xaf, 0x58, 0x35, 0xad, 0xfc, 0x1b,
	0x1b, 0x8e, 0xb5, 0x70, 0xfe, 0x1c, 0x9c, 0xff, 0x72, 0x44, 0xef, 0x41, 0x2b, 0x55, 0x2f, 0x7e,
	0x89, 0x36, 0xf5, 0x97, 0x48, 0x5b, 0xfe, 0x1e, 0xe4, 0x8d, 0xbf, 0x4c, 0x77, 0x66, 0x76, 0x3d,
	0x7d, 0xee, 0xca, 0x77, 0x33, 0x08, 0x6e, 0xb6, 0x73, 0xd6, 0xed, 0x21, 0x56, 0x9d, 0x08, 0xea,
	0xf5, 0x99, 0x46, 0x79, 0x85, 0x24, 0xb1, 0x4f, 0x92, 0x71, 0xde, 0x2f, 0xcc, 0xea, 0x1b, 0x78,
	0xbb, 0x4b, 0x66, 0x30, 0x6d, 0x8e, 0xe8, 0xbe, 0xdd, 0x85, 0x7e, 0x81, 0xb4, 0xb5, 0x69, 0xd0,
	0xdd, 0x19, 0x1f, 0x14, 0xd2, 0xab, 0xb7, 0x07, 0xea, 0x7f, 0x1c, 0x77, 0xff, 0x11, 0xdc, 0xfc,
	0x47, 0x30, 0xa6, 0xb0, 0x12, 0x3f, 0xa2, 0xb1, 0x22, 0xdd, 0x31, 0x3f, 0x9d, 0x42, 0x2a, 0x8e,
	0x85, 0x84, 0x68, 0x02, 0x92, 0x84, 0xa1, 0x1b, 0x0b, 0xb7, 0xcf, 0x5a, 0xde, 0x3a, 0x3b, 0x45,
	0x73, 0x8f, 0x4e, 0x11, 0x9e, 0xd5, 0xd8, 0x4f, 0xf5, 0xdf, 0xc1, 0x93, 0x27, 0x0f, 0x8d, 0xe2,
	0x27, 0xe8, 0xab, 0x7f, 0xd0, 0xc5, 0x15, 0x2e, 0x5c, 0x5e, 0xe1, 0xc2, 0xf5, 0x15, 0x46, 0x7f,
	0x67, 0x18, 0xfd, 0x9f, 0x61, 0x74, 0x9e, 0x61, 0x74, 0x91, 0x61, 0xf4, 0x34, 0xc3, 0xe8, 0x79,
	0x86, 0x0b, 0xd7, 0x19, 0x46, 0xff, 0x76, 0x70, 0xe1, 0xac, 0x83, 0xd1, 0x45, 0x07, 0x17, 0x2e,
	0x3b, 0xb8, 0xf0, 0xeb, 0x5e, 0xc0, 0xf8, 0x1f, 0x81, 0x9d, 0xb2, 0xee, 0x82, 0x49, 0x88, 0xdd,
	0x12, 0x8e, 0x3a, 0x34, 0x58, 0x12, 0x59, 0x3c, 0x61, 0x29, 0xf5, 0x21, 0xb1, 0xfa, 0xb2, 0xc3,
	0xeb, 0x01, 0x73, 0xa0, 0x2d, 0xf3, 0x0d, 0x37, 0xb9, 0x8e, 0xeb, 0xf3, 0x6a, 0xcf, 0x6d, 0xbf,
	0x0e, 0x00, 0x00, 0xff, 0xff, 0x7e, 0xb2, 0x25, 0xcc, 0x4e, 0x06, 0x00, 0x00,
}

func (this *ApiEndpointsStatsNSReq) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ApiEndpointsStatsNSReq)
	if !ok {
		that2, ok := that.(ApiEndpointsStatsNSReq)
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
	if len(this.VhostsFilter) != len(that1.VhostsFilter) {
		return false
	}
	for i := range this.VhostsFilter {
		if this.VhostsFilter[i] != that1.VhostsFilter[i] {
			return false
		}
	}
	if this.Namespace != that1.Namespace {
		return false
	}
	return true
}
func (this *ApiEndpointsStatsAllNSReq) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ApiEndpointsStatsAllNSReq)
	if !ok {
		that2, ok := that.(ApiEndpointsStatsAllNSReq)
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
	return true
}
func (this *ApiEndpointsStatsNSRsp) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ApiEndpointsStatsNSRsp)
	if !ok {
		that2, ok := that.(ApiEndpointsStatsNSRsp)
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
	if this.TotalEndpoints != that1.TotalEndpoints {
		return false
	}
	if this.Discovered != that1.Discovered {
		return false
	}
	if this.Inventory != that1.Inventory {
		return false
	}
	if this.Shadow != that1.Shadow {
		return false
	}
	if this.PiiDetected != that1.PiiDetected {
		return false
	}
	return true
}
func (this *ApiEndpointsStatsNSReq) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&namespace.ApiEndpointsStatsNSReq{")
	s = append(s, "VhostsFilter: "+fmt.Sprintf("%#v", this.VhostsFilter)+",\n")
	s = append(s, "Namespace: "+fmt.Sprintf("%#v", this.Namespace)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *ApiEndpointsStatsAllNSReq) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&namespace.ApiEndpointsStatsAllNSReq{")
	s = append(s, "Namespace: "+fmt.Sprintf("%#v", this.Namespace)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *ApiEndpointsStatsNSRsp) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 9)
	s = append(s, "&namespace.ApiEndpointsStatsNSRsp{")
	s = append(s, "TotalEndpoints: "+fmt.Sprintf("%#v", this.TotalEndpoints)+",\n")
	s = append(s, "Discovered: "+fmt.Sprintf("%#v", this.Discovered)+",\n")
	s = append(s, "Inventory: "+fmt.Sprintf("%#v", this.Inventory)+",\n")
	s = append(s, "Shadow: "+fmt.Sprintf("%#v", this.Shadow)+",\n")
	s = append(s, "PiiDetected: "+fmt.Sprintf("%#v", this.PiiDetected)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringPublicCustomapiAsterix(v interface{}, typ string) string {
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

// NamespaceMLCustomAPIClient is the client API for NamespaceMLCustomAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NamespaceMLCustomAPIClient interface {
	// SuggestValues
	//
	// x-displayName: "Suggest Values"
	// Returns suggested values for the specified field in the given Create/Replace/Custom request
	SuggestValues(ctx context.Context, in *SuggestValuesReq, opts ...grpc.CallOption) (*SuggestValuesResp, error)
	// Get Api Endpoints Stats for Namespace
	//
	// x-displayName: "Get Api Endpoints Stats for Namespace"
	// Get api endpoints stats for the given Namespace
	GetApiEndpointsStats(ctx context.Context, in *ApiEndpointsStatsNSReq, opts ...grpc.CallOption) (*ApiEndpointsStatsNSRsp, error)
	// Get Api Endpoints Stats for All Namespaces
	//
	// x-displayName: "Get Api Endpoints Stats for All Namespaces"
	// Get api endpoints stats for all Namespaces. This API is specific to system namespace
	GetApiEndpointsStatsAllNamespaces(ctx context.Context, in *ApiEndpointsStatsAllNSReq, opts ...grpc.CallOption) (*ApiEndpointsStatsNSRsp, error)
}

type namespaceMLCustomAPIClient struct {
	cc *grpc.ClientConn
}

func NewNamespaceMLCustomAPIClient(cc *grpc.ClientConn) NamespaceMLCustomAPIClient {
	return &namespaceMLCustomAPIClient{cc}
}

func (c *namespaceMLCustomAPIClient) SuggestValues(ctx context.Context, in *SuggestValuesReq, opts ...grpc.CallOption) (*SuggestValuesResp, error) {
	out := new(SuggestValuesResp)
	err := c.cc.Invoke(ctx, "/ves.io.schema.namespace.NamespaceMLCustomAPI/SuggestValues", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *namespaceMLCustomAPIClient) GetApiEndpointsStats(ctx context.Context, in *ApiEndpointsStatsNSReq, opts ...grpc.CallOption) (*ApiEndpointsStatsNSRsp, error) {
	out := new(ApiEndpointsStatsNSRsp)
	err := c.cc.Invoke(ctx, "/ves.io.schema.namespace.NamespaceMLCustomAPI/GetApiEndpointsStats", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *namespaceMLCustomAPIClient) GetApiEndpointsStatsAllNamespaces(ctx context.Context, in *ApiEndpointsStatsAllNSReq, opts ...grpc.CallOption) (*ApiEndpointsStatsNSRsp, error) {
	out := new(ApiEndpointsStatsNSRsp)
	err := c.cc.Invoke(ctx, "/ves.io.schema.namespace.NamespaceMLCustomAPI/GetApiEndpointsStatsAllNamespaces", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NamespaceMLCustomAPIServer is the server API for NamespaceMLCustomAPI service.
type NamespaceMLCustomAPIServer interface {
	// SuggestValues
	//
	// x-displayName: "Suggest Values"
	// Returns suggested values for the specified field in the given Create/Replace/Custom request
	SuggestValues(context.Context, *SuggestValuesReq) (*SuggestValuesResp, error)
	// Get Api Endpoints Stats for Namespace
	//
	// x-displayName: "Get Api Endpoints Stats for Namespace"
	// Get api endpoints stats for the given Namespace
	GetApiEndpointsStats(context.Context, *ApiEndpointsStatsNSReq) (*ApiEndpointsStatsNSRsp, error)
	// Get Api Endpoints Stats for All Namespaces
	//
	// x-displayName: "Get Api Endpoints Stats for All Namespaces"
	// Get api endpoints stats for all Namespaces. This API is specific to system namespace
	GetApiEndpointsStatsAllNamespaces(context.Context, *ApiEndpointsStatsAllNSReq) (*ApiEndpointsStatsNSRsp, error)
}

// UnimplementedNamespaceMLCustomAPIServer can be embedded to have forward compatible implementations.
type UnimplementedNamespaceMLCustomAPIServer struct {
}

func (*UnimplementedNamespaceMLCustomAPIServer) SuggestValues(ctx context.Context, req *SuggestValuesReq) (*SuggestValuesResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SuggestValues not implemented")
}
func (*UnimplementedNamespaceMLCustomAPIServer) GetApiEndpointsStats(ctx context.Context, req *ApiEndpointsStatsNSReq) (*ApiEndpointsStatsNSRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetApiEndpointsStats not implemented")
}
func (*UnimplementedNamespaceMLCustomAPIServer) GetApiEndpointsStatsAllNamespaces(ctx context.Context, req *ApiEndpointsStatsAllNSReq) (*ApiEndpointsStatsNSRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetApiEndpointsStatsAllNamespaces not implemented")
}

func RegisterNamespaceMLCustomAPIServer(s *grpc.Server, srv NamespaceMLCustomAPIServer) {
	s.RegisterService(&_NamespaceMLCustomAPI_serviceDesc, srv)
}

func _NamespaceMLCustomAPI_SuggestValues_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SuggestValuesReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NamespaceMLCustomAPIServer).SuggestValues(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ves.io.schema.namespace.NamespaceMLCustomAPI/SuggestValues",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NamespaceMLCustomAPIServer).SuggestValues(ctx, req.(*SuggestValuesReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _NamespaceMLCustomAPI_GetApiEndpointsStats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApiEndpointsStatsNSReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NamespaceMLCustomAPIServer).GetApiEndpointsStats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ves.io.schema.namespace.NamespaceMLCustomAPI/GetApiEndpointsStats",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NamespaceMLCustomAPIServer).GetApiEndpointsStats(ctx, req.(*ApiEndpointsStatsNSReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _NamespaceMLCustomAPI_GetApiEndpointsStatsAllNamespaces_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApiEndpointsStatsAllNSReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NamespaceMLCustomAPIServer).GetApiEndpointsStatsAllNamespaces(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ves.io.schema.namespace.NamespaceMLCustomAPI/GetApiEndpointsStatsAllNamespaces",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NamespaceMLCustomAPIServer).GetApiEndpointsStatsAllNamespaces(ctx, req.(*ApiEndpointsStatsAllNSReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _NamespaceMLCustomAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ves.io.schema.namespace.NamespaceMLCustomAPI",
	HandlerType: (*NamespaceMLCustomAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SuggestValues",
			Handler:    _NamespaceMLCustomAPI_SuggestValues_Handler,
		},
		{
			MethodName: "GetApiEndpointsStats",
			Handler:    _NamespaceMLCustomAPI_GetApiEndpointsStats_Handler,
		},
		{
			MethodName: "GetApiEndpointsStatsAllNamespaces",
			Handler:    _NamespaceMLCustomAPI_GetApiEndpointsStatsAllNamespaces_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ves.io/schema/namespace/public_customapi_asterix.proto",
}

func (m *ApiEndpointsStatsNSReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ApiEndpointsStatsNSReq) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ApiEndpointsStatsNSReq) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Namespace) > 0 {
		i -= len(m.Namespace)
		copy(dAtA[i:], m.Namespace)
		i = encodeVarintPublicCustomapiAsterix(dAtA, i, uint64(len(m.Namespace)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.VhostsFilter) > 0 {
		for iNdEx := len(m.VhostsFilter) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.VhostsFilter[iNdEx])
			copy(dAtA[i:], m.VhostsFilter[iNdEx])
			i = encodeVarintPublicCustomapiAsterix(dAtA, i, uint64(len(m.VhostsFilter[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *ApiEndpointsStatsAllNSReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ApiEndpointsStatsAllNSReq) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ApiEndpointsStatsAllNSReq) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Namespace) > 0 {
		i -= len(m.Namespace)
		copy(dAtA[i:], m.Namespace)
		i = encodeVarintPublicCustomapiAsterix(dAtA, i, uint64(len(m.Namespace)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ApiEndpointsStatsNSRsp) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ApiEndpointsStatsNSRsp) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ApiEndpointsStatsNSRsp) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.PiiDetected != 0 {
		i = encodeVarintPublicCustomapiAsterix(dAtA, i, uint64(m.PiiDetected))
		i--
		dAtA[i] = 0x28
	}
	if m.Shadow != 0 {
		i = encodeVarintPublicCustomapiAsterix(dAtA, i, uint64(m.Shadow))
		i--
		dAtA[i] = 0x20
	}
	if m.Inventory != 0 {
		i = encodeVarintPublicCustomapiAsterix(dAtA, i, uint64(m.Inventory))
		i--
		dAtA[i] = 0x18
	}
	if m.Discovered != 0 {
		i = encodeVarintPublicCustomapiAsterix(dAtA, i, uint64(m.Discovered))
		i--
		dAtA[i] = 0x10
	}
	if m.TotalEndpoints != 0 {
		i = encodeVarintPublicCustomapiAsterix(dAtA, i, uint64(m.TotalEndpoints))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintPublicCustomapiAsterix(dAtA []byte, offset int, v uint64) int {
	offset -= sovPublicCustomapiAsterix(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ApiEndpointsStatsNSReq) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.VhostsFilter) > 0 {
		for _, s := range m.VhostsFilter {
			l = len(s)
			n += 1 + l + sovPublicCustomapiAsterix(uint64(l))
		}
	}
	l = len(m.Namespace)
	if l > 0 {
		n += 1 + l + sovPublicCustomapiAsterix(uint64(l))
	}
	return n
}

func (m *ApiEndpointsStatsAllNSReq) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Namespace)
	if l > 0 {
		n += 1 + l + sovPublicCustomapiAsterix(uint64(l))
	}
	return n
}

func (m *ApiEndpointsStatsNSRsp) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.TotalEndpoints != 0 {
		n += 1 + sovPublicCustomapiAsterix(uint64(m.TotalEndpoints))
	}
	if m.Discovered != 0 {
		n += 1 + sovPublicCustomapiAsterix(uint64(m.Discovered))
	}
	if m.Inventory != 0 {
		n += 1 + sovPublicCustomapiAsterix(uint64(m.Inventory))
	}
	if m.Shadow != 0 {
		n += 1 + sovPublicCustomapiAsterix(uint64(m.Shadow))
	}
	if m.PiiDetected != 0 {
		n += 1 + sovPublicCustomapiAsterix(uint64(m.PiiDetected))
	}
	return n
}

func sovPublicCustomapiAsterix(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPublicCustomapiAsterix(x uint64) (n int) {
	return sovPublicCustomapiAsterix(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *ApiEndpointsStatsNSReq) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ApiEndpointsStatsNSReq{`,
		`VhostsFilter:` + fmt.Sprintf("%v", this.VhostsFilter) + `,`,
		`Namespace:` + fmt.Sprintf("%v", this.Namespace) + `,`,
		`}`,
	}, "")
	return s
}
func (this *ApiEndpointsStatsAllNSReq) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ApiEndpointsStatsAllNSReq{`,
		`Namespace:` + fmt.Sprintf("%v", this.Namespace) + `,`,
		`}`,
	}, "")
	return s
}
func (this *ApiEndpointsStatsNSRsp) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ApiEndpointsStatsNSRsp{`,
		`TotalEndpoints:` + fmt.Sprintf("%v", this.TotalEndpoints) + `,`,
		`Discovered:` + fmt.Sprintf("%v", this.Discovered) + `,`,
		`Inventory:` + fmt.Sprintf("%v", this.Inventory) + `,`,
		`Shadow:` + fmt.Sprintf("%v", this.Shadow) + `,`,
		`PiiDetected:` + fmt.Sprintf("%v", this.PiiDetected) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringPublicCustomapiAsterix(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *ApiEndpointsStatsNSReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPublicCustomapiAsterix
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
			return fmt.Errorf("proto: ApiEndpointsStatsNSReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ApiEndpointsStatsNSReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VhostsFilter", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPublicCustomapiAsterix
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
				return ErrInvalidLengthPublicCustomapiAsterix
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPublicCustomapiAsterix
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.VhostsFilter = append(m.VhostsFilter, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Namespace", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPublicCustomapiAsterix
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
				return ErrInvalidLengthPublicCustomapiAsterix
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPublicCustomapiAsterix
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Namespace = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPublicCustomapiAsterix(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPublicCustomapiAsterix
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthPublicCustomapiAsterix
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
func (m *ApiEndpointsStatsAllNSReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPublicCustomapiAsterix
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
			return fmt.Errorf("proto: ApiEndpointsStatsAllNSReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ApiEndpointsStatsAllNSReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Namespace", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPublicCustomapiAsterix
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
				return ErrInvalidLengthPublicCustomapiAsterix
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPublicCustomapiAsterix
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Namespace = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPublicCustomapiAsterix(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPublicCustomapiAsterix
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthPublicCustomapiAsterix
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
func (m *ApiEndpointsStatsNSRsp) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPublicCustomapiAsterix
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
			return fmt.Errorf("proto: ApiEndpointsStatsNSRsp: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ApiEndpointsStatsNSRsp: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalEndpoints", wireType)
			}
			m.TotalEndpoints = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPublicCustomapiAsterix
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TotalEndpoints |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Discovered", wireType)
			}
			m.Discovered = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPublicCustomapiAsterix
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Discovered |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Inventory", wireType)
			}
			m.Inventory = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPublicCustomapiAsterix
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Inventory |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Shadow", wireType)
			}
			m.Shadow = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPublicCustomapiAsterix
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Shadow |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PiiDetected", wireType)
			}
			m.PiiDetected = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPublicCustomapiAsterix
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PiiDetected |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipPublicCustomapiAsterix(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPublicCustomapiAsterix
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthPublicCustomapiAsterix
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
func skipPublicCustomapiAsterix(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPublicCustomapiAsterix
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
					return 0, ErrIntOverflowPublicCustomapiAsterix
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
					return 0, ErrIntOverflowPublicCustomapiAsterix
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
				return 0, ErrInvalidLengthPublicCustomapiAsterix
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPublicCustomapiAsterix
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPublicCustomapiAsterix
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPublicCustomapiAsterix        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPublicCustomapiAsterix          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPublicCustomapiAsterix = fmt.Errorf("proto: unexpected end of group")
)
