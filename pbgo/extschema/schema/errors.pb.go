// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ves.io/schema/errors.proto

package schema

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
	golang_proto "github.com/golang/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
	reflect "reflect"
	strconv "strconv"
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

// ErrorCode
//
// x-displayName: "Error Code"
// Union of all possible error-codes from system
type ErrorCode int32

const (
	// No error
	EOK ErrorCode = 0
	// Permissions error
	EPERMS ErrorCode = 1
	// Input is not correct
	EBADINPUT ErrorCode = 2
	// Not found
	ENOTFOUND ErrorCode = 3
	// Already exists
	EEXISTS ErrorCode = 4
	// Unknown/catchall error
	EUNKNOWN ErrorCode = 5
	// Error in serializing/de-serializing
	ESERIALIZE ErrorCode = 6
	// Server error
	EINTERNAL ErrorCode = 7
	// Partial error
	EPARTIAL ErrorCode = 8
)

var ErrorCode_name = map[int32]string{
	0: "EOK",
	1: "EPERMS",
	2: "EBADINPUT",
	3: "ENOTFOUND",
	4: "EEXISTS",
	5: "EUNKNOWN",
	6: "ESERIALIZE",
	7: "EINTERNAL",
	8: "EPARTIAL",
}

var ErrorCode_value = map[string]int32{
	"EOK":        0,
	"EPERMS":     1,
	"EBADINPUT":  2,
	"ENOTFOUND":  3,
	"EEXISTS":    4,
	"EUNKNOWN":   5,
	"ESERIALIZE": 6,
	"EINTERNAL":  7,
	"EPARTIAL":   8,
}

func (ErrorCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f8a113b924da889a, []int{0}
}

// ErrorType
//
// x-displayName: "Error Type"
// Information about a error in API operation
type ErrorType struct {
	// code
	//
	// x-displayName: "Code"
	// A simple general code by category
	Code ErrorCode `protobuf:"varint,1,opt,name=code,proto3,enum=ves.io.schema.ErrorCode" json:"code,omitempty"`
	// message
	//
	// x-displayName: "Message"
	// x-example: "value"
	// A human readable string of the error
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	// error_obj
	//
	// x-displayName: "Error Object"
	// A structured error object for machine parsing
	ErrorObj *types.Any `protobuf:"bytes,3,opt,name=error_obj,json=errorObj,proto3" json:"error_obj,omitempty"`
}

func (m *ErrorType) Reset()      { *m = ErrorType{} }
func (*ErrorType) ProtoMessage() {}
func (*ErrorType) Descriptor() ([]byte, []int) {
	return fileDescriptor_f8a113b924da889a, []int{0}
}
func (m *ErrorType) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ErrorType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ErrorType.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ErrorType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ErrorType.Merge(m, src)
}
func (m *ErrorType) XXX_Size() int {
	return m.Size()
}
func (m *ErrorType) XXX_DiscardUnknown() {
	xxx_messageInfo_ErrorType.DiscardUnknown(m)
}

var xxx_messageInfo_ErrorType proto.InternalMessageInfo

func (m *ErrorType) GetCode() ErrorCode {
	if m != nil {
		return m.Code
	}
	return EOK
}

func (m *ErrorType) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *ErrorType) GetErrorObj() *types.Any {
	if m != nil {
		return m.ErrorObj
	}
	return nil
}

func init() {
	proto.RegisterEnum("ves.io.schema.ErrorCode", ErrorCode_name, ErrorCode_value)
	golang_proto.RegisterEnum("ves.io.schema.ErrorCode", ErrorCode_name, ErrorCode_value)
	proto.RegisterType((*ErrorType)(nil), "ves.io.schema.ErrorType")
	golang_proto.RegisterType((*ErrorType)(nil), "ves.io.schema.ErrorType")
}

func init() { proto.RegisterFile("ves.io/schema/errors.proto", fileDescriptor_f8a113b924da889a) }
func init() { golang_proto.RegisterFile("ves.io/schema/errors.proto", fileDescriptor_f8a113b924da889a) }

