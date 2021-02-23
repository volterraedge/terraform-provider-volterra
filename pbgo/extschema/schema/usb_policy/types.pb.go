// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ves.io/schema/usb_policy/types.proto

package usb_policy

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import _ "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema"
import _ "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema"
import _ "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema"
import _ "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema"
import _ "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/views"

import strings "strings"
import reflect "reflect"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Rule
//
// x-displayName: "usb enablement rule"
// usb enablement rule
type Rule struct {
	// i_serial_number
	//
	// x-displayName: "iSerialNumber"
	// x-example: "0000:00:14.0"
	// Index of Serial Number String Descriptor
	ISerial string `protobuf:"bytes,1,opt,name=i_serial,json=iSerial,proto3" json:"i_serial,omitempty"`
	// id_vendor
	//
	// x-displayName: "Vendor ID"
	// x-example: "0x1d6b"
	// Vendor ID (Assigned by USB Org) in hex
	IdVendor string `protobuf:"bytes,2,opt,name=id_vendor,json=idVendor,proto3" json:"id_vendor,omitempty"`
	// id_product
	//
	// x-displayName: "Product ID"
	// x-example: "0x0002"
	// Product ID (Assigned by Manufacturer) in hex
	IdProduct string `protobuf:"bytes,3,opt,name=id_product,json=idProduct,proto3" json:"id_product,omitempty"`
	// b_device_class
	//
	// x-displayName: "Class"
	// x-example: "hub"
	// The class of this device
	BDeviceClass string `protobuf:"bytes,4,opt,name=b_device_class,json=bDeviceClass,proto3" json:"b_device_class,omitempty"`
	// b_device_sub_class
	//
	// x-displayName: "Subclass"
	// x-example: "hub"
	// The sub-class (within the class) of this device
	BDeviceSubClass string `protobuf:"bytes,5,opt,name=b_device_sub_class,json=bDeviceSubClass,proto3" json:"b_device_sub_class,omitempty"`
	// b_device_protocol
	//
	// x-displayName: "Protocol"
	// x-example: "0002"
	// The protocol (within the sub-class) of this device
	BDeviceProtocol string `protobuf:"bytes,6,opt,name=b_device_protocol,json=bDeviceProtocol,proto3" json:"b_device_protocol,omitempty"`
}

func (m *Rule) Reset()                    { *m = Rule{} }
func (*Rule) ProtoMessage()               {}
func (*Rule) Descriptor() ([]byte, []int) { return fileDescriptorTypes, []int{0} }

func (m *Rule) GetISerial() string {
	if m != nil {
		return m.ISerial
	}
	return ""
}

func (m *Rule) GetIdVendor() string {
	if m != nil {
		return m.IdVendor
	}
	return ""
}

func (m *Rule) GetIdProduct() string {
	if m != nil {
		return m.IdProduct
	}
	return ""
}

func (m *Rule) GetBDeviceClass() string {
	if m != nil {
		return m.BDeviceClass
	}
	return ""
}

func (m *Rule) GetBDeviceSubClass() string {
	if m != nil {
		return m.BDeviceSubClass
	}
	return ""
}

func (m *Rule) GetBDeviceProtocol() string {
	if m != nil {
		return m.BDeviceProtocol
	}
	return ""
}

// Specification for USB policy
//
// x-displayName: "Specification"
// Shape of the USB policy object
type GlobalSpecType struct {
	// Allowed USB devices
	//
	// x-displayName: "Allowed USB devices"
	// x-required
	// List of allowed USB devices
	AllowedDevices []*Rule `protobuf:"bytes,1,rep,name=allowed_devices,json=allowedDevices" json:"allowed_devices,omitempty"`
}

func (m *GlobalSpecType) Reset()                    { *m = GlobalSpecType{} }
func (*GlobalSpecType) ProtoMessage()               {}
func (*GlobalSpecType) Descriptor() ([]byte, []int) { return fileDescriptorTypes, []int{1} }

func (m *GlobalSpecType) GetAllowedDevices() []*Rule {
	if m != nil {
		return m.AllowedDevices
	}
	return nil
}

