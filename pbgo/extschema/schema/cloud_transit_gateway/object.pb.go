// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ves.io/schema/cloud_transit_gateway/object.proto

package cloud_transit_gateway

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	golang_proto "github.com/golang/protobuf/proto"
	schema "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema"
	cloud_connect "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/cloud_connect"
	_ "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/vesenv"
	aws_tgw_site "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/views/aws_tgw_site"
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

// Cloud Transit Gateway Object
//
// x-displayName: "Object"
// Cloud Transit Gateway object
type Object struct {
	// metadata
	//
	// x-displayName: "Metadata"
	// Standard object's metadata
	Metadata *schema.ObjectMetaType `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// system_metadata
	//
	// x-displayName: "System Metadata"
	// System generated object's metadata
	SystemMetadata *schema.SystemObjectMetaType `protobuf:"bytes,2,opt,name=system_metadata,json=systemMetadata,proto3" json:"system_metadata,omitempty"`
	// spec
	//
	// x-displayName: "Spec"
	// Specification of the desired behavior of the Cloud Transit Gateway
	Spec *SpecType `protobuf:"bytes,3,opt,name=spec,proto3" json:"spec,omitempty"`
}

func (m *Object) Reset()      { *m = Object{} }
func (*Object) ProtoMessage() {}
func (*Object) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b3671c15b9c4afb, []int{0}
}
func (m *Object) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Object) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *Object) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Object.Merge(m, src)
}
func (m *Object) XXX_Size() int {
	return m.Size()
}
func (m *Object) XXX_DiscardUnknown() {
	xxx_messageInfo_Object.DiscardUnknown(m)
}

var xxx_messageInfo_Object proto.InternalMessageInfo

func (m *Object) GetMetadata() *schema.ObjectMetaType {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *Object) GetSystemMetadata() *schema.SystemObjectMetaType {
	if m != nil {
		return m.SystemMetadata
	}
	return nil
}

func (m *Object) GetSpec() *SpecType {
	if m != nil {
		return m.Spec
	}
	return nil
}

// Specification for Cloud Transit Gateway
//
// x-displayName: "Specification"
// Shape of the Cloud Transit Gateway  specification
type SpecType struct {
	// gc_spec
	//
	// x-displayName: "GC Spec"
	GcSpec *GlobalSpecType `protobuf:"bytes,1,opt,name=gc_spec,json=gcSpec,proto3" json:"gc_spec,omitempty"`
}

func (m *SpecType) Reset()      { *m = SpecType{} }
func (*SpecType) ProtoMessage() {}
func (*SpecType) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b3671c15b9c4afb, []int{1}
}
func (m *SpecType) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SpecType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *SpecType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SpecType.Merge(m, src)
}
func (m *SpecType) XXX_Size() int {
	return m.Size()
}
func (m *SpecType) XXX_DiscardUnknown() {
	xxx_messageInfo_SpecType.DiscardUnknown(m)
}

var xxx_messageInfo_SpecType proto.InternalMessageInfo

func (m *SpecType) GetGcSpec() *GlobalSpecType {
	if m != nil {
		return m.GcSpec
	}
	return nil
}

// Status for Cloud Transit Gateway
//
// x-displayName: "Status"
// Most recently observed status of object
type StatusObject struct {
	// metadata
	//
	// x-displayName: "Metadata"
	// Standard status's metadata
	Metadata *schema.StatusMetaType `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// object_refs
	//
	// x-displayName: "Config Object"
	// Reference to object for current status
	ObjectRefs []*schema.ObjectRefType `protobuf:"bytes,2,rep,name=object_refs,json=objectRefs,proto3" json:"object_refs,omitempty"`
	// conditions
	//
	// x-displayName: "Conditions"
	// Conditions reported by various component of the system
	Conditions []*schema.ConditionType `protobuf:"bytes,3,rep,name=conditions,proto3" json:"conditions,omitempty"`
	// AWS Transit Gateway Status Type
	//
	// x-displayName: "AWS TGW Status"
	// AWS Transit Gateway Status Type
	Tgw *aws_tgw_site.AWSTGWStatusType `protobuf:"bytes,4,opt,name=tgw,proto3" json:"tgw,omitempty"`
	// AWS TGW Resource Share
	//
	// x-displayName: "AWS TGW Resource Share"
	// AWS TGW Resource Share
	ResourceShareStatus *aws_tgw_site.AWSTGWResourceShareType `protobuf:"bytes,5,opt,name=resource_share_status,json=resourceShareStatus,proto3" json:"resource_share_status,omitempty"`
	// Cloud Connect Attached to AWS TGW Site
	//
	// x-displayName: "Cloud Connect to AWS TGW Site"
	// Cloud Connect to AWS Sites
	VpcAttachments *cloud_connect.AWSAttachmentsListStatusType `protobuf:"bytes,8,opt,name=vpc_attachments,json=vpcAttachments,proto3" json:"vpc_attachments,omitempty"`
}

