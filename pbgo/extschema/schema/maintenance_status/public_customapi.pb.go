// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ves.io/schema/maintenance_status/public_customapi.proto

// Upgrade Status
//
// x-displayName: "F5xC Upgrade Status"

package maintenance_status

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
	schema "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema"
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

// Upgrade Status Request
//
// x-displayName: "Upgrade Status Request"
// Request to get upgrade status.
type GetUpgradeStatusRequest struct {
}

func (m *GetUpgradeStatusRequest) Reset()      { *m = GetUpgradeStatusRequest{} }
func (*GetUpgradeStatusRequest) ProtoMessage() {}
func (*GetUpgradeStatusRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e94ff602ee35b59, []int{0}
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

// Upgrade Status Response
//
// x-displayName: "Upgrade Status Response"
// Response message for Upgrade Status Request.
// Response contain different states in upgrade process.
type GetUpgradeStatusResponse struct {
	// States in the Upgrade status.
	//
	// x-displayName: "Upgrade status States"
	// Upgrade cycle has different status as listed below.
	//
	// Types that are valid to be assigned to UpgradeState:
	//	*GetUpgradeStatusResponse_UpgradeNotInProgress
	//	*GetUpgradeStatusResponse_UpgradeInProgress
	//	*GetUpgradeStatusResponse_UpgradeInProgressWithConfigDowntime
	UpgradeState isGetUpgradeStatusResponse_UpgradeState `protobuf_oneof:"upgrade_state"`
}

func (m *GetUpgradeStatusResponse) Reset()      { *m = GetUpgradeStatusResponse{} }
func (*GetUpgradeStatusResponse) ProtoMessage() {}
func (*GetUpgradeStatusResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e94ff602ee35b59, []int{1}
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

type isGetUpgradeStatusResponse_UpgradeState interface {
	isGetUpgradeStatusResponse_UpgradeState()
	Equal(interface{}) bool
	MarshalTo([]byte) (int, error)
	Size() int
}

type GetUpgradeStatusResponse_UpgradeNotInProgress struct {
	UpgradeNotInProgress *schema.Empty `protobuf:"bytes,1,opt,name=upgrade_not_in_progress,json=upgradeNotInProgress,proto3,oneof" json:"upgrade_not_in_progress,omitempty"`
}
type GetUpgradeStatusResponse_UpgradeInProgress struct {
	UpgradeInProgress *schema.Empty `protobuf:"bytes,2,opt,name=upgrade_in_progress,json=upgradeInProgress,proto3,oneof" json:"upgrade_in_progress,omitempty"`
}
type GetUpgradeStatusResponse_UpgradeInProgressWithConfigDowntime struct {
	UpgradeInProgressWithConfigDowntime *schema.Empty `protobuf:"bytes,3,opt,name=upgrade_in_progress_with_config_downtime,json=upgradeInProgressWithConfigDowntime,proto3,oneof" json:"upgrade_in_progress_with_config_downtime,omitempty"`
}

func (*GetUpgradeStatusResponse_UpgradeNotInProgress) isGetUpgradeStatusResponse_UpgradeState() {}
func (*GetUpgradeStatusResponse_UpgradeInProgress) isGetUpgradeStatusResponse_UpgradeState()    {}
func (*GetUpgradeStatusResponse_UpgradeInProgressWithConfigDowntime) isGetUpgradeStatusResponse_UpgradeState() {
}

func (m *GetUpgradeStatusResponse) GetUpgradeState() isGetUpgradeStatusResponse_UpgradeState {
	if m != nil {
		return m.UpgradeState
	}
	return nil
}

func (m *GetUpgradeStatusResponse) GetUpgradeNotInProgress() *schema.Empty {
	if x, ok := m.GetUpgradeState().(*GetUpgradeStatusResponse_UpgradeNotInProgress); ok {
		return x.UpgradeNotInProgress
	}
	return nil
}

func (m *GetUpgradeStatusResponse) GetUpgradeInProgress() *schema.Empty {
	if x, ok := m.GetUpgradeState().(*GetUpgradeStatusResponse_UpgradeInProgress); ok {
		return x.UpgradeInProgress
	}
	return nil
}

func (m *GetUpgradeStatusResponse) GetUpgradeInProgressWithConfigDowntime() *schema.Empty {
	if x, ok := m.GetUpgradeState().(*GetUpgradeStatusResponse_UpgradeInProgressWithConfigDowntime); ok {
		return x.UpgradeInProgressWithConfigDowntime
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*GetUpgradeStatusResponse) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*GetUpgradeStatusResponse_UpgradeNotInProgress)(nil),
		(*GetUpgradeStatusResponse_UpgradeInProgress)(nil),
		(*GetUpgradeStatusResponse_UpgradeInProgressWithConfigDowntime)(nil),
	}
}

