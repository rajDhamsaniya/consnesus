// Code generated by protoc-gen-go. DO NOT EDIT.
// source: orderer.proto

package orderer

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
	return fileDescriptor_00d11f8df639b0fc, []int{0}
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

type BlockId struct {
	OffSet               string   `protobuf:"bytes,1,opt,name=OffSet,proto3" json:"OffSet,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BlockId) Reset()         { *m = BlockId{} }
func (m *BlockId) String() string { return proto.CompactTextString(m) }
func (*BlockId) ProtoMessage()    {}
func (*BlockId) Descriptor() ([]byte, []int) {
	return fileDescriptor_00d11f8df639b0fc, []int{1}
}

func (m *BlockId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlockId.Unmarshal(m, b)
}
func (m *BlockId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlockId.Marshal(b, m, deterministic)
}
func (m *BlockId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockId.Merge(m, src)
}
func (m *BlockId) XXX_Size() int {
	return xxx_messageInfo_BlockId.Size(m)
}
func (m *BlockId) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockId.DiscardUnknown(m)
}

var xxx_messageInfo_BlockId proto.InternalMessageInfo

func (m *BlockId) GetOffSet() string {
	if m != nil {
		return m.OffSet
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
	return fileDescriptor_00d11f8df639b0fc, []int{2}
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
	proto.RegisterType((*EndorsedTx)(nil), "orderer.EndorsedTx")
	proto.RegisterType((*BlockId)(nil), "orderer.BlockId")
	proto.RegisterType((*Block)(nil), "orderer.Block")
}

func init() { proto.RegisterFile("orderer.proto", fileDescriptor_00d11f8df639b0fc) }

var fileDescriptor_00d11f8df639b0fc = []byte{
	// 287 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x50, 0x4d, 0x4b, 0xc3, 0x40,
	0x14, 0xec, 0xb6, 0xb6, 0x31, 0xcf, 0x0f, 0xca, 0x13, 0x4a, 0x88, 0x1e, 0x62, 0x4e, 0xf1, 0x92,
	0x6a, 0x05, 0xc1, 0xab, 0x50, 0x8a, 0xa7, 0x4a, 0xd2, 0x83, 0x27, 0x21, 0x1f, 0x2f, 0x31, 0x98,
	0x64, 0x43, 0xba, 0x81, 0xf6, 0x57, 0xf9, 0x17, 0xc5, 0xdd, 0xa4, 0x51, 0xf0, 0xe0, 0x6d, 0x67,
	0x98, 0x9d, 0x37, 0x33, 0x70, 0xc6, 0xeb, 0x98, 0x6a, 0xaa, 0xdd, 0xaa, 0xe6, 0x82, 0xa3, 0xd6,
	0x42, 0xf3, 0x32, 0xe5, 0x3c, 0xcd, 0x69, 0x2e, 0xe9, 0xb0, 0x49, 0xe6, 0x54, 0x54, 0x62, 0xaf,
	0x54, 0xf6, 0x2b, 0xc0, 0xb2, 0x8c, 0x79, 0xbd, 0xa5, 0x78, 0xb3, 0x43, 0x84, 0x23, 0x3f, 0x4b,
	0x4b, 0x83, 0x59, 0x23, 0x47, 0xf7, 0xe4, 0x1b, 0x0d, 0xd0, 0x5e, 0x82, 0x7d, 0xce, 0x83, 0xd8,
	0x18, 0x5a, 0xcc, 0x39, 0xf5, 0x3a, 0x88, 0x57, 0xa0, 0x8b, 0xac, 0x20, 0x5f, 0x04, 0x45, 0x65,
	0x8c, 0x2c, 0xe6, 0xe8, 0x5e, 0x4f, 0xd8, 0xd7, 0xa0, 0x3d, 0xe5, 0x3c, 0xfa, 0x78, 0x8e, 0x71,
	0x06, 0x93, 0x75, 0x92, 0xf8, 0x24, 0x0c, 0x26, 0x55, 0x2d, 0xb2, 0xdf, 0x60, 0x2c, 0x25, 0x78,
	0x03, 0xe3, 0xb0, 0x29, 0xa3, 0x77, 0x79, 0xf8, 0x64, 0x71, 0xe1, 0x76, 0x55, 0xfa, 0x6c, 0x9e,
	0x52, 0x1c, 0x22, 0x0e, 0xa5, 0x93, 0x8a, 0xd8, 0xfb, 0x8f, 0x7e, 0xfa, 0x2f, 0x3e, 0x19, 0x68,
	0x6b, 0xe5, 0x84, 0x8f, 0x70, 0xec, 0x37, 0x61, 0x91, 0x89, 0xcd, 0x0e, 0xff, 0xf2, 0x37, 0x67,
	0xae, 0xda, 0xc9, 0xed, 0x76, 0x72, 0x97, 0xdf, 0x3b, 0xd9, 0x03, 0xbc, 0x03, 0x7d, 0x45, 0x42,
	0x26, 0xdd, 0xe2, 0xf4, 0xf0, 0xb7, 0x6d, 0x67, 0x9e, 0xff, 0x66, 0xec, 0xc1, 0x2d, 0xc3, 0x07,
	0x98, 0xae, 0x48, 0xf8, 0x15, 0x45, 0x59, 0x92, 0x45, 0xaa, 0xe4, 0x3f, 0x7e, 0x86, 0x13, 0x79,
	0xfc, 0xfe, 0x2b, 0x00, 0x00, 0xff, 0xff, 0xbc, 0xd2, 0xf7, 0x51, 0xcc, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// OrdererClient is the client API for Orderer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type OrdererClient interface {
	SubmitTx(ctx context.Context, in *EndorsedTx, opts ...grpc.CallOption) (*empty.Empty, error)
	GetBlocks(ctx context.Context, in *BlockId, opts ...grpc.CallOption) (Orderer_GetBlocksClient, error)
	GetSpecificBlock(ctx context.Context, in *BlockId, opts ...grpc.CallOption) (*Block, error)
}

type ordererClient struct {
	cc *grpc.ClientConn
}

func NewOrdererClient(cc *grpc.ClientConn) OrdererClient {
	return &ordererClient{cc}
}

func (c *ordererClient) SubmitTx(ctx context.Context, in *EndorsedTx, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/orderer.Orderer/SubmitTx", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ordererClient) GetBlocks(ctx context.Context, in *BlockId, opts ...grpc.CallOption) (Orderer_GetBlocksClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Orderer_serviceDesc.Streams[0], "/orderer.Orderer/GetBlocks", opts...)
	if err != nil {
		return nil, err
	}
	x := &ordererGetBlocksClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Orderer_GetBlocksClient interface {
	Recv() (*Block, error)
	grpc.ClientStream
}

type ordererGetBlocksClient struct {
	grpc.ClientStream
}

func (x *ordererGetBlocksClient) Recv() (*Block, error) {
	m := new(Block)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *ordererClient) GetSpecificBlock(ctx context.Context, in *BlockId, opts ...grpc.CallOption) (*Block, error) {
	out := new(Block)
	err := c.cc.Invoke(ctx, "/orderer.Orderer/GetSpecificBlock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrdererServer is the server API for Orderer service.
type OrdererServer interface {
	SubmitTx(context.Context, *EndorsedTx) (*empty.Empty, error)
	GetBlocks(*BlockId, Orderer_GetBlocksServer) error
	GetSpecificBlock(context.Context, *BlockId) (*Block, error)
}

func RegisterOrdererServer(s *grpc.Server, srv OrdererServer) {
	s.RegisterService(&_Orderer_serviceDesc, srv)
}

func _Orderer_SubmitTx_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EndorsedTx)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrdererServer).SubmitTx(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/orderer.Orderer/SubmitTx",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrdererServer).SubmitTx(ctx, req.(*EndorsedTx))
	}
	return interceptor(ctx, in, info, handler)
}

func _Orderer_GetBlocks_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(BlockId)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(OrdererServer).GetBlocks(m, &ordererGetBlocksServer{stream})
}

type Orderer_GetBlocksServer interface {
	Send(*Block) error
	grpc.ServerStream
}

type ordererGetBlocksServer struct {
	grpc.ServerStream
}

func (x *ordererGetBlocksServer) Send(m *Block) error {
	return x.ServerStream.SendMsg(m)
}

func _Orderer_GetSpecificBlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BlockId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrdererServer).GetSpecificBlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/orderer.Orderer/GetSpecificBlock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrdererServer).GetSpecificBlock(ctx, req.(*BlockId))
	}
	return interceptor(ctx, in, info, handler)
}

var _Orderer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "orderer.Orderer",
	HandlerType: (*OrdererServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SubmitTx",
			Handler:    _Orderer_SubmitTx_Handler,
		},
		{
			MethodName: "GetSpecificBlock",
			Handler:    _Orderer_GetSpecificBlock_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetBlocks",
			Handler:       _Orderer_GetBlocks_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "orderer.proto",
}