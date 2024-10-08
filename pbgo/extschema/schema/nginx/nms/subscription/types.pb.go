// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ves.io/schema/nginx/nms/subscription/types.proto

package subscription

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema"
	io "io"
	math "math"
	math_bits "math/bits"
	reflect "reflect"
	strings "strings"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// SubscribeRequest
//
// x-displayName: "Subscribe Request"
// Request to subscribe to NGINX Management Suite
type SubscribeRequest struct {
}

func (m *SubscribeRequest) Reset()      { *m = SubscribeRequest{} }
func (*SubscribeRequest) ProtoMessage() {}
func (*SubscribeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d8b2d35681b544d1, []int{0}
}
func (m *SubscribeRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SubscribeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SubscribeRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SubscribeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubscribeRequest.Merge(m, src)
}
func (m *SubscribeRequest) XXX_Size() int {
	return m.Size()
}
func (m *SubscribeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SubscribeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SubscribeRequest proto.InternalMessageInfo

// SubscribeResponse
//
// x-displayName: "Subscribe Response"
// Response of subscribe to NGINX Management Suite
type SubscribeResponse struct {
}

func (m *SubscribeResponse) Reset()      { *m = SubscribeResponse{} }
func (*SubscribeResponse) ProtoMessage() {}
func (*SubscribeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d8b2d35681b544d1, []int{1}
}
func (m *SubscribeResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SubscribeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SubscribeResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SubscribeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubscribeResponse.Merge(m, src)
}
func (m *SubscribeResponse) XXX_Size() int {
	return m.Size()
}
func (m *SubscribeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SubscribeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SubscribeResponse proto.InternalMessageInfo

// UnsubscribeRequest
//
// x-displayName: "Unsubscribe Request"
// Request to unsubscribe to NGINX Management Suite
type UnsubscribeRequest struct {
}

func (m *UnsubscribeRequest) Reset()      { *m = UnsubscribeRequest{} }
func (*UnsubscribeRequest) ProtoMessage() {}
func (*UnsubscribeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d8b2d35681b544d1, []int{2}
}
func (m *UnsubscribeRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UnsubscribeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UnsubscribeRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *UnsubscribeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UnsubscribeRequest.Merge(m, src)
}
func (m *UnsubscribeRequest) XXX_Size() int {
	return m.Size()
}
func (m *UnsubscribeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UnsubscribeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UnsubscribeRequest proto.InternalMessageInfo

// UnsubscribeResponse
//
// x-displayName: "Unsubscribe Response"
// Response of unsubscribe to NGINX Management Suite
type UnsubscribeResponse struct {
}

func (m *UnsubscribeResponse) Reset()      { *m = UnsubscribeResponse{} }
func (*UnsubscribeResponse) ProtoMessage() {}
func (*UnsubscribeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d8b2d35681b544d1, []int{3}
}
func (m *UnsubscribeResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UnsubscribeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UnsubscribeResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *UnsubscribeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UnsubscribeResponse.Merge(m, src)
}
func (m *UnsubscribeResponse) XXX_Size() int {
	return m.Size()
}
func (m *UnsubscribeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UnsubscribeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UnsubscribeResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*SubscribeRequest)(nil), "ves.io.schema.nginx.nms.subscription.SubscribeRequest")
	proto.RegisterType((*SubscribeResponse)(nil), "ves.io.schema.nginx.nms.subscription.SubscribeResponse")
	proto.RegisterType((*UnsubscribeRequest)(nil), "ves.io.schema.nginx.nms.subscription.UnsubscribeRequest")
	proto.RegisterType((*UnsubscribeResponse)(nil), "ves.io.schema.nginx.nms.subscription.UnsubscribeResponse")
}

func init() {
	proto.RegisterFile("ves.io/schema/nginx/nms/subscription/types.proto", fileDescriptor_d8b2d35681b544d1)
}

var fileDescriptor_d8b2d35681b544d1 = []byte{
	// 269 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0x31, 0x4e, 0xc3, 0x30,
	0x14, 0x86, 0xed, 0x85, 0x21, 0x13, 0xa4, 0x30, 0x10, 0xd0, 0x1b, 0x22, 0xd6, 0xda, 0x48, 0xdc,
	0x80, 0x23, 0x80, 0x58, 0xba, 0x25, 0xad, 0x09, 0x16, 0x8d, 0x9f, 0xf1, 0x73, 0xa2, 0xb2, 0x71,
	0x84, 0x1e, 0x83, 0xa3, 0x30, 0x66, 0xec, 0x48, 0x9c, 0x85, 0xb1, 0x47, 0x40, 0x4a, 0x8a, 0x94,
	0x30, 0x74, 0xb3, 0xfe, 0xef, 0x7d, 0xbf, 0xf5, 0x5e, 0x74, 0x5b, 0x2b, 0x12, 0x1a, 0x25, 0x2d,
	0x5f, 0x54, 0x99, 0x49, 0x53, 0x68, 0xb3, 0x91, 0xa6, 0x24, 0x49, 0x55, 0x4e, 0x4b, 0xa7, 0xad,
	0xd7, 0x68, 0xa4, 0x7f, 0xb7, 0x8a, 0x84, 0x75, 0xe8, 0x31, 0xbe, 0x19, 0x0c, 0x31, 0x18, 0xa2,
	0x37, 0x84, 0x29, 0x49, 0x8c, 0x8d, 0xe4, 0x6a, 0xda, 0x8b, 0x7d, 0x7a, 0xa8, 0x48, 0x2e, 0xa7,
	0x70, 0xd4, 0x9e, 0x5c, 0x4f, 0x51, 0x9d, 0xad, 0xf5, 0x2a, 0xf3, 0x6a, 0xa0, 0x69, 0x1c, 0x9d,
	0x3e, 0x0e, 0xbf, 0xe4, 0xea, 0x41, 0xbd, 0x55, 0x8a, 0x7c, 0x3a, 0x8b, 0xce, 0x46, 0x19, 0x59,
	0x34, 0xa4, 0xd2, 0xf3, 0x28, 0x7e, 0x32, 0xf4, 0x7f, 0xf4, 0x22, 0x9a, 0x4d, 0xd2, 0x61, 0xf8,
	0x7e, 0xcb, 0x9b, 0x16, 0xd8, 0xae, 0x05, 0xb6, 0x6f, 0x81, 0x7f, 0x04, 0xe0, 0x9f, 0x01, 0xf8,
	0x57, 0x00, 0xde, 0x04, 0xe0, 0xdf, 0x01, 0xf8, 0x4f, 0x00, 0xb6, 0x0f, 0xc0, 0xb7, 0x1d, 0xb0,
	0xa6, 0x03, 0xb6, 0xeb, 0x80, 0x2d, 0x16, 0x05, 0xda, 0xd7, 0x42, 0xd4, 0xb8, 0xf6, 0xca, 0xb9,
	0x4c, 0x54, 0x24, 0xfb, 0xc7, 0x33, 0xba, 0x72, 0x6e, 0x1d, 0xd6, 0x7a, 0xa5, 0xdc, 0xfc, 0x0f,
	0x4b, 0x9b, 0x17, 0x28, 0xd5, 0xc6, 0x1f, 0x36, 0x3b, 0x7a, 0xf0, 0xfc, 0xa4, 0xdf, 0xf7, 0xee,
	0x37, 0x00, 0x00, 0xff, 0xff, 0x06, 0x03, 0x75, 0x9f, 0x9f, 0x01, 0x00, 0x00,
}

func (this *SubscribeRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*SubscribeRequest)
	if !ok {
		that2, ok := that.(SubscribeRequest)
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
func (this *SubscribeResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*SubscribeResponse)
	if !ok {
		that2, ok := that.(SubscribeResponse)
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
func (this *UnsubscribeRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*UnsubscribeRequest)
	if !ok {
		that2, ok := that.(UnsubscribeRequest)
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
func (this *UnsubscribeResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*UnsubscribeResponse)
	if !ok {
		that2, ok := that.(UnsubscribeResponse)
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
func (this *SubscribeRequest) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 4)
	s = append(s, "&subscription.SubscribeRequest{")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *SubscribeResponse) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 4)
	s = append(s, "&subscription.SubscribeResponse{")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *UnsubscribeRequest) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 4)
	s = append(s, "&subscription.UnsubscribeRequest{")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *UnsubscribeResponse) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 4)
	s = append(s, "&subscription.UnsubscribeResponse{")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringTypes(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *SubscribeRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SubscribeRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SubscribeRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *SubscribeResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SubscribeResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SubscribeResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *UnsubscribeRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UnsubscribeRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *UnsubscribeRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *UnsubscribeResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UnsubscribeResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *UnsubscribeResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTypes(dAtA []byte, offset int, v uint64) int {
	offset -= sovTypes(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *SubscribeRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *SubscribeResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *UnsubscribeRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *UnsubscribeResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTypes(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTypes(x uint64) (n int) {
	return sovTypes(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *SubscribeRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&SubscribeRequest{`,
		`}`,
	}, "")
	return s
}
func (this *SubscribeResponse) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&SubscribeResponse{`,
		`}`,
	}, "")
	return s
}
func (this *UnsubscribeRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&UnsubscribeRequest{`,
		`}`,
	}, "")
	return s
}
func (this *UnsubscribeResponse) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&UnsubscribeResponse{`,
		`}`,
	}, "")
	return s
}
func valueToStringTypes(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *SubscribeRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: SubscribeRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SubscribeRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTypes
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
func (m *SubscribeResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: SubscribeResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SubscribeResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTypes
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
func (m *UnsubscribeRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: UnsubscribeRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UnsubscribeRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTypes
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
func (m *UnsubscribeResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: UnsubscribeResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UnsubscribeResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTypes
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
func skipTypes(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTypes
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
					return 0, ErrIntOverflowTypes
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
					return 0, ErrIntOverflowTypes
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
				return 0, ErrInvalidLengthTypes
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTypes
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTypes
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTypes        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTypes          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTypes = fmt.Errorf("proto: unexpected end of group")
)
