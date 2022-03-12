// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: search.proto

package pb

import (
	context "context"
	models "github.com/tanan/fast-writing/pkg/pb/models"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// SearchServiceClient is the client API for SearchService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SearchServiceClient interface {
	FindContentsIdListByTitle(ctx context.Context, in *models.TitleQueryParams, opts ...grpc.CallOption) (*ContentsScoreList, error)
}

type searchServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSearchServiceClient(cc grpc.ClientConnInterface) SearchServiceClient {
	return &searchServiceClient{cc}
}

func (c *searchServiceClient) FindContentsIdListByTitle(ctx context.Context, in *models.TitleQueryParams, opts ...grpc.CallOption) (*ContentsScoreList, error) {
	out := new(ContentsScoreList)
	err := c.cc.Invoke(ctx, "/fastwriting.SearchService/FindContentsIdListByTitle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SearchServiceServer is the server API for SearchService service.
// All implementations should embed UnimplementedSearchServiceServer
// for forward compatibility
type SearchServiceServer interface {
	FindContentsIdListByTitle(context.Context, *models.TitleQueryParams) (*ContentsScoreList, error)
}

// UnimplementedSearchServiceServer should be embedded to have forward compatible implementations.
type UnimplementedSearchServiceServer struct {
}

func (UnimplementedSearchServiceServer) FindContentsIdListByTitle(context.Context, *models.TitleQueryParams) (*ContentsScoreList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindContentsIdListByTitle not implemented")
}

// UnsafeSearchServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SearchServiceServer will
// result in compilation errors.
type UnsafeSearchServiceServer interface {
	mustEmbedUnimplementedSearchServiceServer()
}

func RegisterSearchServiceServer(s grpc.ServiceRegistrar, srv SearchServiceServer) {
	s.RegisterService(&SearchService_ServiceDesc, srv)
}

func _SearchService_FindContentsIdListByTitle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(models.TitleQueryParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearchServiceServer).FindContentsIdListByTitle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fastwriting.SearchService/FindContentsIdListByTitle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearchServiceServer).FindContentsIdListByTitle(ctx, req.(*models.TitleQueryParams))
	}
	return interceptor(ctx, in, info, handler)
}

// SearchService_ServiceDesc is the grpc.ServiceDesc for SearchService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SearchService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "fastwriting.SearchService",
	HandlerType: (*SearchServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FindContentsIdListByTitle",
			Handler:    _SearchService_FindContentsIdListByTitle_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "search.proto",
}