// Create USB policy
//
// x-displayName: "Create USB policy"
// Creates a new USB policy object
type CreateSpecType struct {
	AllowedDevices []*Rule `protobuf:"bytes,1,rep,name=allowed_devices,json=allowedDevices" json:"allowed_devices,omitempty"`
}

func (m *CreateSpecType) Reset()                    { *m = CreateSpecType{} }
func (*CreateSpecType) ProtoMessage()               {}
func (*CreateSpecType) Descriptor() ([]byte, []int) { return fileDescriptorTypes, []int{2} }

func (m *CreateSpecType) GetAllowedDevices() []*Rule {
	if m != nil {
		return m.AllowedDevices
	}
	return nil
}

// Replace USB policy
//
// x-displayName: "Replace USB policy"
// Replaces the content of an USB policy object
type ReplaceSpecType struct {
	AllowedDevices []*Rule `protobuf:"bytes,1,rep,name=allowed_devices,json=allowedDevices" json:"allowed_devices,omitempty"`
}

func (m *ReplaceSpecType) Reset()                    { *m = ReplaceSpecType{} }
func (*ReplaceSpecType) ProtoMessage()               {}
func (*ReplaceSpecType) Descriptor() ([]byte, []int) { return fileDescriptorTypes, []int{3} }

func (m *ReplaceSpecType) GetAllowedDevices() []*Rule {
	if m != nil {
		return m.AllowedDevices
	}
	return nil
}

// Get USB policy
//
// x-displayName: "Get USB policy"
// Get the USB policy object
type GetSpecType struct {
	AllowedDevices []*Rule `protobuf:"bytes,1,rep,name=allowed_devices,json=allowedDevices" json:"allowed_devices,omitempty"`
}

func (m *GetSpecType) Reset()                    { *m = GetSpecType{} }
func (*GetSpecType) ProtoMessage()               {}
func (*GetSpecType) Descriptor() ([]byte, []int) { return fileDescriptorTypes, []int{4} }

func (m *GetSpecType) GetAllowedDevices() []*Rule {
	if m != nil {
		return m.AllowedDevices
	}
	return nil
}

