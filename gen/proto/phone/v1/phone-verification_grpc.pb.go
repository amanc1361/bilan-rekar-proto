// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: proto/phone/v1/phone-verification.proto

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
	PhoneVerificationService_SendOTP_FullMethodName           = "/auth.v1.PhoneVerificationService/SendOTP"
	PhoneVerificationService_VerifyOTP_FullMethodName         = "/auth.v1.PhoneVerificationService/VerifyOTP"
	PhoneVerificationService_RegisterWithPhone_FullMethodName = "/auth.v1.PhoneVerificationService/RegisterWithPhone"
)

// PhoneVerificationServiceClient is the client API for PhoneVerificationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// سرویس مربوط به تایید شماره تلفن و OTP
type PhoneVerificationServiceClient interface {
	// ارسال کد OTP به شماره تلفن
	SendOTP(ctx context.Context, in *SendOTPRequest, opts ...grpc.CallOption) (*SendOTPResponse, error)
	// تایید کد OTP
	VerifyOTP(ctx context.Context, in *VerifyOTPRequest, opts ...grpc.CallOption) (*VerifyOTPResponse, error)
	// ثبت‌نام کاربر با شماره موبایل تایید شده
	RegisterWithPhone(ctx context.Context, in *RegisterWithPhoneRequest, opts ...grpc.CallOption) (*RegisterWithPhoneResponse, error)
}

type phoneVerificationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPhoneVerificationServiceClient(cc grpc.ClientConnInterface) PhoneVerificationServiceClient {
	return &phoneVerificationServiceClient{cc}
}

func (c *phoneVerificationServiceClient) SendOTP(ctx context.Context, in *SendOTPRequest, opts ...grpc.CallOption) (*SendOTPResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SendOTPResponse)
	err := c.cc.Invoke(ctx, PhoneVerificationService_SendOTP_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *phoneVerificationServiceClient) VerifyOTP(ctx context.Context, in *VerifyOTPRequest, opts ...grpc.CallOption) (*VerifyOTPResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(VerifyOTPResponse)
	err := c.cc.Invoke(ctx, PhoneVerificationService_VerifyOTP_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *phoneVerificationServiceClient) RegisterWithPhone(ctx context.Context, in *RegisterWithPhoneRequest, opts ...grpc.CallOption) (*RegisterWithPhoneResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RegisterWithPhoneResponse)
	err := c.cc.Invoke(ctx, PhoneVerificationService_RegisterWithPhone_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PhoneVerificationServiceServer is the server API for PhoneVerificationService service.
// All implementations must embed UnimplementedPhoneVerificationServiceServer
// for forward compatibility.
//
// سرویس مربوط به تایید شماره تلفن و OTP
type PhoneVerificationServiceServer interface {
	// ارسال کد OTP به شماره تلفن
	SendOTP(context.Context, *SendOTPRequest) (*SendOTPResponse, error)
	// تایید کد OTP
	VerifyOTP(context.Context, *VerifyOTPRequest) (*VerifyOTPResponse, error)
	// ثبت‌نام کاربر با شماره موبایل تایید شده
	RegisterWithPhone(context.Context, *RegisterWithPhoneRequest) (*RegisterWithPhoneResponse, error)
	mustEmbedUnimplementedPhoneVerificationServiceServer()
}

// UnimplementedPhoneVerificationServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedPhoneVerificationServiceServer struct{}

func (UnimplementedPhoneVerificationServiceServer) SendOTP(context.Context, *SendOTPRequest) (*SendOTPResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendOTP not implemented")
}
func (UnimplementedPhoneVerificationServiceServer) VerifyOTP(context.Context, *VerifyOTPRequest) (*VerifyOTPResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyOTP not implemented")
}
func (UnimplementedPhoneVerificationServiceServer) RegisterWithPhone(context.Context, *RegisterWithPhoneRequest) (*RegisterWithPhoneResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterWithPhone not implemented")
}
func (UnimplementedPhoneVerificationServiceServer) mustEmbedUnimplementedPhoneVerificationServiceServer() {
}
func (UnimplementedPhoneVerificationServiceServer) testEmbeddedByValue() {}

// UnsafePhoneVerificationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PhoneVerificationServiceServer will
// result in compilation errors.
type UnsafePhoneVerificationServiceServer interface {
	mustEmbedUnimplementedPhoneVerificationServiceServer()
}

func RegisterPhoneVerificationServiceServer(s grpc.ServiceRegistrar, srv PhoneVerificationServiceServer) {
	// If the following call pancis, it indicates UnimplementedPhoneVerificationServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&PhoneVerificationService_ServiceDesc, srv)
}

func _PhoneVerificationService_SendOTP_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendOTPRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PhoneVerificationServiceServer).SendOTP(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PhoneVerificationService_SendOTP_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PhoneVerificationServiceServer).SendOTP(ctx, req.(*SendOTPRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PhoneVerificationService_VerifyOTP_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyOTPRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PhoneVerificationServiceServer).VerifyOTP(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PhoneVerificationService_VerifyOTP_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PhoneVerificationServiceServer).VerifyOTP(ctx, req.(*VerifyOTPRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PhoneVerificationService_RegisterWithPhone_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterWithPhoneRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PhoneVerificationServiceServer).RegisterWithPhone(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PhoneVerificationService_RegisterWithPhone_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PhoneVerificationServiceServer).RegisterWithPhone(ctx, req.(*RegisterWithPhoneRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PhoneVerificationService_ServiceDesc is the grpc.ServiceDesc for PhoneVerificationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PhoneVerificationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "auth.v1.PhoneVerificationService",
	HandlerType: (*PhoneVerificationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendOTP",
			Handler:    _PhoneVerificationService_SendOTP_Handler,
		},
		{
			MethodName: "VerifyOTP",
			Handler:    _PhoneVerificationService_VerifyOTP_Handler,
		},
		{
			MethodName: "RegisterWithPhone",
			Handler:    _PhoneVerificationService_RegisterWithPhone_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/phone/v1/phone-verification.proto",
}
