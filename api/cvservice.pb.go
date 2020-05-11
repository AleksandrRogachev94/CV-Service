// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cvservice.proto

package main

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

type RecognizeResponseItem struct {
	Location             *FileLocation `protobuf:"bytes,1,opt,name=location,proto3" json:"location,omitempty"`
	Label                string        `protobuf:"bytes,2,opt,name=label,proto3" json:"label,omitempty"`
	Conf                 float64       `protobuf:"fixed64,3,opt,name=conf,proto3" json:"conf,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *RecognizeResponseItem) Reset()         { *m = RecognizeResponseItem{} }
func (m *RecognizeResponseItem) String() string { return proto.CompactTextString(m) }
func (*RecognizeResponseItem) ProtoMessage()    {}
func (*RecognizeResponseItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_2c056965df545451, []int{0}
}

func (m *RecognizeResponseItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RecognizeResponseItem.Unmarshal(m, b)
}
func (m *RecognizeResponseItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RecognizeResponseItem.Marshal(b, m, deterministic)
}
func (m *RecognizeResponseItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RecognizeResponseItem.Merge(m, src)
}
func (m *RecognizeResponseItem) XXX_Size() int {
	return xxx_messageInfo_RecognizeResponseItem.Size(m)
}
func (m *RecognizeResponseItem) XXX_DiscardUnknown() {
	xxx_messageInfo_RecognizeResponseItem.DiscardUnknown(m)
}

var xxx_messageInfo_RecognizeResponseItem proto.InternalMessageInfo

func (m *RecognizeResponseItem) GetLocation() *FileLocation {
	if m != nil {
		return m.Location
	}
	return nil
}

func (m *RecognizeResponseItem) GetLabel() string {
	if m != nil {
		return m.Label
	}
	return ""
}

func (m *RecognizeResponseItem) GetConf() float64 {
	if m != nil {
		return m.Conf
	}
	return 0
}

type FileLocation struct {
	Bucket               string   `protobuf:"bytes,1,opt,name=bucket,proto3" json:"bucket,omitempty"`
	Key                  string   `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FileLocation) Reset()         { *m = FileLocation{} }
func (m *FileLocation) String() string { return proto.CompactTextString(m) }
func (*FileLocation) ProtoMessage()    {}
func (*FileLocation) Descriptor() ([]byte, []int) {
	return fileDescriptor_2c056965df545451, []int{1}
}

func (m *FileLocation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FileLocation.Unmarshal(m, b)
}
func (m *FileLocation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FileLocation.Marshal(b, m, deterministic)
}
func (m *FileLocation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FileLocation.Merge(m, src)
}
func (m *FileLocation) XXX_Size() int {
	return xxx_messageInfo_FileLocation.Size(m)
}
func (m *FileLocation) XXX_DiscardUnknown() {
	xxx_messageInfo_FileLocation.DiscardUnknown(m)
}

var xxx_messageInfo_FileLocation proto.InternalMessageInfo

func (m *FileLocation) GetBucket() string {
	if m != nil {
		return m.Bucket
	}
	return ""
}

func (m *FileLocation) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

type RecognizeRequest struct {
	File                 *FileLocation `protobuf:"bytes,1,opt,name=file,proto3" json:"file,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *RecognizeRequest) Reset()         { *m = RecognizeRequest{} }
func (m *RecognizeRequest) String() string { return proto.CompactTextString(m) }
func (*RecognizeRequest) ProtoMessage()    {}
func (*RecognizeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2c056965df545451, []int{2}
}

func (m *RecognizeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RecognizeRequest.Unmarshal(m, b)
}
func (m *RecognizeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RecognizeRequest.Marshal(b, m, deterministic)
}
func (m *RecognizeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RecognizeRequest.Merge(m, src)
}
func (m *RecognizeRequest) XXX_Size() int {
	return xxx_messageInfo_RecognizeRequest.Size(m)
}
func (m *RecognizeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RecognizeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RecognizeRequest proto.InternalMessageInfo

func (m *RecognizeRequest) GetFile() *FileLocation {
	if m != nil {
		return m.File
	}
	return nil
}

type RecognizeResponse struct {
	Items                []*RecognizeResponseItem `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *RecognizeResponse) Reset()         { *m = RecognizeResponse{} }
func (m *RecognizeResponse) String() string { return proto.CompactTextString(m) }
func (*RecognizeResponse) ProtoMessage()    {}
func (*RecognizeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2c056965df545451, []int{3}
}

