// Code generated by protoc-gen-go.
// source: proto/service.proto
// DO NOT EDIT!

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	proto/service.proto

It has these top-level messages:
	Candidate
	AddCandidateRequest
	CandidateResponse
	DeleteCandidateRequest
	DeleteCandidateResponse
	ListCandidateRequest
	ListCandidateResponse
*/
package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto1.ProtoPackageIsVersion2 // please upgrade the proto package

type Candidate struct {
	Name      string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Votes     int32  `protobuf:"varint,2,opt,name=votes" json:"votes,omitempty"`
	UpdatedAt int32  `protobuf:"varint,3,opt,name=updated_at,json=updatedAt" json:"updated_at,omitempty"`
}

func (m *Candidate) Reset()                    { *m = Candidate{} }
func (m *Candidate) String() string            { return proto1.CompactTextString(m) }
func (*Candidate) ProtoMessage()               {}
func (*Candidate) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type AddCandidateRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *AddCandidateRequest) Reset()                    { *m = AddCandidateRequest{} }
func (m *AddCandidateRequest) String() string            { return proto1.CompactTextString(m) }
func (*AddCandidateRequest) ProtoMessage()               {}
func (*AddCandidateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type CandidateResponse struct {
	Data *Candidate `protobuf:"bytes,1,opt,name=data" json:"data,omitempty"`
}

func (m *CandidateResponse) Reset()                    { *m = CandidateResponse{} }
func (m *CandidateResponse) String() string            { return proto1.CompactTextString(m) }
func (*CandidateResponse) ProtoMessage()               {}
func (*CandidateResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *CandidateResponse) GetData() *Candidate {
	if m != nil {
		return m.Data
	}
	return nil
}

type DeleteCandidateRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *DeleteCandidateRequest) Reset()                    { *m = DeleteCandidateRequest{} }
func (m *DeleteCandidateRequest) String() string            { return proto1.CompactTextString(m) }
func (*DeleteCandidateRequest) ProtoMessage()               {}
func (*DeleteCandidateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type DeleteCandidateResponse struct {
}

func (m *DeleteCandidateResponse) Reset()                    { *m = DeleteCandidateResponse{} }
func (m *DeleteCandidateResponse) String() string            { return proto1.CompactTextString(m) }
func (*DeleteCandidateResponse) ProtoMessage()               {}
func (*DeleteCandidateResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type ListCandidateRequest struct {
}

func (m *ListCandidateRequest) Reset()                    { *m = ListCandidateRequest{} }
func (m *ListCandidateRequest) String() string            { return proto1.CompactTextString(m) }
func (*ListCandidateRequest) ProtoMessage()               {}
func (*ListCandidateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

type ListCandidateResponse struct {
	Data []*Candidate `protobuf:"bytes,1,rep,name=data" json:"data,omitempty"`
}

func (m *ListCandidateResponse) Reset()                    { *m = ListCandidateResponse{} }
func (m *ListCandidateResponse) String() string            { return proto1.CompactTextString(m) }
func (*ListCandidateResponse) ProtoMessage()               {}
func (*ListCandidateResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *ListCandidateResponse) GetData() []*Candidate {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto1.RegisterType((*Candidate)(nil), "proto.Candidate")
	proto1.RegisterType((*AddCandidateRequest)(nil), "proto.AddCandidateRequest")
	proto1.RegisterType((*CandidateResponse)(nil), "proto.CandidateResponse")
	proto1.RegisterType((*DeleteCandidateRequest)(nil), "proto.DeleteCandidateRequest")
	proto1.RegisterType((*DeleteCandidateResponse)(nil), "proto.DeleteCandidateResponse")
	proto1.RegisterType((*ListCandidateRequest)(nil), "proto.ListCandidateRequest")
	proto1.RegisterType((*ListCandidateResponse)(nil), "proto.ListCandidateResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for Poll service

type PollClient interface {
	AddCandidate(ctx context.Context, in *AddCandidateRequest, opts ...grpc.CallOption) (*CandidateResponse, error)
	DeleteCandidate(ctx context.Context, in *DeleteCandidateRequest, opts ...grpc.CallOption) (*DeleteCandidateResponse, error)
	ListCandidate(ctx context.Context, in *ListCandidateRequest, opts ...grpc.CallOption) (*ListCandidateResponse, error)
}

type pollClient struct {
	cc *grpc.ClientConn
}

func NewPollClient(cc *grpc.ClientConn) PollClient {
	return &pollClient{cc}
}

func (c *pollClient) AddCandidate(ctx context.Context, in *AddCandidateRequest, opts ...grpc.CallOption) (*CandidateResponse, error) {
	out := new(CandidateResponse)
	err := grpc.Invoke(ctx, "/proto.Poll/AddCandidate", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pollClient) DeleteCandidate(ctx context.Context, in *DeleteCandidateRequest, opts ...grpc.CallOption) (*DeleteCandidateResponse, error) {
	out := new(DeleteCandidateResponse)
	err := grpc.Invoke(ctx, "/proto.Poll/DeleteCandidate", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pollClient) ListCandidate(ctx context.Context, in *ListCandidateRequest, opts ...grpc.CallOption) (*ListCandidateResponse, error) {
	out := new(ListCandidateResponse)
	err := grpc.Invoke(ctx, "/proto.Poll/ListCandidate", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Poll service

type PollServer interface {
	AddCandidate(context.Context, *AddCandidateRequest) (*CandidateResponse, error)
	DeleteCandidate(context.Context, *DeleteCandidateRequest) (*DeleteCandidateResponse, error)
	ListCandidate(context.Context, *ListCandidateRequest) (*ListCandidateResponse, error)
}

func RegisterPollServer(s *grpc.Server, srv PollServer) {
	s.RegisterService(&_Poll_serviceDesc, srv)
}

func _Poll_AddCandidate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddCandidateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PollServer).AddCandidate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Poll/AddCandidate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PollServer).AddCandidate(ctx, req.(*AddCandidateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Poll_DeleteCandidate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCandidateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PollServer).DeleteCandidate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Poll/DeleteCandidate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PollServer).DeleteCandidate(ctx, req.(*DeleteCandidateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Poll_ListCandidate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListCandidateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PollServer).ListCandidate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Poll/ListCandidate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PollServer).ListCandidate(ctx, req.(*ListCandidateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Poll_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Poll",
	HandlerType: (*PollServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddCandidate",
			Handler:    _Poll_AddCandidate_Handler,
		},
		{
			MethodName: "DeleteCandidate",
			Handler:    _Poll_DeleteCandidate_Handler,
		},
		{
			MethodName: "ListCandidate",
			Handler:    _Poll_ListCandidate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor0,
}

func init() { proto1.RegisterFile("proto/service.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 275 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x50, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x35, 0x36, 0x11, 0x32, 0x2a, 0xea, 0xb4, 0xd6, 0x18, 0xad, 0xc8, 0xe2, 0x41, 0x41, 0x2a,
	0xd4, 0x93, 0x07, 0x0f, 0x45, 0xf1, 0xd4, 0x83, 0x04, 0xef, 0xb2, 0xba, 0x73, 0x08, 0xc4, 0x6c,
	0xec, 0x4e, 0xfb, 0xdd, 0x7e, 0x82, 0x75, 0x93, 0xd6, 0x9a, 0x6c, 0xc0, 0xd3, 0xee, 0xcc, 0x7b,
	0xf3, 0xe6, 0xcd, 0x83, 0x6e, 0x31, 0xd5, 0xac, 0x6f, 0x0c, 0x4d, 0xe7, 0xe9, 0x3b, 0x0d, 0x6d,
	0x85, 0x81, 0x7d, 0xc4, 0x0b, 0x84, 0x0f, 0x32, 0x57, 0xa9, 0x92, 0x4c, 0x88, 0xe0, 0xe7, 0xf2,
	0x83, 0x22, 0xef, 0xdc, 0xbb, 0x0c, 0x13, 0xfb, 0xc7, 0x1e, 0x04, 0x73, 0xcd, 0x64, 0xa2, 0xcd,
	0x45, 0x33, 0x48, 0xca, 0x02, 0x07, 0x00, 0xb3, 0xe2, 0x67, 0x46, 0xbd, 0x4a, 0x8e, 0x3a, 0x16,
	0x0a, 0xab, 0xce, 0x98, 0xc5, 0x15, 0x74, 0xc7, 0x4a, 0xad, 0x84, 0x13, 0xfa, 0x9c, 0x91, 0x61,
	0x97, 0xbe, 0xb8, 0x83, 0x83, 0x35, 0x9e, 0x29, 0x74, 0x6e, 0x08, 0x2f, 0xc0, 0x5f, 0xd4, 0xd2,
	0x12, 0xb7, 0x47, 0xfb, 0xa5, 0xe5, 0xe1, 0x2f, 0xcf, 0xa2, 0xe2, 0x1a, 0xfa, 0x8f, 0x94, 0x11,
	0xd3, 0xbf, 0x16, 0x1d, 0xc3, 0x51, 0x83, 0x5d, 0xae, 0x13, 0x7d, 0xe8, 0x4d, 0x52, 0xc3, 0x75,
	0x19, 0x71, 0x0f, 0x87, 0xb5, 0x7e, 0xc3, 0x5f, 0xa7, 0xdd, 0xdf, 0xe8, 0xcb, 0x03, 0xff, 0x59,
	0x67, 0x19, 0x3e, 0xc1, 0xce, 0x7a, 0x1c, 0x18, 0x57, 0x03, 0x8e, 0x8c, 0xe2, 0xa8, 0x21, 0xb6,
	0x74, 0xb9, 0x81, 0x09, 0xec, 0xd5, 0x4e, 0xc0, 0x41, 0x45, 0x77, 0x07, 0x11, 0x9f, 0xb5, 0xc1,
	0x2b, 0xcd, 0x09, 0xec, 0xfe, 0xb9, 0x11, 0x4f, 0xaa, 0x11, 0x57, 0x22, 0xf1, 0xa9, 0x1b, 0x5c,
	0xaa, 0xbd, 0x6d, 0x59, 0xf8, 0xf6, 0x3b, 0x00, 0x00, 0xff, 0xff, 0xde, 0x1d, 0x0d, 0xdb, 0x73,
	0x02, 0x00, 0x00,
}
