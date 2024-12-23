// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ves.io/schema/topology/topology_region/object.proto

package topology_region

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	golang_proto "github.com/golang/protobuf/proto"
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

// Topology Region
//
// x-displayName: "Topology Region"
// topology region object
type Object struct {
	// metadata
	//
	// x-displayName: "Metadata"
	// Standard object's metadata
	Metadata *schema.ObjectMetaType `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// system_metadata
	//
	// x-displayName: "System Metadata"
	// System generated object's metadata
	SystemMetadata *schema.SystemObjectMetaType `protobuf:"bytes,2,opt,name=system_metadata,json=systemMetadata,proto3" json:"system_metadata,omitempty"`
	// spec
	//
	// x-displayName: "Spec"
	// Specification of the desired behavior of the topology region
	Spec *SpecType `protobuf:"bytes,3,opt,name=spec,proto3" json:"spec,omitempty"`
}

func (m *Object) Reset()      { *m = Object{} }
func (*Object) ProtoMessage() {}
func (*Object) Descriptor() ([]byte, []int) {
	return fileDescriptor_9ec8d5e3ffc7cc29, []int{0}
}
func (m *Object) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Object) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *Object) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Object.Merge(m, src)
}
func (m *Object) XXX_Size() int {
	return m.Size()
}
func (m *Object) XXX_DiscardUnknown() {
	xxx_messageInfo_Object.DiscardUnknown(m)
}

var xxx_messageInfo_Object proto.InternalMessageInfo

func (m *Object) GetMetadata() *schema.ObjectMetaType {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *Object) GetSystemMetadata() *schema.SystemObjectMetaType {
	if m != nil {
		return m.SystemMetadata
	}
	return nil
}

func (m *Object) GetSpec() *SpecType {
	if m != nil {
		return m.Spec
	}
	return nil
}

// Topology region specification
//
// x-displayName: "region specification"
// Shape of the topology region specification
type SpecType struct {
	// gc_spec
	//
	// x-displayName: "GC Spec"
	GcSpec *GlobalSpecType `protobuf:"bytes,2,opt,name=gc_spec,json=gcSpec,proto3" json:"gc_spec,omitempty"`
}

func (m *SpecType) Reset()      { *m = SpecType{} }
func (*SpecType) ProtoMessage() {}
func (*SpecType) Descriptor() ([]byte, []int) {
	return fileDescriptor_9ec8d5e3ffc7cc29, []int{1}
}
func (m *SpecType) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SpecType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *SpecType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SpecType.Merge(m, src)
}
func (m *SpecType) XXX_Size() int {
	return m.Size()
}
func (m *SpecType) XXX_DiscardUnknown() {
	xxx_messageInfo_SpecType.DiscardUnknown(m)
}

var xxx_messageInfo_SpecType proto.InternalMessageInfo

func (m *SpecType) GetGcSpec() *GlobalSpecType {
	if m != nil {
		return m.GcSpec
	}
	return nil
}

func init() {
	proto.RegisterType((*Object)(nil), "ves.io.schema.topology.topology_region.Object")
	golang_proto.RegisterType((*Object)(nil), "ves.io.schema.topology.topology_region.Object")
	proto.RegisterType((*SpecType)(nil), "ves.io.schema.topology.topology_region.SpecType")
	golang_proto.RegisterType((*SpecType)(nil), "ves.io.schema.topology.topology_region.SpecType")
}

func init() {
	proto.RegisterFile("ves.io/schema/topology/topology_region/object.proto", fileDescriptor_9ec8d5e3ffc7cc29)
}
func init() {
	golang_proto.RegisterFile("ves.io/schema/topology/topology_region/object.proto", fileDescriptor_9ec8d5e3ffc7cc29)
}

var fileDescriptor_9ec8d5e3ffc7cc29 = []byte{
	// 507 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x93, 0x41, 0x6b, 0x13, 0x41,
	0x14, 0xc7, 0x77, 0xd2, 0x90, 0x96, 0x11, 0x14, 0xf6, 0x14, 0x2b, 0x0e, 0x25, 0x82, 0x78, 0x70,
	0x77, 0xa5, 0x45, 0xa1, 0x0a, 0x1e, 0xda, 0x40, 0x2f, 0x2d, 0x85, 0xd6, 0x93, 0x0a, 0x61, 0x76,
	0xf3, 0x32, 0x1d, 0xdd, 0xdd, 0x37, 0xcc, 0x4c, 0xd6, 0xe6, 0xe6, 0x47, 0xf0, 0x0b, 0x78, 0xf7,
	0x23, 0x78, 0xac, 0x37, 0x8f, 0xc1, 0x53, 0xbc, 0x88, 0xd9, 0x5c, 0x3c, 0xf6, 0x23, 0x88, 0xb3,
	0x4d, 0x30, 0x41, 0x6c, 0x20, 0xb7, 0x37, 0xbc, 0xff, 0xff, 0xb7, 0x6f, 0xdf, 0x9f, 0x47, 0x77,
	0x0a, 0x30, 0xa1, 0xc4, 0xc8, 0x24, 0x67, 0x90, 0xf1, 0xc8, 0xa2, 0xc2, 0x14, 0xc5, 0x60, 0x56,
	0x74, 0x34, 0x08, 0x89, 0x79, 0x84, 0xf1, 0x1b, 0x48, 0x6c, 0xa8, 0x34, 0x5a, 0xf4, 0xef, 0x57,
	0xa6, 0xb0, 0x32, 0x85, 0x53, 0x6d, 0xb8, 0x60, 0xda, 0x0c, 0x84, 0xb4, 0x67, 0xfd, 0x38, 0x4c,
	0x30, 0x8b, 0x04, 0x0a, 0x8c, 0x9c, 0x3d, 0xee, 0xf7, 0xdc, 0xcb, 0x3d, 0x5c, 0x55, 0x61, 0x37,
	0xef, 0xcc, 0xcf, 0x82, 0xca, 0x4a, 0xcc, 0xcd, 0x55, 0x73, 0x7b, 0xc9, 0x41, 0xed, 0x40, 0xc1,
	0xd4, 0x73, 0x7b, 0xc1, 0xf3, 0x57, 0xab, 0x35, 0xdf, 0x2a, 0xc0, 0x40, 0x5e, 0xcc, 0x7f, 0xb2,
	0xf5, 0xb9, 0x4e, 0x1b, 0xc7, 0xee, 0xbf, 0xfd, 0x5d, 0xba, 0x91, 0x81, 0xe5, 0x5d, 0x6e, 0x79,
	0x93, 0x6c, 0x91, 0x07, 0x37, 0xb6, 0xef, 0x86, 0xf3, 0x4b, 0xa8, 0x84, 0x47, 0x60, 0xf9, 0x8b,
	0x81, 0x82, 0x93, 0x99, 0xdc, 0x3f, 0xa4, 0xb7, 0xcc, 0xc0, 0x58, 0xc8, 0x3a, 0x33, 0x42, 0xcd,
	0x11, 0xee, 0x2d, 0x10, 0x4e, 0x9d, 0x6a, 0x81, 0x73, 0xb3, 0xf2, 0x1e, 0x4d, 0x69, 0x6d, 0x5a,
	0x37, 0x0a, 0x92, 0xe6, 0x9a, 0x43, 0x3c, 0x0a, 0x97, 0x4b, 0x22, 0x3c, 0x55, 0x90, 0x38, 0x9e,
	0x73, 0x3f, 0xfd, 0x51, 0xfb, 0xf6, 0xa5, 0xf9, 0xbd, 0x46, 0xf7, 0xe8, 0xda, 0x01, 0x58, 0xff,
	0x59, 0xef, 0xf1, 0x79, 0x12, 0x18, 0x69, 0x21, 0xc8, 0x78, 0xce, 0x05, 0x64, 0x90, 0xdb, 0xc0,
	0x58, 0x9e, 0x77, 0xb9, 0xee, 0x06, 0x19, 0xe6, 0xd2, 0xa2, 0x7e, 0xb8, 0x55, 0x80, 0x09, 0x24,
	0x06, 0x32, 0xb7, 0xa0, 0x73, 0x9e, 0x06, 0x1a, 0x78, 0x97, 0xb6, 0xe9, 0xfa, 0x09, 0xa8, 0x94,
	0x27, 0xe0, 0xef, 0xfe, 0x9f, 0xd3, 0x37, 0xf0, 0x0f, 0xc8, 0x3b, 0x2d, 0x2d, 0xd0, 0x7d, 0xda,
	0xd8, 0xd7, 0xc0, 0xed, 0xaa, 0x90, 0x36, 0xa4, 0xb0, 0x2a, 0xa4, 0x7e, 0x28, 0xcd, 0x6a, 0x4b,
	0x69, 0xbd, 0xa2, 0x1b, 0xd3, 0x95, 0xfb, 0xc7, 0x74, 0x5d, 0x24, 0x1d, 0x97, 0x5a, 0x15, 0xfc,
	0x93, 0x65, 0x53, 0x3b, 0x48, 0x31, 0xe6, 0xe9, 0x2c, 0xbb, 0x86, 0x48, 0xfe, 0xd4, 0x7b, 0x1f,
	0xc9, 0xc5, 0x73, 0x6f, 0x38, 0x66, 0xde, 0x68, 0xcc, 0xbc, 0xcb, 0x31, 0x23, 0xef, 0x4b, 0x46,
	0x3e, 0x95, 0x8c, 0x7c, 0x2d, 0x19, 0x19, 0x96, 0x8c, 0x8c, 0x4a, 0x46, 0x7e, 0x96, 0x8c, 0xfc,
	0x2a, 0x99, 0x77, 0x59, 0x32, 0xf2, 0x61, 0xc2, 0xbc, 0x8b, 0x09, 0x23, 0xc3, 0x09, 0xf3, 0x46,
	0x13, 0xe6, 0xbd, 0x7c, 0x2d, 0x50, 0xbd, 0x15, 0x61, 0x81, 0xa9, 0x05, 0xad, 0x79, 0xd8, 0x37,
	0x91, 0x2b, 0x7a, 0xa8, 0xb3, 0x40, 0x69, 0x2c, 0x64, 0x17, 0x74, 0x30, 0x6d, 0x47, 0x2a, 0x16,
	0x18, 0xc1, 0xb9, 0xbd, 0x3a, 0x96, 0x6b, 0x4e, 0x30, 0x6e, 0xb8, 0xf3, 0xd9, 0xf9, 0x1d, 0x00,
	0x00, 0xff, 0xff, 0x16, 0xd4, 0x9c, 0xfe, 0x5c, 0x04, 0x00, 0x00,
}

func (this *Object) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Object)
	if !ok {
		that2, ok := that.(Object)
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
	if !this.Metadata.Equal(that1.Metadata) {
		return false
	}
	if !this.SystemMetadata.Equal(that1.SystemMetadata) {
		return false
	}
	if !this.Spec.Equal(that1.Spec) {
		return false
	}
	return true
}
func (this *SpecType) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*SpecType)
	if !ok {
		that2, ok := that.(SpecType)
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
	if !this.GcSpec.Equal(that1.GcSpec) {
		return false
	}
	return true
}
func (this *Object) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 7)
	s = append(s, "&topology_region.Object{")
	if this.Metadata != nil {
		s = append(s, "Metadata: "+fmt.Sprintf("%#v", this.Metadata)+",\n")
	}
	if this.SystemMetadata != nil {
		s = append(s, "SystemMetadata: "+fmt.Sprintf("%#v", this.SystemMetadata)+",\n")
	}
	if this.Spec != nil {
		s = append(s, "Spec: "+fmt.Sprintf("%#v", this.Spec)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *SpecType) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&topology_region.SpecType{")
	if this.GcSpec != nil {
		s = append(s, "GcSpec: "+fmt.Sprintf("%#v", this.GcSpec)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringObject(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *Object) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Object) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Object) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Spec != nil {
		{
			size, err := m.Spec.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintObject(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.SystemMetadata != nil {
		{
			size, err := m.SystemMetadata.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintObject(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.Metadata != nil {
		{
			size, err := m.Metadata.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintObject(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *SpecType) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SpecType) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SpecType) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.GcSpec != nil {
		{
			size, err := m.GcSpec.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintObject(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	return len(dAtA) - i, nil
}

func encodeVarintObject(dAtA []byte, offset int, v uint64) int {
	offset -= sovObject(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Object) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Metadata != nil {
		l = m.Metadata.Size()
		n += 1 + l + sovObject(uint64(l))
	}
	if m.SystemMetadata != nil {
		l = m.SystemMetadata.Size()
		n += 1 + l + sovObject(uint64(l))
	}
	if m.Spec != nil {
		l = m.Spec.Size()
		n += 1 + l + sovObject(uint64(l))
	}
	return n
}

func (m *SpecType) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.GcSpec != nil {
		l = m.GcSpec.Size()
		n += 1 + l + sovObject(uint64(l))
	}
	return n
}

func sovObject(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozObject(x uint64) (n int) {
	return sovObject(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *Object) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Object{`,
		`Metadata:` + strings.Replace(fmt.Sprintf("%v", this.Metadata), "ObjectMetaType", "schema.ObjectMetaType", 1) + `,`,
		`SystemMetadata:` + strings.Replace(fmt.Sprintf("%v", this.SystemMetadata), "SystemObjectMetaType", "schema.SystemObjectMetaType", 1) + `,`,
		`Spec:` + strings.Replace(this.Spec.String(), "SpecType", "SpecType", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *SpecType) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&SpecType{`,
		`GcSpec:` + strings.Replace(fmt.Sprintf("%v", this.GcSpec), "GlobalSpecType", "GlobalSpecType", 1) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringObject(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *Object) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowObject
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
			return fmt.Errorf("proto: Object: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Object: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Metadata", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowObject
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
				return ErrInvalidLengthObject
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthObject
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Metadata == nil {
				m.Metadata = &schema.ObjectMetaType{}
			}
			if err := m.Metadata.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SystemMetadata", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowObject
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
				return ErrInvalidLengthObject
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthObject
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.SystemMetadata == nil {
				m.SystemMetadata = &schema.SystemObjectMetaType{}
			}
			if err := m.SystemMetadata.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Spec", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowObject
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
				return ErrInvalidLengthObject
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthObject
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Spec == nil {
				m.Spec = &SpecType{}
			}
			if err := m.Spec.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipObject(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthObject
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthObject
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
func (m *SpecType) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowObject
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
			return fmt.Errorf("proto: SpecType: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SpecType: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GcSpec", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowObject
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
				return ErrInvalidLengthObject
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthObject
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.GcSpec == nil {
				m.GcSpec = &GlobalSpecType{}
			}
			if err := m.GcSpec.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipObject(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthObject
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthObject
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
func skipObject(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowObject
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
					return 0, ErrIntOverflowObject
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
					return 0, ErrIntOverflowObject
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
				return 0, ErrInvalidLengthObject
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupObject
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthObject
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthObject        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowObject          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupObject = fmt.Errorf("proto: unexpected end of group")
)
