// Code generated by protoc-gen-go. DO NOT EDIT.
// source: registry.proto

package registry

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// we define what a registration needs to have
// all optional and self explanatory
type Registration struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Ipv4                 string   `protobuf:"bytes,2,opt,name=ipv4,proto3" json:"ipv4,omitempty"`
	Port                 string   `protobuf:"bytes,3,opt,name=port,proto3" json:"port,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Registration) Reset()         { *m = Registration{} }
func (m *Registration) String() string { return proto.CompactTextString(m) }
func (*Registration) ProtoMessage()    {}
func (*Registration) Descriptor() ([]byte, []int) {
	return fileDescriptor_41af05d40a615591, []int{0}
}

func (m *Registration) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Registration.Unmarshal(m, b)
}
func (m *Registration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Registration.Marshal(b, m, deterministic)
}
func (m *Registration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Registration.Merge(m, src)
}
func (m *Registration) XXX_Size() int {
	return xxx_messageInfo_Registration.Size(m)
}
func (m *Registration) XXX_DiscardUnknown() {
	xxx_messageInfo_Registration.DiscardUnknown(m)
}

var xxx_messageInfo_Registration proto.InternalMessageInfo

func (m *Registration) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Registration) GetIpv4() string {
	if m != nil {
		return m.Ipv4
	}
	return ""
}

func (m *Registration) GetPort() string {
	if m != nil {
		return m.Port
	}
	return ""
}

// a list/array of registrations
type RegistrationList struct {
	Registrations        []*Registration `protobuf:"bytes,1,rep,name=registrations,proto3" json:"registrations,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *RegistrationList) Reset()         { *m = RegistrationList{} }
func (m *RegistrationList) String() string { return proto.CompactTextString(m) }
func (*RegistrationList) ProtoMessage()    {}
func (*RegistrationList) Descriptor() ([]byte, []int) {
	return fileDescriptor_41af05d40a615591, []int{1}
}

func (m *RegistrationList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegistrationList.Unmarshal(m, b)
}
func (m *RegistrationList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegistrationList.Marshal(b, m, deterministic)
}
func (m *RegistrationList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegistrationList.Merge(m, src)
}
func (m *RegistrationList) XXX_Size() int {
	return xxx_messageInfo_RegistrationList.Size(m)
}
func (m *RegistrationList) XXX_DiscardUnknown() {
	xxx_messageInfo_RegistrationList.DiscardUnknown(m)
}

var xxx_messageInfo_RegistrationList proto.InternalMessageInfo

func (m *RegistrationList) GetRegistrations() []*Registration {
	if m != nil {
		return m.Registrations
	}
	return nil
}

