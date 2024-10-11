// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ves.io/schema/data_privacy/geo_config/types.proto

package geo_config

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	golang_proto "github.com/golang/protobuf/proto"
	schema "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema"
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

// Geo Config specification
//
// x-displayName: "Specification"
// Desired state for geo config
type GlobalSpecType struct {
	// Region Preference
	//
	// x-displayName: "Region Preference"
	// x-example: "us-west-1"
	// Preferred region to store LMA data
	PreferredRegion []*schema.ObjectRefType `protobuf:"bytes,1,rep,name=preferred_region,json=preferredRegion,proto3" json:"preferred_region,omitempty"`
}

func (m *GlobalSpecType) Reset()      { *m = GlobalSpecType{} }
func (*GlobalSpecType) ProtoMessage() {}
func (*GlobalSpecType) Descriptor() ([]byte, []int) {
	return fileDescriptor_255abf386cba8fac, []int{0}
}
func (m *GlobalSpecType) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GlobalSpecType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *GlobalSpecType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GlobalSpecType.Merge(m, src)
}
func (m *GlobalSpecType) XXX_Size() int {
	return m.Size()
}
func (m *GlobalSpecType) XXX_DiscardUnknown() {
	xxx_messageInfo_GlobalSpecType.DiscardUnknown(m)
}

var xxx_messageInfo_GlobalSpecType proto.InternalMessageInfo

func (m *GlobalSpecType) GetPreferredRegion() []*schema.ObjectRefType {
	if m != nil {
		return m.PreferredRegion
	}
	return nil
}

// Get Geo Config
//
// x-displayName: "Get Geo Config"
// Shape of the geo config specification
type GetSpecType struct {
	PreferredRegion []*schema.ObjectRefType `protobuf:"bytes,1,rep,name=preferred_region,json=preferredRegion,proto3" json:"preferred_region,omitempty"`
}

func (m *GetSpecType) Reset()      { *m = GetSpecType{} }
func (*GetSpecType) ProtoMessage() {}
func (*GetSpecType) Descriptor() ([]byte, []int) {
	return fileDescriptor_255abf386cba8fac, []int{1}
}
func (m *GetSpecType) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetSpecType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *GetSpecType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetSpecType.Merge(m, src)
}
func (m *GetSpecType) XXX_Size() int {
	return m.Size()
}
func (m *GetSpecType) XXX_DiscardUnknown() {
	xxx_messageInfo_GetSpecType.DiscardUnknown(m)
}

var xxx_messageInfo_GetSpecType proto.InternalMessageInfo

func (m *GetSpecType) GetPreferredRegion() []*schema.ObjectRefType {
	if m != nil {
		return m.PreferredRegion
	}
	return nil
}

func init() {
	proto.RegisterType((*GlobalSpecType)(nil), "ves.io.schema.data_privacy.geo_config.GlobalSpecType")
	golang_proto.RegisterType((*GlobalSpecType)(nil), "ves.io.schema.data_privacy.geo_config.GlobalSpecType")
	proto.RegisterType((*GetSpecType)(nil), "ves.io.schema.data_privacy.geo_config.GetSpecType")
	golang_proto.RegisterType((*GetSpecType)(nil), "ves.io.schema.data_privacy.geo_config.GetSpecType")
}

func init() {
	proto.RegisterFile("ves.io/schema/data_privacy/geo_config/types.proto", fileDescriptor_255abf386cba8fac)
}
func init() {
	golang_proto.RegisterFile("ves.io/schema/data_privacy/geo_config/types.proto", fileDescriptor_255abf386cba8fac)
}

