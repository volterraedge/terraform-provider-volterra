// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ves.io/schema/site/public_customapi.proto

package site

import proto "github.com/gogo/protobuf/proto"
import golang_proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import _ "github.com/gogo/googleapis/google/api"
import _ "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema"
import _ "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema"
import _ "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/vesenv"

import strings "strings"
import reflect "reflect"

import context "golang.org/x/net/context"
import grpc "google.golang.org/grpc"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = golang_proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Set state request
//
// x-displayName: "Set Status Request"
// Set status of the site
type SetStateReq struct {
	// Namespace
	//
	// x-displayName: "Namespace"
	// x-required
	// x-example: "system"
	// Site namespace
	Namespace string `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	// Name
	//
	// x-displayName: "Name"
	// x-required
	// x-example: "ce398"
	// Site name
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// State
	//
	// x-displayName: "State"
	// x-required
	// x-example: 3
	// Desired (target) state for site (3 = STANDBY)
	State SiteState `protobuf:"varint,3,opt,name=state,proto3,enum=ves.io.schema.site.SiteState" json:"state,omitempty"`
}

func (m *SetStateReq) Reset()                    { *m = SetStateReq{} }
func (*SetStateReq) ProtoMessage()               {}
func (*SetStateReq) Descriptor() ([]byte, []int) { return fileDescriptorPublicCustomapi, []int{0} }

func (m *SetStateReq) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *SetStateReq) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SetStateReq) GetState() SiteState {
	if m != nil {
		return m.State
	}
	return ONLINE
}

// Set state responde
//
// x-displayName: "Set Status Response"
// Response for set state request, empty because the only resturned information
// is currently error message
type SetStateResp struct {
}

func (m *SetStateResp) Reset()                    { *m = SetStateResp{} }
func (*SetStateResp) ProtoMessage()               {}
func (*SetStateResp) Descriptor() ([]byte, []int) { return fileDescriptorPublicCustomapi, []int{1} }

func init() {
	proto.RegisterType((*SetStateReq)(nil), "ves.io.schema.site.SetStateReq")
	golang_proto.RegisterType((*SetStateReq)(nil), "ves.io.schema.site.SetStateReq")
	proto.RegisterType((*SetStateResp)(nil), "ves.io.schema.site.SetStateResp")
	golang_proto.RegisterType((*SetStateResp)(nil), "ves.io.schema.site.SetStateResp")
}
func (this *SetStateReq) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*SetStateReq)
	if !ok {
		that2, ok := that.(SetStateReq)
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
func (this *SetStateResp) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*SetStateResp)
	if !ok {
		that2, ok := that.(SetStateResp)
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
func (this *SetStateReq) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 7)
	s = append(s, "&site.SetStateReq{")
	s = append(s, "Namespace: "+fmt.Sprintf("%#v", this.Namespace)+",\n")
	s = append(s, "Name: "+fmt.Sprintf("%#v", this.Name)+",\n")
	s = append(s, "State: "+fmt.Sprintf("%#v", this.State)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *SetStateResp) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 4)
	s = append(s, "&site.SetStateResp{")
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

// Client API for CustomStateAPI service

type CustomStateAPIClient interface {
	// Set site state
	//
	// x-displayName: "Set site state"
	// Request changing site state but this request goes through validation as some
	// trainsitions are not allowed.
	// It can be used to decomission site by sending state DECOMISSIONING. Example of
	// forbidden state is PROVISIONING and UPGRADING.
	SetState(ctx context.Context, in *SetStateReq, opts ...grpc.CallOption) (*SetStateResp, error)
}

type customStateAPIClient struct {
	cc *grpc.ClientConn
}

func NewCustomStateAPIClient(cc *grpc.ClientConn) CustomStateAPIClient {
	return &customStateAPIClient{cc}
}

func (c *customStateAPIClient) SetState(ctx context.Context, in *SetStateReq, opts ...grpc.CallOption) (*SetStateResp, error) {
	out := new(SetStateResp)
	err := grpc.Invoke(ctx, "/ves.io.schema.site.CustomStateAPI/SetState", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for CustomStateAPI service

type CustomStateAPIServer interface {
	// Set site state
	//
	// x-displayName: "Set site state"
	// Request changing site state but this request goes through validation as some
	// trainsitions are not allowed.
	// It can be used to decomission site by sending state DECOMISSIONING. Example of
	// forbidden state is PROVISIONING and UPGRADING.
	SetState(context.Context, *SetStateReq) (*SetStateResp, error)
}

func RegisterCustomStateAPIServer(s *grpc.Server, srv CustomStateAPIServer) {
	s.RegisterService(&_CustomStateAPI_serviceDesc, srv)
}

func _CustomStateAPI_SetState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetStateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomStateAPIServer).SetState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ves.io.schema.site.CustomStateAPI/SetState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomStateAPIServer).SetState(ctx, req.(*SetStateReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _CustomStateAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ves.io.schema.site.CustomStateAPI",
	HandlerType: (*CustomStateAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetState",
			Handler:    _CustomStateAPI_SetState_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ves.io/schema/site/public_customapi.proto",
}

func (m *SetStateReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SetStateReq) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Namespace) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintPublicCustomapi(dAtA, i, uint64(len(m.Namespace)))
		i += copy(dAtA[i:], m.Namespace)
	}
	if len(m.Name) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintPublicCustomapi(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	if m.State != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintPublicCustomapi(dAtA, i, uint64(m.State))
	}
	return i, nil
}

func (m *SetStateResp) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SetStateResp) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func encodeVarintPublicCustomapi(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *SetStateReq) Size() (n int) {
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

func (m *SetStateResp) Size() (n int) {
	var l int
	_ = l
	return n
}

func sovPublicCustomapi(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozPublicCustomapi(x uint64) (n int) {
	return sovPublicCustomapi(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *SetStateReq) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&SetStateReq{`,
		`Namespace:` + fmt.Sprintf("%v", this.Namespace) + `,`,
		`Name:` + fmt.Sprintf("%v", this.Name) + `,`,
		`State:` + fmt.Sprintf("%v", this.State) + `,`,
		`}`,
	}, "")
	return s
}
func (this *SetStateResp) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&SetStateResp{`,
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
func (m *SetStateReq) Unmarshal(dAtA []byte) error {
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
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: SetStateReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SetStateReq: illegal tag %d (wire type %d)", fieldNum, wire)
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
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPublicCustomapi
			}
			postIndex := iNdEx + intStringLen
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
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPublicCustomapi
			}
			postIndex := iNdEx + intStringLen
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
				m.State |= (SiteState(b) & 0x7F) << shift
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
func (m *SetStateResp) Unmarshal(dAtA []byte) error {
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
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: SetStateResp: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SetStateResp: illegal tag %d (wire type %d)", fieldNum, wire)
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
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthPublicCustomapi
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowPublicCustomapi
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipPublicCustomapi(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthPublicCustomapi = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPublicCustomapi   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("ves.io/schema/site/public_customapi.proto", fileDescriptorPublicCustomapi)
}
func init() {
	golang_proto.RegisterFile("ves.io/schema/site/public_customapi.proto", fileDescriptorPublicCustomapi)
}

var fileDescriptorPublicCustomapi = []byte{
	// 460 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xbf, 0x6e, 0xd4, 0x4e,
	0x10, 0xc7, 0x6f, 0x2e, 0xbf, 0x1f, 0x22, 0x06, 0x5d, 0x61, 0x51, 0x38, 0x47, 0x58, 0x4e, 0x6e,
	0x08, 0x48, 0xde, 0x45, 0x89, 0x68, 0xa8, 0xf8, 0xd3, 0x40, 0x05, 0xca, 0x75, 0x34, 0x68, 0x6d,
	0x26, 0x9b, 0x85, 0xb3, 0x67, 0xf1, 0xae, 0xcd, 0x3f, 0x45, 0x42, 0x79, 0x02, 0x24, 0xde, 0x80,
	0x8a, 0x77, 0x48, 0x93, 0x0e, 0x2a, 0x14, 0x41, 0x43, 0x85, 0x38, 0x87, 0x82, 0x32, 0x8f, 0x80,
	0x6e, 0x9d, 0xcb, 0xe5, 0xc4, 0x89, 0xee, 0x3b, 0xf3, 0xf9, 0xee, 0xcc, 0xec, 0xec, 0x06, 0x57,
	0x6b, 0xb4, 0x5c, 0x93, 0xb0, 0xd9, 0x36, 0xe6, 0x52, 0x58, 0xed, 0x50, 0x98, 0x2a, 0x1d, 0xe9,
	0xec, 0x71, 0x56, 0x59, 0x47, 0xb9, 0x34, 0x9a, 0x9b, 0x92, 0x1c, 0x85, 0x61, 0x6b, 0xe5, 0xad,
	0x95, 0x4f, 0xac, 0xfd, 0x44, 0x69, 0xb7, 0x5d, 0xa5, 0x3c, 0xa3, 0x5c, 0x28, 0x52, 0x24, 0xbc,
	0x35, 0xad, 0xb6, 0x7c, 0xe4, 0x03, 0xaf, 0xda, 0x12, 0xfd, 0x55, 0x45, 0xa4, 0x46, 0x28, 0xa4,
	0xd1, 0x42, 0x16, 0x05, 0x39, 0xe9, 0x34, 0x15, 0xf6, 0x98, 0x5e, 0x9c, 0x9f, 0x85, 0xcc, 0x69,
	0xc8, 0x16, 0x0c, 0xea, 0x5e, 0x19, 0x9c, 0xf2, 0x95, 0x79, 0x7e, 0x1a, 0xc5, 0xf3, 0xa8, 0x46,
	0x8b, 0x45, 0x3d, 0x5f, 0x3e, 0x76, 0xc1, 0xb9, 0x21, 0xba, 0xa1, 0x93, 0x0e, 0x37, 0xf1, 0x79,
	0xb8, 0x1a, 0x2c, 0x17, 0x32, 0x47, 0x6b, 0x64, 0x86, 0x11, 0x0c, 0x60, 0x6d, 0x79, 0x73, 0x96,
	0x08, 0xc3, 0xe0, 0xbf, 0x49, 0x10, 0x75, 0x3d, 0xf0, 0x3a, 0xdc, 0x08, 0xfe, 0xb7, 0x93, 0xd3,
	0xd1, 0xd2, 0x00, 0xd6, 0x7a, 0xeb, 0x97, 0xf8, 0xdf, 0xdb, 0xe2, 0x43, 0xed, 0xb0, 0x6d, 0xd1,
	0x7a, 0xe3, 0x5e, 0x70, 0x7e, 0xd6, 0xd5, 0x9a, 0xf5, 0x1f, 0x10, 0xf4, 0xee, 0xfa, 0xb5, 0xfb,
	0xdc, 0xed, 0x87, 0xf7, 0xc3, 0x0f, 0x10, 0x9c, 0x9d, 0x7a, 0xc2, 0xcb, 0x0b, 0xab, 0xce, 0xe6,
	0xee, 0x0f, 0xfe, 0x6d, 0xb0, 0x26, 0x1e, 0x36, 0x9f, 0xa2, 0x0b, 0x35, 0xda, 0x44, 0x53, 0xa2,
	0xb0, 0xc0, 0x52, 0x8e, 0x92, 0x17, 0xa5, 0x76, 0xb8, 0xfb, 0xed, 0xd7, 0xfb, 0xee, 0x8d, 0xf8,
	0xfa, 0xf1, 0xf3, 0x8b, 0x93, 0xeb, 0x5a, 0xf1, 0xe6, 0x44, 0xef, 0xb4, 0x8b, 0xf7, 0x89, 0x1d,
	0xe1, 0x2f, 0x71, 0x13, 0xae, 0xf5, 0xaf, 0xec, 0xef, 0xc1, 0xd2, 0xd7, 0x3d, 0x58, 0x59, 0xd0,
	0xfd, 0x41, 0xfa, 0x14, 0x33, 0xb7, 0xfb, 0x25, 0xea, 0xde, 0x82, 0x3b, 0xaf, 0x0f, 0xc6, 0xac,
	0xf3, 0x7d, 0xcc, 0x3a, 0x47, 0x63, 0x06, 0x6f, 0x1b, 0x06, 0x1f, 0x1b, 0x06, 0x9f, 0x1b, 0x06,
	0x07, 0x0d, 0x83, 0x9f, 0x0d, 0x83, 0xdf, 0x0d, 0xeb, 0x1c, 0x35, 0x0c, 0xde, 0x1d, 0xb2, 0xce,
	0xfe, 0x21, 0x83, 0x47, 0xf7, 0x14, 0x99, 0x67, 0x8a, 0xd7, 0x34, 0x72, 0x58, 0x96, 0x92, 0x57,
	0x56, 0x78, 0xb1, 0x45, 0x65, 0x9e, 0x98, 0x92, 0x6a, 0xfd, 0x04, 0xcb, 0x64, 0x8a, 0x85, 0x49,
	0x15, 0x09, 0x7c, 0xe9, 0xa6, 0xdf, 0x64, 0xf6, 0x5b, 0xd2, 0x33, 0xfe, 0xa5, 0x37, 0xfe, 0x04,
	0x00, 0x00, 0xff, 0xff, 0xb1, 0x3f, 0x84, 0x3e, 0xf3, 0x02, 0x00, 0x00,
}
