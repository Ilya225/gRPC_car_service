// Code generated by protoc-gen-go. DO NOT EDIT.
// source: car_service.proto

package car_service_client

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

type GetCarByIDRequest struct {
	ID                   string   `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetCarByIDRequest) Reset()         { *m = GetCarByIDRequest{} }
func (m *GetCarByIDRequest) String() string { return proto.CompactTextString(m) }
func (*GetCarByIDRequest) ProtoMessage()    {}
func (*GetCarByIDRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_19805971e15d66a1, []int{0}
}

func (m *GetCarByIDRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetCarByIDRequest.Unmarshal(m, b)
}
func (m *GetCarByIDRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetCarByIDRequest.Marshal(b, m, deterministic)
}
func (m *GetCarByIDRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCarByIDRequest.Merge(m, src)
}
func (m *GetCarByIDRequest) XXX_Size() int {
	return xxx_messageInfo_GetCarByIDRequest.Size(m)
}
func (m *GetCarByIDRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCarByIDRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetCarByIDRequest proto.InternalMessageInfo

func (m *GetCarByIDRequest) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

type CarResponse struct {
	ID                   string   `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CarResponse) Reset()         { *m = CarResponse{} }
func (m *CarResponse) String() string { return proto.CompactTextString(m) }
func (*CarResponse) ProtoMessage()    {}
func (*CarResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_19805971e15d66a1, []int{1}
}

func (m *CarResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CarResponse.Unmarshal(m, b)
}
func (m *CarResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CarResponse.Marshal(b, m, deterministic)
}
func (m *CarResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CarResponse.Merge(m, src)
}
func (m *CarResponse) XXX_Size() int {
	return xxx_messageInfo_CarResponse.Size(m)
}
func (m *CarResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CarResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CarResponse proto.InternalMessageInfo

func (m *CarResponse) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func init() {
	proto.RegisterType((*GetCarByIDRequest)(nil), "car_service_client.GetCarByIDRequest")
	proto.RegisterType((*CarResponse)(nil), "car_service_client.CarResponse")
}

func init() {
	proto.RegisterFile("car_service.proto", fileDescriptor_19805971e15d66a1)
}

var fileDescriptor_19805971e15d66a1 = []byte{
	// 193 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0x4c, 0x4e, 0x2c, 0x8a,
	0x2f, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x42,
	0x12, 0x8a, 0x4f, 0xce, 0xc9, 0x4c, 0xcd, 0x2b, 0x51, 0x52, 0xe6, 0x12, 0x74, 0x4f, 0x2d, 0x71,
	0x4e, 0x2c, 0x72, 0xaa, 0xf4, 0x74, 0x09, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0xe2, 0xe3,
	0x62, 0xf2, 0x74, 0x91, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x02, 0xb2, 0x94, 0x64, 0xb9, 0xb8,
	0x81, 0x2a, 0x82, 0x52, 0x8b, 0x0b, 0xf2, 0xf3, 0x8a, 0x53, 0xd1, 0xa5, 0x8d, 0x12, 0xb8, 0xb8,
	0x80, 0xd2, 0xc1, 0x10, 0x83, 0x85, 0x82, 0xb8, 0xd8, 0x20, 0x26, 0x0a, 0xa9, 0xea, 0x61, 0x5a,
	0xa8, 0x87, 0x61, 0x9b, 0x94, 0x3c, 0x36, 0x65, 0x48, 0xf6, 0x29, 0x31, 0x38, 0x05, 0x70, 0xc9,
	0xa5, 0x56, 0x24, 0xe6, 0x16, 0xe4, 0xa4, 0xea, 0xe5, 0x17, 0xa5, 0x83, 0xd4, 0xeb, 0xc1, 0xbc,
	0x06, 0x51, 0xef, 0x24, 0x80, 0x70, 0x81, 0x33, 0x58, 0x24, 0x80, 0x31, 0x4a, 0x2a, 0x33, 0xa7,
	0x32, 0xd1, 0xc8, 0xc8, 0x54, 0x1f, 0xa8, 0x5e, 0x17, 0xaa, 0x5e, 0x17, 0xa2, 0x3e, 0x89, 0x0d,
	0x1c, 0x24, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x31, 0xe8, 0xfc, 0x35, 0x27, 0x01, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CarServiceClient is the client API for CarService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CarServiceClient interface {
	GetCar(ctx context.Context, in *GetCarByIDRequest, opts ...grpc.CallOption) (*CarResponse, error)
}

type carServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCarServiceClient(cc grpc.ClientConnInterface) CarServiceClient {
	return &carServiceClient{cc}
}

func (c *carServiceClient) GetCar(ctx context.Context, in *GetCarByIDRequest, opts ...grpc.CallOption) (*CarResponse, error) {
	out := new(CarResponse)
	err := c.cc.Invoke(ctx, "/car_service_client.CarService/GetCar", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CarServiceServer is the server API for CarService service.
type CarServiceServer interface {
	GetCar(context.Context, *GetCarByIDRequest) (*CarResponse, error)
}

// UnimplementedCarServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCarServiceServer struct {
}

func (*UnimplementedCarServiceServer) GetCar(ctx context.Context, req *GetCarByIDRequest) (*CarResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCar not implemented")
}

func RegisterCarServiceServer(s *grpc.Server, srv CarServiceServer) {
	s.RegisterService(&_CarService_serviceDesc, srv)
}

func _CarService_GetCar_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCarByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CarServiceServer).GetCar(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/car_service_client.CarService/GetCar",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CarServiceServer).GetCar(ctx, req.(*GetCarByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CarService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "car_service_client.CarService",
	HandlerType: (*CarServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCar",
			Handler:    _CarService_GetCar_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "car_service.proto",
}