func (m *StatusObject) Reset()      { *m = StatusObject{} }
func (*StatusObject) ProtoMessage() {}
func (*StatusObject) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b3671c15b9c4afb, []int{2}
}
func (m *StatusObject) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *StatusObject) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *StatusObject) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatusObject.Merge(m, src)
}
func (m *StatusObject) XXX_Size() int {
	return m.Size()
}
func (m *StatusObject) XXX_DiscardUnknown() {
	xxx_messageInfo_StatusObject.DiscardUnknown(m)
}

var xxx_messageInfo_StatusObject proto.InternalMessageInfo

func (m *StatusObject) GetMetadata() *schema.StatusMetaType {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *StatusObject) GetObjectRefs() []*schema.ObjectRefType {
	if m != nil {
		return m.ObjectRefs
	}
	return nil
}

func (m *StatusObject) GetConditions() []*schema.ConditionType {
	if m != nil {
		return m.Conditions
	}
	return nil
}

func (m *StatusObject) GetTgw() *aws_tgw_site.AWSTGWStatusType {
	if m != nil {
		return m.Tgw
	}
	return nil
}

func (m *StatusObject) GetResourceShareStatus() *aws_tgw_site.AWSTGWResourceShareType {
	if m != nil {
		return m.ResourceShareStatus
	}
	return nil
}

func (m *StatusObject) GetVpcAttachments() *cloud_connect.AWSAttachmentsListStatusType {
	if m != nil {
		return m.VpcAttachments
	}
	return nil
}

func init() {
	proto.RegisterType((*Object)(nil), "ves.io.schema.cloud_transit_gateway.Object")
	golang_proto.RegisterType((*Object)(nil), "ves.io.schema.cloud_transit_gateway.Object")
	proto.RegisterType((*SpecType)(nil), "ves.io.schema.cloud_transit_gateway.SpecType")
	golang_proto.RegisterType((*SpecType)(nil), "ves.io.schema.cloud_transit_gateway.SpecType")
	proto.RegisterType((*StatusObject)(nil), "ves.io.schema.cloud_transit_gateway.StatusObject")
	golang_proto.RegisterType((*StatusObject)(nil), "ves.io.schema.cloud_transit_gateway.StatusObject")
}

func init() {
	proto.RegisterFile("ves.io/schema/cloud_transit_gateway/object.proto", fileDescriptor_9b3671c15b9c4afb)
}
func init() {
	golang_proto.RegisterFile("ves.io/schema/cloud_transit_gateway/object.proto", fileDescriptor_9b3671c15b9c4afb)
}

