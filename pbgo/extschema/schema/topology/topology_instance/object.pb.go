// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ves.io/schema/topology/topology_instance/object.proto

package topology_instance

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

// Topology Instance
//
// x-displayName: "Topology Instance"
// topology instance object
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
	// Specification of the desired behavior of the topology instance
	Spec *SpecType `protobuf:"bytes,3,opt,name=spec,proto3" json:"spec,omitempty"`
}

func (m *Object) Reset()      { *m = Object{} }
func (*Object) ProtoMessage() {}
func (*Object) Descriptor() ([]byte, []int) {
	return fileDescriptor_f53d96024d9e4d34, []int{0}
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

// Topology instance specification
//
// x-displayName: "instance specification"
// Shape of the topology instance specification
type SpecType struct {
	// gc_spec
	//
	// x-displayName: "GC Spec"
	GcSpec *GlobalSpecType `protobuf:"bytes,2,opt,name=gc_spec,json=gcSpec,proto3" json:"gc_spec,omitempty"`
}

func (m *SpecType) Reset()      { *m = SpecType{} }
func (*SpecType) ProtoMessage() {}
func (*SpecType) Descriptor() ([]byte, []int) {
	return fileDescriptor_f53d96024d9e4d34, []int{1}
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
	proto.RegisterType((*Object)(nil), "ves.io.schema.topology.topology_instance.Object")
	golang_proto.RegisterType((*Object)(nil), "ves.io.schema.topology.topology_instance.Object")
	proto.RegisterType((*SpecType)(nil), "ves.io.schema.topology.topology_instance.SpecType")
	golang_proto.RegisterType((*SpecType)(nil), "ves.io.schema.topology.topology_instance.SpecType")
}

func init() {
	proto.RegisterFile("ves.io/schema/topology/topology_instance/object.proto", fileDescriptor_f53d96024d9e4d34)
}
func init() {
	golang_proto.RegisterFile("ves.io/schema/topology/topology_instance/object.proto", fileDescriptor_f53d96024d9e4d34)
}

var fileDescriptor_f53d96024d9e4d34 = []byte{
	// 522 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x94, 0x31, 0x6f, 0x13, 0x31,
	0x14, 0xc7, 0xcf, 0x34, 0xa4, 0x95, 0x91, 0x40, 0xba, 0x29, 0x14, 0x61, 0x55, 0x61, 0xe9, 0xc0,
	0xf9, 0xa4, 0x42, 0x25, 0x60, 0x60, 0x00, 0x44, 0x97, 0x56, 0x88, 0x96, 0xa9, 0x12, 0x44, 0x3e,
	0xdf, 0x8b, 0x6b, 0xb8, 0x3b, 0x5b, 0xb6, 0x73, 0x34, 0x1b, 0x03, 0x0b, 0x1b, 0xdf, 0x80, 0x95,
	0x8f, 0x51, 0x36, 0xc6, 0x88, 0x29, 0x23, 0xb9, 0x08, 0x89, 0xb1, 0x1f, 0x01, 0xd5, 0xd7, 0x04,
	0x12, 0x10, 0xcd, 0x92, 0xed, 0x59, 0xef, 0xff, 0xff, 0xfd, 0x7d, 0xe7, 0x67, 0xe3, 0xed, 0x12,
	0x2c, 0x95, 0x2a, 0xb6, 0xfc, 0x08, 0x72, 0x16, 0x3b, 0xa5, 0x55, 0xa6, 0x44, 0x7f, 0x5a, 0x74,
	0x64, 0x61, 0x1d, 0x2b, 0x38, 0xc4, 0x2a, 0x79, 0x0d, 0xdc, 0x51, 0x6d, 0x94, 0x53, 0xe1, 0x66,
	0x6d, 0xa3, 0xb5, 0x8d, 0x4e, 0xd4, 0xf4, 0x2f, 0xdb, 0x7a, 0x24, 0xa4, 0x3b, 0xea, 0x25, 0x94,
	0xab, 0x3c, 0x16, 0x4a, 0xa8, 0xd8, 0x03, 0x92, 0x5e, 0xd7, 0xaf, 0xfc, 0xc2, 0x57, 0x35, 0x78,
	0xfd, 0xc6, 0xec, 0x7e, 0x94, 0x76, 0x52, 0x15, 0xf6, 0xbc, 0x79, 0x77, 0xe1, 0xcd, 0xba, 0xbe,
	0x86, 0x89, 0xeb, 0xfa, 0x9c, 0xeb, 0x8f, 0x56, 0x7b, 0xb6, 0x55, 0x82, 0x85, 0xa2, 0x9c, 0x0d,
	0x6d, 0xff, 0xb8, 0x8c, 0x9b, 0xcf, 0xfc, 0xb7, 0x87, 0xf7, 0xf1, 0x5a, 0x0e, 0x8e, 0xa5, 0xcc,
	0xb1, 0x16, 0xda, 0x40, 0x9b, 0x57, 0xb6, 0x6e, 0xd2, 0xd9, 0x1f, 0x51, 0x0b, 0xf7, 0xc0, 0xb1,
	0x17, 0x7d, 0x0d, 0xfb, 0x53, 0x79, 0xb8, 0x8b, 0xaf, 0xd9, 0xbe, 0x75, 0x90, 0x77, 0xa6, 0x84,
	0x4b, 0x9e, 0x70, 0x6b, 0x8e, 0x70, 0xe0, 0x55, 0x73, 0x9c, 0xab, 0xb5, 0x77, 0x6f, 0x42, 0x7b,
	0x8a, 0x1b, 0x56, 0x03, 0x6f, 0xad, 0x78, 0xc4, 0x16, 0x5d, 0xf4, 0x34, 0xe8, 0x81, 0x06, 0xee,
	0x89, 0xde, 0xff, 0xe0, 0x43, 0xe3, 0xdb, 0x97, 0xd6, 0xfb, 0x06, 0xee, 0xe2, 0x95, 0x1d, 0x70,
	0x61, 0xa7, 0xbb, 0x7d, 0xcc, 0x23, 0x2b, 0x1d, 0x44, 0x39, 0x2b, 0x98, 0x80, 0x1c, 0x0a, 0x17,
	0x25, 0xcc, 0x4a, 0x1e, 0xe5, 0xaa, 0x90, 0x4e, 0x99, 0xdb, 0x1b, 0xff, 0xd4, 0x9c, 0xe1, 0x53,
	0x66, 0xd2, 0xdf, 0xb2, 0x12, 0x6c, 0x24, 0x55, 0x24, 0x0b, 0x07, 0xa6, 0x60, 0x59, 0x64, 0x80,
	0xa5, 0x18, 0xf0, 0xea, 0x3e, 0xe8, 0x8c, 0x71, 0x08, 0x0f, 0xff, 0x93, 0xd5, 0xb3, 0x70, 0x61,
	0x50, 0xad, 0x99, 0x4f, 0x79, 0x6b, 0xa4, 0x03, 0x9c, 0xe2, 0xe6, 0x63, 0x03, 0xcc, 0x2d, 0x3d,
	0xe5, 0x09, 0x64, 0xb0, 0xe4, 0x14, 0x81, 0x1b, 0xbb, 0xd2, 0x2e, 0xff, 0x6c, 0xda, 0x2f, 0xf1,
	0xda, 0x64, 0x3a, 0xc2, 0xe7, 0x78, 0x55, 0xf0, 0x8e, 0x1f, 0xb1, 0x7a, 0x4a, 0xef, 0x2d, 0x3e,
	0x62, 0x3b, 0x99, 0x4a, 0x58, 0x36, 0x1d, 0xb4, 0xa6, 0xe0, 0x67, 0xf5, 0xa3, 0x4f, 0xe8, 0xe4,
	0x61, 0x30, 0x18, 0x91, 0x60, 0x38, 0x22, 0xc1, 0xe9, 0x88, 0xa0, 0x77, 0x15, 0x41, 0x9f, 0x2b,
	0x82, 0xbe, 0x56, 0x04, 0x0d, 0x2a, 0x82, 0x86, 0x15, 0x41, 0xdf, 0x2b, 0x82, 0x7e, 0x56, 0x24,
	0x38, 0xad, 0x08, 0xfa, 0x38, 0x26, 0xc1, 0xc9, 0x98, 0xa0, 0xc1, 0x98, 0x04, 0xc3, 0x31, 0x09,
	0x0e, 0x5f, 0x09, 0xa5, 0xdf, 0x08, 0x5a, 0xaa, 0xcc, 0x81, 0x31, 0x8c, 0xf6, 0x6c, 0xec, 0x8b,
	0xae, 0x32, 0x79, 0xa4, 0x8d, 0x2a, 0x65, 0x0a, 0x26, 0x9a, 0xb4, 0x63, 0x9d, 0x08, 0x15, 0xc3,
	0xb1, 0x3b, 0xbf, 0xdb, 0x17, 0xbe, 0x19, 0x49, 0xd3, 0xdf, 0xf7, 0x3b, 0xbf, 0x02, 0x00, 0x00,
	0xff, 0xff, 0xa1, 0x96, 0x2d, 0xbc, 0x13, 0x05, 0x00, 0x00,
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
	s = append(s, "&topology_instance.Object{")
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
	s = append(s, "&topology_instance.SpecType{")
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
