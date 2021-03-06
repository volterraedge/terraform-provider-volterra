// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ves.io/schema/views/storage_types.proto

package views

import proto "github.com/gogo/protobuf/proto"
import golang_proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import _ "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema"
import _ "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema"
import _ "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema"
import _ "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema"

import strings "strings"
import reflect "reflect"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = golang_proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Custom Storage Class List
//
// x-displayName: "Custom Storage Class List"
// Add additional custom storage classes in kubernetes for this site
type StorageClassListType struct {
	// List of Storage Classes
	//
	// x-displayName: "List of Storage Classes"
	// List of custom storage classes
	StorageClasses []*StorageClassType `protobuf:"bytes,1,rep,name=storage_classes,json=storageClasses" json:"storage_classes,omitempty"`
}

func (m *StorageClassListType) Reset()      { *m = StorageClassListType{} }
func (*StorageClassListType) ProtoMessage() {}
func (*StorageClassListType) Descriptor() ([]byte, []int) {
	return fileDescriptorStorageTypes, []int{0}
}

func (m *StorageClassListType) GetStorageClasses() []*StorageClassType {
	if m != nil {
		return m.StorageClasses
	}
	return nil
}

// Custom Storage Class
//
// x-displayName: "Custom Storage Class"
// Configuration of custom storage class
type StorageClassType struct {
	// Storage Class Name
	//
	// x-displayName: "Storage Class Name"
	// x-example: "premium"
	// x-required
	// Name of the storage class as it will appear in K8s.
	StorageClassName string `protobuf:"bytes,1,opt,name=storage_class_name,json=storageClassName,proto3" json:"storage_class_name,omitempty"`
	// Default Storage Class
	//
	// x-displayName: "Default Storage Class"
	// Make this storage class default storage class for the K8s cluster
	DefaultStorageClass bool `protobuf:"varint,2,opt,name=default_storage_class,json=defaultStorageClass,proto3" json:"default_storage_class,omitempty"`
	// Select Storage Class Configuration
	//
	// x-displayName: "Select Storage Class Configuration"
	// x-required
	// Select storage Class configuration
	//
	// Types that are valid to be assigned to DeviceChoice:
	//	*StorageClassType_OpenebsEnterprise
	DeviceChoice isStorageClassType_DeviceChoice `protobuf_oneof:"device_choice"`
}

func (m *StorageClassType) Reset()                    { *m = StorageClassType{} }
func (*StorageClassType) ProtoMessage()               {}
func (*StorageClassType) Descriptor() ([]byte, []int) { return fileDescriptorStorageTypes, []int{1} }

type isStorageClassType_DeviceChoice interface {
	isStorageClassType_DeviceChoice()
	Equal(interface{}) bool
	MarshalTo([]byte) (int, error)
	Size() int
}

type StorageClassType_OpenebsEnterprise struct {
	OpenebsEnterprise *StorageClassOpenebsEnterpriseType `protobuf:"bytes,6,opt,name=openebs_enterprise,json=openebsEnterprise,oneof"`
}

func (*StorageClassType_OpenebsEnterprise) isStorageClassType_DeviceChoice() {}

func (m *StorageClassType) GetDeviceChoice() isStorageClassType_DeviceChoice {
	if m != nil {
		return m.DeviceChoice
	}
	return nil
}

func (m *StorageClassType) GetStorageClassName() string {
	if m != nil {
		return m.StorageClassName
	}
	return ""
}

func (m *StorageClassType) GetDefaultStorageClass() bool {
	if m != nil {
		return m.DefaultStorageClass
	}
	return false
}

func (m *StorageClassType) GetOpenebsEnterprise() *StorageClassOpenebsEnterpriseType {
	if x, ok := m.GetDeviceChoice().(*StorageClassType_OpenebsEnterprise); ok {
		return x.OpenebsEnterprise
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*StorageClassType) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _StorageClassType_OneofMarshaler, _StorageClassType_OneofUnmarshaler, _StorageClassType_OneofSizer, []interface{}{
		(*StorageClassType_OpenebsEnterprise)(nil),
	}
}

