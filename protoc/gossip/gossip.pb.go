// Code generated by protoc-gen-go. DO NOT EDIT.
// source: gossip.proto

package gossip

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
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

type EndorsedTx struct {
	Sign                 []string `protobuf:"bytes,1,rep,name=Sign,proto3" json:"Sign,omitempty"`
	Payload              []byte   `protobuf:"bytes,2,opt,name=Payload,proto3" json:"Payload,omitempty"`
	TimeStamp            string   `protobuf:"bytes,3,opt,name=timeStamp,proto3" json:"timeStamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EndorsedTx) Reset()         { *m = EndorsedTx{} }
func (m *EndorsedTx) String() string { return proto.CompactTextString(m) }
func (*EndorsedTx) ProtoMessage()    {}
func (*EndorsedTx) Descriptor() ([]byte, []int) {
	return fileDescriptor_878fa4887b90140c, []int{0}
}

func (m *EndorsedTx) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EndorsedTx.Unmarshal(m, b)
}
func (m *EndorsedTx) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EndorsedTx.Marshal(b, m, deterministic)
}
func (m *EndorsedTx) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EndorsedTx.Merge(m, src)
}
func (m *EndorsedTx) XXX_Size() int {
	return xxx_messageInfo_EndorsedTx.Size(m)
}
func (m *EndorsedTx) XXX_DiscardUnknown() {
	xxx_messageInfo_EndorsedTx.DiscardUnknown(m)
}

var xxx_messageInfo_EndorsedTx proto.InternalMessageInfo

func (m *EndorsedTx) GetSign() []string {
	if m != nil {
		return m.Sign
	}
	return nil
}

func (m *EndorsedTx) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *EndorsedTx) GetTimeStamp() string {
	if m != nil {
		return m.TimeStamp
	}
	return ""
}

type Block struct {
	Bunch                []*EndorsedTx `protobuf:"bytes,1,rep,name=bunch,proto3" json:"bunch,omitempty"`
	Sign                 string        `protobuf:"bytes,2,opt,name=Sign,proto3" json:"Sign,omitempty"`
	OffSet               string        `protobuf:"bytes,3,opt,name=OffSet,proto3" json:"OffSet,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Block) Reset()         { *m = Block{} }
func (m *Block) String() string { return proto.CompactTextString(m) }
func (*Block) ProtoMessage()    {}
func (*Block) Descriptor() ([]byte, []int) {
	return fileDescriptor_878fa4887b90140c, []int{1}
}

func (m *Block) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Block.Unmarshal(m, b)
}
func (m *Block) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Block.Marshal(b, m, deterministic)
}
func (m *Block) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Block.Merge(m, src)
}
func (m *Block) XXX_Size() int {
	return xxx_messageInfo_Block.Size(m)
}
func (m *Block) XXX_DiscardUnknown() {
	xxx_messageInfo_Block.DiscardUnknown(m)
}

var xxx_messageInfo_Block proto.InternalMessageInfo

func (m *Block) GetBunch() []*EndorsedTx {
	if m != nil {
		return m.Bunch
	}
	return nil
}

func (m *Block) GetSign() string {
	if m != nil {
		return m.Sign
	}
	return ""
}

func (m *Block) GetOffSet() string {
	if m != nil {
		return m.OffSet
	}
	return ""
}

func init() {
	proto.RegisterType((*EndorsedTx)(nil), "gossip.EndorsedTx")
	proto.RegisterType((*Block)(nil), "gossip.Block")
}

func init() { proto.RegisterFile("gossip.proto", fileDescriptor_878fa4887b90140c) }

