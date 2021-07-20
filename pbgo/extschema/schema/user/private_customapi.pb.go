// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ves.io/schema/user/private_customapi.proto

package user

import proto "github.com/gogo/protobuf/proto"
import golang_proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import _ "github.com/gogo/googleapis/google/api"
import google_protobuf1 "github.com/gogo/protobuf/types"
import _ "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema"
import _ "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema"

import strings "strings"
import reflect "reflect"

import context "golang.org/x/net/context"
import grpc "google.golang.org/grpc"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = golang_proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// LastLoginUpdateRequest
//
// x-displayName: "Last Login Update Request"
// Request to update user login timestamp.
type LastLoginUpdateRequest struct {
	// user
	//
	// x-displayName: "User"
	// x-example: "user@company1.com"
	// User ID of the user. typically email id
	User string `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	// tenant
	//
	// x-displayName: "Tenant"
	// x-example: "company1-as432s"
	// Tenant ID of the tenant user belongs to.
	Tenant string `protobuf:"bytes,2,opt,name=tenant,proto3" json:"tenant,omitempty"`
	// last_login_timestamp
	//
	// x-displayName: "Last Login Timestamp"
	// Last successful login timestamp of the user .
	LastLoginTimestamp *google_protobuf1.Timestamp `protobuf:"bytes,3,opt,name=last_login_timestamp,json=lastLoginTimestamp" json:"last_login_timestamp,omitempty"`
}

func (m *LastLoginUpdateRequest) Reset()      { *m = LastLoginUpdateRequest{} }
func (*LastLoginUpdateRequest) ProtoMessage() {}
func (*LastLoginUpdateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptorPrivateCustomapi, []int{0}
}

func (m *LastLoginUpdateRequest) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *LastLoginUpdateRequest) GetTenant() string {
	if m != nil {
		return m.Tenant
	}
	return ""
}

func (m *LastLoginUpdateRequest) GetLastLoginTimestamp() *google_protobuf1.Timestamp {
	if m != nil {
		return m.LastLoginTimestamp
	}
	return nil
}

// LastLoginUpdateResponse
//
// x-displayName: "Last Login Update Response"
type LastLoginUpdateResponse struct {
}

func (m *LastLoginUpdateResponse) Reset()      { *m = LastLoginUpdateResponse{} }
func (*LastLoginUpdateResponse) ProtoMessage() {}
func (*LastLoginUpdateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptorPrivateCustomapi, []int{1}
}

// PrivateCascadeDeleteRequest
//
// x-displayName: "Delete the User and Associated Namespace Roles"
// PrivateCascadeDeleteRequest is the request to delete the user along with the associated namespace role objects.
type PrivateCascadeDeleteRequest struct {
	// Namespace
	//
	// x-displayName: "Namespace"
	// x-example: "value"
	// Value of namespace is always "system"
	Namespace string `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	// Tenant name
	//
	// x-displayName: "Tenant name"
	// x-required
	// User deletion will be executed within this tenant.
	TenantName string `protobuf:"bytes,2,opt,name=tenant_name,json=tenantName,proto3" json:"tenant_name,omitempty"`
	// email of the user
	//
	// x-displayName: "Email"
	// x-example: "value"
	// x-required
	// email of the user requesting for
	Email string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
}

func (m *PrivateCascadeDeleteRequest) Reset()      { *m = PrivateCascadeDeleteRequest{} }
func (*PrivateCascadeDeleteRequest) ProtoMessage() {}
func (*PrivateCascadeDeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptorPrivateCustomapi, []int{2}
}

func (m *PrivateCascadeDeleteRequest) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *PrivateCascadeDeleteRequest) GetTenantName() string {
	if m != nil {
		return m.TenantName
	}
	return ""
}

func (m *PrivateCascadeDeleteRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func init() {
	proto.RegisterType((*LastLoginUpdateRequest)(nil), "ves.io.schema.user.LastLoginUpdateRequest")
	golang_proto.RegisterType((*LastLoginUpdateRequest)(nil), "ves.io.schema.user.LastLoginUpdateRequest")
	proto.RegisterType((*LastLoginUpdateResponse)(nil), "ves.io.schema.user.LastLoginUpdateResponse")
	golang_proto.RegisterType((*LastLoginUpdateResponse)(nil), "ves.io.schema.user.LastLoginUpdateResponse")
	proto.RegisterType((*PrivateCascadeDeleteRequest)(nil), "ves.io.schema.user.PrivateCascadeDeleteRequest")
	golang_proto.RegisterType((*PrivateCascadeDeleteRequest)(nil), "ves.io.schema.user.PrivateCascadeDeleteRequest")
}
func (this *LastLoginUpdateRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*LastLoginUpdateRequest)
	if !ok {
		that2, ok := that.(LastLoginUpdateRequest)
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
	if this.User != that1.User {
		return false
	}
	if this.Tenant != that1.Tenant {
		return false
	}
	if !this.LastLoginTimestamp.Equal(that1.LastLoginTimestamp) {
		return false
	}
	return true
}
func (this *LastLoginUpdateResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*LastLoginUpdateResponse)
	if !ok {
		that2, ok := that.(LastLoginUpdateResponse)
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
	return true
}
func (this *PrivateCascadeDeleteRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*PrivateCascadeDeleteRequest)
	if !ok {
		that2, ok := that.(PrivateCascadeDeleteRequest)
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
	if this.TenantName != that1.TenantName {
		return false
	}
	if this.Email != that1.Email {
		return false
	}
	return true
}
func (this *LastLoginUpdateRequest) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 7)
	s = append(s, "&user.LastLoginUpdateRequest{")
	s = append(s, "User: "+fmt.Sprintf("%#v", this.User)+",\n")
	s = append(s, "Tenant: "+fmt.Sprintf("%#v", this.Tenant)+",\n")
	if this.LastLoginTimestamp != nil {
		s = append(s, "LastLoginTimestamp: "+fmt.Sprintf("%#v", this.LastLoginTimestamp)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *LastLoginUpdateResponse) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 4)
	s = append(s, "&user.LastLoginUpdateResponse{")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *PrivateCascadeDeleteRequest) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 7)
	s = append(s, "&user.PrivateCascadeDeleteRequest{")
	s = append(s, "Namespace: "+fmt.Sprintf("%#v", this.Namespace)+",\n")
	s = append(s, "TenantName: "+fmt.Sprintf("%#v", this.TenantName)+",\n")
	s = append(s, "Email: "+fmt.Sprintf("%#v", this.Email)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringPrivateCustomapi(v interface{}, typ string) string {
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

// Client API for CustomPrivateAPI service

type CustomPrivateAPIClient interface {
	// Update Last Login
	//
	// x-displayName: "Update Last Login"
	// API to update last login timestamp of user
	UpdateLastLogin(ctx context.Context, in *LastLoginUpdateRequest, opts ...grpc.CallOption) (*LastLoginUpdateResponse, error)
	// CascadeDelete
	//
	// x-displayName: "Delete User and Related Objects"
	// CascadeDelete deletes the user and associated namespace roles for this user.
	// Use this only if the user and its referenced objects need to be wiped out altogether.
	// Note: users will always be in the system namespace.
	CascadeDelete(ctx context.Context, in *PrivateCascadeDeleteRequest, opts ...grpc.CallOption) (*CascadeDeleteResponse, error)
}

type customPrivateAPIClient struct {
	cc *grpc.ClientConn
}

func NewCustomPrivateAPIClient(cc *grpc.ClientConn) CustomPrivateAPIClient {
	return &customPrivateAPIClient{cc}
}

func (c *customPrivateAPIClient) UpdateLastLogin(ctx context.Context, in *LastLoginUpdateRequest, opts ...grpc.CallOption) (*LastLoginUpdateResponse, error) {
	out := new(LastLoginUpdateResponse)
	err := grpc.Invoke(ctx, "/ves.io.schema.user.CustomPrivateAPI/UpdateLastLogin", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customPrivateAPIClient) CascadeDelete(ctx context.Context, in *PrivateCascadeDeleteRequest, opts ...grpc.CallOption) (*CascadeDeleteResponse, error) {
	out := new(CascadeDeleteResponse)
	err := grpc.Invoke(ctx, "/ves.io.schema.user.CustomPrivateAPI/CascadeDelete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for CustomPrivateAPI service

type CustomPrivateAPIServer interface {
	// Update Last Login
	//
	// x-displayName: "Update Last Login"
	// API to update last login timestamp of user
	UpdateLastLogin(context.Context, *LastLoginUpdateRequest) (*LastLoginUpdateResponse, error)
	// CascadeDelete
	//
	// x-displayName: "Delete User and Related Objects"
	// CascadeDelete deletes the user and associated namespace roles for this user.
	// Use this only if the user and its referenced objects need to be wiped out altogether.
	// Note: users will always be in the system namespace.
	CascadeDelete(context.Context, *PrivateCascadeDeleteRequest) (*CascadeDeleteResponse, error)
}

func RegisterCustomPrivateAPIServer(s *grpc.Server, srv CustomPrivateAPIServer) {
	s.RegisterService(&_CustomPrivateAPI_serviceDesc, srv)
}

func _CustomPrivateAPI_UpdateLastLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LastLoginUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomPrivateAPIServer).UpdateLastLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ves.io.schema.user.CustomPrivateAPI/UpdateLastLogin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomPrivateAPIServer).UpdateLastLogin(ctx, req.(*LastLoginUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CustomPrivateAPI_CascadeDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PrivateCascadeDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomPrivateAPIServer).CascadeDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ves.io.schema.user.CustomPrivateAPI/CascadeDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomPrivateAPIServer).CascadeDelete(ctx, req.(*PrivateCascadeDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CustomPrivateAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ves.io.schema.user.CustomPrivateAPI",
	HandlerType: (*CustomPrivateAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateLastLogin",
			Handler:    _CustomPrivateAPI_UpdateLastLogin_Handler,
		},
		{
			MethodName: "CascadeDelete",
			Handler:    _CustomPrivateAPI_CascadeDelete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ves.io/schema/user/private_customapi.proto",
}

func (m *LastLoginUpdateRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LastLoginUpdateRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.User) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintPrivateCustomapi(dAtA, i, uint64(len(m.User)))
		i += copy(dAtA[i:], m.User)
	}
	if len(m.Tenant) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintPrivateCustomapi(dAtA, i, uint64(len(m.Tenant)))
		i += copy(dAtA[i:], m.Tenant)
	}
	if m.LastLoginTimestamp != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintPrivateCustomapi(dAtA, i, uint64(m.LastLoginTimestamp.Size()))
		n1, err := m.LastLoginTimestamp.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	return i, nil
}

func (m *LastLoginUpdateResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LastLoginUpdateResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func (m *PrivateCascadeDeleteRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PrivateCascadeDeleteRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Namespace) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintPrivateCustomapi(dAtA, i, uint64(len(m.Namespace)))
		i += copy(dAtA[i:], m.Namespace)
	}
	if len(m.TenantName) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintPrivateCustomapi(dAtA, i, uint64(len(m.TenantName)))
		i += copy(dAtA[i:], m.TenantName)
	}
	if len(m.Email) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintPrivateCustomapi(dAtA, i, uint64(len(m.Email)))
		i += copy(dAtA[i:], m.Email)
	}
	return i, nil
}

