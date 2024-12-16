// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ves.io/schema/upgrade_status/public_customapi.proto

// Upgrade status related Custom APIs
//
// x-displayName: "Upgrade Status"
// Upgrade status custom APIs

package upgrade_status

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

// GetRequest
//
// x-displayName: "Get upgrade status request"
// Request to get the upgrade status of request
type GetUpgradeStatusRequest struct {
	// namespace
	//
	// x-displayName: "Namespace"
	// x-example: "shared"
	// Fetch upgrade status for the given namespace
	Namespace string `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	// vh_name
	//
	// x-displayName: "Virtual Host Name"
	// x-example: "blogging-app"
	// Fetch upgrade status for the name of site
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (m *GetUpgradeStatusRequest) Reset()      { *m = GetUpgradeStatusRequest{} }
func (*GetUpgradeStatusRequest) ProtoMessage() {}
func (*GetUpgradeStatusRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_140a3d2649a346ad, []int{0}
}
func (m *GetUpgradeStatusRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetUpgradeStatusRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetUpgradeStatusRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetUpgradeStatusRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUpgradeStatusRequest.Merge(m, src)
}
func (m *GetUpgradeStatusRequest) XXX_Size() int {
	return m.Size()
}
func (m *GetUpgradeStatusRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUpgradeStatusRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetUpgradeStatusRequest proto.InternalMessageInfo

func (m *GetUpgradeStatusRequest) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *GetUpgradeStatusRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// GetUpgradeStatusResponse
//
// x-displayName: "Get upgrade status response"
// Response to get the upgrade status
type GetUpgradeStatusResponse struct {
	// released_signatures
	//
	// x-displayName: "Upgrade status"
	// Upgrade Status
	UpgradeStatus *GlobalSpecType `protobuf:"bytes,1,opt,name=upgrade_status,json=upgradeStatus,proto3" json:"upgrade_status,omitempty"`
}

func (m *GetUpgradeStatusResponse) Reset()      { *m = GetUpgradeStatusResponse{} }
func (*GetUpgradeStatusResponse) ProtoMessage() {}
func (*GetUpgradeStatusResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_140a3d2649a346ad, []int{1}
}
func (m *GetUpgradeStatusResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetUpgradeStatusResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetUpgradeStatusResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetUpgradeStatusResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUpgradeStatusResponse.Merge(m, src)
}
func (m *GetUpgradeStatusResponse) XXX_Size() int {
	return m.Size()
}
func (m *GetUpgradeStatusResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUpgradeStatusResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetUpgradeStatusResponse proto.InternalMessageInfo

func (m *GetUpgradeStatusResponse) GetUpgradeStatus() *GlobalSpecType {
	if m != nil {
		return m.UpgradeStatus
	}
	return nil
}

func init() {
	proto.RegisterType((*GetUpgradeStatusRequest)(nil), "ves.io.schema.upgrade_status.GetUpgradeStatusRequest")
	golang_proto.RegisterType((*GetUpgradeStatusRequest)(nil), "ves.io.schema.upgrade_status.GetUpgradeStatusRequest")
	proto.RegisterType((*GetUpgradeStatusResponse)(nil), "ves.io.schema.upgrade_status.GetUpgradeStatusResponse")
	golang_proto.RegisterType((*GetUpgradeStatusResponse)(nil), "ves.io.schema.upgrade_status.GetUpgradeStatusResponse")
}

func init() {
	proto.RegisterFile("ves.io/schema/upgrade_status/public_customapi.proto", fileDescriptor_140a3d2649a346ad)
}
func init() {
	golang_proto.RegisterFile("ves.io/schema/upgrade_status/public_customapi.proto", fileDescriptor_140a3d2649a346ad)
}

var fileDescriptor_140a3d2649a346ad = []byte{
	// 525 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x52, 0xbf, 0x6b, 0x14, 0x41,
	0x14, 0xde, 0xd9, 0x88, 0x90, 0x15, 0x45, 0xb6, 0xd0, 0xe3, 0x3c, 0x06, 0x39, 0x9b, 0x14, 0xd9,
	0x1d, 0xb8, 0x10, 0x0b, 0xb1, 0xf0, 0x47, 0x91, 0xc2, 0x42, 0xb8, 0xd3, 0xc6, 0x26, 0xcc, 0xee,
	0xbd, 0xdb, 0x8c, 0xde, 0xce, 0x1b, 0x67, 0x66, 0x97, 0x04, 0x09, 0x48, 0x4a, 0x11, 0x11, 0xfc,
	0x27, 0xfc, 0x13, 0x84, 0x34, 0xe9, 0x4c, 0x25, 0x87, 0x36, 0x69, 0x04, 0x6f, 0xcf, 0xc2, 0xf2,
	0xfe, 0x04, 0xb9, 0xdd, 0x4b, 0xcc, 0xc6, 0x70, 0x62, 0xf7, 0xde, 0x7c, 0xf3, 0x7d, 0xef, 0x9b,
	0x79, 0x9f, 0xb7, 0x96, 0x83, 0x09, 0x05, 0x32, 0x13, 0x6f, 0x41, 0xca, 0x59, 0xa6, 0x12, 0xcd,
	0xfb, 0xb0, 0x69, 0x2c, 0xb7, 0x99, 0x61, 0x2a, 0x8b, 0x86, 0x22, 0xde, 0x8c, 0x33, 0x63, 0x31,
	0xe5, 0x4a, 0x84, 0x4a, 0xa3, 0x45, 0xbf, 0x55, 0x91, 0xc2, 0x8a, 0x14, 0xd6, 0x49, 0xcd, 0x20,
	0x11, 0x76, 0x2b, 0x8b, 0xc2, 0x18, 0x53, 0x96, 0x60, 0x82, 0xac, 0x24, 0x45, 0xd9, 0xa0, 0xec,
	0xca, 0xa6, 0xac, 0x2a, 0xb1, 0x66, 0x2b, 0x41, 0x4c, 0x86, 0xc0, 0xb8, 0x12, 0x8c, 0x4b, 0x89,
	0x96, 0x5b, 0x81, 0xd2, 0xcc, 0xd1, 0x1b, 0x75, 0x7f, 0xa8, 0x4e, 0x83, 0x2b, 0x0b, 0xcd, 0xdb,
	0x1d, 0x05, 0xc7, 0x37, 0xdb, 0xf5, 0x9b, 0x39, 0x18, 0x90, 0x79, 0x5d, 0xad, 0xfd, 0xc8, 0xbb,
	0xbe, 0x01, 0xf6, 0x69, 0x25, 0xd2, 0x2b, 0x35, 0xba, 0xf0, 0x32, 0x03, 0x63, 0xfd, 0x96, 0xb7,
	0x2c, 0x79, 0x0a, 0x46, 0xf1, 0x18, 0x1a, 0xe4, 0x26, 0x59, 0x59, 0xee, 0xfe, 0x39, 0xf0, 0x7d,
	0xef, 0xc2, 0xac, 0x69, 0xb8, 0x25, 0x50, 0xd6, 0x6d, 0xf4, 0x1a, 0x7f, 0x8b, 0x19, 0x85, 0xd2,
	0x80, 0xdf, 0xf3, 0xae, 0xd4, 0xad, 0x96, 0x92, 0x97, 0x3a, 0xab, 0xe1, 0xa2, 0x7f, 0x0d, 0x37,
	0x86, 0x18, 0xf1, 0x61, 0x4f, 0x41, 0xfc, 0x64, 0x47, 0x41, 0xf7, 0x72, 0x76, 0x5a, 0xbc, 0xf3,
	0xdd, 0xf5, 0xae, 0xd5, 0xc6, 0x3d, 0x2c, 0x97, 0x76, 0x5f, 0x09, 0xff, 0x9d, 0xeb, 0x5d, 0x3d,
	0x6b, 0xc6, 0x5f, 0xff, 0xc7, 0xb0, 0xf3, 0x7f, 0xa2, 0x79, 0xfb, 0x7f, 0x69, 0xd5, 0x9b, 0xdb,
	0x6f, 0xc8, 0xe1, 0x27, 0x97, 0x14, 0x9f, 0x1b, 0xf7, 0x06, 0xeb, 0xdb, 0x71, 0x60, 0x84, 0x85,
	0x20, 0xe5, 0x92, 0x27, 0x90, 0x82, 0xb4, 0x81, 0xb1, 0x5c, 0xf6, 0xb9, 0xee, 0x07, 0x29, 0x4a,
	0x61, 0x51, 0xaf, 0xe6, 0x60, 0x02, 0x81, 0x81, 0x90, 0x03, 0xcd, 0x8d, 0xd5, 0x59, 0x6c, 0x33,
	0x0d, 0x81, 0x06, 0xde, 0xdf, 0xfb, 0xf6, 0xf3, 0x83, 0x7b, 0xd7, 0xbf, 0x33, 0x0f, 0x27, 0x3b,
	0xd9, 0x83, 0x61, 0xaf, 0x4e, 0xea, 0x5d, 0x36, 0x9b, 0x30, 0x3f, 0xd9, 0x3d, 0x93, 0x8c, 0x66,
	0xe7, 0x60, 0x9f, 0x2c, 0x7d, 0xdd, 0x27, 0xb7, 0x16, 0xbe, 0xe5, 0x71, 0xf4, 0x1c, 0x62, 0xbb,
	0xf7, 0xa5, 0xb1, 0x34, 0x75, 0xc9, 0x83, 0xb7, 0x64, 0x34, 0xa6, 0xce, 0xd1, 0x98, 0x3a, 0xd3,
	0x31, 0x25, 0xaf, 0x0b, 0x4a, 0x3e, 0x16, 0x94, 0x1c, 0x16, 0x94, 0x8c, 0x0a, 0x4a, 0x7e, 0x14,
	0x94, 0xfc, 0x2a, 0xa8, 0x33, 0x2d, 0x28, 0x79, 0x3f, 0xa1, 0xce, 0xc1, 0x84, 0x92, 0xd1, 0x84,
	0x3a, 0x47, 0x13, 0xea, 0x3c, 0xeb, 0x26, 0xa8, 0x5e, 0x24, 0x61, 0x8e, 0x43, 0x0b, 0x5a, 0xf3,
	0x70, 0x96, 0xcd, 0x59, 0x31, 0x40, 0x9d, 0x06, 0x4a, 0x63, 0x2e, 0xfa, 0xa0, 0x83, 0x63, 0x98,
	0xa9, 0x28, 0x41, 0x06, 0xdb, 0x76, 0x9e, 0xd5, 0x73, 0xc3, 0x1d, 0x5d, 0x2c, 0x33, 0xbb, 0xf6,
	0x3b, 0x00, 0x00, 0xff, 0xff, 0x1c, 0x03, 0x92, 0x0e, 0xc0, 0x03, 0x00, 0x00,
}

func (this *GetUpgradeStatusRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*GetUpgradeStatusRequest)
	if !ok {
		that2, ok := that.(GetUpgradeStatusRequest)
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
func (this *GetUpgradeStatusResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*GetUpgradeStatusResponse)
	if !ok {
		that2, ok := that.(GetUpgradeStatusResponse)
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
	if !this.UpgradeStatus.Equal(that1.UpgradeStatus) {
		return false
	}
	return true
}
func (this *GetUpgradeStatusRequest) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&upgrade_status.GetUpgradeStatusRequest{")
	s = append(s, "Namespace: "+fmt.Sprintf("%#v", this.Namespace)+",\n")
	s = append(s, "Name: "+fmt.Sprintf("%#v", this.Name)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *GetUpgradeStatusResponse) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&upgrade_status.GetUpgradeStatusResponse{")
	if this.UpgradeStatus != nil {
		s = append(s, "UpgradeStatus: "+fmt.Sprintf("%#v", this.UpgradeStatus)+",\n")
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

// UpgradeStatusCustomApiClient is the client API for UpgradeStatusCustomApi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UpgradeStatusCustomApiClient interface {
	// ReleasedSignatures
	//
	// x-displayName: "Released Signatures"
	// API to get Released Signatures
	GetUpgradeStatus(ctx context.Context, in *GetUpgradeStatusRequest, opts ...grpc.CallOption) (*GetUpgradeStatusResponse, error)
}

type upgradeStatusCustomApiClient struct {
	cc *grpc.ClientConn
}

func NewUpgradeStatusCustomApiClient(cc *grpc.ClientConn) UpgradeStatusCustomApiClient {
	return &upgradeStatusCustomApiClient{cc}
}

func (c *upgradeStatusCustomApiClient) GetUpgradeStatus(ctx context.Context, in *GetUpgradeStatusRequest, opts ...grpc.CallOption) (*GetUpgradeStatusResponse, error) {
	out := new(GetUpgradeStatusResponse)
	err := c.cc.Invoke(ctx, "/ves.io.schema.upgrade_status.UpgradeStatusCustomApi/GetUpgradeStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UpgradeStatusCustomApiServer is the server API for UpgradeStatusCustomApi service.
type UpgradeStatusCustomApiServer interface {
	// ReleasedSignatures
	//
	// x-displayName: "Released Signatures"
	// API to get Released Signatures
	GetUpgradeStatus(context.Context, *GetUpgradeStatusRequest) (*GetUpgradeStatusResponse, error)
}

// UnimplementedUpgradeStatusCustomApiServer can be embedded to have forward compatible implementations.
type UnimplementedUpgradeStatusCustomApiServer struct {
}

func (*UnimplementedUpgradeStatusCustomApiServer) GetUpgradeStatus(ctx context.Context, req *GetUpgradeStatusRequest) (*GetUpgradeStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUpgradeStatus not implemented")
}

func RegisterUpgradeStatusCustomApiServer(s *grpc.Server, srv UpgradeStatusCustomApiServer) {
	s.RegisterService(&_UpgradeStatusCustomApi_serviceDesc, srv)
}

func _UpgradeStatusCustomApi_GetUpgradeStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUpgradeStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UpgradeStatusCustomApiServer).GetUpgradeStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ves.io.schema.upgrade_status.UpgradeStatusCustomApi/GetUpgradeStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UpgradeStatusCustomApiServer).GetUpgradeStatus(ctx, req.(*GetUpgradeStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UpgradeStatusCustomApi_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ves.io.schema.upgrade_status.UpgradeStatusCustomApi",
	HandlerType: (*UpgradeStatusCustomApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUpgradeStatus",
			Handler:    _UpgradeStatusCustomApi_GetUpgradeStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ves.io/schema/upgrade_status/public_customapi.proto",
}

func (m *GetUpgradeStatusRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetUpgradeStatusRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetUpgradeStatusRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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

func (m *GetUpgradeStatusResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetUpgradeStatusResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetUpgradeStatusResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.UpgradeStatus != nil {
		{
			size, err := m.UpgradeStatus.MarshalToSizedBuffer(dAtA[:i])
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
func (m *GetUpgradeStatusRequest) Size() (n int) {
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

func (m *GetUpgradeStatusResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.UpgradeStatus != nil {
		l = m.UpgradeStatus.Size()
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
func (this *GetUpgradeStatusRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&GetUpgradeStatusRequest{`,
		`Namespace:` + fmt.Sprintf("%v", this.Namespace) + `,`,
		`Name:` + fmt.Sprintf("%v", this.Name) + `,`,
		`}`,
	}, "")
	return s
}
func (this *GetUpgradeStatusResponse) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&GetUpgradeStatusResponse{`,
		`UpgradeStatus:` + strings.Replace(fmt.Sprintf("%v", this.UpgradeStatus), "GlobalSpecType", "GlobalSpecType", 1) + `,`,
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
func (m *GetUpgradeStatusRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: GetUpgradeStatusRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetUpgradeStatusRequest: illegal tag %d (wire type %d)", fieldNum, wire)
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
func (m *GetUpgradeStatusResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: GetUpgradeStatusResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetUpgradeStatusResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UpgradeStatus", wireType)
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
			if m.UpgradeStatus == nil {
				m.UpgradeStatus = &GlobalSpecType{}
			}
			if err := m.UpgradeStatus.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
