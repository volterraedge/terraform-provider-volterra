// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ves.io/schema/views/aws_tgw_site/private_customapi.proto

package aws_tgw_site

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
	proto.RegisterFile("ves.io/schema/views/aws_tgw_site/private_customapi.proto", fileDescriptor_1b3948f72e872ec9)
}
func init() {
	golang_proto.RegisterFile("ves.io/schema/views/aws_tgw_site/private_customapi.proto", fileDescriptor_1b3948f72e872ec9)
}

var fileDescriptor_1b3948f72e872ec9 = []byte{
	// 509 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x52, 0x41, 0x6b, 0x13, 0x41,
	0x14, 0xde, 0x89, 0xe2, 0x21, 0x17, 0x4b, 0x4e, 0x35, 0x96, 0xb1, 0x04, 0x51, 0xd0, 0xee, 0x0e,
	0x28, 0x05, 0xf1, 0xa6, 0xe2, 0x41, 0x10, 0x2c, 0x8a, 0x1e, 0x7a, 0x59, 0x66, 0xb7, 0x2f, 0x93,
	0xd1, 0xec, 0xbe, 0x71, 0xe6, 0xed, 0xa6, 0x22, 0x05, 0xe9, 0x2f, 0x28, 0xf8, 0x07, 0x3c, 0xfa,
	0x1f, 0x72, 0xe9, 0x4d, 0x4f, 0x12, 0xd4, 0x43, 0x8f, 0x66, 0xa3, 0xe0, 0xb1, 0x3f, 0x41, 0x32,
	0x49, 0x24, 0x89, 0x95, 0x7a, 0xf0, 0xf6, 0x66, 0xbf, 0xef, 0x7b, 0xfb, 0xbd, 0xf7, 0xbe, 0xfa,
	0xad, 0x12, 0x5c, 0xa4, 0x51, 0xb8, 0xb4, 0x03, 0x99, 0x14, 0xa5, 0x86, 0x9e, 0x13, 0xb2, 0xe7,
	0x62, 0x52, 0xbd, 0xd8, 0x69, 0x02, 0x61, 0xac, 0x2e, 0x25, 0x41, 0x9c, 0x16, 0x8e, 0x30, 0x93,
	0x46, 0x47, 0xc6, 0x22, 0x61, 0x63, 0x7d, 0xa2, 0x8c, 0x26, 0xca, 0xc8, 0x2b, 0xa3, 0x79, 0x65,
	0x33, 0x54, 0x9a, 0x3a, 0x45, 0x12, 0xa5, 0x98, 0x09, 0x85, 0x0a, 0x85, 0x17, 0x26, 0x45, 0xdb,
	0xbf, 0xfc, 0xc3, 0x57, 0x93, 0x86, 0xcd, 0x35, 0x85, 0xa8, 0xba, 0x20, 0xa4, 0xd1, 0x42, 0xe6,
	0x39, 0x92, 0x24, 0x8d, 0xb9, 0x9b, 0xa2, 0x17, 0x17, 0x8d, 0xa2, 0x99, 0x07, 0x2f, 0x2c, 0x82,
	0xf4, 0xca, 0xc0, 0x0c, 0x5a, 0x5b, 0x1a, 0x50, 0x76, 0xf5, 0x8e, 0x24, 0x98, 0xa2, 0xad, 0x25,
	0x14, 0x1c, 0xe4, 0xe5, 0x52, 0xf3, 0xf5, 0x3f, 0x57, 0x14, 0x2f, 0x32, 0x36, 0x4e, 0x5d, 0xe2,
	0xbc, 0xa3, 0xcb, 0x27, 0xb1, 0xc7, 0xac, 0x78, 0x9e, 0x75, 0x65, 0x99, 0x65, 0xa9, 0x90, 0xdd,
	0xb8, 0x83, 0x8e, 0xe6, 0xbb, 0xdd, 0x78, 0x77, 0xa6, 0xbe, 0xb2, 0x35, 0x39, 0xd1, 0x3d, 0x7f,
	0xa1, 0x3b, 0x5b, 0x0f, 0x1a, 0x5f, 0x6b, 0xf5, 0xf3, 0x4f, 0xcd, 0x78, 0xce, 0x27, 0x9a, 0xe0,
	0xbe, 0xb5, 0x68, 0x1b, 0xd7, 0xa3, 0x93, 0x0e, 0xb6, 0xc4, 0x7a, 0x0c, 0x2f, 0x0b, 0x70, 0xd4,
	0xdc, 0xf8, 0x37, 0xb2, 0x33, 0x98, 0x3b, 0x68, 0xfd, 0x60, 0xd5, 0x87, 0xd5, 0x4b, 0xed, 0xcd,
	0xdd, 0x34, 0x1c, 0xcf, 0x11, 0x66, 0x32, 0x97, 0x0a, 0x32, 0xc8, 0x29, 0x4c, 0xa4, 0xd3, 0x69,
	0x58, 0x38, 0xb0, 0xfb, 0x5f, 0xbe, 0xbf, 0xad, 0xf5, 0x59, 0xeb, 0xe1, 0x2c, 0x51, 0x62, 0x92,
	0x28, 0x91, 0xcb, 0x0c, 0x9c, 0x91, 0x29, 0x38, 0xf1, 0xfa, 0x77, 0xbd, 0xb7, 0xb8, 0x3f, 0x0f,
	0xec, 0x89, 0xc2, 0x1b, 0xf0, 0x9f, 0x62, 0x18, 0x5b, 0xb8, 0xcd, 0xae, 0x6d, 0x77, 0x5a, 0xa9,
	0x58, 0x70, 0x2b, 0x74, 0x4e, 0x16, 0x9d, 0x81, 0x94, 0x44, 0xcf, 0xea, 0xff, 0xf6, 0xa7, 0xe6,
	0xe6, 0x61, 0x9f, 0x9d, 0xfd, 0xdc, 0x67, 0x57, 0x4f, 0x8b, 0x7e, 0xf4, 0x28, 0x79, 0x0e, 0x29,
	0xed, 0x7f, 0x5a, 0xad, 0xad, 0xb0, 0xbb, 0x07, 0x6c, 0x30, 0xe4, 0xc1, 0xd1, 0x90, 0x07, 0xc7,
	0x43, 0xce, 0xde, 0x54, 0x9c, 0xbd, 0xaf, 0x38, 0xfb, 0x58, 0x71, 0x36, 0xa8, 0x38, 0xfb, 0x56,
	0x71, 0xf6, 0xb3, 0xe2, 0xc1, 0x71, 0xc5, 0xd9, 0xc1, 0x88, 0x07, 0x87, 0x23, 0xce, 0x06, 0x23,
	0x1e, 0x1c, 0x8d, 0x78, 0xb0, 0xfd, 0x4c, 0xa1, 0x79, 0xa1, 0xa2, 0x12, 0xbb, 0x04, 0xd6, 0xca,
	0xa8, 0x70, 0xc2, 0x17, 0x6d, 0xb4, 0x59, 0x68, 0x2c, 0x96, 0x7a, 0x07, 0x6c, 0x38, 0x83, 0x85,
	0x49, 0x14, 0x0a, 0xd8, 0xa5, 0xe9, 0x1e, 0xfe, 0x9a, 0xc7, 0xe4, 0x9c, 0x0f, 0xcf, 0xcd, 0x5f,
	0x01, 0x00, 0x00, 0xff, 0xff, 0xeb, 0x14, 0x5e, 0x3c, 0xff, 0x03, 0x00, 0x00,
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
	err := c.cc.Invoke(ctx, "/ves.io.schema.views.aws_tgw_site.PrivateCustomAPI/UpdateSiteError", in, out, opts...)
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
		FullMethod: "/ves.io.schema.views.aws_tgw_site.PrivateCustomAPI/UpdateSiteError",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PrivateCustomAPIServer).UpdateSiteError(ctx, req.(*views.UpdateSiteErrorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PrivateCustomAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ves.io.schema.views.aws_tgw_site.PrivateCustomAPI",
	HandlerType: (*PrivateCustomAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateSiteError",
			Handler:    _PrivateCustomAPI_UpdateSiteError_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ves.io/schema/views/aws_tgw_site/private_customapi.proto",
}