func encodeVarintPrivateCustomapi(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *LastLoginUpdateRequest) Size() (n int) {
	var l int
	_ = l
	l = len(m.User)
	if l > 0 {
		n += 1 + l + sovPrivateCustomapi(uint64(l))
	}
	l = len(m.Tenant)
	if l > 0 {
		n += 1 + l + sovPrivateCustomapi(uint64(l))
	}
	if m.LastLoginTimestamp != nil {
		l = m.LastLoginTimestamp.Size()
		n += 1 + l + sovPrivateCustomapi(uint64(l))
	}
	return n
}

func (m *LastLoginUpdateResponse) Size() (n int) {
	var l int
	_ = l
	return n
}

func (m *PrivateCascadeDeleteRequest) Size() (n int) {
	var l int
	_ = l
	l = len(m.Namespace)
	if l > 0 {
		n += 1 + l + sovPrivateCustomapi(uint64(l))
	}
	l = len(m.TenantName)
	if l > 0 {
		n += 1 + l + sovPrivateCustomapi(uint64(l))
	}
	l = len(m.Email)
	if l > 0 {
		n += 1 + l + sovPrivateCustomapi(uint64(l))
	}
	return n
}

func sovPrivateCustomapi(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozPrivateCustomapi(x uint64) (n int) {
	return sovPrivateCustomapi(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *LastLoginUpdateRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&LastLoginUpdateRequest{`,
		`User:` + fmt.Sprintf("%v", this.User) + `,`,
		`Tenant:` + fmt.Sprintf("%v", this.Tenant) + `,`,
		`LastLoginTimestamp:` + strings.Replace(fmt.Sprintf("%v", this.LastLoginTimestamp), "Timestamp", "google_protobuf1.Timestamp", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *LastLoginUpdateResponse) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&LastLoginUpdateResponse{`,
		`}`,
	}, "")
	return s
}
func (this *PrivateCascadeDeleteRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&PrivateCascadeDeleteRequest{`,
		`Namespace:` + fmt.Sprintf("%v", this.Namespace) + `,`,
		`TenantName:` + fmt.Sprintf("%v", this.TenantName) + `,`,
		`Email:` + fmt.Sprintf("%v", this.Email) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringPrivateCustomapi(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *LastLoginUpdateRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPrivateCustomapi
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
			return fmt.Errorf("proto: LastLoginUpdateRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LastLoginUpdateRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field User", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPrivateCustomapi
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
				return ErrInvalidLengthPrivateCustomapi
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.User = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tenant", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPrivateCustomapi
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
				return ErrInvalidLengthPrivateCustomapi
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Tenant = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastLoginTimestamp", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPrivateCustomapi
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
				return ErrInvalidLengthPrivateCustomapi
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.LastLoginTimestamp == nil {
				m.LastLoginTimestamp = &google_protobuf1.Timestamp{}
			}
			if err := m.LastLoginTimestamp.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPrivateCustomapi(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPrivateCustomapi
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
func (m *LastLoginUpdateResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPrivateCustomapi
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
			return fmt.Errorf("proto: LastLoginUpdateResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LastLoginUpdateResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipPrivateCustomapi(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPrivateCustomapi
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
func (m *PrivateCascadeDeleteRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPrivateCustomapi
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
			return fmt.Errorf("proto: PrivateCascadeDeleteRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PrivateCascadeDeleteRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Namespace", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPrivateCustomapi
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
				return ErrInvalidLengthPrivateCustomapi
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Namespace = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TenantName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPrivateCustomapi
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
				return ErrInvalidLengthPrivateCustomapi
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TenantName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Email", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPrivateCustomapi
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
				return ErrInvalidLengthPrivateCustomapi
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Email = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPrivateCustomapi(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPrivateCustomapi
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
func skipPrivateCustomapi(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPrivateCustomapi
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
					return 0, ErrIntOverflowPrivateCustomapi
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
					return 0, ErrIntOverflowPrivateCustomapi
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
				return 0, ErrInvalidLengthPrivateCustomapi
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowPrivateCustomapi
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
				next, err := skipPrivateCustomapi(dAtA[start:])
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
	ErrInvalidLengthPrivateCustomapi = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPrivateCustomapi   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("ves.io/schema/user/private_customapi.proto", fileDescriptorPrivateCustomapi)
}
func init() {
	golang_proto.RegisterFile("ves.io/schema/user/private_customapi.proto", fileDescriptorPrivateCustomapi)
}

var fileDescriptorPrivateCustomapi = []byte{
	// 619 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x53, 0xcf, 0x4f, 0xd4, 0x40,
	0x14, 0xde, 0x41, 0x24, 0x61, 0x88, 0xd1, 0x4c, 0x08, 0x42, 0x21, 0x03, 0xee, 0x09, 0xd6, 0xd0,
	0x49, 0xf0, 0x24, 0x37, 0xc5, 0x83, 0x46, 0xa2, 0x64, 0xa3, 0xc6, 0x90, 0x98, 0xcd, 0xb4, 0xfb,
	0x28, 0xa3, 0x6d, 0x67, 0xec, 0x4c, 0x57, 0xd1, 0x98, 0x18, 0xff, 0x01, 0x7f, 0x9d, 0xbd, 0x79,
	0xf0, 0x7f, 0xe0, 0x42, 0xe2, 0xc5, 0x23, 0xd1, 0x8b, 0x47, 0x29, 0x1e, 0x3c, 0xf2, 0x27, 0x98,
	0x9d, 0x76, 0xbb, 0xbb, 0xd0, 0xa8, 0x89, 0xb7, 0x99, 0xf7, 0xbe, 0x7e, 0xef, 0xfb, 0xbe, 0xd7,
	0xc1, 0x8d, 0x0e, 0x68, 0x57, 0x48, 0xa6, 0xfd, 0x6d, 0x88, 0x38, 0x4b, 0x35, 0x24, 0x4c, 0x25,
	0xa2, 0xc3, 0x0d, 0xb4, 0xfc, 0x54, 0x1b, 0x19, 0x71, 0x25, 0x5c, 0x95, 0x48, 0x23, 0x09, 0xc9,
	0xb1, 0x6e, 0x8e, 0x75, 0xbb, 0x58, 0x67, 0x39, 0x10, 0x66, 0x3b, 0xf5, 0x5c, 0x5f, 0x46, 0x2c,
	0x90, 0x81, 0x64, 0x16, 0xea, 0xa5, 0x5b, 0xf6, 0x66, 0x2f, 0xf6, 0x94, 0x53, 0x38, 0x73, 0x81,
	0x94, 0x41, 0x08, 0x8c, 0x2b, 0xc1, 0x78, 0x1c, 0x4b, 0xc3, 0x8d, 0x90, 0xb1, 0x2e, 0xba, 0xf3,
	0x45, 0xb7, 0xe4, 0x30, 0x22, 0x02, 0x6d, 0x78, 0xa4, 0x0a, 0xc0, 0xec, 0xb0, 0x5a, 0xa9, 0x06,
	0xbf, 0x9e, 0x19, 0x6e, 0x9a, 0x1d, 0x05, 0x25, 0x71, 0x85, 0x4b, 0xe9, 0x3d, 0x04, 0xdf, 0x14,
	0x80, 0xa5, 0xaa, 0x18, 0x52, 0x2f, 0x14, 0xfe, 0xf1, 0x14, 0x1c, 0x5a, 0x01, 0x1d, 0x98, 0x55,
	0x7f, 0x87, 0xf0, 0xd4, 0x3a, 0xd7, 0x66, 0x5d, 0x06, 0x22, 0xbe, 0xab, 0xda, 0xdc, 0x40, 0x13,
	0x1e, 0xa7, 0xa0, 0x0d, 0x21, 0x78, 0xb4, 0x0b, 0x9f, 0x46, 0x0b, 0x68, 0x71, 0xbc, 0x69, 0xcf,
	0x64, 0x0a, 0x8f, 0x19, 0x88, 0x79, 0x6c, 0xa6, 0x47, 0x6c, 0xb5, 0xb8, 0x91, 0x75, 0x3c, 0x19,
	0x72, 0x6d, 0x5a, 0x61, 0x97, 0xa6, 0x55, 0x06, 0x31, 0x7d, 0x6a, 0x01, 0x2d, 0x4e, 0xac, 0x38,
	0x6e, 0x1e, 0x95, 0xdb, 0x8b, 0xca, 0xbd, 0xd3, 0x43, 0x34, 0x49, 0xd8, 0x9b, 0x5e, 0xd6, 0xea,
	0x33, 0xf8, 0xfc, 0x09, 0x4d, 0x5a, 0xc9, 0x58, 0x43, 0x3d, 0xc1, 0xb3, 0x1b, 0xf9, 0xc2, 0xd7,
	0xb8, 0xf6, 0x79, 0x1b, 0xae, 0x41, 0x08, 0x7d, 0xcd, 0x73, 0x78, 0x3c, 0xe6, 0x11, 0x68, 0xc5,
	0x7d, 0x28, 0x84, 0xf7, 0x0b, 0x64, 0x1e, 0x4f, 0xe4, 0x7a, 0x5b, 0xdd, 0x5a, 0x61, 0x01, 0xe7,
	0xa5, 0x5b, 0x3c, 0x02, 0x32, 0x89, 0x4f, 0x43, 0xc4, 0x45, 0x68, 0x75, 0x8f, 0x37, 0xf3, 0xcb,
	0xca, 0xeb, 0x51, 0x7c, 0x6e, 0xcd, 0xe6, 0x5a, 0x8c, 0xbe, 0xb2, 0x71, 0x83, 0xbc, 0x1d, 0xc1,
	0x67, 0x73, 0x6d, 0xa5, 0x54, 0xd2, 0x70, 0x4f, 0xfe, 0x73, 0x6e, 0x75, 0xba, 0xce, 0xc5, 0x7f,
	0xc2, 0x16, 0xae, 0x3f, 0xa2, 0x57, 0xdf, 0x7e, 0xbe, 0x1f, 0xf9, 0x80, 0xea, 0x97, 0x7b, 0xff,
	0x3b, 0xcb, 0x37, 0xcd, 0x4a, 0x77, 0x9a, 0xe9, 0x1d, 0x6d, 0x20, 0xb2, 0x3b, 0xd6, 0x2c, 0xb5,
	0x0c, 0xad, 0xfe, 0x56, 0x56, 0x51, 0x63, 0xf3, 0x41, 0xfd, 0x3e, 0x1b, 0x1a, 0xcb, 0x44, 0x6c,
	0x12, 0xa9, 0x15, 0xf8, 0x86, 0x3d, 0x49, 0x84, 0x81, 0xff, 0xa1, 0x27, 0x9f, 0x11, 0x3e, 0x33,
	0xb4, 0x16, 0xc2, 0xaa, 0x5c, 0xfe, 0x61, 0x81, 0xce, 0x52, 0xd5, 0x07, 0xc7, 0x90, 0x45, 0x28,
	0xf7, 0x6c, 0x26, 0x1b, 0xf5, 0x9b, 0x7f, 0xb3, 0x34, 0xe0, 0xe1, 0x79, 0x79, 0x7e, 0x51, 0x18,
	0xf1, 0x73, 0xee, 0x56, 0xdb, 0x92, 0xaf, 0xa2, 0x86, 0x73, 0x61, 0x6f, 0x17, 0x8d, 0x7e, 0xdd,
	0x45, 0x33, 0x15, 0x4a, 0x6e, 0xdb, 0x67, 0x78, 0xf5, 0xd9, 0xfe, 0x01, 0xad, 0x7d, 0x3f, 0xa0,
	0xb5, 0xa3, 0x03, 0x8a, 0x5e, 0x66, 0x14, 0x7d, 0xca, 0x28, 0xfa, 0x92, 0x51, 0xb4, 0x9f, 0x51,
	0xf4, 0x23, 0xa3, 0xe8, 0x57, 0x46, 0x6b, 0x47, 0x19, 0x45, 0x6f, 0x0e, 0x69, 0x6d, 0xef, 0x90,
	0xa2, 0xcd, 0xeb, 0x81, 0x54, 0x8f, 0x02, 0xb7, 0x23, 0x43, 0x03, 0x49, 0xd2, 0x65, 0x63, 0xf6,
	0xb0, 0x25, 0x93, 0x68, 0x59, 0x25, 0xb2, 0x23, 0xda, 0x90, 0x2c, 0xf7, 0xda, 0x4c, 0x79, 0x81,
	0x64, 0xf0, 0xd4, 0x14, 0x96, 0x06, 0x5e, 0xaf, 0x37, 0x66, 0x1f, 0xd1, 0xa5, 0xdf, 0x01, 0x00,
	0x00, 0xff, 0xff, 0x10, 0x43, 0x60, 0x62, 0x0c, 0x05, 0x00, 0x00,
}
