// Code generated by protoc-gen-go. DO NOT EDIT.
// source: eventServer.proto

/*
Package serverproto is a generated protocol buffer package.

It is generated from these files:
	eventServer.proto

It has these top-level messages:
	ClientRegisterAddressReq
	ClientRegisterAddressRes
	ClientTransactionReq
	ClientTransactionRes
	ChainTranscationReq
	ChainTranscationRes
	ClientTransactionJavaReq
	ClientTransactionJavaRes
*/
package serverproto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// 客户端注册地址
type ClientRegisterAddressReq struct {
	AddRessIpReq   string `protobuf:"bytes,1,opt,name=AddRessIpReq" json:"AddRessIpReq,omitempty"`
	AddRessPortReq string `protobuf:"bytes,2,opt,name=AddRessPortReq" json:"AddRessPortReq,omitempty"`
	RemarkReq      string `protobuf:"bytes,3,opt,name=RemarkReq" json:"RemarkReq,omitempty"`
}

func (m *ClientRegisterAddressReq) Reset()                    { *m = ClientRegisterAddressReq{} }
func (m *ClientRegisterAddressReq) String() string            { return proto.CompactTextString(m) }
func (*ClientRegisterAddressReq) ProtoMessage()               {}
func (*ClientRegisterAddressReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ClientRegisterAddressReq) GetAddRessIpReq() string {
	if m != nil {
		return m.AddRessIpReq
	}
	return ""
}

func (m *ClientRegisterAddressReq) GetAddRessPortReq() string {
	if m != nil {
		return m.AddRessPortReq
	}
	return ""
}

func (m *ClientRegisterAddressReq) GetRemarkReq() string {
	if m != nil {
		return m.RemarkReq
	}
	return ""
}

// 服务端注册地址返回值
type ClientRegisterAddressRes struct {
	MessageRes   string `protobuf:"bytes,1,opt,name=MessageRes" json:"MessageRes,omitempty"`
	IsSuccess    bool   `protobuf:"varint,2,opt,name=IsSuccess" json:"IsSuccess,omitempty"`
	MessageIDRes string `protobuf:"bytes,3,opt,name=MessageIDRes" json:"MessageIDRes,omitempty"`
}

func (m *ClientRegisterAddressRes) Reset()                    { *m = ClientRegisterAddressRes{} }
func (m *ClientRegisterAddressRes) String() string            { return proto.CompactTextString(m) }
func (*ClientRegisterAddressRes) ProtoMessage()               {}
func (*ClientRegisterAddressRes) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ClientRegisterAddressRes) GetMessageRes() string {
	if m != nil {
		return m.MessageRes
	}
	return ""
}

func (m *ClientRegisterAddressRes) GetIsSuccess() bool {
	if m != nil {
		return m.IsSuccess
	}
	return false
}

func (m *ClientRegisterAddressRes) GetMessageIDRes() string {
	if m != nil {
		return m.MessageIDRes
	}
	return ""
}

// 客户端请求交易体
type ClientTransactionReq struct {
	TxidReq string `protobuf:"bytes,1,opt,name=TxidReq" json:"TxidReq,omitempty"`
}

func (m *ClientTransactionReq) Reset()                    { *m = ClientTransactionReq{} }
func (m *ClientTransactionReq) String() string            { return proto.CompactTextString(m) }
func (*ClientTransactionReq) ProtoMessage()               {}
func (*ClientTransactionReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ClientTransactionReq) GetTxidReq() string {
	if m != nil {
		return m.TxidReq
	}
	return ""
}

// 服务端请求交易返回值
type ClientTransactionRes struct {
	TxidRes    string `protobuf:"bytes,1,opt,name=TxidRes" json:"TxidRes,omitempty"`
	CodeRes    string `protobuf:"bytes,2,opt,name=CodeRes" json:"CodeRes,omitempty"`
	MessageRes string `protobuf:"bytes,3,opt,name=MessageRes" json:"MessageRes,omitempty"`
}

