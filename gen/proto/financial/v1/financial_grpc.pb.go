// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: proto/financial/v1/financial.proto

package v1

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
	FinancialPeriodService_CreateFinancialPeriod_FullMethodName                 = "/financial.v1.FinancialPeriodService/CreateFinancialPeriod"
	FinancialPeriodService_GetFinancialPeriod_FullMethodName                    = "/financial.v1.FinancialPeriodService/GetFinancialPeriod"
	FinancialPeriodService_UpdateFinancialPeriod_FullMethodName                 = "/financial.v1.FinancialPeriodService/UpdateFinancialPeriod"
	FinancialPeriodService_DeleteFinancialPeriod_FullMethodName                 = "/financial.v1.FinancialPeriodService/DeleteFinancialPeriod"
	FinancialPeriodService_ListFinancialPeriods_FullMethodName                  = "/financial.v1.FinancialPeriodService/ListFinancialPeriods"
	FinancialPeriodService_GetFinancialPeriodsByBusinessID_FullMethodName       = "/financial.v1.FinancialPeriodService/GetFinancialPeriodsByBusinessID"
	FinancialPeriodService_GetActiveFinancialPeriodsByBusinessID_FullMethodName = "/financial.v1.FinancialPeriodService/GetActiveFinancialPeriodsByBusinessID"
	FinancialPeriodService_GetFinancialPeriodByBusinessAndName_FullMethodName   = "/financial.v1.FinancialPeriodService/GetFinancialPeriodByBusinessAndName"
	FinancialPeriodService_GetFinancialPeriodsByDateRange_FullMethodName        = "/financial.v1.FinancialPeriodService/GetFinancialPeriodsByDateRange"
	FinancialPeriodService_GetOverlappingFinancialPeriods_FullMethodName        = "/financial.v1.FinancialPeriodService/GetOverlappingFinancialPeriods"
	FinancialPeriodService_CloseFinancialPeriod_FullMethodName                  = "/financial.v1.FinancialPeriodService/CloseFinancialPeriod"
	FinancialPeriodService_ReopenFinancialPeriod_FullMethodName                 = "/financial.v1.FinancialPeriodService/ReopenFinancialPeriod"
	FinancialPeriodService_GetFinancialPeriodsByUserID_FullMethodName           = "/financial.v1.FinancialPeriodService/GetFinancialPeriodsByUserID"
	FinancialPeriodService_AddUserToFinancialPeriod_FullMethodName              = "/financial.v1.FinancialPeriodService/AddUserToFinancialPeriod"
	FinancialPeriodService_RemoveUserFromFinancialPeriod_FullMethodName         = "/financial.v1.FinancialPeriodService/RemoveUserFromFinancialPeriod"
)