func (m *RecognizeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RecognizeResponse.Unmarshal(m, b)
}
func (m *RecognizeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RecognizeResponse.Marshal(b, m, deterministic)
}
func (m *RecognizeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RecognizeResponse.Merge(m, src)
}
func (m *RecognizeResponse) XXX_Size() int {
	return xxx_messageInfo_RecognizeResponse.Size(m)
}
func (m *RecognizeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RecognizeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RecognizeResponse proto.InternalMessageInfo

func (m *RecognizeResponse) GetItems() []*RecognizeResponseItem {
	if m != nil {
		return m.Items
	}
	return nil
}

func init() {
	proto.RegisterType((*RecognizeResponseItem)(nil), "main.RecognizeResponseItem")
	proto.RegisterType((*FileLocation)(nil), "main.FileLocation")
	proto.RegisterType((*RecognizeRequest)(nil), "main.RecognizeRequest")
	proto.RegisterType((*RecognizeResponse)(nil), "main.RecognizeResponse")
}

func init() {
	proto.RegisterFile("cvservice.proto", fileDescriptor_2c056965df545451)
}

var fileDescriptor_2c056965df545451 = []byte{
	// 254 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x51, 0xcf, 0x4b, 0xc3, 0x30,
	0x14, 0x36, 0xb6, 0x1b, 0xf6, 0x4d, 0x70, 0x3e, 0x74, 0x16, 0xbd, 0x94, 0x1c, 0xa4, 0xa7, 0x82,
	0xf5, 0x22, 0x1e, 0xbc, 0x08, 0x03, 0xd1, 0x53, 0x04, 0xef, 0x6d, 0x78, 0x93, 0xb0, 0x34, 0xd9,
	0x96, 0x6c, 0xa0, 0x7f, 0xbd, 0x2c, 0x29, 0x63, 0x30, 0xf5, 0xf6, 0x1e, 0xdf, 0x8f, 0x7c, 0xf9,
	0x1e, 0x9c, 0xc9, 0x8d, 0xa3, 0xd5, 0x46, 0x49, 0xaa, 0x16, 0x2b, 0xeb, 0x2d, 0xa6, 0x5d, 0xa3,
	0x0c, 0x5f, 0xc2, 0xa5, 0x20, 0x69, 0x3f, 0x8d, 0xfa, 0x26, 0x41, 0x6e, 0x61, 0x8d, 0xa3, 0x17,
	0x4f, 0x1d, 0x56, 0x70, 0xa2, 0xad, 0x6c, 0xbc, 0xb2, 0x26, 0x67, 0x05, 0x2b, 0x47, 0x35, 0x56,
	0x5b, 0x45, 0x35, 0x55, 0x9a, 0xde, 0x7a, 0x44, 0xec, 0x38, 0x78, 0x01, 0x03, 0xdd, 0xb4, 0xa4,
	0xf3, 0xe3, 0x82, 0x95, 0x99, 0x88, 0x0b, 0x22, 0xa4, 0xd2, 0x9a, 0x59, 0x9e, 0x14, 0xac, 0x64,
	0x22, 0xcc, 0xfc, 0x01, 0x4e, 0xf7, 0x3d, 0x70, 0x02, 0xc3, 0x76, 0x2d, 0xe7, 0xe4, 0xc3, 0x3b,
	0x99, 0xe8, 0x37, 0x1c, 0x43, 0x32, 0xa7, 0xaf, 0xde, 0x6f, 0x3b, 0xf2, 0x47, 0x18, 0xef, 0x85,
	0x5d, 0xae, 0xc9, 0x79, 0xbc, 0x85, 0x74, 0xa6, 0x34, 0xfd, 0x93, 0x31, 0xe0, 0x7c, 0x0a, 0xe7,
	0x07, 0x1f, 0xc5, 0x3b, 0x18, 0x28, 0x4f, 0x9d, 0xcb, 0x59, 0x91, 0x94, 0xa3, 0xfa, 0x26, 0xaa,
	0x7f, 0x2d, 0x44, 0x44, 0x66, 0xfd, 0x0a, 0xd9, 0xf3, 0xc7, 0x7b, 0x6c, 0x12, 0x9f, 0x20, 0xdb,
	0x91, 0x71, 0x72, 0xa0, 0x0e, 0x09, 0xaf, 0xaf, 0xfe, 0x70, 0xe5, 0x47, 0xed, 0x30, 0x9c, 0xe2,
	0xfe, 0x27, 0x00, 0x00, 0xff, 0xff, 0x4e, 0x63, 0x9f, 0xa4, 0x9d, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CVServiceClient is the client API for CVService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CVServiceClient interface {
	Recognize(ctx context.Context, in *RecognizeRequest, opts ...grpc.CallOption) (*RecognizeResponse, error)
}

type cVServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCVServiceClient(cc grpc.ClientConnInterface) CVServiceClient {
	return &cVServiceClient{cc}
}

func (c *cVServiceClient) Recognize(ctx context.Context, in *RecognizeRequest, opts ...grpc.CallOption) (*RecognizeResponse, error) {
	out := new(RecognizeResponse)
	err := c.cc.Invoke(ctx, "/main.CVService/Recognize", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CVServiceServer is the server API for CVService service.
type CVServiceServer interface {
	Recognize(context.Context, *RecognizeRequest) (*RecognizeResponse, error)
}

// UnimplementedCVServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCVServiceServer struct {
}

func (*UnimplementedCVServiceServer) Recognize(ctx context.Context, req *RecognizeRequest) (*RecognizeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Recognize not implemented")
}

func RegisterCVServiceServer(s *grpc.Server, srv CVServiceServer) {
	s.RegisterService(&_CVService_serviceDesc, srv)
}

func _CVService_Recognize_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RecognizeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CVServiceServer).Recognize(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.CVService/Recognize",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CVServiceServer).Recognize(ctx, req.(*RecognizeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CVService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "main.CVService",
	HandlerType: (*CVServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Recognize",
			Handler:    _CVService_Recognize_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cvservice.proto",
}
