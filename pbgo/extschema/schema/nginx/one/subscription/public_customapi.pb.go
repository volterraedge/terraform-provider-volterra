// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ves.io/schema/nginx/one/subscription/public_customapi.proto

// NGINX One Subscription API
//
// x-displayName: "NGINX One Subscription API"
// Use this API to subscribe to NGINX One

package subscription

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

// SubscribeRequest
//
// x-displayName: "Subscribe Request"
// Request to subscribe to NGINX One
type SubscribeRequest struct {
}

func (m *SubscribeRequest) Reset()      { *m = SubscribeRequest{} }
func (*SubscribeRequest) ProtoMessage() {}
func (*SubscribeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4221319b380ec252, []int{0}
}
func (m *SubscribeRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SubscribeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SubscribeRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SubscribeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubscribeRequest.Merge(m, src)
}
func (m *SubscribeRequest) XXX_Size() int {
	return m.Size()
}
func (m *SubscribeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SubscribeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SubscribeRequest proto.InternalMessageInfo

// SubscribeResponse
//
// x-displayName: "Subscribe Response"
// Response of subscribe to NGINX One
type SubscribeResponse struct {
}

func (m *SubscribeResponse) Reset()      { *m = SubscribeResponse{} }
func (*SubscribeResponse) ProtoMessage() {}
func (*SubscribeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4221319b380ec252, []int{1}
}
func (m *SubscribeResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SubscribeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SubscribeResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SubscribeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubscribeResponse.Merge(m, src)
}
func (m *SubscribeResponse) XXX_Size() int {
	return m.Size()
}
func (m *SubscribeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SubscribeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SubscribeResponse proto.InternalMessageInfo

// UnsubscribeRequest
//
// x-displayName: "Unsubscribe Request"
// Request to unsubscribe to NGINX One
type UnsubscribeRequest struct {
}

func (m *UnsubscribeRequest) Reset()      { *m = UnsubscribeRequest{} }
func (*UnsubscribeRequest) ProtoMessage() {}
func (*UnsubscribeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4221319b380ec252, []int{2}
}
func (m *UnsubscribeRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UnsubscribeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UnsubscribeRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *UnsubscribeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UnsubscribeRequest.Merge(m, src)
}
func (m *UnsubscribeRequest) XXX_Size() int {
	return m.Size()
}
func (m *UnsubscribeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UnsubscribeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UnsubscribeRequest proto.InternalMessageInfo

// UnsubscribeResponse
//
// x-displayName: "Unsubscribe Response"
// Response of unsubscribe to NGINX One
type UnsubscribeResponse struct {
}

func (m *UnsubscribeResponse) Reset()      { *m = UnsubscribeResponse{} }
func (*UnsubscribeResponse) ProtoMessage() {}
func (*UnsubscribeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4221319b380ec252, []int{3}
}
func (m *UnsubscribeResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UnsubscribeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UnsubscribeResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *UnsubscribeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UnsubscribeResponse.Merge(m, src)
}
func (m *UnsubscribeResponse) XXX_Size() int {
	return m.Size()
}
func (m *UnsubscribeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UnsubscribeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UnsubscribeResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*SubscribeRequest)(nil), "ves.io.schema.nginx.one.subscription.SubscribeRequest")
	golang_proto.RegisterType((*SubscribeRequest)(nil), "ves.io.schema.nginx.one.subscription.SubscribeRequest")
	proto.RegisterType((*SubscribeResponse)(nil), "ves.io.schema.nginx.one.subscription.SubscribeResponse")
	golang_proto.RegisterType((*SubscribeResponse)(nil), "ves.io.schema.nginx.one.subscription.SubscribeResponse")
	proto.RegisterType((*UnsubscribeRequest)(nil), "ves.io.schema.nginx.one.subscription.UnsubscribeRequest")
	golang_proto.RegisterType((*UnsubscribeRequest)(nil), "ves.io.schema.nginx.one.subscription.UnsubscribeRequest")
	proto.RegisterType((*UnsubscribeResponse)(nil), "ves.io.schema.nginx.one.subscription.UnsubscribeResponse")
	golang_proto.RegisterType((*UnsubscribeResponse)(nil), "ves.io.schema.nginx.one.subscription.UnsubscribeResponse")
}

func init() {
	proto.RegisterFile("ves.io/schema/nginx/one/subscription/public_customapi.proto", fileDescriptor_4221319b380ec252)
}
func init() {
	golang_proto.RegisterFile("ves.io/schema/nginx/one/subscription/public_customapi.proto", fileDescriptor_4221319b380ec252)
}

var fileDescriptor_4221319b380ec252 = []byte{
	// 457 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0xbd, 0x6e, 0xd4, 0x40,
	0x10, 0xf6, 0xe6, 0x24, 0xa4, 0x38, 0x14, 0xe0, 0x80, 0x14, 0x19, 0xb4, 0x85, 0x85, 0x10, 0x02,
	0x79, 0x57, 0x80, 0xf8, 0xaf, 0x42, 0x2a, 0x3a, 0x04, 0xa2, 0x49, 0x83, 0x6c, 0x33, 0xb7, 0x59,
	0x61, 0xef, 0x2c, 0x9e, 0xb5, 0x75, 0x74, 0x28, 0x4f, 0x80, 0x04, 0x8f, 0x40, 0xc1, 0x3b, 0x50,
	0x90, 0x2e, 0xa9, 0xd0, 0x49, 0x34, 0x29, 0x39, 0x1f, 0x05, 0x65, 0x1e, 0x01, 0x61, 0x9b, 0xc8,
	0x17, 0x50, 0x14, 0xe8, 0x66, 0xe6, 0x9b, 0x6f, 0xf4, 0x7d, 0xb3, 0xb3, 0xfe, 0x83, 0x1a, 0x48,
	0x68, 0x94, 0x94, 0x6d, 0x41, 0x91, 0x48, 0xa3, 0xb4, 0x99, 0x48, 0x34, 0x20, 0xa9, 0x4a, 0x29,
	0x2b, 0xb5, 0x75, 0x1a, 0x8d, 0xb4, 0x55, 0x9a, 0xeb, 0xec, 0x79, 0x56, 0x91, 0xc3, 0x22, 0xb1,
	0x5a, 0xd8, 0x12, 0x1d, 0x06, 0x97, 0x3a, 0xb2, 0xe8, 0xc8, 0xa2, 0x25, 0x0b, 0x34, 0x20, 0x86,
	0xe4, 0x30, 0x56, 0xda, 0x6d, 0x55, 0xa9, 0xc8, 0xb0, 0x90, 0x0a, 0x15, 0xca, 0x96, 0x9c, 0x56,
	0xe3, 0x36, 0x6b, 0x93, 0x36, 0xea, 0x86, 0x86, 0x17, 0x15, 0xa2, 0xca, 0x41, 0x26, 0x56, 0xcb,
	0xc4, 0x18, 0x74, 0xc9, 0xaf, 0x29, 0xd4, 0xa3, 0x17, 0x16, 0xf5, 0xa2, 0x1d, 0x82, 0xd1, 0x22,
	0x58, 0x03, 0x81, 0xa9, 0x17, 0x7b, 0xa2, 0xc0, 0x3f, 0xf3, 0xb4, 0x53, 0x97, 0xc2, 0x13, 0x78,
	0x55, 0x01, 0xb9, 0x68, 0xd5, 0x3f, 0x3b, 0xa8, 0x91, 0x45, 0x43, 0x10, 0x9d, 0xf3, 0x83, 0x67,
	0x86, 0x8e, 0xb6, 0x9e, 0xf7, 0x57, 0x17, 0xaa, 0x5d, 0xf3, 0x8d, 0x0f, 0x23, 0x7f, 0x79, 0xa3,
	0xdd, 0xce, 0xfa, 0xe3, 0x47, 0xc1, 0x67, 0xe6, 0x2f, 0x1f, 0x0e, 0x0c, 0x6e, 0x8b, 0x93, 0xac,
	0x49, 0x1c, 0x55, 0x15, 0xde, 0xf9, 0x67, 0x5e, 0xaf, 0x7c, 0xbd, 0xd9, 0x5d, 0x3b, 0x3d, 0xbe,
	0x35, 0xc9, 0xe2, 0x54, 0xe7, 0xb9, 0x36, 0x6a, 0xfb, 0xeb, 0xf7, 0x77, 0x4b, 0xd7, 0xa2, 0xcb,
	0xfd, 0x3b, 0x4a, 0x93, 0x14, 0x40, 0x36, 0xc9, 0x80, 0x24, 0xbd, 0x26, 0x07, 0x85, 0x34, 0xd7,
	0xe5, 0xa1, 0xa9, 0xfb, 0xec, 0x6a, 0xb0, 0xcb, 0xfc, 0x95, 0x81, 0xcf, 0xe0, 0xee, 0xc9, 0xb4,
	0xfc, 0xb9, 0xb0, 0xf0, 0xde, 0x7f, 0x30, 0x7b, 0x1f, 0x1b, 0x7f, 0xf5, 0x11, 0x47, 0x57, 0x8e,
	0xf5, 0x51, 0x99, 0xa1, 0x93, 0x70, 0x65, 0xe7, 0x13, 0x1b, 0x6d, 0x7f, 0x59, 0x1b, 0xed, 0x2d,
	0xb1, 0x87, 0xef, 0xd9, 0x74, 0xc6, 0xbd, 0xfd, 0x19, 0xf7, 0x0e, 0x66, 0x9c, 0xbd, 0x69, 0x38,
	0xfb, 0xd8, 0x70, 0xb6, 0xd7, 0x70, 0x36, 0x6d, 0x38, 0xfb, 0xd6, 0x70, 0xf6, 0xa3, 0xe1, 0xde,
	0x41, 0xc3, 0xd9, 0xdb, 0x39, 0xf7, 0x76, 0xe6, 0x9c, 0x4d, 0xe7, 0xdc, 0xdb, 0x9f, 0x73, 0x6f,
	0x73, 0x53, 0xa1, 0x7d, 0xa9, 0x44, 0x8d, 0xb9, 0x83, 0xb2, 0x4c, 0x44, 0x45, 0xb2, 0x0d, 0xc6,
	0x58, 0x16, 0xb1, 0x2d, 0xb1, 0xd6, 0x2f, 0xa0, 0x8c, 0x7f, 0xc3, 0xd2, 0xa6, 0x0a, 0x25, 0x4c,
	0x5c, 0x7f, 0x8a, 0xc7, 0x7e, 0xaf, 0xf4, 0x54, 0x7b, 0x9a, 0x37, 0x7f, 0x06, 0x00, 0x00, 0xff,
	0xff, 0x30, 0x09, 0x96, 0x82, 0x8d, 0x03, 0x00, 0x00,
}

func (this *SubscribeRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*SubscribeRequest)
	if !ok {
		that2, ok := that.(SubscribeRequest)
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
	return true
}
func (this *SubscribeResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*SubscribeResponse)
	if !ok {
		that2, ok := that.(SubscribeResponse)
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
	return true
}
func (this *UnsubscribeRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*UnsubscribeRequest)
	if !ok {
		that2, ok := that.(UnsubscribeRequest)
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
	return true
}
func (this *UnsubscribeResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*UnsubscribeResponse)
	if !ok {
		that2, ok := that.(UnsubscribeResponse)
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
	return true
}
func (this *SubscribeRequest) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 4)
	s = append(s, "&subscription.SubscribeRequest{")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *SubscribeResponse) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 4)
	s = append(s, "&subscription.SubscribeResponse{")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *UnsubscribeRequest) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 4)
	s = append(s, "&subscription.UnsubscribeRequest{")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *UnsubscribeResponse) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 4)
	s = append(s, "&subscription.UnsubscribeResponse{")
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
	// Subscribe
	//
	// x-displayName: "Subscribe to NGINX One"
	// Subscribe to NGINX One
	Subscribe(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (*SubscribeResponse, error)
	// Unsubscribe
	//
	// x-displayName: "Unsubscribe to NGINX One"
	// Unsubscribe to NGINX One
	Unsubscribe(ctx context.Context, in *UnsubscribeRequest, opts ...grpc.CallOption) (*UnsubscribeResponse, error)
}

