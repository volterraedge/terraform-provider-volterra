// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ves.io/schema/uztna/uztna_active_sessions/public_uztna_terminate_sessions.proto

// manage active sessions
//
// x-displayName: "Manage Active Sessions"
// APIs to monitor UZTNA active sessions on all applications.

package uztna_active_sessions

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/googleapis/google/api"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
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

// TerminateSessionsRequest
//
// x-displayName: "Terminate Access Session Request"
// Request structure for Terminate Access Session
type TerminateSessionsRequest struct {
	// namespace
	//
	// x-displayName: "Namespace"
	// x-example: "ns1"
	// Namespace of the App type for the current request
	Namespace string `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	// session_id
	//
	// x-displayName: "Session ID"
	// x-required
	// x-example: "9d6d591"
	// ID of the session
	SessionId string `protobuf:"bytes,2,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
}

func (m *TerminateSessionsRequest) Reset()      { *m = TerminateSessionsRequest{} }
func (*TerminateSessionsRequest) ProtoMessage() {}
func (*TerminateSessionsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b0d7c0fdffabdce, []int{0}
}
func (m *TerminateSessionsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TerminateSessionsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TerminateSessionsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TerminateSessionsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TerminateSessionsRequest.Merge(m, src)
}
func (m *TerminateSessionsRequest) XXX_Size() int {
	return m.Size()
}
func (m *TerminateSessionsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TerminateSessionsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TerminateSessionsRequest proto.InternalMessageInfo

func (m *TerminateSessionsRequest) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *TerminateSessionsRequest) GetSessionId() string {
	if m != nil {
		return m.SessionId
	}
	return ""
}

func init() {
	proto.RegisterType((*TerminateSessionsRequest)(nil), "ves.io.schema.uztna.uztna_active_sessions.TerminateSessionsRequest")
	golang_proto.RegisterType((*TerminateSessionsRequest)(nil), "ves.io.schema.uztna.uztna_active_sessions.TerminateSessionsRequest")
}

func init() {
	proto.RegisterFile("ves.io/schema/uztna/uztna_active_sessions/public_uztna_terminate_sessions.proto", fileDescriptor_9b0d7c0fdffabdce)
}
func init() {
	golang_proto.RegisterFile("ves.io/schema/uztna/uztna_active_sessions/public_uztna_terminate_sessions.proto", fileDescriptor_9b0d7c0fdffabdce)
}

