// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ves.io/schema/ai_assistant/gen_dashboard_filter/gen_dashboard_filter.proto

// Generate Dashboard Filter
//
// x-displayName: "Generate Dashboard Filter"
// AI Assistant generates dashboard filter response

package gen_dashboard_filter

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema"
	common "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/ai_assistant/common"
	io "io"
	math "math"
	math_bits "math/bits"
	reflect "reflect"
	strings "strings"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// GenDashboardFilterResponse
//
// x-displayName: "Generate Dashboard Filter"
// Generate dashboard filter response
type GenDashboardFilterResponse struct {
	// summary
	//
	// x-example: "135 events matching your query were found"
	// x-displayName: "Summary"
	// Events summary
	Summary string `protobuf:"bytes,1,opt,name=summary,proto3" json:"summary,omitempty"`
	// actions
	//
	// x-example: "To view these events go to the ${{link:dashboard_link}} page"
	// x-displayName: "Actions"
	// Follow-up actions
	Actions string `protobuf:"bytes,2,opt,name=actions,proto3" json:"actions,omitempty"`
	// links
	//
	// x-displayName: "Links"
	// Links that will be presented to the user
	Links []*common.Link `protobuf:"bytes,3,rep,name=links,proto3" json:"links,omitempty"`
}

func (m *GenDashboardFilterResponse) Reset()      { *m = GenDashboardFilterResponse{} }
func (*GenDashboardFilterResponse) ProtoMessage() {}
func (*GenDashboardFilterResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4d7f3e2e0a806465, []int{0}
}
func (m *GenDashboardFilterResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenDashboardFilterResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenDashboardFilterResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenDashboardFilterResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenDashboardFilterResponse.Merge(m, src)
}
func (m *GenDashboardFilterResponse) XXX_Size() int {
	return m.Size()
}
func (m *GenDashboardFilterResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GenDashboardFilterResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GenDashboardFilterResponse proto.InternalMessageInfo

func (m *GenDashboardFilterResponse) GetSummary() string {
	if m != nil {
		return m.Summary
	}
	return ""
}

func (m *GenDashboardFilterResponse) GetActions() string {
	if m != nil {
		return m.Actions
	}
	return ""
}

func (m *GenDashboardFilterResponse) GetLinks() []*common.Link {
	if m != nil {
		return m.Links
	}
	return nil
}

func init() {
	proto.RegisterType((*GenDashboardFilterResponse)(nil), "ves.io.schema.ai_assistant.gen_dashboard_filter.GenDashboardFilterResponse")
}

func init() {
	proto.RegisterFile("ves.io/schema/ai_assistant/gen_dashboard_filter/gen_dashboard_filter.proto", fileDescriptor_4d7f3e2e0a806465)
}

var fileDescriptor_4d7f3e2e0a806465 = []byte{
	// 347 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xb1, 0x4e, 0xeb, 0x30,
	0x14, 0x86, 0xe3, 0x5b, 0xdd, 0x7b, 0x45, 0xd8, 0x3a, 0x45, 0x45, 0xb2, 0x2a, 0x16, 0x3a, 0x50,
	0x5b, 0x82, 0x99, 0x05, 0x21, 0x90, 0x10, 0x53, 0x47, 0x96, 0xc8, 0x49, 0x5d, 0xd7, 0x6a, 0xed,
	0x13, 0xd9, 0x4e, 0x80, 0x8d, 0x37, 0x80, 0x9d, 0x17, 0xe0, 0x51, 0x18, 0x3b, 0x76, 0xa4, 0xee,
	0xc2, 0xd8, 0x47, 0x40, 0x75, 0x52, 0x89, 0x4a, 0x55, 0xc5, 0x76, 0xfe, 0xfc, 0xe7, 0x7c, 0xfa,
	0xf3, 0xcb, 0xf1, 0x6d, 0xc5, 0x2d, 0x91, 0x40, 0x6d, 0x3e, 0xe6, 0x8a, 0x51, 0x26, 0x53, 0x66,
	0xad, 0xb4, 0x8e, 0x69, 0x47, 0x05, 0xd7, 0xe9, 0x90, 0xd9, 0x71, 0x06, 0xcc, 0x0c, 0xd3, 0x91,
	0x9c, 0x3a, 0x6e, 0x76, 0x7e, 0x24, 0x85, 0x01, 0x07, 0x6d, 0x5a, 0xb3, 0x48, 0xcd, 0x22, 0x3f,
	0x59, 0x64, 0xd7, 0x59, 0xa7, 0x2f, 0xa4, 0x1b, 0x97, 0x19, 0xc9, 0x41, 0x51, 0x01, 0x02, 0x68,
	0xe0, 0x64, 0xe5, 0x28, 0xa8, 0x20, 0xc2, 0x54, 0xf3, 0x3b, 0xa7, 0x7b, 0xb2, 0xe6, 0xa0, 0x14,
	0x68, 0x3a, 0x95, 0x7a, 0xd2, 0x6c, 0x1f, 0x6d, 0x6f, 0x43, 0xe1, 0x24, 0x68, 0xdb, 0x98, 0xdd,
	0x6d, 0xb3, 0x92, 0xfc, 0x21, 0xdd, 0xda, 0x38, 0x7e, 0x41, 0x71, 0xe7, 0x86, 0xeb, 0xab, 0x4d,
	0xe6, 0xeb, 0x10, 0x79, 0xc0, 0x6d, 0x01, 0xda, 0xf2, 0x76, 0x12, 0xff, 0xb7, 0xa5, 0x52, 0xcc,
	0x3c, 0x25, 0xa8, 0x8b, 0x7a, 0x07, 0x83, 0x8d, 0x5c, 0x3b, 0x2c, 0x0f, 0xa4, 0xe4, 0x4f, 0xed,
	0x34, 0xb2, 0x7d, 0x11, 0xff, 0x5d, 0xe7, 0xb3, 0x49, 0xab, 0xdb, 0xea, 0x1d, 0x9e, 0x9d, 0x90,
	0x3d, 0x7d, 0xd5, 0xff, 0x43, 0xee, 0xa4, 0x9e, 0x0c, 0xea, 0xab, 0xcb, 0x37, 0x34, 0x5b, 0xe0,
	0x68, 0xbe, 0xc0, 0xd1, 0x6a, 0x81, 0xd1, 0xb3, 0xc7, 0xe8, 0xdd, 0x63, 0xf4, 0xe1, 0x31, 0x9a,
	0x79, 0x8c, 0x3e, 0x3d, 0x46, 0x5f, 0x1e, 0x47, 0x2b, 0x8f, 0xd1, 0xeb, 0x12, 0x47, 0xb3, 0x25,
	0x8e, 0xe6, 0x4b, 0x1c, 0xdd, 0x73, 0x01, 0xc5, 0x44, 0x90, 0x0a, 0xd6, 0xf9, 0x0d, 0x23, 0xa5,
	0xa5, 0x61, 0x18, 0x81, 0x51, 0xfd, 0xc2, 0x40, 0x25, 0x87, 0xdc, 0xf4, 0x37, 0x36, 0x2d, 0x32,
	0x01, 0x94, 0x3f, 0xba, 0xa6, 0x99, 0xdf, 0xbe, 0x8b, 0xec, 0x5f, 0xa8, 0xed, 0xfc, 0x3b, 0x00,
	0x00, 0xff, 0xff, 0x46, 0xd3, 0x46, 0x17, 0x51, 0x02, 0x00, 0x00,
}

func (this *GenDashboardFilterResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*GenDashboardFilterResponse)
	if !ok {
		that2, ok := that.(GenDashboardFilterResponse)
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
	if this.Summary != that1.Summary {
		return false
	}
	if this.Actions != that1.Actions {
		return false
	}
	if len(this.Links) != len(that1.Links) {
		return false
	}
	for i := range this.Links {
		if !this.Links[i].Equal(that1.Links[i]) {
			return false
		}
	}
	return true
}
func (this *GenDashboardFilterResponse) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 7)
	s = append(s, "&gen_dashboard_filter.GenDashboardFilterResponse{")
	s = append(s, "Summary: "+fmt.Sprintf("%#v", this.Summary)+",\n")
	s = append(s, "Actions: "+fmt.Sprintf("%#v", this.Actions)+",\n")
	if this.Links != nil {
		s = append(s, "Links: "+fmt.Sprintf("%#v", this.Links)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringGenDashboardFilter(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *GenDashboardFilterResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenDashboardFilterResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenDashboardFilterResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Links) > 0 {
		for iNdEx := len(m.Links) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Links[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenDashboardFilter(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Actions) > 0 {
		i -= len(m.Actions)
		copy(dAtA[i:], m.Actions)
		i = encodeVarintGenDashboardFilter(dAtA, i, uint64(len(m.Actions)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Summary) > 0 {
		i -= len(m.Summary)
		copy(dAtA[i:], m.Summary)
		i = encodeVarintGenDashboardFilter(dAtA, i, uint64(len(m.Summary)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintGenDashboardFilter(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenDashboardFilter(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenDashboardFilterResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Summary)
	if l > 0 {
		n += 1 + l + sovGenDashboardFilter(uint64(l))
	}
	l = len(m.Actions)
	if l > 0 {
		n += 1 + l + sovGenDashboardFilter(uint64(l))
	}
	if len(m.Links) > 0 {
		for _, e := range m.Links {
			l = e.Size()
			n += 1 + l + sovGenDashboardFilter(uint64(l))
		}
	}
	return n
}

func sovGenDashboardFilter(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenDashboardFilter(x uint64) (n int) {
	return sovGenDashboardFilter(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *GenDashboardFilterResponse) String() string {
	if this == nil {
		return "nil"
	}
	repeatedStringForLinks := "[]*Link{"
	for _, f := range this.Links {
		repeatedStringForLinks += strings.Replace(fmt.Sprintf("%v", f), "Link", "common.Link", 1) + ","
	}
	repeatedStringForLinks += "}"
	s := strings.Join([]string{`&GenDashboardFilterResponse{`,
		`Summary:` + fmt.Sprintf("%v", this.Summary) + `,`,
		`Actions:` + fmt.Sprintf("%v", this.Actions) + `,`,
		`Links:` + repeatedStringForLinks + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringGenDashboardFilter(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *GenDashboardFilterResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenDashboardFilter
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
			return fmt.Errorf("proto: GenDashboardFilterResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenDashboardFilterResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Summary", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenDashboardFilter
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
				return ErrInvalidLengthGenDashboardFilter
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenDashboardFilter
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Summary = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Actions", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenDashboardFilter
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
				return ErrInvalidLengthGenDashboardFilter
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenDashboardFilter
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Actions = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Links", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenDashboardFilter
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
				return ErrInvalidLengthGenDashboardFilter
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenDashboardFilter
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Links = append(m.Links, &common.Link{})
			if err := m.Links[len(m.Links)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenDashboardFilter(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenDashboardFilter
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthGenDashboardFilter
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
func skipGenDashboardFilter(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenDashboardFilter
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
					return 0, ErrIntOverflowGenDashboardFilter
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
					return 0, ErrIntOverflowGenDashboardFilter
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
				return 0, ErrInvalidLengthGenDashboardFilter
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenDashboardFilter
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenDashboardFilter
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenDashboardFilter        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenDashboardFilter          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenDashboardFilter = fmt.Errorf("proto: unexpected end of group")
)
