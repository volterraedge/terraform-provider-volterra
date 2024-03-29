// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ves.io/schema/topology/topology_site/types.proto

package topology_site

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	site "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/site"
	topology "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/topology"
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

// Topology Site
//
// x-displayName: "Site"
//
//	A node represents a site in the topology graph.
type GlobalSpecType struct {
	// Topology Metadata
	//
	// x-displayName: "Metadata"
	// A common metadata for topology.
	TopologyMetadata *topology.MetaType `protobuf:"bytes,1,opt,name=topology_metadata,json=topologyMetadata,proto3" json:"topology_metadata,omitempty"`
	// Topology Spec
	//
	// x-displayName: "Spec"
	// A canonical spec for this topology node.
	TopologySpec *topology.SiteType `protobuf:"bytes,2,opt,name=topology_spec,json=topologySpec,proto3" json:"topology_spec,omitempty"`
	// main_nodes
	//
	// x-displayName: "Main Nodes"
	// Connectivity information of main/master nodes to create a full mesh of Phobos services across all CEs in a site-mesh-group or dc-cluster-group.
	MainNodes []*site.Node `protobuf:"bytes,3,rep,name=main_nodes,json=mainNodes,proto3" json:"main_nodes,omitempty"`
}

func (m *GlobalSpecType) Reset()      { *m = GlobalSpecType{} }
func (*GlobalSpecType) ProtoMessage() {}
func (*GlobalSpecType) Descriptor() ([]byte, []int) {
	return fileDescriptor_04c03c060760c5fe, []int{0}
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

func (m *GlobalSpecType) GetTopologyMetadata() *topology.MetaType {
	if m != nil {
		return m.TopologyMetadata
	}
	return nil
}

func (m *GlobalSpecType) GetTopologySpec() *topology.SiteType {
	if m != nil {
		return m.TopologySpec
	}
	return nil
}

func (m *GlobalSpecType) GetMainNodes() []*site.Node {
	if m != nil {
		return m.MainNodes
	}
	return nil
}

func init() {
	proto.RegisterType((*GlobalSpecType)(nil), "ves.io.schema.topology.topology_site.GlobalSpecType")
}

func init() {
	proto.RegisterFile("ves.io/schema/topology/topology_site/types.proto", fileDescriptor_04c03c060760c5fe)
}

var fileDescriptor_04c03c060760c5fe = []byte{
	// 352 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x52, 0xb1, 0x6e, 0xe2, 0x40,
	0x10, 0xf5, 0x1e, 0xd2, 0x49, 0x67, 0xee, 0x4e, 0x89, 0x2b, 0x8b, 0x62, 0x84, 0x50, 0x0a, 0x1a,
	0x76, 0x23, 0x52, 0xa4, 0x8f, 0x14, 0xa5, 0x22, 0x05, 0xa4, 0xa2, 0x41, 0x6b, 0x7b, 0x30, 0x56,
	0x6c, 0x66, 0x65, 0x2f, 0x28, 0x74, 0xf9, 0x84, 0x14, 0xf9, 0x88, 0x7c, 0x4a, 0x4a, 0x4a, 0xba,
	0x84, 0xa5, 0x49, 0xc9, 0x27, 0x44, 0x36, 0x31, 0x08, 0x24, 0xe8, 0xde, 0xcc, 0xbe, 0xf7, 0xf6,
	0xcd, 0xec, 0xda, 0x97, 0x53, 0xcc, 0x78, 0x44, 0x22, 0xf3, 0x47, 0x98, 0x48, 0xa1, 0x49, 0x51,
	0x4c, 0xe1, 0x6c, 0x0b, 0x06, 0x59, 0xa4, 0x51, 0xe8, 0x99, 0xc2, 0x8c, 0xab, 0x94, 0x34, 0x39,
	0x17, 0x1b, 0x05, 0xdf, 0x28, 0x78, 0x49, 0xe4, 0x7b, 0x8a, 0x5a, 0x2b, 0x8c, 0xf4, 0x68, 0xe2,
	0x71, 0x9f, 0x12, 0x11, 0x52, 0x48, 0xa2, 0x10, 0x7b, 0x93, 0x61, 0x51, 0x15, 0x45, 0x81, 0x36,
	0xa6, 0x35, 0xd8, 0x8f, 0x71, 0x78, 0x69, 0xad, 0x71, 0x2c, 0xe6, 0x8e, 0xd3, 0xf8, 0x60, 0xf6,
	0xff, 0xbb, 0x98, 0x3c, 0x19, 0xf7, 0x14, 0xfa, 0x0f, 0x33, 0x85, 0x4e, 0xc7, 0x3e, 0xdf, 0xc6,
	0x4a, 0x50, 0xcb, 0x40, 0x6a, 0xe9, 0xb2, 0x3a, 0x6b, 0x56, 0xdb, 0x75, 0x7e, 0x64, 0x8e, 0x0e,
	0x6a, 0x99, 0x8b, 0xbb, 0x67, 0x65, 0xab, 0xf3, 0xa3, 0x74, 0x6e, 0xed, 0x7f, 0xbb, 0x29, 0x15,
	0xfa, 0xee, 0xaf, 0xd3, 0x56, 0xbd, 0x48, 0x63, 0x61, 0xf5, 0xb7, 0x6c, 0xe5, 0xc9, 0x9c, 0x6b,
	0xdb, 0x4e, 0x64, 0x34, 0x1e, 0x8c, 0x29, 0xc0, 0xcc, 0xad, 0xd4, 0x2b, 0xcd, 0x6a, 0xdb, 0x3d,
	0xf0, 0xc8, 0x37, 0xc0, 0xef, 0x29, 0xc0, 0xee, 0x9f, 0x9c, 0x9b, 0xa3, 0xec, 0xe6, 0x95, 0xcd,
	0x97, 0x60, 0x2d, 0x96, 0x60, 0xad, 0x97, 0xc0, 0x9e, 0x0d, 0xb0, 0x37, 0x03, 0xec, 0xdd, 0x00,
	0x9b, 0x1b, 0x60, 0x0b, 0x03, 0xec, 0xd3, 0x00, 0xfb, 0x32, 0x60, 0xad, 0x0d, 0xb0, 0x97, 0x15,
	0x58, 0xf3, 0x15, 0x58, 0x8b, 0x15, 0x58, 0xfd, 0x7e, 0x48, 0xea, 0x31, 0xe4, 0x53, 0x8a, 0x35,
	0xa6, 0xa9, 0xe4, 0x93, 0x4c, 0x14, 0x60, 0x48, 0x69, 0xd2, 0x52, 0x29, 0x4d, 0xa3, 0x00, 0xd3,
	0x56, 0x79, 0x2c, 0x94, 0x17, 0x92, 0xc0, 0x27, 0x5d, 0xbe, 0xc7, 0xa9, 0xdf, 0xe1, 0xfd, 0x2e,
	0xf6, 0x7f, 0xf5, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x75, 0x49, 0xe3, 0xdf, 0x4c, 0x02, 0x00, 0x00,
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
	if !this.TopologyMetadata.Equal(that1.TopologyMetadata) {
		return false
	}
	if !this.TopologySpec.Equal(that1.TopologySpec) {
		return false
	}
	if len(this.MainNodes) != len(that1.MainNodes) {
		return false
	}
	for i := range this.MainNodes {
		if !this.MainNodes[i].Equal(that1.MainNodes[i]) {
			return false
		}
	}
	return true
}
func (this *GlobalSpecType) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 7)
	s = append(s, "&topology_site.GlobalSpecType{")
	if this.TopologyMetadata != nil {
		s = append(s, "TopologyMetadata: "+fmt.Sprintf("%#v", this.TopologyMetadata)+",\n")
	}
	if this.TopologySpec != nil {
		s = append(s, "TopologySpec: "+fmt.Sprintf("%#v", this.TopologySpec)+",\n")
	}
	if this.MainNodes != nil {
		s = append(s, "MainNodes: "+fmt.Sprintf("%#v", this.MainNodes)+",\n")
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
	if len(m.MainNodes) > 0 {
		for iNdEx := len(m.MainNodes) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.MainNodes[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTypes(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.TopologySpec != nil {
		{
			size, err := m.TopologySpec.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.TopologyMetadata != nil {
		{
			size, err := m.TopologyMetadata.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
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
	if m.TopologyMetadata != nil {
		l = m.TopologyMetadata.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.TopologySpec != nil {
		l = m.TopologySpec.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	if len(m.MainNodes) > 0 {
		for _, e := range m.MainNodes {
			l = e.Size()
			n += 1 + l + sovTypes(uint64(l))
		}
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
	repeatedStringForMainNodes := "[]*Node{"
	for _, f := range this.MainNodes {
		repeatedStringForMainNodes += strings.Replace(fmt.Sprintf("%v", f), "Node", "site.Node", 1) + ","
	}
	repeatedStringForMainNodes += "}"
	s := strings.Join([]string{`&GlobalSpecType{`,
		`TopologyMetadata:` + strings.Replace(fmt.Sprintf("%v", this.TopologyMetadata), "MetaType", "topology.MetaType", 1) + `,`,
		`TopologySpec:` + strings.Replace(fmt.Sprintf("%v", this.TopologySpec), "SiteType", "topology.SiteType", 1) + `,`,
		`MainNodes:` + repeatedStringForMainNodes + `,`,
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
				return fmt.Errorf("proto: wrong wireType = %d for field TopologyMetadata", wireType)
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
			if m.TopologyMetadata == nil {
				m.TopologyMetadata = &topology.MetaType{}
			}
			if err := m.TopologyMetadata.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TopologySpec", wireType)
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
			if m.TopologySpec == nil {
				m.TopologySpec = &topology.SiteType{}
			}
			if err := m.TopologySpec.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MainNodes", wireType)
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
			m.MainNodes = append(m.MainNodes, &site.Node{})
			if err := m.MainNodes[len(m.MainNodes)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