func _StorageClassType_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*StorageClassType)
	// device_choice
	switch x := m.DeviceChoice.(type) {
	case *StorageClassType_OpenebsEnterprise:
		_ = b.EncodeVarint(6<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.OpenebsEnterprise); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("StorageClassType.DeviceChoice has unexpected type %T", x)
	}
	return nil
}

func _StorageClassType_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*StorageClassType)
	switch tag {
	case 6: // device_choice.openebs_enterprise
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(StorageClassOpenebsEnterpriseType)
		err := b.DecodeMessage(msg)
		m.DeviceChoice = &StorageClassType_OpenebsEnterprise{msg}
		return true, err
	default:
		return false, nil
	}
}

func _StorageClassType_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*StorageClassType)
	// device_choice
	switch x := m.DeviceChoice.(type) {
	case *StorageClassType_OpenebsEnterprise:
		s := proto.Size(x.OpenebsEnterprise)
		n += proto.SizeVarint(6<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// OpenEBS Enterprise
//
// x-displayName: "OpenEBS Enterprise"
// Storage class Device configuration for OpenEBS Enterprise
type StorageClassOpenebsEnterpriseType struct {
	// Replication
	//
	// x-displayName: "Replication"
	// x-example: "1"
	// Replication sets the replication factor of the PV, i.e. the number of data replicas to be maintained for it such as 1 or 3.
	Replication int32 `protobuf:"varint,2,opt,name=replication,proto3" json:"replication,omitempty"`
	// Storage Size
	//
	// x-displayName: "Storage Size"
	// x-example "10"
	// Size of each node of storage class. e.g If "Storage Class Replicas" will be set to 3 and "Storage Class Size" to 10GB.
	// Three 10GB disk will be created and assigned to nodes.
	StorageClassSize uint32 `protobuf:"varint,4,opt,name=storage_class_size,json=storageClassSize,proto3" json:"storage_class_size,omitempty"`
}

func (m *StorageClassOpenebsEnterpriseType) Reset()      { *m = StorageClassOpenebsEnterpriseType{} }
func (*StorageClassOpenebsEnterpriseType) ProtoMessage() {}
func (*StorageClassOpenebsEnterpriseType) Descriptor() ([]byte, []int) {
	return fileDescriptorStorageTypes, []int{2}
}

func (m *StorageClassOpenebsEnterpriseType) GetReplication() int32 {
	if m != nil {
		return m.Replication
	}
	return 0
}

func (m *StorageClassOpenebsEnterpriseType) GetStorageClassSize() uint32 {
	if m != nil {
		return m.StorageClassSize
	}
	return 0
}

func init() {
	proto.RegisterType((*StorageClassListType)(nil), "ves.io.schema.views.StorageClassListType")
	golang_proto.RegisterType((*StorageClassListType)(nil), "ves.io.schema.views.StorageClassListType")
	proto.RegisterType((*StorageClassType)(nil), "ves.io.schema.views.StorageClassType")
	golang_proto.RegisterType((*StorageClassType)(nil), "ves.io.schema.views.StorageClassType")
	proto.RegisterType((*StorageClassOpenebsEnterpriseType)(nil), "ves.io.schema.views.StorageClassOpenebsEnterpriseType")
	golang_proto.RegisterType((*StorageClassOpenebsEnterpriseType)(nil), "ves.io.schema.views.StorageClassOpenebsEnterpriseType")
}
func (this *StorageClassListType) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*StorageClassListType)
	if !ok {
		that2, ok := that.(StorageClassListType)
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
	if len(this.StorageClasses) != len(that1.StorageClasses) {
		return false
	}
	for i := range this.StorageClasses {
		if !this.StorageClasses[i].Equal(that1.StorageClasses[i]) {
			return false
		}
	}
	return true
}
func (this *StorageClassType) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*StorageClassType)
	if !ok {
		that2, ok := that.(StorageClassType)
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
	if this.StorageClassName != that1.StorageClassName {
		return false
	}
	if this.DefaultStorageClass != that1.DefaultStorageClass {
		return false
	}
	if that1.DeviceChoice == nil {
		if this.DeviceChoice != nil {
			return false
		}
	} else if this.DeviceChoice == nil {
		return false
	} else if !this.DeviceChoice.Equal(that1.DeviceChoice) {
		return false
	}
	return true
}
func (this *StorageClassType_OpenebsEnterprise) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*StorageClassType_OpenebsEnterprise)
	if !ok {
		that2, ok := that.(StorageClassType_OpenebsEnterprise)
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
	if !this.OpenebsEnterprise.Equal(that1.OpenebsEnterprise) {
		return false
	}
	return true
}
func (this *StorageClassOpenebsEnterpriseType) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*StorageClassOpenebsEnterpriseType)
	if !ok {
		that2, ok := that.(StorageClassOpenebsEnterpriseType)
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
	if this.Replication != that1.Replication {
		return false
	}
	if this.StorageClassSize != that1.StorageClassSize {
		return false
	}
	return true
}
func (this *StorageClassListType) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&views.StorageClassListType{")
	if this.StorageClasses != nil {
		s = append(s, "StorageClasses: "+fmt.Sprintf("%#v", this.StorageClasses)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *StorageClassType) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 7)
	s = append(s, "&views.StorageClassType{")
	s = append(s, "StorageClassName: "+fmt.Sprintf("%#v", this.StorageClassName)+",\n")
	s = append(s, "DefaultStorageClass: "+fmt.Sprintf("%#v", this.DefaultStorageClass)+",\n")
	if this.DeviceChoice != nil {
		s = append(s, "DeviceChoice: "+fmt.Sprintf("%#v", this.DeviceChoice)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *StorageClassType_OpenebsEnterprise) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&views.StorageClassType_OpenebsEnterprise{` +
		`OpenebsEnterprise:` + fmt.Sprintf("%#v", this.OpenebsEnterprise) + `}`}, ", ")
	return s
}
func (this *StorageClassOpenebsEnterpriseType) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&views.StorageClassOpenebsEnterpriseType{")
	s = append(s, "Replication: "+fmt.Sprintf("%#v", this.Replication)+",\n")
	s = append(s, "StorageClassSize: "+fmt.Sprintf("%#v", this.StorageClassSize)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringStorageTypes(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *StorageClassListType) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StorageClassListType) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.StorageClasses) > 0 {
		for _, msg := range m.StorageClasses {
			dAtA[i] = 0xa
			i++
			i = encodeVarintStorageTypes(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *StorageClassType) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StorageClassType) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.StorageClassName) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintStorageTypes(dAtA, i, uint64(len(m.StorageClassName)))
		i += copy(dAtA[i:], m.StorageClassName)
	}
	if m.DefaultStorageClass {
		dAtA[i] = 0x10
		i++
		if m.DefaultStorageClass {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.DeviceChoice != nil {
		nn1, err := m.DeviceChoice.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += nn1
	}
	return i, nil
}

func (m *StorageClassType_OpenebsEnterprise) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	if m.OpenebsEnterprise != nil {
		dAtA[i] = 0x32
		i++
		i = encodeVarintStorageTypes(dAtA, i, uint64(m.OpenebsEnterprise.Size()))
		n2, err := m.OpenebsEnterprise.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	return i, nil
}
func (m *StorageClassOpenebsEnterpriseType) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StorageClassOpenebsEnterpriseType) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Replication != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintStorageTypes(dAtA, i, uint64(m.Replication))
	}
	if m.StorageClassSize != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintStorageTypes(dAtA, i, uint64(m.StorageClassSize))
	}
	return i, nil
}

func encodeVarintStorageTypes(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *StorageClassListType) Size() (n int) {
	var l int
	_ = l
	if len(m.StorageClasses) > 0 {
		for _, e := range m.StorageClasses {
			l = e.Size()
			n += 1 + l + sovStorageTypes(uint64(l))
		}
	}
	return n
}

func (m *StorageClassType) Size() (n int) {
	var l int
	_ = l
	l = len(m.StorageClassName)
	if l > 0 {
		n += 1 + l + sovStorageTypes(uint64(l))
	}
	if m.DefaultStorageClass {
		n += 2
	}
	if m.DeviceChoice != nil {
		n += m.DeviceChoice.Size()
	}
	return n
}

func (m *StorageClassType_OpenebsEnterprise) Size() (n int) {
	var l int
	_ = l
	if m.OpenebsEnterprise != nil {
		l = m.OpenebsEnterprise.Size()
		n += 1 + l + sovStorageTypes(uint64(l))
	}
	return n
}
func (m *StorageClassOpenebsEnterpriseType) Size() (n int) {
	var l int
	_ = l
	if m.Replication != 0 {
		n += 1 + sovStorageTypes(uint64(m.Replication))
	}
	if m.StorageClassSize != 0 {
		n += 1 + sovStorageTypes(uint64(m.StorageClassSize))
	}
	return n
}

func sovStorageTypes(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozStorageTypes(x uint64) (n int) {
	return sovStorageTypes(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *StorageClassListType) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&StorageClassListType{`,
		`StorageClasses:` + strings.Replace(fmt.Sprintf("%v", this.StorageClasses), "StorageClassType", "StorageClassType", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *StorageClassType) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&StorageClassType{`,
		`StorageClassName:` + fmt.Sprintf("%v", this.StorageClassName) + `,`,
		`DefaultStorageClass:` + fmt.Sprintf("%v", this.DefaultStorageClass) + `,`,
		`DeviceChoice:` + fmt.Sprintf("%v", this.DeviceChoice) + `,`,
		`}`,
	}, "")
	return s
}
func (this *StorageClassType_OpenebsEnterprise) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&StorageClassType_OpenebsEnterprise{`,
		`OpenebsEnterprise:` + strings.Replace(fmt.Sprintf("%v", this.OpenebsEnterprise), "StorageClassOpenebsEnterpriseType", "StorageClassOpenebsEnterpriseType", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *StorageClassOpenebsEnterpriseType) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&StorageClassOpenebsEnterpriseType{`,
		`Replication:` + fmt.Sprintf("%v", this.Replication) + `,`,
		`StorageClassSize:` + fmt.Sprintf("%v", this.StorageClassSize) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringStorageTypes(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *StorageClassListType) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStorageTypes
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
			return fmt.Errorf("proto: StorageClassListType: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StorageClassListType: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StorageClasses", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStorageTypes
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
				return ErrInvalidLengthStorageTypes
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StorageClasses = append(m.StorageClasses, &StorageClassType{})
			if err := m.StorageClasses[len(m.StorageClasses)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipStorageTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthStorageTypes
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
func (m *StorageClassType) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStorageTypes
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
			return fmt.Errorf("proto: StorageClassType: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StorageClassType: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StorageClassName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStorageTypes
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
				return ErrInvalidLengthStorageTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StorageClassName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DefaultStorageClass", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStorageTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.DefaultStorageClass = bool(v != 0)
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OpenebsEnterprise", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStorageTypes
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
				return ErrInvalidLengthStorageTypes
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &StorageClassOpenebsEnterpriseType{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.DeviceChoice = &StorageClassType_OpenebsEnterprise{v}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipStorageTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthStorageTypes
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
func (m *StorageClassOpenebsEnterpriseType) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStorageTypes
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
			return fmt.Errorf("proto: StorageClassOpenebsEnterpriseType: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StorageClassOpenebsEnterpriseType: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Replication", wireType)
			}
			m.Replication = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStorageTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Replication |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StorageClassSize", wireType)
			}
			m.StorageClassSize = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStorageTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StorageClassSize |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipStorageTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthStorageTypes
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
func skipStorageTypes(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowStorageTypes
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
					return 0, ErrIntOverflowStorageTypes
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
					return 0, ErrIntOverflowStorageTypes
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
				return 0, ErrInvalidLengthStorageTypes
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowStorageTypes
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
				next, err := skipStorageTypes(dAtA[start:])
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
	ErrInvalidLengthStorageTypes = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowStorageTypes   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("ves.io/schema/views/storage_types.proto", fileDescriptorStorageTypes)
}
func init() {
	golang_proto.RegisterFile("ves.io/schema/views/storage_types.proto", fileDescriptorStorageTypes)
}

var fileDescriptorStorageTypes = []byte{
	// 571 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x53, 0x41, 0x6b, 0x13, 0x41,
	0x14, 0xce, 0x4b, 0xd2, 0x5a, 0xa7, 0xb4, 0xc6, 0xa9, 0xc2, 0x1a, 0x65, 0x5c, 0x0b, 0x6a, 0x10,
	0xbb, 0xab, 0x29, 0x14, 0xe9, 0x49, 0xa2, 0x82, 0x8a, 0xa8, 0xa4, 0x9e, 0xbc, 0x2c, 0x93, 0xcd,
	0xeb, 0x66, 0x30, 0x9b, 0x59, 0x66, 0x26, 0xd1, 0xf6, 0x62, 0xff, 0x80, 0x20, 0xfe, 0x02, 0x8f,
	0xfe, 0x06, 0xf7, 0x52, 0x3c, 0x09, 0x5e, 0x72, 0x2c, 0x9e, 0xec, 0xf6, 0xa2, 0xb7, 0x1e, 0x3d,
	0xca, 0x6e, 0x12, 0xd9, 0x44, 0xa1, 0xb7, 0xf7, 0xf2, 0x7d, 0xef, 0xfb, 0xe6, 0xcb, 0x7b, 0x4b,
	0xae, 0x0f, 0x50, 0x3b, 0x42, 0xba, 0xda, 0xef, 0x60, 0xc8, 0xdd, 0x81, 0xc0, 0xd7, 0xda, 0xd5,
	0x46, 0x2a, 0x1e, 0xa0, 0x67, 0x76, 0x22, 0xd4, 0x4e, 0xa4, 0xa4, 0x91, 0x74, 0x65, 0x44, 0x74,
	0x46, 0x44, 0x27, 0x23, 0x56, 0xd7, 0x02, 0x61, 0x3a, 0xfd, 0x96, 0xe3, 0xcb, 0xd0, 0x0d, 0x64,
	0x20, 0xdd, 0x8c, 0xdb, 0xea, 0x6f, 0x67, 0x5d, 0xd6, 0x64, 0xd5, 0x48, 0xa3, 0x7a, 0x71, 0xda,
	0x4c, 0x46, 0x46, 0xc8, 0xde, 0xd8, 0xa0, 0x7a, 0x61, 0x1a, 0xcc, 0x79, 0x57, 0x2f, 0xcd, 0x3c,
	0x92, 0x77, 0x45, 0x9b, 0x1b, 0x1c, 0xa3, 0xf6, 0xbf, 0x11, 0xbc, 0x29, 0xe9, 0xd5, 0x3d, 0x20,
	0xe7, 0xb6, 0x46, 0x99, 0xee, 0x75, 0xb9, 0xd6, 0x4f, 0x84, 0x36, 0x2f, 0x76, 0x22, 0xa4, 0x1d,
	0x72, 0x66, 0x92, 0xd5, 0x4f, 0x01, 0xd4, 0x16, 0xd8, 0xa5, 0xda, 0x62, 0xfd, 0xaa, 0xf3, 0x9f,
	0xb8, 0x4e, 0x5e, 0x23, 0x9d, 0x6f, 0x58, 0xc3, 0x18, 0xe0, 0xf3, 0xaf, 0xfd, 0xd2, 0xdc, 0x07,
	0x28, 0x56, 0xca, 0x93, 0xca, 0x82, 0xe6, 0xb2, 0xce, 0x71, 0x51, 0xaf, 0x7e, 0x2b, 0x92, 0xca,
	0xec, 0x38, 0xbd, 0x43, 0xe8, 0x94, 0xbd, 0xd7, 0xe3, 0x21, 0x5a, 0x60, 0x43, 0xed, 0x74, 0x83,
	0x64, 0x62, 0xaa, 0xf4, 0xb1, 0x08, 0xcd, 0x4a, 0x5e, 0xec, 0x29, 0x0f, 0x91, 0xd6, 0xc9, 0xf9,
	0x36, 0x6e, 0xf3, 0x7e, 0xd7, 0x78, 0x53, 0x0a, 0x56, 0xd1, 0x86, 0xda, 0x42, 0x73, 0x65, 0x0c,
	0xe6, 0x1d, 0x69, 0x40, 0xa8, 0x8c, 0xb0, 0x87, 0x2d, 0xed, 0x61, 0xcf, 0xa0, 0x8a, 0x94, 0xd0,
	0x68, 0xcd, 0xdb, 0x50, 0x5b, 0xac, 0x6f, 0x9c, 0x98, 0xf7, 0xd9, 0x68, 0xf4, 0xc1, 0xdf, 0xc9,
	0x34, 0xc1, 0xc3, 0x42, 0xf3, 0xac, 0x9c, 0x05, 0x36, 0x37, 0xbf, 0xc4, 0xb0, 0x41, 0x2e, 0x13,
	0x6b, 0x3c, 0x6f, 0x67, 0x02, 0xf6, 0x73, 0xae, 0x78, 0x88, 0x06, 0x95, 0xa6, 0xa5, 0xdb, 0x37,
	0xeb, 0x64, 0x85, 0x2c, 0x4f, 0x08, 0xf7, 0x71, 0x20, 0x7c, 0xa4, 0xb0, 0xde, 0x60, 0x64, 0xa9,
	0x9d, 0x35, 0x9e, 0xdf, 0x91, 0xe9, 0x6f, 0x4b, 0xfb, 0x31, 0xcc, 0x0f, 0x63, 0x28, 0x25, 0x31,
	0xc0, 0xc6, 0xe3, 0xf2, 0x42, 0xa9, 0x52, 0x5e, 0x7d, 0x07, 0xe4, 0xca, 0x89, 0x8f, 0xa3, 0xd7,
	0xc8, 0xa2, 0xc2, 0xa8, 0x2b, 0x7c, 0x9e, 0x1e, 0x43, 0xf6, 0xd7, 0xcc, 0x35, 0xca, 0xbf, 0x63,
	0x28, 0x34, 0xf3, 0x00, 0xbd, 0x3b, 0xbb, 0x06, 0x2d, 0x76, 0xd1, 0x2a, 0xdb, 0x50, 0x5b, 0x6a,
	0xd0, 0xef, 0x31, 0x14, 0xeb, 0xb7, 0xd2, 0x65, 0x9c, 0xba, 0x31, 0x67, 0xed, 0x2d, 0xd4, 0x66,
	0xd6, 0xb1, 0x25, 0x76, 0xb1, 0xf1, 0x76, 0x78, 0xc8, 0x0a, 0x07, 0x87, 0xac, 0x70, 0x7c, 0xc8,
	0x60, 0x2f, 0x61, 0xf0, 0x29, 0x61, 0xf0, 0x35, 0x61, 0x30, 0x4c, 0x18, 0x1c, 0x24, 0x0c, 0x7e,
	0x24, 0x0c, 0x7e, 0x26, 0xac, 0x70, 0x9c, 0x30, 0x78, 0x7f, 0xc4, 0x0a, 0xfb, 0x47, 0x0c, 0x5e,
	0x3e, 0x0a, 0x64, 0xf4, 0x2a, 0x70, 0x06, 0xb2, 0x6b, 0x50, 0x29, 0xee, 0xf4, 0xb5, 0x9b, 0x15,
	0xdb, 0x52, 0x85, 0x6b, 0x91, 0x92, 0x03, 0xd1, 0x46, 0xb5, 0x36, 0x81, 0xdd, 0xa8, 0x15, 0x48,
	0x17, 0xdf, 0x98, 0xf1, 0x9d, 0xe7, 0xbf, 0xd8, 0xd6, 0x7c, 0x76, 0xe8, 0xeb, 0x7f, 0x02, 0x00,
	0x00, 0xff, 0xff, 0xb4, 0x45, 0xba, 0xc3, 0xcf, 0x03, 0x00, 0x00,
}
