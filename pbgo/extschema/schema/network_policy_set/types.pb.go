// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ves.io/schema/network_policy_set/types.proto

package network_policy_set

import (
	proto "github.com/gogo/protobuf/proto"
	golang_proto "github.com/golang/protobuf/proto"

	fmt "fmt"

	math "math"

	_ "github.com/gogo/protobuf/gogoproto"

	_ "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema"

	_ "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema"

	ves_io_schema4 "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema"

	_ "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema"

	_ "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema"

	strconv "strconv"

	strings "strings"

	reflect "reflect"

	io "io"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = golang_proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Network Policy Set Type
//
// x-displayName: "Network Policy Set Type"
// Internal field to dis how this policy set is used.
type NetworkPolicySetType int32

const (
	// x-displayName: "Network Firewall"
	// Used for network firewall
	NETWORK_FIREWALL NetworkPolicySetType = 0
	// x-displayName: "Policy Based Routing"
	// Used for network policy based routing
	POLICY_BASED_ROUTING NetworkPolicySetType = 1
)

var NetworkPolicySetType_name = map[int32]string{
	0: "NETWORK_FIREWALL",
	1: "POLICY_BASED_ROUTING",
}
var NetworkPolicySetType_value = map[string]int32{
	"NETWORK_FIREWALL":     0,
	"POLICY_BASED_ROUTING": 1,
}

func (NetworkPolicySetType) EnumDescriptor() ([]byte, []int) { return fileDescriptorTypes, []int{0} }

// Network policy set specification
//
// x-displayName: "Specification"
// Desired state of network policy set
type GlobalSpecType struct {
	// policies
	//
	// x-displayName: "Policies"
	// Ordered list of references to the network policy that make up this Network policy set.
	Policies []*ves_io_schema4.ObjectRefType `protobuf:"bytes,1,rep,name=policies" json:"policies,omitempty"`
	// Network Policy Set Type
	//
	// x-displayName: "Network Policy Set Type"
	// Network policy set type lets VER know what this policy set is used for.
	NpsType NetworkPolicySetType `protobuf:"varint,2,opt,name=nps_type,json=npsType,proto3,enum=ves.io.schema.network_policy_set.NetworkPolicySetType" json:"nps_type,omitempty"`
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

func (m *GlobalSpecType) GetNpsType() NetworkPolicySetType {
	if m != nil {
		return m.NpsType
	}
	return NETWORK_FIREWALL
}

// Create Network Policy Set
//
// x-displayName: "Create Network Policy Set"
// Create Network Policy Set in a given namespace. If one already exist it will give a error.
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

// Replace Network Policy Set
//
// x-displayName: "Replace Network Policy Set"
// Replace Network Policy Set in a given namespace.
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

// Get Network Policy Set
//
// x-displayName: "Get Network Policy Set"
// Get network policy set in a given namespace.
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
	proto.RegisterType((*GlobalSpecType)(nil), "ves.io.schema.network_policy_set.GlobalSpecType")
	golang_proto.RegisterType((*GlobalSpecType)(nil), "ves.io.schema.network_policy_set.GlobalSpecType")
	proto.RegisterType((*CreateSpecType)(nil), "ves.io.schema.network_policy_set.CreateSpecType")
	golang_proto.RegisterType((*CreateSpecType)(nil), "ves.io.schema.network_policy_set.CreateSpecType")
	proto.RegisterType((*ReplaceSpecType)(nil), "ves.io.schema.network_policy_set.ReplaceSpecType")
	golang_proto.RegisterType((*ReplaceSpecType)(nil), "ves.io.schema.network_policy_set.ReplaceSpecType")
	proto.RegisterType((*GetSpecType)(nil), "ves.io.schema.network_policy_set.GetSpecType")
	golang_proto.RegisterType((*GetSpecType)(nil), "ves.io.schema.network_policy_set.GetSpecType")
	proto.RegisterEnum("ves.io.schema.network_policy_set.NetworkPolicySetType", NetworkPolicySetType_name, NetworkPolicySetType_value)
	golang_proto.RegisterEnum("ves.io.schema.network_policy_set.NetworkPolicySetType", NetworkPolicySetType_name, NetworkPolicySetType_value)
}
func (x NetworkPolicySetType) String() string {
	s, ok := NetworkPolicySetType_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
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
	if this.NpsType != that1.NpsType {
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
	s := make([]string, 0, 6)
	s = append(s, "&network_policy_set.GlobalSpecType{")
	if this.Policies != nil {
		s = append(s, "Policies: "+fmt.Sprintf("%#v", this.Policies)+",\n")
	}
	s = append(s, "NpsType: "+fmt.Sprintf("%#v", this.NpsType)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *CreateSpecType) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&network_policy_set.CreateSpecType{")
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
	s = append(s, "&network_policy_set.ReplaceSpecType{")
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
	s = append(s, "&network_policy_set.GetSpecType{")
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
	if m.NpsType != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintTypes(dAtA, i, uint64(m.NpsType))
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
	this.NpsType = NetworkPolicySetType([]int32{0, 1}[r.Intn(2)])
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
	if m.NpsType != 0 {
		n += 1 + sovTypes(uint64(m.NpsType))
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
		`NpsType:` + fmt.Sprintf("%v", this.NpsType) + `,`,
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
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NpsType", wireType)
			}
			m.NpsType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NpsType |= (NetworkPolicySetType(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
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

func init() { proto.RegisterFile("ves.io/schema/network_policy_set/types.proto", fileDescriptorTypes) }
func init() {
	golang_proto.RegisterFile("ves.io/schema/network_policy_set/types.proto", fileDescriptorTypes)
}

var fileDescriptorTypes = []byte{
	// 506 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x53, 0x4f, 0x6b, 0x13, 0x41,
	0x14, 0xcf, 0x33, 0xd0, 0x96, 0x29, 0xc4, 0x34, 0x04, 0x8c, 0x31, 0x0c, 0x21, 0xa7, 0x22, 0x66,
	0x16, 0x2a, 0x78, 0xf0, 0x20, 0x34, 0x35, 0x0d, 0xa1, 0x21, 0x29, 0x9b, 0xd8, 0x62, 0x2f, 0xcb,
	0xee, 0xf6, 0x65, 0xbb, 0x76, 0x93, 0x19, 0x77, 0x27, 0xa9, 0x39, 0x14, 0xfa, 0x11, 0xc4, 0x93,
	0x1f, 0x41, 0xfc, 0x04, 0xa2, 0x97, 0x1c, 0x3c, 0x88, 0xa7, 0x1e, 0x73, 0x34, 0xd3, 0x4b, 0xbd,
	0xf5, 0xe8, 0x51, 0xb2, 0x89, 0xd5, 0x6d, 0x14, 0xe9, 0xa5, 0xb7, 0x99, 0xf7, 0xfb, 0xed, 0xfb,
	0xfd, 0x61, 0x87, 0x3c, 0xe8, 0x63, 0xc0, 0x5c, 0xae, 0x05, 0xf6, 0x01, 0x76, 0x4c, 0xad, 0x8b,
	0xf2, 0x88, 0xfb, 0x87, 0x86, 0xe0, 0x9e, 0x6b, 0x0f, 0x8c, 0x00, 0xa5, 0x26, 0x07, 0x02, 0x03,
	0x26, 0x7c, 0x2e, 0x79, 0x2a, 0x3f, 0x65, 0xb3, 0x29, 0x9b, 0xcd, 0xb3, 0xb3, 0x45, 0xc7, 0x95,
	0x07, 0x3d, 0x8b, 0xd9, 0xbc, 0xa3, 0x39, 0xdc, 0xe1, 0x5a, 0xf8, 0xa1, 0xd5, 0x6b, 0x87, 0xb7,
	0xf0, 0x12, 0x9e, 0xa6, 0x0b, 0xb3, 0x77, 0xe6, 0xe4, 0x67, 0xc0, 0xbd, 0x28, 0xc0, 0x85, 0x74,
	0x79, 0x77, 0x66, 0x23, 0x7b, 0x37, 0x0a, 0xfe, 0xe1, 0x30, 0x9b, 0x8b, 0x42, 0x7d, 0xd3, 0x73,
	0xf7, 0x4d, 0x89, 0x33, 0x34, 0x7f, 0x05, 0x75, 0xf1, 0xc8, 0x88, 0xac, 0x2e, 0x7c, 0x06, 0x92,
	0xa8, 0x78, 0xdc, 0x32, 0xbd, 0xa6, 0x40, 0xbb, 0x35, 0x10, 0x98, 0xda, 0x21, 0x4b, 0x61, 0x40,
	0x17, 0x83, 0x0c, 0xe4, 0xe3, 0xab, 0xcb, 0x6b, 0x39, 0x16, 0xed, 0xa1, 0x61, 0xbd, 0x40, 0x5b,
	0xea, 0xd8, 0x9e, 0xf0, 0x4b, 0xb9, 0xf7, 0xc7, 0x89, 0x68, 0x33, 0x1f, 0xbf, 0x0f, 0xe3, 0x0b,
	0x6f, 0x20, 0x9e, 0x3c, 0x01, 0xfd, 0x72, 0x57, 0x6a, 0x8f, 0x2c, 0x75, 0x45, 0x60, 0x4c, 0xdc,
	0x67, 0x6e, 0xe5, 0x61, 0x35, 0xb1, 0xf6, 0x88, 0xfd, 0xaf, 0x5f, 0x56, 0x9f, 0x8e, 0xb6, 0xc3,
	0x49, 0x13, 0x65, 0xa8, 0xb8, 0x38, 0x3a, 0x86, 0xf3, 0x4f, 0x00, 0xfa, 0x62, 0x57, 0x04, 0x93,
	0x49, 0x41, 0x90, 0xc4, 0x86, 0x8f, 0xa6, 0xc4, 0xcb, 0x14, 0x5b, 0xd7, 0x4c, 0xb1, 0x32, 0x97,
	0xe2, 0xb7, 0xf5, 0xc7, 0x2b, 0x5f, 0x9f, 0x5c, 0x69, 0xa9, 0xf0, 0x92, 0xdc, 0xd6, 0x51, 0x78,
	0xa6, 0x7d, 0x73, 0x92, 0x1d, 0xb2, 0x5c, 0x41, 0x79, 0x53, 0x72, 0xf7, 0x37, 0x49, 0xfa, 0x6f,
	0xed, 0xa7, 0xd2, 0x24, 0x59, 0x2f, 0xb7, 0x76, 0x1b, 0xfa, 0x96, 0xb1, 0x59, 0xd5, 0xcb, 0xbb,
	0xeb, 0xb5, 0x5a, 0x32, 0x96, 0xca, 0x90, 0xf4, 0x76, 0xa3, 0x56, 0xdd, 0x78, 0x6e, 0x94, 0xd6,
	0x9b, 0xe5, 0xa7, 0x86, 0xde, 0x78, 0xd6, 0xaa, 0xd6, 0x2b, 0x49, 0x28, 0xbd, 0x85, 0xd3, 0x31,
	0x8d, 0x8d, 0xc6, 0x34, 0x76, 0x31, 0xa6, 0xf0, 0x63, 0x4c, 0xe1, 0x44, 0x51, 0x78, 0xa7, 0x28,
	0x7c, 0x50, 0x14, 0x86, 0x8a, 0xc2, 0x17, 0x45, 0xe1, 0x54, 0x51, 0x18, 0x29, 0x0a, 0xdf, 0x14,
	0x85, 0x73, 0x45, 0x63, 0x17, 0x8a, 0xc2, 0xeb, 0x33, 0x1a, 0x1b, 0x9e, 0x51, 0xd8, 0xdb, 0x71,
	0xb8, 0x38, 0x74, 0x58, 0x9f, 0x7b, 0x12, 0x7d, 0xdf, 0x64, 0xbd, 0x40, 0x0b, 0x0f, 0x6d, 0xee,
	0x77, 0x8a, 0xc2, 0xe7, 0x7d, 0x77, 0x1f, 0xfd, 0xe2, 0x2f, 0x58, 0x13, 0x96, 0xc3, 0x35, 0x7c,
	0x25, 0x67, 0xbf, 0xfd, 0x3f, 0xdf, 0xba, 0xb5, 0x10, 0x3e, 0x82, 0x87, 0x3f, 0x03, 0x00, 0x00,
	0xff, 0xff, 0x39, 0xc8, 0x9a, 0xf9, 0x16, 0x04, 0x00, 0x00,
}