// FinancialPeriodServiceClient is the client API for FinancialPeriodService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// FinancialPeriodService provides financial period management functionality
type FinancialPeriodServiceClient interface {
	// Basic CRUD operations
	CreateFinancialPeriod(ctx context.Context, in *CreateFinancialPeriodRequest, opts ...grpc.CallOption) (*CreateFinancialPeriodResponse, error)
	GetFinancialPeriod(ctx context.Context, in *GetFinancialPeriodRequest, opts ...grpc.CallOption) (*GetFinancialPeriodResponse, error)
	UpdateFinancialPeriod(ctx context.Context, in *UpdateFinancialPeriodRequest, opts ...grpc.CallOption) (*UpdateFinancialPeriodResponse, error)
	DeleteFinancialPeriod(ctx context.Context, in *DeleteFinancialPeriodRequest, opts ...grpc.CallOption) (*DeleteFinancialPeriodResponse, error)
	ListFinancialPeriods(ctx context.Context, in *ListFinancialPeriodsRequest, opts ...grpc.CallOption) (*ListFinancialPeriodsResponse, error)
	// Business related operations
	GetFinancialPeriodsByBusinessID(ctx context.Context, in *GetFinancialPeriodsByBusinessIDRequest, opts ...grpc.CallOption) (*GetFinancialPeriodsByBusinessIDResponse, error)
	GetActiveFinancialPeriodsByBusinessID(ctx context.Context, in *GetActiveFinancialPeriodsByBusinessIDRequest, opts ...grpc.CallOption) (*GetActiveFinancialPeriodsByBusinessIDResponse, error)
	GetFinancialPeriodByBusinessAndName(ctx context.Context, in *GetFinancialPeriodByBusinessAndNameRequest, opts ...grpc.CallOption) (*GetFinancialPeriodResponse, error)
	// Time related operations
	GetFinancialPeriodsByDateRange(ctx context.Context, in *GetFinancialPeriodsByDateRangeRequest, opts ...grpc.CallOption) (*GetFinancialPeriodsByDateRangeResponse, error)
	GetOverlappingFinancialPeriods(ctx context.Context, in *GetOverlappingFinancialPeriodsRequest, opts ...grpc.CallOption) (*GetOverlappingFinancialPeriodsResponse, error)
	// Status operations
	CloseFinancialPeriod(ctx context.Context, in *CloseFinancialPeriodRequest, opts ...grpc.CallOption) (*CloseFinancialPeriodResponse, error)
	ReopenFinancialPeriod(ctx context.Context, in *ReopenFinancialPeriodRequest, opts ...grpc.CallOption) (*ReopenFinancialPeriodResponse, error)
	// User access operations
	GetFinancialPeriodsByUserID(ctx context.Context, in *GetFinancialPeriodsByUserIDRequest, opts ...grpc.CallOption) (*GetFinancialPeriodsByUserIDResponse, error)
	AddUserToFinancialPeriod(ctx context.Context, in *AddUserToFinancialPeriodRequest, opts ...grpc.CallOption) (*AddUserToFinancialPeriodResponse, error)
	RemoveUserFromFinancialPeriod(ctx context.Context, in *RemoveUserFromFinancialPeriodRequest, opts ...grpc.CallOption) (*RemoveUserFromFinancialPeriodResponse, error)
}

type financialPeriodServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFinancialPeriodServiceClient(cc grpc.ClientConnInterface) FinancialPeriodServiceClient {
	return &financialPeriodServiceClient{cc}
}

