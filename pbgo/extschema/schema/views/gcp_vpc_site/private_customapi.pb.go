// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ves.io/schema/views/gcp_vpc_site/private_customapi.proto

package gcp_vpc_site

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
	proto.RegisterFile("ves.io/schema/views/gcp_vpc_site/private_customapi.proto", fileDescriptor_6dadb40db7df8027)
}
func init() {
	golang_proto.RegisterFile("ves.io/schema/views/gcp_vpc_site/private_customapi.proto", fileDescriptor_6dadb40db7df8027)
}

var fileDescriptor_6dadb40db7df8027 = []byte{
	// 515 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x53, 0x31, 0x8b, 0x13, 0x41,
	0x18, 0xdd, 0x89, 0x62, 0x91, 0xc6, 0x23, 0xd5, 0x19, 0x8f, 0x21, 0x04, 0x51, 0xd0, 0xdb, 0x1d,
	0x50, 0x0e, 0xc4, 0x4e, 0x0f, 0x0b, 0x41, 0xf0, 0x50, 0xb4, 0xb8, 0x66, 0x99, 0x9d, 0x7c, 0xd9,
	0x8c, 0x66, 0xe7, 0x1b, 0x67, 0x66, 0xf7, 0x4e, 0xe4, 0x40, 0xae, 0xb3, 0x3b, 0xf0, 0x4f, 0xf8,
	0x1f, 0xd2, 0x5c, 0x21, 0x68, 0x25, 0x41, 0x9b, 0x2b, 0xcd, 0xc6, 0xc2, 0xf2, 0x4a, 0x4b, 0xc9,
	0x24, 0x91, 0x64, 0x8d, 0x68, 0x71, 0xdd, 0x37, 0xfb, 0xde, 0x9b, 0x7d, 0xdf, 0xdb, 0xb7, 0xf5,
	0xdb, 0x05, 0xd8, 0x48, 0x22, 0xb3, 0xa2, 0x07, 0x19, 0x67, 0x85, 0x84, 0x3d, 0xcb, 0x52, 0xa1,
	0xe3, 0x42, 0x8b, 0xd8, 0x4a, 0x07, 0x4c, 0x1b, 0x59, 0x70, 0x07, 0xb1, 0xc8, 0xad, 0xc3, 0x8c,
	0x6b, 0x19, 0x69, 0x83, 0x0e, 0x1b, 0xad, 0xa9, 0x32, 0x9a, 0x2a, 0x23, 0xaf, 0x8c, 0x16, 0x95,
	0xcd, 0x30, 0x95, 0xae, 0x97, 0x27, 0x91, 0xc0, 0x8c, 0xa5, 0x98, 0x22, 0xf3, 0xc2, 0x24, 0xef,
	0xfa, 0x93, 0x3f, 0xf8, 0x69, 0x7a, 0x61, 0x73, 0x23, 0x45, 0x4c, 0xfb, 0xc0, 0xb8, 0x96, 0x8c,
	0x2b, 0x85, 0x8e, 0x3b, 0x89, 0xca, 0xce, 0xd0, 0xcb, 0xcb, 0x46, 0x51, 0x2f, 0x82, 0x97, 0x96,
	0x41, 0xf7, 0x4a, 0xc3, 0x1c, 0xda, 0xa8, 0x2c, 0xc8, 0xfb, 0xb2, 0xc3, 0x1d, 0xcc, 0xd0, 0x76,
	0x05, 0x05, 0x0b, 0xaa, 0xa8, 0x5c, 0xde, 0xfa, 0x33, 0xa2, 0x78, 0x99, 0x71, 0x65, 0x55, 0x88,
	0x93, 0x08, 0xe2, 0x45, 0x27, 0x57, 0xab, 0x2c, 0xe3, 0x72, 0xde, 0x8f, 0x7b, 0x68, 0xdd, 0xa2,
	0xe3, 0x9b, 0x1f, 0xce, 0xd5, 0xd7, 0x76, 0xa6, 0xa1, 0x6f, 0xfb, 0xcc, 0xef, 0xee, 0x3c, 0x68,
	0xfc, 0xac, 0xd5, 0x2f, 0x3e, 0xd5, 0x13, 0xe7, 0x4f, 0xa4, 0x83, 0xfb, 0xc6, 0xa0, 0x69, 0xdc,
	0x88, 0x56, 0x7d, 0x82, 0x0a, 0xeb, 0x31, 0xbc, 0xcc, 0xc1, 0xba, 0xe6, 0xe6, 0xff, 0x91, 0xad,
	0x46, 0x65, 0xa1, 0xfd, 0xb6, 0x56, 0x7e, 0x5c, 0xdf, 0xee, 0x6e, 0xed, 0x8b, 0x70, 0xb2, 0x47,
	0x98, 0x71, 0xc5, 0x53, 0xc8, 0x40, 0xb9, 0x30, 0xe1, 0x56, 0x8a, 0x30, 0xb7, 0x60, 0x36, 0x5b,
	0x2b, 0x09, 0xd6, 0x71, 0xd5, 0xe1, 0xa6, 0xe3, 0x39, 0x87, 0x5f, 0xbf, 0xbf, 0xab, 0x0d, 0x48,
	0xfb, 0xe1, 0xbc, 0x47, 0x6c, 0xda, 0x23, 0xa6, 0x78, 0x06, 0x56, 0x73, 0x01, 0x96, 0xbd, 0xfe,
	0x3d, 0x1f, 0x2c, 0x57, 0xcf, 0x03, 0x07, 0x2c, 0xf7, 0x26, 0xfd, 0xa3, 0x18, 0x26, 0x36, 0xef,
	0x90, 0xeb, 0xbb, 0xbd, 0xb6, 0x60, 0x4b, 0x1b, 0x31, 0xa9, 0x9c, 0x41, 0xab, 0x41, 0x38, 0xb6,
	0x67, 0xe4, 0x99, 0xbd, 0xa9, 0xb9, 0x75, 0x3c, 0x20, 0xe7, 0xbf, 0x0c, 0xc8, 0xb5, 0x7f, 0x15,
	0x3e, 0x7a, 0x94, 0x3c, 0x07, 0xe1, 0x0e, 0x3f, 0xaf, 0xd7, 0xd6, 0xc8, 0xbd, 0x23, 0x32, 0x1c,
	0xd1, 0xe0, 0x64, 0x44, 0x83, 0xd3, 0x11, 0x25, 0x6f, 0x4a, 0x4a, 0xde, 0x97, 0x94, 0x7c, 0x2a,
	0x29, 0x19, 0x96, 0x94, 0x7c, 0x2b, 0x29, 0xf9, 0x51, 0xd2, 0xe0, 0xb4, 0xa4, 0xe4, 0x68, 0x4c,
	0x83, 0xe3, 0x31, 0x25, 0xc3, 0x31, 0x0d, 0x4e, 0xc6, 0x34, 0xd8, 0x7d, 0x96, 0xa2, 0x7e, 0x91,
	0x46, 0x05, 0xf6, 0x1d, 0x18, 0xc3, 0xa3, 0xdc, 0x32, 0x3f, 0x74, 0xd1, 0x64, 0xa1, 0x36, 0x58,
	0xc8, 0x0e, 0x98, 0x70, 0x0e, 0x33, 0x9d, 0xa4, 0xc8, 0x60, 0xdf, 0xcd, 0x72, 0xf8, 0xeb, 0xaf,
	0x9c, 0x5c, 0xf0, 0x05, 0xbb, 0xf5, 0x2b, 0x00, 0x00, 0xff, 0xff, 0xca, 0xab, 0x6b, 0x1c, 0xf5,
	0x03, 0x00, 0x00,
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
	err := c.cc.Invoke(ctx, "/ves.io.schema.views.gcp_vpc_site.PrivateCustomAPI/UpdateSiteError", in, out, opts...)
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
		FullMethod: "/ves.io.schema.views.gcp_vpc_site.PrivateCustomAPI/UpdateSiteError",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PrivateCustomAPIServer).UpdateSiteError(ctx, req.(*views.UpdateSiteErrorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PrivateCustomAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ves.io.schema.views.gcp_vpc_site.PrivateCustomAPI",
	HandlerType: (*PrivateCustomAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateSiteError",
			Handler:    _PrivateCustomAPI_UpdateSiteError_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ves.io/schema/views/gcp_vpc_site/private_customapi.proto",
}