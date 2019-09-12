// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ip.proto

package proto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Resp_MessageType int32

const (
	Resp_SUCCESS    Resp_MessageType = 0
	Resp_FAILED     Resp_MessageType = 1
	Resp_NOT_ENOUGH Resp_MessageType = 2
)

var Resp_MessageType_name = map[int32]string{
	0: "SUCCESS",
	1: "FAILED",
	2: "NOT_ENOUGH",
}

var Resp_MessageType_value = map[string]int32{
	"SUCCESS":    0,
	"FAILED":     1,
	"NOT_ENOUGH": 2,
}

func (x Resp_MessageType) String() string {
	return proto.EnumName(Resp_MessageType_name, int32(x))
}

func (Resp_MessageType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0035787f70b0c50f, []int{1, 0}
}

type Req struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Req) Reset()         { *m = Req{} }
func (m *Req) String() string { return proto.CompactTextString(m) }
func (*Req) ProtoMessage()    {}
func (*Req) Descriptor() ([]byte, []int) {
	return fileDescriptor_0035787f70b0c50f, []int{0}
}

func (m *Req) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Req.Unmarshal(m, b)
}
func (m *Req) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Req.Marshal(b, m, deterministic)
}
func (m *Req) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Req.Merge(m, src)
}
func (m *Req) XXX_Size() int {
	return xxx_messageInfo_Req.Size(m)
}
func (m *Req) XXX_DiscardUnknown() {
	xxx_messageInfo_Req.DiscardUnknown(m)
}

var xxx_messageInfo_Req proto.InternalMessageInfo

type Resp struct {
	Proxy                string           `protobuf:"bytes,1,opt,name=proxy,proto3" json:"proxy,omitempty"`
	Message              Resp_MessageType `protobuf:"varint,2,opt,name=message,proto3,enum=proto.Resp_MessageType" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Resp) Reset()         { *m = Resp{} }
func (m *Resp) String() string { return proto.CompactTextString(m) }
func (*Resp) ProtoMessage()    {}
func (*Resp) Descriptor() ([]byte, []int) {
	return fileDescriptor_0035787f70b0c50f, []int{1}
}

func (m *Resp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Resp.Unmarshal(m, b)
}
func (m *Resp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Resp.Marshal(b, m, deterministic)
}
func (m *Resp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Resp.Merge(m, src)
}
func (m *Resp) XXX_Size() int {
	return xxx_messageInfo_Resp.Size(m)
}
func (m *Resp) XXX_DiscardUnknown() {
	xxx_messageInfo_Resp.DiscardUnknown(m)
}

var xxx_messageInfo_Resp proto.InternalMessageInfo

func (m *Resp) GetProxy() string {
	if m != nil {
		return m.Proxy
	}
	return ""
}

func (m *Resp) GetMessage() Resp_MessageType {
	if m != nil {
		return m.Message
	}
	return Resp_SUCCESS
}

func init() {
	proto.RegisterEnum("proto.Resp_MessageType", Resp_MessageType_name, Resp_MessageType_value)
	proto.RegisterType((*Req)(nil), "proto.Req")
	proto.RegisterType((*Resp)(nil), "proto.Resp")
}

func init() { proto.RegisterFile("ip.proto", fileDescriptor_0035787f70b0c50f) }

var fileDescriptor_0035787f70b0c50f = []byte{
	// 188 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xc8, 0x2c, 0xd0, 0x2b,
	0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x53, 0x4a, 0xac, 0x5c, 0xcc, 0x41, 0xa9, 0x85, 0x4a,
	0xed, 0x8c, 0x5c, 0x2c, 0x41, 0xa9, 0xc5, 0x05, 0x42, 0x22, 0x5c, 0x20, 0x89, 0x8a, 0x4a, 0x09,
	0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x08, 0x47, 0xc8, 0x90, 0x8b, 0x3d, 0x37, 0xb5, 0xb8, 0x38,
	0x31, 0x3d, 0x55, 0x82, 0x49, 0x81, 0x51, 0x83, 0xcf, 0x48, 0x1c, 0x62, 0x8a, 0x1e, 0x48, 0x8f,
	0x9e, 0x2f, 0x44, 0x2a, 0xa4, 0xb2, 0x20, 0x35, 0x08, 0xa6, 0x4e, 0xc9, 0x8c, 0x8b, 0x1b, 0x49,
	0x5c, 0x88, 0x9b, 0x8b, 0x3d, 0x38, 0xd4, 0xd9, 0xd9, 0x35, 0x38, 0x58, 0x80, 0x41, 0x88, 0x8b,
	0x8b, 0xcd, 0xcd, 0xd1, 0xd3, 0xc7, 0xd5, 0x45, 0x80, 0x51, 0x88, 0x8f, 0x8b, 0xcb, 0xcf, 0x3f,
	0x24, 0xde, 0xd5, 0xcf, 0x3f, 0xd4, 0xdd, 0x43, 0x80, 0xc9, 0xc8, 0x98, 0x8b, 0x33, 0x00, 0x64,
	0x67, 0x40, 0x7e, 0x7e, 0x8e, 0x90, 0x1a, 0x17, 0xa7, 0x73, 0x62, 0x4e, 0x0e, 0x58, 0x40, 0x88,
	0x0b, 0x6e, 0x67, 0xa1, 0x14, 0x37, 0x92, 0xfd, 0x4a, 0x0c, 0x49, 0x6c, 0x60, 0x9e, 0x31, 0x20,
	0x00, 0x00, 0xff, 0xff, 0x8d, 0x1d, 0x8a, 0x2a, 0xdf, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ProxyPoolClient is the client API for ProxyPool service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ProxyPoolClient interface {
	CallProxy(ctx context.Context, in *Req, opts ...grpc.CallOption) (*Resp, error)
}

type proxyPoolClient struct {
	cc *grpc.ClientConn
}

func NewProxyPoolClient(cc *grpc.ClientConn) ProxyPoolClient {
	return &proxyPoolClient{cc}
}

func (c *proxyPoolClient) CallProxy(ctx context.Context, in *Req, opts ...grpc.CallOption) (*Resp, error) {
	out := new(Resp)
	err := c.cc.Invoke(ctx, "/proto.ProxyPool/CallProxy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProxyPoolServer is the server API for ProxyPool service.
type ProxyPoolServer interface {
	CallProxy(context.Context, *Req) (*Resp, error)
}

// UnimplementedProxyPoolServer can be embedded to have forward compatible implementations.
type UnimplementedProxyPoolServer struct {
}

func (*UnimplementedProxyPoolServer) CallProxy(ctx context.Context, req *Req) (*Resp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CallProxy not implemented")
}

func RegisterProxyPoolServer(s *grpc.Server, srv ProxyPoolServer) {
	s.RegisterService(&_ProxyPool_serviceDesc, srv)
}

func _ProxyPool_CallProxy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Req)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProxyPoolServer).CallProxy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ProxyPool/CallProxy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProxyPoolServer).CallProxy(ctx, req.(*Req))
	}
	return interceptor(ctx, in, info, handler)
}

var _ProxyPool_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.ProxyPool",
	HandlerType: (*ProxyPoolServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CallProxy",
			Handler:    _ProxyPool_CallProxy_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ip.proto",
}
