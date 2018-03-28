// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/consignment/consignment.proto

/*
Package go_micro_srv_consignment is a generated protocol buffer package.

It is generated from these files:
	proto/consignment/consignment.proto

It has these top-level messages:
	Consignment
	Container
	GetRequest
	Response
*/
package go_micro_srv_consignment

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

type Consignment struct {
	Id          string       `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Description string       `protobuf:"bytes,2,opt,name=description" json:"description,omitempty"`
	Weight      int32        `protobuf:"varint,3,opt,name=weight" json:"weight,omitempty"`
	Containers  []*Container `protobuf:"bytes,4,rep,name=containers" json:"containers,omitempty"`
	VesselId    string       `protobuf:"bytes,5,opt,name=vessel_id,json=vesselId" json:"vessel_id,omitempty"`
}

func (m *Consignment) Reset()                    { *m = Consignment{} }
func (m *Consignment) String() string            { return proto.CompactTextString(m) }
func (*Consignment) ProtoMessage()               {}
func (*Consignment) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Consignment) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Consignment) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Consignment) GetWeight() int32 {
	if m != nil {
		return m.Weight
	}
	return 0
}

func (m *Consignment) GetContainers() []*Container {
	if m != nil {
		return m.Containers
	}
	return nil
}

func (m *Consignment) GetVesselId() string {
	if m != nil {
		return m.VesselId
	}
	return ""
}

type Container struct {
	Id         string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	CustomerId string `protobuf:"bytes,2,opt,name=customer_id,json=customerId" json:"customer_id,omitempty"`
	Origin     string `protobuf:"bytes,3,opt,name=origin" json:"origin,omitempty"`
	UserId     string `protobuf:"bytes,4,opt,name=user_id,json=userId" json:"user_id,omitempty"`
}

func (m *Container) Reset()                    { *m = Container{} }
func (m *Container) String() string            { return proto.CompactTextString(m) }
func (*Container) ProtoMessage()               {}
func (*Container) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Container) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Container) GetCustomerId() string {
	if m != nil {
		return m.CustomerId
	}
	return ""
}

func (m *Container) GetOrigin() string {
	if m != nil {
		return m.Origin
	}
	return ""
}

func (m *Container) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

// Created a blanck get request
type GetRequest struct {
}

func (m *GetRequest) Reset()                    { *m = GetRequest{} }
func (m *GetRequest) String() string            { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()               {}
func (*GetRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type Response struct {
	Created     bool         `protobuf:"varint,1,opt,name=created" json:"created,omitempty"`
	Consignment *Consignment `protobuf:"bytes,2,opt,name=consignment" json:"consignment,omitempty"`
	// Added a pluralised consignment to our generic response message
	Consignments []*Consignment `protobuf:"bytes,3,rep,name=consignments" json:"consignments,omitempty"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Response) GetCreated() bool {
	if m != nil {
		return m.Created
	}
	return false
}

func (m *Response) GetConsignment() *Consignment {
	if m != nil {
		return m.Consignment
	}
	return nil
}

func (m *Response) GetConsignments() []*Consignment {
	if m != nil {
		return m.Consignments
	}
	return nil
}

func init() {
	proto.RegisterType((*Consignment)(nil), "go.micro.srv.consignment.Consignment")
	proto.RegisterType((*Container)(nil), "go.micro.srv.consignment.Container")
	proto.RegisterType((*GetRequest)(nil), "go.micro.srv.consignment.GetRequest")
	proto.RegisterType((*Response)(nil), "go.micro.srv.consignment.Response")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for ShippingService service

type ShippingServiceClient interface {
	CreateConsignment(ctx context.Context, in *Consignment, opts ...grpc.CallOption) (*Response, error)
	// Create a new method
	GetConsignments(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*Response, error)
}

type shippingServiceClient struct {
	cc *grpc.ClientConn
}

func NewShippingServiceClient(cc *grpc.ClientConn) ShippingServiceClient {
	return &shippingServiceClient{cc}
}

func (c *shippingServiceClient) CreateConsignment(ctx context.Context, in *Consignment, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/go.micro.srv.consignment.ShippingService/CreateConsignment", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shippingServiceClient) GetConsignments(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/go.micro.srv.consignment.ShippingService/GetConsignments", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ShippingService service

type ShippingServiceServer interface {
	CreateConsignment(context.Context, *Consignment) (*Response, error)
	// Create a new method
	GetConsignments(context.Context, *GetRequest) (*Response, error)
}

func RegisterShippingServiceServer(s *grpc.Server, srv ShippingServiceServer) {
	s.RegisterService(&_ShippingService_serviceDesc, srv)
}

func _ShippingService_CreateConsignment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Consignment)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShippingServiceServer).CreateConsignment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/go.micro.srv.consignment.ShippingService/CreateConsignment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShippingServiceServer).CreateConsignment(ctx, req.(*Consignment))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShippingService_GetConsignments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShippingServiceServer).GetConsignments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/go.micro.srv.consignment.ShippingService/GetConsignments",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShippingServiceServer).GetConsignments(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ShippingService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "go.micro.srv.consignment.ShippingService",
	HandlerType: (*ShippingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateConsignment",
			Handler:    _ShippingService_CreateConsignment_Handler,
		},
		{
			MethodName: "GetConsignments",
			Handler:    _ShippingService_GetConsignments_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/consignment/consignment.proto",
}

