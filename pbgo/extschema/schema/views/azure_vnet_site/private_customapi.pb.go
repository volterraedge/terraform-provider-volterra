// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ves.io/schema/views/azure_vnet_site/private_customapi.proto

package azure_vnet_site

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
	views "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/views"
	_ "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/virtual_host"
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
	proto.RegisterFile("ves.io/schema/views/azure_vnet_site/private_customapi.proto", fileDescriptor_825e212eba761e7f)
}
func init() {
	golang_proto.RegisterFile("ves.io/schema/views/azure_vnet_site/private_customapi.proto", fileDescriptor_825e212eba761e7f)
}

var fileDescriptor_825e212eba761e7f = []byte{
	// 518 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x53, 0x31, 0x6f, 0xd3, 0x40,
	0x18, 0xf5, 0x05, 0xc4, 0x90, 0x85, 0x2a, 0x53, 0x09, 0xd5, 0xa9, 0x0a, 0x88, 0xa1, 0xd4, 0x3e,
	0x09, 0x84, 0x84, 0x60, 0x82, 0x8a, 0x81, 0xa9, 0x15, 0x88, 0x81, 0x2e, 0xd6, 0xd9, 0xf9, 0xe2,
	0x1c, 0x8d, 0xfd, 0x1d, 0x77, 0x9f, 0xdd, 0x02, 0xaa, 0x84, 0xfa, 0x07, 0x40, 0xea, 0x9f, 0xe0,
	0x3f, 0x74, 0x29, 0x13, 0x2c, 0xa0, 0x08, 0x96, 0x8e, 0xc4, 0x61, 0x60, 0xec, 0x2f, 0x40, 0x28,
	0x97, 0x04, 0x25, 0x6e, 0x84, 0x3a, 0xb0, 0x7d, 0xe7, 0xf7, 0xde, 0xf9, 0x7d, 0xcf, 0xcf, 0xf5,
	0xfb, 0x05, 0xd8, 0x40, 0xa1, 0xb0, 0x71, 0x17, 0x52, 0x29, 0x0a, 0x05, 0xbb, 0x56, 0xc8, 0xd7,
	0xb9, 0x81, 0xb0, 0xc8, 0x80, 0x42, 0xab, 0x08, 0x84, 0x36, 0xaa, 0x90, 0x04, 0x61, 0x9c, 0x5b,
	0xc2, 0x54, 0x6a, 0x15, 0x68, 0x83, 0x84, 0x8d, 0x6b, 0x63, 0x71, 0x30, 0x16, 0x07, 0x4e, 0x1c,
	0x54, 0xc4, 0x4d, 0x3f, 0x51, 0xd4, 0xcd, 0xa3, 0x20, 0xc6, 0x54, 0x24, 0x98, 0xa0, 0x70, 0xda,
	0x28, 0xef, 0xb8, 0x93, 0x3b, 0xb8, 0x69, 0x7c, 0x67, 0x73, 0x25, 0x41, 0x4c, 0x7a, 0x20, 0xa4,
	0x56, 0x42, 0x66, 0x19, 0x92, 0x24, 0x85, 0x99, 0x9d, 0xa0, 0x57, 0xe7, 0xed, 0xa2, 0x9e, 0x05,
	0xaf, 0xcc, 0x83, 0xf4, 0x4a, 0xc3, 0x14, 0x5a, 0xa9, 0xac, 0x29, 0x7b, 0xaa, 0x2d, 0x09, 0x26,
	0x68, 0xab, 0x82, 0x82, 0x85, 0xac, 0xa8, 0x5c, 0xbe, 0x7a, 0x36, 0xa8, 0x70, 0x9e, 0x71, 0x7d,
	0x51, 0x94, 0xa3, 0x08, 0xc2, 0x59, 0x27, 0x37, 0xaa, 0x2c, 0x43, 0xb9, 0xec, 0x85, 0x5d, 0xb4,
	0x34, 0xeb, 0xf8, 0xd6, 0x97, 0x0b, 0xf5, 0xa5, 0xad, 0x71, 0xee, 0x1b, 0x2e, 0xf6, 0x07, 0x5b,
	0x8f, 0x1b, 0xbf, 0x6b, 0xf5, 0xcb, 0xcf, 0xf4, 0xc8, 0xf9, 0x53, 0x45, 0xf0, 0xc8, 0x18, 0x34,
	0x8d, 0x9b, 0xc1, 0xa2, 0xaf, 0x50, 0x61, 0x3d, 0x81, 0x97, 0x39, 0x58, 0x6a, 0xae, 0x9f, 0x8f,
	0x6c, 0x35, 0x66, 0x16, 0x5a, 0xef, 0x6a, 0xe5, 0xa7, 0xe5, 0x8d, 0xce, 0x9d, 0xbd, 0xd8, 0x1f,
	0xed, 0xe1, 0xa7, 0x32, 0x93, 0x09, 0xa4, 0x90, 0x91, 0x1f, 0x49, 0xab, 0x62, 0x3f, 0xb7, 0x60,
	0xd6, 0x57, 0x17, 0x12, 0x2c, 0xc9, 0xac, 0x2d, 0x4d, 0xdb, 0x71, 0x0e, 0xbe, 0xff, 0x3c, 0xac,
	0x7d, 0x64, 0xad, 0xcd, 0x69, 0x95, 0xc4, 0xb8, 0x4a, 0x22, 0x93, 0x29, 0x58, 0x2d, 0x63, 0xb0,
	0xe2, 0xcd, 0xdf, 0x79, 0xff, 0x4c, 0x01, 0x1d, 0xb6, 0x2f, 0x72, 0xe7, 0xd3, 0x3d, 0x0a, 0x61,
	0xe4, 0xf4, 0x1e, 0x5b, 0xdb, 0xde, 0x69, 0x75, 0xc4, 0xdc, 0x52, 0x42, 0x65, 0x64, 0xd0, 0x6a,
	0x88, 0x49, 0xec, 0x1a, 0xf5, 0x3f, 0x5f, 0xd6, 0xbc, 0x7b, 0x7c, 0xc4, 0x2e, 0x7e, 0x3b, 0x62,
	0x6b, 0xe7, 0x68, 0x7e, 0xb0, 0x19, 0xbd, 0x80, 0x98, 0x0e, 0xbe, 0x2e, 0xd7, 0x96, 0xd8, 0xc3,
	0x43, 0xd6, 0x1f, 0x70, 0xef, 0x64, 0xc0, 0xbd, 0xd3, 0x01, 0x67, 0x6f, 0x4b, 0xce, 0x3e, 0x94,
	0x9c, 0x7d, 0x2e, 0x39, 0xeb, 0x97, 0x9c, 0xfd, 0x28, 0x39, 0xfb, 0x55, 0x72, 0xef, 0xb4, 0xe4,
	0xec, 0xfd, 0x90, 0x7b, 0xc7, 0x43, 0xce, 0xfa, 0x43, 0xee, 0x9d, 0x0c, 0xb9, 0xb7, 0xfd, 0x3c,
	0x41, 0xbd, 0x93, 0x04, 0x05, 0xf6, 0x08, 0x8c, 0x91, 0x41, 0x6e, 0x85, 0x1b, 0x3a, 0x68, 0x52,
	0x5f, 0x1b, 0x2c, 0x54, 0x1b, 0x8c, 0x3f, 0x85, 0x85, 0x8e, 0x12, 0x14, 0xb0, 0x47, 0x93, 0x34,
	0xfe, 0xf5, 0x67, 0x47, 0x97, 0x5c, 0xd9, 0x6e, 0xff, 0x09, 0x00, 0x00, 0xff, 0xff, 0xbe, 0xce,
	0xa9, 0xbb, 0x07, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PrivateCustomAPIClient is the client API for PrivateCustomAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PrivateCustomAPIClient interface {
	// UpdateSiteError
	//
	// x-displayName: "UpdateSiteError"
	// Private API to update service domains from service.
	UpdateSiteError(ctx context.Context, in *views.UpdateSiteErrorRequest, opts ...grpc.CallOption) (*views.UpdateSiteErrorResponse, error)
}

type privateCustomAPIClient struct {
	cc *grpc.ClientConn
}

func NewPrivateCustomAPIClient(cc *grpc.ClientConn) PrivateCustomAPIClient {
	return &privateCustomAPIClient{cc}
}

func (c *privateCustomAPIClient) UpdateSiteError(ctx context.Context, in *views.UpdateSiteErrorRequest, opts ...grpc.CallOption) (*views.UpdateSiteErrorResponse, error) {
	out := new(views.UpdateSiteErrorResponse)
	err := c.cc.Invoke(ctx, "/ves.io.schema.views.azure_vnet_site.PrivateCustomAPI/UpdateSiteError", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PrivateCustomAPIServer is the server API for PrivateCustomAPI service.
type PrivateCustomAPIServer interface {
	// UpdateSiteError
	//
	// x-displayName: "UpdateSiteError"
	// Private API to update service domains from service.
	UpdateSiteError(context.Context, *views.UpdateSiteErrorRequest) (*views.UpdateSiteErrorResponse, error)
}

// UnimplementedPrivateCustomAPIServer can be embedded to have forward compatible implementations.
type UnimplementedPrivateCustomAPIServer struct {
}

func (*UnimplementedPrivateCustomAPIServer) UpdateSiteError(ctx context.Context, req *views.UpdateSiteErrorRequest) (*views.UpdateSiteErrorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateSiteError not implemented")
}

func RegisterPrivateCustomAPIServer(s *grpc.Server, srv PrivateCustomAPIServer) {
	s.RegisterService(&_PrivateCustomAPI_serviceDesc, srv)
}

func _PrivateCustomAPI_UpdateSiteError_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(views.UpdateSiteErrorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PrivateCustomAPIServer).UpdateSiteError(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ves.io.schema.views.azure_vnet_site.PrivateCustomAPI/UpdateSiteError",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PrivateCustomAPIServer).UpdateSiteError(ctx, req.(*views.UpdateSiteErrorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PrivateCustomAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ves.io.schema.views.azure_vnet_site.PrivateCustomAPI",
	HandlerType: (*PrivateCustomAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateSiteError",
			Handler:    _PrivateCustomAPI_UpdateSiteError_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ves.io/schema/views/azure_vnet_site/private_customapi.proto",
}