var fileDescriptor_255abf386cba8fac = []byte{
	// 412 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0x31, 0xaf, 0xd3, 0x30,
	0x10, 0xc7, 0x63, 0x9e, 0x60, 0xc8, 0x93, 0xe0, 0xf1, 0x16, 0x20, 0x3c, 0xf9, 0x55, 0x15, 0x48,
	0x2c, 0xb5, 0x05, 0x6c, 0x0c, 0x0c, 0x5d, 0x3a, 0x22, 0x15, 0x26, 0x18, 0x22, 0x27, 0xb9, 0xb8,
	0x86, 0xa4, 0x67, 0x39, 0x6e, 0xa0, 0x03, 0x82, 0x85, 0x1d, 0xb1, 0xf1, 0x0d, 0x10, 0x1f, 0x81,
	0xa9, 0x23, 0x62, 0xea, 0xd8, 0x91, 0x3a, 0x0b, 0x6c, 0xfd, 0x08, 0x88, 0x34, 0xb4, 0x4d, 0x2b,
	0x18, 0xd8, 0xee, 0xf4, 0xbf, 0xdf, 0xfd, 0xcf, 0x77, 0xf6, 0xef, 0x96, 0x50, 0x30, 0x85, 0xbc,
	0x88, 0x47, 0x90, 0x0b, 0x9e, 0x08, 0x2b, 0x42, 0x6d, 0x54, 0x29, 0xe2, 0x29, 0x97, 0x80, 0x61,
	0x8c, 0xe3, 0x54, 0x49, 0x6e, 0xa7, 0x1a, 0x0a, 0xa6, 0x0d, 0x5a, 0x3c, 0xbd, 0xbd, 0x46, 0xd8,
	0x1a, 0x61, 0xbb, 0x08, 0xdb, 0x22, 0x41, 0x4f, 0x2a, 0x3b, 0x9a, 0x44, 0x2c, 0xc6, 0x9c, 0x4b,
	0x94, 0xc8, 0x6b, 0x3a, 0x9a, 0xa4, 0x75, 0x56, 0x27, 0x75, 0xb4, 0xee, 0x1a, 0x9c, 0x4b, 0x44,
	0x99, 0xc1, 0xb6, 0xca, 0xaa, 0x1c, 0x0a, 0x2b, 0x72, 0xdd, 0x14, 0xdc, 0x6c, 0x4f, 0x8a, 0xda,
	0x2a, 0x1c, 0x37, 0x33, 0x05, 0x37, 0xda, 0xe2, 0xce, 0xb8, 0xc1, 0x59, 0x5b, 0x2a, 0x45, 0xa6,
	0x12, 0x61, 0xa1, 0x51, 0x3b, 0x7b, 0xaa, 0x82, 0x97, 0x61, 0xbb, 0xf5, 0xf9, 0x61, 0x45, 0xb1,
	0x6b, 0xd0, 0x7d, 0xe3, 0x5f, 0x1e, 0x64, 0x18, 0x89, 0xec, 0xb1, 0x86, 0xf8, 0xc9, 0x54, 0xc3,
	0x69, 0xee, 0x9f, 0x68, 0x03, 0x29, 0x18, 0x03, 0x49, 0x68, 0x40, 0x2a, 0x1c, 0x5f, 0x27, 0x9d,
	0xa3, 0x3b, 0xc7, 0xf7, 0xce, 0x58, 0x7b, 0x79, 0x8f, 0xa2, 0xe7, 0x10, 0xdb, 0x21, 0xa4, 0xbf,
	0xb9, 0xfe, 0xad, 0xcf, 0xaf, 0xaf, 0xb5, 0xd6, 0x99, 0xe5, 0xa2, 0xe1, 0xbf, 0xfc, 0x9c, 0x1d,
	0x5d, 0xfc, 0x40, 0x2e, 0x9c, 0x90, 0xe1, 0x95, 0x4d, 0xef, 0x61, 0x2d, 0x75, 0xdf, 0x11, 0xff,
	0x78, 0x00, 0x76, 0x63, 0x1f, 0xff, 0xa7, 0x7d, 0xf0, 0x77, 0xfb, 0x03, 0xd3, 0x07, 0x57, 0xbf,
	0x3d, 0xdc, 0x7b, 0x76, 0xff, 0x23, 0x99, 0x2f, 0xa9, 0xb7, 0x58, 0x52, 0x6f, 0xb5, 0xa4, 0xe4,
	0xad, 0xa3, 0xe4, 0x93, 0xa3, 0xe4, 0xab, 0xa3, 0x64, 0xee, 0x28, 0x59, 0x38, 0x4a, 0xbe, 0x3b,
	0x4a, 0x7e, 0x38, 0xea, 0xad, 0x1c, 0x25, 0xef, 0x2b, 0xea, 0xcd, 0x2a, 0x4a, 0xe6, 0x15, 0xf5,
	0x16, 0x15, 0xf5, 0x9e, 0x3e, 0x93, 0xa8, 0x5f, 0x48, 0x56, 0x62, 0x66, 0xc1, 0x18, 0xc1, 0x26,
	0x05, 0xaf, 0x83, 0x14, 0x4d, 0xde, 0xd3, 0x06, 0x4b, 0x95, 0x80, 0xe9, 0xfd, 0x91, 0xb9, 0x8e,
	0x24, 0x72, 0x78, 0x65, 0x9b, 0xbb, 0xfc, 0xfb, 0x03, 0x47, 0x97, 0xea, 0x5b, 0xdd, 0xff, 0x15,
	0x00, 0x00, 0xff, 0xff, 0x7a, 0x65, 0x95, 0x1b, 0xf0, 0x02, 0x00, 0x00,
}

