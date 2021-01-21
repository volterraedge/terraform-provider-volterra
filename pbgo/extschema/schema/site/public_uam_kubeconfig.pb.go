// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ves.io/schema/site/public_uam_kubeconfig.proto

package site

import proto "github.com/gogo/protobuf/proto"
import golang_proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import _ "github.com/gogo/googleapis/google/api"
import google_api2 "google.golang.org/genproto/googleapis/api/httpbody"
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

// Client API for UamKubeConfigAPI service

type UamKubeConfigAPIClient interface {
	// Get K8s Cluster Global Kube Config
	//
	// x-displayName: "Get K8s Cluster Global Kube Config"
	// Down load kube config for global k8s cluster access
	GetGlobalKubeConfig(ctx context.Context, in *GetKubeConfigReq, opts ...grpc.CallOption) (*google_api2.HttpBody, error)
	// List Global Kube Configs
	//
	// x-displayName: "List Global Kube Configs"
	// Returns list of all global active kubeconfig minted for this site
	ListGlobalKubeConfig(ctx context.Context, in *ListKubeConfigReq, opts ...grpc.CallOption) (*ListKubeConfigRsp, error)
}

type uamKubeConfigAPIClient struct {
	cc *grpc.ClientConn
}

func NewUamKubeConfigAPIClient(cc *grpc.ClientConn) UamKubeConfigAPIClient {
	return &uamKubeConfigAPIClient{cc}
}