func (c *financialPeriodServiceClient) CreateFinancialPeriod(ctx context.Context, in *CreateFinancialPeriodRequest, opts ...grpc.CallOption) (*CreateFinancialPeriodResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateFinancialPeriodResponse)
	err := c.cc.Invoke(ctx, FinancialPeriodService_CreateFinancialPeriod_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *financialPeriodServiceClient) GetFinancialPeriod(ctx context.Context, in *GetFinancialPeriodRequest, opts ...grpc.CallOption) (*GetFinancialPeriodResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetFinancialPeriodResponse)
	err := c.cc.Invoke(ctx, FinancialPeriodService_GetFinancialPeriod_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *financialPeriodServiceClient) UpdateFinancialPeriod(ctx context.Context, in *UpdateFinancialPeriodRequest, opts ...grpc.CallOption) (*UpdateFinancialPeriodResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateFinancialPeriodResponse)
	err := c.cc.Invoke(ctx, FinancialPeriodService_UpdateFinancialPeriod_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *financialPeriodServiceClient) DeleteFinancialPeriod(ctx context.Context, in *DeleteFinancialPeriodRequest, opts ...grpc.CallOption) (*DeleteFinancialPeriodResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteFinancialPeriodResponse)
	err := c.cc.Invoke(ctx, FinancialPeriodService_DeleteFinancialPeriod_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *financialPeriodServiceClient) ListFinancialPeriods(ctx context.Context, in *ListFinancialPeriodsRequest, opts ...grpc.CallOption) (*ListFinancialPeriodsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListFinancialPeriodsResponse)
	err := c.cc.Invoke(ctx, FinancialPeriodService_ListFinancialPeriods_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *financialPeriodServiceClient) GetFinancialPeriodsByBusinessID(ctx context.Context, in *GetFinancialPeriodsByBusinessIDRequest, opts ...grpc.CallOption) (*GetFinancialPeriodsByBusinessIDResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetFinancialPeriodsByBusinessIDResponse)
	err := c.cc.Invoke(ctx, FinancialPeriodService_GetFinancialPeriodsByBusinessID_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *financialPeriodServiceClient) GetActiveFinancialPeriodsByBusinessID(ctx context.Context, in *GetActiveFinancialPeriodsByBusinessIDRequest, opts ...grpc.CallOption) (*GetActiveFinancialPeriodsByBusinessIDResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetActiveFinancialPeriodsByBusinessIDResponse)
	err := c.cc.Invoke(ctx, FinancialPeriodService_GetActiveFinancialPeriodsByBusinessID_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *financialPeriodServiceClient) GetFinancialPeriodByBusinessAndName(ctx context.Context, in *GetFinancialPeriodByBusinessAndNameRequest, opts ...grpc.CallOption) (*GetFinancialPeriodResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetFinancialPeriodResponse)
	err := c.cc.Invoke(ctx, FinancialPeriodService_GetFinancialPeriodByBusinessAndName_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *financialPeriodServiceClient) GetFinancialPeriodsByDateRange(ctx context.Context, in *GetFinancialPeriodsByDateRangeRequest, opts ...grpc.CallOption) (*GetFinancialPeriodsByDateRangeResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetFinancialPeriodsByDateRangeResponse)
	err := c.cc.Invoke(ctx, FinancialPeriodService_GetFinancialPeriodsByDateRange_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *financialPeriodServiceClient) GetOverlappingFinancialPeriods(ctx context.Context, in *GetOverlappingFinancialPeriodsRequest, opts ...grpc.CallOption) (*GetOverlappingFinancialPeriodsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetOverlappingFinancialPeriodsResponse)
	err := c.cc.Invoke(ctx, FinancialPeriodService_GetOverlappingFinancialPeriods_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *financialPeriodServiceClient) CloseFinancialPeriod(ctx context.Context, in *CloseFinancialPeriodRequest, opts ...grpc.CallOption) (*CloseFinancialPeriodResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CloseFinancialPeriodResponse)
	err := c.cc.Invoke(ctx, FinancialPeriodService_CloseFinancialPeriod_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *financialPeriodServiceClient) ReopenFinancialPeriod(ctx context.Context, in *ReopenFinancialPeriodRequest, opts ...grpc.CallOption) (*ReopenFinancialPeriodResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ReopenFinancialPeriodResponse)
	err := c.cc.Invoke(ctx, FinancialPeriodService_ReopenFinancialPeriod_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *financialPeriodServiceClient) GetFinancialPeriodsByUserID(ctx context.Context, in *GetFinancialPeriodsByUserIDRequest, opts ...grpc.CallOption) (*GetFinancialPeriodsByUserIDResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetFinancialPeriodsByUserIDResponse)
	err := c.cc.Invoke(ctx, FinancialPeriodService_GetFinancialPeriodsByUserID_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *financialPeriodServiceClient) AddUserToFinancialPeriod(ctx context.Context, in *AddUserToFinancialPeriodRequest, opts ...grpc.CallOption) (*AddUserToFinancialPeriodResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddUserToFinancialPeriodResponse)
	err := c.cc.Invoke(ctx, FinancialPeriodService_AddUserToFinancialPeriod_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *financialPeriodServiceClient) RemoveUserFromFinancialPeriod(ctx context.Context, in *RemoveUserFromFinancialPeriodRequest, opts ...grpc.CallOption) (*RemoveUserFromFinancialPeriodResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RemoveUserFromFinancialPeriodResponse)
	err := c.cc.Invoke(ctx, FinancialPeriodService_RemoveUserFromFinancialPeriod_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FinancialPeriodServiceServer is the server API for FinancialPeriodService service.
// All implementations must embed UnimplementedFinancialPeriodServiceServer
// for forward compatibility.
//
// FinancialPeriodService provides financial period management functionality
type FinancialPeriodServiceServer interface {
	// Basic CRUD operations
	CreateFinancialPeriod(context.Context, *CreateFinancialPeriodRequest) (*CreateFinancialPeriodResponse, error)
	GetFinancialPeriod(context.Context, *GetFinancialPeriodRequest) (*GetFinancialPeriodResponse, error)
	UpdateFinancialPeriod(context.Context, *UpdateFinancialPeriodRequest) (*UpdateFinancialPeriodResponse, error)
	DeleteFinancialPeriod(context.Context, *DeleteFinancialPeriodRequest) (*DeleteFinancialPeriodResponse, error)
	ListFinancialPeriods(context.Context, *ListFinancialPeriodsRequest) (*ListFinancialPeriodsResponse, error)
	// Business related operations
	GetFinancialPeriodsByBusinessID(context.Context, *GetFinancialPeriodsByBusinessIDRequest) (*GetFinancialPeriodsByBusinessIDResponse, error)
	GetActiveFinancialPeriodsByBusinessID(context.Context, *GetActiveFinancialPeriodsByBusinessIDRequest) (*GetActiveFinancialPeriodsByBusinessIDResponse, error)
	GetFinancialPeriodByBusinessAndName(context.Context, *GetFinancialPeriodByBusinessAndNameRequest) (*GetFinancialPeriodResponse, error)
	// Time related operations
	GetFinancialPeriodsByDateRange(context.Context, *GetFinancialPeriodsByDateRangeRequest) (*GetFinancialPeriodsByDateRangeResponse, error)
	GetOverlappingFinancialPeriods(context.Context, *GetOverlappingFinancialPeriodsRequest) (*GetOverlappingFinancialPeriodsResponse, error)
	// Status operations
	CloseFinancialPeriod(context.Context, *CloseFinancialPeriodRequest) (*CloseFinancialPeriodResponse, error)
	ReopenFinancialPeriod(context.Context, *ReopenFinancialPeriodRequest) (*ReopenFinancialPeriodResponse, error)
	// User access operations
	GetFinancialPeriodsByUserID(context.Context, *GetFinancialPeriodsByUserIDRequest) (*GetFinancialPeriodsByUserIDResponse, error)
	AddUserToFinancialPeriod(context.Context, *AddUserToFinancialPeriodRequest) (*AddUserToFinancialPeriodResponse, error)
	RemoveUserFromFinancialPeriod(context.Context, *RemoveUserFromFinancialPeriodRequest) (*RemoveUserFromFinancialPeriodResponse, error)
	mustEmbedUnimplementedFinancialPeriodServiceServer()
}

// UnimplementedFinancialPeriodServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedFinancialPeriodServiceServer struct{}

func (UnimplementedFinancialPeriodServiceServer) CreateFinancialPeriod(context.Context, *CreateFinancialPeriodRequest) (*CreateFinancialPeriodResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateFinancialPeriod not implemented")
}
func (UnimplementedFinancialPeriodServiceServer) GetFinancialPeriod(context.Context, *GetFinancialPeriodRequest) (*GetFinancialPeriodResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFinancialPeriod not implemented")
}
func (UnimplementedFinancialPeriodServiceServer) UpdateFinancialPeriod(context.Context, *UpdateFinancialPeriodRequest) (*UpdateFinancialPeriodResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateFinancialPeriod not implemented")
}
func (UnimplementedFinancialPeriodServiceServer) DeleteFinancialPeriod(context.Context, *DeleteFinancialPeriodRequest) (*DeleteFinancialPeriodResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteFinancialPeriod not implemented")
}
func (UnimplementedFinancialPeriodServiceServer) ListFinancialPeriods(context.Context, *ListFinancialPeriodsRequest) (*ListFinancialPeriodsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListFinancialPeriods not implemented")
}
func (UnimplementedFinancialPeriodServiceServer) GetFinancialPeriodsByBusinessID(context.Context, *GetFinancialPeriodsByBusinessIDRequest) (*GetFinancialPeriodsByBusinessIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFinancialPeriodsByBusinessID not implemented")
}
func (UnimplementedFinancialPeriodServiceServer) GetActiveFinancialPeriodsByBusinessID(context.Context, *GetActiveFinancialPeriodsByBusinessIDRequest) (*GetActiveFinancialPeriodsByBusinessIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetActiveFinancialPeriodsByBusinessID not implemented")
}
func (UnimplementedFinancialPeriodServiceServer) GetFinancialPeriodByBusinessAndName(context.Context, *GetFinancialPeriodByBusinessAndNameRequest) (*GetFinancialPeriodResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFinancialPeriodByBusinessAndName not implemented")
}
func (UnimplementedFinancialPeriodServiceServer) GetFinancialPeriodsByDateRange(context.Context, *GetFinancialPeriodsByDateRangeRequest) (*GetFinancialPeriodsByDateRangeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFinancialPeriodsByDateRange not implemented")
}
func (UnimplementedFinancialPeriodServiceServer) GetOverlappingFinancialPeriods(context.Context, *GetOverlappingFinancialPeriodsRequest) (*GetOverlappingFinancialPeriodsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOverlappingFinancialPeriods not implemented")
}
func (UnimplementedFinancialPeriodServiceServer) CloseFinancialPeriod(context.Context, *CloseFinancialPeriodRequest) (*CloseFinancialPeriodResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CloseFinancialPeriod not implemented")
}
func (UnimplementedFinancialPeriodServiceServer) ReopenFinancialPeriod(context.Context, *ReopenFinancialPeriodRequest) (*ReopenFinancialPeriodResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReopenFinancialPeriod not implemented")
}
func (UnimplementedFinancialPeriodServiceServer) GetFinancialPeriodsByUserID(context.Context, *GetFinancialPeriodsByUserIDRequest) (*GetFinancialPeriodsByUserIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFinancialPeriodsByUserID not implemented")
}
func (UnimplementedFinancialPeriodServiceServer) AddUserToFinancialPeriod(context.Context, *AddUserToFinancialPeriodRequest) (*AddUserToFinancialPeriodResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddUserToFinancialPeriod not implemented")
}
func (UnimplementedFinancialPeriodServiceServer) RemoveUserFromFinancialPeriod(context.Context, *RemoveUserFromFinancialPeriodRequest) (*RemoveUserFromFinancialPeriodResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveUserFromFinancialPeriod not implemented")
}
func (UnimplementedFinancialPeriodServiceServer) mustEmbedUnimplementedFinancialPeriodServiceServer() {
}
func (UnimplementedFinancialPeriodServiceServer) testEmbeddedByValue() {}

// UnsafeFinancialPeriodServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FinancialPeriodServiceServer will
// result in compilation errors.
type UnsafeFinancialPeriodServiceServer interface {
	mustEmbedUnimplementedFinancialPeriodServiceServer()
}

func RegisterFinancialPeriodServiceServer(s grpc.ServiceRegistrar, srv FinancialPeriodServiceServer) {
	// If the following call pancis, it indicates UnimplementedFinancialPeriodServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&FinancialPeriodService_ServiceDesc, srv)
}

func _FinancialPeriodService_CreateFinancialPeriod_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateFinancialPeriodRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FinancialPeriodServiceServer).CreateFinancialPeriod(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FinancialPeriodService_CreateFinancialPeriod_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FinancialPeriodServiceServer).CreateFinancialPeriod(ctx, req.(*CreateFinancialPeriodRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FinancialPeriodService_GetFinancialPeriod_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFinancialPeriodRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FinancialPeriodServiceServer).GetFinancialPeriod(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FinancialPeriodService_GetFinancialPeriod_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FinancialPeriodServiceServer).GetFinancialPeriod(ctx, req.(*GetFinancialPeriodRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FinancialPeriodService_UpdateFinancialPeriod_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateFinancialPeriodRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FinancialPeriodServiceServer).UpdateFinancialPeriod(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FinancialPeriodService_UpdateFinancialPeriod_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FinancialPeriodServiceServer).UpdateFinancialPeriod(ctx, req.(*UpdateFinancialPeriodRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FinancialPeriodService_DeleteFinancialPeriod_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteFinancialPeriodRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FinancialPeriodServiceServer).DeleteFinancialPeriod(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FinancialPeriodService_DeleteFinancialPeriod_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FinancialPeriodServiceServer).DeleteFinancialPeriod(ctx, req.(*DeleteFinancialPeriodRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FinancialPeriodService_ListFinancialPeriods_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListFinancialPeriodsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FinancialPeriodServiceServer).ListFinancialPeriods(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FinancialPeriodService_ListFinancialPeriods_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FinancialPeriodServiceServer).ListFinancialPeriods(ctx, req.(*ListFinancialPeriodsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FinancialPeriodService_GetFinancialPeriodsByBusinessID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFinancialPeriodsByBusinessIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FinancialPeriodServiceServer).GetFinancialPeriodsByBusinessID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FinancialPeriodService_GetFinancialPeriodsByBusinessID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FinancialPeriodServiceServer).GetFinancialPeriodsByBusinessID(ctx, req.(*GetFinancialPeriodsByBusinessIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FinancialPeriodService_GetActiveFinancialPeriodsByBusinessID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetActiveFinancialPeriodsByBusinessIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FinancialPeriodServiceServer).GetActiveFinancialPeriodsByBusinessID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FinancialPeriodService_GetActiveFinancialPeriodsByBusinessID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FinancialPeriodServiceServer).GetActiveFinancialPeriodsByBusinessID(ctx, req.(*GetActiveFinancialPeriodsByBusinessIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FinancialPeriodService_GetFinancialPeriodByBusinessAndName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFinancialPeriodByBusinessAndNameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FinancialPeriodServiceServer).GetFinancialPeriodByBusinessAndName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FinancialPeriodService_GetFinancialPeriodByBusinessAndName_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FinancialPeriodServiceServer).GetFinancialPeriodByBusinessAndName(ctx, req.(*GetFinancialPeriodByBusinessAndNameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FinancialPeriodService_GetFinancialPeriodsByDateRange_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFinancialPeriodsByDateRangeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FinancialPeriodServiceServer).GetFinancialPeriodsByDateRange(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FinancialPeriodService_GetFinancialPeriodsByDateRange_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FinancialPeriodServiceServer).GetFinancialPeriodsByDateRange(ctx, req.(*GetFinancialPeriodsByDateRangeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FinancialPeriodService_GetOverlappingFinancialPeriods_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOverlappingFinancialPeriodsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FinancialPeriodServiceServer).GetOverlappingFinancialPeriods(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FinancialPeriodService_GetOverlappingFinancialPeriods_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FinancialPeriodServiceServer).GetOverlappingFinancialPeriods(ctx, req.(*GetOverlappingFinancialPeriodsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FinancialPeriodService_CloseFinancialPeriod_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CloseFinancialPeriodRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FinancialPeriodServiceServer).CloseFinancialPeriod(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FinancialPeriodService_CloseFinancialPeriod_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FinancialPeriodServiceServer).CloseFinancialPeriod(ctx, req.(*CloseFinancialPeriodRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FinancialPeriodService_ReopenFinancialPeriod_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReopenFinancialPeriodRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FinancialPeriodServiceServer).ReopenFinancialPeriod(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FinancialPeriodService_ReopenFinancialPeriod_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FinancialPeriodServiceServer).ReopenFinancialPeriod(ctx, req.(*ReopenFinancialPeriodRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FinancialPeriodService_GetFinancialPeriodsByUserID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFinancialPeriodsByUserIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FinancialPeriodServiceServer).GetFinancialPeriodsByUserID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FinancialPeriodService_GetFinancialPeriodsByUserID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FinancialPeriodServiceServer).GetFinancialPeriodsByUserID(ctx, req.(*GetFinancialPeriodsByUserIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FinancialPeriodService_AddUserToFinancialPeriod_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddUserToFinancialPeriodRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FinancialPeriodServiceServer).AddUserToFinancialPeriod(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FinancialPeriodService_AddUserToFinancialPeriod_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FinancialPeriodServiceServer).AddUserToFinancialPeriod(ctx, req.(*AddUserToFinancialPeriodRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FinancialPeriodService_RemoveUserFromFinancialPeriod_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveUserFromFinancialPeriodRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FinancialPeriodServiceServer).RemoveUserFromFinancialPeriod(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FinancialPeriodService_RemoveUserFromFinancialPeriod_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FinancialPeriodServiceServer).RemoveUserFromFinancialPeriod(ctx, req.(*RemoveUserFromFinancialPeriodRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// FinancialPeriodService_ServiceDesc is the grpc.ServiceDesc for FinancialPeriodService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FinancialPeriodService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "financial.v1.FinancialPeriodService",
	HandlerType: (*FinancialPeriodServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateFinancialPeriod",
			Handler:    _FinancialPeriodService_CreateFinancialPeriod_Handler,
		},
		{
			MethodName: "GetFinancialPeriod",
			Handler:    _FinancialPeriodService_GetFinancialPeriod_Handler,
		},
		{
			MethodName: "UpdateFinancialPeriod",
			Handler:    _FinancialPeriodService_UpdateFinancialPeriod_Handler,
		},
		{
			MethodName: "DeleteFinancialPeriod",
			Handler:    _FinancialPeriodService_DeleteFinancialPeriod_Handler,
		},
		{
			MethodName: "ListFinancialPeriods",
			Handler:    _FinancialPeriodService_ListFinancialPeriods_Handler,
		},
		{
			MethodName: "GetFinancialPeriodsByBusinessID",
			Handler:    _FinancialPeriodService_GetFinancialPeriodsByBusinessID_Handler,
		},
		{
			MethodName: "GetActiveFinancialPeriodsByBusinessID",
			Handler:    _FinancialPeriodService_GetActiveFinancialPeriodsByBusinessID_Handler,
		},
		{
			MethodName: "GetFinancialPeriodByBusinessAndName",
			Handler:    _FinancialPeriodService_GetFinancialPeriodByBusinessAndName_Handler,
		},
		{
			MethodName: "GetFinancialPeriodsByDateRange",
			Handler:    _FinancialPeriodService_GetFinancialPeriodsByDateRange_Handler,
		},
		{
			MethodName: "GetOverlappingFinancialPeriods",
			Handler:    _FinancialPeriodService_GetOverlappingFinancialPeriods_Handler,
		},
		{
			MethodName: "CloseFinancialPeriod",
			Handler:    _FinancialPeriodService_CloseFinancialPeriod_Handler,
		},
		{
			MethodName: "ReopenFinancialPeriod",
			Handler:    _FinancialPeriodService_ReopenFinancialPeriod_Handler,
		},
		{
			MethodName: "GetFinancialPeriodsByUserID",
			Handler:    _FinancialPeriodService_GetFinancialPeriodsByUserID_Handler,
		},
		{
			MethodName: "AddUserToFinancialPeriod",
			Handler:    _FinancialPeriodService_AddUserToFinancialPeriod_Handler,
		},
		{
			MethodName: "RemoveUserFromFinancialPeriod",
			Handler:    _FinancialPeriodService_RemoveUserFromFinancialPeriod_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/financial/v1/financial.proto",
}
