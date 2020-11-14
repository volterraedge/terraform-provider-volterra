// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ves.io/schema/namespace/types.proto

package namespace

import (
	proto "github.com/gogo/protobuf/proto"
	golang_proto "github.com/golang/protobuf/proto"

	fmt "fmt"

	math "math"

	_ "github.com/gogo/protobuf/gogoproto"

	_ "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema"

	_ "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema"

	strings "strings"

	reflect "reflect"

	io "io"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = golang_proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// GlobalSpecType
//
// x-displayName: "Global Specification"
// This is the shape of the namespace representation in the database at Global Controller
type GlobalSpecType struct {
}

func (m *GlobalSpecType) Reset()                    { *m = GlobalSpecType{} }
func (*GlobalSpecType) ProtoMessage()               {}
func (*GlobalSpecType) Descriptor() ([]byte, []int) { return fileDescriptorTypes, []int{0} }

// Create namespace
//
// x-displayName: "Create Specification"
// Creates a new namespace. Name of the object is name of the name space.
type CreateSpecType struct {
}

func (m *CreateSpecType) Reset()                    { *m = CreateSpecType{} }
func (*CreateSpecType) ProtoMessage()               {}
func (*CreateSpecType) Descriptor() ([]byte, []int) { return fileDescriptorTypes, []int{1} }

// Replace namespace
//
// x-displayName: "Replace Specification"
// Replaces attributes of a namespace including its metadata like labels, description etc.
type ReplaceSpecType struct {
}

func (m *ReplaceSpecType) Reset()                    { *m = ReplaceSpecType{} }
func (*ReplaceSpecType) ProtoMessage()               {}
func (*ReplaceSpecType) Descriptor() ([]byte, []int) { return fileDescriptorTypes, []int{2} }

// Get namespace
//
// x-displayName: "Get Specification"
// This is the read representation of the namespace object.
type GetSpecType struct {
}

func (m *GetSpecType) Reset()                    { *m = GetSpecType{} }
func (*GetSpecType) ProtoMessage()               {}
func (*GetSpecType) Descriptor() ([]byte, []int) { return fileDescriptorTypes, []int{3} }

func init() {
	proto.RegisterType((*GlobalSpecType)(nil), "ves.io.schema.namespace.GlobalSpecType")
	golang_proto.RegisterType((*GlobalSpecType)(nil), "ves.io.schema.namespace.GlobalSpecType")
	proto.RegisterType((*CreateSpecType)(nil), "ves.io.schema.namespace.CreateSpecType")
	golang_proto.RegisterType((*CreateSpecType)(nil), "ves.io.schema.namespace.CreateSpecType")
	proto.RegisterType((*ReplaceSpecType)(nil), "ves.io.schema.namespace.ReplaceSpecType")
	golang_proto.RegisterType((*ReplaceSpecType)(nil), "ves.io.schema.namespace.ReplaceSpecType")
	proto.RegisterType((*GetSpecType)(nil), "ves.io.schema.namespace.GetSpecType")
	golang_proto.RegisterType((*GetSpecType)(nil), "ves.io.schema.namespace.GetSpecType")
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
	return true
}
func (this *CreateSpecType) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*CreateSpecType)
	if !ok {
		that2, ok := that.(CreateSpecType)
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
func (this *ReplaceSpecType) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ReplaceSpecType)
	if !ok {
		that2, ok := that.(ReplaceSpecType)
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
	return true
}
func (this *GlobalSpecType) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 4)
	s = append(s, "&namespace.GlobalSpecType{")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *CreateSpecType) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 4)
	s = append(s, "&namespace.CreateSpecType{")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *ReplaceSpecType) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 4)
	s = append(s, "&namespace.ReplaceSpecType{")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *GetSpecType) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 4)
	s = append(s, "&namespace.GetSpecType{")
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
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GlobalSpecType) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func (m *CreateSpecType) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CreateSpecType) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func (m *ReplaceSpecType) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ReplaceSpecType) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func (m *GetSpecType) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetSpecType) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func encodeVarintTypes(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func NewPopulatedGlobalSpecType(r randyTypes, easy bool) *GlobalSpecType {
	this := &GlobalSpecType{}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedCreateSpecType(r randyTypes, easy bool) *CreateSpecType {
	this := &CreateSpecType{}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedReplaceSpecType(r randyTypes, easy bool) *ReplaceSpecType {
	this := &ReplaceSpecType{}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedGetSpecType(r randyTypes, easy bool) *GetSpecType {
	this := &GetSpecType{}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

type randyTypes interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneTypes(r randyTypes) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringTypes(r randyTypes) string {
	v1 := r.Intn(100)
	tmps := make([]rune, v1)
	for i := 0; i < v1; i++ {
		tmps[i] = randUTF8RuneTypes(r)
	}
	return string(tmps)
}
func randUnrecognizedTypes(r randyTypes, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldTypes(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldTypes(dAtA []byte, r randyTypes, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateTypes(dAtA, uint64(key))
		v2 := r.Int63()
		if r.Intn(2) == 0 {
			v2 *= -1
		}
		dAtA = encodeVarintPopulateTypes(dAtA, uint64(v2))
	case 1:
		dAtA = encodeVarintPopulateTypes(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateTypes(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateTypes(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateTypes(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateTypes(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *GlobalSpecType) Size() (n int) {
	var l int
	_ = l
	return n
}

func (m *CreateSpecType) Size() (n int) {
	var l int
	_ = l
	return n
}

func (m *ReplaceSpecType) Size() (n int) {
	var l int
	_ = l
	return n
}

func (m *GetSpecType) Size() (n int) {
	var l int
	_ = l
	return n
}

func sovTypes(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozTypes(x uint64) (n int) {
	return sovTypes(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *GlobalSpecType) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&GlobalSpecType{`,
		`}`,
	}, "")
	return s
}
func (this *CreateSpecType) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&CreateSpecType{`,
		`}`,
	}, "")
	return s
}
func (this *ReplaceSpecType) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ReplaceSpecType{`,
		`}`,
	}, "")
	return s
}
func (this *GetSpecType) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&GetSpecType{`,
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
			wire |= (uint64(b) & 0x7F) << shift
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
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
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
func (m *CreateSpecType) Unmarshal(dAtA []byte) error {
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
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CreateSpecType: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CreateSpecType: illegal tag %d (wire type %d)", fieldNum, wire)
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
func (m *ReplaceSpecType) Unmarshal(dAtA []byte) error {
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
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ReplaceSpecType: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ReplaceSpecType: illegal tag %d (wire type %d)", fieldNum, wire)
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
			wire |= (uint64(b) & 0x7F) << shift
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
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
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
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthTypes
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowTypes
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
				next, err := skipTypes(dAtA[start:])
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
	ErrInvalidLengthTypes = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTypes   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("ves.io/schema/namespace/types.proto", fileDescriptorTypes) }
func init() { golang_proto.RegisterFile("ves.io/schema/namespace/types.proto", fileDescriptorTypes) }

var fileDescriptorTypes = []byte{
	// 297 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0xaf, 0x4e, 0x33, 0x41,
	0x14, 0xc5, 0xf7, 0x9a, 0x4f, 0xec, 0x97, 0x94, 0x3f, 0x86, 0x50, 0x92, 0x1b, 0xd2, 0x62, 0x3b,
	0x23, 0x70, 0x08, 0x04, 0x88, 0x2a, 0x0c, 0xa0, 0x70, 0xb3, 0xcb, 0xed, 0x74, 0xc3, 0x6e, 0xef,
	0x64, 0x76, 0xda, 0x50, 0xc7, 0x1b, 0xc0, 0x63, 0xf0, 0x08, 0xc8, 0x4a, 0x82, 0xaa, 0xac, 0x64,
	0xa7, 0x06, 0x59, 0x89, 0x24, 0x99, 0x6d, 0x09, 0xad, 0xa8, 0x3b, 0x67, 0xce, 0x2f, 0x73, 0x92,
	0x73, 0xe3, 0xf6, 0x88, 0x4a, 0x91, 0xb1, 0x2c, 0xd3, 0x3e, 0x15, 0x4a, 0x0e, 0x54, 0x41, 0xa5,
	0x51, 0x29, 0x49, 0x37, 0x36, 0x54, 0x0a, 0x63, 0xd9, 0xf1, 0xfe, 0x41, 0x0d, 0x89, 0x1a, 0x12,
	0xbf, 0x50, 0xb3, 0xa3, 0x33, 0xd7, 0x1f, 0x26, 0x22, 0xe5, 0x42, 0x6a, 0xd6, 0x2c, 0x03, 0x9f,
	0x0c, 0x7b, 0xc1, 0x05, 0x13, 0x54, 0xfd, 0x4f, 0xf3, 0x68, 0xbd, 0x8c, 0x8d, 0xcb, 0x78, 0xb0,
	0x2c, 0x69, 0x1e, 0xae, 0x87, 0x7f, 0xfa, 0x5b, 0xbb, 0x71, 0xa3, 0x9b, 0x73, 0xa2, 0xf2, 0x1b,
	0x43, 0xe9, 0xed, 0xd8, 0x50, 0xab, 0x1d, 0x37, 0x2e, 0x2d, 0x29, 0x47, 0xab, 0x97, 0xb3, 0xbd,
	0x8f, 0xf3, 0x4d, 0xe8, 0x24, 0xde, 0xb9, 0x26, 0x93, 0xab, 0x74, 0x2b, 0x75, 0x1c, 0xff, 0xef,
	0x92, 0xdb, 0x42, 0x5c, 0x3c, 0xc3, 0xb4, 0xc2, 0x68, 0x56, 0x61, 0xb4, 0xa8, 0x10, 0xbe, 0x2b,
	0x84, 0x27, 0x8f, 0xf0, 0xea, 0x11, 0xde, 0x3c, 0xc2, 0xc4, 0x23, 0xbc, 0x7b, 0x84, 0xa9, 0x47,
	0x98, 0x79, 0x84, 0x4f, 0x8f, 0xf0, 0xe5, 0x31, 0x5a, 0x78, 0x84, 0x97, 0x39, 0x46, 0x93, 0x39,
	0xc2, 0xdd, 0x95, 0x66, 0xf3, 0xa0, 0xc5, 0x88, 0x73, 0x47, 0xd6, 0x2a, 0x31, 0x2c, 0x65, 0x10,
	0x3d, 0xb6, 0x45, 0xc7, 0x58, 0x1e, 0x65, 0xf7, 0x64, 0x3b, 0xab, 0x58, 0x9a, 0x44, 0xb3, 0xa4,
	0x47, 0xb7, 0xdc, 0x62, 0xf3, 0x38, 0xc9, 0xbf, 0xb0, 0xcb, 0xe9, 0x4f, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x11, 0xd1, 0xf2, 0x42, 0xbe, 0x01, 0x00, 0x00,
}
