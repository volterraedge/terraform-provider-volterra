// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ves.io/schema/views/terraform_parameters/public_customapi.proto

// View Terraform Parameters
//
// x-displayName: "View Terraform Parameters"
// View Terraform Parameters is set of parameters that are used by terraform scripts
// to instantiate view objects external to volterra

package terraform_parameters

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
	_ "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/views"
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

// Get Terraform Parameters for a View
//
// x-displayName: "Get Terraform Parameters for view"
// returned from list of terraform parameter objects for a given view.
type GetRequest struct {
	// Namespace
	//
	// x-displayName: "Namespace"
	// x-example: "value"
	// Namespace for the label to be retrieved
	Namespace string `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	// Kind of View
	//
	// x-displayName: "Kind of View"
	// x-example: "value"
	// Kind of view of which terraform parameters are requested e.g. aws_vpc_site, azure_vnet_site
	ViewKind string `protobuf:"bytes,2,opt,name=view_kind,json=viewKind,proto3" json:"view_kind,omitempty"`
	// Name of view
	//
	// x-displayName: "Name of view"
	// x-example: "value"
	// Name of the view for which terraform parameters are requested
	ViewName string `protobuf:"bytes,3,opt,name=view_name,json=viewName,proto3" json:"view_name,omitempty"`
}

func (m *GetRequest) Reset()      { *m = GetRequest{} }
func (*GetRequest) ProtoMessage() {}
func (*GetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7c13d7ef60096627, []int{0}
}
func (m *GetRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRequest.Merge(m, src)
}
func (m *GetRequest) XXX_Size() int {
	return m.Size()
}
func (m *GetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRequest proto.InternalMessageInfo

func (m *GetRequest) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *GetRequest) GetViewKind() string {
	if m != nil {
		return m.ViewKind
	}
	return ""
}

func (m *GetRequest) GetViewName() string {
	if m != nil {
		return m.ViewName
	}
	return ""
}

// GetResponse
//
// x-displayName: "Get Response"
// Response for Get API
type GetResponse struct {
	// Terraform Parameters
	//
	// x-displayName: "Terraform Parameters"
	// Terraform Parameters details
	TerraformParameters *GlobalSpecType `protobuf:"bytes,1,opt,name=terraform_parameters,json=terraformParameters,proto3" json:"terraform_parameters,omitempty"`
}

func (m *GetResponse) Reset()      { *m = GetResponse{} }
func (*GetResponse) ProtoMessage() {}
func (*GetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7c13d7ef60096627, []int{1}
}
func (m *GetResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetResponse.Merge(m, src)
}
func (m *GetResponse) XXX_Size() int {
	return m.Size()
}
func (m *GetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetResponse proto.InternalMessageInfo

func (m *GetResponse) GetTerraformParameters() *GlobalSpecType {
	if m != nil {
		return m.TerraformParameters
	}
	return nil
}

// GetStatusResponse
//
// x-displayName: "Get Status Response"
// Response for GetStatus API
type GetStatusResponse struct {
	// Status Object
	//
	// x-displayName: "Terraform Status Object"
	// Status Object of Terraform for a view
	Status *StatusObject `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (m *GetStatusResponse) Reset()      { *m = GetStatusResponse{} }
func (*GetStatusResponse) ProtoMessage() {}
func (*GetStatusResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7c13d7ef60096627, []int{2}
}
func (m *GetStatusResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetStatusResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetStatusResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetStatusResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetStatusResponse.Merge(m, src)
}
func (m *GetStatusResponse) XXX_Size() int {
	return m.Size()
}
func (m *GetStatusResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetStatusResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetStatusResponse proto.InternalMessageInfo

func (m *GetStatusResponse) GetStatus() *StatusObject {
	if m != nil {
		return m.Status
	}
	return nil
}

func init() {
	proto.RegisterType((*GetRequest)(nil), "ves.io.schema.views.terraform_parameters.GetRequest")
	golang_proto.RegisterType((*GetRequest)(nil), "ves.io.schema.views.terraform_parameters.GetRequest")
	proto.RegisterType((*GetResponse)(nil), "ves.io.schema.views.terraform_parameters.GetResponse")
	golang_proto.RegisterType((*GetResponse)(nil), "ves.io.schema.views.terraform_parameters.GetResponse")
	proto.RegisterType((*GetStatusResponse)(nil), "ves.io.schema.views.terraform_parameters.GetStatusResponse")
	golang_proto.RegisterType((*GetStatusResponse)(nil), "ves.io.schema.views.terraform_parameters.GetStatusResponse")
}

func init() {
	proto.RegisterFile("ves.io/schema/views/terraform_parameters/public_customapi.proto", fileDescriptor_7c13d7ef60096627)
}
func init() {
	golang_proto.RegisterFile("ves.io/schema/views/terraform_parameters/public_customapi.proto", fileDescriptor_7c13d7ef60096627)
}

var fileDescriptor_7c13d7ef60096627 = []byte{
	// 650 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0x3f, 0x4f, 0x14, 0x41,
	0x14, 0xbf, 0xb9, 0x8b, 0xc4, 0x5b, 0x1a, 0x5d, 0x2d, 0xce, 0x83, 0x8c, 0xe4, 0x2a, 0x0a, 0x77,
	0x27, 0x41, 0x30, 0x26, 0xc6, 0xa8, 0x58, 0x5c, 0x0c, 0x51, 0xc9, 0x41, 0x65, 0x21, 0x99, 0xdb,
	0x7d, 0xb7, 0x8c, 0xdc, 0xee, 0x8c, 0x33, 0xb3, 0x27, 0x48, 0x2e, 0x31, 0xf8, 0x05, 0x48, 0x6c,
	0x4c, 0xec, 0xb4, 0xf1, 0x23, 0x98, 0xd0, 0xd0, 0x49, 0x65, 0x88, 0x36, 0x94, 0xb2, 0x67, 0x61,
	0xc9, 0x47, 0x30, 0xcc, 0x2e, 0x0b, 0x87, 0x97, 0x78, 0x18, 0xba, 0x99, 0xf9, 0xbd, 0xf7, 0x7b,
	0xff, 0x7e, 0xf3, 0xac, 0x7b, 0x1d, 0x50, 0x2e, 0xe3, 0x44, 0x79, 0xcb, 0x10, 0x52, 0xd2, 0x61,
	0xf0, 0x4a, 0x11, 0x0d, 0x52, 0xd2, 0x16, 0x97, 0xe1, 0x92, 0xa0, 0x92, 0x86, 0xa0, 0x41, 0x2a,
	0x22, 0xe2, 0x66, 0x9b, 0x79, 0x4b, 0x5e, 0xac, 0x34, 0x0f, 0xa9, 0x60, 0xae, 0x90, 0x5c, 0x73,
	0x7b, 0x32, 0x25, 0x70, 0x53, 0x02, 0xd7, 0x10, 0xb8, 0x83, 0x08, 0xaa, 0x4e, 0xc0, 0xf4, 0x72,
	0xdc, 0x74, 0x3d, 0x1e, 0x92, 0x80, 0x07, 0x9c, 0x18, 0x82, 0x66, 0xdc, 0x32, 0x37, 0x73, 0x31,
	0xa7, 0x94, 0xb8, 0x3a, 0x1e, 0x70, 0x1e, 0xb4, 0x81, 0x50, 0xc1, 0x08, 0x8d, 0x22, 0xae, 0xa9,
	0x66, 0x3c, 0x52, 0x19, 0x3a, 0xd6, 0x9f, 0x37, 0x17, 0x27, 0xc1, 0x6b, 0xfd, 0xa0, 0x5e, 0x13,
	0x70, 0x04, 0x8d, 0x9f, 0xaa, 0x97, 0xb6, 0x99, 0x4f, 0x35, 0x64, 0x68, 0xed, 0x14, 0x0a, 0x0a,
	0xa2, 0xce, 0x29, 0xf2, 0x89, 0xbf, 0x3b, 0xb6, 0xd4, 0x6f, 0x31, 0x33, 0x74, 0x4f, 0x79, 0xf3,
	0x05, 0x78, 0x3a, 0x73, 0x9b, 0x1e, 0xda, 0xed, 0x64, 0x41, 0xd7, 0x07, 0x7a, 0x1d, 0x1b, 0xd4,
	0x7c, 0xcb, 0xaa, 0x83, 0x6e, 0xc0, 0xcb, 0x18, 0x94, 0xb6, 0xc7, 0xad, 0x72, 0x44, 0x43, 0x50,
	0x82, 0x7a, 0x50, 0x41, 0x13, 0x68, 0xb2, 0xdc, 0x38, 0x7e, 0xb0, 0xc7, 0xac, 0xb2, 0xa9, 0x67,
	0x85, 0x45, 0x7e, 0xa5, 0x68, 0xd0, 0x8b, 0x87, 0x0f, 0x73, 0x2c, 0xf2, 0x73, 0xf0, 0xd0, 0xbc,
	0x52, 0x3a, 0x06, 0x9f, 0xd0, 0x10, 0x6a, 0xaf, 0xad, 0x51, 0x13, 0x45, 0x09, 0x1e, 0x29, 0xb0,
	0x57, 0xac, 0xab, 0x83, 0x32, 0x37, 0x11, 0x47, 0xa7, 0x6e, 0xbb, 0xc3, 0x8a, 0xc6, 0xad, 0xb7,
	0x79, 0x93, 0xb6, 0x17, 0x04, 0x78, 0x8b, 0x6b, 0x02, 0x1a, 0x57, 0x72, 0xa3, 0xf9, 0xdc, 0xa6,
	0xb6, 0x6c, 0x5d, 0xae, 0x83, 0x5e, 0xd0, 0x54, 0xc7, 0x2a, 0xcf, 0x60, 0xc1, 0x1a, 0x51, 0xe6,
	0x25, 0x8b, 0x79, 0x6b, 0xf8, 0x98, 0x29, 0xd3, 0x53, 0x33, 0x9b, 0xd9, 0xd2, 0x76, 0x17, 0x35,
	0x32, 0xaa, 0xa9, 0xb7, 0x17, 0xac, 0xf2, 0x43, 0xf3, 0x01, 0x1e, 0xcc, 0x3f, 0xb2, 0x37, 0x8b,
	0x56, 0xa9, 0x0e, 0xda, 0x9e, 0x3e, 0x43, 0x39, 0xf9, 0x24, 0xaa, 0x33, 0x67, 0xf4, 0x4a, 0xeb,
	0xaa, 0x7d, 0x40, 0x3b, 0x5f, 0x8a, 0x28, 0xf9, 0x5a, 0xb9, 0xdf, 0x9a, 0x59, 0xf5, 0x1c, 0xc5,
	0x34, 0x38, 0x21, 0x8d, 0x68, 0x00, 0x21, 0x44, 0xda, 0x51, 0x9a, 0x46, 0x3e, 0x95, 0xbe, 0x13,
	0xf2, 0x88, 0x69, 0x2e, 0x6f, 0x74, 0x40, 0x39, 0x8c, 0x3b, 0x2c, 0x6a, 0x49, 0xaa, 0xb4, 0x8c,
	0x3d, 0x1d, 0x4b, 0x70, 0x24, 0x50, 0x7f, 0xe3, 0xc7, 0xaf, 0x77, 0xc5, 0xc7, 0xf6, 0x5c, 0xf6,
	0xb3, 0x49, 0x2e, 0x03, 0x45, 0xd6, 0xf3, 0x73, 0x77, 0xb0, 0x02, 0xd7, 0x73, 0x99, 0x74, 0xb3,
	0xf3, 0xa1, 0x47, 0xd7, 0xfe, 0x54, 0xb4, 0xca, 0xf9, 0x2c, 0xfe, 0xb3, 0x31, 0x77, 0xce, 0xe4,
	0xd5, 0x3f, 0xf6, 0xda, 0xc7, 0xf3, 0x6d, 0xcf, 0xa2, 0xdd, 0x38, 0xc7, 0xf6, 0x90, 0x54, 0x46,
	0xd5, 0xbb, 0xdb, 0x5b, 0xa8, 0xf4, 0x7d, 0x0b, 0x91, 0xa1, 0x0b, 0x4d, 0xc5, 0xb8, 0xf1, 0xad,
	0x52, 0xbc, 0x84, 0x66, 0xdf, 0xa3, 0xdd, 0x7d, 0x5c, 0xd8, 0xdb, 0xc7, 0x85, 0x83, 0x7d, 0x8c,
	0xde, 0x24, 0x18, 0x7d, 0x4e, 0x30, 0xda, 0x49, 0x30, 0xda, 0x4d, 0x30, 0xfa, 0x99, 0x60, 0xf4,
	0x3b, 0xc1, 0x85, 0x83, 0x04, 0xa3, 0xcd, 0x1e, 0x2e, 0x6c, 0xf7, 0x30, 0xda, 0xed, 0xe1, 0xc2,
	0x5e, 0x0f, 0x17, 0x9e, 0x3d, 0x0f, 0xb8, 0x58, 0x09, 0xdc, 0x0e, 0x6f, 0x9b, 0x30, 0x6e, 0x7c,
	0x62, 0xc3, 0x38, 0x42, 0xf2, 0x0e, 0xf3, 0x41, 0x3a, 0x47, 0x30, 0x11, 0xcd, 0x80, 0x13, 0x58,
	0xd5, 0xd9, 0x72, 0xf9, 0xe7, 0x66, 0x6a, 0x8e, 0x98, 0x9d, 0x73, 0xf3, 0x4f, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x49, 0xb9, 0x14, 0x9b, 0x57, 0x06, 0x00, 0x00,
}

func (this *GetRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*GetRequest)
	if !ok {
		that2, ok := that.(GetRequest)
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
	if this.ViewKind != that1.ViewKind {
		return false
	}
	if this.ViewName != that1.ViewName {
		return false
	}
	return true
}
func (this *GetResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*GetResponse)
	if !ok {
		that2, ok := that.(GetResponse)
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
	if !this.TerraformParameters.Equal(that1.TerraformParameters) {
		return false
	}
	return true
}
func (this *GetStatusResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*GetStatusResponse)
	if !ok {
		that2, ok := that.(GetStatusResponse)
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
	if !this.Status.Equal(that1.Status) {
		return false
	}
	return true
}
func (this *GetRequest) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 7)
	s = append(s, "&terraform_parameters.GetRequest{")
	s = append(s, "Namespace: "+fmt.Sprintf("%#v", this.Namespace)+",\n")
	s = append(s, "ViewKind: "+fmt.Sprintf("%#v", this.ViewKind)+",\n")
	s = append(s, "ViewName: "+fmt.Sprintf("%#v", this.ViewName)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *GetResponse) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&terraform_parameters.GetResponse{")
	if this.TerraformParameters != nil {
		s = append(s, "TerraformParameters: "+fmt.Sprintf("%#v", this.TerraformParameters)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *GetStatusResponse) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&terraform_parameters.GetStatusResponse{")
	if this.Status != nil {
		s = append(s, "Status: "+fmt.Sprintf("%#v", this.Status)+",\n")
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
	// Get Terraform Parameters for a View
	//
	// x-displayName: "Get Terraform Parameters for view"
	// returned from list of terraform parameter objects for a given view.
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
	// Get Status of Terraform for a View
	//
	// x-displayName: "Get Status of Terraform for view"
	// returned from list of terraform parameter status objects for a given view.
	GetStatus(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetStatusResponse, error)
}

type customAPIClient struct {
	cc *grpc.ClientConn
}

func NewCustomAPIClient(cc *grpc.ClientConn) CustomAPIClient {
	return &customAPIClient{cc}
}

func (c *customAPIClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, "/ves.io.schema.views.terraform_parameters.CustomAPI/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customAPIClient) GetStatus(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetStatusResponse, error) {
	out := new(GetStatusResponse)
	err := c.cc.Invoke(ctx, "/ves.io.schema.views.terraform_parameters.CustomAPI/GetStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CustomAPIServer is the server API for CustomAPI service.
type CustomAPIServer interface {
	// Get Terraform Parameters for a View
	//
	// x-displayName: "Get Terraform Parameters for view"
	// returned from list of terraform parameter objects for a given view.
	Get(context.Context, *GetRequest) (*GetResponse, error)
	// Get Status of Terraform for a View
	//
	// x-displayName: "Get Status of Terraform for view"
	// returned from list of terraform parameter status objects for a given view.
	GetStatus(context.Context, *GetRequest) (*GetStatusResponse, error)
}

// UnimplementedCustomAPIServer can be embedded to have forward compatible implementations.
type UnimplementedCustomAPIServer struct {
}

func (*UnimplementedCustomAPIServer) Get(ctx context.Context, req *GetRequest) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedCustomAPIServer) GetStatus(ctx context.Context, req *GetRequest) (*GetStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStatus not implemented")
}

func RegisterCustomAPIServer(s *grpc.Server, srv CustomAPIServer) {
	s.RegisterService(&_CustomAPI_serviceDesc, srv)
}

func _CustomAPI_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomAPIServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ves.io.schema.views.terraform_parameters.CustomAPI/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomAPIServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CustomAPI_GetStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomAPIServer).GetStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ves.io.schema.views.terraform_parameters.CustomAPI/GetStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomAPIServer).GetStatus(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CustomAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ves.io.schema.views.terraform_parameters.CustomAPI",
	HandlerType: (*CustomAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _CustomAPI_Get_Handler,
		},
		{
			MethodName: "GetStatus",
			Handler:    _CustomAPI_GetStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ves.io/schema/views/terraform_parameters/public_customapi.proto",
}

func (m *GetRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ViewName) > 0 {
		i -= len(m.ViewName)
		copy(dAtA[i:], m.ViewName)
		i = encodeVarintPublicCustomapi(dAtA, i, uint64(len(m.ViewName)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.ViewKind) > 0 {
		i -= len(m.ViewKind)
		copy(dAtA[i:], m.ViewKind)
		i = encodeVarintPublicCustomapi(dAtA, i, uint64(len(m.ViewKind)))
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

func (m *GetResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.TerraformParameters != nil {
		{
			size, err := m.TerraformParameters.MarshalToSizedBuffer(dAtA[:i])
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

func (m *GetStatusResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetStatusResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetStatusResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Status != nil {
		{
			size, err := m.Status.MarshalToSizedBuffer(dAtA[:i])
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
func (m *GetRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Namespace)
	if l > 0 {
		n += 1 + l + sovPublicCustomapi(uint64(l))
	}
	l = len(m.ViewKind)
	if l > 0 {
		n += 1 + l + sovPublicCustomapi(uint64(l))
	}
	l = len(m.ViewName)
	if l > 0 {
		n += 1 + l + sovPublicCustomapi(uint64(l))
	}
	return n
}

func (m *GetResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.TerraformParameters != nil {
		l = m.TerraformParameters.Size()
		n += 1 + l + sovPublicCustomapi(uint64(l))
	}
	return n
}

func (m *GetStatusResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Status != nil {
		l = m.Status.Size()
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
func (this *GetRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&GetRequest{`,
		`Namespace:` + fmt.Sprintf("%v", this.Namespace) + `,`,
		`ViewKind:` + fmt.Sprintf("%v", this.ViewKind) + `,`,
		`ViewName:` + fmt.Sprintf("%v", this.ViewName) + `,`,
		`}`,
	}, "")
	return s
}
func (this *GetResponse) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&GetResponse{`,
		`TerraformParameters:` + strings.Replace(fmt.Sprintf("%v", this.TerraformParameters), "GlobalSpecType", "GlobalSpecType", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *GetStatusResponse) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&GetStatusResponse{`,
		`Status:` + strings.Replace(fmt.Sprintf("%v", this.Status), "StatusObject", "StatusObject", 1) + `,`,
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
func (m *GetRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: GetRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetRequest: illegal tag %d (wire type %d)", fieldNum, wire)
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
				return fmt.Errorf("proto: wrong wireType = %d for field ViewKind", wireType)
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
			m.ViewKind = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ViewName", wireType)
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
			m.ViewName = string(dAtA[iNdEx:postIndex])
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
func (m *GetResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: GetResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TerraformParameters", wireType)
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
			if m.TerraformParameters == nil {
				m.TerraformParameters = &GlobalSpecType{}
			}
			if err := m.TerraformParameters.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *GetStatusResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: GetStatusResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetStatusResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
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
			if m.Status == nil {
				m.Status = &StatusObject{}
			}
			if err := m.Status.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