func init() {
	proto.RegisterType((*Rule)(nil), "ves.io.schema.usb_policy.Rule")
	proto.RegisterType((*GlobalSpecType)(nil), "ves.io.schema.usb_policy.GlobalSpecType")
	proto.RegisterType((*CreateSpecType)(nil), "ves.io.schema.usb_policy.CreateSpecType")
	proto.RegisterType((*ReplaceSpecType)(nil), "ves.io.schema.usb_policy.ReplaceSpecType")
	proto.RegisterType((*GetSpecType)(nil), "ves.io.schema.usb_policy.GetSpecType")
}
func (this *Rule) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Rule)
	if !ok {
		that2, ok := that.(Rule)
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
	if this.ISerial != that1.ISerial {
		return false
	}
	if this.IdVendor != that1.IdVendor {
		return false
	}
	if this.IdProduct != that1.IdProduct {
		return false
	}
	if this.BDeviceClass != that1.BDeviceClass {
		return false
	}
	if this.BDeviceSubClass != that1.BDeviceSubClass {
		return false
	}
	if this.BDeviceProtocol != that1.BDeviceProtocol {
		return false
	}
	return true
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
	if len(this.AllowedDevices) != len(that1.AllowedDevices) {
		return false
	}
	for i := range this.AllowedDevices {
		if !this.AllowedDevices[i].Equal(that1.AllowedDevices[i]) {
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
	if len(this.AllowedDevices) != len(that1.AllowedDevices) {
		return false
	}
	for i := range this.AllowedDevices {
		if !this.AllowedDevices[i].Equal(that1.AllowedDevices[i]) {
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
	if len(this.AllowedDevices) != len(that1.AllowedDevices) {
		return false
	}
	for i := range this.AllowedDevices {
		if !this.AllowedDevices[i].Equal(that1.AllowedDevices[i]) {
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
	if len(this.AllowedDevices) != len(that1.AllowedDevices) {
		return false
	}
	for i := range this.AllowedDevices {
		if !this.AllowedDevices[i].Equal(that1.AllowedDevices[i]) {
			return false
		}
	}
	return true
}
func (this *Rule) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 10)
	s = append(s, "&usb_policy.Rule{")
	s = append(s, "ISerial: "+fmt.Sprintf("%#v", this.ISerial)+",\n")
	s = append(s, "IdVendor: "+fmt.Sprintf("%#v", this.IdVendor)+",\n")
	s = append(s, "IdProduct: "+fmt.Sprintf("%#v", this.IdProduct)+",\n")
	s = append(s, "BDeviceClass: "+fmt.Sprintf("%#v", this.BDeviceClass)+",\n")
	s = append(s, "BDeviceSubClass: "+fmt.Sprintf("%#v", this.BDeviceSubClass)+",\n")
	s = append(s, "BDeviceProtocol: "+fmt.Sprintf("%#v", this.BDeviceProtocol)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *GlobalSpecType) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&usb_policy.GlobalSpecType{")
	if this.AllowedDevices != nil {
		s = append(s, "AllowedDevices: "+fmt.Sprintf("%#v", this.AllowedDevices)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *CreateSpecType) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&usb_policy.CreateSpecType{")
	if this.AllowedDevices != nil {
		s = append(s, "AllowedDevices: "+fmt.Sprintf("%#v", this.AllowedDevices)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *ReplaceSpecType) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&usb_policy.ReplaceSpecType{")
	if this.AllowedDevices != nil {
		s = append(s, "AllowedDevices: "+fmt.Sprintf("%#v", this.AllowedDevices)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *GetSpecType) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&usb_policy.GetSpecType{")
	if this.AllowedDevices != nil {
		s = append(s, "AllowedDevices: "+fmt.Sprintf("%#v", this.AllowedDevices)+",\n")
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
func (m *Rule) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Rule) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.ISerial) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintTypes(dAtA, i, uint64(len(m.ISerial)))
		i += copy(dAtA[i:], m.ISerial)
	}
	if len(m.IdVendor) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintTypes(dAtA, i, uint64(len(m.IdVendor)))
		i += copy(dAtA[i:], m.IdVendor)
	}
	if len(m.IdProduct) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintTypes(dAtA, i, uint64(len(m.IdProduct)))
		i += copy(dAtA[i:], m.IdProduct)
	}
	if len(m.BDeviceClass) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintTypes(dAtA, i, uint64(len(m.BDeviceClass)))
		i += copy(dAtA[i:], m.BDeviceClass)
	}
	if len(m.BDeviceSubClass) > 0 {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintTypes(dAtA, i, uint64(len(m.BDeviceSubClass)))
		i += copy(dAtA[i:], m.BDeviceSubClass)
	}
	if len(m.BDeviceProtocol) > 0 {
		dAtA[i] = 0x32
		i++
		i = encodeVarintTypes(dAtA, i, uint64(len(m.BDeviceProtocol)))
		i += copy(dAtA[i:], m.BDeviceProtocol)
	}
	return i, nil
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
	if len(m.AllowedDevices) > 0 {
		for _, msg := range m.AllowedDevices {
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
	if len(m.AllowedDevices) > 0 {
		for _, msg := range m.AllowedDevices {
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
	if len(m.AllowedDevices) > 0 {
		for _, msg := range m.AllowedDevices {
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
	if len(m.AllowedDevices) > 0 {
		for _, msg := range m.AllowedDevices {
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
func NewPopulatedRule(r randyTypes, easy bool) *Rule {
	this := &Rule{}
	this.ISerial = string(randStringTypes(r))
	this.IdVendor = string(randStringTypes(r))
	this.IdProduct = string(randStringTypes(r))
	this.BDeviceClass = string(randStringTypes(r))
	this.BDeviceSubClass = string(randStringTypes(r))
	this.BDeviceProtocol = string(randStringTypes(r))
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedGlobalSpecType(r randyTypes, easy bool) *GlobalSpecType {
	this := &GlobalSpecType{}
	if r.Intn(10) != 0 {
		v1 := r.Intn(5)
		this.AllowedDevices = make([]*Rule, v1)
		for i := 0; i < v1; i++ {
			this.AllowedDevices[i] = NewPopulatedRule(r, easy)
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
		this.AllowedDevices = make([]*Rule, v2)
		for i := 0; i < v2; i++ {
			this.AllowedDevices[i] = NewPopulatedRule(r, easy)
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
		this.AllowedDevices = make([]*Rule, v3)
		for i := 0; i < v3; i++ {
			this.AllowedDevices[i] = NewPopulatedRule(r, easy)
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
		this.AllowedDevices = make([]*Rule, v4)
		for i := 0; i < v4; i++ {
			this.AllowedDevices[i] = NewPopulatedRule(r, easy)
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
func (m *Rule) Size() (n int) {
	var l int
	_ = l
	l = len(m.ISerial)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.IdVendor)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.IdProduct)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.BDeviceClass)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.BDeviceSubClass)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.BDeviceProtocol)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}

func (m *GlobalSpecType) Size() (n int) {
	var l int
	_ = l
	if len(m.AllowedDevices) > 0 {
		for _, e := range m.AllowedDevices {
			l = e.Size()
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	return n
}

func (m *CreateSpecType) Size() (n int) {
	var l int
	_ = l
	if len(m.AllowedDevices) > 0 {
		for _, e := range m.AllowedDevices {
			l = e.Size()
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	return n
}

func (m *ReplaceSpecType) Size() (n int) {
	var l int
	_ = l
	if len(m.AllowedDevices) > 0 {
		for _, e := range m.AllowedDevices {
			l = e.Size()
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	return n
}

func (m *GetSpecType) Size() (n int) {
	var l int
	_ = l
	if len(m.AllowedDevices) > 0 {
		for _, e := range m.AllowedDevices {
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
func (this *Rule) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Rule{`,
		`ISerial:` + fmt.Sprintf("%v", this.ISerial) + `,`,
		`IdVendor:` + fmt.Sprintf("%v", this.IdVendor) + `,`,
		`IdProduct:` + fmt.Sprintf("%v", this.IdProduct) + `,`,
		`BDeviceClass:` + fmt.Sprintf("%v", this.BDeviceClass) + `,`,
		`BDeviceSubClass:` + fmt.Sprintf("%v", this.BDeviceSubClass) + `,`,
		`BDeviceProtocol:` + fmt.Sprintf("%v", this.BDeviceProtocol) + `,`,
		`}`,
	}, "")
	return s
}
func (this *GlobalSpecType) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&GlobalSpecType{`,
		`AllowedDevices:` + strings.Replace(fmt.Sprintf("%v", this.AllowedDevices), "Rule", "Rule", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *CreateSpecType) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&CreateSpecType{`,
		`AllowedDevices:` + strings.Replace(fmt.Sprintf("%v", this.AllowedDevices), "Rule", "Rule", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *ReplaceSpecType) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ReplaceSpecType{`,
		`AllowedDevices:` + strings.Replace(fmt.Sprintf("%v", this.AllowedDevices), "Rule", "Rule", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *GetSpecType) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&GetSpecType{`,
		`AllowedDevices:` + strings.Replace(fmt.Sprintf("%v", this.AllowedDevices), "Rule", "Rule", 1) + `,`,
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
func (m *Rule) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: Rule: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Rule: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ISerial", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ISerial = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IdVendor", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.IdVendor = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IdProduct", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.IdProduct = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BDeviceClass", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BDeviceClass = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BDeviceSubClass", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BDeviceSubClass = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BDeviceProtocol", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BDeviceProtocol = string(dAtA[iNdEx:postIndex])
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
				return fmt.Errorf("proto: wrong wireType = %d for field AllowedDevices", wireType)
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
			m.AllowedDevices = append(m.AllowedDevices, &Rule{})
			if err := m.AllowedDevices[len(m.AllowedDevices)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
				return fmt.Errorf("proto: wrong wireType = %d for field AllowedDevices", wireType)
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
			m.AllowedDevices = append(m.AllowedDevices, &Rule{})
			if err := m.AllowedDevices[len(m.AllowedDevices)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
				return fmt.Errorf("proto: wrong wireType = %d for field AllowedDevices", wireType)
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
			m.AllowedDevices = append(m.AllowedDevices, &Rule{})
			if err := m.AllowedDevices[len(m.AllowedDevices)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
				return fmt.Errorf("proto: wrong wireType = %d for field AllowedDevices", wireType)
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
			m.AllowedDevices = append(m.AllowedDevices, &Rule{})
			if err := m.AllowedDevices[len(m.AllowedDevices)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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

func init() { proto.RegisterFile("ves.io/schema/usb_policy/types.proto", fileDescriptorTypes) }

var fileDescriptorTypes = []byte{
	// 534 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x94, 0xbf, 0x6f, 0xd3, 0x40,
	0x14, 0xc7, 0xf3, 0x9a, 0xb6, 0x24, 0x57, 0x94, 0x50, 0xc3, 0xe0, 0xa6, 0x70, 0x58, 0x51, 0x87,
	0x0a, 0x14, 0x5b, 0xc0, 0xc6, 0x80, 0x44, 0x8a, 0x94, 0x0d, 0x55, 0x09, 0x30, 0x20, 0x21, 0xcb,
	0x3f, 0x5e, 0xdd, 0x13, 0x97, 0xde, 0xc9, 0x67, 0xbb, 0x64, 0xe3, 0x0f, 0x40, 0x02, 0xf1, 0x57,
	0xf0, 0x27, 0x20, 0x58, 0x2a, 0x26, 0xc4, 0xd4, 0xb1, 0x03, 0x03, 0x31, 0x0b, 0x6c, 0x1d, 0x19,
	0x51, 0xce, 0x6e, 0x68, 0x1a, 0xb1, 0x74, 0xe8, 0xf6, 0xde, 0xf7, 0xfb, 0xf1, 0xfb, 0xbe, 0xf3,
	0x2f, 0xb2, 0x91, 0xa1, 0xb2, 0x99, 0x70, 0x54, 0xb0, 0x8b, 0x43, 0xcf, 0x49, 0x95, 0xef, 0x4a,
	0xc1, 0x59, 0x30, 0x72, 0x92, 0x91, 0x44, 0x65, 0xcb, 0x58, 0x24, 0xc2, 0x30, 0x0b, 0xca, 0x2e,
	0x28, 0xfb, 0x1f, 0xd5, 0xea, 0x44, 0x2c, 0xd9, 0x4d, 0x7d, 0x3b, 0x10, 0x43, 0x27, 0x12, 0x91,
	0x70, 0xf4, 0x05, 0x7e, 0xba, 0xa3, 0x3b, 0xdd, 0xe8, 0xaa, 0x18, 0xd4, 0x5a, 0x9f, 0x8d, 0x13,
	0x32, 0x61, 0x62, 0xaf, 0x4c, 0x69, 0xad, 0xcd, 0x9a, 0xa7, 0x16, 0x68, 0x5d, 0x9f, 0xb5, 0x32,
	0x8f, 0xb3, 0xd0, 0x4b, 0xb0, 0x74, 0xad, 0x33, 0x2e, 0xc3, 0x7d, 0x77, 0x76, 0xf4, 0xcd, 0x79,
	0x42, 0x9d, 0x0e, 0x68, 0x7f, 0x07, 0xb2, 0xd8, 0x4f, 0x39, 0x1a, 0x6b, 0xa4, 0xc6, 0x5c, 0x85,
	0x31, 0xf3, 0xb8, 0x09, 0x16, 0x6c, 0xd6, 0xfb, 0x97, 0xd8, 0x40, 0xb7, 0xc6, 0x3a, 0xa9, 0xb3,
	0xd0, 0xcd, 0x70, 0x2f, 0x14, 0xb1, 0xb9, 0xa0, 0xbd, 0x1a, 0x0b, 0x9f, 0xe9, 0xde, 0xb8, 0x41,
	0x08, 0x0b, 0x5d, 0x19, 0x8b, 0x30, 0x0d, 0x12, 0xb3, 0xaa, 0xdd, 0x3a, 0x0b, 0xb7, 0x0b, 0xc1,
	0xd8, 0x20, 0x0d, 0xdf, 0x0d, 0x31, 0x63, 0x01, 0xba, 0x01, 0xf7, 0x94, 0x32, 0x17, 0x35, 0x72,
	0xd9, 0x7f, 0xa4, 0xc5, 0xad, 0x89, 0x66, 0xdc, 0x26, 0xc6, 0x94, 0x52, 0xa9, 0x5f, 0x92, 0x4b,
	0x9a, 0x6c, 0x96, 0xe4, 0x20, 0xf5, 0x0b, 0xf8, 0x16, 0x59, 0x9d, 0xc2, 0xfa, 0x10, 0x81, 0xe0,
	0xe6, 0xf2, 0x0c, 0xbb, 0x5d, 0xca, 0xed, 0xb7, 0x40, 0x1a, 0x3d, 0x2e, 0x7c, 0x8f, 0x0f, 0x24,
	0x06, 0x4f, 0x46, 0x12, 0x8d, 0x17, 0xa4, 0xe9, 0x71, 0x2e, 0xf6, 0x31, 0x2c, 0x87, 0x28, 0x13,
	0xac, 0xea, 0xe6, 0xca, 0x5d, 0x6a, 0xff, 0xef, 0x69, 0xdb, 0x93, 0x3b, 0xd4, 0xbd, 0xf6, 0xe9,
	0xf7, 0x41, 0x75, 0xe9, 0x3d, 0x2c, 0x5c, 0xb1, 0x4e, 0xaa, 0x1a, 0xf4, 0x1b, 0xe5, 0xb0, 0x22,
	0x58, 0xdd, 0xa7, 0x5f, 0x3e, 0x43, 0x8b, 0x98, 0xe4, 0xea, 0xc3, 0x42, 0xb7, 0x9e, 0x0e, 0xba,
	0x56, 0x19, 0x64, 0xc0, 0x9d, 0x36, 0x27, 0x8d, 0xad, 0x18, 0xbd, 0x04, 0xa7, 0x0b, 0xf5, 0xce,
	0xb9, 0xd0, 0x5c, 0xf4, 0xea, 0xb7, 0x07, 0x67, 0x0e, 0xdb, 0x1e, 0x92, 0x66, 0x1f, 0x25, 0xf7,
	0x82, 0x8b, 0x89, 0x63, 0x64, 0xa5, 0x87, 0xc9, 0x45, 0x44, 0x75, 0xdf, 0xc0, 0xe1, 0x98, 0x56,
	0x8e, 0xc6, 0xb4, 0x72, 0x3c, 0xa6, 0xf0, 0x67, 0x4c, 0xe1, 0x75, 0x4e, 0xe1, 0x43, 0x4e, 0xe1,
	0x63, 0x4e, 0xe1, 0x20, 0xa7, 0xf0, 0x35, 0xa7, 0x70, 0x98, 0x53, 0x38, 0xca, 0x29, 0xfc, 0xc8,
	0x29, 0xfc, 0xca, 0x69, 0xe5, 0x38, 0xa7, 0xf0, 0xee, 0x27, 0xad, 0x3c, 0x7f, 0x1c, 0x09, 0xf9,
	0x32, 0xb2, 0x33, 0xc1, 0x13, 0x8c, 0xe3, 0xc9, 0x06, 0x8e, 0x2e, 0x76, 0x44, 0x3c, 0xec, 0xc8,
	0x58, 0x64, 0x2c, 0xc4, 0xb8, 0x73, 0x62, 0x3b, 0xd2, 0x8f, 0x84, 0x83, 0xaf, 0x92, 0xf2, 0x3b,
	0x9a, 0xfb, 0x6b, 0xf8, 0xcb, 0xfa, 0x4d, 0xbc, 0xf7, 0x37, 0x00, 0x00, 0xff, 0xff, 0x9f, 0x27,
	0xf5, 0x7e, 0x58, 0x04, 0x00, 0x00,
}