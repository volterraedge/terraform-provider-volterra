// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ves.io/schema/saas_config/types.proto

package saas_config

import (
	fmt "fmt"
	_ "github.com/gogo/googleapis/google/api"
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

type GlobalSpecType struct {
	// plan_type
	//
	// x-displayName: "Plan Type"
	// Type of the billing plan the tenant is subscribed to. Value can be one of FREE, INDIVIDUAL, TEAM, ORGANIZATION.
	PlanType schema.PlanType `protobuf:"varint,1,opt,name=plan_type,json=planType,proto3,enum=ves.io.schema.PlanType" json:"plan_type,omitempty"`
	// crm_info
	//
	// x-displayName: "CRM Information"
	// This field holds CRM information
	CrmInfo *schema.CRMInfo `protobuf:"bytes,2,opt,name=crm_info,json=crmInfo,proto3" json:"crm_info,omitempty"`
	// addon_services_subscribed
	//
	// x-displayName: "Addon Services Subscribed"
	// x-example: "shape-bot"
	// List of addon service names currently subscribed by the tenant.
	AddonServicesSubscribed []string `protobuf:"bytes,3,rep,name=addon_services_subscribed,json=addonServicesSubscribed,proto3" json:"addon_services_subscribed,omitempty"`
	// Origin
	//
	// x-displayName: "Origin"
	// x-required
	// origin of the signup, from which platform signup is originated, example f5xc, aws..etc
	Origin schema.SignupOrigin `protobuf:"varint,4,opt,name=origin,proto3,enum=ves.io.schema.SignupOrigin" json:"origin,omitempty"`
	// Tenant Type
	//
	// x-displayName: "Tenant Type"
	// Type of the tenant (FREEMIUM, ENTERPRISE)
	// Freemium type includes Free or individual account signups and
	// Enterprise type includes all tenants signed up for a team account with a team plan or organizational plan
	TenantType schema.TenantType `protobuf:"varint,5,opt,name=tenant_type,json=tenantType,proto3,enum=ves.io.schema.TenantType" json:"tenant_type,omitempty"`
}

func (m *GlobalSpecType) Reset()      { *m = GlobalSpecType{} }
func (*GlobalSpecType) ProtoMessage() {}
func (*GlobalSpecType) Descriptor() ([]byte, []int) {
	return fileDescriptor_28e5e1e4640e9fe5, []int{0}
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

func (m *GlobalSpecType) GetPlanType() schema.PlanType {
	if m != nil {
		return m.PlanType
	}
	return schema.FREE
}

func (m *GlobalSpecType) GetCrmInfo() *schema.CRMInfo {
	if m != nil {
		return m.CrmInfo
	}
	return nil
}

func (m *GlobalSpecType) GetAddonServicesSubscribed() []string {
	if m != nil {
		return m.AddonServicesSubscribed
	}
	return nil
}

func (m *GlobalSpecType) GetOrigin() schema.SignupOrigin {
	if m != nil {
		return m.Origin
	}
	return schema.ORIGIN_UNKNOWN
}

func (m *GlobalSpecType) GetTenantType() schema.TenantType {
	if m != nil {
		return m.TenantType
	}
	return schema.UNKNOWN
}

func init() {
	proto.RegisterType((*GlobalSpecType)(nil), "ves.io.schema.saas_config.GlobalSpecType")
	golang_proto.RegisterType((*GlobalSpecType)(nil), "ves.io.schema.saas_config.GlobalSpecType")
}

func init() {
	proto.RegisterFile("ves.io/schema/saas_config/types.proto", fileDescriptor_28e5e1e4640e9fe5)
}
func init() {
	golang_proto.RegisterFile("ves.io/schema/saas_config/types.proto", fileDescriptor_28e5e1e4640e9fe5)
}

var fileDescriptor_28e5e1e4640e9fe5 = []byte{
	// 463 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x92, 0x3f, 0x6f, 0xd3, 0x40,
	0x18, 0xc6, 0x7d, 0x0d, 0x94, 0xd6, 0x91, 0x2a, 0xe4, 0x81, 0x26, 0x69, 0x75, 0x8a, 0x2a, 0x21,
	0x65, 0x89, 0x4f, 0xb4, 0x4c, 0x8c, 0x65, 0xa8, 0x18, 0x50, 0x51, 0xd2, 0x89, 0xc5, 0x3a, 0xdb,
	0xe7, 0xeb, 0x09, 0xfb, 0xde, 0xd3, 0xdd, 0xd9, 0xa2, 0x5b, 0xbf, 0x00, 0x12, 0xe2, 0x53, 0xb0,
	0xb3, 0x31, 0x75, 0x64, 0xcc, 0x98, 0x91, 0x38, 0x0b, 0x6c, 0xfd, 0x08, 0x28, 0x17, 0xb7, 0x8a,
	0xc3, 0xf6, 0x9c, 0x9e, 0xdf, 0xf3, 0xfe, 0xb1, 0x5f, 0xff, 0x65, 0xc5, 0x4c, 0x28, 0x80, 0x98,
	0xe4, 0x9a, 0x15, 0x94, 0x18, 0x4a, 0x4d, 0x94, 0x80, 0xcc, 0x04, 0x27, 0xf6, 0x46, 0x31, 0x13,
	0x2a, 0x0d, 0x16, 0x82, 0xfe, 0x1a, 0x0b, 0xd7, 0x58, 0xb8, 0x81, 0x0d, 0xc6, 0x5c, 0xd8, 0xeb,
	0x32, 0x0e, 0x13, 0x28, 0x08, 0x07, 0x0e, 0xc4, 0x25, 0xe2, 0x32, 0x73, 0x2f, 0xf7, 0x70, 0x6a,
	0x5d, 0x69, 0x70, 0xcc, 0x01, 0x78, 0xce, 0x08, 0x55, 0x82, 0x50, 0x29, 0xc1, 0x52, 0x2b, 0x40,
	0x36, 0x7d, 0x06, 0x47, 0xed, 0x71, 0x40, 0x6d, 0x9a, 0xfd, 0xb6, 0xb9, 0x31, 0xdf, 0xe0, 0xb8,
	0x6d, 0x55, 0x34, 0x17, 0x29, 0xb5, 0xac, 0x71, 0x4f, 0xb6, 0x5c, 0x66, 0x98, 0xac, 0xda, 0xc5,
	0x4f, 0x7e, 0xec, 0xf8, 0x07, 0x17, 0x39, 0xc4, 0x34, 0x9f, 0x2a, 0x96, 0x5c, 0xdd, 0x28, 0x16,
	0xbc, 0xf6, 0xf7, 0x55, 0x4e, 0x65, 0xb4, 0x6a, 0xd4, 0x43, 0x43, 0x34, 0x3a, 0x38, 0x3d, 0x0c,
	0xdb, 0x1f, 0xe2, 0x43, 0x4e, 0xe5, 0x8a, 0x9d, 0xec, 0xa9, 0x46, 0x05, 0xaf, 0xfc, 0xbd, 0x44,
	0x17, 0x91, 0x90, 0x19, 0xf4, 0x76, 0x86, 0x68, 0xd4, 0x3d, 0x7d, 0xb1, 0x15, 0x7a, 0x3b, 0x79,
	0xff, 0x4e, 0x66, 0x30, 0x79, 0x96, 0xe8, 0x62, 0x25, 0x82, 0x0b, 0xbf, 0x4f, 0xd3, 0x14, 0x64,
	0x64, 0x98, 0xae, 0x44, 0xc2, 0x4c, 0x64, 0xca, 0xd8, 0x24, 0x5a, 0xc4, 0x2c, 0xed, 0x75, 0x86,
	0x9d, 0xd1, 0xfe, 0x79, 0xf7, 0xe7, 0xdf, 0xbb, 0xce, 0xee, 0x37, 0xd4, 0x79, 0x7e, 0x8b, 0x26,
	0x87, 0x8e, 0x9e, 0x36, 0xf0, 0xf4, 0x91, 0x0d, 0xce, 0xfc, 0x5d, 0xd0, 0x82, 0x0b, 0xd9, 0x7b,
	0xe2, 0xc6, 0x3d, 0xda, 0xea, 0x3c, 0x15, 0x5c, 0x96, 0xea, 0xd2, 0x21, 0x93, 0x06, 0x0d, 0xde,
	0xf8, 0x5d, 0xcb, 0x24, 0x95, 0x76, 0xbd, 0xe8, 0x53, 0x97, 0xec, 0x6f, 0x25, 0xaf, 0x1c, 0xe1,
	0x56, 0xf5, 0xed, 0xa3, 0x3e, 0xff, 0x82, 0x66, 0x0b, 0xec, 0xcd, 0x17, 0xd8, 0xbb, 0x5f, 0x60,
	0x74, 0x5b, 0x63, 0xf4, 0xbd, 0xc6, 0xe8, 0x57, 0x8d, 0xd1, 0xac, 0xc6, 0x68, 0x5e, 0x63, 0xf4,
	0xbb, 0xc6, 0xe8, 0x4f, 0x8d, 0xbd, 0xfb, 0x1a, 0xa3, 0xaf, 0x4b, 0xec, 0xdd, 0x2d, 0x31, 0x9a,
	0x2d, 0xb1, 0x37, 0x5f, 0x62, 0xef, 0xe3, 0x25, 0x07, 0xf5, 0x89, 0x87, 0x15, 0xe4, 0x96, 0x69,
	0x4d, 0xc3, 0xd2, 0x10, 0x27, 0x32, 0xd0, 0xc5, 0x58, 0x69, 0xa8, 0x44, 0xca, 0xf4, 0xf8, 0xc1,
	0x26, 0x2a, 0xe6, 0x40, 0xd8, 0x67, 0xfb, 0x70, 0xac, 0xff, 0xdd, 0x6c, 0xbc, 0xeb, 0x7e, 0xe6,
	0xd9, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa3, 0x13, 0xf3, 0x61, 0xd7, 0x02, 0x00, 0x00,
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
	if this.PlanType != that1.PlanType {
		return false
	}
	if !this.CrmInfo.Equal(that1.CrmInfo) {
		return false
	}
	if len(this.AddonServicesSubscribed) != len(that1.AddonServicesSubscribed) {
		return false
	}
	for i := range this.AddonServicesSubscribed {
		if this.AddonServicesSubscribed[i] != that1.AddonServicesSubscribed[i] {
			return false
		}
	}
	if this.Origin != that1.Origin {
		return false
	}
	if this.TenantType != that1.TenantType {
		return false
	}
	return true
}
func (this *GlobalSpecType) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 9)
	s = append(s, "&saas_config.GlobalSpecType{")
	s = append(s, "PlanType: "+fmt.Sprintf("%#v", this.PlanType)+",\n")
	if this.CrmInfo != nil {
		s = append(s, "CrmInfo: "+fmt.Sprintf("%#v", this.CrmInfo)+",\n")
	}
	s = append(s, "AddonServicesSubscribed: "+fmt.Sprintf("%#v", this.AddonServicesSubscribed)+",\n")
	s = append(s, "Origin: "+fmt.Sprintf("%#v", this.Origin)+",\n")
	s = append(s, "TenantType: "+fmt.Sprintf("%#v", this.TenantType)+",\n")
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
	if m.TenantType != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.TenantType))
		i--
		dAtA[i] = 0x28
	}
	if m.Origin != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.Origin))
		i--
		dAtA[i] = 0x20
	}
	if len(m.AddonServicesSubscribed) > 0 {
		for iNdEx := len(m.AddonServicesSubscribed) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.AddonServicesSubscribed[iNdEx])
			copy(dAtA[i:], m.AddonServicesSubscribed[iNdEx])
			i = encodeVarintTypes(dAtA, i, uint64(len(m.AddonServicesSubscribed[iNdEx])))
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.CrmInfo != nil {
		{
			size, err := m.CrmInfo.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.PlanType != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.PlanType))
		i--
		dAtA[i] = 0x8
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
	if m.PlanType != 0 {
		n += 1 + sovTypes(uint64(m.PlanType))
	}
	if m.CrmInfo != nil {
		l = m.CrmInfo.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	if len(m.AddonServicesSubscribed) > 0 {
		for _, s := range m.AddonServicesSubscribed {
			l = len(s)
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	if m.Origin != 0 {
		n += 1 + sovTypes(uint64(m.Origin))
	}
	if m.TenantType != 0 {
		n += 1 + sovTypes(uint64(m.TenantType))
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
	s := strings.Join([]string{`&GlobalSpecType{`,
		`PlanType:` + fmt.Sprintf("%v", this.PlanType) + `,`,
		`CrmInfo:` + strings.Replace(fmt.Sprintf("%v", this.CrmInfo), "CRMInfo", "schema.CRMInfo", 1) + `,`,
		`AddonServicesSubscribed:` + fmt.Sprintf("%v", this.AddonServicesSubscribed) + `,`,
		`Origin:` + fmt.Sprintf("%v", this.Origin) + `,`,
		`TenantType:` + fmt.Sprintf("%v", this.TenantType) + `,`,
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
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PlanType", wireType)
			}
			m.PlanType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PlanType |= schema.PlanType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CrmInfo", wireType)
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
			if m.CrmInfo == nil {
				m.CrmInfo = &schema.CRMInfo{}
			}
			if err := m.CrmInfo.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AddonServicesSubscribed", wireType)
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
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AddonServicesSubscribed = append(m.AddonServicesSubscribed, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Origin", wireType)
			}
			m.Origin = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Origin |= schema.SignupOrigin(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TenantType", wireType)
			}
			m.TenantType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TenantType |= schema.TenantType(b&0x7F) << shift
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
