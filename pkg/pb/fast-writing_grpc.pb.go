// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.3
// source: fast-writing.proto

package pb

import (
	context "context"
	models "fast-writing/pkg/pb/models"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServiceClient interface {
	GetUser(ctx context.Context, in *models.UserId, opts ...grpc.CallOption) (*models.User, error)
	CreateUser(ctx context.Context, in *models.User, opts ...grpc.CallOption) (*models.UserId, error)
	UpdateUser(ctx context.Context, in *models.User, opts ...grpc.CallOption) (*models.UserId, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) GetUser(ctx context.Context, in *models.UserId, opts ...grpc.CallOption) (*models.User, error) {
	out := new(models.User)
	err := c.cc.Invoke(ctx, "/fastwriting.UserService/GetUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) CreateUser(ctx context.Context, in *models.User, opts ...grpc.CallOption) (*models.UserId, error) {
	out := new(models.UserId)
	err := c.cc.Invoke(ctx, "/fastwriting.UserService/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UpdateUser(ctx context.Context, in *models.User, opts ...grpc.CallOption) (*models.UserId, error) {
	out := new(models.UserId)
	err := c.cc.Invoke(ctx, "/fastwriting.UserService/UpdateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
// All implementations should embed UnimplementedUserServiceServer
// for forward compatibility
type UserServiceServer interface {
	GetUser(context.Context, *models.UserId) (*models.User, error)
	CreateUser(context.Context, *models.User) (*models.UserId, error)
	UpdateUser(context.Context, *models.User) (*models.UserId, error)
}

// UnimplementedUserServiceServer should be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (UnimplementedUserServiceServer) GetUser(context.Context, *models.UserId) (*models.User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (UnimplementedUserServiceServer) CreateUser(context.Context, *models.User) (*models.UserId, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedUserServiceServer) UpdateUser(context.Context, *models.User) (*models.UserId, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}

// UnsafeUserServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServiceServer will
// result in compilation errors.
type UnsafeUserServiceServer interface {
	mustEmbedUnimplementedUserServiceServer()
}

func RegisterUserServiceServer(s grpc.ServiceRegistrar, srv UserServiceServer) {
	s.RegisterService(&UserService_ServiceDesc, srv)
}

func _UserService_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(models.UserId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fastwriting.UserService/GetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUser(ctx, req.(*models.UserId))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(models.User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fastwriting.UserService/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).CreateUser(ctx, req.(*models.User))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(models.User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fastwriting.UserService/UpdateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UpdateUser(ctx, req.(*models.User))
	}
	return interceptor(ctx, in, info, handler)
}

// UserService_ServiceDesc is the grpc.ServiceDesc for UserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "fastwriting.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUser",
			Handler:    _UserService_GetUser_Handler,
		},
		{
			MethodName: "CreateUser",
			Handler:    _UserService_CreateUser_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _UserService_UpdateUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "fast-writing.proto",
}

// WritingServiceClient is the client API for WritingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WritingServiceClient interface {
	GetContents(ctx context.Context, in *models.ContentsId, opts ...grpc.CallOption) (*models.Contents, error)
	GetContentsList(ctx context.Context, in *models.ContentsQueryParams, opts ...grpc.CallOption) (*models.ContentsList, error)
	GetContentsListByTags(ctx context.Context, in *models.TagQueryParams, opts ...grpc.CallOption) (*models.ContentsList, error)
	GetUserContentsList(ctx context.Context, in *models.UserContentsQueryParams, opts ...grpc.CallOption) (*models.ContentsList, error)
	CreateUserContents(ctx context.Context, in *CreateContentsRequest, opts ...grpc.CallOption) (*CreateContentsResponse, error)
	DeleteUserContents(ctx context.Context, in *DeleteContentsRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
}

type writingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewWritingServiceClient(cc grpc.ClientConnInterface) WritingServiceClient {
	return &writingServiceClient{cc}
}

func (c *writingServiceClient) GetContents(ctx context.Context, in *models.ContentsId, opts ...grpc.CallOption) (*models.Contents, error) {
	out := new(models.Contents)
	err := c.cc.Invoke(ctx, "/fastwriting.WritingService/GetContents", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *writingServiceClient) GetContentsList(ctx context.Context, in *models.ContentsQueryParams, opts ...grpc.CallOption) (*models.ContentsList, error) {
	out := new(models.ContentsList)
	err := c.cc.Invoke(ctx, "/fastwriting.WritingService/GetContentsList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *writingServiceClient) GetContentsListByTags(ctx context.Context, in *models.TagQueryParams, opts ...grpc.CallOption) (*models.ContentsList, error) {
	out := new(models.ContentsList)
	err := c.cc.Invoke(ctx, "/fastwriting.WritingService/GetContentsListByTags", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *writingServiceClient) GetUserContentsList(ctx context.Context, in *models.UserContentsQueryParams, opts ...grpc.CallOption) (*models.ContentsList, error) {
	out := new(models.ContentsList)
	err := c.cc.Invoke(ctx, "/fastwriting.WritingService/GetUserContentsList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *writingServiceClient) CreateUserContents(ctx context.Context, in *CreateContentsRequest, opts ...grpc.CallOption) (*CreateContentsResponse, error) {
	out := new(CreateContentsResponse)
	err := c.cc.Invoke(ctx, "/fastwriting.WritingService/CreateUserContents", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *writingServiceClient) DeleteUserContents(ctx context.Context, in *DeleteContentsRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, "/fastwriting.WritingService/DeleteUserContents", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WritingServiceServer is the server API for WritingService service.
// All implementations should embed UnimplementedWritingServiceServer
// for forward compatibility
type WritingServiceServer interface {
	GetContents(context.Context, *models.ContentsId) (*models.Contents, error)
	GetContentsList(context.Context, *models.ContentsQueryParams) (*models.ContentsList, error)
	GetContentsListByTags(context.Context, *models.TagQueryParams) (*models.ContentsList, error)
	GetUserContentsList(context.Context, *models.UserContentsQueryParams) (*models.ContentsList, error)
	CreateUserContents(context.Context, *CreateContentsRequest) (*CreateContentsResponse, error)
	DeleteUserContents(context.Context, *DeleteContentsRequest) (*DeleteResponse, error)
}

// UnimplementedWritingServiceServer should be embedded to have forward compatible implementations.
type UnimplementedWritingServiceServer struct {
}

func (UnimplementedWritingServiceServer) GetContents(context.Context, *models.ContentsId) (*models.Contents, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetContents not implemented")
}
func (UnimplementedWritingServiceServer) GetContentsList(context.Context, *models.ContentsQueryParams) (*models.ContentsList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetContentsList not implemented")
}
func (UnimplementedWritingServiceServer) GetContentsListByTags(context.Context, *models.TagQueryParams) (*models.ContentsList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetContentsListByTags not implemented")
}
func (UnimplementedWritingServiceServer) GetUserContentsList(context.Context, *models.UserContentsQueryParams) (*models.ContentsList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserContentsList not implemented")
}
func (UnimplementedWritingServiceServer) CreateUserContents(context.Context, *CreateContentsRequest) (*CreateContentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUserContents not implemented")
}
func (UnimplementedWritingServiceServer) DeleteUserContents(context.Context, *DeleteContentsRequest) (*DeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUserContents not implemented")
}

// UnsafeWritingServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WritingServiceServer will
// result in compilation errors.
type UnsafeWritingServiceServer interface {
	mustEmbedUnimplementedWritingServiceServer()
}

func RegisterWritingServiceServer(s grpc.ServiceRegistrar, srv WritingServiceServer) {
	s.RegisterService(&WritingService_ServiceDesc, srv)
}

func _WritingService_GetContents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(models.ContentsId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WritingServiceServer).GetContents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fastwriting.WritingService/GetContents",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WritingServiceServer).GetContents(ctx, req.(*models.ContentsId))
	}
	return interceptor(ctx, in, info, handler)
}

func _WritingService_GetContentsList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(models.ContentsQueryParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WritingServiceServer).GetContentsList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fastwriting.WritingService/GetContentsList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WritingServiceServer).GetContentsList(ctx, req.(*models.ContentsQueryParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _WritingService_GetContentsListByTags_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(models.TagQueryParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WritingServiceServer).GetContentsListByTags(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fastwriting.WritingService/GetContentsListByTags",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WritingServiceServer).GetContentsListByTags(ctx, req.(*models.TagQueryParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _WritingService_GetUserContentsList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(models.UserContentsQueryParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WritingServiceServer).GetUserContentsList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fastwriting.WritingService/GetUserContentsList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WritingServiceServer).GetUserContentsList(ctx, req.(*models.UserContentsQueryParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _WritingService_CreateUserContents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateContentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WritingServiceServer).CreateUserContents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fastwriting.WritingService/CreateUserContents",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WritingServiceServer).CreateUserContents(ctx, req.(*CreateContentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WritingService_DeleteUserContents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteContentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WritingServiceServer).DeleteUserContents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fastwriting.WritingService/DeleteUserContents",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WritingServiceServer).DeleteUserContents(ctx, req.(*DeleteContentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// WritingService_ServiceDesc is the grpc.ServiceDesc for WritingService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WritingService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "fastwriting.WritingService",
	HandlerType: (*WritingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetContents",
			Handler:    _WritingService_GetContents_Handler,
		},
		{
			MethodName: "GetContentsList",
			Handler:    _WritingService_GetContentsList_Handler,
		},
		{
			MethodName: "GetContentsListByTags",
			Handler:    _WritingService_GetContentsListByTags_Handler,
		},
		{
			MethodName: "GetUserContentsList",
			Handler:    _WritingService_GetUserContentsList_Handler,
		},
		{
			MethodName: "CreateUserContents",
			Handler:    _WritingService_CreateUserContents_Handler,
		},
		{
			MethodName: "DeleteUserContents",
			Handler:    _WritingService_DeleteUserContents_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "fast-writing.proto",
}
