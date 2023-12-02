// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.1
// source: docs.proto

package docs

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// DocumentServiceClient is the client API for DocumentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DocumentServiceClient interface {
	GetAllByLimitAndOffset(ctx context.Context, in *GetAllRequest, opts ...grpc.CallOption) (*GetAllResponse, error)
	Save(ctx context.Context, in *SaveRequest, opts ...grpc.CallOption) (*SaveResponse, error)
	Validate(ctx context.Context, in *ValidateRequest, opts ...grpc.CallOption) (*ValidateResponse, error)
}

type documentServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDocumentServiceClient(cc grpc.ClientConnInterface) DocumentServiceClient {
	return &documentServiceClient{cc}
}

func (c *documentServiceClient) GetAllByLimitAndOffset(ctx context.Context, in *GetAllRequest, opts ...grpc.CallOption) (*GetAllResponse, error) {
	out := new(GetAllResponse)
	err := c.cc.Invoke(ctx, "/DocumentService/GetAllByLimitAndOffset", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *documentServiceClient) Save(ctx context.Context, in *SaveRequest, opts ...grpc.CallOption) (*SaveResponse, error) {
	out := new(SaveResponse)
	err := c.cc.Invoke(ctx, "/DocumentService/Save", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *documentServiceClient) Validate(ctx context.Context, in *ValidateRequest, opts ...grpc.CallOption) (*ValidateResponse, error) {
	out := new(ValidateResponse)
	err := c.cc.Invoke(ctx, "/DocumentService/Validate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DocumentServiceServer is the server API for DocumentService service.
// All implementations must embed UnimplementedDocumentServiceServer
// for forward compatibility
type DocumentServiceServer interface {
	GetAllByLimitAndOffset(context.Context, *GetAllRequest) (*GetAllResponse, error)
	Save(context.Context, *SaveRequest) (*SaveResponse, error)
	Validate(context.Context, *ValidateRequest) (*ValidateResponse, error)
	mustEmbedUnimplementedDocumentServiceServer()
}

// UnimplementedDocumentServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDocumentServiceServer struct {
}

func (UnimplementedDocumentServiceServer) GetAllByLimitAndOffset(context.Context, *GetAllRequest) (*GetAllResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllByLimitAndOffset not implemented")
}
func (UnimplementedDocumentServiceServer) Save(context.Context, *SaveRequest) (*SaveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Save not implemented")
}
func (UnimplementedDocumentServiceServer) Validate(context.Context, *ValidateRequest) (*ValidateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Validate not implemented")
}
func (UnimplementedDocumentServiceServer) mustEmbedUnimplementedDocumentServiceServer() {}

// UnsafeDocumentServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DocumentServiceServer will
// result in compilation errors.
type UnsafeDocumentServiceServer interface {
	mustEmbedUnimplementedDocumentServiceServer()
}

func RegisterDocumentServiceServer(s grpc.ServiceRegistrar, srv DocumentServiceServer) {
	s.RegisterService(&DocumentService_ServiceDesc, srv)
}

func _DocumentService_GetAllByLimitAndOffset_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DocumentServiceServer).GetAllByLimitAndOffset(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/DocumentService/GetAllByLimitAndOffset",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DocumentServiceServer).GetAllByLimitAndOffset(ctx, req.(*GetAllRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DocumentService_Save_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SaveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DocumentServiceServer).Save(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/DocumentService/Save",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DocumentServiceServer).Save(ctx, req.(*SaveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DocumentService_Validate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DocumentServiceServer).Validate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/DocumentService/Validate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DocumentServiceServer).Validate(ctx, req.(*ValidateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DocumentService_ServiceDesc is the grpc.ServiceDesc for DocumentService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DocumentService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "DocumentService",
	HandlerType: (*DocumentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAllByLimitAndOffset",
			Handler:    _DocumentService_GetAllByLimitAndOffset_Handler,
		},
		{
			MethodName: "Save",
			Handler:    _DocumentService_Save_Handler,
		},
		{
			MethodName: "Validate",
			Handler:    _DocumentService_Validate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "docs.proto",
}