func (c *uamKubeConfigAPIClient) GetGlobalKubeConfig(ctx context.Context, in *GetKubeConfigReq, opts ...grpc.CallOption) (*google_api2.HttpBody, error) {
	out := new(google_api2.HttpBody)
	err := grpc.Invoke(ctx, "/ves.io.schema.site.UamKubeConfigAPI/GetGlobalKubeConfig", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *uamKubeConfigAPIClient) ListGlobalKubeConfig(ctx context.Context, in *ListKubeConfigReq, opts ...grpc.CallOption) (*ListKubeConfigRsp, error) {
	out := new(ListKubeConfigRsp)
	err := grpc.Invoke(ctx, "/ves.io.schema.site.UamKubeConfigAPI/ListGlobalKubeConfig", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for UamKubeConfigAPI service

type UamKubeConfigAPIServer interface {
	// Get K8s Cluster Global Kube Config
	//
	// x-displayName: "Get K8s Cluster Global Kube Config"
	// Down load kube config for global k8s cluster access
	GetGlobalKubeConfig(context.Context, *GetKubeConfigReq) (*google_api2.HttpBody, error)
	// List Global Kube Configs
	//
	// x-displayName: "List Global Kube Configs"
	// Returns list of all global active kubeconfig minted for this site
	ListGlobalKubeConfig(context.Context, *ListKubeConfigReq) (*ListKubeConfigRsp, error)
}

func RegisterUamKubeConfigAPIServer(s *grpc.Server, srv UamKubeConfigAPIServer) {
	s.RegisterService(&_UamKubeConfigAPI_serviceDesc, srv)
}

func _UamKubeConfigAPI_GetGlobalKubeConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetKubeConfigReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UamKubeConfigAPIServer).GetGlobalKubeConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ves.io.schema.site.UamKubeConfigAPI/GetGlobalKubeConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UamKubeConfigAPIServer).GetGlobalKubeConfig(ctx, req.(*GetKubeConfigReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UamKubeConfigAPI_ListGlobalKubeConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListKubeConfigReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UamKubeConfigAPIServer).ListGlobalKubeConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ves.io.schema.site.UamKubeConfigAPI/ListGlobalKubeConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UamKubeConfigAPIServer).ListGlobalKubeConfig(ctx, req.(*ListKubeConfigReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _UamKubeConfigAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ves.io.schema.site.UamKubeConfigAPI",
	HandlerType: (*UamKubeConfigAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetGlobalKubeConfig",
			Handler:    _UamKubeConfigAPI_GetGlobalKubeConfig_Handler,
		},
		{
			MethodName: "ListGlobalKubeConfig",
			Handler:    _UamKubeConfigAPI_ListGlobalKubeConfig_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ves.io/schema/site/public_uam_kubeconfig.proto",
}

func init() {
	proto.RegisterFile("ves.io/schema/site/public_uam_kubeconfig.proto", fileDescriptorPublicUamKubeconfig)
}
func init() {
	golang_proto.RegisterFile("ves.io/schema/site/public_uam_kubeconfig.proto", fileDescriptorPublicUamKubeconfig)
}

var fileDescriptorPublicUamKubeconfig = []byte{
	// 463 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x92, 0xbf, 0x6b, 0x14, 0x41,
	0x14, 0xc7, 0x6f, 0x2e, 0x60, 0x71, 0x55, 0x58, 0x53, 0x24, 0xa7, 0x4c, 0x71, 0x28, 0x56, 0x3b,
	0x03, 0x5a, 0xfb, 0x2b, 0x16, 0x17, 0x51, 0x50, 0x04, 0x0b, 0x2d, 0x8c, 0x33, 0x7b, 0xef, 0xe6,
	0xc6, 0xec, 0xee, 0x1b, 0x77, 0x66, 0x17, 0x4f, 0x09, 0x48, 0xfe, 0x02, 0xc1, 0x7f, 0xc2, 0xc2,
	0xc6, 0x4e, 0x48, 0x93, 0x42, 0x30, 0x95, 0x04, 0x6d, 0x2c, 0xbd, 0x89, 0x85, 0x65, 0xfe, 0x04,
	0xc9, 0xdc, 0x9d, 0x9b, 0x25, 0x27, 0x08, 0x76, 0xef, 0xed, 0xf7, 0xb3, 0x8f, 0xef, 0xf7, 0xcd,
	0xeb, 0xb0, 0x0a, 0x2c, 0xd3, 0xc8, 0x6d, 0x32, 0x82, 0x4c, 0x70, 0xab, 0x1d, 0x70, 0x53, 0xca,
	0x54, 0x27, 0x9b, 0xa5, 0xc8, 0x36, 0xb7, 0x4a, 0x09, 0x09, 0xe6, 0x43, 0xad, 0x98, 0x29, 0xd0,
	0x61, 0x14, 0x4d, 0x79, 0x36, 0xe5, 0xd9, 0x31, 0xdf, 0x8d, 0x95, 0x76, 0xa3, 0x52, 0xb2, 0x04,
	0x33, 0xae, 0x50, 0x21, 0x0f, 0xa8, 0x2c, 0x87, 0xa1, 0x0b, 0x4d, 0xa8, 0xa6, 0x23, 0xba, 0xe7,
	0x15, 0xa2, 0x4a, 0x81, 0x0b, 0xa3, 0xb9, 0xc8, 0x73, 0x74, 0xc2, 0x69, 0xcc, 0xed, 0x4c, 0x5d,
	0x3b, 0xa1, 0x8e, 0x9c, 0x33, 0x12, 0x07, 0xe3, 0x99, 0x74, 0xae, 0xe9, 0x15, 0xcd, 0xc9, 0xff,
	0xe8, 0x82, 0x20, 0x6e, 0x6c, 0x60, 0xae, 0xf7, 0x9a, 0x7a, 0x05, 0x16, 0xf2, 0xaa, 0x39, 0xe3,
	0xf2, 0x87, 0xa5, 0xce, 0xf2, 0x43, 0x91, 0xdd, 0x29, 0x25, 0xdc, 0x0a, 0xa1, 0x6f, 0xde, 0xbf,
	0x1d, 0xbd, 0x27, 0x9d, 0xb3, 0x7d, 0x70, 0xfd, 0x14, 0xa5, 0x48, 0x6b, 0x29, 0xba, 0xc0, 0x4e,
	0xaf, 0x82, 0xf5, 0xc1, 0xd5, 0xc8, 0x03, 0x78, 0xde, 0x5d, 0x61, 0xd3, 0x3c, 0x4c, 0x18, 0xcd,
	0x36, 0x9c, 0x33, 0xeb, 0x38, 0x18, 0xf7, 0x1e, 0xf9, 0xcf, 0xab, 0xcb, 0x15, 0xd8, 0x58, 0x63,
	0xac, 0x85, 0xb0, 0x71, 0x22, 0x84, 0xdd, 0xf9, 0xf6, 0xf3, 0x6d, 0xfb, 0x7a, 0x74, 0x75, 0xb6,
	0x7f, 0x9e, 0x8b, 0x0c, 0xac, 0x11, 0x09, 0x58, 0xfe, 0xea, 0x4f, 0xbd, 0x1d, 0x92, 0xcd, 0xbe,
	0x6c, 0x73, 0x15, 0x5c, 0xc5, 0xf5, 0x33, 0x45, 0x9f, 0x48, 0x67, 0xe5, 0xae, 0xb6, 0xa7, 0xfd,
	0x5e, 0x5c, 0xe4, 0xf7, 0x98, 0x6c, 0x1a, 0xfe, 0x17, 0xcc, 0x9a, 0xde, 0x93, 0xfd, 0x8f, 0x6d,
	0xf2, 0xd7, 0x14, 0x37, 0xa2, 0x6b, 0xff, 0x95, 0xc2, 0x76, 0x2f, 0xed, 0xed, 0x92, 0xa5, 0xaf,
	0xbb, 0x64, 0x6d, 0x81, 0x9b, 0x7b, 0xf2, 0x19, 0x24, 0x6e, 0xe7, 0xcb, 0x6a, 0xfb, 0x29, 0x59,
	0x7f, 0x79, 0x30, 0xa1, 0xad, 0xef, 0x13, 0xda, 0x3a, 0x9a, 0x50, 0xf2, 0xda, 0x53, 0xf2, 0xce,
	0x53, 0xb2, 0xef, 0x29, 0x39, 0xf0, 0x94, 0xfc, 0xf0, 0x94, 0xfc, 0xf2, 0xb4, 0x75, 0xe4, 0x29,
	0x79, 0x73, 0x48, 0x5b, 0x7b, 0x87, 0x94, 0x3c, 0xde, 0x50, 0x68, 0xb6, 0x14, 0xab, 0x30, 0x75,
	0x50, 0x14, 0x82, 0x95, 0x96, 0x87, 0x62, 0x88, 0x45, 0x16, 0x9b, 0x02, 0x2b, 0x3d, 0x80, 0x22,
	0x9e, 0xcb, 0xdc, 0x48, 0x85, 0x1c, 0x5e, 0xb8, 0xf9, 0x61, 0xd5, 0xf7, 0x25, 0xcf, 0x84, 0xb3,
	0xb9, 0xf2, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x7c, 0xd2, 0xf9, 0x44, 0x45, 0x03, 0x00, 0x00,
}
