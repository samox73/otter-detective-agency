// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.0
// source: case/case.proto

package casepb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	CaseService_CreateCase_FullMethodName         = "/casepb.CaseService/CreateCase"
	CaseService_GetCase_FullMethodName            = "/casepb.CaseService/GetCase"
	CaseService_ListCases_FullMethodName          = "/casepb.CaseService/ListCases"
	CaseService_AssignCaseToPlayer_FullMethodName = "/casepb.CaseService/AssignCaseToPlayer"
	CaseService_GetPlayerCase_FullMethodName      = "/casepb.CaseService/GetPlayerCase"
)

// CaseServiceClient is the client API for CaseService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CaseServiceClient interface {
	CreateCase(ctx context.Context, in *CreateCaseRequest, opts ...grpc.CallOption) (*Case, error)
	GetCase(ctx context.Context, in *GetCaseRequest, opts ...grpc.CallOption) (*Case, error)
	ListCases(ctx context.Context, in *ListCasesRequest, opts ...grpc.CallOption) (*CaseList, error)
	AssignCaseToPlayer(ctx context.Context, in *AssignCaseRequest, opts ...grpc.CallOption) (*PlayerCase, error)
	GetPlayerCase(ctx context.Context, in *GetPlayerCaseRequest, opts ...grpc.CallOption) (*GetPlayerCaseResponse, error)
}

type caseServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCaseServiceClient(cc grpc.ClientConnInterface) CaseServiceClient {
	return &caseServiceClient{cc}
}

func (c *caseServiceClient) CreateCase(ctx context.Context, in *CreateCaseRequest, opts ...grpc.CallOption) (*Case, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Case)
	err := c.cc.Invoke(ctx, CaseService_CreateCase_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *caseServiceClient) GetCase(ctx context.Context, in *GetCaseRequest, opts ...grpc.CallOption) (*Case, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Case)
	err := c.cc.Invoke(ctx, CaseService_GetCase_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *caseServiceClient) ListCases(ctx context.Context, in *ListCasesRequest, opts ...grpc.CallOption) (*CaseList, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CaseList)
	err := c.cc.Invoke(ctx, CaseService_ListCases_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *caseServiceClient) AssignCaseToPlayer(ctx context.Context, in *AssignCaseRequest, opts ...grpc.CallOption) (*PlayerCase, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PlayerCase)
	err := c.cc.Invoke(ctx, CaseService_AssignCaseToPlayer_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *caseServiceClient) GetPlayerCase(ctx context.Context, in *GetPlayerCaseRequest, opts ...grpc.CallOption) (*GetPlayerCaseResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetPlayerCaseResponse)
	err := c.cc.Invoke(ctx, CaseService_GetPlayerCase_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CaseServiceServer is the server API for CaseService service.
// All implementations must embed UnimplementedCaseServiceServer
// for forward compatibility.
type CaseServiceServer interface {
	CreateCase(context.Context, *CreateCaseRequest) (*Case, error)
	GetCase(context.Context, *GetCaseRequest) (*Case, error)
	ListCases(context.Context, *ListCasesRequest) (*CaseList, error)
	AssignCaseToPlayer(context.Context, *AssignCaseRequest) (*PlayerCase, error)
	GetPlayerCase(context.Context, *GetPlayerCaseRequest) (*GetPlayerCaseResponse, error)
	mustEmbedUnimplementedCaseServiceServer()
}

// UnimplementedCaseServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedCaseServiceServer struct{}

func (UnimplementedCaseServiceServer) CreateCase(context.Context, *CreateCaseRequest) (*Case, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCase not implemented")
}
func (UnimplementedCaseServiceServer) GetCase(context.Context, *GetCaseRequest) (*Case, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCase not implemented")
}
func (UnimplementedCaseServiceServer) ListCases(context.Context, *ListCasesRequest) (*CaseList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCases not implemented")
}
func (UnimplementedCaseServiceServer) AssignCaseToPlayer(context.Context, *AssignCaseRequest) (*PlayerCase, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AssignCaseToPlayer not implemented")
}
func (UnimplementedCaseServiceServer) GetPlayerCase(context.Context, *GetPlayerCaseRequest) (*GetPlayerCaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPlayerCase not implemented")
}
func (UnimplementedCaseServiceServer) mustEmbedUnimplementedCaseServiceServer() {}
func (UnimplementedCaseServiceServer) testEmbeddedByValue()                     {}

// UnsafeCaseServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CaseServiceServer will
// result in compilation errors.
type UnsafeCaseServiceServer interface {
	mustEmbedUnimplementedCaseServiceServer()
}

func RegisterCaseServiceServer(s grpc.ServiceRegistrar, srv CaseServiceServer) {
	// If the following call pancis, it indicates UnimplementedCaseServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&CaseService_ServiceDesc, srv)
}

func _CaseService_CreateCase_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CaseServiceServer).CreateCase(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CaseService_CreateCase_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CaseServiceServer).CreateCase(ctx, req.(*CreateCaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CaseService_GetCase_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CaseServiceServer).GetCase(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CaseService_GetCase_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CaseServiceServer).GetCase(ctx, req.(*GetCaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CaseService_ListCases_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListCasesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CaseServiceServer).ListCases(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CaseService_ListCases_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CaseServiceServer).ListCases(ctx, req.(*ListCasesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CaseService_AssignCaseToPlayer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AssignCaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CaseServiceServer).AssignCaseToPlayer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CaseService_AssignCaseToPlayer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CaseServiceServer).AssignCaseToPlayer(ctx, req.(*AssignCaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CaseService_GetPlayerCase_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPlayerCaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CaseServiceServer).GetPlayerCase(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CaseService_GetPlayerCase_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CaseServiceServer).GetPlayerCase(ctx, req.(*GetPlayerCaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CaseService_ServiceDesc is the grpc.ServiceDesc for CaseService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CaseService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "casepb.CaseService",
	HandlerType: (*CaseServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCase",
			Handler:    _CaseService_CreateCase_Handler,
		},
		{
			MethodName: "GetCase",
			Handler:    _CaseService_GetCase_Handler,
		},
		{
			MethodName: "ListCases",
			Handler:    _CaseService_ListCases_Handler,
		},
		{
			MethodName: "AssignCaseToPlayer",
			Handler:    _CaseService_AssignCaseToPlayer_Handler,
		},
		{
			MethodName: "GetPlayerCase",
			Handler:    _CaseService_GetPlayerCase_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "case/case.proto",
}
