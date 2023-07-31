// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.4
// source: internal/handlers/article/grpc/article.proto

package grpc

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
	Article_GetAllArticles_FullMethodName = "/article_grpc.Article/GetAllArticles"
	Article_GetArticleById_FullMethodName = "/article_grpc.Article/GetArticleById"
	Article_CreateArticle_FullMethodName  = "/article_grpc.Article/CreateArticle"
	Article_UpdateArticle_FullMethodName  = "/article_grpc.Article/UpdateArticle"
	Article_DeleteArticle_FullMethodName  = "/article_grpc.Article/DeleteArticle"
)

// ArticleClient is the client API for Article service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ArticleClient interface {
	GetAllArticles(ctx context.Context, in *GetAllArticlesRequest, opts ...grpc.CallOption) (*GetAllArticlesResponse, error)
	GetArticleById(ctx context.Context, in *GetArticleByIdRequest, opts ...grpc.CallOption) (*GetArticleByIdResponse, error)
	CreateArticle(ctx context.Context, in *CreateArticleRequest, opts ...grpc.CallOption) (*CreateArticleResponse, error)
	UpdateArticle(ctx context.Context, in *UpdateArticleRequest, opts ...grpc.CallOption) (*UpdateArticleResponse, error)
	DeleteArticle(ctx context.Context, in *DeleteArticleRequest, opts ...grpc.CallOption) (*DeleteArticleResponse, error)
}

type articleClient struct {
	cc grpc.ClientConnInterface
}

func NewArticleClient(cc grpc.ClientConnInterface) ArticleClient {
	return &articleClient{cc}
}

func (c *articleClient) GetAllArticles(ctx context.Context, in *GetAllArticlesRequest, opts ...grpc.CallOption) (*GetAllArticlesResponse, error) {
	out := new(GetAllArticlesResponse)
	err := c.cc.Invoke(ctx, Article_GetAllArticles_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *articleClient) GetArticleById(ctx context.Context, in *GetArticleByIdRequest, opts ...grpc.CallOption) (*GetArticleByIdResponse, error) {
	out := new(GetArticleByIdResponse)
	err := c.cc.Invoke(ctx, Article_GetArticleById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *articleClient) CreateArticle(ctx context.Context, in *CreateArticleRequest, opts ...grpc.CallOption) (*CreateArticleResponse, error) {
	out := new(CreateArticleResponse)
	err := c.cc.Invoke(ctx, Article_CreateArticle_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *articleClient) UpdateArticle(ctx context.Context, in *UpdateArticleRequest, opts ...grpc.CallOption) (*UpdateArticleResponse, error) {
	out := new(UpdateArticleResponse)
	err := c.cc.Invoke(ctx, Article_UpdateArticle_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *articleClient) DeleteArticle(ctx context.Context, in *DeleteArticleRequest, opts ...grpc.CallOption) (*DeleteArticleResponse, error) {
	out := new(DeleteArticleResponse)
	err := c.cc.Invoke(ctx, Article_DeleteArticle_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ArticleServer is the server API for Article service.
// All implementations must embed UnimplementedArticleServer
// for forward compatibility
type ArticleServer interface {
	GetAllArticles(context.Context, *GetAllArticlesRequest) (*GetAllArticlesResponse, error)
	GetArticleById(context.Context, *GetArticleByIdRequest) (*GetArticleByIdResponse, error)
	CreateArticle(context.Context, *CreateArticleRequest) (*CreateArticleResponse, error)
	UpdateArticle(context.Context, *UpdateArticleRequest) (*UpdateArticleResponse, error)
	DeleteArticle(context.Context, *DeleteArticleRequest) (*DeleteArticleResponse, error)
	mustEmbedUnimplementedArticleServer()
}

// UnimplementedArticleServer must be embedded to have forward compatible implementations.
type UnimplementedArticleServer struct {
}

func (UnimplementedArticleServer) GetAllArticles(context.Context, *GetAllArticlesRequest) (*GetAllArticlesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllArticles not implemented")
}
func (UnimplementedArticleServer) GetArticleById(context.Context, *GetArticleByIdRequest) (*GetArticleByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetArticleById not implemented")
}
func (UnimplementedArticleServer) CreateArticle(context.Context, *CreateArticleRequest) (*CreateArticleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateArticle not implemented")
}
func (UnimplementedArticleServer) UpdateArticle(context.Context, *UpdateArticleRequest) (*UpdateArticleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateArticle not implemented")
}
func (UnimplementedArticleServer) DeleteArticle(context.Context, *DeleteArticleRequest) (*DeleteArticleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteArticle not implemented")
}
func (UnimplementedArticleServer) mustEmbedUnimplementedArticleServer() {}

// UnsafeArticleServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ArticleServer will
// result in compilation errors.
type UnsafeArticleServer interface {
	mustEmbedUnimplementedArticleServer()
}

func RegisterArticleServer(s grpc.ServiceRegistrar, srv ArticleServer) {
	s.RegisterService(&Article_ServiceDesc, srv)
}

func _Article_GetAllArticles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllArticlesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticleServer).GetAllArticles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Article_GetAllArticles_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticleServer).GetAllArticles(ctx, req.(*GetAllArticlesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Article_GetArticleById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetArticleByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticleServer).GetArticleById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Article_GetArticleById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticleServer).GetArticleById(ctx, req.(*GetArticleByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Article_CreateArticle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateArticleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticleServer).CreateArticle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Article_CreateArticle_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticleServer).CreateArticle(ctx, req.(*CreateArticleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Article_UpdateArticle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateArticleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticleServer).UpdateArticle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Article_UpdateArticle_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticleServer).UpdateArticle(ctx, req.(*UpdateArticleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Article_DeleteArticle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteArticleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticleServer).DeleteArticle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Article_DeleteArticle_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticleServer).DeleteArticle(ctx, req.(*DeleteArticleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Article_ServiceDesc is the grpc.ServiceDesc for Article service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Article_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "article_grpc.Article",
	HandlerType: (*ArticleServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAllArticles",
			Handler:    _Article_GetAllArticles_Handler,
		},
		{
			MethodName: "GetArticleById",
			Handler:    _Article_GetArticleById_Handler,
		},
		{
			MethodName: "CreateArticle",
			Handler:    _Article_CreateArticle_Handler,
		},
		{
			MethodName: "UpdateArticle",
			Handler:    _Article_UpdateArticle_Handler,
		},
		{
			MethodName: "DeleteArticle",
			Handler:    _Article_DeleteArticle_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "internal/handlers/article/grpc/article.proto",
}