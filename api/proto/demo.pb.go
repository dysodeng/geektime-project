// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/proto/demo.proto

package proto

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

type UserRequest struct {
	UserId               uint64   `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserRequest) Reset()         { *m = UserRequest{} }
func (m *UserRequest) String() string { return proto.CompactTextString(m) }
func (*UserRequest) ProtoMessage()    {}
func (*UserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_32bace1b3e25dcdc, []int{0}
}

func (m *UserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserRequest.Unmarshal(m, b)
}
func (m *UserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserRequest.Marshal(b, m, deterministic)
}
func (m *UserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserRequest.Merge(m, src)
}
func (m *UserRequest) XXX_Size() int {
	return xxx_messageInfo_UserRequest.Size(m)
}
func (m *UserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UserRequest proto.InternalMessageInfo

func (m *UserRequest) GetUserId() uint64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

type UserResponse struct {
	UserId               uint64   `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Username             string   `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Nickname             string   `protobuf:"bytes,3,opt,name=nickname,proto3" json:"nickname,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserResponse) Reset()         { *m = UserResponse{} }
func (m *UserResponse) String() string { return proto.CompactTextString(m) }
func (*UserResponse) ProtoMessage()    {}
func (*UserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_32bace1b3e25dcdc, []int{1}
}

func (m *UserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserResponse.Unmarshal(m, b)
}
func (m *UserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserResponse.Marshal(b, m, deterministic)
}
func (m *UserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserResponse.Merge(m, src)
}
func (m *UserResponse) XXX_Size() int {
	return xxx_messageInfo_UserResponse.Size(m)
}
func (m *UserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UserResponse proto.InternalMessageInfo

func (m *UserResponse) GetUserId() uint64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *UserResponse) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *UserResponse) GetNickname() string {
	if m != nil {
		return m.Nickname
	}
	return ""
}

func init() {
	proto.RegisterType((*UserRequest)(nil), "proto.UserRequest")
	proto.RegisterType((*UserResponse)(nil), "proto.UserResponse")
}

func init() { proto.RegisterFile("api/proto/demo.proto", fileDescriptor_32bace1b3e25dcdc) }

var fileDescriptor_32bace1b3e25dcdc = []byte{
	// 178 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x49, 0x2c, 0xc8, 0xd4,
	0x2f, 0x28, 0xca, 0x2f, 0xc9, 0xd7, 0x4f, 0x49, 0xcd, 0xcd, 0xd7, 0x03, 0x33, 0x85, 0x58, 0xc1,
	0x94, 0x92, 0x1a, 0x17, 0x77, 0x68, 0x71, 0x6a, 0x51, 0x50, 0x6a, 0x61, 0x69, 0x6a, 0x71, 0x89,
	0x90, 0x38, 0x17, 0x7b, 0x69, 0x71, 0x6a, 0x51, 0x7c, 0x66, 0x8a, 0x04, 0xa3, 0x02, 0xa3, 0x06,
	0x4b, 0x10, 0x1b, 0x88, 0xeb, 0x99, 0xa2, 0x14, 0xcf, 0xc5, 0x03, 0x51, 0x57, 0x5c, 0x90, 0x9f,
	0x57, 0x9c, 0x8a, 0x53, 0xa1, 0x90, 0x14, 0x17, 0x07, 0x88, 0x95, 0x97, 0x98, 0x9b, 0x2a, 0xc1,
	0xa4, 0xc0, 0xa8, 0xc1, 0x19, 0x04, 0xe7, 0x83, 0xe4, 0xf2, 0x32, 0x93, 0xb3, 0xc1, 0x72, 0xcc,
	0x10, 0x39, 0x18, 0xdf, 0xc8, 0x89, 0x8b, 0xdb, 0x25, 0x35, 0x37, 0x3f, 0x38, 0xb5, 0xa8, 0x2c,
	0x33, 0x39, 0x55, 0xc8, 0x98, 0x8b, 0x03, 0x64, 0x9f, 0x67, 0x5e, 0x5a, 0xbe, 0x90, 0x10, 0xc4,
	0xc9, 0x7a, 0x48, 0x0e, 0x95, 0x12, 0x46, 0x11, 0x83, 0x38, 0xca, 0x89, 0x3d, 0x0a, 0xe2, 0xab,
	0x24, 0x36, 0x30, 0x65, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x18, 0xba, 0x1d, 0x34, 0xfb, 0x00,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DemoServiceClient is the client API for DemoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DemoServiceClient interface {
	UserInfo(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserResponse, error)
}

type demoServiceClient struct {
	cc *grpc.ClientConn
}

func NewDemoServiceClient(cc *grpc.ClientConn) DemoServiceClient {
	return &demoServiceClient{cc}
}

func (c *demoServiceClient) UserInfo(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, "/proto.DemoService/UserInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DemoServiceServer is the server API for DemoService service.
type DemoServiceServer interface {
	UserInfo(context.Context, *UserRequest) (*UserResponse, error)
}

func RegisterDemoServiceServer(s *grpc.Server, srv DemoServiceServer) {
	s.RegisterService(&_DemoService_serviceDesc, srv)
}

func _DemoService_UserInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DemoServiceServer).UserInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.DemoService/UserInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DemoServiceServer).UserInfo(ctx, req.(*UserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DemoService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.DemoService",
	HandlerType: (*DemoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UserInfo",
			Handler:    _DemoService_UserInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/proto/demo.proto",
}