type customAPIClient struct {
	cc *grpc.ClientConn
}

func NewCustomAPIClient(cc *grpc.ClientConn) CustomAPIClient {
	return &customAPIClient{cc}
}

func (c *customAPIClient) Subscribe(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (*SubscribeResponse, error) {
	out := new(SubscribeResponse)
	err := c.cc.Invoke(ctx, "/ves.io.schema.nginx.one.subscription.CustomAPI/Subscribe", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customAPIClient) Unsubscribe(ctx context.Context, in *UnsubscribeRequest, opts ...grpc.CallOption) (*UnsubscribeResponse, error) {
	out := new(UnsubscribeResponse)
	err := c.cc.Invoke(ctx, "/ves.io.schema.nginx.one.subscription.CustomAPI/Unsubscribe", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CustomAPIServer is the server API for CustomAPI service.
type CustomAPIServer interface {
	// Subscribe
	//
	// x-displayName: "Subscribe to NGINX One"
	// Subscribe to NGINX One
	Subscribe(context.Context, *SubscribeRequest) (*SubscribeResponse, error)
	// Unsubscribe
	//
	// x-displayName: "Unsubscribe to NGINX One"
	// Unsubscribe to NGINX One
	Unsubscribe(context.Context, *UnsubscribeRequest) (*UnsubscribeResponse, error)
}

// UnimplementedCustomAPIServer can be embedded to have forward compatible implementations.
type UnimplementedCustomAPIServer struct {
}

func (*UnimplementedCustomAPIServer) Subscribe(ctx context.Context, req *SubscribeRequest) (*SubscribeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Subscribe not implemented")
}
func (*UnimplementedCustomAPIServer) Unsubscribe(ctx context.Context, req *UnsubscribeRequest) (*UnsubscribeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Unsubscribe not implemented")
}

func RegisterCustomAPIServer(s *grpc.Server, srv CustomAPIServer) {
	s.RegisterService(&_CustomAPI_serviceDesc, srv)
}

func _CustomAPI_Subscribe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubscribeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomAPIServer).Subscribe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ves.io.schema.nginx.one.subscription.CustomAPI/Subscribe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomAPIServer).Subscribe(ctx, req.(*SubscribeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CustomAPI_Unsubscribe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnsubscribeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomAPIServer).Unsubscribe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ves.io.schema.nginx.one.subscription.CustomAPI/Unsubscribe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomAPIServer).Unsubscribe(ctx, req.(*UnsubscribeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CustomAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ves.io.schema.nginx.one.subscription.CustomAPI",
	HandlerType: (*CustomAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Subscribe",
			Handler:    _CustomAPI_Subscribe_Handler,
		},
		{
			MethodName: "Unsubscribe",
			Handler:    _CustomAPI_Unsubscribe_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ves.io/schema/nginx/one/subscription/public_customapi.proto",
}

func (m *SubscribeRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SubscribeRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SubscribeRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *SubscribeResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SubscribeResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SubscribeResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *UnsubscribeRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UnsubscribeRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *UnsubscribeRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *UnsubscribeResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UnsubscribeResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *UnsubscribeResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
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
func (m *SubscribeRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *SubscribeResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *UnsubscribeRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *UnsubscribeResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovPublicCustomapi(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPublicCustomapi(x uint64) (n int) {
	return sovPublicCustomapi(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *SubscribeRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&SubscribeRequest{`,
		`}`,
	}, "")
	return s
}
func (this *SubscribeResponse) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&SubscribeResponse{`,
		`}`,
	}, "")
	return s
}
func (this *UnsubscribeRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&UnsubscribeRequest{`,
		`}`,
	}, "")
	return s
}
func (this *UnsubscribeResponse) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&UnsubscribeResponse{`,
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
func (m *SubscribeRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: SubscribeRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SubscribeRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
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
func (m *SubscribeResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: SubscribeResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SubscribeResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
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
func (m *UnsubscribeRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: UnsubscribeRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UnsubscribeRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
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
func (m *UnsubscribeResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: UnsubscribeResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UnsubscribeResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
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
