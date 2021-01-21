// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ves.io/schema/alert_policy_set/types.proto

package alert_policy_set

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import _ "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema"
import ves_io_schema4 "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema"
import _ "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema"
import _ "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema"

import strings "strings"
import reflect "reflect"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Specification for Alert Policy Set
//
// x-displayName: "Specification"
// Shape of the Alert Policy Set object
type GlobalSpecType struct {
	// Alert Policies
	//
	// x-displayName: "Alert Policies"
	// x-required
	// Reference to a list of Alert Policy objects to send the alert.
	Policies []*ves_io_schema4.ObjectRefType `protobuf:"bytes,1,rep,name=policies" json:"policies,omitempty"`
}

func (m *GlobalSpecType) Reset()                    { *m = GlobalSpecType{} }
func (*GlobalSpecType) ProtoMessage()               {}
func (*GlobalSpecType) Descriptor() ([]byte, []int) { return fileDescriptorTypes, []int{0} }

func (m *GlobalSpecType) GetPolicies() []*ves_io_schema4.ObjectRefType {
	if m != nil {
		return m.Policies
	}
	return nil
}

// Create Alert Policy Set
//
// x-displayName: "Create Alert Policy Set"
// Creates a new Alert Policy Set Object
type CreateSpecType struct {
	Policies []*ves_io_schema4.ObjectRefType `protobuf:"bytes,1,rep,name=policies" json:"policies,omitempty"`
}

func (m *CreateSpecType) Reset()                    { *m = CreateSpecType{} }
func (*CreateSpecType) ProtoMessage()               {}
func (*CreateSpecType) Descriptor() ([]byte, []int) { return fileDescriptorTypes, []int{1} }

func (m *CreateSpecType) GetPolicies() []*ves_io_schema4.ObjectRefType {
	if m != nil {
		return m.Policies
	}
	return nil
}

// Replace Alert Policy Set
//
// x-displayName: "Replace Alert Policy Set"
// Replaces the content of the Alert Policy Set Object
type ReplaceSpecType struct {
	Policies []*ves_io_schema4.ObjectRefType `protobuf:"bytes,1,rep,name=policies" json:"policies,omitempty"`
}

func (m *ReplaceSpecType) Reset()                    { *m = ReplaceSpecType{} }
func (*ReplaceSpecType) ProtoMessage()               {}
func (*ReplaceSpecType) Descriptor() ([]byte, []int) { return fileDescriptorTypes, []int{2} }

func (m *ReplaceSpecType) GetPolicies() []*ves_io_schema4.ObjectRefType {
	if m != nil {
		return m.Policies
	}
	return nil
}

// Get Alert Policy Set
//
// x-displayName: "Get Alert Policy Set"
// Get the Alert Policy Set Object
type GetSpecType struct {
	Policies []*ves_io_schema4.ObjectRefType `protobuf:"bytes,1,rep,name=policies" json:"policies,omitempty"`
}

func (m *GetSpecType) Reset()                    { *m = GetSpecType{} }
func (*GetSpecType) ProtoMessage()               {}
func (*GetSpecType) Descriptor() ([]byte, []int) { return fileDescriptorTypes, []int{3} }

func (m *GetSpecType) GetPolicies() []*ves_io_schema4.ObjectRefType {
	if m != nil {
		return m.Policies
	}
	return nil
}

