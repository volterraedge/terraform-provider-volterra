// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ves.io/schema/dns_domain/public_customapi.proto

// Verify Dns Domain
//
// x-displayName: "DNS Domain"
// Verify DNS Domain is used to trigger verification DNS domain on demand

package dns_domain

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

// Verify DNS Domain Request
//
// x-displayName: "Verify DNS Domain"
// DNS domain request verification shape
type VerifyDnsDomainRequest struct {
	// Namespace
	//
	// x-displayName: "Namespace"
	// x-example: "system"
	// Namespace is always system for dns_domain
	Namespace string `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	// Name
	//
	// x-displayName: "Name"
	// x-example: "example.com"
	// Name dns_domain object, which also domain name
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (m *VerifyDnsDomainRequest) Reset()      { *m = VerifyDnsDomainRequest{} }
func (*VerifyDnsDomainRequest) ProtoMessage() {}
func (*VerifyDnsDomainRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_fc937e91d9921f51, []int{0}
}
func (m *VerifyDnsDomainRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *VerifyDnsDomainRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_VerifyDnsDomainRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *VerifyDnsDomainRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VerifyDnsDomainRequest.Merge(m, src)
}
func (m *VerifyDnsDomainRequest) XXX_Size() int {
	return m.Size()
}
func (m *VerifyDnsDomainRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_VerifyDnsDomainRequest.DiscardUnknown(m)
}

var xxx_messageInfo_VerifyDnsDomainRequest proto.InternalMessageInfo

func (m *VerifyDnsDomainRequest) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *VerifyDnsDomainRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// Verify Response
//
// x-displayName: "Verify Response"
// Verify DNS domain Response
type VerifyDnsDomainResponse struct {
}

func (m *VerifyDnsDomainResponse) Reset()      { *m = VerifyDnsDomainResponse{} }
func (*VerifyDnsDomainResponse) ProtoMessage() {}
func (*VerifyDnsDomainResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_fc937e91d9921f51, []int{1}
}
func (m *VerifyDnsDomainResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *VerifyDnsDomainResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_VerifyDnsDomainResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *VerifyDnsDomainResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VerifyDnsDomainResponse.Merge(m, src)
}
func (m *VerifyDnsDomainResponse) XXX_Size() int {
	return m.Size()
}
func (m *VerifyDnsDomainResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_VerifyDnsDomainResponse.DiscardUnknown(m)
}

var xxx_messageInfo_VerifyDnsDomainResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*VerifyDnsDomainRequest)(nil), "ves.io.schema.dns_domain.VerifyDnsDomainRequest")
	golang_proto.RegisterType((*VerifyDnsDomainRequest)(nil), "ves.io.schema.dns_domain.VerifyDnsDomainRequest")
	proto.RegisterType((*VerifyDnsDomainResponse)(nil), "ves.io.schema.dns_domain.VerifyDnsDomainResponse")
	golang_proto.RegisterType((*VerifyDnsDomainResponse)(nil), "ves.io.schema.dns_domain.VerifyDnsDomainResponse")
}

func init() {
	proto.RegisterFile("ves.io/schema/dns_domain/public_customapi.proto", fileDescriptor_fc937e91d9921f51)
}
func init() {
	golang_proto.RegisterFile("ves.io/schema/dns_domain/public_customapi.proto", fileDescriptor_fc937e91d9921f51)
}

var fileDescriptor_fc937e91d9921f51 = []byte{
	// 496 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0xbf, 0x6f, 0xd3, 0x40,
	0x14, 0xc7, 0x7d, 0x01, 0x21, 0xc5, 0x0b, 0xc8, 0x03, 0x84, 0x10, 0x1d, 0x55, 0x26, 0x84, 0xb0,
	0x8f, 0x1f, 0x42, 0x48, 0x88, 0x05, 0xe8, 0x02, 0x03, 0xa0, 0x0e, 0x0c, 0x30, 0x54, 0xe7, 0xf3,
	0xb3, 0x73, 0x10, 0xdf, 0x3b, 0xee, 0xce, 0x56, 0x11, 0xaa, 0x84, 0x3a, 0x31, 0x22, 0xb1, 0xb1,
	0xb1, 0xf1, 0x3f, 0x74, 0xe9, 0x06, 0x13, 0x8a, 0x60, 0xa0, 0x23, 0x71, 0x18, 0x18, 0xfb, 0x27,
	0x20, 0x2e, 0xa1, 0x69, 0x52, 0x3a, 0xb0, 0xbd, 0x77, 0x9f, 0xf7, 0x7d, 0xfa, 0xfa, 0xe9, 0xeb,
	0x90, 0xd5, 0x60, 0x13, 0x89, 0xcc, 0x8a, 0x01, 0x94, 0x9c, 0x65, 0xca, 0xae, 0x67, 0x58, 0x72,
	0xa9, 0x98, 0xae, 0xd2, 0xa1, 0x14, 0xeb, 0xa2, 0xb2, 0x0e, 0x4b, 0xae, 0x65, 0xa2, 0x0d, 0x3a,
	0x8c, 0x3a, 0x53, 0x41, 0x32, 0x15, 0x24, 0x73, 0x41, 0x37, 0x2e, 0xa4, 0x1b, 0x54, 0x69, 0x22,
	0xb0, 0x64, 0x05, 0x16, 0xc8, 0xbc, 0x20, 0xad, 0x72, 0xdf, 0xf9, 0xc6, 0x57, 0xd3, 0x45, 0xdd,
	0x5e, 0x81, 0x58, 0x0c, 0x81, 0x71, 0x2d, 0x19, 0x57, 0x0a, 0x1d, 0x77, 0x12, 0x95, 0x9d, 0xd1,
	0x73, 0x8b, 0xbe, 0x50, 0x1f, 0x84, 0xbd, 0x45, 0x58, 0xf3, 0xa1, 0xcc, 0xb8, 0x83, 0x19, 0xed,
	0x2f, 0x51, 0xb0, 0xa0, 0xea, 0xc5, 0x0d, 0xfd, 0xfb, 0xe1, 0xe9, 0xc7, 0x60, 0x64, 0xfe, 0x72,
	0x55, 0xd9, 0x55, 0x6f, 0x7f, 0x0d, 0x5e, 0x54, 0x60, 0x5d, 0xd4, 0x0b, 0xdb, 0x8a, 0x97, 0x60,
	0x35, 0x17, 0xd0, 0x21, 0x2b, 0xe4, 0x42, 0x7b, 0x6d, 0xfe, 0x10, 0x45, 0xe1, 0xf1, 0x3f, 0x4d,
	0xa7, 0xe5, 0x81, 0xaf, 0xfb, 0x67, 0xc3, 0x33, 0x87, 0x76, 0x59, 0x8d, 0xca, 0xc2, 0xd5, 0xef,
	0xad, 0xb0, 0x7d, 0xd7, 0x1f, 0xf0, 0xf6, 0xa3, 0x7b, 0xd1, 0xfb, 0x56, 0x78, 0x72, 0x69, 0x32,
	0xba, 0x9c, 0x1c, 0x75, 0xcf, 0xe4, 0xdf, 0x06, 0xbb, 0x57, 0xfe, 0x43, 0x31, 0xb5, 0xd1, 0xff,
	0x40, 0x9a, 0x4f, 0x9d, 0xa7, 0xf9, 0xf5, 0x0d, 0x11, 0x5b, 0x10, 0x95, 0x81, 0x12, 0xec, 0x20,
	0x4e, 0xb9, 0x95, 0x22, 0x2e, 0x51, 0x49, 0x87, 0xe6, 0xd2, 0xca, 0x32, 0xb6, 0x8e, 0xab, 0x8c,
	0x9b, 0x6c, 0x3e, 0x51, 0x83, 0x8d, 0x25, 0xc6, 0x52, 0xe5, 0x86, 0x5b, 0x67, 0x2a, 0xe1, 0x2a,
	0x03, 0xb1, 0x01, 0x9e, 0x6d, 0x7d, 0xfb, 0xf9, 0xae, 0x75, 0xab, 0x7f, 0x63, 0x16, 0x1a, 0xb6,
	0x7f, 0x30, 0xcb, 0x5e, 0xed, 0xd7, 0x9b, 0x07, 0xd3, 0xe5, 0x9f, 0x37, 0x59, 0xed, 0x1d, 0xdf,
	0x24, 0x17, 0xbb, 0xf1, 0xce, 0x36, 0x39, 0xf6, 0x75, 0x9b, 0x9c, 0x3f, 0xf2, 0xeb, 0x1e, 0xa6,
	0xcf, 0x40, 0xb8, 0xad, 0x2f, 0x9d, 0xd6, 0x29, 0x72, 0xe7, 0x0d, 0x19, 0x8d, 0x69, 0xb0, 0x3b,
	0xa6, 0xc1, 0xde, 0x98, 0x92, 0xd7, 0x0d, 0x25, 0x1f, 0x1b, 0x4a, 0x3e, 0x37, 0x94, 0x8c, 0x1a,
	0x4a, 0x7e, 0x34, 0x94, 0xfc, 0x6a, 0x68, 0xb0, 0xd7, 0x50, 0xf2, 0x76, 0x42, 0x83, 0x9d, 0x09,
	0x25, 0xa3, 0x09, 0x0d, 0x76, 0x27, 0x34, 0x78, 0xf2, 0xa0, 0x40, 0xfd, 0xbc, 0x48, 0x6a, 0x1c,
	0x3a, 0x30, 0x86, 0x27, 0x95, 0x65, 0xbe, 0xc8, 0xd1, 0x94, 0xb1, 0x36, 0x58, 0xcb, 0x0c, 0x4c,
	0xfc, 0x17, 0x33, 0x9d, 0x16, 0xc8, 0x60, 0xc3, 0xcd, 0xe2, 0x74, 0xe8, 0x47, 0x49, 0x4f, 0xf8,
	0x48, 0x5d, 0xfb, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x33, 0x20, 0x4d, 0x91, 0x4b, 0x03, 0x00, 0x00,
}

func (this *VerifyDnsDomainRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*VerifyDnsDomainRequest)
	if !ok {
		that2, ok := that.(VerifyDnsDomainRequest)
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
	return true
}
func (this *VerifyDnsDomainResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*VerifyDnsDomainResponse)
	if !ok {
		that2, ok := that.(VerifyDnsDomainResponse)
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
func (this *VerifyDnsDomainRequest) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&dns_domain.VerifyDnsDomainRequest{")
	s = append(s, "Namespace: "+fmt.Sprintf("%#v", this.Namespace)+",\n")
	s = append(s, "Name: "+fmt.Sprintf("%#v", this.Name)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *VerifyDnsDomainResponse) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 4)
	s = append(s, "&dns_domain.VerifyDnsDomainResponse{")
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
	// Verify
	//
	// x-displayName: "Verify Dns Domain"
	// Verify DNS Domain for a given dns_domain object
	VerifyDnsDomain(ctx context.Context, in *VerifyDnsDomainRequest, opts ...grpc.CallOption) (*VerifyDnsDomainResponse, error)
}

type customAPIClient struct {
	cc *grpc.ClientConn
}

func NewCustomAPIClient(cc *grpc.ClientConn) CustomAPIClient {
	return &customAPIClient{cc}
}

func (c *customAPIClient) VerifyDnsDomain(ctx context.Context, in *VerifyDnsDomainRequest, opts ...grpc.CallOption) (*VerifyDnsDomainResponse, error) {
	out := new(VerifyDnsDomainResponse)
	err := c.cc.Invoke(ctx, "/ves.io.schema.dns_domain.CustomAPI/VerifyDnsDomain", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CustomAPIServer is the server API for CustomAPI service.
type CustomAPIServer interface {
	// Verify
	//
	// x-displayName: "Verify Dns Domain"
	// Verify DNS Domain for a given dns_domain object
	VerifyDnsDomain(context.Context, *VerifyDnsDomainRequest) (*VerifyDnsDomainResponse, error)
}

// UnimplementedCustomAPIServer can be embedded to have forward compatible implementations.
type UnimplementedCustomAPIServer struct {
}

func (*UnimplementedCustomAPIServer) VerifyDnsDomain(ctx context.Context, req *VerifyDnsDomainRequest) (*VerifyDnsDomainResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyDnsDomain not implemented")
}

func RegisterCustomAPIServer(s *grpc.Server, srv CustomAPIServer) {
	s.RegisterService(&_CustomAPI_serviceDesc, srv)
}

func _CustomAPI_VerifyDnsDomain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyDnsDomainRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomAPIServer).VerifyDnsDomain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ves.io.schema.dns_domain.CustomAPI/VerifyDnsDomain",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomAPIServer).VerifyDnsDomain(ctx, req.(*VerifyDnsDomainRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CustomAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ves.io.schema.dns_domain.CustomAPI",
	HandlerType: (*CustomAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "VerifyDnsDomain",
			Handler:    _CustomAPI_VerifyDnsDomain_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ves.io/schema/dns_domain/public_customapi.proto",
}

func (m *VerifyDnsDomainRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *VerifyDnsDomainRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *VerifyDnsDomainRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
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

func (m *VerifyDnsDomainResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *VerifyDnsDomainResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *VerifyDnsDomainResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
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
func (m *VerifyDnsDomainRequest) Size() (n int) {
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
	return n
}

func (m *VerifyDnsDomainResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovPublicCustomapi(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPublicCustomapi(x uint64) (n int) {
	return sovPublicCustomapi(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *VerifyDnsDomainRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&VerifyDnsDomainRequest{`,
		`Namespace:` + fmt.Sprintf("%v", this.Namespace) + `,`,
		`Name:` + fmt.Sprintf("%v", this.Name) + `,`,
		`}`,
	}, "")
	return s
}
func (this *VerifyDnsDomainResponse) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&VerifyDnsDomainResponse{`,
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
func (m *VerifyDnsDomainRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: VerifyDnsDomainRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: VerifyDnsDomainRequest: illegal tag %d (wire type %d)", fieldNum, wire)
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
func (m *VerifyDnsDomainResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: VerifyDnsDomainResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: VerifyDnsDomainResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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