var fileDescriptor_9b3671c15b9c4afb = []byte{
	// 747 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x94, 0xcf, 0x4e, 0xdb, 0x4a,
	0x14, 0xc6, 0xe3, 0x24, 0x04, 0x34, 0x20, 0x40, 0x46, 0x48, 0xbe, 0xdc, 0x7b, 0x2d, 0x14, 0x16,
	0x65, 0x81, 0xed, 0x16, 0xd4, 0x05, 0x08, 0x15, 0xf1, 0xa7, 0x62, 0x03, 0xaa, 0x94, 0x20, 0xd1,
	0x76, 0x63, 0x4d, 0xc6, 0x27, 0x8e, 0xdb, 0xd8, 0x63, 0xcd, 0x8c, 0x1d, 0xb2, 0xa8, 0x54, 0xf5,
	0x01, 0xaa, 0xee, 0x2a, 0xf5, 0x09, 0xaa, 0x3e, 0x45, 0xe9, 0xaa, 0x4b, 0xca, 0x8a, 0x55, 0x55,
	0x9c, 0x4d, 0x97, 0xac, 0xbb, 0xaa, 0x32, 0x26, 0xd4, 0x89, 0x42, 0x15, 0x89, 0xdd, 0xd8, 0xe7,
	0xfb, 0x7e, 0x73, 0x66, 0xce, 0x99, 0x83, 0xee, 0xc7, 0xc0, 0x4d, 0x8f, 0x5a, 0x9c, 0x34, 0xc0,
	0xc7, 0x16, 0x69, 0xd2, 0xc8, 0xb1, 0x05, 0xc3, 0x01, 0xf7, 0x84, 0xed, 0x62, 0x01, 0x2d, 0xdc,
	0xb6, 0x68, 0xed, 0x05, 0x10, 0x61, 0x86, 0x8c, 0x0a, 0xaa, 0x2e, 0xa5, 0x0e, 0x33, 0x75, 0x98,
	0x43, 0x1d, 0x0b, 0x86, 0xeb, 0x89, 0x46, 0x54, 0x33, 0x09, 0xf5, 0x2d, 0x97, 0xba, 0xd4, 0x92,
	0xde, 0x5a, 0x54, 0x97, 0x5f, 0xf2, 0x43, 0xae, 0x52, 0xe6, 0xc2, 0xbd, 0x61, 0x59, 0x10, 0x1a,
	0x04, 0x40, 0x84, 0x25, 0xda, 0x21, 0xf0, 0x6b, 0xa1, 0x35, 0x4a, 0xba, 0x59, 0xc3, 0xbf, 0xfd,
	0x06, 0x1a, 0x0a, 0x8f, 0x06, 0xbd, 0xe0, 0x3f, 0xfd, 0xc1, 0xac, 0xaf, 0xdc, 0x1f, 0x8a, 0x81,
	0x43, 0x10, 0x0f, 0xd8, 0x57, 0x06, 0x34, 0x1e, 0xb4, 0xb8, 0x85, 0x5b, 0xdc, 0x16, 0x6e, 0xcb,
	0xe6, 0x9e, 0x80, 0x2c, 0xb1, 0xfc, 0xad, 0x88, 0x4a, 0x4f, 0xe4, 0x45, 0xaa, 0xeb, 0x68, 0xc2,
	0x07, 0x81, 0x1d, 0x2c, 0xb0, 0xa6, 0x2c, 0x2a, 0xcb, 0x93, 0xab, 0xff, 0x9b, 0xfd, 0xb7, 0x9a,
	0x0a, 0x0f, 0x41, 0xe0, 0xa3, 0x76, 0x08, 0x95, 0x1b, 0xb9, 0x7a, 0x80, 0x66, 0x78, 0x9b, 0x0b,
	0xf0, 0xed, 0x1b, 0x42, 0x5e, 0x12, 0x96, 0x06, 0x08, 0x55, 0xa9, 0x1a, 0xe0, 0x4c, 0xa7, 0xde,
	0xc3, 0x1e, 0x6d, 0x1b, 0x15, 0x79, 0x08, 0x44, 0x2b, 0x48, 0x84, 0x61, 0x8e, 0x50, 0x5a, 0xb3,
	0x1a, 0x02, 0x91, 0x30, 0x69, 0xdd, 0xf8, 0x95, 0x7f, 0xb3, 0x55, 0x4a, 0xb1, 0x1f, 0xb6, 0xc6,
	0x1c, 0xf0, 0xe9, 0x83, 0x2f, 0xa7, 0x5a, 0x7e, 0x56, 0x39, 0x3f, 0xd5, 0xbe, 0xe7, 0xd1, 0x2e,
	0x2a, 0xec, 0x83, 0x50, 0x37, 0xeb, 0x0f, 0x4f, 0x88, 0xc1, 0x81, 0x44, 0x0c, 0x7c, 0xe0, 0x0d,
	0x83, 0x0b, 0x1c, 0x38, 0x98, 0x39, 0x86, 0x4f, 0x03, 0x4f, 0x50, 0xb6, 0xb2, 0x18, 0x03, 0x37,
	0x3c, 0x6a, 0x78, 0x41, 0x9d, 0x61, 0x2e, 0x58, 0x44, 0x44, 0xc4, 0xc0, 0x60, 0x80, 0x1d, 0xf4,
	0x18, 0x8d, 0x57, 0x20, 0x6c, 0x62, 0x02, 0xea, 0xc6, 0xad, 0xa0, 0x88, 0xc3, 0xad, 0x94, 0x16,
	0xf3, 0x04, 0xa0, 0x3d, 0x54, 0xda, 0x65, 0x80, 0xc5, 0x9d, 0x29, 0x7b, 0xd0, 0x84, 0x3b, 0x53,
	0x8a, 0x07, 0x1e, 0xbf, 0xe3, 0xc5, 0x94, 0x9f, 0xa2, 0x89, 0x5e, 0x39, 0xd4, 0x03, 0x34, 0xee,
	0x12, 0x5b, 0x96, 0x33, 0xed, 0xa9, 0xb5, 0x91, 0xca, 0xb9, 0xdf, 0xa4, 0x35, 0xdc, 0xbc, 0x29,
	0x6a, 0xc9, 0x25, 0xdd, 0x75, 0xf9, 0x6d, 0x11, 0x4d, 0x55, 0x05, 0x16, 0x11, 0x1f, 0xb9, 0x67,
	0x53, 0xf9, 0x90, 0x9e, 0x3d, 0x46, 0x93, 0xe9, 0x04, 0xb1, 0x19, 0xd4, 0xb9, 0x96, 0x5f, 0x2c,
	0x2c, 0x4f, 0xae, 0xfe, 0x37, 0xb4, 0xe3, 0x2b, 0x50, 0xef, 0x9a, 0x77, 0xb4, 0x4f, 0xaf, 0xe6,
	0x87, 0xe6, 0x5b, 0x41, 0xb4, 0x27, 0xe4, 0xea, 0x26, 0x42, 0x84, 0x06, 0x8e, 0x27, 0x1f, 0xa5,
	0x56, 0x18, 0xca, 0xdd, 0xed, 0x09, 0x64, 0x52, 0x19, 0xbd, 0xba, 0x87, 0x0a, 0xc2, 0x6d, 0x69,
	0x45, 0x79, 0x98, 0xd5, 0x01, 0x9b, 0x7c, 0xcc, 0x66, 0xf6, 0x31, 0x9b, 0xdb, 0xc7, 0xd5, 0xa3,
	0xfd, 0xe3, 0xf4, 0x94, 0x12, 0xd6, 0xb5, 0xab, 0x3e, 0x9a, 0x67, 0xc0, 0x69, 0xc4, 0x08, 0xd8,
	0xbc, 0x81, 0x19, 0xd8, 0x5c, 0x2a, 0xb4, 0x31, 0xc9, 0x5d, 0x1f, 0x95, 0x5b, 0xb9, 0x86, 0x54,
	0xbb, 0x0c, 0x89, 0x9f, 0x63, 0xd9, 0x5f, 0xe9, 0xbe, 0x6a, 0x0d, 0xcd, 0xc4, 0x21, 0xb1, 0xb1,
	0x10, 0x98, 0x34, 0x7c, 0x08, 0x04, 0xd7, 0x26, 0x86, 0x6e, 0xd4, 0x37, 0x43, 0xbb, 0x7b, 0x6c,
	0xff, 0xb1, 0x74, 0x3b, 0x2f, 0x73, 0x8e, 0xe9, 0x38, 0x24, 0x99, 0xe8, 0xc6, 0xdc, 0xf9, 0xa3,
	0x59, 0x34, 0x8d, 0xa6, 0x7a, 0x05, 0x34, 0x23, 0xcf, 0xd9, 0x79, 0xaf, 0x9c, 0x5d, 0xea, 0xb9,
	0x8b, 0x4b, 0x3d, 0x77, 0x75, 0xa9, 0x2b, 0xaf, 0x13, 0x5d, 0xf9, 0x98, 0xe8, 0xca, 0xd7, 0x44,
	0x57, 0xce, 0x12, 0x5d, 0xb9, 0x48, 0x74, 0xe5, 0x47, 0xa2, 0x2b, 0x3f, 0x13, 0x3d, 0x77, 0x95,
	0xe8, 0xca, 0xbb, 0x8e, 0x9e, 0xfb, 0xdc, 0xd1, 0x95, 0xb3, 0x8e, 0x9e, 0xbb, 0xe8, 0xe8, 0xb9,
	0xe7, 0xcf, 0x5c, 0x1a, 0xbe, 0x74, 0xcd, 0x98, 0x36, 0x05, 0x30, 0x86, 0xcd, 0x88, 0x5b, 0x72,
	0x51, 0xa7, 0xcc, 0x37, 0x42, 0x46, 0x63, 0xcf, 0x01, 0x66, 0xf4, 0xc2, 0x56, 0x58, 0x73, 0xa9,
	0x05, 0x27, 0xe2, 0x7a, 0xb4, 0xfe, 0x6d, 0xdc, 0xd7, 0x4a, 0x72, 0xbe, 0xae, 0xfd, 0x0e, 0x00,
	0x00, 0xff, 0xff, 0x1c, 0x3f, 0x1a, 0xad, 0xcb, 0x06, 0x00, 0x00,
}