var fileDescriptor_9b0d7c0fdffabdce = []byte{
	// 592 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0xcf, 0x4e, 0xd5, 0x4e,
	0x14, 0xee, 0xdc, 0xfb, 0xfb, 0x91, 0xd0, 0x9d, 0x8d, 0x31, 0xcd, 0x95, 0x4c, 0x90, 0x15, 0x97,
	0x64, 0x66, 0x88, 0xc6, 0x85, 0x2b, 0x22, 0xc6, 0x18, 0x56, 0x10, 0x74, 0x25, 0x21, 0xcd, 0xb4,
	0x3d, 0x94, 0x91, 0xb6, 0x53, 0x3b, 0xd3, 0x8a, 0x10, 0x12, 0xc3, 0x0b, 0x68, 0xe2, 0xca, 0x37,
	0xf0, 0x05, 0xdc, 0xc8, 0x86, 0xb8, 0x41, 0x37, 0x86, 0xe8, 0x86, 0xa5, 0xf4, 0xba, 0x40, 0x57,
	0x3c, 0x82, 0xb9, 0xfd, 0x73, 0xaf, 0xc0, 0x55, 0xd9, 0x34, 0x67, 0xce, 0x77, 0xce, 0x77, 0xce,
	0x7c, 0xe7, 0x4c, 0xcd, 0xc5, 0x1c, 0x14, 0x15, 0x92, 0x29, 0x6f, 0x1d, 0x22, 0xce, 0xb2, 0x2d,
	0x1d, 0xd7, 0x5f, 0x87, 0x7b, 0x5a, 0xe4, 0xe0, 0x28, 0x50, 0x4a, 0xc8, 0x58, 0xb1, 0x24, 0x73,
	0x43, 0xe1, 0x39, 0x15, 0xa8, 0x21, 0x8d, 0x44, 0xcc, 0xf5, 0x10, 0xa7, 0x49, 0x2a, 0xb5, 0xb4,
	0xba, 0x15, 0x21, 0xad, 0x08, 0x69, 0x19, 0x4d, 0x47, 0x12, 0x76, 0x48, 0x20, 0xf4, 0x7a, 0xe6,
	0x52, 0x4f, 0x46, 0x2c, 0x90, 0x81, 0x64, 0x25, 0x83, 0x9b, 0xad, 0x95, 0xa7, 0xf2, 0x50, 0x5a,
	0x15, 0x73, 0x67, 0x22, 0x90, 0x32, 0x08, 0x81, 0xf1, 0x44, 0x30, 0x1e, 0xc7, 0x52, 0x73, 0x3d,
	0xac, 0xdb, 0xb9, 0x5e, 0xa3, 0x03, 0x0e, 0x88, 0x12, 0xfd, 0xbc, 0x01, 0xcf, 0xde, 0x52, 0x26,
	0xbf, 0x67, 0x4e, 0x9c, 0x05, 0x73, 0x1e, 0x0a, 0x9f, 0x6b, 0xa8, 0xd1, 0xa9, 0x73, 0x28, 0x28,
	0x88, 0xf3, 0x73, 0x0c, 0x93, 0xe7, 0x62, 0x04, 0x3c, 0x73, 0xce, 0x44, 0x4c, 0x79, 0xa6, 0xfd,
	0xa8, 0x51, 0xec, 0x61, 0x7d, 0xff, 0x65, 0x78, 0x9a, 0x81, 0xd2, 0xd6, 0x84, 0x39, 0x1e, 0xf3,
	0x08, 0x54, 0xc2, 0x3d, 0xb0, 0xd1, 0x24, 0x9a, 0x1e, 0x5f, 0x1e, 0x3a, 0xac, 0xae, 0x69, 0xd6,
	0x82, 0x39, 0xc2, 0xb7, 0x5b, 0x7d, 0x78, 0xde, 0x7c, 0xff, 0x63, 0xbf, 0xfd, 0x7f, 0xda, 0x3e,
	0x42, 0x68, 0x79, 0xbc, 0x46, 0x17, 0xfc, 0x9b, 0x2f, 0xff, 0x33, 0xaf, 0x5e, 0xa8, 0x72, 0x77,
	0x69, 0xc1, 0x7a, 0xd7, 0x36, 0xaf, 0x5c, 0x00, 0xac, 0x7b, 0xf4, 0xd2, 0xa3, 0xa2, 0x7f, 0x6a,
	0xbe, 0x73, 0x8d, 0x56, 0xba, 0xd3, 0x46, 0x77, 0x7a, 0xbf, 0xaf, 0xfb, 0xd4, 0x87, 0x56, 0x71,
	0x60, 0xdf, 0x58, 0xbb, 0xbd, 0xe9, 0x91, 0x2c, 0x16, 0x39, 0xa4, 0x8a, 0x87, 0xa4, 0x4f, 0x4d,
	0x94, 0xe6, 0xb1, 0xcf, 0x53, 0x9f, 0x64, 0x0a, 0xd2, 0x9f, 0x07, 0xf6, 0x27, 0x64, 0x2e, 0xe5,
	0xa0, 0x88, 0x90, 0xa4, 0x6a, 0x84, 0x94, 0x2d, 0xd4, 0xdf, 0xaa, 0x11, 0xd2, 0x34, 0x42, 0x06,
	0x7b, 0xd7, 0x78, 0x78, 0x22, 0x2e, 0x3a, 0xad, 0xcc, 0x1c, 0xf3, 0x21, 0x04, 0x0d, 0xd6, 0x46,
	0xbd, 0xb7, 0x6c, 0x20, 0xa9, 0x62, 0xd3, 0x2b, 0x9c, 0x6c, 0xad, 0x4e, 0xaf, 0x10, 0x4e, 0xb6,
	0x66, 0xc9, 0x9d, 0xd5, 0x99, 0x95, 0xda, 0xe8, 0xce, 0x75, 0xeb, 0x27, 0x50, 0x65, 0x3b, 0xdc,
	0xf3, 0x40, 0xa9, 0x46, 0x87, 0x2a, 0xb1, 0x1f, 0x38, 0x3a, 0x77, 0xf7, 0xeb, 0xf7, 0xd7, 0xad,
	0x85, 0x99, 0x07, 0x23, 0x4a, 0x6e, 0x0f, 0xec, 0x9d, 0xbf, 0x56, 0xd8, 0x1e, 0x0e, 0x7b, 0xa7,
	0x33, 0xb7, 0xbf, 0x87, 0xda, 0x5f, 0xf6, 0xd0, 0xec, 0xe5, 0x07, 0xb5, 0xe8, 0x3e, 0x01, 0x4f,
	0xef, 0x7e, 0xb6, 0xdb, 0x27, 0x2d, 0x34, 0xff, 0x06, 0x1d, 0x1e, 0x63, 0xe3, 0xe8, 0x18, 0x1b,
	0xa7, 0xc7, 0x18, 0xbd, 0x28, 0x30, 0x7a, 0x5b, 0x60, 0xf4, 0xb1, 0xc0, 0xe8, 0xb0, 0xc0, 0xe8,
	0x5b, 0x81, 0xd1, 0x49, 0x81, 0x8d, 0xd3, 0x02, 0xa3, 0x57, 0x3d, 0x6c, 0xec, 0xf7, 0x30, 0x3a,
	0xec, 0x61, 0xe3, 0xa8, 0x87, 0x8d, 0xc7, 0x4e, 0x20, 0x93, 0x8d, 0x80, 0xe6, 0x32, 0xd4, 0x90,
	0xa6, 0x9c, 0x66, 0x8a, 0x95, 0xc6, 0x9a, 0x4c, 0x23, 0x92, 0xa4, 0x32, 0x17, 0x3e, 0xa4, 0xa4,
	0x81, 0x59, 0xe2, 0x06, 0x92, 0xc1, 0xa6, 0xae, 0x9f, 0xc0, 0xbf, 0x7f, 0x27, 0xee, 0x58, 0xb9,
	0x31, 0xb7, 0x7e, 0x05, 0x00, 0x00, 0xff, 0xff, 0xff, 0x86, 0xdf, 0xc7, 0x82, 0x04, 0x00, 0x00,
}

