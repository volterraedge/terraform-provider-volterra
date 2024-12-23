// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ves.io/schema/views/aws_vpc_site/private_customapi.proto

package aws_vpc_site

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
	proto.RegisterFile("ves.io/schema/views/aws_vpc_site/private_customapi.proto", fileDescriptor_1dc7fe64d1504e69)
}
func init() {
	golang_proto.RegisterFile("ves.io/schema/views/aws_vpc_site/private_customapi.proto", fileDescriptor_1dc7fe64d1504e69)
}

var fileDescriptor_1dc7fe64d1504e69 = []byte{
	// 506 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x53, 0x31, 0x8b, 0x13, 0x41,
	0x18, 0xdd, 0x89, 0x62, 0x91, 0xc6, 0x23, 0xd5, 0x19, 0x8f, 0xe1, 0x08, 0xa2, 0xa0, 0xee, 0x0e,
	0x28, 0x07, 0x62, 0xa7, 0x62, 0x21, 0x08, 0x1e, 0x8a, 0x16, 0xd7, 0x2c, 0x93, 0xcd, 0x97, 0xcd,
	0x68, 0x76, 0xbe, 0x71, 0xe6, 0xdb, 0xcd, 0x89, 0x1c, 0xc8, 0xfd, 0x82, 0x03, 0x7f, 0x82, 0x8d,
	0xff, 0x21, 0xcd, 0x75, 0x5a, 0x49, 0xd0, 0xc2, 0x2b, 0xcd, 0xc6, 0xe2, 0xca, 0xfb, 0x09, 0x92,
	0x49, 0x22, 0xc9, 0x7a, 0xa2, 0xc5, 0x75, 0xdf, 0xec, 0x7b, 0x6f, 0xf6, 0x7d, 0x6f, 0xdf, 0xd6,
	0xef, 0x14, 0xe0, 0x22, 0x85, 0xc2, 0x25, 0x3d, 0xc8, 0xa4, 0x28, 0x14, 0x0c, 0x9c, 0x90, 0x03,
	0x17, 0x17, 0x26, 0x89, 0x9d, 0x22, 0x10, 0xc6, 0xaa, 0x42, 0x12, 0xc4, 0x49, 0xee, 0x08, 0x33,
	0x69, 0x54, 0x64, 0x2c, 0x12, 0x36, 0x36, 0x67, 0xca, 0x68, 0xa6, 0x8c, 0xbc, 0x32, 0x5a, 0x56,
	0x36, 0xc3, 0x54, 0x51, 0x2f, 0x6f, 0x47, 0x09, 0x66, 0x22, 0xc5, 0x14, 0x85, 0x17, 0xb6, 0xf3,
	0xae, 0x3f, 0xf9, 0x83, 0x9f, 0x66, 0x17, 0x36, 0x37, 0x52, 0xc4, 0xb4, 0x0f, 0x42, 0x1a, 0x25,
	0xa4, 0xd6, 0x48, 0x92, 0x14, 0x6a, 0x37, 0x47, 0x2f, 0xaf, 0x1a, 0x45, 0xb3, 0x0c, 0x5e, 0x5a,
	0x05, 0xe9, 0x8d, 0x81, 0x05, 0xb4, 0x51, 0x59, 0x50, 0xf6, 0x55, 0x47, 0x12, 0xcc, 0xd1, 0x56,
	0x05, 0x05, 0x07, 0xba, 0xa8, 0x5c, 0xbe, 0xf9, 0x67, 0x44, 0xf1, 0x2a, 0xe3, 0xca, 0x69, 0x21,
	0x4e, 0x23, 0x88, 0x97, 0x9d, 0x5c, 0xad, 0xb2, 0x2c, 0xe5, 0xb2, 0x1f, 0xf7, 0xd0, 0xd1, 0xb2,
	0xe3, 0x5b, 0x1f, 0xce, 0xd5, 0xd7, 0xb6, 0x67, 0xa1, 0x3f, 0xf0, 0x99, 0xdf, 0xdb, 0x7e, 0xd4,
	0xf8, 0x5e, 0xab, 0x5f, 0x7c, 0x6e, 0xa6, 0xce, 0x9f, 0x29, 0x82, 0x87, 0xd6, 0xa2, 0x6d, 0xdc,
	0x88, 0x4e, 0xfb, 0x04, 0x15, 0xd6, 0x53, 0x78, 0x9d, 0x83, 0xa3, 0xe6, 0xcd, 0xff, 0x23, 0x3b,
	0x83, 0xda, 0x41, 0xeb, 0x98, 0x95, 0x9f, 0xd6, 0x5b, 0xdd, 0xad, 0xdd, 0x24, 0x9c, 0xee, 0x11,
	0x66, 0x52, 0xcb, 0x14, 0x32, 0xd0, 0x14, 0x3a, 0x92, 0xba, 0x23, 0x6d, 0x27, 0xcc, 0x1d, 0xd8,
	0xfd, 0x6f, 0x3f, 0xdf, 0xd7, 0x86, 0xac, 0xf5, 0x78, 0x51, 0x13, 0x31, 0xab, 0x89, 0xd0, 0x32,
	0x03, 0x67, 0x64, 0x02, 0x4e, 0xbc, 0xfd, 0x3d, 0xef, 0xad, 0x36, 0xcb, 0x03, 0x7b, 0x22, 0xf7,
	0x1e, 0xfc, 0xa3, 0x18, 0xa6, 0x2e, 0xee, 0xb2, 0xeb, 0x3b, 0xbd, 0x56, 0x22, 0x56, 0x0c, 0x0b,
	0xa5, 0xc9, 0xa2, 0x33, 0x90, 0x90, 0x18, 0x58, 0x75, 0x66, 0x6f, 0x6a, 0x6e, 0x1d, 0x0e, 0xd9,
	0xf9, 0xaf, 0x43, 0x76, 0xed, 0x5f, 0x7d, 0x8e, 0x9e, 0xb4, 0x5f, 0x42, 0x42, 0xfb, 0x5f, 0xd6,
	0x6b, 0x6b, 0xec, 0xfe, 0x01, 0x1b, 0x8d, 0x79, 0x70, 0x34, 0xe6, 0xc1, 0xc9, 0x98, 0xb3, 0x77,
	0x25, 0x67, 0x1f, 0x4b, 0xce, 0x3e, 0x97, 0x9c, 0x8d, 0x4a, 0xce, 0x7e, 0x94, 0x9c, 0x1d, 0x97,
	0x3c, 0x38, 0x29, 0x39, 0x3b, 0x98, 0xf0, 0xe0, 0x70, 0xc2, 0xd9, 0x68, 0xc2, 0x83, 0xa3, 0x09,
	0x0f, 0x76, 0x5e, 0xa4, 0x68, 0x5e, 0xa5, 0x51, 0x81, 0x7d, 0x02, 0x6b, 0x65, 0x94, 0x3b, 0xe1,
	0x87, 0x2e, 0xda, 0x2c, 0x34, 0x16, 0x0b, 0xd5, 0x01, 0x1b, 0x2e, 0x60, 0x61, 0xda, 0x29, 0x0a,
	0xd8, 0xa5, 0x79, 0x0e, 0x7f, 0xfd, 0x53, 0xdb, 0x17, 0x7c, 0x7f, 0x6e, 0xff, 0x0a, 0x00, 0x00,
	0xff, 0xff, 0xb0, 0x5d, 0x1f, 0x68, 0xd4, 0x03, 0x00, 0x00,
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
	err := c.cc.Invoke(ctx, "/ves.io.schema.views.aws_vpc_site.PrivateCustomAPI/UpdateSiteError", in, out, opts...)
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
		FullMethod: "/ves.io.schema.views.aws_vpc_site.PrivateCustomAPI/UpdateSiteError",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PrivateCustomAPIServer).UpdateSiteError(ctx, req.(*views.UpdateSiteErrorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PrivateCustomAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ves.io.schema.views.aws_vpc_site.PrivateCustomAPI",
	HandlerType: (*PrivateCustomAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateSiteError",
			Handler:    _PrivateCustomAPI_UpdateSiteError_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ves.io/schema/views/aws_vpc_site/private_customapi.proto",
}
