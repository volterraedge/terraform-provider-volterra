// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ves.io/schema/views/api_inventory/types.proto

package api_inventory

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	golang_proto "github.com/golang/protobuf/proto"
	_ "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema"
	views "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/views"
	api_definition "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/views/api_definition"
	common_waf "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/views/common_waf"
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

// GlobalSpecType
//
// x-displayName: "Specification"
// x-required
// Shape of api_inventory in the storage backend.
type GlobalSpecType struct {
	// Internal API Groups Builders
	//
	// x-displayName: "Internal API Groups Builders"
	// API Groups builders define how to create API groups from a list of endpoints.
	// Note: the builders are managed via custom API.
	InternalApiGroupsBuilders []*api_definition.ApiGroupBuilder `protobuf:"bytes,1,rep,name=internal_api_groups_builders,json=internalApiGroupsBuilders,proto3" json:"internal_api_groups_builders,omitempty"` // Deprecated: Do not use.
	// API Definition
	//
	// x-displayName: "API Definition"
	// A reference to an API Definition object to be used as a source of user OpenAPI specifications
	ApiDefinitionRef *views.ObjectRefType `protobuf:"bytes,2,opt,name=api_definition_ref,json=apiDefinitionRef,proto3" json:"api_definition_ref,omitempty"`
	// Sensitive Data Discovery
	//
	// x-displayName: "Sensitive Data Discovery"
	// A reference to an Sensitive Data Discovery object
	SensitiveDataPolicyRef *views.ObjectRefType `protobuf:"bytes,3,opt,name=sensitive_data_policy_ref,json=sensitiveDataPolicyRef,proto3" json:"sensitive_data_policy_ref,omitempty"`
	// Code Scan Integrations
	//
	// x-displayName: "Code Scan Integrations"
	// A reference to an API Code Base Integration object
	SelectedCodeBaseIntegrations []*common_waf.CodeBaseIntegrationSelection `protobuf:"bytes,5,rep,name=selected_code_base_integrations,json=selectedCodeBaseIntegrations,proto3" json:"selected_code_base_integrations,omitempty"`
	// view_internal
	//
	// x-displayName: "View Internal"
	// Reference to view internal object.
	ViewInternal *views.ObjectRefType `protobuf:"bytes,1000,opt,name=view_internal,json=viewInternal,proto3" json:"view_internal,omitempty"`
}

