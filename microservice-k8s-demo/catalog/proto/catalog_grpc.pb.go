// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.3
// source: proto/catalog.proto

package proto

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

const (
	CatalogService_GetCatalogItem_FullMethodName         = "/catalog.CatalogService/GetCatalogItem"
	CatalogService_ListCatalogItems_FullMethodName       = "/catalog.CatalogService/ListCatalogItems"
	CatalogService_ListCatalogItemsByName_FullMethodName = "/catalog.CatalogService/ListCatalogItemsByName"
	CatalogService_CreateCatalogItem_FullMethodName      = "/catalog.CatalogService/CreateCatalogItem"
	CatalogService_UpdateCatalogItem_FullMethodName      = "/catalog.CatalogService/UpdateCatalogItem"
	CatalogService_DeleteCatalogItem_FullMethodName      = "/catalog.CatalogService/DeleteCatalogItem"
)

// CatalogServiceClient is the client API for CatalogService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CatalogServiceClient interface {
	GetCatalogItem(ctx context.Context, in *GetCatalogItemRequest, opts ...grpc.CallOption) (*GetCatalogItemResponse, error)
	ListCatalogItems(ctx context.Context, in *ListCatalogItemsRequest, opts ...grpc.CallOption) (*ListCatalogItemsResponse, error)
	ListCatalogItemsByName(ctx context.Context, in *ListCatalogItemsByNameRequest, opts ...grpc.CallOption) (*ListCatalogItemsByNameResponse, error)
	CreateCatalogItem(ctx context.Context, in *CreateCatalogItemRequest, opts ...grpc.CallOption) (*CreateCatalogItemResponse, error)
	UpdateCatalogItem(ctx context.Context, in *UpdateCatalogItemRequest, opts ...grpc.CallOption) (*UpdateCatalogItemResponse, error)
	DeleteCatalogItem(ctx context.Context, in *DeleteCatalogItemRequest, opts ...grpc.CallOption) (*DeleteCatalogItemResponse, error)
}

type catalogServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCatalogServiceClient(cc grpc.ClientConnInterface) CatalogServiceClient {
	return &catalogServiceClient{cc}
}