// w/ this message we define what the server needs
// to know what you want to have
type RegistrationFetchRequest struct {
	Registrations        []*Registration `protobuf:"bytes,1,rep,name=registrations,proto3" json:"registrations,omitempty"`
	FetchAll             bool            `protobuf:"varint,2,opt,name=fetchAll,proto3" json:"fetchAll,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *RegistrationFetchRequest) Reset()         { *m = RegistrationFetchRequest{} }
func (m *RegistrationFetchRequest) String() string { return proto.CompactTextString(m) }
func (*RegistrationFetchRequest) ProtoMessage()    {}
func (*RegistrationFetchRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_41af05d40a615591, []int{2}
}

func (m *RegistrationFetchRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegistrationFetchRequest.Unmarshal(m, b)
}
func (m *RegistrationFetchRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegistrationFetchRequest.Marshal(b, m, deterministic)
}
func (m *RegistrationFetchRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegistrationFetchRequest.Merge(m, src)
}
func (m *RegistrationFetchRequest) XXX_Size() int {
	return xxx_messageInfo_RegistrationFetchRequest.Size(m)
}
func (m *RegistrationFetchRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RegistrationFetchRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RegistrationFetchRequest proto.InternalMessageInfo

func (m *RegistrationFetchRequest) GetRegistrations() []*Registration {
	if m != nil {
		return m.Registrations
	}
	return nil
}

func (m *RegistrationFetchRequest) GetFetchAll() bool {
	if m != nil {
		return m.FetchAll
	}
	return false
}

// simple response when you sign up
type RegisterResponse struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterResponse) Reset()         { *m = RegisterResponse{} }
func (m *RegisterResponse) String() string { return proto.CompactTextString(m) }
func (*RegisterResponse) ProtoMessage()    {}
func (*RegisterResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_41af05d40a615591, []int{3}
}

func (m *RegisterResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterResponse.Unmarshal(m, b)
}
func (m *RegisterResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterResponse.Marshal(b, m, deterministic)
}
func (m *RegisterResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterResponse.Merge(m, src)
}
func (m *RegisterResponse) XXX_Size() int {
	return xxx_messageInfo_RegisterResponse.Size(m)
}
func (m *RegisterResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterResponse proto.InternalMessageInfo

func (m *RegisterResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*Registration)(nil), "registry.Registration")
	proto.RegisterType((*RegistrationList)(nil), "registry.RegistrationList")
	proto.RegisterType((*RegistrationFetchRequest)(nil), "registry.RegistrationFetchRequest")
	proto.RegisterType((*RegisterResponse)(nil), "registry.RegisterResponse")
}

func init() { proto.RegisterFile("registry.proto", fileDescriptor_41af05d40a615591) }

var fileDescriptor_41af05d40a615591 = []byte{
	// 258 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x92, 0x4b, 0x4b, 0x03, 0x31,
	0x10, 0xc7, 0x59, 0x2b, 0xba, 0x8e, 0x0f, 0x64, 0x10, 0x09, 0x7b, 0x2a, 0x39, 0xf5, 0x20, 0x3d,
	0x54, 0x8f, 0x22, 0x7a, 0xf1, 0x20, 0x3d, 0x48, 0x44, 0xef, 0xeb, 0x32, 0xd6, 0x40, 0xbb, 0x59,
	0x27, 0xb1, 0xd0, 0xef, 0xeb, 0x07, 0x91, 0x24, 0xa6, 0xc6, 0x47, 0x4f, 0x7b, 0xfb, 0xcf, 0x3f,
	0xc9, 0x6f, 0x5e, 0x81, 0x23, 0xa6, 0x99, 0xb6, 0x8e, 0x57, 0xe3, 0x8e, 0x8d, 0x33, 0x58, 0xa6,
	0x58, 0xde, 0xc1, 0x81, 0x8a, 0xba, 0x76, 0xda, 0xb4, 0x88, 0xb0, 0xdd, 0xd6, 0x0b, 0x12, 0xc5,
	0xb0, 0x18, 0xed, 0xa9, 0xa0, 0xbd, 0xa7, 0xbb, 0xe5, 0x85, 0xd8, 0x8a, 0x9e, 0xd7, 0xde, 0xeb,
	0x0c, 0x3b, 0x31, 0x88, 0x9e, 0xd7, 0xf2, 0x1e, 0x8e, 0x73, 0xd6, 0x54, 0x5b, 0x87, 0x97, 0x70,
	0xc8, 0x99, 0x67, 0x45, 0x31, 0x1c, 0x8c, 0xf6, 0x27, 0xa7, 0xe3, 0x75, 0x45, 0xf9, 0x13, 0xf5,
	0xf3, 0xb2, 0x74, 0x20, 0xf2, 0xe3, 0x5b, 0x72, 0xcd, 0xab, 0xa2, 0xb7, 0x77, 0xea, 0x4b, 0xc6,
	0x0a, 0xca, 0x17, 0x4f, 0xbb, 0x99, 0xcf, 0x43, 0x5f, 0xa5, 0x5a, 0xc7, 0xf2, 0x2c, 0xf5, 0x41,
	0xac, 0xc8, 0x76, 0xa6, 0xb5, 0x84, 0x02, 0x76, 0x17, 0x64, 0x6d, 0x3d, 0x4b, 0xa3, 0x49, 0xe1,
	0xe4, 0xa3, 0x80, 0xf2, 0x2b, 0xd3, 0x0a, 0xaf, 0x92, 0x26, 0xc6, 0x0d, 0x95, 0x54, 0xd5, 0x6f,
	0x3f, 0x4b, 0x73, 0x0d, 0xf0, 0xd8, 0x72, 0x1f, 0xc2, 0x13, 0x9c, 0x84, 0x46, 0x1e, 0x88, 0x97,
	0xba, 0xa1, 0xa9, 0x69, 0xe2, 0x62, 0xe5, 0xff, 0xac, 0x7c, 0xa4, 0x7f, 0xb9, 0xdf, 0x8b, 0x7c,
	0xde, 0x09, 0x3f, 0xe7, 0xfc, 0x33, 0x00, 0x00, 0xff, 0xff, 0x8c, 0xa7, 0x6c, 0x79, 0x4b, 0x02,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RegistryClient is the client API for Registry service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RegistryClient interface {
	Register(ctx context.Context, in *Registration, opts ...grpc.CallOption) (*RegisterResponse, error)
	Unregister(ctx context.Context, in *Registration, opts ...grpc.CallOption) (*RegisterResponse, error)
	FetchServiceLocation(ctx context.Context, in *RegistrationFetchRequest, opts ...grpc.CallOption) (*RegistrationList, error)
}

type registryClient struct {
	cc *grpc.ClientConn
}

func NewRegistryClient(cc *grpc.ClientConn) RegistryClient {
	return &registryClient{cc}
}

func (c *registryClient) Register(ctx context.Context, in *Registration, opts ...grpc.CallOption) (*RegisterResponse, error) {
	out := new(RegisterResponse)
	err := c.cc.Invoke(ctx, "/registry.Registry/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registryClient) Unregister(ctx context.Context, in *Registration, opts ...grpc.CallOption) (*RegisterResponse, error) {
	out := new(RegisterResponse)
	err := c.cc.Invoke(ctx, "/registry.Registry/Unregister", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registryClient) FetchServiceLocation(ctx context.Context, in *RegistrationFetchRequest, opts ...grpc.CallOption) (*RegistrationList, error) {
	out := new(RegistrationList)
	err := c.cc.Invoke(ctx, "/registry.Registry/fetchServiceLocation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RegistryServer is the server API for Registry service.
type RegistryServer interface {
	Register(context.Context, *Registration) (*RegisterResponse, error)
	Unregister(context.Context, *Registration) (*RegisterResponse, error)
	FetchServiceLocation(context.Context, *RegistrationFetchRequest) (*RegistrationList, error)
}

func RegisterRegistryServer(s *grpc.Server, srv RegistryServer) {
	s.RegisterService(&_Registry_serviceDesc, srv)
}

func _Registry_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Registration)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistryServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/registry.Registry/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistryServer).Register(ctx, req.(*Registration))
	}
	return interceptor(ctx, in, info, handler)
}

func _Registry_Unregister_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Registration)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistryServer).Unregister(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/registry.Registry/Unregister",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistryServer).Unregister(ctx, req.(*Registration))
	}
	return interceptor(ctx, in, info, handler)
}

func _Registry_FetchServiceLocation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegistrationFetchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistryServer).FetchServiceLocation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/registry.Registry/FetchServiceLocation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistryServer).FetchServiceLocation(ctx, req.(*RegistrationFetchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Registry_serviceDesc = grpc.ServiceDesc{
	ServiceName: "registry.Registry",
	HandlerType: (*RegistryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _Registry_Register_Handler,
		},
		{
			MethodName: "Unregister",
			Handler:    _Registry_Unregister_Handler,
		},
		{
			MethodName: "fetchServiceLocation",
			Handler:    _Registry_FetchServiceLocation_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "registry.proto",
}
