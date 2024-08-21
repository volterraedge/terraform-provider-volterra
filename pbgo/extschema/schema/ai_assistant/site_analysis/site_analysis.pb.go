// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ves.io/schema/ai_assistant/site_analysis/site_analysis.proto

// site_analysis
//
// x-displayName: "Site Analysis"
// AI Assistant site analysis response

package site_analysis

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema"
	common "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/ai_assistant/common"
	_ "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/policy"
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

// Analysis and Action
//
// x-displayName: "Analysis and Action"
// Analysis and Action
type AnalysisAndAction struct {
	// analysis
	//
	// x-displayName: "Analysis"
	// x-example: "siteA is in provisioning state due to connectivity issues"
	Analysis string `protobuf:"bytes,1,opt,name=analysis,proto3" json:"analysis,omitempty"`
	// action
	//
	// x-displayName: "Action"
	// x-example: "Retry provisioning the site"
	Action string `protobuf:"bytes,2,opt,name=action,proto3" json:"action,omitempty"`
}

func (m *AnalysisAndAction) Reset()      { *m = AnalysisAndAction{} }
func (*AnalysisAndAction) ProtoMessage() {}
func (*AnalysisAndAction) Descriptor() ([]byte, []int) {
	return fileDescriptor_15ed2a11dbdd73e8, []int{0}
}
func (m *AnalysisAndAction) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AnalysisAndAction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AnalysisAndAction.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AnalysisAndAction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AnalysisAndAction.Merge(m, src)
}
func (m *AnalysisAndAction) XXX_Size() int {
	return m.Size()
}
func (m *AnalysisAndAction) XXX_DiscardUnknown() {
	xxx_messageInfo_AnalysisAndAction.DiscardUnknown(m)
}

var xxx_messageInfo_AnalysisAndAction proto.InternalMessageInfo

func (m *AnalysisAndAction) GetAnalysis() string {
	if m != nil {
		return m.Analysis
	}
	return ""
}

func (m *AnalysisAndAction) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

// Site Analysis Response
//
// x-displayName: "Site Analysis Response"
// Site analysis response
type SiteAnalysisResponse struct {
	// summary
	//
	// x-displayName: "Summary"
	// x-example: "This site analysis response provides status of sites."
	Summary string `protobuf:"bytes,1,opt,name=summary,proto3" json:"summary,omitempty"`
	// analysis and action list
	//
	// x-example: "The site failed because it was stuck in provisioning state. Retry provisioning the site"
	// x-displayName: "Analysis and Action list"
	// Site Status Analysis and Actions list
	AnalysisAndActions []*AnalysisAndAction `protobuf:"bytes,2,rep,name=analysis_and_actions,json=analysisAndActions,proto3" json:"analysis_and_actions,omitempty"`
	// external doc link
	//
	// x-displayName: "External Link"
	// External doc link, that will be presented to the user
	ExternalLink *common.Link `protobuf:"bytes,3,opt,name=external_link,json=externalLink,proto3" json:"external_link,omitempty"`
	// internal link
	//
	// x-displayName: "Internal Link"
	// Internal Link like dashboard link, that will be presented to the user
	InternalLink *common.Link `protobuf:"bytes,4,opt,name=internal_link,json=internalLink,proto3" json:"internal_link,omitempty"`
}

