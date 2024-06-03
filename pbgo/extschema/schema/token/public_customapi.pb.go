// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ves.io/schema/token/public_customapi.proto

// Token API
//
// x-displayName: "Token"
// Token API is used to allocate, manage and delete tokens used for
// registration of new Customer Edge sites. Single token can be used to register multiple sites

package token

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

// Change token
//
// x-displayName: "Change Token"
// Generic response when object is changed
type ObjectChangeResp struct {
	// Token object
	//
	// x-displayName: "Token Object"
	Obj *Object `protobuf:"bytes,1,opt,name=obj,proto3" json:"obj,omitempty"`
}

func (m *ObjectChangeResp) Reset()      { *m = ObjectChangeResp{} }
func (*ObjectChangeResp) ProtoMessage() {}
func (*ObjectChangeResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_1daf03efacb073ee, []int{0}
}
func (m *ObjectChangeResp) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ObjectChangeResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ObjectChangeResp.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ObjectChangeResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ObjectChangeResp.Merge(m, src)
}
func (m *ObjectChangeResp) XXX_Size() int {
	return m.Size()
}
func (m *ObjectChangeResp) XXX_DiscardUnknown() {
	xxx_messageInfo_ObjectChangeResp.DiscardUnknown(m)
}

var xxx_messageInfo_ObjectChangeResp proto.InternalMessageInfo

func (m *ObjectChangeResp) GetObj() *Object {
	if m != nil {
		return m.Obj
	}
	return nil
}

