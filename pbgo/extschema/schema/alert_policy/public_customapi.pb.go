// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ves.io/schema/alert_policy/public_customapi.proto

// Alert Policy
//
// x-displayName: "Alert Policy"
// Custom APIs for alert policy object

package alert_policy

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/googleapis/google/api"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_sortkeys "github.com/gogo/protobuf/sortkeys"
	golang_proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	_ "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema"
	_ "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/vesenv"
	io "io"
	math "math"
	math_bits "math/bits"
	reflect "reflect"
	strconv "strconv"
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

// AlertPolicyStatus
//
// x-displayName: "AlertPolicyStatus"
// Indicates if Alert Policy is in Alert Policy Set, PolicyStatus is ACTIVE or INACTIVE.
type AlertPolicyStatus int32

const (
	// x-displayName: "INACTIVE"
	// INACTIVE means Alert Policy is not in Alert Policy Set, and will
	// not be used to send Alerts to Alert Receivers.
	INACTIVE AlertPolicyStatus = 0
	// x-displayName: "ACTIVE"
	// ACTIVE means Alert Policy is in Alert Policy Set, and is used
	// to send out Alert to Alert Receivers.
	ACTIVE AlertPolicyStatus = 1
)

var AlertPolicyStatus_name = map[int32]string{
	0: "INACTIVE",
	1: "ACTIVE",
}

var AlertPolicyStatus_value = map[string]int32{
	"INACTIVE": 0,
	"ACTIVE":   1,
}

func (AlertPolicyStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_54d5737f5424cdab, []int{0}
}

// Alert Policy Match Request
//
// x-displayName: "Alert Policy Match Request"
// Request message for GetAlertPolicyMatch RPC,
// describing alert to match against alert policies
type AlertPolicyMatchRequest struct {
	// namespace
	//
	// x-displayName: "Namespace"
	// x-required
	// x-example: "ns1"
	// The namespace in which the configuration object is present
	Namespace string `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	// labels
	//
	// x-displayName: "Labels"
	// x-required
	// x-example:   "labels": {
	//                 "namespace": "system",
	//                 "tenant" :"tenant2",
	//                 "alertname": "podXcrash",
	//                 "group": "IaaS",
	//                 "severity": "critical"
	//               },
	// Alert labels to find match against alert policies
	// requires tenant and namespace to be defined labels map
	Labels map[string]string `protobuf:"bytes,2,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (m *AlertPolicyMatchRequest) Reset()      { *m = AlertPolicyMatchRequest{} }
func (*AlertPolicyMatchRequest) ProtoMessage() {}
func (*AlertPolicyMatchRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_54d5737f5424cdab, []int{0}
}
func (m *AlertPolicyMatchRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AlertPolicyMatchRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AlertPolicyMatchRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AlertPolicyMatchRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AlertPolicyMatchRequest.Merge(m, src)
}
func (m *AlertPolicyMatchRequest) XXX_Size() int {
	return m.Size()
}
func (m *AlertPolicyMatchRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AlertPolicyMatchRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AlertPolicyMatchRequest proto.InternalMessageInfo

func (m *AlertPolicyMatchRequest) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *AlertPolicyMatchRequest) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

// Alert Policy Match
//
// x-displayName: "Alert Policy Match"
// Alert Policy info that matches AlertPolicyMatchRequest giving
//
//	alert policy name,
//	alert policy activation state
type AlertPolicyMatch struct {
	// policy_name
	//
	// x-displayName: "Policy Name"
	// x-example: samplepolicy
	PolicyName string `protobuf:"bytes,1,opt,name=policy_name,json=policyName,proto3" json:"policy_name,omitempty"`
	// policy_status
	//
	// x-displayName: "Policy Status"
	// x-example: 1, ACTIVE
	PolicyStatus AlertPolicyStatus `protobuf:"varint,2,opt,name=policy_status,json=policyStatus,proto3,enum=ves.io.schema.alert_policy.AlertPolicyStatus" json:"policy_status,omitempty"`
}

func (m *AlertPolicyMatch) Reset()      { *m = AlertPolicyMatch{} }
func (*AlertPolicyMatch) ProtoMessage() {}
func (*AlertPolicyMatch) Descriptor() ([]byte, []int) {
	return fileDescriptor_54d5737f5424cdab, []int{1}
}
func (m *AlertPolicyMatch) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AlertPolicyMatch) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AlertPolicyMatch.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AlertPolicyMatch) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AlertPolicyMatch.Merge(m, src)
}
func (m *AlertPolicyMatch) XXX_Size() int {
	return m.Size()
}
func (m *AlertPolicyMatch) XXX_DiscardUnknown() {
	xxx_messageInfo_AlertPolicyMatch.DiscardUnknown(m)
}

