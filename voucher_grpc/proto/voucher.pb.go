// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/voucher.proto

package proto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type Error struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Error) Reset()         { *m = Error{} }
func (m *Error) String() string { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()    {}
func (*Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_60f00a0a2a5aeccc, []int{0}
}

func (m *Error) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Error.Unmarshal(m, b)
}
func (m *Error) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Error.Marshal(b, m, deterministic)
}
func (m *Error) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Error.Merge(m, src)
}
func (m *Error) XXX_Size() int {
	return xxx_messageInfo_Error.Size(m)
}
func (m *Error) XXX_DiscardUnknown() {
	xxx_messageInfo_Error.DiscardUnknown(m)
}

var xxx_messageInfo_Error proto.InternalMessageInfo

func (m *Error) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Error) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type Voucher struct {
	Id                   int32                `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Code                 string               `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	Discount             float32              `protobuf:"fixed32,3,opt,name=discount,proto3" json:"discount,omitempty"`
	Start                *timestamp.Timestamp `protobuf:"bytes,4,opt,name=start,proto3" json:"start,omitempty"`
	End                  *timestamp.Timestamp `protobuf:"bytes,5,opt,name=end,proto3" json:"end,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Voucher) Reset()         { *m = Voucher{} }
func (m *Voucher) String() string { return proto.CompactTextString(m) }
func (*Voucher) ProtoMessage()    {}
func (*Voucher) Descriptor() ([]byte, []int) {
	return fileDescriptor_60f00a0a2a5aeccc, []int{1}
}

func (m *Voucher) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Voucher.Unmarshal(m, b)
}
func (m *Voucher) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Voucher.Marshal(b, m, deterministic)
}
func (m *Voucher) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Voucher.Merge(m, src)
}
func (m *Voucher) XXX_Size() int {
	return xxx_messageInfo_Voucher.Size(m)
}
func (m *Voucher) XXX_DiscardUnknown() {
	xxx_messageInfo_Voucher.DiscardUnknown(m)
}

var xxx_messageInfo_Voucher proto.InternalMessageInfo

func (m *Voucher) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Voucher) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *Voucher) GetDiscount() float32 {
	if m != nil {
		return m.Discount
	}
	return 0
}

func (m *Voucher) GetStart() *timestamp.Timestamp {
	if m != nil {
		return m.Start
	}
	return nil
}

func (m *Voucher) GetEnd() *timestamp.Timestamp {
	if m != nil {
		return m.End
	}
	return nil
}

// Neu da gan cho no mot gia tri la so = x, no never dc thay doi
type VoucherReq struct {
	Code                 string               `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Discount             float32              `protobuf:"fixed32,2,opt,name=discount,proto3" json:"discount,omitempty"`
	Start                *timestamp.Timestamp `protobuf:"bytes,3,opt,name=start,proto3" json:"start,omitempty"`
	End                  *timestamp.Timestamp `protobuf:"bytes,4,opt,name=end,proto3" json:"end,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *VoucherReq) Reset()         { *m = VoucherReq{} }
func (m *VoucherReq) String() string { return proto.CompactTextString(m) }
func (*VoucherReq) ProtoMessage()    {}
func (*VoucherReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_60f00a0a2a5aeccc, []int{2}
}

func (m *VoucherReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VoucherReq.Unmarshal(m, b)
}
func (m *VoucherReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VoucherReq.Marshal(b, m, deterministic)
}
func (m *VoucherReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VoucherReq.Merge(m, src)
}
func (m *VoucherReq) XXX_Size() int {
	return xxx_messageInfo_VoucherReq.Size(m)
}
func (m *VoucherReq) XXX_DiscardUnknown() {
	xxx_messageInfo_VoucherReq.DiscardUnknown(m)
}

var xxx_messageInfo_VoucherReq proto.InternalMessageInfo

func (m *VoucherReq) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *VoucherReq) GetDiscount() float32 {
	if m != nil {
		return m.Discount
	}
	return 0
}

func (m *VoucherReq) GetStart() *timestamp.Timestamp {
	if m != nil {
		return m.Start
	}
	return nil
}

func (m *VoucherReq) GetEnd() *timestamp.Timestamp {
	if m != nil {
		return m.End
	}
	return nil
}

type VoucherRes struct {
	Error                *Error   `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	Data                 *Voucher `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VoucherRes) Reset()         { *m = VoucherRes{} }
func (m *VoucherRes) String() string { return proto.CompactTextString(m) }
func (*VoucherRes) ProtoMessage()    {}
func (*VoucherRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_60f00a0a2a5aeccc, []int{3}
}

func (m *VoucherRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VoucherRes.Unmarshal(m, b)
}
func (m *VoucherRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VoucherRes.Marshal(b, m, deterministic)
}
func (m *VoucherRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VoucherRes.Merge(m, src)
}
func (m *VoucherRes) XXX_Size() int {
	return xxx_messageInfo_VoucherRes.Size(m)
}
func (m *VoucherRes) XXX_DiscardUnknown() {
	xxx_messageInfo_VoucherRes.DiscardUnknown(m)
}

var xxx_messageInfo_VoucherRes proto.InternalMessageInfo

func (m *VoucherRes) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *VoucherRes) GetData() *Voucher {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*Error)(nil), "proto.Error")
	proto.RegisterType((*Voucher)(nil), "proto.Voucher")
	proto.RegisterType((*VoucherReq)(nil), "proto.VoucherReq")
	proto.RegisterType((*VoucherRes)(nil), "proto.VoucherRes")
}

func init() { proto.RegisterFile("proto/voucher.proto", fileDescriptor_60f00a0a2a5aeccc) }

var fileDescriptor_60f00a0a2a5aeccc = []byte{
	// 298 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0x4d, 0x4e, 0xc3, 0x30,
	0x10, 0x85, 0x71, 0x1a, 0xd3, 0x32, 0x45, 0x95, 0x30, 0x1b, 0x2b, 0x1b, 0x22, 0xaf, 0xb2, 0x40,
	0x29, 0x0a, 0xe2, 0x06, 0x70, 0x01, 0x53, 0xb1, 0x4f, 0xe3, 0x21, 0x58, 0xa2, 0x75, 0x6b, 0x3b,
	0x3d, 0x0d, 0x17, 0xe0, 0x96, 0xa8, 0x76, 0x52, 0xf1, 0x2b, 0x60, 0x65, 0xbf, 0x99, 0xf7, 0xc6,
	0x9f, 0x35, 0x70, 0xbe, 0xb1, 0xc6, 0x9b, 0xf9, 0xce, 0x74, 0xcd, 0x13, 0xda, 0x32, 0x28, 0x46,
	0xc3, 0x91, 0x5d, 0xb4, 0xc6, 0xb4, 0xcf, 0x38, 0x0f, 0x6a, 0xd9, 0x3d, 0xce, 0xbd, 0x5e, 0xa1,
	0xf3, 0xf5, 0x6a, 0x13, 0x7d, 0xe2, 0x06, 0xe8, 0x9d, 0xb5, 0xc6, 0x32, 0x06, 0x69, 0x63, 0x14,
	0x72, 0x92, 0x93, 0x82, 0xca, 0x70, 0x67, 0x1c, 0xc6, 0x2b, 0x74, 0xae, 0x6e, 0x91, 0x27, 0x39,
	0x29, 0x4e, 0xe4, 0x20, 0xc5, 0x2b, 0x81, 0xf1, 0x43, 0x7c, 0x90, 0xcd, 0x20, 0xd1, 0xaa, 0xcf,
	0x25, 0x5a, 0x1d, 0x26, 0xc5, 0x48, 0x9c, 0x94, 0xc1, 0x44, 0x69, 0xd7, 0x98, 0x6e, 0xed, 0xf9,
	0x28, 0x27, 0x45, 0x22, 0x0f, 0x9a, 0x5d, 0x01, 0x75, 0xbe, 0xb6, 0x9e, 0xa7, 0x39, 0x29, 0xa6,
	0x55, 0x56, 0x46, 0xe6, 0x72, 0x60, 0x2e, 0x17, 0x03, 0xb3, 0x8c, 0x46, 0x76, 0x09, 0x23, 0x5c,
	0x2b, 0x4e, 0x7f, 0xf5, 0xef, 0x6d, 0xe2, 0x85, 0x00, 0xf4, 0xac, 0x12, 0xb7, 0x1f, 0x3e, 0xfa,
	0x1d, 0x5e, 0xf2, 0x13, 0xde, 0xe8, 0x9f, 0x78, 0xe9, 0xdf, 0xf0, 0x16, 0xef, 0xe8, 0x1c, 0x13,
	0x40, 0x71, 0xbf, 0x8f, 0x80, 0x37, 0xad, 0x4e, 0x63, 0xac, 0x0c, 0x3b, 0x92, 0xb1, 0xc5, 0x04,
	0xa4, 0xaa, 0xf6, 0x75, 0x20, 0x9d, 0x56, 0xb3, 0xde, 0x32, 0x0c, 0x09, 0xbd, 0xea, 0x16, 0x66,
	0x7d, 0xe1, 0x1e, 0xed, 0x4e, 0x37, 0xc8, 0x2a, 0x98, 0x48, 0x6c, 0xb5, 0xf3, 0x68, 0xd9, 0xd9,
	0xa7, 0x0c, 0x6e, 0xb3, 0x2f, 0x25, 0x27, 0x8e, 0x96, 0xc7, 0xa1, 0x76, 0xfd, 0x16, 0x00, 0x00,
	0xff, 0xff, 0xa3, 0x17, 0xb1, 0x26, 0x63, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// VoucherServiceClient is the client API for VoucherService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type VoucherServiceClient interface {
	Register(ctx context.Context, in *VoucherReq, opts ...grpc.CallOption) (*VoucherRes, error)
}

type voucherServiceClient struct {
	cc *grpc.ClientConn
}

func NewVoucherServiceClient(cc *grpc.ClientConn) VoucherServiceClient {
	return &voucherServiceClient{cc}
}

func (c *voucherServiceClient) Register(ctx context.Context, in *VoucherReq, opts ...grpc.CallOption) (*VoucherRes, error) {
	out := new(VoucherRes)
	err := c.cc.Invoke(ctx, "/proto.VoucherService/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VoucherServiceServer is the server API for VoucherService service.
type VoucherServiceServer interface {
	Register(context.Context, *VoucherReq) (*VoucherRes, error)
}

// UnimplementedVoucherServiceServer can be embedded to have forward compatible implementations.
type UnimplementedVoucherServiceServer struct {
}

func (*UnimplementedVoucherServiceServer) Register(ctx context.Context, req *VoucherReq) (*VoucherRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}

func RegisterVoucherServiceServer(s *grpc.Server, srv VoucherServiceServer) {
	s.RegisterService(&_VoucherService_serviceDesc, srv)
}

func _VoucherService_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VoucherReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VoucherServiceServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.VoucherService/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VoucherServiceServer).Register(ctx, req.(*VoucherReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _VoucherService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.VoucherService",
	HandlerType: (*VoucherServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _VoucherService_Register_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/voucher.proto",
}