func init() { proto.RegisterFile("proto/consignment/consignment.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 351 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x4f, 0x4e, 0xeb, 0x30,
	0x10, 0x87, 0x5f, 0xfa, 0x3f, 0x93, 0xea, 0x55, 0x78, 0x01, 0x16, 0x2c, 0x88, 0x52, 0x90, 0xba,
	0x0a, 0x52, 0x39, 0x42, 0x16, 0x55, 0xb6, 0xee, 0x1a, 0x41, 0x89, 0x47, 0xa9, 0x25, 0x62, 0x07,
	0xdb, 0x2d, 0x57, 0x83, 0x6b, 0x70, 0x22, 0x54, 0xa7, 0xa1, 0x46, 0xa8, 0xa8, 0x3b, 0xff, 0x66,
	0xfc, 0x4d, 0xbe, 0x8c, 0x0c, 0xd3, 0x5a, 0x2b, 0xab, 0xee, 0x0a, 0x25, 0x8d, 0x28, 0x65, 0x85,
	0xd2, 0xfa, 0xe7, 0xd4, 0x75, 0x09, 0x2d, 0x55, 0x5a, 0x89, 0x42, 0xab, 0xd4, 0xe8, 0x6d, 0xea,
	0xf5, 0x93, 0x8f, 0x00, 0xa2, 0xec, 0x90, 0xc9, 0x7f, 0xe8, 0x08, 0x4e, 0x83, 0x38, 0x98, 0x85,
	0xac, 0x23, 0x38, 0x89, 0x21, 0xe2, 0x68, 0x0a, 0x2d, 0x6a, 0x2b, 0x94, 0xa4, 0x1d, 0xd7, 0xf0,
	0x4b, 0xe4, 0x1c, 0x06, 0x6f, 0x28, 0xca, 0xb5, 0xa5, 0xdd, 0x38, 0x98, 0xf5, 0xd9, 0x3e, 0x91,
	0x0c, 0xa0, 0x50, 0xd2, 0xae, 0x84, 0x44, 0x6d, 0x68, 0x2f, 0xee, 0xce, 0xa2, 0xf9, 0x34, 0x3d,
	0x26, 0x92, 0x66, 0xed, 0x5d, 0xe6, 0x61, 0xe4, 0x0a, 0xc2, 0x2d, 0x1a, 0x83, 0x2f, 0x8f, 0x82,
	0xd3, 0xbe, 0xfb, 0xf8, 0xa8, 0x29, 0xe4, 0x3c, 0xa9, 0x20, 0xfc, 0xa6, 0x7e, 0x89, 0x5f, 0x43,
	0x54, 0x6c, 0x8c, 0x55, 0x15, 0xea, 0x1d, 0xdb, 0x88, 0x43, 0x5b, 0xca, 0xf9, 0xce, 0x5b, 0x69,
	0x51, 0x0a, 0xe9, 0xbc, 0x43, 0xb6, 0x4f, 0xe4, 0x02, 0x86, 0x1b, 0xd3, 0x40, 0xbd, 0xa6, 0xb1,
	0x8b, 0x39, 0x4f, 0xc6, 0x00, 0x0b, 0xb4, 0x0c, 0x5f, 0x37, 0x68, 0x6c, 0xf2, 0x1e, 0xc0, 0x88,
	0xa1, 0xa9, 0x95, 0x34, 0x48, 0x28, 0x0c, 0x0b, 0x8d, 0x2b, 0x8b, 0x8d, 0xc1, 0x88, 0xb5, 0x91,
	0x2c, 0x20, 0xf2, 0xfe, 0xd2, 0x69, 0x44, 0xf3, 0xdb, 0x3f, 0xd7, 0xd0, 0x9e, 0x99, 0x4f, 0x92,
	0x1c, 0xc6, 0x5e, 0x34, 0xb4, 0xeb, 0x16, 0x7a, 0xe2, 0xa4, 0x1f, 0xe8, 0xfc, 0x33, 0x80, 0xc9,
	0x72, 0x2d, 0xea, 0x5a, 0xc8, 0x72, 0x89, 0x7a, 0x2b, 0x0a, 0x24, 0x4f, 0x70, 0x96, 0x39, 0x65,
	0xff, 0x31, 0x9c, 0x36, 0xfd, 0x32, 0x39, 0x7e, 0xad, 0xdd, 0x50, 0xf2, 0x8f, 0x3c, 0xc0, 0x64,
	0x81, 0xd6, 0xe3, 0x0c, 0xb9, 0x39, 0x0e, 0x1e, 0x36, 0x7d, 0xda, 0xf8, 0xe7, 0x81, 0x7b, 0xe9,
	0xf7, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x86, 0x3d, 0x82, 0x7e, 0x10, 0x03, 0x00, 0x00,
}