func (m *GlobalSpecType) Reset()      { *m = GlobalSpecType{} }
func (*GlobalSpecType) ProtoMessage() {}
func (*GlobalSpecType) Descriptor() ([]byte, []int) {
	return fileDescriptor_a97fa5c4b65be4fc, []int{0}
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

// Deprecated: Do not use.
func (m *GlobalSpecType) GetInternalApiGroupsBuilders() []*api_definition.ApiGroupBuilder {
	if m != nil {
		return m.InternalApiGroupsBuilders
	}
	return nil
}

func (m *GlobalSpecType) GetApiDefinitionRef() *views.ObjectRefType {
	if m != nil {
		return m.ApiDefinitionRef
	}
	return nil
}

func (m *GlobalSpecType) GetSensitiveDataPolicyRef() *views.ObjectRefType {
	if m != nil {
		return m.SensitiveDataPolicyRef
	}
	return nil
}

func (m *GlobalSpecType) GetSelectedCodeBaseIntegrations() []*common_waf.CodeBaseIntegrationSelection {
	if m != nil {
		return m.SelectedCodeBaseIntegrations
	}
	return nil
}

func (m *GlobalSpecType) GetViewInternal() *views.ObjectRefType {
	if m != nil {
		return m.ViewInternal
	}
	return nil
}

func init() {
	proto.RegisterType((*GlobalSpecType)(nil), "ves.io.schema.views.api_inventory.GlobalSpecType")
	golang_proto.RegisterType((*GlobalSpecType)(nil), "ves.io.schema.views.api_inventory.GlobalSpecType")
}

func init() {
	proto.RegisterFile("ves.io/schema/views/api_inventory/types.proto", fileDescriptor_a97fa5c4b65be4fc)
}
func init() {
	golang_proto.RegisterFile("ves.io/schema/views/api_inventory/types.proto", fileDescriptor_a97fa5c4b65be4fc)
}

var fileDescriptor_a97fa5c4b65be4fc = []byte{
	// 575 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0x3f, 0x6f, 0xd3, 0x40,
	0x14, 0xcf, 0xd1, 0x94, 0x3f, 0x2e, 0x54, 0x55, 0x44, 0x51, 0x1a, 0xa2, 0x6b, 0xe9, 0xd4, 0x25,
	0x67, 0xa9, 0x15, 0x1b, 0x0b, 0xa6, 0x52, 0x55, 0x16, 0x50, 0x8a, 0x84, 0xc4, 0xc0, 0xe9, 0x6c,
	0x3f, 0xbb, 0x07, 0x8e, 0xef, 0x74, 0x77, 0x71, 0x89, 0x50, 0x04, 0x2b, 0x1b, 0x0b, 0xdf, 0x01,
	0xf5, 0x2b, 0xb0, 0x30, 0x32, 0x76, 0xcc, 0x48, 0x9d, 0xa5, 0x6c, 0xfd, 0x08, 0xc8, 0x97, 0xa4,
	0x8d, 0xc1, 0x95, 0xba, 0xdd, 0xdd, 0xfb, 0xfd, 0x7b, 0x77, 0xcf, 0x76, 0x3a, 0x19, 0x68, 0xc2,
	0x85, 0xab, 0x83, 0x43, 0xe8, 0x31, 0x37, 0xe3, 0x70, 0xa4, 0x5d, 0x26, 0x39, 0xe5, 0x69, 0x06,
	0xa9, 0x11, 0x6a, 0xe0, 0x9a, 0x81, 0x04, 0x4d, 0xa4, 0x12, 0x46, 0x34, 0x1e, 0x4d, 0xe0, 0x64,
	0x02, 0x27, 0x16, 0x4e, 0x4a, 0xf0, 0x56, 0x27, 0xe6, 0xe6, 0xb0, 0xef, 0x93, 0x40, 0xf4, 0xdc,
	0x58, 0xc4, 0xc2, 0xb5, 0x4c, 0xbf, 0x1f, 0xd9, 0x9d, 0xdd, 0xd8, 0xd5, 0x44, 0xb1, 0xf5, 0xb0,
	0x1c, 0x40, 0x48, 0xc3, 0x45, 0x3a, 0xb5, 0x6b, 0xad, 0x95, 0x8b, 0x73, 0x49, 0x5a, 0xed, 0x7f,
	0x82, 0xb3, 0x84, 0x87, 0xcc, 0xc0, 0xb4, 0xba, 0xf1, 0x7f, 0x5b, 0xb4, 0x2c, 0x4d, 0xae, 0x6a,
	0x3c, 0x84, 0x88, 0xa7, 0xbc, 0x80, 0x96, 0xfc, 0x1e, 0x57, 0xe1, 0x03, 0xd1, 0xeb, 0x89, 0x94,
	0x1e, 0xb1, 0x68, 0x6e, 0x49, 0xe7, 0x69, 0xeb, 0x55, 0xb4, 0x39, 0xc0, 0xe6, 0x9f, 0xba, 0xb3,
	0xbc, 0x97, 0x08, 0x9f, 0x25, 0x07, 0x12, 0x82, 0x57, 0x03, 0x09, 0x8d, 0x4f, 0x4e, 0x9b, 0xa7,
	0x06, 0x54, 0xca, 0x12, 0x5a, 0x24, 0x8a, 0x95, 0xe8, 0x4b, 0x4d, 0xfd, 0x3e, 0x4f, 0x42, 0x50,
	0xba, 0x89, 0x36, 0x16, 0xb6, 0x96, 0xb6, 0x77, 0xc8, 0x55, 0x6f, 0x71, 0xd9, 0x01, 0x79, 0x2a,
	0xf9, 0x5e, 0xc1, 0xf6, 0x26, 0x5c, 0xef, 0xce, 0x68, 0x88, 0xce, 0x7e, 0x20, 0xd4, 0x44, 0xdd,
	0xb5, 0x99, 0xc7, 0x0c, 0xa3, 0xa7, 0x20, 0xdd, 0x88, 0x9d, 0x46, 0x59, 0x87, 0x2a, 0x88, 0x9a,
	0x37, 0x36, 0xd0, 0xd6, 0xd2, 0xf6, 0x66, 0xa5, 0xed, 0x0b, 0xff, 0x1d, 0x04, 0xa6, 0x0b, 0x51,
	0xd1, 0x80, 0xb7, 0x7a, 0x3c, 0x5c, 0x2e, 0x0b, 0x14, 0x9e, 0xdd, 0x15, 0x26, 0xf9, 0xee, 0xc5,
	0x51, 0x17, 0xa2, 0xc6, 0x47, 0x67, 0x4d, 0x43, 0xaa, 0xb9, 0xe1, 0x19, 0xd0, 0x90, 0x19, 0x46,
	0xa5, 0x48, 0x78, 0x30, 0xb0, 0x7e, 0x0b, 0xd7, 0xf6, 0x6b, 0x1f, 0x0f, 0x57, 0x2b, 0x75, 0xac,
	0xed, 0x83, 0x8b, 0xd2, 0x2e, 0x33, 0xec, 0xa5, 0x2d, 0x14, 0xe6, 0x5f, 0x90, 0xb3, 0xae, 0x21,
	0x81, 0xc0, 0x40, 0x48, 0x03, 0x11, 0x02, 0xf5, 0x99, 0x06, 0x5a, 0x5c, 0x4b, 0xac, 0x98, 0x9d,
	0x95, 0xe6, 0xa2, 0xbd, 0xea, 0x27, 0x95, 0x19, 0x2e, 0x5f, 0x9c, 0x3c, 0x13, 0x21, 0x78, 0x4c,
	0xc3, 0xfe, 0x25, 0xf7, 0xc0, 0x2a, 0x73, 0x91, 0x7a, 0x75, 0x9b, 0xa2, 0x3d, 0xb3, 0xaa, 0xc0,
	0xea, 0xc6, 0x5b, 0xe7, 0x9e, 0x9d, 0xd1, 0xd9, 0x9b, 0x34, 0xcf, 0x6e, 0x5d, 0xbb, 0xfb, 0xfb,
	0xc7, 0xc3, 0x32, 0xd9, 0xfa, 0xdd, 0x2d, 0x8e, 0xf6, 0xa7, 0x27, 0xcf, 0xeb, 0xb7, 0xeb, 0x2b,
	0x8b, 0xde, 0x37, 0x74, 0x72, 0x8a, 0x6b, 0xa3, 0x53, 0x5c, 0x3b, 0x3f, 0xc5, 0xe8, 0x73, 0x8e,
	0xd1, 0xf7, 0x1c, 0xa3, 0x5f, 0x39, 0x46, 0x27, 0x39, 0x46, 0xa3, 0x1c, 0xa3, 0xdf, 0x39, 0x46,
	0x67, 0x39, 0xae, 0x9d, 0xe7, 0x18, 0x7d, 0x1d, 0xe3, 0xda, 0xcf, 0x31, 0x46, 0x27, 0x63, 0x5c,
	0x1b, 0x8d, 0x71, 0xed, 0xcd, 0xeb, 0x58, 0xc8, 0xf7, 0x31, 0xc9, 0x44, 0x62, 0x40, 0x29, 0x46,
	0xfa, 0xda, 0xb5, 0x8b, 0x48, 0xa8, 0x5e, 0x47, 0x2a, 0x91, 0xf1, 0x10, 0x54, 0x67, 0x56, 0x76,
	0xa5, 0x1f, 0x0b, 0x17, 0x3e, 0x98, 0xe9, 0xe8, 0x5f, 0xfd, 0x87, 0xf1, 0x6f, 0xda, 0x4f, 0x61,
	0xe7, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x8a, 0x15, 0x34, 0xd1, 0x8d, 0x04, 0x00, 0x00,
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
	if len(this.InternalApiGroupsBuilders) != len(that1.InternalApiGroupsBuilders) {
		return false
	}
	for i := range this.InternalApiGroupsBuilders {
		if !this.InternalApiGroupsBuilders[i].Equal(that1.InternalApiGroupsBuilders[i]) {
			return false
		}
	}
	if !this.ApiDefinitionRef.Equal(that1.ApiDefinitionRef) {
		return false
	}
	if !this.SensitiveDataPolicyRef.Equal(that1.SensitiveDataPolicyRef) {
		return false
	}
	if len(this.SelectedCodeBaseIntegrations) != len(that1.SelectedCodeBaseIntegrations) {
		return false
	}
	for i := range this.SelectedCodeBaseIntegrations {
		if !this.SelectedCodeBaseIntegrations[i].Equal(that1.SelectedCodeBaseIntegrations[i]) {
			return false
		}
	}
	if !this.ViewInternal.Equal(that1.ViewInternal) {
		return false
	}
	return true
}
func (this *GlobalSpecType) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 9)
	s = append(s, "&api_inventory.GlobalSpecType{")
	if this.InternalApiGroupsBuilders != nil {
		s = append(s, "InternalApiGroupsBuilders: "+fmt.Sprintf("%#v", this.InternalApiGroupsBuilders)+",\n")
	}
	if this.ApiDefinitionRef != nil {
		s = append(s, "ApiDefinitionRef: "+fmt.Sprintf("%#v", this.ApiDefinitionRef)+",\n")
	}
	if this.SensitiveDataPolicyRef != nil {
		s = append(s, "SensitiveDataPolicyRef: "+fmt.Sprintf("%#v", this.SensitiveDataPolicyRef)+",\n")
	}
	if this.SelectedCodeBaseIntegrations != nil {
		s = append(s, "SelectedCodeBaseIntegrations: "+fmt.Sprintf("%#v", this.SelectedCodeBaseIntegrations)+",\n")
	}
	if this.ViewInternal != nil {
		s = append(s, "ViewInternal: "+fmt.Sprintf("%#v", this.ViewInternal)+",\n")
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
	if m.ViewInternal != nil {
		{
			size, err := m.ViewInternal.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x3e
		i--
		dAtA[i] = 0xc2
	}
	if len(m.SelectedCodeBaseIntegrations) > 0 {
		for iNdEx := len(m.SelectedCodeBaseIntegrations) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.SelectedCodeBaseIntegrations[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTypes(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if m.SensitiveDataPolicyRef != nil {
		{
			size, err := m.SensitiveDataPolicyRef.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.ApiDefinitionRef != nil {
		{
			size, err := m.ApiDefinitionRef.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.InternalApiGroupsBuilders) > 0 {
		for iNdEx := len(m.InternalApiGroupsBuilders) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.InternalApiGroupsBuilders[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.InternalApiGroupsBuilders) > 0 {
		for _, e := range m.InternalApiGroupsBuilders {
			l = e.Size()
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	if m.ApiDefinitionRef != nil {
		l = m.ApiDefinitionRef.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.SensitiveDataPolicyRef != nil {
		l = m.SensitiveDataPolicyRef.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	if len(m.SelectedCodeBaseIntegrations) > 0 {
		for _, e := range m.SelectedCodeBaseIntegrations {
			l = e.Size()
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	if m.ViewInternal != nil {
		l = m.ViewInternal.Size()
		n += 2 + l + sovTypes(uint64(l))
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
	repeatedStringForInternalApiGroupsBuilders := "[]*ApiGroupBuilder{"
	for _, f := range this.InternalApiGroupsBuilders {
		repeatedStringForInternalApiGroupsBuilders += strings.Replace(fmt.Sprintf("%v", f), "ApiGroupBuilder", "api_definition.ApiGroupBuilder", 1) + ","
	}
	repeatedStringForInternalApiGroupsBuilders += "}"
	repeatedStringForSelectedCodeBaseIntegrations := "[]*CodeBaseIntegrationSelection{"
	for _, f := range this.SelectedCodeBaseIntegrations {
		repeatedStringForSelectedCodeBaseIntegrations += strings.Replace(fmt.Sprintf("%v", f), "CodeBaseIntegrationSelection", "common_waf.CodeBaseIntegrationSelection", 1) + ","
	}
	repeatedStringForSelectedCodeBaseIntegrations += "}"
	s := strings.Join([]string{`&GlobalSpecType{`,
		`InternalApiGroupsBuilders:` + repeatedStringForInternalApiGroupsBuilders + `,`,
		`ApiDefinitionRef:` + strings.Replace(fmt.Sprintf("%v", this.ApiDefinitionRef), "ObjectRefType", "views.ObjectRefType", 1) + `,`,
		`SensitiveDataPolicyRef:` + strings.Replace(fmt.Sprintf("%v", this.SensitiveDataPolicyRef), "ObjectRefType", "views.ObjectRefType", 1) + `,`,
		`SelectedCodeBaseIntegrations:` + repeatedStringForSelectedCodeBaseIntegrations + `,`,
		`ViewInternal:` + strings.Replace(fmt.Sprintf("%v", this.ViewInternal), "ObjectRefType", "views.ObjectRefType", 1) + `,`,
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
				return fmt.Errorf("proto: wrong wireType = %d for field InternalApiGroupsBuilders", wireType)
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
			m.InternalApiGroupsBuilders = append(m.InternalApiGroupsBuilders, &api_definition.ApiGroupBuilder{})
			if err := m.InternalApiGroupsBuilders[len(m.InternalApiGroupsBuilders)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ApiDefinitionRef", wireType)
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
			if m.ApiDefinitionRef == nil {
				m.ApiDefinitionRef = &views.ObjectRefType{}
			}
			if err := m.ApiDefinitionRef.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SensitiveDataPolicyRef", wireType)
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
			if m.SensitiveDataPolicyRef == nil {
				m.SensitiveDataPolicyRef = &views.ObjectRefType{}
			}
			if err := m.SensitiveDataPolicyRef.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SelectedCodeBaseIntegrations", wireType)
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
			m.SelectedCodeBaseIntegrations = append(m.SelectedCodeBaseIntegrations, &common_waf.CodeBaseIntegrationSelection{})
			if err := m.SelectedCodeBaseIntegrations[len(m.SelectedCodeBaseIntegrations)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 1000:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ViewInternal", wireType)
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
			if m.ViewInternal == nil {
				m.ViewInternal = &views.ObjectRefType{}
			}
			if err := m.ViewInternal.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
