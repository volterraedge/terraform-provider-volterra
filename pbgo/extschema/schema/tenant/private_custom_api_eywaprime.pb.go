// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ves.io/schema/tenant/private_custom_api_eywaprime.proto

// Tenant
//
// x-displayName: "Tenant"
// Package for working with Tenant representation.

package tenant

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
	math "math"
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

func init() {
	proto.RegisterFile("ves.io/schema/tenant/private_custom_api_eywaprime.proto", fileDescriptor_841d7cf0a846746a)
}
func init() {
	golang_proto.RegisterFile("ves.io/schema/tenant/private_custom_api_eywaprime.proto", fileDescriptor_841d7cf0a846746a)
}

var fileDescriptor_841d7cf0a846746a = []byte{
	// 404 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x92, 0xbf, 0xaa, 0x13, 0x41,
	0x14, 0xc6, 0x77, 0xf4, 0x62, 0xb1, 0x95, 0x2c, 0x36, 0xae, 0x32, 0x48, 0x4a, 0x61, 0x67, 0x50,
	0x41, 0x41, 0x6e, 0x73, 0xfd, 0x83, 0x68, 0x75, 0xb9, 0x58, 0xd9, 0x84, 0xd9, 0xcd, 0xb9, 0x73,
	0x07, 0xef, 0xee, 0x19, 0x67, 0xce, 0x6e, 0x4c, 0x27, 0xe2, 0x03, 0x08, 0x76, 0x3e, 0x81, 0x9d,
	0x0f, 0x60, 0x93, 0xc2, 0xc2, 0x32, 0x60, 0x93, 0xd2, 0xec, 0x5a, 0x58, 0xe6, 0x11, 0x84, 0xd9,
	0x24, 0x18, 0x49, 0xc0, 0xca, 0xee, 0x1c, 0xce, 0xf7, 0x7d, 0xfc, 0x3e, 0x38, 0xf1, 0xbd, 0x06,
	0xbc, 0x30, 0x28, 0x7d, 0x71, 0x06, 0xa5, 0x92, 0x04, 0x95, 0xaa, 0x48, 0x5a, 0x67, 0x1a, 0x45,
	0x30, 0x2c, 0x6a, 0x4f, 0x58, 0x0e, 0x95, 0x35, 0x43, 0x98, 0x8c, 0x95, 0x75, 0xa6, 0x04, 0x61,
	0x1d, 0x12, 0x26, 0x57, 0x7a, 0xa3, 0xe8, 0x8d, 0xa2, 0x37, 0xa6, 0x99, 0x36, 0x74, 0x56, 0xe7,
	0xa2, 0xc0, 0x52, 0x6a, 0xd4, 0x28, 0x83, 0x38, 0xaf, 0x4f, 0xc3, 0x16, 0x96, 0x30, 0xf5, 0x21,
	0xe9, 0x75, 0x8d, 0xa8, 0xcf, 0x41, 0x2a, 0x6b, 0xa4, 0xaa, 0x2a, 0x24, 0x45, 0x06, 0x2b, 0xbf,
	0xba, 0x5e, 0xdb, 0x66, 0x43, 0xfb, 0xe7, 0xf1, 0xc6, 0x4e, 0x70, 0x9a, 0x58, 0x58, 0x29, 0x6e,
	0x7f, 0xbc, 0x18, 0x5f, 0x7d, 0x18, 0x0a, 0x1c, 0xf7, 0x75, 0x8e, 0x8e, 0x9f, 0x3e, 0x5e, 0xb7,
	0x48, 0xbe, 0xb2, 0x38, 0x3d, 0x01, 0x4f, 0xce, 0x14, 0x04, 0xa3, 0x47, 0xa0, 0x0a, 0x0a, 0xa2,
	0xe7, 0x21, 0x28, 0xc9, 0xc4, 0xae, 0x7e, 0xe2, 0x6f, 0xdd, 0x09, 0xbc, 0xaa, 0xc1, 0x53, 0x2a,
	0xfe, 0x55, 0xee, 0x2d, 0x56, 0x1e, 0x06, 0x4f, 0xde, 0x7e, 0xff, 0xf9, 0xe1, 0xc2, 0xd1, 0xe0,
	0x50, 0x6e, 0xf9, 0xa4, 0xa9, 0xc8, 0xa1, 0xb7, 0x50, 0x90, 0x74, 0x1b, 0x38, 0x39, 0x76, 0x86,
	0x60, 0x5d, 0x71, 0xb4, 0xc9, 0xbc, 0xcf, 0x6e, 0x26, 0x9f, 0x59, 0x7c, 0xf9, 0x7f, 0xc3, 0x1f,
	0x06, 0xf8, 0xbb, 0x83, 0x5b, 0xfb, 0xe1, 0xf7, 0x13, 0xa7, 0x07, 0xd3, 0x2f, 0xec, 0xe0, 0xc1,
	0x3b, 0x36, 0x5b, 0xf0, 0x68, 0xbe, 0xe0, 0xd1, 0x72, 0xc1, 0xd9, 0x9b, 0x96, 0xb3, 0x4f, 0x2d,
	0x67, 0xdf, 0x5a, 0xce, 0x66, 0x2d, 0x67, 0x3f, 0x5a, 0xce, 0x7e, 0xb5, 0x3c, 0x5a, 0xb6, 0x9c,
	0xbd, 0xef, 0x78, 0x34, 0xed, 0x38, 0x9b, 0x75, 0x3c, 0x9a, 0x77, 0x3c, 0x7a, 0xf1, 0x4c, 0xa3,
	0x7d, 0xa9, 0x45, 0x83, 0xe7, 0x04, 0xce, 0x29, 0x51, 0x7b, 0x19, 0x86, 0x53, 0x74, 0x65, 0x66,
	0x1d, 0x36, 0x66, 0x04, 0x2e, 0x5b, 0x9f, 0xa5, 0xcd, 0x35, 0x4a, 0x78, 0x4d, 0x2b, 0xc8, 0xad,
	0x7f, 0xc9, 0x2f, 0x85, 0x57, 0xb9, 0xf3, 0x3b, 0x00, 0x00, 0xff, 0xff, 0xbd, 0x0a, 0xe7, 0x7b,
	0x07, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CustomPrivateAPIEywaprimeClient is the client API for CustomPrivateAPIEywaprime service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CustomPrivateAPIEywaprimeClient interface {
	// RestrictedDeactivateTenant
	//
	// x-displayName: "Restricted Deactivate Tenant"
	// This API mark tenant for deletion queue, after approve it will completely removed from the system. This API proxies the request to eywa’s tenant disable API and allowed on restricted access in all environments, users with access to restricted APIs can only trigger this.
	RestrictedDeactivateTenant(ctx context.Context, in *DeactivateTenantRequest, opts ...grpc.CallOption) (*DeactivateTenantResponse, error)
	// DeactivateTenant
	//
	// x-displayName: "Deactivate Tenant"
	// This API mark tenant for deletion queue, after approve it will completely removed from the system. This API proxies the request to eywa’s tenant disable API and not allowed to be triggered in production, users with normal introspection access can trigger this.
	DeactivateTenant(ctx context.Context, in *DeactivateTenantRequest, opts ...grpc.CallOption) (*DeactivateTenantResponse, error)
}

type customPrivateAPIEywaprimeClient struct {
	cc *grpc.ClientConn
}

func NewCustomPrivateAPIEywaprimeClient(cc *grpc.ClientConn) CustomPrivateAPIEywaprimeClient {
	return &customPrivateAPIEywaprimeClient{cc}
}

func (c *customPrivateAPIEywaprimeClient) RestrictedDeactivateTenant(ctx context.Context, in *DeactivateTenantRequest, opts ...grpc.CallOption) (*DeactivateTenantResponse, error) {
	out := new(DeactivateTenantResponse)
	err := c.cc.Invoke(ctx, "/ves.io.schema.tenant.CustomPrivateAPIEywaprime/RestrictedDeactivateTenant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customPrivateAPIEywaprimeClient) DeactivateTenant(ctx context.Context, in *DeactivateTenantRequest, opts ...grpc.CallOption) (*DeactivateTenantResponse, error) {
	out := new(DeactivateTenantResponse)
	err := c.cc.Invoke(ctx, "/ves.io.schema.tenant.CustomPrivateAPIEywaprime/DeactivateTenant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CustomPrivateAPIEywaprimeServer is the server API for CustomPrivateAPIEywaprime service.
type CustomPrivateAPIEywaprimeServer interface {
	// RestrictedDeactivateTenant
	//
	// x-displayName: "Restricted Deactivate Tenant"
	// This API mark tenant for deletion queue, after approve it will completely removed from the system. This API proxies the request to eywa’s tenant disable API and allowed on restricted access in all environments, users with access to restricted APIs can only trigger this.
	RestrictedDeactivateTenant(context.Context, *DeactivateTenantRequest) (*DeactivateTenantResponse, error)
	// DeactivateTenant
	//
	// x-displayName: "Deactivate Tenant"
	// This API mark tenant for deletion queue, after approve it will completely removed from the system. This API proxies the request to eywa’s tenant disable API and not allowed to be triggered in production, users with normal introspection access can trigger this.
	DeactivateTenant(context.Context, *DeactivateTenantRequest) (*DeactivateTenantResponse, error)
}

// UnimplementedCustomPrivateAPIEywaprimeServer can be embedded to have forward compatible implementations.
type UnimplementedCustomPrivateAPIEywaprimeServer struct {
}

func (*UnimplementedCustomPrivateAPIEywaprimeServer) RestrictedDeactivateTenant(ctx context.Context, req *DeactivateTenantRequest) (*DeactivateTenantResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RestrictedDeactivateTenant not implemented")
}
func (*UnimplementedCustomPrivateAPIEywaprimeServer) DeactivateTenant(ctx context.Context, req *DeactivateTenantRequest) (*DeactivateTenantResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeactivateTenant not implemented")
}

func RegisterCustomPrivateAPIEywaprimeServer(s *grpc.Server, srv CustomPrivateAPIEywaprimeServer) {
	s.RegisterService(&_CustomPrivateAPIEywaprime_serviceDesc, srv)
}

func _CustomPrivateAPIEywaprime_RestrictedDeactivateTenant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeactivateTenantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomPrivateAPIEywaprimeServer).RestrictedDeactivateTenant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ves.io.schema.tenant.CustomPrivateAPIEywaprime/RestrictedDeactivateTenant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomPrivateAPIEywaprimeServer).RestrictedDeactivateTenant(ctx, req.(*DeactivateTenantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CustomPrivateAPIEywaprime_DeactivateTenant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeactivateTenantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomPrivateAPIEywaprimeServer).DeactivateTenant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ves.io.schema.tenant.CustomPrivateAPIEywaprime/DeactivateTenant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomPrivateAPIEywaprimeServer).DeactivateTenant(ctx, req.(*DeactivateTenantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CustomPrivateAPIEywaprime_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ves.io.schema.tenant.CustomPrivateAPIEywaprime",
	HandlerType: (*CustomPrivateAPIEywaprimeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RestrictedDeactivateTenant",
			Handler:    _CustomPrivateAPIEywaprime_RestrictedDeactivateTenant_Handler,
		},
		{
			MethodName: "DeactivateTenant",
			Handler:    _CustomPrivateAPIEywaprime_DeactivateTenant_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ves.io/schema/tenant/private_custom_api_eywaprime.proto",
}