func (m *SiteAnalysisResponse) Reset()      { *m = SiteAnalysisResponse{} }
func (*SiteAnalysisResponse) ProtoMessage() {}
func (*SiteAnalysisResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_15ed2a11dbdd73e8, []int{1}
}
func (m *SiteAnalysisResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SiteAnalysisResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SiteAnalysisResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SiteAnalysisResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SiteAnalysisResponse.Merge(m, src)
}
func (m *SiteAnalysisResponse) XXX_Size() int {
	return m.Size()
}
func (m *SiteAnalysisResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SiteAnalysisResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SiteAnalysisResponse proto.InternalMessageInfo

func (m *SiteAnalysisResponse) GetSummary() string {
	if m != nil {
		return m.Summary
	}
	return ""
}

func (m *SiteAnalysisResponse) GetAnalysisAndActions() []*AnalysisAndAction {
	if m != nil {
		return m.AnalysisAndActions
	}
	return nil
}

func (m *SiteAnalysisResponse) GetExternalLink() *common.Link {
	if m != nil {
		return m.ExternalLink
	}
	return nil
}

func (m *SiteAnalysisResponse) GetInternalLink() *common.Link {
	if m != nil {
		return m.InternalLink
	}
	return nil
}

func init() {
	proto.RegisterType((*AnalysisAndAction)(nil), "ves.io.schema.ai_assistant.site_analysis.AnalysisAndAction")
	proto.RegisterType((*SiteAnalysisResponse)(nil), "ves.io.schema.ai_assistant.site_analysis.SiteAnalysisResponse")
}

func init() {
	proto.RegisterFile("ves.io/schema/ai_assistant/site_analysis/site_analysis.proto", fileDescriptor_15ed2a11dbdd73e8)
}

var fileDescriptor_15ed2a11dbdd73e8 = []byte{
	// 453 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0xb1, 0x8e, 0xd3, 0x30,
	0x18, 0xc7, 0xe3, 0x36, 0x1c, 0x25, 0x07, 0xd2, 0x11, 0x9d, 0x50, 0x54, 0x90, 0x55, 0xdd, 0x42,
	0x07, 0x6a, 0x4b, 0xc7, 0x08, 0x4b, 0x19, 0x4f, 0x37, 0xf5, 0x36, 0x06, 0x22, 0x37, 0xf5, 0xe5,
	0xac, 0x26, 0xfe, 0x22, 0xdb, 0x0d, 0xd7, 0x8d, 0x47, 0x60, 0xe0, 0x21, 0x78, 0x07, 0x5e, 0x80,
	0xb1, 0xe3, 0x8d, 0x34, 0x5d, 0x18, 0xef, 0x11, 0x50, 0x9c, 0x18, 0x5d, 0x00, 0x15, 0xd8, 0xfc,
	0xcf, 0xf7, 0xff, 0xff, 0x3e, 0x7f, 0xfa, 0x9c, 0xe0, 0x75, 0xc9, 0x35, 0x11, 0x40, 0x75, 0x72,
	0xc5, 0x73, 0x46, 0x99, 0x88, 0x99, 0xd6, 0x42, 0x1b, 0x26, 0x0d, 0xd5, 0xc2, 0xf0, 0x98, 0x49,
	0x96, 0xad, 0xb5, 0xd0, 0x5d, 0x45, 0x0a, 0x05, 0x06, 0xc2, 0x71, 0x93, 0x26, 0x4d, 0x9a, 0xdc,
	0x4d, 0x93, 0x8e, 0x7f, 0x38, 0x49, 0x85, 0xb9, 0x5a, 0xcd, 0x49, 0x02, 0x39, 0x4d, 0x21, 0x05,
	0x6a, 0x01, 0xf3, 0xd5, 0xa5, 0x55, 0x56, 0xd8, 0x53, 0x03, 0x1e, 0xbe, 0xd8, 0x73, 0xad, 0x04,
	0xf2, 0x1c, 0x24, 0xcd, 0x84, 0x5c, 0xb6, 0xee, 0xa7, 0x5d, 0x37, 0x14, 0x46, 0x80, 0x6c, 0xef,
	0x38, 0x1c, 0x75, 0x8b, 0x05, 0x64, 0x22, 0x59, 0x53, 0xb3, 0x2e, 0xb8, 0x73, 0x3c, 0xeb, 0x3a,
	0x4a, 0x96, 0x89, 0x05, 0x33, 0xfc, 0xcf, 0xf9, 0x52, 0xf0, 0xf7, 0x71, 0xa7, 0xc3, 0xc9, 0x45,
	0xf0, 0x78, 0xda, 0xce, 0x39, 0x95, 0x8b, 0x69, 0x52, 0xd7, 0xc2, 0x61, 0x30, 0x70, 0xc3, 0x47,
	0x68, 0x84, 0xc6, 0x0f, 0x66, 0x3f, 0x75, 0xf8, 0x24, 0x38, 0x60, 0xd6, 0x15, 0xf5, 0x6c, 0xa5,
	0x55, 0x67, 0xfe, 0xa0, 0x7f, 0xe4, 0x9f, 0xf9, 0x03, 0xff, 0xe8, 0xde, 0xc9, 0x97, 0x5e, 0x70,
	0x7c, 0x21, 0x0c, 0x77, 0xe4, 0x19, 0xd7, 0x05, 0x48, 0xcd, 0xc3, 0x28, 0xb8, 0xaf, 0x57, 0x79,
	0xce, 0xd4, 0xba, 0xe5, 0x3a, 0x19, 0xe6, 0xc1, 0xb1, 0x6b, 0x11, 0x33, 0xb9, 0x88, 0x1b, 0xaa,
	0x8e, 0x7a, 0xa3, 0xfe, 0xf8, 0xf0, 0xf4, 0x15, 0xf9, 0xd7, 0x65, 0x91, 0xdf, 0xa6, 0x99, 0x85,
	0xec, 0xd7, 0x4f, 0x3a, 0x3c, 0x0f, 0x1e, 0xf1, 0x6b, 0xc3, 0x95, 0x64, 0x59, 0x5c, 0x2f, 0x23,
	0xea, 0x8f, 0xd0, 0xf8, 0xf0, 0xf4, 0xf9, 0xbe, 0x3e, 0xcd, 0xee, 0xc8, 0xb9, 0x90, 0xcb, 0xd9,
	0x43, 0x97, 0xae, 0x55, 0x4d, 0x13, 0xf2, 0x2e, 0xcd, 0xff, 0x4f, 0x9a, 0x4b, 0xd7, 0xea, 0xcd,
	0x27, 0xb4, 0xd9, 0x62, 0xef, 0x66, 0x8b, 0xbd, 0xdb, 0x2d, 0x46, 0x1f, 0x2a, 0x8c, 0x3e, 0x57,
	0x18, 0x7d, 0xad, 0x30, 0xda, 0x54, 0x18, 0x7d, 0xab, 0x30, 0xfa, 0x5e, 0x61, 0xef, 0xb6, 0xc2,
	0xe8, 0xe3, 0x0e, 0x7b, 0x9b, 0x1d, 0xf6, 0x6e, 0x76, 0xd8, 0x7b, 0xfb, 0x2e, 0x85, 0x62, 0x99,
	0x92, 0x12, 0x32, 0xc3, 0x95, 0x62, 0x64, 0xa5, 0xa9, 0x3d, 0x5c, 0x82, 0xca, 0x27, 0x85, 0x82,
	0x52, 0x2c, 0xb8, 0x9a, 0xb8, 0x32, 0x2d, 0xe6, 0x29, 0x50, 0x7e, 0x6d, 0xda, 0xa7, 0xf1, 0xd7,
	0x7f, 0x68, 0x7e, 0x60, 0x1f, 0xcc, 0xcb, 0x1f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x0b, 0xea, 0xba,
	0x17, 0x76, 0x03, 0x00, 0x00,
}

func (this *AnalysisAndAction) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*AnalysisAndAction)
	if !ok {
		that2, ok := that.(AnalysisAndAction)
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
	if this.Analysis != that1.Analysis {
		return false
	}
	if this.Action != that1.Action {
		return false
	}
	return true
}
func (this *SiteAnalysisResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*SiteAnalysisResponse)
	if !ok {
		that2, ok := that.(SiteAnalysisResponse)
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
	if len(this.AnalysisAndActions) != len(that1.AnalysisAndActions) {
		return false
	}
	for i := range this.AnalysisAndActions {
		if !this.AnalysisAndActions[i].Equal(that1.AnalysisAndActions[i]) {
			return false
		}
	}
	if !this.ExternalLink.Equal(that1.ExternalLink) {
		return false
	}
	if !this.InternalLink.Equal(that1.InternalLink) {
		return false
	}
	return true
}
func (this *AnalysisAndAction) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&site_analysis.AnalysisAndAction{")
	s = append(s, "Analysis: "+fmt.Sprintf("%#v", this.Analysis)+",\n")
	s = append(s, "Action: "+fmt.Sprintf("%#v", this.Action)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *SiteAnalysisResponse) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 8)
	s = append(s, "&site_analysis.SiteAnalysisResponse{")
	s = append(s, "Summary: "+fmt.Sprintf("%#v", this.Summary)+",\n")
	if this.AnalysisAndActions != nil {
		s = append(s, "AnalysisAndActions: "+fmt.Sprintf("%#v", this.AnalysisAndActions)+",\n")
	}
	if this.ExternalLink != nil {
		s = append(s, "ExternalLink: "+fmt.Sprintf("%#v", this.ExternalLink)+",\n")
	}
	if this.InternalLink != nil {
		s = append(s, "InternalLink: "+fmt.Sprintf("%#v", this.InternalLink)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringSiteAnalysis(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *AnalysisAndAction) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AnalysisAndAction) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AnalysisAndAction) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Action) > 0 {
		i -= len(m.Action)
		copy(dAtA[i:], m.Action)
		i = encodeVarintSiteAnalysis(dAtA, i, uint64(len(m.Action)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Analysis) > 0 {
		i -= len(m.Analysis)
		copy(dAtA[i:], m.Analysis)
		i = encodeVarintSiteAnalysis(dAtA, i, uint64(len(m.Analysis)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *SiteAnalysisResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SiteAnalysisResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SiteAnalysisResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.InternalLink != nil {
		{
			size, err := m.InternalLink.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintSiteAnalysis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if m.ExternalLink != nil {
		{
			size, err := m.ExternalLink.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintSiteAnalysis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if len(m.AnalysisAndActions) > 0 {
		for iNdEx := len(m.AnalysisAndActions) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.AnalysisAndActions[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintSiteAnalysis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Summary) > 0 {
		i -= len(m.Summary)
		copy(dAtA[i:], m.Summary)
		i = encodeVarintSiteAnalysis(dAtA, i, uint64(len(m.Summary)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintSiteAnalysis(dAtA []byte, offset int, v uint64) int {
	offset -= sovSiteAnalysis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *AnalysisAndAction) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Analysis)
	if l > 0 {
		n += 1 + l + sovSiteAnalysis(uint64(l))
	}
	l = len(m.Action)
	if l > 0 {
		n += 1 + l + sovSiteAnalysis(uint64(l))
	}
	return n
}

func (m *SiteAnalysisResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Summary)
	if l > 0 {
		n += 1 + l + sovSiteAnalysis(uint64(l))
	}
	if len(m.AnalysisAndActions) > 0 {
		for _, e := range m.AnalysisAndActions {
			l = e.Size()
			n += 1 + l + sovSiteAnalysis(uint64(l))
		}
	}
	if m.ExternalLink != nil {
		l = m.ExternalLink.Size()
		n += 1 + l + sovSiteAnalysis(uint64(l))
	}
	if m.InternalLink != nil {
		l = m.InternalLink.Size()
		n += 1 + l + sovSiteAnalysis(uint64(l))
	}
	return n
}

func sovSiteAnalysis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozSiteAnalysis(x uint64) (n int) {
	return sovSiteAnalysis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *AnalysisAndAction) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&AnalysisAndAction{`,
		`Analysis:` + fmt.Sprintf("%v", this.Analysis) + `,`,
		`Action:` + fmt.Sprintf("%v", this.Action) + `,`,
		`}`,
	}, "")
	return s
}
func (this *SiteAnalysisResponse) String() string {
	if this == nil {
		return "nil"
	}
	repeatedStringForAnalysisAndActions := "[]*AnalysisAndAction{"
	for _, f := range this.AnalysisAndActions {
		repeatedStringForAnalysisAndActions += strings.Replace(f.String(), "AnalysisAndAction", "AnalysisAndAction", 1) + ","
	}
	repeatedStringForAnalysisAndActions += "}"
	s := strings.Join([]string{`&SiteAnalysisResponse{`,
		`Summary:` + fmt.Sprintf("%v", this.Summary) + `,`,
		`AnalysisAndActions:` + repeatedStringForAnalysisAndActions + `,`,
		`ExternalLink:` + strings.Replace(fmt.Sprintf("%v", this.ExternalLink), "Link", "common.Link", 1) + `,`,
		`InternalLink:` + strings.Replace(fmt.Sprintf("%v", this.InternalLink), "Link", "common.Link", 1) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringSiteAnalysis(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *AnalysisAndAction) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSiteAnalysis
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
			return fmt.Errorf("proto: AnalysisAndAction: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AnalysisAndAction: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Analysis", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSiteAnalysis
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
				return ErrInvalidLengthSiteAnalysis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSiteAnalysis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Analysis = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Action", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSiteAnalysis
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
				return ErrInvalidLengthSiteAnalysis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSiteAnalysis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Action = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSiteAnalysis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSiteAnalysis
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthSiteAnalysis
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
func (m *SiteAnalysisResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSiteAnalysis
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
			return fmt.Errorf("proto: SiteAnalysisResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SiteAnalysisResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Summary", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSiteAnalysis
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
				return ErrInvalidLengthSiteAnalysis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSiteAnalysis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Summary = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AnalysisAndActions", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSiteAnalysis
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
				return ErrInvalidLengthSiteAnalysis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSiteAnalysis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AnalysisAndActions = append(m.AnalysisAndActions, &AnalysisAndAction{})
			if err := m.AnalysisAndActions[len(m.AnalysisAndActions)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExternalLink", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSiteAnalysis
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
				return ErrInvalidLengthSiteAnalysis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSiteAnalysis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ExternalLink == nil {
				m.ExternalLink = &common.Link{}
			}
			if err := m.ExternalLink.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InternalLink", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSiteAnalysis
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
				return ErrInvalidLengthSiteAnalysis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSiteAnalysis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.InternalLink == nil {
				m.InternalLink = &common.Link{}
			}
			if err := m.InternalLink.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSiteAnalysis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSiteAnalysis
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthSiteAnalysis
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
func skipSiteAnalysis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSiteAnalysis
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
					return 0, ErrIntOverflowSiteAnalysis
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
					return 0, ErrIntOverflowSiteAnalysis
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
				return 0, ErrInvalidLengthSiteAnalysis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupSiteAnalysis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthSiteAnalysis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthSiteAnalysis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSiteAnalysis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupSiteAnalysis = fmt.Errorf("proto: unexpected end of group")
)