func (this *TerminateSessionsRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*TerminateSessionsRequest)
	if !ok {
		that2, ok := that.(TerminateSessionsRequest)
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
	if this.SessionId != that1.SessionId {
		return false
	}
	return true
}
func (this *TerminateSessionsRequest) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&uztna_active_sessions.TerminateSessionsRequest{")
	s = append(s, "Namespace: "+fmt.Sprintf("%#v", this.Namespace)+",\n")
	s = append(s, "SessionId: "+fmt.Sprintf("%#v", this.SessionId)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringPublicUztnaTerminateSessions(v interface{}, typ string) string {
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

// TerminateSessionsAPIClient is the client API for TerminateSessionsAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TerminateSessionsAPIClient interface {
	// TerminateSessions
	//
	// x-displayName: "Terminate Access Sessions"
	// Terminate uztna access sessions
	TerminateSessions(ctx context.Context, in *TerminateSessionsRequest, opts ...grpc.CallOption) (*types.Empty, error)
}

type terminateSessionsAPIClient struct {
	cc *grpc.ClientConn
}

func NewTerminateSessionsAPIClient(cc *grpc.ClientConn) TerminateSessionsAPIClient {
	return &terminateSessionsAPIClient{cc}
}

func (c *terminateSessionsAPIClient) TerminateSessions(ctx context.Context, in *TerminateSessionsRequest, opts ...grpc.CallOption) (*types.Empty, error) {
	out := new(types.Empty)
	err := c.cc.Invoke(ctx, "/ves.io.schema.uztna.uztna_active_sessions.TerminateSessionsAPI/TerminateSessions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TerminateSessionsAPIServer is the server API for TerminateSessionsAPI service.
type TerminateSessionsAPIServer interface {
	// TerminateSessions
	//
	// x-displayName: "Terminate Access Sessions"
	// Terminate uztna access sessions
	TerminateSessions(context.Context, *TerminateSessionsRequest) (*types.Empty, error)
}

// UnimplementedTerminateSessionsAPIServer can be embedded to have forward compatible implementations.
type UnimplementedTerminateSessionsAPIServer struct {
}

func (*UnimplementedTerminateSessionsAPIServer) TerminateSessions(ctx context.Context, req *TerminateSessionsRequest) (*types.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TerminateSessions not implemented")
}

func RegisterTerminateSessionsAPIServer(s *grpc.Server, srv TerminateSessionsAPIServer) {
	s.RegisterService(&_TerminateSessionsAPI_serviceDesc, srv)
}

func _TerminateSessionsAPI_TerminateSessions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TerminateSessionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TerminateSessionsAPIServer).TerminateSessions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ves.io.schema.uztna.uztna_active_sessions.TerminateSessionsAPI/TerminateSessions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TerminateSessionsAPIServer).TerminateSessions(ctx, req.(*TerminateSessionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TerminateSessionsAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ves.io.schema.uztna.uztna_active_sessions.TerminateSessionsAPI",
	HandlerType: (*TerminateSessionsAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "TerminateSessions",
			Handler:    _TerminateSessionsAPI_TerminateSessions_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ves.io/schema/uztna/uztna_active_sessions/public_uztna_terminate_sessions.proto",
}

func (m *TerminateSessionsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TerminateSessionsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TerminateSessionsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.SessionId) > 0 {
		i -= len(m.SessionId)
		copy(dAtA[i:], m.SessionId)
		i = encodeVarintPublicUztnaTerminateSessions(dAtA, i, uint64(len(m.SessionId)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Namespace) > 0 {
		i -= len(m.Namespace)
		copy(dAtA[i:], m.Namespace)
		i = encodeVarintPublicUztnaTerminateSessions(dAtA, i, uint64(len(m.Namespace)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintPublicUztnaTerminateSessions(dAtA []byte, offset int, v uint64) int {
	offset -= sovPublicUztnaTerminateSessions(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *TerminateSessionsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Namespace)
	if l > 0 {
		n += 1 + l + sovPublicUztnaTerminateSessions(uint64(l))
	}
	l = len(m.SessionId)
	if l > 0 {
		n += 1 + l + sovPublicUztnaTerminateSessions(uint64(l))
	}
	return n
}

func sovPublicUztnaTerminateSessions(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPublicUztnaTerminateSessions(x uint64) (n int) {
	return sovPublicUztnaTerminateSessions(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *TerminateSessionsRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&TerminateSessionsRequest{`,
		`Namespace:` + fmt.Sprintf("%v", this.Namespace) + `,`,
		`SessionId:` + fmt.Sprintf("%v", this.SessionId) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringPublicUztnaTerminateSessions(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *TerminateSessionsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPublicUztnaTerminateSessions
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
			return fmt.Errorf("proto: TerminateSessionsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TerminateSessionsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Namespace", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPublicUztnaTerminateSessions
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
				return ErrInvalidLengthPublicUztnaTerminateSessions
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPublicUztnaTerminateSessions
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Namespace = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SessionId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPublicUztnaTerminateSessions
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
				return ErrInvalidLengthPublicUztnaTerminateSessions
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPublicUztnaTerminateSessions
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SessionId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPublicUztnaTerminateSessions(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPublicUztnaTerminateSessions
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthPublicUztnaTerminateSessions
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
func skipPublicUztnaTerminateSessions(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPublicUztnaTerminateSessions
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
					return 0, ErrIntOverflowPublicUztnaTerminateSessions
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
					return 0, ErrIntOverflowPublicUztnaTerminateSessions
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
				return 0, ErrInvalidLengthPublicUztnaTerminateSessions
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPublicUztnaTerminateSessions
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPublicUztnaTerminateSessions
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPublicUztnaTerminateSessions        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPublicUztnaTerminateSessions          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPublicUztnaTerminateSessions = fmt.Errorf("proto: unexpected end of group")
)