// Change token state
//
// x-displayName: "State Request"
// Changes token state, only state is changing, namespace and name is used to find token to change
type StateReq struct {
	// Token namespace
	//
	// x-displayName: "Namespace"
	// x-example: "value"
	// Namespace where token is residing
	Namespace string `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	// Token name
	//
	// x-displayName: "Name"
	// x-example: "value"
	// Token object name to change, it's usually same as Metadata.Uid
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Token state
	//
	// x-displayName: "Token State"
	// Change token state to this state
	State State `protobuf:"varint,3,opt,name=state,proto3,enum=ves.io.schema.token.State" json:"state,omitempty"`
}

func (m *StateReq) Reset()      { *m = StateReq{} }
func (*StateReq) ProtoMessage() {}
func (*StateReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_1daf03efacb073ee, []int{1}
}
func (m *StateReq) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *StateReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_StateReq.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *StateReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StateReq.Merge(m, src)
}
func (m *StateReq) XXX_Size() int {
	return m.Size()
}
func (m *StateReq) XXX_DiscardUnknown() {
	xxx_messageInfo_StateReq.DiscardUnknown(m)
}

var xxx_messageInfo_StateReq proto.InternalMessageInfo

func (m *StateReq) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *StateReq) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *StateReq) GetState() State {
	if m != nil {
		return m.State
	}
	return UNKNOWN
}

func init() {
	proto.RegisterType((*ObjectChangeResp)(nil), "ves.io.schema.token.ObjectChangeResp")
	golang_proto.RegisterType((*ObjectChangeResp)(nil), "ves.io.schema.token.ObjectChangeResp")
	proto.RegisterType((*StateReq)(nil), "ves.io.schema.token.StateReq")
	golang_proto.RegisterType((*StateReq)(nil), "ves.io.schema.token.StateReq")
}

func init() {
	proto.RegisterFile("ves.io/schema/token/public_customapi.proto", fileDescriptor_1daf03efacb073ee)
}
func init() {
	golang_proto.RegisterFile("ves.io/schema/token/public_customapi.proto", fileDescriptor_1daf03efacb073ee)
}

var fileDescriptor_1daf03efacb073ee = []byte{
	// 535 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xbf, 0x8b, 0x13, 0x41,
	0x14, 0xc7, 0x77, 0x12, 0x15, 0x6f, 0x05, 0x91, 0xb5, 0x09, 0x7b, 0xe7, 0x18, 0x02, 0x42, 0x38,
	0xdc, 0x1d, 0x89, 0x88, 0x60, 0xa5, 0x77, 0x20, 0x5c, 0xa5, 0x44, 0x2b, 0x1b, 0x9d, 0xdd, 0xbc,
	0x6c, 0xe6, 0x2e, 0x3b, 0x33, 0x37, 0x33, 0xbb, 0x9e, 0x48, 0x40, 0xce, 0x7f, 0x40, 0xf0, 0x2f,
	0xb0, 0xb3, 0xf2, 0x1f, 0xb8, 0x26, 0x9d, 0x56, 0x12, 0xb4, 0xb9, 0xd2, 0x6c, 0x2c, 0x2c, 0xaf,
	0xb2, 0x96, 0xcc, 0xee, 0xfd, 0x08, 0xc4, 0xeb, 0xde, 0x9b, 0xf7, 0x79, 0xdf, 0x99, 0x79, 0xef,
	0xeb, 0xae, 0xe7, 0xa0, 0x43, 0x26, 0x88, 0x8e, 0x07, 0x90, 0x52, 0x62, 0xc4, 0x0e, 0x70, 0x22,
	0xb3, 0x68, 0xc8, 0xe2, 0x97, 0x71, 0xa6, 0x8d, 0x48, 0xa9, 0x64, 0xa1, 0x54, 0xc2, 0x08, 0xef,
	0x7a, 0xc9, 0x86, 0x25, 0x1b, 0x5a, 0xd6, 0x0f, 0x12, 0x66, 0x06, 0x59, 0x14, 0xc6, 0x22, 0x25,
	0x89, 0x48, 0x04, 0xb1, 0x6c, 0x94, 0xf5, 0x6d, 0x66, 0x13, 0x1b, 0x95, 0x1a, 0xfe, 0x5a, 0x22,
	0x44, 0x32, 0x04, 0x42, 0x25, 0x23, 0x94, 0x73, 0x61, 0xa8, 0x61, 0x82, 0xeb, 0xaa, 0xba, 0xba,
	0xf8, 0x1a, 0x21, 0xcf, 0x16, 0x9b, 0xcb, 0x9e, 0x2a, 0xa2, 0x6d, 0x88, 0x4d, 0x45, 0xdc, 0x5c,
	0x46, 0x98, 0x37, 0x12, 0x8e, 0x25, 0x5a, 0x8b, 0x40, 0x0e, 0x1a, 0x78, 0xbe, 0x78, 0x4d, 0xeb,
	0xb1, 0x7b, 0xed, 0x89, 0x15, 0xdd, 0x1c, 0x50, 0x9e, 0x40, 0x17, 0xb4, 0xf4, 0x3a, 0x6e, 0x5d,
	0x44, 0xdb, 0x0d, 0xd4, 0x44, 0xed, 0x2b, 0x9d, 0xd5, 0x70, 0xc9, 0x1c, 0xc2, 0xb2, 0x67, 0xa3,
	0x3e, 0x1e, 0xa1, 0xee, 0x1c, 0x6e, 0x71, 0xf7, 0xf2, 0x33, 0x43, 0x0d, 0x74, 0x61, 0xd7, 0x5b,
	0x73, 0x57, 0x38, 0x4d, 0x41, 0x4b, 0x1a, 0x83, 0x55, 0x59, 0xe9, 0x9e, 0x1e, 0x78, 0x9e, 0x7b,
	0x61, 0x9e, 0x34, 0x6a, 0xb6, 0x60, 0x63, 0xef, 0x8e, 0x7b, 0x51, 0xcf, 0xbb, 0x1b, 0xf5, 0x26,
	0x6a, 0x5f, 0xed, 0xf8, 0x4b, 0xef, 0x2c, 0xf5, 0x4b, 0xb0, 0xf3, 0xa5, 0xe6, 0xae, 0x6c, 0xda,
	0x8d, 0x3d, 0x7a, 0xba, 0xe5, 0xfd, 0x45, 0xae, 0xfb, 0x7c, 0x0e, 0x59, 0xc6, 0xbb, 0x71, 0x4e,
	0x3f, 0xec, 0xfa, 0xb7, 0xce, 0xf9, 0xd2, 0xe9, 0x18, 0x5a, 0x9f, 0x50, 0xf1, 0xb5, 0xf1, 0xaa,
	0x7f, 0x6f, 0x2f, 0x0e, 0x34, 0x33, 0x10, 0xa4, 0x94, 0xd3, 0x04, 0x52, 0xe0, 0x26, 0x88, 0xa8,
	0x66, 0x71, 0x90, 0x69, 0x50, 0xb7, 0x9b, 0x4b, 0x01, 0x6d, 0x28, 0xef, 0x51, 0xd5, 0xab, 0x98,
	0x1c, 0x74, 0xc0, 0x44, 0xc0, 0x78, 0x5f, 0x51, 0x6d, 0x54, 0x16, 0x9b, 0x4c, 0x41, 0xf0, 0x5a,
	0x31, 0x03, 0xfb, 0x3f, 0x7f, 0x7f, 0xac, 0xdd, 0x6f, 0x75, 0x2a, 0x1b, 0x92, 0x93, 0x41, 0x69,
	0xf2, 0xf6, 0x24, 0x1e, 0x95, 0xdb, 0xad, 0x8e, 0x46, 0xc4, 0xce, 0xe0, 0x01, 0x5a, 0xf7, 0xdb,
	0xe3, 0x03, 0x54, 0xff, 0x71, 0x80, 0xfc, 0xff, 0xff, 0x68, 0xff, 0x7b, 0xa3, 0xf6, 0x10, 0x6d,
	0xbc, 0x47, 0x93, 0x29, 0x76, 0x0e, 0xa7, 0xd8, 0x39, 0x9a, 0x62, 0xf4, 0xae, 0xc0, 0xe8, 0x73,
	0x81, 0xd1, 0xb7, 0x02, 0xa3, 0x49, 0x81, 0xd1, 0xaf, 0x02, 0xa3, 0x3f, 0x05, 0x76, 0x8e, 0x0a,
	0x8c, 0x3e, 0xcc, 0xb0, 0x33, 0x9e, 0x61, 0x34, 0x99, 0x61, 0xe7, 0x70, 0x86, 0x9d, 0x17, 0x5b,
	0x89, 0x90, 0x3b, 0x49, 0x98, 0x8b, 0xa1, 0x01, 0xa5, 0x68, 0x98, 0x69, 0x62, 0x83, 0xbe, 0x50,
	0x69, 0x20, 0x95, 0xc8, 0x59, 0x0f, 0x54, 0x70, 0x5c, 0x26, 0x32, 0x4a, 0x04, 0x81, 0x3d, 0x53,
	0xd9, 0xee, 0xac, 0x3d, 0xa3, 0x4b, 0xd6, 0x75, 0x77, 0xff, 0x05, 0x00, 0x00, 0xff, 0xff, 0xd6,
	0x8f, 0xdb, 0x29, 0x89, 0x03, 0x00, 0x00,
}

func (this *ObjectChangeResp) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ObjectChangeResp)
	if !ok {
		that2, ok := that.(ObjectChangeResp)
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
	if !this.Obj.Equal(that1.Obj) {
		return false
	}
	return true
}
func (this *StateReq) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*StateReq)
	if !ok {
		that2, ok := that.(StateReq)
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
	if this.Name != that1.Name {
		return false
	}
	if this.State != that1.State {
		return false
	}
	return true
}
func (this *ObjectChangeResp) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&token.ObjectChangeResp{")
	if this.Obj != nil {
		s = append(s, "Obj: "+fmt.Sprintf("%#v", this.Obj)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *StateReq) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 7)
	s = append(s, "&token.StateReq{")
	s = append(s, "Namespace: "+fmt.Sprintf("%#v", this.Namespace)+",\n")
	s = append(s, "Name: "+fmt.Sprintf("%#v", this.Name)+",\n")
	s = append(s, "State: "+fmt.Sprintf("%#v", this.State)+",\n")
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
	// Token states
	//
	// x-displayName: "Set Token State"
	// TokenState changes token status, it can be used to disable token.
	TokenState(ctx context.Context, in *StateReq, opts ...grpc.CallOption) (*ObjectChangeResp, error)
}

type customAPIClient struct {
	cc *grpc.ClientConn
}

func NewCustomAPIClient(cc *grpc.ClientConn) CustomAPIClient {
	return &customAPIClient{cc}
}

func (c *customAPIClient) TokenState(ctx context.Context, in *StateReq, opts ...grpc.CallOption) (*ObjectChangeResp, error) {
	out := new(ObjectChangeResp)
	err := c.cc.Invoke(ctx, "/ves.io.schema.token.CustomAPI/TokenState", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CustomAPIServer is the server API for CustomAPI service.
type CustomAPIServer interface {
	// Token states
	//
	// x-displayName: "Set Token State"
	// TokenState changes token status, it can be used to disable token.
	TokenState(context.Context, *StateReq) (*ObjectChangeResp, error)
}

// UnimplementedCustomAPIServer can be embedded to have forward compatible implementations.
type UnimplementedCustomAPIServer struct {
}

func (*UnimplementedCustomAPIServer) TokenState(ctx context.Context, req *StateReq) (*ObjectChangeResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TokenState not implemented")
}

func RegisterCustomAPIServer(s *grpc.Server, srv CustomAPIServer) {
	s.RegisterService(&_CustomAPI_serviceDesc, srv)
}

func _CustomAPI_TokenState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomAPIServer).TokenState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ves.io.schema.token.CustomAPI/TokenState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomAPIServer).TokenState(ctx, req.(*StateReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _CustomAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ves.io.schema.token.CustomAPI",
	HandlerType: (*CustomAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "TokenState",
			Handler:    _CustomAPI_TokenState_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ves.io/schema/token/public_customapi.proto",
}

func (m *ObjectChangeResp) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ObjectChangeResp) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ObjectChangeResp) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Obj != nil {
		{
			size, err := m.Obj.MarshalToSizedBuffer(dAtA[:i])
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

func (m *StateReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StateReq) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *StateReq) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.State != 0 {
		i = encodeVarintPublicCustomapi(dAtA, i, uint64(m.State))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintPublicCustomapi(dAtA, i, uint64(len(m.Name)))
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
func (m *ObjectChangeResp) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Obj != nil {
		l = m.Obj.Size()
		n += 1 + l + sovPublicCustomapi(uint64(l))
	}
	return n
}

func (m *StateReq) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Namespace)
	if l > 0 {
		n += 1 + l + sovPublicCustomapi(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovPublicCustomapi(uint64(l))
	}
	if m.State != 0 {
		n += 1 + sovPublicCustomapi(uint64(m.State))
	}
	return n
}

func sovPublicCustomapi(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPublicCustomapi(x uint64) (n int) {
	return sovPublicCustomapi(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *ObjectChangeResp) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ObjectChangeResp{`,
		`Obj:` + strings.Replace(fmt.Sprintf("%v", this.Obj), "Object", "Object", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *StateReq) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&StateReq{`,
		`Namespace:` + fmt.Sprintf("%v", this.Namespace) + `,`,
		`Name:` + fmt.Sprintf("%v", this.Name) + `,`,
		`State:` + fmt.Sprintf("%v", this.State) + `,`,
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
func (m *ObjectChangeResp) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: ObjectChangeResp: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ObjectChangeResp: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Obj", wireType)
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
			if m.Obj == nil {
				m.Obj = &Object{}
			}
			if err := m.Obj.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *StateReq) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: StateReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StateReq: illegal tag %d (wire type %d)", fieldNum, wire)
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
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
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
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field State", wireType)
			}
			m.State = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPublicCustomapi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.State |= State(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
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