func init() {
	proto.RegisterType((*GetUpgradeStatusRequest)(nil), "ves.io.schema.maintenance_status.GetUpgradeStatusRequest")
	golang_proto.RegisterType((*GetUpgradeStatusRequest)(nil), "ves.io.schema.maintenance_status.GetUpgradeStatusRequest")
	proto.RegisterType((*GetUpgradeStatusResponse)(nil), "ves.io.schema.maintenance_status.GetUpgradeStatusResponse")
	golang_proto.RegisterType((*GetUpgradeStatusResponse)(nil), "ves.io.schema.maintenance_status.GetUpgradeStatusResponse")
}

func init() {
	proto.RegisterFile("ves.io/schema/maintenance_status/public_customapi.proto", fileDescriptor_0e94ff602ee35b59)
}
func init() {
	golang_proto.RegisterFile("ves.io/schema/maintenance_status/public_customapi.proto", fileDescriptor_0e94ff602ee35b59)
}

var fileDescriptor_0e94ff602ee35b59 = []byte{
	// 564 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x94, 0x3d, 0x6b, 0x14, 0x41,
	0x18, 0xc7, 0x77, 0x36, 0x20, 0xb8, 0x22, 0xea, 0x19, 0xc8, 0xe5, 0x0c, 0x43, 0x38, 0x9b, 0x20,
	0xee, 0x0e, 0x44, 0x44, 0xb4, 0x33, 0xf1, 0x25, 0x29, 0x94, 0x10, 0x51, 0xc1, 0x66, 0x99, 0xdb,
	0x7b, 0x6e, 0x6e, 0xf0, 0x76, 0x9e, 0x71, 0x66, 0x76, 0x93, 0x74, 0x21, 0xb5, 0x45, 0xc0, 0xef,
	0x20, 0xb6, 0x76, 0x82, 0x4d, 0x3a, 0x53, 0x49, 0xc0, 0x26, 0xa5, 0xd9, 0xb3, 0xb0, 0xcc, 0x47,
	0x10, 0xf7, 0xf6, 0xe4, 0x2e, 0x67, 0x22, 0xd8, 0xcd, 0xf0, 0xe3, 0xff, 0xdb, 0xdd, 0xe7, 0x65,
	0x83, 0x3b, 0x39, 0xd8, 0x48, 0x22, 0xb3, 0x49, 0x17, 0x52, 0xce, 0x52, 0x2e, 0x95, 0x03, 0xc5,
	0x55, 0x02, 0xb1, 0x75, 0xdc, 0x65, 0x96, 0xe9, 0xac, 0xd5, 0x93, 0x49, 0x9c, 0x64, 0xd6, 0x61,
	0xca, 0xb5, 0x8c, 0xb4, 0x41, 0x87, 0xb5, 0xf9, 0x41, 0x30, 0x1a, 0x04, 0xa3, 0xc9, 0x60, 0x23,
	0x14, 0xd2, 0x75, 0xb3, 0x56, 0x94, 0x60, 0xca, 0x04, 0x0a, 0x64, 0x65, 0xb0, 0x95, 0x75, 0xca,
	0x5b, 0x79, 0x29, 0x4f, 0x03, 0x61, 0x63, 0x4e, 0x20, 0x8a, 0x1e, 0x30, 0xae, 0x25, 0xe3, 0x4a,
	0xa1, 0xe3, 0x4e, 0xa2, 0xb2, 0x15, 0xbd, 0x36, 0xfe, 0x9e, 0xa8, 0x47, 0xe1, 0xec, 0x38, 0x74,
	0x5b, 0x1a, 0x86, 0x68, 0x6e, 0x1c, 0xe5, 0xbc, 0x27, 0xdb, 0xdc, 0x41, 0x45, 0x9b, 0x27, 0x28,
	0x58, 0x50, 0xf9, 0xb8, 0xbc, 0x39, 0x1b, 0xcc, 0x3c, 0x06, 0xf7, 0x5c, 0x0b, 0xc3, 0xdb, 0xf0,
	0xac, 0xfc, 0xb4, 0x75, 0x78, 0x93, 0x81, 0x75, 0xcd, 0x8f, 0x7e, 0x50, 0x9f, 0x64, 0x56, 0xa3,
	0xb2, 0x50, 0x7b, 0x12, 0xcc, 0x64, 0x03, 0x10, 0x2b, 0x74, 0xb1, 0x54, 0xb1, 0x36, 0x28, 0x0c,
	0x58, 0x5b, 0x27, 0xf3, 0x64, 0xe1, 0xc2, 0xe2, 0x74, 0x34, 0x5e, 0xc2, 0x87, 0xa9, 0x76, 0x5b,
	0x2b, 0xde, 0xfa, 0x74, 0x15, 0x7b, 0x8a, 0x6e, 0x55, 0xad, 0x55, 0x99, 0xda, 0xa3, 0xe0, 0xea,
	0x50, 0x37, 0xaa, 0xf2, 0xcf, 0x54, 0x5d, 0xa9, 0x22, 0x23, 0x9e, 0x34, 0x58, 0xf8, 0x8b, 0x27,
	0xde, 0x90, 0xae, 0x1b, 0x27, 0xa8, 0x3a, 0x52, 0xc4, 0x6d, 0xdc, 0x50, 0x4e, 0xa6, 0x50, 0x9f,
	0x3a, 0x53, 0x7e, 0x7d, 0x42, 0xfe, 0x52, 0xba, 0xee, 0x72, 0xe9, 0x78, 0x50, 0x29, 0x96, 0x2e,
	0x05, 0x17, 0x87, 0x8f, 0xfb, 0x3d, 0x16, 0xb0, 0xf8, 0xde, 0x0f, 0xce, 0x2f, 0x97, 0xb3, 0x74,
	0x7f, 0x6d, 0xb5, 0xf6, 0xd6, 0x0f, 0x2e, 0x9f, 0xac, 0x60, 0xed, 0x6e, 0xf4, 0xaf, 0xd9, 0x8a,
	0x4e, 0xe9, 0x48, 0xe3, 0xde, 0xff, 0x44, 0x07, 0x0d, 0x6b, 0x6e, 0x93, 0xfd, 0x4f, 0x3e, 0x29,
	0xbe, 0xd4, 0x57, 0x3a, 0xb7, 0x37, 0x93, 0x30, 0x41, 0x65, 0xb1, 0x07, 0x61, 0x8b, 0x5b, 0x99,
	0x84, 0x29, 0x2a, 0xe9, 0xd0, 0xdc, 0x3c, 0x89, 0xe0, 0x0f, 0xc9, 0xc1, 0x86, 0x12, 0x43, 0x01,
	0x0a, 0x0c, 0xef, 0x85, 0x06, 0x78, 0x7b, 0xe7, 0xdb, 0x8f, 0x77, 0xfe, 0x8d, 0xda, 0x42, 0xb5,
	0x43, 0x4c, 0xf1, 0x14, 0xac, 0xe6, 0x09, 0x58, 0x66, 0xb7, 0xac, 0x83, 0x94, 0x8d, 0xd6, 0x26,
	0xb3, 0x8d, 0x60, 0xef, 0x33, 0x99, 0xda, 0xf9, 0x5a, 0xf7, 0xeb, 0x64, 0x69, 0x97, 0x1c, 0x1c,
	0x51, 0xef, 0xf0, 0x88, 0x7a, 0xc7, 0x47, 0x94, 0x6c, 0x17, 0x94, 0x7c, 0x28, 0x28, 0xd9, 0x2f,
	0x28, 0x39, 0x28, 0x28, 0xf9, 0x5e, 0x50, 0xf2, 0xb3, 0xa0, 0xde, 0x71, 0x41, 0xc9, 0x6e, 0x9f,
	0x7a, 0x7b, 0x7d, 0x4a, 0x0e, 0xfa, 0xd4, 0x3b, 0xec, 0x53, 0xef, 0xd5, 0x0b, 0x81, 0xfa, 0xb5,
	0x88, 0x72, 0xec, 0x39, 0x30, 0x86, 0x47, 0x99, 0x65, 0xe5, 0xa1, 0x83, 0x26, 0x0d, 0xb5, 0xc1,
	0x5c, 0xb6, 0xc1, 0x84, 0x43, 0xcc, 0x74, 0x4b, 0x20, 0x83, 0x4d, 0x57, 0x6d, 0xc1, 0xa9, 0xbf,
	0x82, 0xd6, 0xb9, 0x72, 0x23, 0x6e, 0xfd, 0x0a, 0x00, 0x00, 0xff, 0xff, 0x5a, 0xb6, 0x38, 0xe3,
	0x35, 0x04, 0x00, 0x00,
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
	if that1.UpgradeState == nil {
		if this.UpgradeState != nil {
			return false
		}
	} else if this.UpgradeState == nil {
		return false
	} else if !this.UpgradeState.Equal(that1.UpgradeState) {
		return false
	}
	return true
}
func (this *GetUpgradeStatusResponse_UpgradeNotInProgress) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*GetUpgradeStatusResponse_UpgradeNotInProgress)
	if !ok {
		that2, ok := that.(GetUpgradeStatusResponse_UpgradeNotInProgress)
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
	if !this.UpgradeNotInProgress.Equal(that1.UpgradeNotInProgress) {
		return false
	}
	return true
}
func (this *GetUpgradeStatusResponse_UpgradeInProgress) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*GetUpgradeStatusResponse_UpgradeInProgress)
	if !ok {
		that2, ok := that.(GetUpgradeStatusResponse_UpgradeInProgress)
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
	if !this.UpgradeInProgress.Equal(that1.UpgradeInProgress) {
		return false
	}
	return true
}
func (this *GetUpgradeStatusResponse_UpgradeInProgressWithConfigDowntime) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*GetUpgradeStatusResponse_UpgradeInProgressWithConfigDowntime)
	if !ok {
		that2, ok := that.(GetUpgradeStatusResponse_UpgradeInProgressWithConfigDowntime)
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
	if !this.UpgradeInProgressWithConfigDowntime.Equal(that1.UpgradeInProgressWithConfigDowntime) {
		return false
	}
	return true
}
func (this *GetUpgradeStatusRequest) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 4)
	s = append(s, "&maintenance_status.GetUpgradeStatusRequest{")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *GetUpgradeStatusResponse) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 7)
	s = append(s, "&maintenance_status.GetUpgradeStatusResponse{")
	if this.UpgradeState != nil {
		s = append(s, "UpgradeState: "+fmt.Sprintf("%#v", this.UpgradeState)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *GetUpgradeStatusResponse_UpgradeNotInProgress) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&maintenance_status.GetUpgradeStatusResponse_UpgradeNotInProgress{` +
		`UpgradeNotInProgress:` + fmt.Sprintf("%#v", this.UpgradeNotInProgress) + `}`}, ", ")
	return s
}
func (this *GetUpgradeStatusResponse_UpgradeInProgress) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&maintenance_status.GetUpgradeStatusResponse_UpgradeInProgress{` +
		`UpgradeInProgress:` + fmt.Sprintf("%#v", this.UpgradeInProgress) + `}`}, ", ")
	return s
}
func (this *GetUpgradeStatusResponse_UpgradeInProgressWithConfigDowntime) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&maintenance_status.GetUpgradeStatusResponse_UpgradeInProgressWithConfigDowntime{` +
		`UpgradeInProgressWithConfigDowntime:` + fmt.Sprintf("%#v", this.UpgradeInProgressWithConfigDowntime) + `}`}, ", ")
	return s
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
	// Upgrade Status
	//
	// x-displayName: "Upgrade Status"
	// Request to get the upgrade status
	GetUpgradeStatus(ctx context.Context, in *GetUpgradeStatusRequest, opts ...grpc.CallOption) (*GetUpgradeStatusResponse, error)
}

type customAPIClient struct {
	cc *grpc.ClientConn
}

func NewCustomAPIClient(cc *grpc.ClientConn) CustomAPIClient {
	return &customAPIClient{cc}
}

func (c *customAPIClient) GetUpgradeStatus(ctx context.Context, in *GetUpgradeStatusRequest, opts ...grpc.CallOption) (*GetUpgradeStatusResponse, error) {
	out := new(GetUpgradeStatusResponse)
	err := c.cc.Invoke(ctx, "/ves.io.schema.maintenance_status.CustomAPI/GetUpgradeStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CustomAPIServer is the server API for CustomAPI service.
type CustomAPIServer interface {
	// Upgrade Status
	//
	// x-displayName: "Upgrade Status"
	// Request to get the upgrade status
	GetUpgradeStatus(context.Context, *GetUpgradeStatusRequest) (*GetUpgradeStatusResponse, error)
}

// UnimplementedCustomAPIServer can be embedded to have forward compatible implementations.
type UnimplementedCustomAPIServer struct {
}

func (*UnimplementedCustomAPIServer) GetUpgradeStatus(ctx context.Context, req *GetUpgradeStatusRequest) (*GetUpgradeStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUpgradeStatus not implemented")
}

func RegisterCustomAPIServer(s *grpc.Server, srv CustomAPIServer) {
	s.RegisterService(&_CustomAPI_serviceDesc, srv)
}

func _CustomAPI_GetUpgradeStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUpgradeStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomAPIServer).GetUpgradeStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ves.io.schema.maintenance_status.CustomAPI/GetUpgradeStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomAPIServer).GetUpgradeStatus(ctx, req.(*GetUpgradeStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CustomAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ves.io.schema.maintenance_status.CustomAPI",
	HandlerType: (*CustomAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUpgradeStatus",
			Handler:    _CustomAPI_GetUpgradeStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ves.io/schema/maintenance_status/public_customapi.proto",
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
	if m.UpgradeState != nil {
		{
			size := m.UpgradeState.Size()
			i -= size
			if _, err := m.UpgradeState.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	return len(dAtA) - i, nil
}

func (m *GetUpgradeStatusResponse_UpgradeNotInProgress) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetUpgradeStatusResponse_UpgradeNotInProgress) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.UpgradeNotInProgress != nil {
		{
			size, err := m.UpgradeNotInProgress.MarshalToSizedBuffer(dAtA[:i])
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
func (m *GetUpgradeStatusResponse_UpgradeInProgress) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetUpgradeStatusResponse_UpgradeInProgress) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.UpgradeInProgress != nil {
		{
			size, err := m.UpgradeInProgress.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintPublicCustomapi(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	return len(dAtA) - i, nil
}
func (m *GetUpgradeStatusResponse_UpgradeInProgressWithConfigDowntime) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetUpgradeStatusResponse_UpgradeInProgressWithConfigDowntime) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.UpgradeInProgressWithConfigDowntime != nil {
		{
			size, err := m.UpgradeInProgressWithConfigDowntime.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintPublicCustomapi(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
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
	return n
}

func (m *GetUpgradeStatusResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.UpgradeState != nil {
		n += m.UpgradeState.Size()
	}
	return n
}

func (m *GetUpgradeStatusResponse_UpgradeNotInProgress) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.UpgradeNotInProgress != nil {
		l = m.UpgradeNotInProgress.Size()
		n += 1 + l + sovPublicCustomapi(uint64(l))
	}
	return n
}
func (m *GetUpgradeStatusResponse_UpgradeInProgress) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.UpgradeInProgress != nil {
		l = m.UpgradeInProgress.Size()
		n += 1 + l + sovPublicCustomapi(uint64(l))
	}
	return n
}
func (m *GetUpgradeStatusResponse_UpgradeInProgressWithConfigDowntime) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.UpgradeInProgressWithConfigDowntime != nil {
		l = m.UpgradeInProgressWithConfigDowntime.Size()
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
		`}`,
	}, "")
	return s
}
func (this *GetUpgradeStatusResponse) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&GetUpgradeStatusResponse{`,
		`UpgradeState:` + fmt.Sprintf("%v", this.UpgradeState) + `,`,
		`}`,
	}, "")
	return s
}
func (this *GetUpgradeStatusResponse_UpgradeNotInProgress) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&GetUpgradeStatusResponse_UpgradeNotInProgress{`,
		`UpgradeNotInProgress:` + strings.Replace(fmt.Sprintf("%v", this.UpgradeNotInProgress), "Empty", "schema.Empty", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *GetUpgradeStatusResponse_UpgradeInProgress) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&GetUpgradeStatusResponse_UpgradeInProgress{`,
		`UpgradeInProgress:` + strings.Replace(fmt.Sprintf("%v", this.UpgradeInProgress), "Empty", "schema.Empty", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *GetUpgradeStatusResponse_UpgradeInProgressWithConfigDowntime) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&GetUpgradeStatusResponse_UpgradeInProgressWithConfigDowntime{`,
		`UpgradeInProgressWithConfigDowntime:` + strings.Replace(fmt.Sprintf("%v", this.UpgradeInProgressWithConfigDowntime), "Empty", "schema.Empty", 1) + `,`,
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
				return fmt.Errorf("proto: wrong wireType = %d for field UpgradeNotInProgress", wireType)
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
			v := &schema.Empty{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.UpgradeState = &GetUpgradeStatusResponse_UpgradeNotInProgress{v}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UpgradeInProgress", wireType)
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
			v := &schema.Empty{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.UpgradeState = &GetUpgradeStatusResponse_UpgradeInProgress{v}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UpgradeInProgressWithConfigDowntime", wireType)
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
			v := &schema.Empty{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.UpgradeState = &GetUpgradeStatusResponse_UpgradeInProgressWithConfigDowntime{v}
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