var fileDescriptor_f8a113b924da889a = []byte{
	// 421 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x91, 0x3f, 0x6f, 0xd3, 0x40,
	0x18, 0xc6, 0x7d, 0x4d, 0xc9, 0x9f, 0x2b, 0xad, 0x4e, 0x27, 0x06, 0x93, 0xe1, 0x14, 0x31, 0x45,
	0x88, 0x9c, 0x45, 0xf9, 0x04, 0x6e, 0x7b, 0x48, 0x56, 0xc3, 0x39, 0x72, 0x1c, 0x81, 0xba, 0xa0,
	0x38, 0xb9, 0x5e, 0x53, 0xea, 0xbe, 0xd6, 0xd9, 0x89, 0xc8, 0x80, 0xc4, 0xc4, 0xcc, 0xc7, 0xe0,
	0x63, 0x30, 0x32, 0x66, 0xec, 0x48, 0xec, 0x85, 0xb1, 0x1f, 0x01, 0xf5, 0x4c, 0x40, 0x4c, 0xf7,
	0x3e, 0x7a, 0x7e, 0xaf, 0x9e, 0x47, 0xf7, 0xe2, 0xee, 0x4a, 0xe5, 0x7c, 0x01, 0x5e, 0x3e, 0xbb,
	0x52, 0xe9, 0xd4, 0x53, 0xc6, 0x80, 0xc9, 0x79, 0x66, 0xa0, 0x00, 0x7a, 0x58, 0x7b, 0xbc, 0xf6,
	0xba, 0x03, 0xbd, 0x28, 0xae, 0x96, 0x09, 0x9f, 0x41, 0xea, 0x69, 0xd0, 0xe0, 0x59, 0x2a, 0x59,
	0x5e, 0x5a, 0x65, 0x85, 0x9d, 0xea, 0xed, 0xee, 0x53, 0x0d, 0xa0, 0x6f, 0xd4, 0x3f, 0x6a, 0x7a,
	0xbb, 0xae, 0xad, 0x67, 0x5f, 0x10, 0xee, 0x88, 0x87, 0xa4, 0x78, 0x9d, 0x29, 0xfa, 0x02, 0xef,
	0xcf, 0x60, 0xae, 0x5c, 0xd4, 0x43, 0xfd, 0xa3, 0x63, 0x97, 0xff, 0x97, 0xca, 0x2d, 0x77, 0x0a,
	0x73, 0x15, 0x59, 0x8a, 0xba, 0xb8, 0x95, 0xaa, 0x3c, 0x9f, 0x6a, 0xe5, 0xee, 0xf5, 0x50, 0xbf,
	0x13, 0xed, 0x24, 0x7d, 0x89, 0x3b, 0xb6, 0xfe, 0x7b, 0x48, 0xae, 0xdd, 0x46, 0x0f, 0xf5, 0x0f,
	0x8e, 0x9f, 0xf0, 0xba, 0x04, 0xdf, 0x95, 0xe0, 0xfe, 0xed, 0x3a, 0x6a, 0x5b, 0x2c, 0x4c, 0xae,
	0x9f, 0xff, 0x2d, 0xf2, 0x10, 0x40, 0x5b, 0xb8, 0x21, 0xc2, 0x73, 0xe2, 0x50, 0x8c, 0x9b, 0x62,
	0x24, 0xa2, 0x37, 0x63, 0x82, 0xe8, 0x21, 0xee, 0x88, 0x13, 0xff, 0x2c, 0x90, 0xa3, 0x49, 0x4c,
	0xf6, 0xac, 0x94, 0x61, 0xfc, 0x3a, 0x9c, 0xc8, 0x33, 0xd2, 0xa0, 0x07, 0xb8, 0x25, 0xc4, 0xbb,
	0x60, 0x1c, 0x8f, 0xc9, 0x3e, 0x7d, 0x8c, 0xdb, 0x62, 0x22, 0xcf, 0x65, 0xf8, 0x56, 0x92, 0x47,
	0xf4, 0x08, 0x63, 0x31, 0x16, 0x51, 0xe0, 0x0f, 0x83, 0x0b, 0x41, 0x9a, 0x76, 0x33, 0x90, 0xb1,
	0x88, 0xa4, 0x3f, 0x24, 0x2d, 0x0b, 0x8f, 0xfc, 0x28, 0x0e, 0xfc, 0x21, 0x69, 0x9f, 0x7c, 0xda,
	0x6c, 0x99, 0x73, 0xb7, 0x65, 0xce, 0xfd, 0x96, 0xa1, 0xcf, 0x25, 0x43, 0xdf, 0x4a, 0x86, 0x7e,
	0x94, 0x0c, 0x6d, 0x4a, 0x86, 0x7e, 0x96, 0x0c, 0xfd, 0x2a, 0x99, 0x73, 0x5f, 0x32, 0xf4, 0xb5,
	0x62, 0xce, 0xf7, 0x8a, 0xa1, 0x4d, 0xc5, 0x9c, 0xbb, 0x8a, 0x39, 0x17, 0xa7, 0x1a, 0xb2, 0x0f,
	0x9a, 0xaf, 0xe0, 0xa6, 0x50, 0xc6, 0x4c, 0xf9, 0x32, 0xf7, 0xec, 0x70, 0x09, 0x26, 0x1d, 0x64,
	0x06, 0x56, 0x8b, 0xb9, 0x32, 0x83, 0x9d, 0xed, 0x65, 0x89, 0x06, 0x4f, 0x7d, 0x2c, 0xfe, 0xdc,
	0xbb, 0x7e, 0x92, 0xa6, 0xfd, 0x9f, 0x57, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0xd9, 0x5b, 0xeb,
	0x7a, 0x0e, 0x02, 0x00, 0x00,
}