func (this *GlobalSpecType) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*GlobalSpecType)
	if !ok {
		that2, ok := that.(GlobalSpecType)
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
	if len(this.PreferredRegion) != len(that1.PreferredRegion) {
		return false
	}
	for i := range this.PreferredRegion {
		if !this.PreferredRegion[i].Equal(that1.PreferredRegion[i]) {
			return false
		}
	}
	return true
}
func (this *GetSpecType) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*GetSpecType)
	if !ok {
		that2, ok := that.(GetSpecType)
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
	if len(this.PreferredRegion) != len(that1.PreferredRegion) {
		return false
	}
	for i := range this.PreferredRegion {
		if !this.PreferredRegion[i].Equal(that1.PreferredRegion[i]) {
			return false
		}
	}
	return true
}
func (this *GlobalSpecType) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&geo_config.GlobalSpecType{")
	if this.PreferredRegion != nil {
		s = append(s, "PreferredRegion: "+fmt.Sprintf("%#v", this.PreferredRegion)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *GetSpecType) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&geo_config.GetSpecType{")
	if this.PreferredRegion != nil {
		s = append(s, "PreferredRegion: "+fmt.Sprintf("%#v", this.PreferredRegion)+",\n")
	}
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
func (m *GlobalSpecType) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GlobalSpecType) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GlobalSpecType) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.PreferredRegion) > 0 {
		for iNdEx := len(m.PreferredRegion) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.PreferredRegion[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTypes(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *GetSpecType) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetSpecType) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetSpecType) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.PreferredRegion) > 0 {
		for iNdEx := len(m.PreferredRegion) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.PreferredRegion[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTypes(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
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
func (m *GlobalSpecType) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.PreferredRegion) > 0 {
		for _, e := range m.PreferredRegion {
			l = e.Size()
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	return n
}

func (m *GetSpecType) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.PreferredRegion) > 0 {
		for _, e := range m.PreferredRegion {
			l = e.Size()
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	return n
}

func sovTypes(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTypes(x uint64) (n int) {
	return sovTypes(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *GlobalSpecType) String() string {
	if this == nil {
		return "nil"
	}
	repeatedStringForPreferredRegion := "[]*ObjectRefType{"
	for _, f := range this.PreferredRegion {
		repeatedStringForPreferredRegion += strings.Replace(fmt.Sprintf("%v", f), "ObjectRefType", "schema.ObjectRefType", 1) + ","
	}
	repeatedStringForPreferredRegion += "}"
	s := strings.Join([]string{`&GlobalSpecType{`,
		`PreferredRegion:` + repeatedStringForPreferredRegion + `,`,
		`}`,
	}, "")
	return s
}
func (this *GetSpecType) String() string {
	if this == nil {
		return "nil"
	}
	repeatedStringForPreferredRegion := "[]*ObjectRefType{"
	for _, f := range this.PreferredRegion {
		repeatedStringForPreferredRegion += strings.Replace(fmt.Sprintf("%v", f), "ObjectRefType", "schema.ObjectRefType", 1) + ","
	}
	repeatedStringForPreferredRegion += "}"
	s := strings.Join([]string{`&GetSpecType{`,
		`PreferredRegion:` + repeatedStringForPreferredRegion + `,`,
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
func (m *GlobalSpecType) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: GlobalSpecType: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GlobalSpecType: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PreferredRegion", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PreferredRegion = append(m.PreferredRegion, &schema.ObjectRefType{})
			if err := m.PreferredRegion[len(m.PreferredRegion)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
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
func (m *GetSpecType) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: GetSpecType: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetSpecType: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PreferredRegion", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PreferredRegion = append(m.PreferredRegion, &schema.ObjectRefType{})
			if err := m.PreferredRegion[len(m.PreferredRegion)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
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
