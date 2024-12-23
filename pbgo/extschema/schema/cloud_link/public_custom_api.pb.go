// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ves.io/schema/cloud_link/public_custom_api.proto

// CloudLink
//
// x-displayName: "CloudLink"
// APIs to manage Cloud Link resources

package cloud_link

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

// CloudLink Reapply Config Request
//
// x-displayName: "CloudLink Reapply Config Request"
// Reapply CloudLink Config
type ReapplyConfigRequest struct {
	// Name
	//
	// x-displayName: "Name"
	// x-example: "aws-cloud-link-east"
	// Name of Cloud Link.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (m *ReapplyConfigRequest) Reset()      { *m = ReapplyConfigRequest{} }
func (*ReapplyConfigRequest) ProtoMessage() {}
func (*ReapplyConfigRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cc9d11a2d081fbcb, []int{0}
}
func (m *ReapplyConfigRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ReapplyConfigRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ReapplyConfigRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ReapplyConfigRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReapplyConfigRequest.Merge(m, src)
}
func (m *ReapplyConfigRequest) XXX_Size() int {
	return m.Size()
}
func (m *ReapplyConfigRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReapplyConfigRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReapplyConfigRequest proto.InternalMessageInfo

func (m *ReapplyConfigRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// CloudLink Reapply Config Response
//
// x-displayName: "CloudLink Reapply Config Response"
// Reapply CloudLink Config
type ReapplyConfigResponse struct {
}

func (m *ReapplyConfigResponse) Reset()      { *m = ReapplyConfigResponse{} }
func (*ReapplyConfigResponse) ProtoMessage() {}
func (*ReapplyConfigResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_cc9d11a2d081fbcb, []int{1}
}
func (m *ReapplyConfigResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ReapplyConfigResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ReapplyConfigResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ReapplyConfigResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReapplyConfigResponse.Merge(m, src)
}
func (m *ReapplyConfigResponse) XXX_Size() int {
	return m.Size()
}
func (m *ReapplyConfigResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ReapplyConfigResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ReapplyConfigResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ReapplyConfigRequest)(nil), "ves.io.schema.cloud_link.ReapplyConfigRequest")
	golang_proto.RegisterType((*ReapplyConfigRequest)(nil), "ves.io.schema.cloud_link.ReapplyConfigRequest")
	proto.RegisterType((*ReapplyConfigResponse)(nil), "ves.io.schema.cloud_link.ReapplyConfigResponse")
	golang_proto.RegisterType((*ReapplyConfigResponse)(nil), "ves.io.schema.cloud_link.ReapplyConfigResponse")
}

func init() {
	proto.RegisterFile("ves.io/schema/cloud_link/public_custom_api.proto", fileDescriptor_cc9d11a2d081fbcb)
}
func init() {
	golang_proto.RegisterFile("ves.io/schema/cloud_link/public_custom_api.proto", fileDescriptor_cc9d11a2d081fbcb)
}

var fileDescriptor_cc9d11a2d081fbcb = []byte{
	// 494 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0xb1, 0x8b, 0x13, 0x41,
	0x14, 0xc6, 0x77, 0x4e, 0x11, 0x5c, 0x38, 0x90, 0xa0, 0x18, 0xe3, 0x31, 0x27, 0xa9, 0xe4, 0x70,
	0x77, 0xe4, 0xc4, 0x46, 0x1b, 0xf5, 0x6c, 0x6c, 0x54, 0x52, 0xda, 0x84, 0xd9, 0xcd, 0xcb, 0x66,
	0xcc, 0xee, 0xbc, 0x71, 0xde, 0xec, 0xe6, 0x82, 0x0a, 0x72, 0x95, 0xa5, 0xe0, 0x3f, 0x61, 0xe7,
	0x1f, 0x70, 0xcd, 0x75, 0x5a, 0x49, 0xd0, 0xe6, 0x4a, 0xb3, 0xb1, 0xb0, 0x92, 0xfb, 0x13, 0xe4,
	0x76, 0x73, 0x98, 0x44, 0x0f, 0xec, 0xde, 0xf0, 0xfb, 0xbe, 0xd9, 0x6f, 0x1f, 0xdf, 0xf8, 0x37,
	0x0b, 0xa0, 0x50, 0xa1, 0xa0, 0x78, 0x00, 0x99, 0x14, 0x71, 0x8a, 0x79, 0xaf, 0x9b, 0x2a, 0x3d,
	0x14, 0x26, 0x8f, 0x52, 0x15, 0x77, 0xe3, 0x9c, 0x1c, 0x66, 0x5d, 0x69, 0x54, 0x68, 0x2c, 0x3a,
	0x6c, 0x34, 0x6b, 0x47, 0x58, 0x3b, 0xc2, 0x3f, 0x8e, 0x56, 0x90, 0x28, 0x37, 0xc8, 0xa3, 0x30,
	0xc6, 0x4c, 0x24, 0x98, 0xa0, 0xa8, 0x0c, 0x51, 0xde, 0xaf, 0x4e, 0xd5, 0xa1, 0x9a, 0xea, 0x8b,
	0x5a, 0x1b, 0x09, 0x62, 0x92, 0x82, 0x90, 0x46, 0x09, 0xa9, 0x35, 0x3a, 0xe9, 0x14, 0x6a, 0x9a,
	0xd3, 0xab, 0xcb, 0xc1, 0xd0, 0x2c, 0xc2, 0x2b, 0xcb, 0xd0, 0x8d, 0x0d, 0x9c, 0xa0, 0x8d, 0x65,
	0x54, 0xc8, 0x54, 0xf5, 0xa4, 0x83, 0x39, 0x6d, 0xaf, 0x50, 0x20, 0xd0, 0xc5, 0xca, 0xe5, 0x9b,
	0x2b, 0x1a, 0x05, 0x23, 0x5a, 0xfc, 0x44, 0x7b, 0xcb, 0xbf, 0xd8, 0x01, 0x69, 0x4c, 0x3a, 0xde,
	0x41, 0xdd, 0x57, 0x49, 0x07, 0x5e, 0xe4, 0x40, 0xae, 0xd1, 0xf0, 0xcf, 0x6a, 0x99, 0x41, 0x93,
	0x5d, 0x63, 0xd7, 0xcf, 0x77, 0xaa, 0xb9, 0x7d, 0xd9, 0xbf, 0xb4, 0xa2, 0x25, 0x83, 0x9a, 0x60,
	0xfb, 0xe3, 0x9a, 0xbf, 0xbe, 0x53, 0xed, 0xf6, 0xa1, 0x74, 0xf2, 0xfe, 0xd3, 0x47, 0x8d, 0x5f,
	0xcc, 0x5f, 0x5f, 0xd2, 0x36, 0xc2, 0xf0, 0xb4, 0x5d, 0x87, 0xff, 0x0a, 0xd0, 0x12, 0xff, 0xad,
	0xaf, 0x43, 0xb4, 0x5f, 0x95, 0x9f, 0x9a, 0xdb, 0xfd, 0xdb, 0xbb, 0x71, 0x40, 0x10, 0xe7, 0x16,
	0x32, 0xa0, 0x41, 0x40, 0x4e, 0xea, 0x9e, 0xb4, 0xbd, 0x20, 0x27, 0xb0, 0x37, 0x0a, 0xa0, 0x40,
	0x61, 0xa0, 0xc1, 0x8d, 0xd0, 0x0e, 0x83, 0x91, 0x55, 0x0e, 0xf6, 0xbe, 0xfd, 0x78, 0xbf, 0x76,
	0xaf, 0x7d, 0x77, 0xde, 0x11, 0x71, 0xfc, 0xa7, 0x64, 0x64, 0x0c, 0x24, 0x68, 0x4c, 0x0e, 0xb2,
	0x85, 0x1e, 0x91, 0x78, 0x79, 0x8c, 0x5f, 0x0b, 0x5b, 0x27, 0xe8, 0xc6, 0x55, 0x84, 0x3b, 0x6c,
	0xab, 0x15, 0x1c, 0xec, 0xb3, 0x33, 0x5f, 0xf7, 0xd9, 0xe6, 0xa9, 0xa9, 0x9f, 0x44, 0xcf, 0x21,
	0x76, 0x7b, 0x5f, 0x9a, 0x6b, 0x17, 0xd8, 0x83, 0xb7, 0x6c, 0x32, 0xe5, 0xde, 0xe1, 0x94, 0x7b,
	0x47, 0x53, 0xce, 0xde, 0x94, 0x9c, 0x7d, 0x28, 0x39, 0xfb, 0x5c, 0x72, 0x36, 0x29, 0x39, 0xfb,
	0x5e, 0x72, 0xf6, 0xb3, 0xe4, 0xde, 0x51, 0xc9, 0xd9, 0xbb, 0x19, 0xf7, 0x0e, 0x66, 0x9c, 0x4d,
	0x66, 0xdc, 0x3b, 0x9c, 0x71, 0xef, 0xd9, 0xe3, 0x04, 0xcd, 0x30, 0x09, 0x0b, 0x4c, 0x1d, 0x58,
	0x2b, 0xc3, 0x9c, 0x44, 0x35, 0xf4, 0xd1, 0x66, 0x81, 0xb1, 0x58, 0xa8, 0x1e, 0xd8, 0xe0, 0x04,
	0x0b, 0x13, 0x25, 0x28, 0x60, 0xd7, 0xcd, 0x1b, 0xf0, 0xd7, 0xdb, 0x88, 0xce, 0x55, 0x45, 0xb8,
	0xf5, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x0f, 0xbc, 0x28, 0x21, 0x3e, 0x03, 0x00, 0x00,
}

func (this *ReapplyConfigRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ReapplyConfigRequest)
	if !ok {
		that2, ok := that.(ReapplyConfigRequest)
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
	if this.Name != that1.Name {
		return false
	}
	return true
}
func (this *ReapplyConfigResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ReapplyConfigResponse)
	if !ok {
		that2, ok := that.(ReapplyConfigResponse)
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
func (this *ReapplyConfigRequest) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&cloud_link.ReapplyConfigRequest{")
	s = append(s, "Name: "+fmt.Sprintf("%#v", this.Name)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *ReapplyConfigResponse) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 4)
	s = append(s, "&cloud_link.ReapplyConfigResponse{")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringPublicCustomApi(v interface{}, typ string) string {
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

// CustomDataAPIClient is the client API for CustomDataAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CustomDataAPIClient interface {
	// CloudLink Reapply Config
	//
	// x-displayName: "CloudLink "
	// Reapply CloudLink Config
	ReapplyConfig(ctx context.Context, in *ReapplyConfigRequest, opts ...grpc.CallOption) (*ReapplyConfigResponse, error)
}

type customDataAPIClient struct {
	cc *grpc.ClientConn
}

func NewCustomDataAPIClient(cc *grpc.ClientConn) CustomDataAPIClient {
	return &customDataAPIClient{cc}
}

func (c *customDataAPIClient) ReapplyConfig(ctx context.Context, in *ReapplyConfigRequest, opts ...grpc.CallOption) (*ReapplyConfigResponse, error) {
	out := new(ReapplyConfigResponse)
	err := c.cc.Invoke(ctx, "/ves.io.schema.cloud_link.CustomDataAPI/ReapplyConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CustomDataAPIServer is the server API for CustomDataAPI service.
type CustomDataAPIServer interface {
	// CloudLink Reapply Config
	//
	// x-displayName: "CloudLink "
	// Reapply CloudLink Config
	ReapplyConfig(context.Context, *ReapplyConfigRequest) (*ReapplyConfigResponse, error)
}

// UnimplementedCustomDataAPIServer can be embedded to have forward compatible implementations.
type UnimplementedCustomDataAPIServer struct {
}

func (*UnimplementedCustomDataAPIServer) ReapplyConfig(ctx context.Context, req *ReapplyConfigRequest) (*ReapplyConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReapplyConfig not implemented")
}

func RegisterCustomDataAPIServer(s *grpc.Server, srv CustomDataAPIServer) {
	s.RegisterService(&_CustomDataAPI_serviceDesc, srv)
}

func _CustomDataAPI_ReapplyConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReapplyConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomDataAPIServer).ReapplyConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ves.io.schema.cloud_link.CustomDataAPI/ReapplyConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomDataAPIServer).ReapplyConfig(ctx, req.(*ReapplyConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CustomDataAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ves.io.schema.cloud_link.CustomDataAPI",
	HandlerType: (*CustomDataAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ReapplyConfig",
			Handler:    _CustomDataAPI_ReapplyConfig_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ves.io/schema/cloud_link/public_custom_api.proto",
}

func (m *ReapplyConfigRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ReapplyConfigRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ReapplyConfigRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintPublicCustomApi(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ReapplyConfigResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ReapplyConfigResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ReapplyConfigResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintPublicCustomApi(dAtA []byte, offset int, v uint64) int {
	offset -= sovPublicCustomApi(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ReapplyConfigRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovPublicCustomApi(uint64(l))
	}
	return n
}

func (m *ReapplyConfigResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovPublicCustomApi(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPublicCustomApi(x uint64) (n int) {
	return sovPublicCustomApi(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *ReapplyConfigRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ReapplyConfigRequest{`,
		`Name:` + fmt.Sprintf("%v", this.Name) + `,`,
		`}`,
	}, "")
	return s
}
func (this *ReapplyConfigResponse) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ReapplyConfigResponse{`,
		`}`,
	}, "")
	return s
}
func valueToStringPublicCustomApi(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *ReapplyConfigRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPublicCustomApi
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
			return fmt.Errorf("proto: ReapplyConfigRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ReapplyConfigRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPublicCustomApi
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
				return ErrInvalidLengthPublicCustomApi
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPublicCustomApi
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPublicCustomApi(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPublicCustomApi
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthPublicCustomApi
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
func (m *ReapplyConfigResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPublicCustomApi
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
			return fmt.Errorf("proto: ReapplyConfigResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ReapplyConfigResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipPublicCustomApi(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPublicCustomApi
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthPublicCustomApi
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
func skipPublicCustomApi(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPublicCustomApi
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
					return 0, ErrIntOverflowPublicCustomApi
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
					return 0, ErrIntOverflowPublicCustomApi
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
				return 0, ErrInvalidLengthPublicCustomApi
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPublicCustomApi
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPublicCustomApi
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPublicCustomApi        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPublicCustomApi          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPublicCustomApi = fmt.Errorf("proto: unexpected end of group")
)
