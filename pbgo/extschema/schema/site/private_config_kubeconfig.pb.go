// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ves.io/schema/site/private_config_kubeconfig.proto

package site

import proto "github.com/gogo/protobuf/proto"
import golang_proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import _ "github.com/gogo/googleapis/google/api"
import _ "google.golang.org/genproto/googleapis/api/httpbody"
import _ "github.com/gogo/protobuf/types"
import _ "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema"
import _ "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/vesenv"

import context "golang.org/x/net/context"
import grpc "google.golang.org/grpc"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = golang_proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for PrivateConfigKubeConfigAPI service

type PrivateConfigKubeConfigAPIClient interface {
	// Check Global Accesss Enabled
	//
	// x-displayName: "Global Access Check"
	// API to check global access is enabled or not.
	GlobalAccessEnabled(ctx context.Context, in *GlobalAccessCheckRequest, opts ...grpc.CallOption) (*GlobalAccessCheckResponse, error)
}

type privateConfigKubeConfigAPIClient struct {
	cc *grpc.ClientConn
}

func NewPrivateConfigKubeConfigAPIClient(cc *grpc.ClientConn) PrivateConfigKubeConfigAPIClient {
	return &privateConfigKubeConfigAPIClient{cc}
}

func (c *privateConfigKubeConfigAPIClient) GlobalAccessEnabled(ctx context.Context, in *GlobalAccessCheckRequest, opts ...grpc.CallOption) (*GlobalAccessCheckResponse, error) {
	out := new(GlobalAccessCheckResponse)
	err := grpc.Invoke(ctx, "/ves.io.schema.site.PrivateConfigKubeConfigAPI/GlobalAccessEnabled", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for PrivateConfigKubeConfigAPI service

type PrivateConfigKubeConfigAPIServer interface {
	// Check Global Accesss Enabled
	//
	// x-displayName: "Global Access Check"
	// API to check global access is enabled or not.
	GlobalAccessEnabled(context.Context, *GlobalAccessCheckRequest) (*GlobalAccessCheckResponse, error)
}

func RegisterPrivateConfigKubeConfigAPIServer(s *grpc.Server, srv PrivateConfigKubeConfigAPIServer) {
	s.RegisterService(&_PrivateConfigKubeConfigAPI_serviceDesc, srv)
}

func _PrivateConfigKubeConfigAPI_GlobalAccessEnabled_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GlobalAccessCheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PrivateConfigKubeConfigAPIServer).GlobalAccessEnabled(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ves.io.schema.site.PrivateConfigKubeConfigAPI/GlobalAccessEnabled",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PrivateConfigKubeConfigAPIServer).GlobalAccessEnabled(ctx, req.(*GlobalAccessCheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PrivateConfigKubeConfigAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ves.io.schema.site.PrivateConfigKubeConfigAPI",
	HandlerType: (*PrivateConfigKubeConfigAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GlobalAccessEnabled",
			Handler:    _PrivateConfigKubeConfigAPI_GlobalAccessEnabled_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ves.io/schema/site/private_config_kubeconfig.proto",
}

func init() {
	proto.RegisterFile("ves.io/schema/site/private_config_kubeconfig.proto", fileDescriptorPrivateConfigKubeconfig)
}
func init() {
	golang_proto.RegisterFile("ves.io/schema/site/private_config_kubeconfig.proto", fileDescriptorPrivateConfigKubeconfig)
}

var fileDescriptorPrivateConfigKubeconfig = []byte{
	// 475 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x52, 0x31, 0x6b, 0x14, 0x41,
	0x18, 0xbd, 0x39, 0xc5, 0xe2, 0x2a, 0x59, 0x11, 0x92, 0x55, 0x46, 0x48, 0x63, 0xe3, 0xee, 0x40,
	0xfc, 0x03, 0xc6, 0x20, 0x2a, 0x16, 0x06, 0xcb, 0x14, 0x9e, 0x33, 0x73, 0xdf, 0xcd, 0x8d, 0xb7,
	0xbb, 0xdf, 0x38, 0x33, 0xbb, 0x18, 0x25, 0x20, 0x01, 0x7b, 0xc1, 0xbf, 0x60, 0xe1, 0x4f, 0x10,
	0x82, 0x90, 0xce, 0xab, 0x24, 0x68, 0x63, 0xe9, 0x6d, 0x2c, 0x2c, 0xf3, 0x13, 0xe4, 0x66, 0xf7,
	0xe2, 0x1d, 0x5e, 0x21, 0xa4, 0x59, 0xde, 0xec, 0x7b, 0xdf, 0x7b, 0x3b, 0x6f, 0xbf, 0xde, 0x66,
	0x05, 0x2e, 0xd5, 0xc8, 0x9c, 0x1c, 0x41, 0xce, 0x99, 0xd3, 0x1e, 0x98, 0xb1, 0xba, 0xe2, 0x1e,
	0xfa, 0x12, 0x8b, 0xa1, 0x56, 0xfd, 0x71, 0x29, 0xa0, 0x81, 0xa9, 0xb1, 0xe8, 0x31, 0x8a, 0x9a,
	0x99, 0xb4, 0x99, 0x49, 0x67, 0x33, 0x71, 0xa2, 0xb4, 0x1f, 0x95, 0x22, 0x95, 0x98, 0x33, 0x85,
	0x0a, 0x59, 0x90, 0x8a, 0x72, 0x18, 0x4e, 0xe1, 0x10, 0x50, 0x63, 0x11, 0x5f, 0x57, 0x88, 0x2a,
	0x03, 0xc6, 0x8d, 0x66, 0xbc, 0x28, 0xd0, 0x73, 0xaf, 0xb1, 0x70, 0x2d, 0xbb, 0xbe, 0xc0, 0x8e,
	0xbc, 0x37, 0x02, 0x07, 0x7b, 0x2d, 0x75, 0xa3, 0xa5, 0xce, 0xec, 0xbd, 0xce, 0xc1, 0x79, 0x9e,
	0x9b, 0x56, 0x70, 0x6d, 0xf9, 0x42, 0x68, 0x16, 0x8d, 0xe9, 0x8a, 0xdb, 0xfa, 0x3d, 0x03, 0x73,
	0x7e, 0x63, 0x99, 0xaf, 0xc0, 0x41, 0x51, 0x2d, 0x7b, 0x6c, 0xbe, 0xbd, 0xd0, 0x8b, 0x77, 0x9a,
	0x86, 0xb6, 0x43, 0x2b, 0x8f, 0x4a, 0xd1, 0xa2, 0xad, 0x9d, 0x87, 0xd1, 0xe7, 0x6e, 0xef, 0xca,
	0xfd, 0x0c, 0x05, 0xcf, 0xb6, 0xa4, 0x04, 0xe7, 0xee, 0x15, 0x5c, 0x64, 0x30, 0x88, 0x6e, 0xa5,
	0xff, 0xb6, 0x96, 0x2e, 0x0a, 0xb7, 0x47, 0x20, 0xc7, 0x4f, 0xe0, 0x45, 0x09, 0xce, 0xc7, 0xc9,
	0x7f, 0xaa, 0x9d, 0xc1, 0xc2, 0xc1, 0xc6, 0x84, 0x4c, 0x3e, 0x75, 0x49, 0xfd, 0x65, 0xed, 0x6a,
	0x05, 0x2e, 0xd1, 0x98, 0x68, 0xce, 0x5d, 0x22, 0x67, 0x0f, 0x0b, 0x7c, 0x70, 0xf0, 0xfd, 0xd7,
	0xfb, 0xee, 0x07, 0x12, 0xdd, 0x99, 0xff, 0x5a, 0x56, 0xf0, 0x1c, 0x9c, 0xe1, 0x12, 0x1c, 0x7b,
	0x7d, 0x86, 0xf7, 0x43, 0x21, 0xed, 0x9b, 0x7d, 0xa6, 0x42, 0x5c, 0x9f, 0x87, 0xbc, 0xbe, 0x9c,
	0x05, 0xee, 0x3e, 0x8b, 0x9e, 0xb2, 0xa5, 0x2f, 0x63, 0xba, 0xf0, 0x16, 0x9d, 0x01, 0xe9, 0xd9,
	0x2c, 0xea, 0xbc, 0x09, 0xf1, 0xcd, 0xa3, 0x43, 0x72, 0xf1, 0xdb, 0x21, 0x59, 0x5f, 0x51, 0xc0,
	0x63, 0xf1, 0x1c, 0xa4, 0x3f, 0xf8, 0xba, 0xd6, 0xbd, 0x4c, 0xee, 0xbe, 0x3a, 0x9e, 0xd2, 0xce,
	0x8f, 0x29, 0xed, 0x9c, 0x4e, 0x29, 0x79, 0x53, 0x53, 0xf2, 0xb1, 0xa6, 0x64, 0x52, 0x53, 0x72,
	0x5c, 0x53, 0xf2, 0xb3, 0xa6, 0xe4, 0x77, 0x4d, 0x3b, 0xa7, 0x35, 0x25, 0xef, 0x4e, 0x68, 0xe7,
	0xe8, 0x84, 0x92, 0xdd, 0x07, 0x0a, 0xcd, 0x58, 0xa5, 0x15, 0x66, 0x1e, 0xac, 0xe5, 0x69, 0xe9,
	0x58, 0x00, 0x43, 0xb4, 0x79, 0x62, 0x2c, 0x56, 0x7a, 0x00, 0x36, 0x99, 0xd3, 0xcc, 0x08, 0x85,
	0x0c, 0x5e, 0xfa, 0xf9, 0xb2, 0xfc, 0xdd, 0x19, 0x71, 0x29, 0xac, 0xc2, 0xed, 0x3f, 0x01, 0x00,
	0x00, 0xff, 0xff, 0x65, 0xb4, 0xbf, 0x85, 0x3e, 0x03, 0x00, 0x00,
}