var fileDescriptor_878fa4887b90140c = []byte{
	// 230 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0x8f, 0x31, 0x4f, 0xc3, 0x30,
	0x10, 0x85, 0x49, 0x4b, 0x53, 0xe5, 0x28, 0xcb, 0x0d, 0x55, 0x54, 0x18, 0xa2, 0x4c, 0x99, 0x5c,
	0x29, 0xac, 0x4c, 0x88, 0x8a, 0x11, 0xe4, 0x30, 0xb0, 0x30, 0x24, 0x8d, 0x63, 0x2c, 0x9c, 0x9c,
	0x95, 0xb8, 0x88, 0xfe, 0x7b, 0x84, 0x9d, 0x34, 0x9b, 0xdf, 0xd3, 0xe9, 0xfb, 0xfc, 0x60, 0x23,
	0x69, 0x18, 0x94, 0x61, 0xa6, 0x27, 0x4b, 0x18, 0xfa, 0xb4, 0xbb, 0x93, 0x44, 0x52, 0x8b, 0xbd,
	0x6b, 0xab, 0x53, 0xb3, 0x17, 0xad, 0xb1, 0x67, 0x7f, 0x94, 0x7e, 0x00, 0x1c, 0xba, 0x9a, 0xfa,
	0x41, 0xd4, 0xef, 0xbf, 0x88, 0x70, 0x5d, 0x28, 0xd9, 0xc5, 0x41, 0xb2, 0xcc, 0x22, 0xee, 0xde,
	0x18, 0xc3, 0xfa, 0xad, 0x3c, 0x6b, 0x2a, 0xeb, 0x78, 0x91, 0x04, 0xd9, 0x86, 0x4f, 0x11, 0xef,
	0x21, 0xb2, 0xaa, 0x15, 0x85, 0x2d, 0x5b, 0x13, 0x2f, 0x93, 0x20, 0x8b, 0xf8, 0x5c, 0xa4, 0x9f,
	0xb0, 0x7a, 0xd2, 0x74, 0xfc, 0xc6, 0x0c, 0x56, 0xd5, 0xa9, 0x3b, 0x7e, 0x39, 0xea, 0x4d, 0x8e,
	0x6c, 0xfc, 0xe5, 0xec, 0xe5, 0xfe, 0xe0, 0xa2, 0x5f, 0x38, 0x96, 0xd7, 0x6f, 0x21, 0x7c, 0x6d,
	0x9a, 0x42, 0xd8, 0xd1, 0x30, 0xa6, 0xfc, 0x11, 0xc2, 0x17, 0xc7, 0xc1, 0x1c, 0xd6, 0xcf, 0x42,
	0xab, 0x1f, 0xd1, 0xe3, 0xed, 0xc4, 0x76, 0xe6, 0xdd, 0x96, 0xf9, 0xe9, 0x6c, 0x9a, 0xce, 0x0e,
	0xff, 0xd3, 0xd3, 0xab, 0x2a, 0x74, 0xcd, 0xc3, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x32, 0x6d,
	0x9f, 0x28, 0x32, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GossipClient is the client API for Gossip service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GossipClient interface {
	Deliver(ctx context.Context, in *Block, opts ...grpc.CallOption) (*empty.Empty, error)
}

type gossipClient struct {
	cc *grpc.ClientConn
}

func NewGossipClient(cc *grpc.ClientConn) GossipClient {
	return &gossipClient{cc}
}

func (c *gossipClient) Deliver(ctx context.Context, in *Block, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/gossip.Gossip/Deliver", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GossipServer is the server API for Gossip service.
type GossipServer interface {
	Deliver(context.Context, *Block) (*empty.Empty, error)
}

func RegisterGossipServer(s *grpc.Server, srv GossipServer) {
	s.RegisterService(&_Gossip_serviceDesc, srv)
}

func _Gossip_Deliver_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Block)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GossipServer).Deliver(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gossip.Gossip/Deliver",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GossipServer).Deliver(ctx, req.(*Block))
	}
	return interceptor(ctx, in, info, handler)
}

var _Gossip_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gossip.Gossip",
	HandlerType: (*GossipServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Deliver",
			Handler:    _Gossip_Deliver_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gossip.proto",
}