func init() {
	proto.RegisterType((*GlobalSpecType)(nil), "ves.io.schema.alert_policy_set.GlobalSpecType")
	proto.RegisterType((*CreateSpecType)(nil), "ves.io.schema.alert_policy_set.CreateSpecType")
	proto.RegisterType((*ReplaceSpecType)(nil), "ves.io.schema.alert_policy_set.ReplaceSpecType")
	proto.RegisterType((*GetSpecType)(nil), "ves.io.schema.alert_policy_set.GetSpecType")
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
	if len(this.Policies) != len(that1.Policies) {
		return false
	}
	for i := range this.Policies {
		if !this.Policies[i].Equal(that1.Policies[i]) {
			return false
		}
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
	if len(this.Policies) != len(that1.Policies) {
		return false
	}
	for i := range this.Policies {
		if !this.Policies[i].Equal(that1.Policies[i]) {
			return false
		}
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
	if len(this.Policies) != len(that1.Policies) {
		return false
	}
	for i := range this.Policies {
		if !this.Policies[i].Equal(that1.Policies[i]) {
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
	if len(this.Policies) != len(that1.Policies) {
		return false
	}
	for i := range this.Policies {
		if !this.Policies[i].Equal(that1.Policies[i]) {
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
	s = append(s, "&alert_policy_set.GlobalSpecType{")
	if this.Policies != nil {
		s = append(s, "Policies: "+fmt.Sprintf("%#v", this.Policies)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *CreateSpecType) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&alert_policy_set.CreateSpecType{")
	if this.Policies != nil {
		s = append(s, "Policies: "+fmt.Sprintf("%#v", this.Policies)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *ReplaceSpecType) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&alert_policy_set.ReplaceSpecType{")
	if this.Policies != nil {
		s = append(s, "Policies: "+fmt.Sprintf("%#v", this.Policies)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *GetSpecType) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&alert_policy_set.GetSpecType{")
	if this.Policies != nil {
		s = append(s, "Policies: "+fmt.Sprintf("%#v", this.Policies)+",\n")
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
	if len(m.Policies) > 0 {
		for _, msg := range m.Policies {
			dAtA[i] = 0xa
			i++
			i = encodeVarintTypes(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
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
	if len(m.Policies) > 0 {
		for _, msg := range m.Policies {
			dAtA[i] = 0xa
			i++
			i = encodeVarintTypes(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
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
	if len(m.Policies) > 0 {
		for _, msg := range m.Policies {
			dAtA[i] = 0xa
			i++
			i = encodeVarintTypes(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
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
	if len(m.Policies) > 0 {
		for _, msg := range m.Policies {
			dAtA[i] = 0xa
			i++
			i = encodeVarintTypes(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
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
	if r.Intn(10) != 0 {
		v1 := r.Intn(5)
		this.Policies = make([]*ves_io_schema4.ObjectRefType, v1)
		for i := 0; i < v1; i++ {
			this.Policies[i] = ves_io_schema4.NewPopulatedObjectRefType(r, easy)
		}
	}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedCreateSpecType(r randyTypes, easy bool) *CreateSpecType {
	this := &CreateSpecType{}
	if r.Intn(10) != 0 {
		v2 := r.Intn(5)
		this.Policies = make([]*ves_io_schema4.ObjectRefType, v2)
		for i := 0; i < v2; i++ {
			this.Policies[i] = ves_io_schema4.NewPopulatedObjectRefType(r, easy)
		}
	}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedReplaceSpecType(r randyTypes, easy bool) *ReplaceSpecType {
	this := &ReplaceSpecType{}
	if r.Intn(10) != 0 {
		v3 := r.Intn(5)
		this.Policies = make([]*ves_io_schema4.ObjectRefType, v3)
		for i := 0; i < v3; i++ {
			this.Policies[i] = ves_io_schema4.NewPopulatedObjectRefType(r, easy)
		}
	}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedGetSpecType(r randyTypes, easy bool) *GetSpecType {
	this := &GetSpecType{}
	if r.Intn(10) != 0 {
		v4 := r.Intn(5)
		this.Policies = make([]*ves_io_schema4.ObjectRefType, v4)
		for i := 0; i < v4; i++ {
			this.Policies[i] = ves_io_schema4.NewPopulatedObjectRefType(r, easy)
		}
	}
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
	v5 := r.Intn(100)
	tmps := make([]rune, v5)
	for i := 0; i < v5; i++ {
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
		v6 := r.Int63()
		if r.Intn(2) == 0 {
			v6 *= -1
		}
		dAtA = encodeVarintPopulateTypes(dAtA, uint64(v6))
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
	if len(m.Policies) > 0 {
		for _, e := range m.Policies {
			l = e.Size()
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	return n
}

func (m *CreateSpecType) Size() (n int) {
	var l int
	_ = l
	if len(m.Policies) > 0 {
		for _, e := range m.Policies {
			l = e.Size()
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	return n
}

func (m *ReplaceSpecType) Size() (n int) {
	var l int
	_ = l
	if len(m.Policies) > 0 {
		for _, e := range m.Policies {
			l = e.Size()
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	return n
}

func (m *GetSpecType) Size() (n int) {
	var l int
	_ = l
	if len(m.Policies) > 0 {
		for _, e := range m.Policies {
			l = e.Size()
			n += 1 + l + sovTypes(uint64(l))
		}
	}
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
		`Policies:` + strings.Replace(fmt.Sprintf("%v", this.Policies), "ObjectRefType", "ves_io_schema4.ObjectRefType", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *CreateSpecType) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&CreateSpecType{`,
		`Policies:` + strings.Replace(fmt.Sprintf("%v", this.Policies), "ObjectRefType", "ves_io_schema4.ObjectRefType", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *ReplaceSpecType) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ReplaceSpecType{`,
		`Policies:` + strings.Replace(fmt.Sprintf("%v", this.Policies), "ObjectRefType", "ves_io_schema4.ObjectRefType", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *GetSpecType) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&GetSpecType{`,
		`Policies:` + strings.Replace(fmt.Sprintf("%v", this.Policies), "ObjectRefType", "ves_io_schema4.ObjectRefType", 1) + `,`,
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
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Policies", wireType)
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
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Policies = append(m.Policies, &ves_io_schema4.ObjectRefType{})
			if err := m.Policies[len(m.Policies)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Policies", wireType)
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
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Policies = append(m.Policies, &ves_io_schema4.ObjectRefType{})
			if err := m.Policies[len(m.Policies)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Policies", wireType)
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
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Policies = append(m.Policies, &ves_io_schema4.ObjectRefType{})
			if err := m.Policies[len(m.Policies)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Policies", wireType)
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
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Policies = append(m.Policies, &ves_io_schema4.ObjectRefType{})
			if err := m.Policies[len(m.Policies)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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

func init() { proto.RegisterFile("ves.io/schema/alert_policy_set/types.proto", fileDescriptorTypes) }

var fileDescriptorTypes = []byte{
	// 402 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x93, 0xb1, 0x6b, 0x14, 0x41,
	0x14, 0xc6, 0xef, 0x21, 0x88, 0x4c, 0x24, 0xc1, 0xab, 0xf4, 0x8c, 0x8f, 0xe5, 0xb0, 0x08, 0xc2,
	0xcd, 0xa0, 0x76, 0x16, 0x82, 0xb1, 0x08, 0x36, 0x2a, 0x67, 0x2a, 0x9b, 0x63, 0x76, 0xf3, 0x6e,
	0x33, 0x66, 0x37, 0x6f, 0x98, 0x99, 0x5b, 0xbd, 0x42, 0xb0, 0xb1, 0x17, 0xff, 0x0a, 0xf1, 0x2f,
	0x10, 0x6c, 0xc4, 0x4a, 0xac, 0x52, 0xa6, 0x74, 0xc7, 0xc6, 0x32, 0xa5, 0xa5, 0x64, 0x6f, 0x95,
	0xec, 0x81, 0x65, 0xae, 0x9b, 0xc7, 0xef, 0x7b, 0xef, 0x7b, 0x1f, 0xbc, 0x11, 0xb7, 0x2a, 0xf2,
	0xd2, 0xb0, 0xf2, 0xd9, 0x3e, 0x95, 0x5a, 0xe9, 0x82, 0x5c, 0x98, 0x58, 0x2e, 0x4c, 0x36, 0x9f,
	0x78, 0x0a, 0x2a, 0xcc, 0x2d, 0x79, 0x69, 0x1d, 0x07, 0xee, 0xe3, 0x42, 0x2b, 0x17, 0x5a, 0xb9,
	0xac, 0x1d, 0x8c, 0x72, 0x13, 0xf6, 0x67, 0xa9, 0xcc, 0xb8, 0x54, 0x39, 0xe7, 0xac, 0x9a, 0xb6,
	0x74, 0x36, 0x6d, 0xaa, 0xa6, 0x68, 0x5e, 0x8b, 0x71, 0x83, 0xeb, 0x5d, 0x6b, 0xb6, 0xc1, 0xf0,
	0x61, 0xeb, 0x35, 0xb8, 0xd6, 0x85, 0x67, 0xd6, 0x18, 0x6c, 0x76, 0x51, 0xa5, 0x0b, 0xb3, 0xa7,
	0x03, 0xb5, 0x34, 0x59, 0xa2, 0x86, 0x5e, 0x4e, 0x3a, 0xa3, 0x87, 0x6f, 0x41, 0xac, 0xef, 0x14,
	0x9c, 0xea, 0xe2, 0x99, 0xa5, 0x6c, 0x77, 0x6e, 0xa9, 0xff, 0x48, 0x5c, 0x6a, 0x72, 0x18, 0xf2,
	0x57, 0x21, 0xb9, 0xb0, 0xb5, 0x76, 0x67, 0x53, 0x76, 0xc3, 0x3e, 0x49, 0x5f, 0x50, 0x16, 0xc6,
	0x34, 0x3d, 0xd5, 0x6f, 0x6f, 0x7c, 0x7c, 0x7d, 0xf9, 0x6c, 0xfc, 0xf1, 0xbf, 0xf6, 0x7b, 0x5b,
	0x5f, 0x3f, 0xc3, 0x4d, 0x31, 0x14, 0x37, 0x1e, 0x9c, 0x72, 0x9f, 0x3c, 0x6d, 0x41, 0x32, 0x65,
	0x97, 0x3c, 0xd6, 0x25, 0x79, 0xab, 0x33, 0xea, 0xc3, 0xed, 0xe1, 0xa1, 0x58, 0x7f, 0xe8, 0x48,
	0x07, 0x3a, 0x8f, 0x35, 0xae, 0x7c, 0xbf, 0xbf, 0x14, 0x72, 0xc8, 0x62, 0x63, 0x4c, 0xb6, 0xd0,
	0xd9, 0xaa, 0x0c, 0x0f, 0xc4, 0xda, 0x0e, 0x85, 0xd5, 0x98, 0x6d, 0xbf, 0x87, 0xa3, 0x1a, 0x7b,
	0xc7, 0x35, 0xf6, 0x4e, 0x6a, 0x84, 0xdf, 0x35, 0xc2, 0x9b, 0x88, 0xf0, 0x21, 0x22, 0x7c, 0x8a,
	0x08, 0x5f, 0x22, 0xc2, 0xb7, 0x88, 0x70, 0x14, 0x11, 0x8e, 0x23, 0xc2, 0x8f, 0x88, 0xf0, 0x2b,
	0x62, 0xef, 0x24, 0x22, 0xbc, 0xfb, 0x89, 0xbd, 0xe7, 0xbb, 0x39, 0xdb, 0x83, 0x5c, 0x56, 0x5c,
	0x04, 0x72, 0x4e, 0xcb, 0x99, 0x57, 0xcd, 0x63, 0xca, 0xae, 0x1c, 0x59, 0xc7, 0x95, 0xd9, 0x23,
	0x37, 0xfa, 0x8b, 0x95, 0x4d, 0x73, 0x56, 0xf4, 0x2a, 0xb4, 0x57, 0xf6, 0x9f, 0xdf, 0x93, 0x5e,
	0x6c, 0x2e, 0xee, 0xee, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x5c, 0xf4, 0xbb, 0x54, 0x66, 0x03,
	0x00, 0x00,
}