func (x ErrorCode) String() string {
	s, ok := ErrorCode_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (this *ErrorType) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ErrorType)
	if !ok {
		that2, ok := that.(ErrorType)
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
	if this.Code != that1.Code {
		return false
	}
	if this.Message != that1.Message {
		return false
	}
	if !this.ErrorObj.Equal(that1.ErrorObj) {
		return false
	}
	return true
}
func (this *ErrorType) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 7)
	s = append(s, "&schema.ErrorType{")
	s = append(s, "Code: "+fmt.Sprintf("%#v", this.Code)+",\n")
	s = append(s, "Message: "+fmt.Sprintf("%#v", this.Message)+",\n")
	if this.ErrorObj != nil {
		s = append(s, "ErrorObj: "+fmt.Sprintf("%#v", this.ErrorObj)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringErrors(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *ErrorType) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ErrorType) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ErrorType) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ErrorObj != nil {
		{
			size, err := m.ErrorObj.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintErrors(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Message) > 0 {
		i -= len(m.Message)
		copy(dAtA[i:], m.Message)
		i = encodeVarintErrors(dAtA, i, uint64(len(m.Message)))
		i--
		dAtA[i] = 0x12
	}
	if m.Code != 0 {
		i = encodeVarintErrors(dAtA, i, uint64(m.Code))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintErrors(dAtA []byte, offset int, v uint64) int {
	offset -= sovErrors(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ErrorType) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Code != 0 {
		n += 1 + sovErrors(uint64(m.Code))
	}
	l = len(m.Message)
	if l > 0 {
		n += 1 + l + sovErrors(uint64(l))
	}
	if m.ErrorObj != nil {
		l = m.ErrorObj.Size()
		n += 1 + l + sovErrors(uint64(l))
	}
	return n
}

func sovErrors(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozErrors(x uint64) (n int) {
	return sovErrors(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *ErrorType) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ErrorType{`,
		`Code:` + fmt.Sprintf("%v", this.Code) + `,`,
		`Message:` + fmt.Sprintf("%v", this.Message) + `,`,
		`ErrorObj:` + strings.Replace(fmt.Sprintf("%v", this.ErrorObj), "Any", "types.Any", 1) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringErrors(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *ErrorType) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowErrors
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
			return fmt.Errorf("proto: ErrorType: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ErrorType: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Code", wireType)
			}
			m.Code = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowErrors
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Code |= ErrorCode(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Message", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowErrors
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
				return ErrInvalidLengthErrors
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthErrors
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Message = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ErrorObj", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowErrors
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
				return ErrInvalidLengthErrors
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthErrors
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ErrorObj == nil {
				m.ErrorObj = &types.Any{}
			}
			if err := m.ErrorObj.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipErrors(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthErrors
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthErrors
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
func skipErrors(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowErrors
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
					return 0, ErrIntOverflowErrors
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
					return 0, ErrIntOverflowErrors
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
				return 0, ErrInvalidLengthErrors
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupErrors
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthErrors
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthErrors        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowErrors          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupErrors = fmt.Errorf("proto: unexpected end of group")
)