var xxx_messageInfo_AlertPolicyMatch proto.InternalMessageInfo

func (m *AlertPolicyMatch) GetPolicyName() string {
	if m != nil {
		return m.PolicyName
	}
	return ""
}

func (m *AlertPolicyMatch) GetPolicyStatus() AlertPolicyStatus {
	if m != nil {
		return m.PolicyStatus
	}
	return INACTIVE
}

// Alert Policy Match Response
//
// x-displayName: "Alert Policy Match Response"
// Response of matching Alert Policies from Get Alert Policy Match request
type AlertPolicyMatchResponse struct {
	// alert_match
	//
	// x-displayName: "Alert Match"
	// x-example:{"alert_match":[{"policy_name":"policy1","policy_active":1},{"policy_name":"policy4"}]}
	// List of Alert Policies that match given requested namespace,labels.
	AlertMatch []*AlertPolicyMatch `protobuf:"bytes,1,rep,name=alert_match,json=alertMatch,proto3" json:"alert_match,omitempty"`
}

func (m *AlertPolicyMatchResponse) Reset()      { *m = AlertPolicyMatchResponse{} }
func (*AlertPolicyMatchResponse) ProtoMessage() {}
func (*AlertPolicyMatchResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_54d5737f5424cdab, []int{2}
}
func (m *AlertPolicyMatchResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AlertPolicyMatchResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AlertPolicyMatchResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AlertPolicyMatchResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AlertPolicyMatchResponse.Merge(m, src)
}
func (m *AlertPolicyMatchResponse) XXX_Size() int {
	return m.Size()
}
func (m *AlertPolicyMatchResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AlertPolicyMatchResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AlertPolicyMatchResponse proto.InternalMessageInfo

func (m *AlertPolicyMatchResponse) GetAlertMatch() []*AlertPolicyMatch {
	if m != nil {
		return m.AlertMatch
	}
	return nil
}

func init() {
	proto.RegisterEnum("ves.io.schema.alert_policy.AlertPolicyStatus", AlertPolicyStatus_name, AlertPolicyStatus_value)
	golang_proto.RegisterEnum("ves.io.schema.alert_policy.AlertPolicyStatus", AlertPolicyStatus_name, AlertPolicyStatus_value)
	proto.RegisterType((*AlertPolicyMatchRequest)(nil), "ves.io.schema.alert_policy.AlertPolicyMatchRequest")
	golang_proto.RegisterType((*AlertPolicyMatchRequest)(nil), "ves.io.schema.alert_policy.AlertPolicyMatchRequest")
	proto.RegisterMapType((map[string]string)(nil), "ves.io.schema.alert_policy.AlertPolicyMatchRequest.LabelsEntry")
	golang_proto.RegisterMapType((map[string]string)(nil), "ves.io.schema.alert_policy.AlertPolicyMatchRequest.LabelsEntry")
	proto.RegisterType((*AlertPolicyMatch)(nil), "ves.io.schema.alert_policy.AlertPolicyMatch")
	golang_proto.RegisterType((*AlertPolicyMatch)(nil), "ves.io.schema.alert_policy.AlertPolicyMatch")
	proto.RegisterType((*AlertPolicyMatchResponse)(nil), "ves.io.schema.alert_policy.AlertPolicyMatchResponse")
	golang_proto.RegisterType((*AlertPolicyMatchResponse)(nil), "ves.io.schema.alert_policy.AlertPolicyMatchResponse")
}

func init() {
	proto.RegisterFile("ves.io/schema/alert_policy/public_customapi.proto", fileDescriptor_54d5737f5424cdab)
}
func init() {
	golang_proto.RegisterFile("ves.io/schema/alert_policy/public_customapi.proto", fileDescriptor_54d5737f5424cdab)
}

var fileDescriptor_54d5737f5424cdab = []byte{
	// 648 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0xc1, 0x4f, 0xd4, 0x4e,
	0x14, 0xee, 0x74, 0x03, 0x81, 0x59, 0x7e, 0xbf, 0xac, 0xd5, 0xc4, 0xba, 0x92, 0x11, 0xf7, 0x44,
	0x08, 0x6d, 0xc3, 0xa2, 0x46, 0xb9, 0x18, 0x20, 0xc4, 0x90, 0x08, 0x92, 0xd5, 0x78, 0xf0, 0x42,
	0xa6, 0xdd, 0x47, 0xa9, 0xb4, 0x9d, 0xda, 0x99, 0x36, 0x6c, 0x8c, 0x89, 0xe1, 0x62, 0xe2, 0xc9,
	0xc4, 0x9b, 0x7f, 0x81, 0x77, 0x6f, 0x72, 0xe1, 0xa6, 0x27, 0x43, 0xf4, 0x20, 0x47, 0xe9, 0x7a,
	0xd0, 0x1b, 0x07, 0xff, 0x00, 0xb3, 0xd3, 0x82, 0xbb, 0xac, 0x18, 0xf1, 0xf6, 0xde, 0xfb, 0xe6,
	0xfb, 0xe6, 0xbd, 0xaf, 0xf3, 0x8a, 0xa7, 0x52, 0xe0, 0xa6, 0xc7, 0x2c, 0xee, 0xac, 0x43, 0x40,
	0x2d, 0xea, 0x43, 0x2c, 0x56, 0x23, 0xe6, 0x7b, 0x4e, 0xcb, 0x8a, 0x12, 0xdb, 0xf7, 0x9c, 0x55,
	0x27, 0xe1, 0x82, 0x05, 0x34, 0xf2, 0xcc, 0x28, 0x66, 0x82, 0x69, 0xd5, 0x9c, 0x62, 0xe6, 0x14,
	0xb3, 0x9b, 0x52, 0x35, 0x5c, 0x4f, 0xac, 0x27, 0xb6, 0xe9, 0xb0, 0xc0, 0x72, 0x99, 0xcb, 0x2c,
	0x49, 0xb1, 0x93, 0x35, 0x99, 0xc9, 0x44, 0x46, 0xb9, 0x54, 0x75, 0xd4, 0x65, 0xcc, 0xf5, 0xc1,
	0xa2, 0x91, 0x67, 0xd1, 0x30, 0x64, 0x82, 0x0a, 0x8f, 0x85, 0xbc, 0x40, 0x2f, 0xf6, 0xf6, 0xc6,
	0xa2, 0x6e, 0xf0, 0x42, 0x2f, 0x28, 0x5a, 0x11, 0x1c, 0x42, 0xa3, 0xbd, 0x50, 0x4a, 0x7d, 0xaf,
	0x49, 0x05, 0x14, 0x68, 0xed, 0x18, 0x0a, 0x1c, 0xc2, 0xb4, 0x57, 0xbc, 0xf6, 0x19, 0xe1, 0xf3,
	0xb3, 0x9d, 0xb9, 0x56, 0xe4, 0x58, 0x4b, 0x54, 0x38, 0xeb, 0x0d, 0x78, 0x94, 0x00, 0x17, 0xda,
	0x28, 0x1e, 0x0e, 0x69, 0x00, 0x3c, 0xa2, 0x0e, 0xe8, 0x68, 0x0c, 0x8d, 0x0f, 0x37, 0x7e, 0x15,
	0x34, 0x07, 0x0f, 0xfa, 0xd4, 0x06, 0x9f, 0xeb, 0xea, 0x58, 0x69, 0xbc, 0x5c, 0xbf, 0x69, 0x9e,
	0xec, 0x96, 0x79, 0xc2, 0x15, 0xe6, 0x6d, 0xa9, 0xb0, 0x10, 0x8a, 0xb8, 0x35, 0x87, 0xdf, 0x7e,
	0xdf, 0x29, 0x0d, 0xbc, 0x42, 0x6a, 0xa5, 0xde, 0x28, 0xa4, 0xab, 0x37, 0x70, 0xb9, 0xeb, 0x88,
	0x56, 0xc1, 0xa5, 0x0d, 0x68, 0x15, 0xbd, 0x74, 0x42, 0xed, 0x1c, 0x1e, 0x48, 0xa9, 0x9f, 0x80,
	0xae, 0xca, 0x5a, 0x9e, 0xcc, 0xa8, 0xd7, 0x51, 0xed, 0x19, 0xc2, 0x95, 0xe3, 0xd7, 0x6a, 0x97,
	0x70, 0x39, 0xef, 0x68, 0xb5, 0x33, 0x48, 0x21, 0x84, 0xf3, 0xd2, 0x32, 0x0d, 0x40, 0x6b, 0xe0,
	0xff, 0x8a, 0x03, 0x5c, 0x50, 0x91, 0x70, 0xa9, 0xfb, 0x7f, 0xdd, 0xf8, 0xcb, 0xe1, 0xee, 0x4a,
	0x52, 0x63, 0x24, 0xea, 0xca, 0x6a, 0x1e, 0xd6, 0xfb, 0xe7, 0xe7, 0x11, 0x0b, 0x39, 0x68, 0x4b,
	0xb8, 0x9c, 0x6b, 0x05, 0x9d, 0xb2, 0x8e, 0xa4, 0x95, 0x93, 0xa7, 0xb2, 0x12, 0x4b, 0x58, 0xc6,
	0x13, 0x06, 0x3e, 0xd3, 0xd7, 0x8d, 0x36, 0x82, 0x87, 0x16, 0x97, 0x67, 0xe7, 0xef, 0x2d, 0xde,
	0x5f, 0xa8, 0x28, 0x1a, 0xc6, 0x83, 0x45, 0x8c, 0xea, 0x6f, 0x54, 0x3c, 0x3c, 0x2f, 0x1f, 0xfd,
	0xec, 0xca, 0xa2, 0xf6, 0x03, 0xe1, 0xb3, 0xb7, 0x40, 0xf4, 0x99, 0x36, 0xfd, 0x0f, 0x5f, 0xb6,
	0x7a, 0xe5, 0x74, 0xa4, 0xdc, 0x8e, 0x5a, 0x9c, 0xbd, 0xd3, 0xeb, 0x6b, 0x57, 0x37, 0x1d, 0x83,
	0x26, 0x4d, 0x4f, 0x18, 0x3e, 0x73, 0xb9, 0x21, 0xa9, 0xdc, 0x08, 0x58, 0xe8, 0x09, 0x16, 0x4f,
	0xa6, 0xc0, 0x0d, 0x8f, 0x19, 0x2e, 0x84, 0x10, 0x53, 0xdf, 0x88, 0x81, 0x36, 0xb7, 0x3e, 0x7d,
	0x7d, 0xa9, 0x5e, 0xab, 0x4d, 0x15, 0x8b, 0x6c, 0x1d, 0xbd, 0x52, 0x6e, 0x3d, 0x3e, 0x8a, 0x9f,
	0xf4, 0xee, 0xbc, 0x34, 0x7c, 0x06, 0x4d, 0x54, 0xad, 0x9d, 0x6d, 0x54, 0xfa, 0xb8, 0x8d, 0x2e,
	0xff, 0xa1, 0xe1, 0x3b, 0xf6, 0x43, 0x70, 0xc4, 0xd6, 0x07, 0x5d, 0x1d, 0x42, 0x73, 0xcf, 0xd1,
	0xee, 0x3e, 0x51, 0xf6, 0xf6, 0x89, 0x72, 0xb0, 0x4f, 0xd0, 0xd3, 0x8c, 0xa0, 0xd7, 0x19, 0x41,
	0xef, 0x33, 0x82, 0x76, 0x33, 0x82, 0xbe, 0x64, 0x04, 0x7d, 0xcb, 0x88, 0x72, 0x90, 0x11, 0xf4,
	0xa2, 0x4d, 0x94, 0x9d, 0x36, 0x41, 0xbb, 0x6d, 0xa2, 0xec, 0xb5, 0x89, 0xf2, 0x60, 0xc5, 0x65,
	0xd1, 0x86, 0x6b, 0xa6, 0xcc, 0x17, 0x10, 0xc7, 0xd4, 0x4c, 0xb8, 0x25, 0x83, 0x35, 0x16, 0x07,
	0x46, 0x14, 0xb3, 0xd4, 0x6b, 0x42, 0x6c, 0x1c, 0xc2, 0x56, 0x64, 0xbb, 0xcc, 0x82, 0x4d, 0x51,
	0x6c, 0xf0, 0x6f, 0x7e, 0x5d, 0xf6, 0xa0, 0xdc, 0xe3, 0xe9, 0x9f, 0x01, 0x00, 0x00, 0xff, 0xff,
	0x4d, 0x97, 0x1a, 0x21, 0xdf, 0x04, 0x00, 0x00,
}

func (x AlertPolicyStatus) String() string {
	s, ok := AlertPolicyStatus_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (this *AlertPolicyMatchRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*AlertPolicyMatchRequest)
	if !ok {
		that2, ok := that.(AlertPolicyMatchRequest)
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
	if this.Namespace != that1.Namespace {
		return false
	}
	if len(this.Labels) != len(that1.Labels) {
		return false
	}
	for i := range this.Labels {
		if this.Labels[i] != that1.Labels[i] {
			return false
		}
	}
	return true
}
func (this *AlertPolicyMatch) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*AlertPolicyMatch)
	if !ok {
		that2, ok := that.(AlertPolicyMatch)
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
	if this.PolicyName != that1.PolicyName {
		return false
	}
	if this.PolicyStatus != that1.PolicyStatus {
		return false
	}
	return true
}
func (this *AlertPolicyMatchResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*AlertPolicyMatchResponse)
	if !ok {
		that2, ok := that.(AlertPolicyMatchResponse)
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
	if len(this.AlertMatch) != len(that1.AlertMatch) {
		return false
	}
	for i := range this.AlertMatch {
		if !this.AlertMatch[i].Equal(that1.AlertMatch[i]) {
			return false
		}
	}
	return true
}
func (this *AlertPolicyMatchRequest) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&alert_policy.AlertPolicyMatchRequest{")
	s = append(s, "Namespace: "+fmt.Sprintf("%#v", this.Namespace)+",\n")
	keysForLabels := make([]string, 0, len(this.Labels))
	for k, _ := range this.Labels {
		keysForLabels = append(keysForLabels, k)
	}
	github_com_gogo_protobuf_sortkeys.Strings(keysForLabels)
	mapStringForLabels := "map[string]string{"
	for _, k := range keysForLabels {
		mapStringForLabels += fmt.Sprintf("%#v: %#v,", k, this.Labels[k])
	}
	mapStringForLabels += "}"
	if this.Labels != nil {
		s = append(s, "Labels: "+mapStringForLabels+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *AlertPolicyMatch) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&alert_policy.AlertPolicyMatch{")
	s = append(s, "PolicyName: "+fmt.Sprintf("%#v", this.PolicyName)+",\n")
	s = append(s, "PolicyStatus: "+fmt.Sprintf("%#v", this.PolicyStatus)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *AlertPolicyMatchResponse) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&alert_policy.AlertPolicyMatchResponse{")
	if this.AlertMatch != nil {
		s = append(s, "AlertMatch: "+fmt.Sprintf("%#v", this.AlertMatch)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringPublicCustomapi(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CustomAPIClient is the client API for CustomAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CustomAPIClient interface {
	// GetAlertPolicyMatch
	//
	// x-displayName: "Get Alert Policy Match"
	// Get Alert Policies that match to a set of alert labels for a namespace.
	GetAlertPolicyMatch(ctx context.Context, in *AlertPolicyMatchRequest, opts ...grpc.CallOption) (*AlertPolicyMatchResponse, error)
}

type customAPIClient struct {
	cc *grpc.ClientConn
}

func NewCustomAPIClient(cc *grpc.ClientConn) CustomAPIClient {
	return &customAPIClient{cc}
}

func (c *customAPIClient) GetAlertPolicyMatch(ctx context.Context, in *AlertPolicyMatchRequest, opts ...grpc.CallOption) (*AlertPolicyMatchResponse, error) {
	out := new(AlertPolicyMatchResponse)
	err := c.cc.Invoke(ctx, "/ves.io.schema.alert_policy.CustomAPI/GetAlertPolicyMatch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CustomAPIServer is the server API for CustomAPI service.
type CustomAPIServer interface {
	// GetAlertPolicyMatch
	//
	// x-displayName: "Get Alert Policy Match"
	// Get Alert Policies that match to a set of alert labels for a namespace.
	GetAlertPolicyMatch(context.Context, *AlertPolicyMatchRequest) (*AlertPolicyMatchResponse, error)
}

// UnimplementedCustomAPIServer can be embedded to have forward compatible implementations.
type UnimplementedCustomAPIServer struct {
}

func (*UnimplementedCustomAPIServer) GetAlertPolicyMatch(ctx context.Context, req *AlertPolicyMatchRequest) (*AlertPolicyMatchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAlertPolicyMatch not implemented")
}

func RegisterCustomAPIServer(s *grpc.Server, srv CustomAPIServer) {
	s.RegisterService(&_CustomAPI_serviceDesc, srv)
}

func _CustomAPI_GetAlertPolicyMatch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AlertPolicyMatchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomAPIServer).GetAlertPolicyMatch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ves.io.schema.alert_policy.CustomAPI/GetAlertPolicyMatch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomAPIServer).GetAlertPolicyMatch(ctx, req.(*AlertPolicyMatchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CustomAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ves.io.schema.alert_policy.CustomAPI",
	HandlerType: (*CustomAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAlertPolicyMatch",
			Handler:    _CustomAPI_GetAlertPolicyMatch_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ves.io/schema/alert_policy/public_customapi.proto",
}

func (m *AlertPolicyMatchRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AlertPolicyMatchRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AlertPolicyMatchRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Labels) > 0 {
		for k := range m.Labels {
			v := m.Labels[k]
			baseI := i
			i -= len(v)
			copy(dAtA[i:], v)
			i = encodeVarintPublicCustomapi(dAtA, i, uint64(len(v)))
			i--
			dAtA[i] = 0x12
			i -= len(k)
			copy(dAtA[i:], k)
			i = encodeVarintPublicCustomapi(dAtA, i, uint64(len(k)))
			i--
			dAtA[i] = 0xa
			i = encodeVarintPublicCustomapi(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Namespace) > 0 {
		i -= len(m.Namespace)
		copy(dAtA[i:], m.Namespace)
		i = encodeVarintPublicCustomapi(dAtA, i, uint64(len(m.Namespace)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *AlertPolicyMatch) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AlertPolicyMatch) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AlertPolicyMatch) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.PolicyStatus != 0 {
		i = encodeVarintPublicCustomapi(dAtA, i, uint64(m.PolicyStatus))
		i--
		dAtA[i] = 0x10
	}
	if len(m.PolicyName) > 0 {
		i -= len(m.PolicyName)
		copy(dAtA[i:], m.PolicyName)
		i = encodeVarintPublicCustomapi(dAtA, i, uint64(len(m.PolicyName)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *AlertPolicyMatchResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AlertPolicyMatchResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AlertPolicyMatchResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.AlertMatch) > 0 {
		for iNdEx := len(m.AlertMatch) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.AlertMatch[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintPublicCustomapi(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintPublicCustomapi(dAtA []byte, offset int, v uint64) int {
	offset -= sovPublicCustomapi(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *AlertPolicyMatchRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Namespace)
	if l > 0 {
		n += 1 + l + sovPublicCustomapi(uint64(l))
	}
	if len(m.Labels) > 0 {
		for k, v := range m.Labels {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovPublicCustomapi(uint64(len(k))) + 1 + len(v) + sovPublicCustomapi(uint64(len(v)))
			n += mapEntrySize + 1 + sovPublicCustomapi(uint64(mapEntrySize))
		}
	}
	return n
}

func (m *AlertPolicyMatch) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.PolicyName)
	if l > 0 {
		n += 1 + l + sovPublicCustomapi(uint64(l))
	}
	if m.PolicyStatus != 0 {
		n += 1 + sovPublicCustomapi(uint64(m.PolicyStatus))
	}
	return n
}

func (m *AlertPolicyMatchResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.AlertMatch) > 0 {
		for _, e := range m.AlertMatch {
			l = e.Size()
			n += 1 + l + sovPublicCustomapi(uint64(l))
		}
	}
	return n
}

func sovPublicCustomapi(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPublicCustomapi(x uint64) (n int) {
	return sovPublicCustomapi(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *AlertPolicyMatchRequest) String() string {
	if this == nil {
		return "nil"
	}
	keysForLabels := make([]string, 0, len(this.Labels))
	for k, _ := range this.Labels {
		keysForLabels = append(keysForLabels, k)
	}
	github_com_gogo_protobuf_sortkeys.Strings(keysForLabels)
	mapStringForLabels := "map[string]string{"
	for _, k := range keysForLabels {
		mapStringForLabels += fmt.Sprintf("%v: %v,", k, this.Labels[k])
	}
	mapStringForLabels += "}"
	s := strings.Join([]string{`&AlertPolicyMatchRequest{`,
		`Namespace:` + fmt.Sprintf("%v", this.Namespace) + `,`,
		`Labels:` + mapStringForLabels + `,`,
		`}`,
	}, "")
	return s
}
func (this *AlertPolicyMatch) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&AlertPolicyMatch{`,
		`PolicyName:` + fmt.Sprintf("%v", this.PolicyName) + `,`,
		`PolicyStatus:` + fmt.Sprintf("%v", this.PolicyStatus) + `,`,
		`}`,
	}, "")
	return s
}
func (this *AlertPolicyMatchResponse) String() string {
	if this == nil {
		return "nil"
	}
	repeatedStringForAlertMatch := "[]*AlertPolicyMatch{"
	for _, f := range this.AlertMatch {
		repeatedStringForAlertMatch += strings.Replace(f.String(), "AlertPolicyMatch", "AlertPolicyMatch", 1) + ","
	}
	repeatedStringForAlertMatch += "}"
	s := strings.Join([]string{`&AlertPolicyMatchResponse{`,
		`AlertMatch:` + repeatedStringForAlertMatch + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringPublicCustomapi(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *AlertPolicyMatchRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPublicCustomapi
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
			return fmt.Errorf("proto: AlertPolicyMatchRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AlertPolicyMatchRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Namespace", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPublicCustomapi
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
				return ErrInvalidLengthPublicCustomapi
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPublicCustomapi
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Namespace = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Labels", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPublicCustomapi
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
				return ErrInvalidLengthPublicCustomapi
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPublicCustomapi
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Labels == nil {
				m.Labels = make(map[string]string)
			}
			var mapkey string
			var mapvalue string
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowPublicCustomapi
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
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowPublicCustomapi
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthPublicCustomapi
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey < 0 {
						return ErrInvalidLengthPublicCustomapi
					}
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var stringLenmapvalue uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowPublicCustomapi
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapvalue |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapvalue := int(stringLenmapvalue)
					if intStringLenmapvalue < 0 {
						return ErrInvalidLengthPublicCustomapi
					}
					postStringIndexmapvalue := iNdEx + intStringLenmapvalue
					if postStringIndexmapvalue < 0 {
						return ErrInvalidLengthPublicCustomapi
					}
					if postStringIndexmapvalue > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = string(dAtA[iNdEx:postStringIndexmapvalue])
					iNdEx = postStringIndexmapvalue
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipPublicCustomapi(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if skippy < 0 {
						return ErrInvalidLengthPublicCustomapi
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.Labels[mapkey] = mapvalue
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPublicCustomapi(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPublicCustomapi
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthPublicCustomapi
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
func (m *AlertPolicyMatch) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPublicCustomapi
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
			return fmt.Errorf("proto: AlertPolicyMatch: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AlertPolicyMatch: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PolicyName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPublicCustomapi
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
				return ErrInvalidLengthPublicCustomapi
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPublicCustomapi
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PolicyName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PolicyStatus", wireType)
			}
			m.PolicyStatus = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPublicCustomapi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PolicyStatus |= AlertPolicyStatus(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipPublicCustomapi(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPublicCustomapi
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthPublicCustomapi
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
func (m *AlertPolicyMatchResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPublicCustomapi
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
			return fmt.Errorf("proto: AlertPolicyMatchResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AlertPolicyMatchResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AlertMatch", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPublicCustomapi
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
				return ErrInvalidLengthPublicCustomapi
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPublicCustomapi
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AlertMatch = append(m.AlertMatch, &AlertPolicyMatch{})
			if err := m.AlertMatch[len(m.AlertMatch)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPublicCustomapi(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPublicCustomapi
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthPublicCustomapi
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
func skipPublicCustomapi(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPublicCustomapi
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
					return 0, ErrIntOverflowPublicCustomapi
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
					return 0, ErrIntOverflowPublicCustomapi
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
				return 0, ErrInvalidLengthPublicCustomapi
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPublicCustomapi
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPublicCustomapi
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPublicCustomapi        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPublicCustomapi          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPublicCustomapi = fmt.Errorf("proto: unexpected end of group")
)
