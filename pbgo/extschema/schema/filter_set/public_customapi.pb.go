// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ves.io/schema/filter_set/public_customapi.proto

// filter set
//
// x-displayName: "Filter Set"
// Filter Set is a set of saved filters used in the Console. This allows users to declare named sets
// of filters so that they can be consistently used and shared to quickly reactivate a particular
// view of the data in the Console.
//
// Filter Set has a context key, which maps a console page or view identifier to the filter set for later retrieval,
// and a list of field name/values.
//
// Filter Sets can be created in a custom namespace or the shared namespace.
// Any Filter Set created by a tenant should not have a value starting with "ves-io-"

package filter_set

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

// Find Filter Sets Request
//
// x-displayName: "Find Filter Sets Request"
// Find Filter Sets API returns FilterSets that match the given context key(s)
type FindFilterSetsReq struct {
	// namespace
	//
	// x-displayName: "Namespace For Filter Sets"
	// x-required
	// find filter sets in the given namespace
	// x-example: "bot-defense-apac"
	Namespace string `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	// context_keys
	//
	// x-displayName: "Context Keys"
	// x-required
	// context key(s) that idenfify one or more console pages/views where filters of the same field/values can be applied
	ContextKeys []string `protobuf:"bytes,2,rep,name=context_keys,json=contextKeys,proto3" json:"context_keys,omitempty"`
}

func (m *FindFilterSetsReq) Reset()      { *m = FindFilterSetsReq{} }
func (*FindFilterSetsReq) ProtoMessage() {}
func (*FindFilterSetsReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_714fe4b72c7534c5, []int{0}
}
func (m *FindFilterSetsReq) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *FindFilterSetsReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_FindFilterSetsReq.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *FindFilterSetsReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindFilterSetsReq.Merge(m, src)
}
func (m *FindFilterSetsReq) XXX_Size() int {
	return m.Size()
}
func (m *FindFilterSetsReq) XXX_DiscardUnknown() {
	xxx_messageInfo_FindFilterSetsReq.DiscardUnknown(m)
}

var xxx_messageInfo_FindFilterSetsReq proto.InternalMessageInfo

func (m *FindFilterSetsReq) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *FindFilterSetsReq) GetContextKeys() []string {
	if m != nil {
		return m.ContextKeys
	}
	return nil
}

// Find Filter Sets Response
//
// x-displayName: "Find Filter Sets Response"
// Response for Find Filter Sets API
type FindFilterSetsRsp struct {
	// filter_sets
	//
	// x-displayName: "Filter Sets"
	// List of filter sets with the given context key(s)
	FilterSets []*Object `protobuf:"bytes,1,rep,name=filter_sets,json=filterSets,proto3" json:"filter_sets,omitempty"`
}

func (m *FindFilterSetsRsp) Reset()      { *m = FindFilterSetsRsp{} }
func (*FindFilterSetsRsp) ProtoMessage() {}
func (*FindFilterSetsRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_714fe4b72c7534c5, []int{1}
}
func (m *FindFilterSetsRsp) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *FindFilterSetsRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_FindFilterSetsRsp.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *FindFilterSetsRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindFilterSetsRsp.Merge(m, src)
}
func (m *FindFilterSetsRsp) XXX_Size() int {
	return m.Size()
}
func (m *FindFilterSetsRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_FindFilterSetsRsp.DiscardUnknown(m)
}

var xxx_messageInfo_FindFilterSetsRsp proto.InternalMessageInfo

func (m *FindFilterSetsRsp) GetFilterSets() []*Object {
	if m != nil {
		return m.FilterSets
	}
	return nil
}

func init() {
	proto.RegisterType((*FindFilterSetsReq)(nil), "ves.io.schema.filter_set.FindFilterSetsReq")
	golang_proto.RegisterType((*FindFilterSetsReq)(nil), "ves.io.schema.filter_set.FindFilterSetsReq")
	proto.RegisterType((*FindFilterSetsRsp)(nil), "ves.io.schema.filter_set.FindFilterSetsRsp")
	golang_proto.RegisterType((*FindFilterSetsRsp)(nil), "ves.io.schema.filter_set.FindFilterSetsRsp")
}

func init() {
	proto.RegisterFile("ves.io/schema/filter_set/public_customapi.proto", fileDescriptor_714fe4b72c7534c5)
}
func init() {
	golang_proto.RegisterFile("ves.io/schema/filter_set/public_customapi.proto", fileDescriptor_714fe4b72c7534c5)
}

var fileDescriptor_714fe4b72c7534c5 = []byte{
	// 556 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0xcf, 0x6b, 0xd4, 0x4e,
	0x1c, 0xdd, 0xd9, 0xe5, 0xfb, 0xc5, 0x4d, 0x45, 0x34, 0x78, 0x08, 0x6b, 0x19, 0x97, 0x45, 0xa1,
	0x68, 0x93, 0x81, 0xaa, 0x08, 0xde, 0x5c, 0xa1, 0x20, 0x82, 0xca, 0x7a, 0xd3, 0x43, 0x99, 0x4c,
	0x3e, 0x49, 0xa7, 0x4d, 0xf2, 0x19, 0x33, 0x93, 0xd0, 0x22, 0x05, 0x29, 0x1e, 0x3c, 0x8a, 0xfe,
	0x13, 0x9e, 0x3c, 0x8b, 0xbd, 0xec, 0xcd, 0x9e, 0xa4, 0xe8, 0xa5, 0x47, 0x9b, 0xf5, 0xa0, 0xb7,
	0xfa, 0x1f, 0x88, 0x69, 0xfa, 0x63, 0x5d, 0x16, 0xbd, 0x7d, 0x66, 0xde, 0xfb, 0xbc, 0x97, 0xbc,
	0x79, 0x16, 0x2b, 0x40, 0x7b, 0x12, 0x99, 0x16, 0xcb, 0x90, 0x70, 0x16, 0xca, 0xd8, 0x40, 0xb6,
	0xa4, 0xc1, 0x30, 0x95, 0xfb, 0xb1, 0x14, 0x4b, 0x22, 0xd7, 0x06, 0x13, 0xae, 0xa4, 0xa7, 0x32,
	0x34, 0x68, 0x3b, 0x07, 0x0b, 0xde, 0xc1, 0x82, 0x77, 0xbc, 0xd0, 0x71, 0x23, 0x69, 0x96, 0x73,
	0xdf, 0x13, 0x98, 0xb0, 0x08, 0x23, 0x64, 0xd5, 0x82, 0x9f, 0x87, 0xd5, 0xa9, 0x3a, 0x54, 0xd3,
	0x81, 0x50, 0x67, 0x36, 0x42, 0x8c, 0x62, 0x60, 0x5c, 0x49, 0xc6, 0xd3, 0x14, 0x0d, 0x37, 0x12,
	0x53, 0x5d, 0xa3, 0x97, 0xa7, 0x7e, 0x17, 0xfa, 0x2b, 0x20, 0x4c, 0x4d, 0xbb, 0x34, 0x95, 0x66,
	0xd6, 0x15, 0x1c, 0x8a, 0x5d, 0x18, 0x67, 0xa1, 0x3a, 0xe9, 0x34, 0x3b, 0x0e, 0x16, 0x3c, 0x96,
	0x01, 0x37, 0x50, 0xa3, 0xbd, 0x3f, 0x50, 0xd0, 0x90, 0x16, 0xe3, 0x0a, 0xbd, 0x15, 0xeb, 0xdc,
	0xa2, 0x4c, 0x83, 0xc5, 0xca, 0xfc, 0x11, 0x18, 0x3d, 0x80, 0xa7, 0xf6, 0xac, 0xd5, 0x4e, 0x79,
	0x02, 0x5a, 0x71, 0x01, 0x0e, 0xe9, 0x92, 0xb9, 0xf6, 0xe0, 0xf8, 0xc2, 0xbe, 0x69, 0x9d, 0x16,
	0x98, 0x1a, 0x58, 0x33, 0x4b, 0xab, 0xb0, 0xae, 0x9d, 0x66, 0xb7, 0x35, 0xd7, 0xee, 0x9f, 0xff,
	0xf0, 0x63, 0xd8, 0xfa, 0xef, 0x35, 0x69, 0x9e, 0x22, 0x87, 0x93, 0x43, 0x06, 0x33, 0x35, 0xf3,
	0x1e, 0xac, 0xeb, 0xde, 0x93, 0x09, 0x2f, 0xad, 0xec, 0x45, 0x6b, 0xe6, 0xf8, 0xcf, 0xb5, 0x43,
	0xba, 0xad, 0xb9, 0x99, 0x85, 0xae, 0x37, 0xed, 0xa5, 0xbc, 0x07, 0x55, 0x84, 0xfd, 0xd6, 0x70,
	0x83, 0x0c, 0xac, 0xf0, 0x48, 0x6a, 0xe1, 0x5d, 0xd3, 0x6a, 0xdf, 0xa9, 0xde, 0xfb, 0xf6, 0xc3,
	0xbb, 0xf6, 0x4f, 0x62, 0x9d, 0x19, 0xf7, 0xb2, 0xaf, 0x4e, 0xd7, 0x9c, 0x48, 0xa0, 0xf3, 0xef,
	0x64, 0xad, 0x7a, 0x2f, 0xc8, 0xf6, 0xfb, 0x26, 0x29, 0x3f, 0x3a, 0xfd, 0xf0, 0xc6, 0x9a, 0x70,
	0x03, 0x08, 0x79, 0x1e, 0x1b, 0x97, 0x0b, 0x01, 0x5a, 0xcf, 0x57, 0x77, 0x02, 0x53, 0x8d, 0x31,
	0xb8, 0x3e, 0xd7, 0x52, 0xb8, 0x09, 0xa6, 0xd2, 0x60, 0x36, 0x5f, 0x80, 0x76, 0x25, 0xba, 0x11,
	0xa4, 0x90, 0xf1, 0xd8, 0xcd, 0x80, 0x07, 0x9b, 0x5f, 0xbe, 0xbd, 0x69, 0x5e, 0xef, 0xb1, 0xba,
	0xc2, 0xec, 0x28, 0x76, 0xcd, 0x9e, 0x1d, 0xcd, 0x1b, 0x27, 0xca, 0xa2, 0x59, 0x28, 0xd3, 0xe0,
	0x16, 0xb9, 0xd2, 0x71, 0x87, 0x5b, 0xa4, 0xf5, 0x79, 0x8b, 0x5c, 0xfc, 0x4b, 0x76, 0x9b, 0x9f,
	0x9c, 0xe6, 0x59, 0xd2, 0x7f, 0x49, 0x76, 0xf6, 0x68, 0x63, 0x77, 0x8f, 0x36, 0xf6, 0xf7, 0x28,
	0x79, 0x5e, 0x52, 0xf2, 0xb6, 0xa4, 0x64, 0xbb, 0xa4, 0x64, 0xa7, 0xa4, 0xe4, 0x6b, 0x49, 0xc9,
	0xf7, 0x92, 0x36, 0xf6, 0x4b, 0x4a, 0x5e, 0x8d, 0x68, 0x63, 0x38, 0xa2, 0x64, 0x67, 0x44, 0x1b,
	0xbb, 0x23, 0xda, 0x78, 0x7c, 0x3f, 0x42, 0xb5, 0x1a, 0x79, 0x05, 0xfe, 0x56, 0xce, 0xb8, 0x97,
	0x6b, 0x56, 0x0d, 0x21, 0x66, 0x89, 0xab, 0x32, 0x2c, 0x64, 0x00, 0x99, 0x7b, 0x08, 0x33, 0xe5,
	0x47, 0xc8, 0x60, 0xcd, 0xd4, 0x3d, 0x9c, 0xe8, 0xbb, 0xff, 0x7f, 0xd5, 0xc5, 0x6b, 0xbf, 0x02,
	0x00, 0x00, 0xff, 0xff, 0x9d, 0xd1, 0xd3, 0x29, 0xd1, 0x03, 0x00, 0x00,
}

func (this *FindFilterSetsReq) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*FindFilterSetsReq)
	if !ok {
		that2, ok := that.(FindFilterSetsReq)
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
	if len(this.ContextKeys) != len(that1.ContextKeys) {
		return false
	}
	for i := range this.ContextKeys {
		if this.ContextKeys[i] != that1.ContextKeys[i] {
			return false
		}
	}
	return true
}
func (this *FindFilterSetsRsp) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*FindFilterSetsRsp)
	if !ok {
		that2, ok := that.(FindFilterSetsRsp)
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
	if len(this.FilterSets) != len(that1.FilterSets) {
		return false
	}
	for i := range this.FilterSets {
		if !this.FilterSets[i].Equal(that1.FilterSets[i]) {
			return false
		}
	}
	return true
}
func (this *FindFilterSetsReq) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&filter_set.FindFilterSetsReq{")
	s = append(s, "Namespace: "+fmt.Sprintf("%#v", this.Namespace)+",\n")
	s = append(s, "ContextKeys: "+fmt.Sprintf("%#v", this.ContextKeys)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *FindFilterSetsRsp) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&filter_set.FindFilterSetsRsp{")
	if this.FilterSets != nil {
		s = append(s, "FilterSets: "+fmt.Sprintf("%#v", this.FilterSets)+",\n")
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
	// Find Filter Sets
	//
	// x-displayName: "Find Filter Sets for 1 or More Context Keys"
	// Retrieve any saved filter sets that are applicable for the given context key(s)
	FindFilterSets(ctx context.Context, in *FindFilterSetsReq, opts ...grpc.CallOption) (*FindFilterSetsRsp, error)
}

type customAPIClient struct {
	cc *grpc.ClientConn
}

func NewCustomAPIClient(cc *grpc.ClientConn) CustomAPIClient {
	return &customAPIClient{cc}
}

func (c *customAPIClient) FindFilterSets(ctx context.Context, in *FindFilterSetsReq, opts ...grpc.CallOption) (*FindFilterSetsRsp, error) {
	out := new(FindFilterSetsRsp)
	err := c.cc.Invoke(ctx, "/ves.io.schema.filter_set.CustomAPI/FindFilterSets", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CustomAPIServer is the server API for CustomAPI service.
type CustomAPIServer interface {
	// Find Filter Sets
	//
	// x-displayName: "Find Filter Sets for 1 or More Context Keys"
	// Retrieve any saved filter sets that are applicable for the given context key(s)
	FindFilterSets(context.Context, *FindFilterSetsReq) (*FindFilterSetsRsp, error)
}

// UnimplementedCustomAPIServer can be embedded to have forward compatible implementations.
type UnimplementedCustomAPIServer struct {
}

func (*UnimplementedCustomAPIServer) FindFilterSets(ctx context.Context, req *FindFilterSetsReq) (*FindFilterSetsRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindFilterSets not implemented")
}

func RegisterCustomAPIServer(s *grpc.Server, srv CustomAPIServer) {
	s.RegisterService(&_CustomAPI_serviceDesc, srv)
}

func _CustomAPI_FindFilterSets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindFilterSetsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomAPIServer).FindFilterSets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ves.io.schema.filter_set.CustomAPI/FindFilterSets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomAPIServer).FindFilterSets(ctx, req.(*FindFilterSetsReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _CustomAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ves.io.schema.filter_set.CustomAPI",
	HandlerType: (*CustomAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FindFilterSets",
			Handler:    _CustomAPI_FindFilterSets_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ves.io/schema/filter_set/public_customapi.proto",
}

func (m *FindFilterSetsReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FindFilterSetsReq) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *FindFilterSetsReq) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ContextKeys) > 0 {
		for iNdEx := len(m.ContextKeys) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.ContextKeys[iNdEx])
			copy(dAtA[i:], m.ContextKeys[iNdEx])
			i = encodeVarintPublicCustomapi(dAtA, i, uint64(len(m.ContextKeys[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
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

func (m *FindFilterSetsRsp) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FindFilterSetsRsp) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *FindFilterSetsRsp) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.FilterSets) > 0 {
		for iNdEx := len(m.FilterSets) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.FilterSets[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintPublicCustomapi(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
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
func (m *FindFilterSetsReq) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Namespace)
	if l > 0 {
		n += 1 + l + sovPublicCustomapi(uint64(l))
	}
	if len(m.ContextKeys) > 0 {
		for _, s := range m.ContextKeys {
			l = len(s)
			n += 1 + l + sovPublicCustomapi(uint64(l))
		}
	}
	return n
}

func (m *FindFilterSetsRsp) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.FilterSets) > 0 {
		for _, e := range m.FilterSets {
			l = e.Size()
			n += 1 + l + sovPublicCustomapi(uint64(l))
		}
	}
	return n
}

func sovPublicCustomapi(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPublicCustomapi(x uint64) (n int) {
	return sovPublicCustomapi(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *FindFilterSetsReq) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&FindFilterSetsReq{`,
		`Namespace:` + fmt.Sprintf("%v", this.Namespace) + `,`,
		`ContextKeys:` + fmt.Sprintf("%v", this.ContextKeys) + `,`,
		`}`,
	}, "")
	return s
}
func (this *FindFilterSetsRsp) String() string {
	if this == nil {
		return "nil"
	}
	repeatedStringForFilterSets := "[]*Object{"
	for _, f := range this.FilterSets {
		repeatedStringForFilterSets += strings.Replace(fmt.Sprintf("%v", f), "Object", "Object", 1) + ","
	}
	repeatedStringForFilterSets += "}"
	s := strings.Join([]string{`&FindFilterSetsRsp{`,
		`FilterSets:` + repeatedStringForFilterSets + `,`,
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
func (m *FindFilterSetsReq) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: FindFilterSetsReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FindFilterSetsReq: illegal tag %d (wire type %d)", fieldNum, wire)
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
				return fmt.Errorf("proto: wrong wireType = %d for field ContextKeys", wireType)
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
			m.ContextKeys = append(m.ContextKeys, string(dAtA[iNdEx:postIndex]))
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
func (m *FindFilterSetsRsp) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: FindFilterSetsRsp: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FindFilterSetsRsp: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FilterSets", wireType)
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
			m.FilterSets = append(m.FilterSets, &Object{})
			if err := m.FilterSets[len(m.FilterSets)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
