// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ves.io/schema/views/tcp_loadbalancer/public_customapi.proto

// TCP Proxy
//
// x-displayName: "Configure TCP Load Balancer"
// TCP Load Balancer view defines a required parameters that can be used in CRUD, to create and manage TCP Load Balancer.
// It can be used to create TCP Load Balancer and TCP Load Balancer with SNI.
//
// View will create following child objects.
//
// * Virtual-host
// * default route
// * clusters
// * endpoints
// * advertise policy
//

package tcp_loadbalancer

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
	virtual_host_dns_info "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/virtual_host_dns_info"
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

// Get DNS info Request
//
// x-displayName: "Get DNS Info Request"
// Request message for get-dns-info API
type GetDnsInfoRequest struct {
	// Namespace
	//
	// x-displayName: "Namespace"
	// x-example: "value"
	// Namespace for the TCP load balancer
	Namespace string `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	// Name
	//
	// x-displayName: "Name"
	// x-example: "value"
	// Name of the TCP load balancer
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (m *GetDnsInfoRequest) Reset()      { *m = GetDnsInfoRequest{} }
func (*GetDnsInfoRequest) ProtoMessage() {}
func (*GetDnsInfoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d38103af7fac2e92, []int{0}
}
func (m *GetDnsInfoRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetDnsInfoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetDnsInfoRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetDnsInfoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDnsInfoRequest.Merge(m, src)
}
func (m *GetDnsInfoRequest) XXX_Size() int {
	return m.Size()
}
func (m *GetDnsInfoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDnsInfoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetDnsInfoRequest proto.InternalMessageInfo

func (m *GetDnsInfoRequest) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *GetDnsInfoRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// GetDnsInfoResponse
//
// x-displayName: "Get DNS Info Response"
// Response for get-dns-info API
type GetDnsInfoResponse struct {
	// DNS information
	//
	// x-displayName: "DNS information"
	// DNS information object for this TCP load balancer
	DnsInfo *virtual_host_dns_info.GlobalSpecType `protobuf:"bytes,1,opt,name=dns_info,json=dnsInfo,proto3" json:"dns_info,omitempty"`
}

func (m *GetDnsInfoResponse) Reset()      { *m = GetDnsInfoResponse{} }
func (*GetDnsInfoResponse) ProtoMessage() {}
func (*GetDnsInfoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d38103af7fac2e92, []int{1}
}
func (m *GetDnsInfoResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetDnsInfoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetDnsInfoResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetDnsInfoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDnsInfoResponse.Merge(m, src)
}
func (m *GetDnsInfoResponse) XXX_Size() int {
	return m.Size()
}
func (m *GetDnsInfoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDnsInfoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetDnsInfoResponse proto.InternalMessageInfo

func (m *GetDnsInfoResponse) GetDnsInfo() *virtual_host_dns_info.GlobalSpecType {
	if m != nil {
		return m.DnsInfo
	}
	return nil
}

func init() {
	proto.RegisterType((*GetDnsInfoRequest)(nil), "ves.io.schema.views.tcp_loadbalancer.GetDnsInfoRequest")
	golang_proto.RegisterType((*GetDnsInfoRequest)(nil), "ves.io.schema.views.tcp_loadbalancer.GetDnsInfoRequest")
	proto.RegisterType((*GetDnsInfoResponse)(nil), "ves.io.schema.views.tcp_loadbalancer.GetDnsInfoResponse")
	golang_proto.RegisterType((*GetDnsInfoResponse)(nil), "ves.io.schema.views.tcp_loadbalancer.GetDnsInfoResponse")
}

func init() {
	proto.RegisterFile("ves.io/schema/views/tcp_loadbalancer/public_customapi.proto", fileDescriptor_d38103af7fac2e92)
}
func init() {
	golang_proto.RegisterFile("ves.io/schema/views/tcp_loadbalancer/public_customapi.proto", fileDescriptor_d38103af7fac2e92)
}

var fileDescriptor_d38103af7fac2e92 = []byte{
	// 553 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0x3f, 0x6b, 0x14, 0x4f,
	0x18, 0xbe, 0xb9, 0xdf, 0x0f, 0x35, 0x6b, 0xa3, 0x5b, 0x1d, 0x67, 0x18, 0xc2, 0x61, 0x21, 0xe8,
	0xce, 0x48, 0x82, 0xa8, 0x58, 0xa9, 0xd1, 0x90, 0x46, 0x25, 0x5a, 0xa5, 0x39, 0x66, 0x77, 0xdf,
	0xdb, 0x1b, 0xdd, 0x9d, 0x77, 0x9c, 0x99, 0x3d, 0x13, 0x24, 0x20, 0x69, 0xc4, 0x4e, 0xd0, 0x0f,
	0xe1, 0x47, 0x10, 0xd2, 0xa4, 0x33, 0x95, 0x84, 0xd8, 0xa4, 0x34, 0x7b, 0x16, 0x96, 0xf9, 0x08,
	0x72, 0x73, 0x97, 0x3f, 0x6e, 0x44, 0x62, 0xf7, 0xbe, 0xfb, 0xf0, 0x3c, 0x33, 0xcf, 0x33, 0xcf,
	0x06, 0x77, 0x06, 0x60, 0x99, 0x44, 0x6e, 0x93, 0x3e, 0x14, 0x82, 0x0f, 0x24, 0xbc, 0xb2, 0xdc,
	0x25, 0xba, 0x9b, 0xa3, 0x48, 0x63, 0x91, 0x0b, 0x95, 0x80, 0xe1, 0xba, 0x8c, 0x73, 0x99, 0x74,
	0x93, 0xd2, 0x3a, 0x2c, 0x84, 0x96, 0x4c, 0x1b, 0x74, 0x18, 0x5e, 0x1e, 0x93, 0xd9, 0x98, 0xcc,
	0x3c, 0x99, 0xd5, 0xc9, 0xed, 0x28, 0x93, 0xae, 0x5f, 0xc6, 0x2c, 0xc1, 0x82, 0x67, 0x98, 0x21,
	0xf7, 0xe4, 0xb8, 0xec, 0xf9, 0xcd, 0x2f, 0x7e, 0x1a, 0x8b, 0xb6, 0xa7, 0x33, 0xc4, 0x2c, 0x07,
	0x2e, 0xb4, 0xe4, 0x42, 0x29, 0x74, 0xc2, 0x49, 0x54, 0x76, 0x82, 0x5e, 0xfa, 0xfd, 0xbe, 0xa8,
	0x8f, 0x83, 0xd3, 0x35, 0x33, 0x22, 0x97, 0xa9, 0x70, 0x30, 0x41, 0x3b, 0x35, 0x14, 0x2c, 0xa8,
	0x41, 0x4d, 0xe1, 0x7a, 0x3d, 0x0e, 0xe3, 0x4a, 0x91, 0x77, 0xfb, 0x68, 0x5d, 0x37, 0x55, 0xb6,
	0x2b, 0x55, 0x0f, 0x39, 0xc6, 0xcf, 0x21, 0x71, 0x63, 0x46, 0xe7, 0x41, 0x70, 0x71, 0x01, 0xdc,
	0xbc, 0xb2, 0x8b, 0xaa, 0x87, 0x4b, 0xf0, 0xb2, 0x04, 0xeb, 0xc2, 0xe9, 0x60, 0x4a, 0x89, 0x02,
	0xac, 0x16, 0x09, 0xb4, 0xc8, 0x0c, 0xb9, 0x32, 0xb5, 0x74, 0xf4, 0x21, 0x0c, 0x83, 0xff, 0x47,
	0x4b, 0xab, 0xe9, 0x01, 0x3f, 0x77, 0xd2, 0x20, 0x3c, 0x2e, 0x63, 0x35, 0x2a, 0x0b, 0xe1, 0xa3,
	0xe0, 0xdc, 0xc1, 0xa9, 0x5e, 0xe6, 0xfc, 0xec, 0x1c, 0xab, 0x67, 0xfe, 0x87, 0x1b, 0xb2, 0x85,
	0x1c, 0x63, 0x91, 0x3f, 0xd5, 0x90, 0x3c, 0x5b, 0xd5, 0xb0, 0x74, 0x36, 0x1d, 0xeb, 0xce, 0xee,
	0x34, 0x83, 0xa9, 0xfb, 0xfe, 0x11, 0xef, 0x3e, 0x59, 0x0c, 0xdf, 0x35, 0x83, 0xe0, 0xe8, 0xd0,
	0xf0, 0x26, 0x3b, 0xcd, 0x73, 0xb2, 0x13, 0x6e, 0xdb, 0xb7, 0xfe, 0x9d, 0x38, 0xf6, 0xd7, 0x79,
	0x4b, 0xb6, 0x3e, 0x37, 0x49, 0xf5, 0xa5, 0x35, 0xd7, 0xbb, 0xb1, 0x92, 0x44, 0x16, 0x92, 0xd2,
	0x40, 0x01, 0xb6, 0x1f, 0x59, 0x27, 0x54, 0x2a, 0x4c, 0x1a, 0x15, 0xa8, 0xa4, 0x43, 0x73, 0x6d,
	0x66, 0x00, 0x36, 0x92, 0x18, 0x69, 0x83, 0x2b, 0xab, 0x91, 0x01, 0x91, 0xae, 0x7f, 0xfb, 0xf1,
	0xa1, 0xf9, 0x30, 0x9c, 0x9f, 0xb4, 0x93, 0x1f, 0x06, 0x6c, 0xf9, 0xeb, 0xc3, 0x79, 0xed, 0x44,
	0x99, 0x27, 0xe8, 0x1a, 0xcf, 0xc0, 0x45, 0xa9, 0xb2, 0xd1, 0x28, 0xb1, 0xf6, 0xed, 0xcd, 0x0d,
	0xf2, 0xdf, 0xce, 0x06, 0xb9, 0x7a, 0x2a, 0x2b, 0x8f, 0x7d, 0x03, 0xd6, 0xbf, 0xb6, 0x9a, 0x17,
	0xc8, 0xbd, 0x8f, 0x64, 0x7b, 0x8f, 0x36, 0x76, 0xf7, 0x68, 0x63, 0x7f, 0x8f, 0x92, 0x37, 0x15,
	0x25, 0x9f, 0x2a, 0x4a, 0xb6, 0x2a, 0x4a, 0xb6, 0x2b, 0x4a, 0xbe, 0x57, 0x94, 0xfc, 0xac, 0x68,
	0x63, 0xbf, 0xa2, 0xe4, 0xfd, 0x90, 0x36, 0x36, 0x87, 0x94, 0x6c, 0x0f, 0x69, 0x63, 0x77, 0x48,
	0x1b, 0xcb, 0xcb, 0x19, 0xea, 0x17, 0x19, 0x1b, 0x60, 0xee, 0xc0, 0x18, 0xc1, 0x4a, 0xcb, 0xfd,
	0xd0, 0x43, 0x53, 0x8c, 0x9c, 0x0e, 0x64, 0x0a, 0x26, 0x3a, 0x80, 0xb9, 0x8e, 0x33, 0xe4, 0xb0,
	0xe2, 0x26, 0xed, 0xfc, 0xeb, 0x3f, 0x1b, 0x9f, 0xf1, 0xfd, 0x9c, 0xfb, 0x15, 0x00, 0x00, 0xff,
	0xff, 0xbc, 0x6b, 0x06, 0x37, 0xe2, 0x03, 0x00, 0x00,
}

func (this *GetDnsInfoRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*GetDnsInfoRequest)
	if !ok {
		that2, ok := that.(GetDnsInfoRequest)
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
	if this.Name != that1.Name {
		return false
	}
	return true
}
func (this *GetDnsInfoResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*GetDnsInfoResponse)
	if !ok {
		that2, ok := that.(GetDnsInfoResponse)
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
	if !this.DnsInfo.Equal(that1.DnsInfo) {
		return false
	}
	return true
}
func (this *GetDnsInfoRequest) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&tcp_loadbalancer.GetDnsInfoRequest{")
	s = append(s, "Namespace: "+fmt.Sprintf("%#v", this.Namespace)+",\n")
	s = append(s, "Name: "+fmt.Sprintf("%#v", this.Name)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *GetDnsInfoResponse) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&tcp_loadbalancer.GetDnsInfoResponse{")
	if this.DnsInfo != nil {
		s = append(s, "DnsInfo: "+fmt.Sprintf("%#v", this.DnsInfo)+",\n")
	}
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
	// GetDnsInfo
	//
	// x-displayName: "Get DNS Info"
	// GetDnsInfo is an API to get DNS information for a given TCP load balancer
	GetDnsInfo(ctx context.Context, in *GetDnsInfoRequest, opts ...grpc.CallOption) (*GetDnsInfoResponse, error)
}

type customAPIClient struct {
	cc *grpc.ClientConn
}

func NewCustomAPIClient(cc *grpc.ClientConn) CustomAPIClient {
	return &customAPIClient{cc}
}

func (c *customAPIClient) GetDnsInfo(ctx context.Context, in *GetDnsInfoRequest, opts ...grpc.CallOption) (*GetDnsInfoResponse, error) {
	out := new(GetDnsInfoResponse)
	err := c.cc.Invoke(ctx, "/ves.io.schema.views.tcp_loadbalancer.CustomAPI/GetDnsInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CustomAPIServer is the server API for CustomAPI service.
type CustomAPIServer interface {
	// GetDnsInfo
	//
	// x-displayName: "Get DNS Info"
	// GetDnsInfo is an API to get DNS information for a given TCP load balancer
	GetDnsInfo(context.Context, *GetDnsInfoRequest) (*GetDnsInfoResponse, error)
}

// UnimplementedCustomAPIServer can be embedded to have forward compatible implementations.
type UnimplementedCustomAPIServer struct {
}

func (*UnimplementedCustomAPIServer) GetDnsInfo(ctx context.Context, req *GetDnsInfoRequest) (*GetDnsInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDnsInfo not implemented")
}

func RegisterCustomAPIServer(s *grpc.Server, srv CustomAPIServer) {
	s.RegisterService(&_CustomAPI_serviceDesc, srv)
}

func _CustomAPI_GetDnsInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDnsInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomAPIServer).GetDnsInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ves.io.schema.views.tcp_loadbalancer.CustomAPI/GetDnsInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomAPIServer).GetDnsInfo(ctx, req.(*GetDnsInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CustomAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ves.io.schema.views.tcp_loadbalancer.CustomAPI",
	HandlerType: (*CustomAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetDnsInfo",
			Handler:    _CustomAPI_GetDnsInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ves.io/schema/views/tcp_loadbalancer/public_customapi.proto",
}

func (m *GetDnsInfoRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetDnsInfoRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetDnsInfoRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintPublicCustomapi(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
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

func (m *GetDnsInfoResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetDnsInfoResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetDnsInfoResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.DnsInfo != nil {
		{
			size, err := m.DnsInfo.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintPublicCustomapi(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
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
func (m *GetDnsInfoRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Namespace)
	if l > 0 {
		n += 1 + l + sovPublicCustomapi(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovPublicCustomapi(uint64(l))
	}
	return n
}

func (m *GetDnsInfoResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.DnsInfo != nil {
		l = m.DnsInfo.Size()
		n += 1 + l + sovPublicCustomapi(uint64(l))
	}
	return n
}

func sovPublicCustomapi(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPublicCustomapi(x uint64) (n int) {
	return sovPublicCustomapi(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *GetDnsInfoRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&GetDnsInfoRequest{`,
		`Namespace:` + fmt.Sprintf("%v", this.Namespace) + `,`,
		`Name:` + fmt.Sprintf("%v", this.Name) + `,`,
		`}`,
	}, "")
	return s
}
func (this *GetDnsInfoResponse) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&GetDnsInfoResponse{`,
		`DnsInfo:` + strings.Replace(fmt.Sprintf("%v", this.DnsInfo), "GlobalSpecType", "virtual_host_dns_info.GlobalSpecType", 1) + `,`,
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
func (m *GetDnsInfoRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: GetDnsInfoRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetDnsInfoRequest: illegal tag %d (wire type %d)", fieldNum, wire)
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
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
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
			m.Name = string(dAtA[iNdEx:postIndex])
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
func (m *GetDnsInfoResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: GetDnsInfoResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetDnsInfoResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DnsInfo", wireType)
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
			if m.DnsInfo == nil {
				m.DnsInfo = &virtual_host_dns_info.GlobalSpecType{}
			}
			if err := m.DnsInfo.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
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