func (this *Object) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Object)
	if !ok {
		that2, ok := that.(Object)
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
	if !this.Metadata.Equal(that1.Metadata) {
		return false
	}
	if !this.SystemMetadata.Equal(that1.SystemMetadata) {
		return false
	}
	if !this.Spec.Equal(that1.Spec) {
		return false
	}
	return true
}
func (this *SpecType) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*SpecType)
	if !ok {
		that2, ok := that.(SpecType)
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
	if !this.GcSpec.Equal(that1.GcSpec) {
		return false
	}
	return true
}
func (this *StatusObject) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*StatusObject)
	if !ok {
		that2, ok := that.(StatusObject)
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
	if !this.Metadata.Equal(that1.Metadata) {
		return false
	}
	if len(this.ObjectRefs) != len(that1.ObjectRefs) {
		return false
	}
	for i := range this.ObjectRefs {
		if !this.ObjectRefs[i].Equal(that1.ObjectRefs[i]) {
			return false
		}
	}
	if len(this.Conditions) != len(that1.Conditions) {
		return false
	}
	for i := range this.Conditions {
		if !this.Conditions[i].Equal(that1.Conditions[i]) {
			return false
		}
	}
	if !this.Tgw.Equal(that1.Tgw) {
		return false
	}
	if !this.ResourceShareStatus.Equal(that1.ResourceShareStatus) {
		return false
	}
	if !this.VpcAttachments.Equal(that1.VpcAttachments) {
		return false
	}
	return true
}
func (this *Object) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 7)
	s = append(s, "&cloud_transit_gateway.Object{")
	if this.Metadata != nil {
		s = append(s, "Metadata: "+fmt.Sprintf("%#v", this.Metadata)+",\n")
	}
	if this.SystemMetadata != nil {
		s = append(s, "SystemMetadata: "+fmt.Sprintf("%#v", this.SystemMetadata)+",\n")
	}
	if this.Spec != nil {
		s = append(s, "Spec: "+fmt.Sprintf("%#v", this.Spec)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *SpecType) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&cloud_transit_gateway.SpecType{")
	if this.GcSpec != nil {
		s = append(s, "GcSpec: "+fmt.Sprintf("%#v", this.GcSpec)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *StatusObject) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 10)
	s = append(s, "&cloud_transit_gateway.StatusObject{")
	if this.Metadata != nil {
		s = append(s, "Metadata: "+fmt.Sprintf("%#v", this.Metadata)+",\n")
	}
	if this.ObjectRefs != nil {
		s = append(s, "ObjectRefs: "+fmt.Sprintf("%#v", this.ObjectRefs)+",\n")
	}
	if this.Conditions != nil {
		s = append(s, "Conditions: "+fmt.Sprintf("%#v", this.Conditions)+",\n")
	}
	if this.Tgw != nil {
		s = append(s, "Tgw: "+fmt.Sprintf("%#v", this.Tgw)+",\n")
	}
	if this.ResourceShareStatus != nil {
		s = append(s, "ResourceShareStatus: "+fmt.Sprintf("%#v", this.ResourceShareStatus)+",\n")
	}
	if this.VpcAttachments != nil {
		s = append(s, "VpcAttachments: "+fmt.Sprintf("%#v", this.VpcAttachments)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringObject(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *Object) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Object) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Object) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Spec != nil {
		{
			size, err := m.Spec.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintObject(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.SystemMetadata != nil {
		{
			size, err := m.SystemMetadata.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintObject(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.Metadata != nil {
		{
			size, err := m.Metadata.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintObject(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *SpecType) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SpecType) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SpecType) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.GcSpec != nil {
		{
			size, err := m.GcSpec.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintObject(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *StatusObject) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StatusObject) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *StatusObject) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.VpcAttachments != nil {
		{
			size, err := m.VpcAttachments.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintObject(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x42
	}
	if m.ResourceShareStatus != nil {
		{
			size, err := m.ResourceShareStatus.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintObject(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x2a
	}
	if m.Tgw != nil {
		{
			size, err := m.Tgw.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintObject(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if len(m.Conditions) > 0 {
		for iNdEx := len(m.Conditions) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Conditions[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintObject(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.ObjectRefs) > 0 {
		for iNdEx := len(m.ObjectRefs) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ObjectRefs[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintObject(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if m.Metadata != nil {
		{
			size, err := m.Metadata.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintObject(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintObject(dAtA []byte, offset int, v uint64) int {
	offset -= sovObject(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Object) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Metadata != nil {
		l = m.Metadata.Size()
		n += 1 + l + sovObject(uint64(l))
	}
	if m.SystemMetadata != nil {
		l = m.SystemMetadata.Size()
		n += 1 + l + sovObject(uint64(l))
	}
	if m.Spec != nil {
		l = m.Spec.Size()
		n += 1 + l + sovObject(uint64(l))
	}
	return n
}

func (m *SpecType) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.GcSpec != nil {
		l = m.GcSpec.Size()
		n += 1 + l + sovObject(uint64(l))
	}
	return n
}

func (m *StatusObject) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Metadata != nil {
		l = m.Metadata.Size()
		n += 1 + l + sovObject(uint64(l))
	}
	if len(m.ObjectRefs) > 0 {
		for _, e := range m.ObjectRefs {
			l = e.Size()
			n += 1 + l + sovObject(uint64(l))
		}
	}
	if len(m.Conditions) > 0 {
		for _, e := range m.Conditions {
			l = e.Size()
			n += 1 + l + sovObject(uint64(l))
		}
	}
	if m.Tgw != nil {
		l = m.Tgw.Size()
		n += 1 + l + sovObject(uint64(l))
	}
	if m.ResourceShareStatus != nil {
		l = m.ResourceShareStatus.Size()
		n += 1 + l + sovObject(uint64(l))
	}
	if m.VpcAttachments != nil {
		l = m.VpcAttachments.Size()
		n += 1 + l + sovObject(uint64(l))
	}
	return n
}

func sovObject(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozObject(x uint64) (n int) {
	return sovObject(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *Object) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Object{`,
		`Metadata:` + strings.Replace(fmt.Sprintf("%v", this.Metadata), "ObjectMetaType", "schema.ObjectMetaType", 1) + `,`,
		`SystemMetadata:` + strings.Replace(fmt.Sprintf("%v", this.SystemMetadata), "SystemObjectMetaType", "schema.SystemObjectMetaType", 1) + `,`,
		`Spec:` + strings.Replace(this.Spec.String(), "SpecType", "SpecType", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *SpecType) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&SpecType{`,
		`GcSpec:` + strings.Replace(fmt.Sprintf("%v", this.GcSpec), "GlobalSpecType", "GlobalSpecType", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *StatusObject) String() string {
	if this == nil {
		return "nil"
	}
	repeatedStringForObjectRefs := "[]*ObjectRefType{"
	for _, f := range this.ObjectRefs {
		repeatedStringForObjectRefs += strings.Replace(fmt.Sprintf("%v", f), "ObjectRefType", "schema.ObjectRefType", 1) + ","
	}
	repeatedStringForObjectRefs += "}"
	repeatedStringForConditions := "[]*ConditionType{"
	for _, f := range this.Conditions {
		repeatedStringForConditions += strings.Replace(fmt.Sprintf("%v", f), "ConditionType", "schema.ConditionType", 1) + ","
	}
	repeatedStringForConditions += "}"
	s := strings.Join([]string{`&StatusObject{`,
		`Metadata:` + strings.Replace(fmt.Sprintf("%v", this.Metadata), "StatusMetaType", "schema.StatusMetaType", 1) + `,`,
		`ObjectRefs:` + repeatedStringForObjectRefs + `,`,
		`Conditions:` + repeatedStringForConditions + `,`,
		`Tgw:` + strings.Replace(fmt.Sprintf("%v", this.Tgw), "AWSTGWStatusType", "aws_tgw_site.AWSTGWStatusType", 1) + `,`,
		`ResourceShareStatus:` + strings.Replace(fmt.Sprintf("%v", this.ResourceShareStatus), "AWSTGWResourceShareType", "aws_tgw_site.AWSTGWResourceShareType", 1) + `,`,
		`VpcAttachments:` + strings.Replace(fmt.Sprintf("%v", this.VpcAttachments), "AWSAttachmentsListStatusType", "cloud_connect.AWSAttachmentsListStatusType", 1) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringObject(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *Object) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowObject
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
			return fmt.Errorf("proto: Object: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Object: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Metadata", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowObject
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
				return ErrInvalidLengthObject
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthObject
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Metadata == nil {
				m.Metadata = &schema.ObjectMetaType{}
			}
			if err := m.Metadata.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SystemMetadata", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowObject
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
				return ErrInvalidLengthObject
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthObject
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.SystemMetadata == nil {
				m.SystemMetadata = &schema.SystemObjectMetaType{}
			}
			if err := m.SystemMetadata.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Spec", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowObject
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
				return ErrInvalidLengthObject
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthObject
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Spec == nil {
				m.Spec = &SpecType{}
			}
			if err := m.Spec.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipObject(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthObject
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthObject
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
func (m *SpecType) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowObject
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
			return fmt.Errorf("proto: SpecType: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SpecType: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GcSpec", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowObject
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
				return ErrInvalidLengthObject
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthObject
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.GcSpec == nil {
				m.GcSpec = &GlobalSpecType{}
			}
			if err := m.GcSpec.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipObject(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthObject
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthObject
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
func (m *StatusObject) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowObject
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
			return fmt.Errorf("proto: StatusObject: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StatusObject: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Metadata", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowObject
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
				return ErrInvalidLengthObject
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthObject
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Metadata == nil {
				m.Metadata = &schema.StatusMetaType{}
			}
			if err := m.Metadata.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ObjectRefs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowObject
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
				return ErrInvalidLengthObject
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthObject
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ObjectRefs = append(m.ObjectRefs, &schema.ObjectRefType{})
			if err := m.ObjectRefs[len(m.ObjectRefs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Conditions", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowObject
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
				return ErrInvalidLengthObject
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthObject
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Conditions = append(m.Conditions, &schema.ConditionType{})
			if err := m.Conditions[len(m.Conditions)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tgw", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowObject
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
				return ErrInvalidLengthObject
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthObject
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Tgw == nil {
				m.Tgw = &aws_tgw_site.AWSTGWStatusType{}
			}
			if err := m.Tgw.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ResourceShareStatus", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowObject
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
				return ErrInvalidLengthObject
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthObject
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ResourceShareStatus == nil {
				m.ResourceShareStatus = &aws_tgw_site.AWSTGWResourceShareType{}
			}
			if err := m.ResourceShareStatus.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VpcAttachments", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowObject
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
				return ErrInvalidLengthObject
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthObject
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.VpcAttachments == nil {
				m.VpcAttachments = &cloud_connect.AWSAttachmentsListStatusType{}
			}
			if err := m.VpcAttachments.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipObject(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthObject
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthObject
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
func skipObject(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowObject
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
					return 0, ErrIntOverflowObject
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
					return 0, ErrIntOverflowObject
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
				return 0, ErrInvalidLengthObject
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupObject
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthObject
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthObject        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowObject          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupObject = fmt.Errorf("proto: unexpected end of group")
)