func (m *ClientTransactionRes) Reset()                    { *m = ClientTransactionRes{} }
func (m *ClientTransactionRes) String() string            { return proto.CompactTextString(m) }
func (*ClientTransactionRes) ProtoMessage()               {}
func (*ClientTransactionRes) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ClientTransactionRes) GetTxidRes() string {
	if m != nil {
		return m.TxidRes
	}
	return ""
}

func (m *ClientTransactionRes) GetCodeRes() string {
	if m != nil {
		return m.CodeRes
	}
	return ""
}

func (m *ClientTransactionRes) GetMessageRes() string {
	if m != nil {
		return m.MessageRes
	}
	return ""
}

// 链请求交易体
type ChainTranscationReq struct {
	BatchidReq string            `protobuf:"bytes,1,opt,name=BatchidReq" json:"BatchidReq,omitempty"`
	TxsReq     map[string]string `protobuf:"bytes,2,rep,name=TxsReq" json:"TxsReq,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *ChainTranscationReq) Reset()                    { *m = ChainTranscationReq{} }
func (m *ChainTranscationReq) String() string            { return proto.CompactTextString(m) }
func (*ChainTranscationReq) ProtoMessage()               {}
func (*ChainTranscationReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *ChainTranscationReq) GetBatchidReq() string {
	if m != nil {
		return m.BatchidReq
	}
	return ""
}

func (m *ChainTranscationReq) GetTxsReq() map[string]string {
	if m != nil {
		return m.TxsReq
	}
	return nil
}

// 服务端对链请求交易返回值
type ChainTranscationRes struct {
	BatchidReq string `protobuf:"bytes,1,opt,name=BatchidReq" json:"BatchidReq,omitempty"`
	IsSuccess  bool   `protobuf:"varint,2,opt,name=IsSuccess" json:"IsSuccess,omitempty"`
}

func (m *ChainTranscationRes) Reset()                    { *m = ChainTranscationRes{} }
func (m *ChainTranscationRes) String() string            { return proto.CompactTextString(m) }
func (*ChainTranscationRes) ProtoMessage()               {}
func (*ChainTranscationRes) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *ChainTranscationRes) GetBatchidReq() string {
	if m != nil {
		return m.BatchidReq
	}
	return ""
}

func (m *ChainTranscationRes) GetIsSuccess() bool {
	if m != nil {
		return m.IsSuccess
	}
	return false
}

// 服务端发送java请求参数
type ClientTransactionJavaReq struct {
	TxidRes    string `protobuf:"bytes,1,opt,name=TxidRes" json:"TxidRes,omitempty"`
	CodeRes    string `protobuf:"bytes,2,opt,name=CodeRes" json:"CodeRes,omitempty"`
	MessageRes string `protobuf:"bytes,3,opt,name=MessageRes" json:"MessageRes,omitempty"`
}

func (m *ClientTransactionJavaReq) Reset()                    { *m = ClientTransactionJavaReq{} }
func (m *ClientTransactionJavaReq) String() string            { return proto.CompactTextString(m) }
func (*ClientTransactionJavaReq) ProtoMessage()               {}
func (*ClientTransactionJavaReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *ClientTransactionJavaReq) GetTxidRes() string {
	if m != nil {
		return m.TxidRes
	}
	return ""
}

func (m *ClientTransactionJavaReq) GetCodeRes() string {
	if m != nil {
		return m.CodeRes
	}
	return ""
}

func (m *ClientTransactionJavaReq) GetMessageRes() string {
	if m != nil {
		return m.MessageRes
	}
	return ""
}

// go客户端收到java服务端返回值
type ClientTransactionJavaRes struct {
	TxidRes       string `protobuf:"bytes,1,opt,name=TxidRes" json:"TxidRes,omitempty"`
	IsReceivedRes string `protobuf:"bytes,2,opt,name=IsReceivedRes" json:"IsReceivedRes,omitempty"`
}

func (m *ClientTransactionJavaRes) Reset()                    { *m = ClientTransactionJavaRes{} }
func (m *ClientTransactionJavaRes) String() string            { return proto.CompactTextString(m) }
func (*ClientTransactionJavaRes) ProtoMessage()               {}
func (*ClientTransactionJavaRes) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *ClientTransactionJavaRes) GetTxidRes() string {
	if m != nil {
		return m.TxidRes
	}
	return ""
}

func (m *ClientTransactionJavaRes) GetIsReceivedRes() string {
	if m != nil {
		return m.IsReceivedRes
	}
	return ""
}

func init() {
	proto.RegisterType((*ClientRegisterAddressReq)(nil), "server.ClientRegisterAddressReq")
	proto.RegisterType((*ClientRegisterAddressRes)(nil), "server.ClientRegisterAddressRes")
	proto.RegisterType((*ClientTransactionReq)(nil), "server.ClientTransactionReq")
	proto.RegisterType((*ClientTransactionRes)(nil), "server.ClientTransactionRes")
	proto.RegisterType((*ChainTranscationReq)(nil), "server.ChainTranscationReq")
	proto.RegisterType((*ChainTranscationRes)(nil), "server.ChainTranscationRes")
	proto.RegisterType((*ClientTransactionJavaReq)(nil), "server.ClientTransactionJavaReq")
	proto.RegisterType((*ClientTransactionJavaRes)(nil), "server.ClientTransactionJavaRes")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for GoEventService service

type GoEventServiceClient interface {
	GoClientRegistEvent(ctx context.Context, in *ClientRegisterAddressReq, opts ...grpc.CallOption) (*ClientRegisterAddressRes, error)
	GoClientRequestEvent(ctx context.Context, in *ClientTransactionReq, opts ...grpc.CallOption) (*ClientTransactionRes, error)
	GoChainRequestEvent(ctx context.Context, opts ...grpc.CallOption) (GoEventService_GoChainRequestEventClient, error)
}

type goEventServiceClient struct {
	cc *grpc.ClientConn
}

func NewGoEventServiceClient(cc *grpc.ClientConn) GoEventServiceClient {
	return &goEventServiceClient{cc}
}

func (c *goEventServiceClient) GoClientRegistEvent(ctx context.Context, in *ClientRegisterAddressReq, opts ...grpc.CallOption) (*ClientRegisterAddressRes, error) {
	out := new(ClientRegisterAddressRes)
	err := grpc.Invoke(ctx, "/server.GoEventService/GoClientRegistEvent", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goEventServiceClient) GoClientRequestEvent(ctx context.Context, in *ClientTransactionReq, opts ...grpc.CallOption) (*ClientTransactionRes, error) {
	out := new(ClientTransactionRes)
	err := grpc.Invoke(ctx, "/server.GoEventService/GoClientRequestEvent", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goEventServiceClient) GoChainRequestEvent(ctx context.Context, opts ...grpc.CallOption) (GoEventService_GoChainRequestEventClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_GoEventService_serviceDesc.Streams[0], c.cc, "/server.GoEventService/GoChainRequestEvent", opts...)
	if err != nil {
		return nil, err
	}
	x := &goEventServiceGoChainRequestEventClient{stream}
	return x, nil
}

type GoEventService_GoChainRequestEventClient interface {
	Send(*ChainTranscationReq) error
	Recv() (*ChainTranscationRes, error)
	grpc.ClientStream
}

type goEventServiceGoChainRequestEventClient struct {
	grpc.ClientStream
}

func (x *goEventServiceGoChainRequestEventClient) Send(m *ChainTranscationReq) error {
	return x.ClientStream.SendMsg(m)
}

func (x *goEventServiceGoChainRequestEventClient) Recv() (*ChainTranscationRes, error) {
	m := new(ChainTranscationRes)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for GoEventService service

type GoEventServiceServer interface {
	GoClientRegistEvent(context.Context, *ClientRegisterAddressReq) (*ClientRegisterAddressRes, error)
	GoClientRequestEvent(context.Context, *ClientTransactionReq) (*ClientTransactionRes, error)
	GoChainRequestEvent(GoEventService_GoChainRequestEventServer) error
}

func RegisterGoEventServiceServer(s *grpc.Server, srv GoEventServiceServer) {
	s.RegisterService(&_GoEventService_serviceDesc, srv)
}

func _GoEventService_GoClientRegistEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClientRegisterAddressReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoEventServiceServer).GoClientRegistEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/server.GoEventService/GoClientRegistEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoEventServiceServer).GoClientRegistEvent(ctx, req.(*ClientRegisterAddressReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _GoEventService_GoClientRequestEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClientTransactionReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoEventServiceServer).GoClientRequestEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/server.GoEventService/GoClientRequestEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoEventServiceServer).GoClientRequestEvent(ctx, req.(*ClientTransactionReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _GoEventService_GoChainRequestEvent_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GoEventServiceServer).GoChainRequestEvent(&goEventServiceGoChainRequestEventServer{stream})
}

type GoEventService_GoChainRequestEventServer interface {
	Send(*ChainTranscationRes) error
	Recv() (*ChainTranscationReq, error)
	grpc.ServerStream
}

type goEventServiceGoChainRequestEventServer struct {
	grpc.ServerStream
}

func (x *goEventServiceGoChainRequestEventServer) Send(m *ChainTranscationRes) error {
	return x.ServerStream.SendMsg(m)
}

func (x *goEventServiceGoChainRequestEventServer) Recv() (*ChainTranscationReq, error) {
	m := new(ChainTranscationReq)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _GoEventService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "server.GoEventService",
	HandlerType: (*GoEventServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GoClientRegistEvent",
			Handler:    _GoEventService_GoClientRegistEvent_Handler,
		},
		{
			MethodName: "GoClientRequestEvent",
			Handler:    _GoEventService_GoClientRequestEvent_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GoChainRequestEvent",
			Handler:       _GoEventService_GoChainRequestEvent_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "eventServer.proto",
}

// Client API for JavaEventService service

type JavaEventServiceClient interface {
	JavaClientRequestEvent(ctx context.Context, opts ...grpc.CallOption) (JavaEventService_JavaClientRequestEventClient, error)
}

type javaEventServiceClient struct {
	cc *grpc.ClientConn
}

func NewJavaEventServiceClient(cc *grpc.ClientConn) JavaEventServiceClient {
	return &javaEventServiceClient{cc}
}

func (c *javaEventServiceClient) JavaClientRequestEvent(ctx context.Context, opts ...grpc.CallOption) (JavaEventService_JavaClientRequestEventClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_JavaEventService_serviceDesc.Streams[0], c.cc, "/server.JavaEventService/JavaClientRequestEvent", opts...)
	if err != nil {
		return nil, err
	}
	x := &javaEventServiceJavaClientRequestEventClient{stream}
	return x, nil
}

type JavaEventService_JavaClientRequestEventClient interface {
	Send(*ClientTransactionJavaReq) error
	Recv() (*ClientTransactionJavaRes, error)
	grpc.ClientStream
}

type javaEventServiceJavaClientRequestEventClient struct {
	grpc.ClientStream
}

func (x *javaEventServiceJavaClientRequestEventClient) Send(m *ClientTransactionJavaReq) error {
	return x.ClientStream.SendMsg(m)
}

func (x *javaEventServiceJavaClientRequestEventClient) Recv() (*ClientTransactionJavaRes, error) {
	m := new(ClientTransactionJavaRes)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for JavaEventService service

type JavaEventServiceServer interface {
	JavaClientRequestEvent(JavaEventService_JavaClientRequestEventServer) error
}

func RegisterJavaEventServiceServer(s *grpc.Server, srv JavaEventServiceServer) {
	s.RegisterService(&_JavaEventService_serviceDesc, srv)
}

func _JavaEventService_JavaClientRequestEvent_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(JavaEventServiceServer).JavaClientRequestEvent(&javaEventServiceJavaClientRequestEventServer{stream})
}

type JavaEventService_JavaClientRequestEventServer interface {
	Send(*ClientTransactionJavaRes) error
	Recv() (*ClientTransactionJavaReq, error)
	grpc.ServerStream
}

type javaEventServiceJavaClientRequestEventServer struct {
	grpc.ServerStream
}

func (x *javaEventServiceJavaClientRequestEventServer) Send(m *ClientTransactionJavaRes) error {
	return x.ServerStream.SendMsg(m)
}

func (x *javaEventServiceJavaClientRequestEventServer) Recv() (*ClientTransactionJavaReq, error) {
	m := new(ClientTransactionJavaReq)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _JavaEventService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "server.JavaEventService",
	HandlerType: (*JavaEventServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "JavaClientRequestEvent",
			Handler:       _JavaEventService_JavaClientRequestEvent_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "eventServer.proto",
}

func init() { proto.RegisterFile("eventServer.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 502 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0xad, 0x13, 0x11, 0xe8, 0x04, 0xaa, 0xb2, 0x8d, 0x50, 0x64, 0xaa, 0x2a, 0x5a, 0x21, 0xc8,
	0x29, 0xaa, 0xc2, 0xa5, 0x70, 0x41, 0x4d, 0xa8, 0xaa, 0x20, 0x21, 0x45, 0x9b, 0x70, 0x29, 0xa7,
	0xad, 0x3d, 0x6a, 0x4d, 0x8b, 0xdd, 0xec, 0x38, 0x56, 0x2b, 0x71, 0xe4, 0x2b, 0xf8, 0x0b, 0xfe,
	0x10, 0xed, 0xae, 0x8d, 0xbd, 0x51, 0xe2, 0x5c, 0x7a, 0x4a, 0xe6, 0xcd, 0x9b, 0x99, 0x37, 0x33,
	0x3b, 0x86, 0x97, 0x98, 0x61, 0x9c, 0xce, 0x50, 0x65, 0xa8, 0x06, 0x77, 0x2a, 0x49, 0x13, 0xd6,
	0x22, 0x63, 0xf1, 0xdf, 0x1e, 0x74, 0xc7, 0xb7, 0x11, 0xc6, 0xa9, 0xc0, 0xab, 0x88, 0x52, 0x54,
	0xa7, 0x61, 0xa8, 0x90, 0x48, 0xe0, 0x82, 0x71, 0x78, 0x7e, 0x1a, 0x86, 0x02, 0x89, 0x26, 0x77,
	0x02, 0x17, 0x5d, 0xaf, 0xe7, 0xf5, 0x77, 0x85, 0x83, 0xb1, 0xb7, 0xb0, 0x97, 0xdb, 0xd3, 0x44,
	0xa5, 0x9a, 0xd5, 0x30, 0xac, 0x15, 0x94, 0x1d, 0xc2, 0xae, 0xc0, 0x9f, 0x52, 0xdd, 0x68, 0x4a,
	0xd3, 0x50, 0x4a, 0x80, 0xff, 0xda, 0xa8, 0x82, 0xd8, 0x11, 0xc0, 0x57, 0x24, 0x92, 0x57, 0x28,
	0x90, 0x72, 0x0d, 0x15, 0x44, 0x67, 0x9e, 0xd0, 0x6c, 0x19, 0x04, 0x48, 0x64, 0x8a, 0x3f, 0x13,
	0x25, 0xa0, 0x7b, 0xc8, 0xb9, 0x93, 0xcf, 0x3a, 0xde, 0x96, 0x76, 0x30, 0x7e, 0x0c, 0x1d, 0x5b,
	0x7d, 0xae, 0x64, 0x4c, 0x32, 0x48, 0xa3, 0x24, 0xd6, 0x9a, 0xbb, 0xf0, 0x74, 0x7e, 0x1f, 0x85,
	0x65, 0xeb, 0x85, 0xc9, 0x7f, 0xac, 0x8d, 0xa0, 0x32, 0x82, 0xdc, 0x08, 0xe3, 0x19, 0x27, 0xa1,
	0x69, 0xc1, 0x0e, 0xa8, 0x30, 0x57, 0xfa, 0x6b, 0xae, 0xf6, 0xc7, 0xff, 0x7a, 0x70, 0x30, 0xbe,
	0x96, 0x51, 0x6c, 0x6a, 0x05, 0xb2, 0x50, 0x77, 0x04, 0x30, 0x92, 0x69, 0x70, 0x5d, 0x15, 0x58,
	0x41, 0xd8, 0x27, 0x68, 0xcd, 0xef, 0xc9, 0x6e, 0xa4, 0xd9, 0x6f, 0x0f, 0xdf, 0x0d, 0xec, 0xce,
	0x07, 0x6b, 0x92, 0x0d, 0x2c, 0xf3, 0x2c, 0x4e, 0xd5, 0x83, 0xc8, 0xc3, 0xfc, 0x0f, 0xd0, 0xae,
	0xc0, 0x6c, 0x1f, 0x9a, 0x37, 0xf8, 0x90, 0x17, 0xd2, 0x7f, 0x59, 0x07, 0x9e, 0x64, 0xf2, 0x76,
	0x89, 0x79, 0x47, 0xd6, 0xf8, 0xd8, 0x38, 0xf1, 0xf8, 0x6c, 0x9d, 0x64, 0xda, 0x2a, 0xb9, 0x76,
	0x95, 0x3c, 0x2e, 0x1e, 0x49, 0x65, 0xe8, 0x5f, 0x64, 0x26, 0x9d, 0x55, 0x3d, 0xe2, 0xe0, 0x2f,
	0x36, 0xd6, 0xab, 0x5b, 0xf4, 0x1b, 0x78, 0x31, 0x21, 0x81, 0x01, 0x46, 0x19, 0x86, 0x65, 0x55,
	0x17, 0x1c, 0xfe, 0x69, 0xc0, 0xde, 0x79, 0x72, 0x56, 0xdc, 0x65, 0x14, 0x20, 0xfb, 0x0e, 0x07,
	0xe7, 0x49, 0xf5, 0x0a, 0x8c, 0x97, 0xf5, 0xfe, 0xaf, 0x6d, 0xc3, 0x99, 0xfa, 0xdb, 0x18, 0xc4,
	0x77, 0xd8, 0x1c, 0x3a, 0x65, 0xf2, 0xc5, 0x12, 0x8b, 0xec, 0x87, 0x6e, 0xac, 0x7b, 0x00, 0x7e,
	0x9d, 0x57, 0x67, 0xfd, 0x66, 0x24, 0xeb, 0x45, 0x3b, 0x49, 0x5f, 0xd7, 0xbc, 0x34, 0xbf, 0xc6,
	0x49, 0x7c, 0xa7, 0xef, 0x1d, 0x7b, 0xc3, 0x0c, 0xf6, 0xf5, 0x9c, 0x9d, 0xe9, 0x5c, 0xc2, 0x2b,
	0x8d, 0xad, 0x69, 0xa1, 0xb7, 0x51, 0x64, 0xfe, 0x38, 0xfc, 0x6d, 0x8c, 0xbc, 0xee, 0xe8, 0x04,
	0xda, 0x96, 0x68, 0xbe, 0x91, 0x23, 0x5f, 0x7b, 0x2b, 0x02, 0xed, 0x07, 0x74, 0xaa, 0x7d, 0x53,
	0xef, 0xa2, 0x4a, 0xbd, 0x6c, 0x99, 0x9f, 0xf7, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0x8a, 0x7f,
	0x92, 0x2e, 0x6a, 0x05, 0x00, 0x00,
}