func (c *catalogServiceClient) GetCatalogItem(ctx context.Context, in *GetCatalogItemRequest, opts ...grpc.CallOption) (*GetCatalogItemResponse, error) {
	out := new(GetCatalogItemResponse)
	err := c.cc.Invoke(ctx, CatalogService_GetCatalogItem_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogServiceClient) ListCatalogItems(ctx context.Context, in *ListCatalogItemsRequest, opts ...grpc.CallOption) (*ListCatalogItemsResponse, error) {
	out := new(ListCatalogItemsResponse)
	err := c.cc.Invoke(ctx, CatalogService_ListCatalogItems_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogServiceClient) ListCatalogItemsByName(ctx context.Context, in *ListCatalogItemsByNameRequest, opts ...grpc.CallOption) (*ListCatalogItemsByNameResponse, error) {
	out := new(ListCatalogItemsByNameResponse)
	err := c.cc.Invoke(ctx, CatalogService_ListCatalogItemsByName_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogServiceClient) CreateCatalogItem(ctx context.Context, in *CreateCatalogItemRequest, opts ...grpc.CallOption) (*CreateCatalogItemResponse, error) {
	out := new(CreateCatalogItemResponse)
	err := c.cc.Invoke(ctx, CatalogService_CreateCatalogItem_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogServiceClient) UpdateCatalogItem(ctx context.Context, in *UpdateCatalogItemRequest, opts ...grpc.CallOption) (*UpdateCatalogItemResponse, error) {
	out := new(UpdateCatalogItemResponse)
	err := c.cc.Invoke(ctx, CatalogService_UpdateCatalogItem_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogServiceClient) DeleteCatalogItem(ctx context.Context, in *DeleteCatalogItemRequest, opts ...grpc.CallOption) (*DeleteCatalogItemResponse, error) {
	out := new(DeleteCatalogItemResponse)
	err := c.cc.Invoke(ctx, CatalogService_DeleteCatalogItem_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CatalogServiceServer is the server API for CatalogService service.
// All implementations must embed UnimplementedCatalogServiceServer
// for forward compatibility
type CatalogServiceServer interface {
	GetCatalogItem(context.Context, *GetCatalogItemRequest) (*GetCatalogItemResponse, error)
	ListCatalogItems(context.Context, *ListCatalogItemsRequest) (*ListCatalogItemsResponse, error)
	ListCatalogItemsByName(context.Context, *ListCatalogItemsByNameRequest) (*ListCatalogItemsByNameResponse, error)
	CreateCatalogItem(context.Context, *CreateCatalogItemRequest) (*CreateCatalogItemResponse, error)
	UpdateCatalogItem(context.Context, *UpdateCatalogItemRequest) (*UpdateCatalogItemResponse, error)
	DeleteCatalogItem(context.Context, *DeleteCatalogItemRequest) (*DeleteCatalogItemResponse, error)
	mustEmbedUnimplementedCatalogServiceServer()
}

// UnimplementedCatalogServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCatalogServiceServer struct{}

func (UnimplementedCatalogServiceServer) GetCatalogItem(context.Context, *GetCatalogItemRequest) (*GetCatalogItemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCatalogItem not implemented")
}

func (UnimplementedCatalogServiceServer) ListCatalogItems(context.Context, *ListCatalogItemsRequest) (*ListCatalogItemsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCatalogItems not implemented")
}

func (UnimplementedCatalogServiceServer) ListCatalogItemsByName(context.Context, *ListCatalogItemsByNameRequest) (*ListCatalogItemsByNameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCatalogItemsByName not implemented")
}

func (UnimplementedCatalogServiceServer) CreateCatalogItem(context.Context, *CreateCatalogItemRequest) (*CreateCatalogItemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCatalogItem not implemented")
}

func (UnimplementedCatalogServiceServer) UpdateCatalogItem(context.Context, *UpdateCatalogItemRequest) (*UpdateCatalogItemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCatalogItem not implemented")
}

func (UnimplementedCatalogServiceServer) DeleteCatalogItem(context.Context, *DeleteCatalogItemRequest) (*DeleteCatalogItemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCatalogItem not implemented")
}
func (UnimplementedCatalogServiceServer) mustEmbedUnimplementedCatalogServiceServer() {}

// UnsafeCatalogServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CatalogServiceServer will
// result in compilation errors.
type UnsafeCatalogServiceServer interface {
	mustEmbedUnimplementedCatalogServiceServer()
}

func RegisterCatalogServiceServer(s grpc.ServiceRegistrar, srv CatalogServiceServer) {
	s.RegisterService(&CatalogService_ServiceDesc, srv)
}

func _CatalogService_GetCatalogItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCatalogItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServiceServer).GetCatalogItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CatalogService_GetCatalogItem_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServiceServer).GetCatalogItem(ctx, req.(*GetCatalogItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CatalogService_ListCatalogItems_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListCatalogItemsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServiceServer).ListCatalogItems(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CatalogService_ListCatalogItems_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServiceServer).ListCatalogItems(ctx, req.(*ListCatalogItemsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CatalogService_ListCatalogItemsByName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListCatalogItemsByNameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServiceServer).ListCatalogItemsByName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CatalogService_ListCatalogItemsByName_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServiceServer).ListCatalogItemsByName(ctx, req.(*ListCatalogItemsByNameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CatalogService_CreateCatalogItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCatalogItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServiceServer).CreateCatalogItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CatalogService_CreateCatalogItem_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServiceServer).CreateCatalogItem(ctx, req.(*CreateCatalogItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CatalogService_UpdateCatalogItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCatalogItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServiceServer).UpdateCatalogItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CatalogService_UpdateCatalogItem_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServiceServer).UpdateCatalogItem(ctx, req.(*UpdateCatalogItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CatalogService_DeleteCatalogItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCatalogItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServiceServer).DeleteCatalogItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CatalogService_DeleteCatalogItem_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServiceServer).DeleteCatalogItem(ctx, req.(*DeleteCatalogItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CatalogService_ServiceDesc is the grpc.ServiceDesc for CatalogService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CatalogService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "catalog.CatalogService",
	HandlerType: (*CatalogServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCatalogItem",
			Handler:    _CatalogService_GetCatalogItem_Handler,
		},
		{
			MethodName: "ListCatalogItems",
			Handler:    _CatalogService_ListCatalogItems_Handler,
		},
		{
			MethodName: "ListCatalogItemsByName",
			Handler:    _CatalogService_ListCatalogItemsByName_Handler,
		},
		{
			MethodName: "CreateCatalogItem",
			Handler:    _CatalogService_CreateCatalogItem_Handler,
		},
		{
			MethodName: "UpdateCatalogItem",
			Handler:    _CatalogService_UpdateCatalogItem_Handler,
		},
		{
			MethodName: "DeleteCatalogItem",
			Handler:    _CatalogService_DeleteCatalogItem_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/catalog.proto",
}
