// MIT License

// Copyright (c) 2024 EASE Lab

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.0
// source: proto/compression/compression.proto

package compression

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
	FileCompress_CompressFile_FullMethodName = "/compression.FileCompress/CompressFile"
)

// FileCompressClient is the client API for FileCompress service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FileCompressClient interface {
	CompressFile(ctx context.Context, in *SendFile, opts ...grpc.CallOption) (*GetCompressedFile, error)
}

type fileCompressClient struct {
	cc grpc.ClientConnInterface
}

func NewFileCompressClient(cc grpc.ClientConnInterface) FileCompressClient {
	return &fileCompressClient{cc}
}

func (c *fileCompressClient) CompressFile(ctx context.Context, in *SendFile, opts ...grpc.CallOption) (*GetCompressedFile, error) {
	out := new(GetCompressedFile)
	err := c.cc.Invoke(ctx, FileCompress_CompressFile_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FileCompressServer is the server API for FileCompress service.
// All implementations must embed UnimplementedFileCompressServer
// for forward compatibility
type FileCompressServer interface {
	CompressFile(context.Context, *SendFile) (*GetCompressedFile, error)
	mustEmbedUnimplementedFileCompressServer()
}

// UnimplementedFileCompressServer must be embedded to have forward compatible implementations.
type UnimplementedFileCompressServer struct {
}

func (UnimplementedFileCompressServer) CompressFile(context.Context, *SendFile) (*GetCompressedFile, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CompressFile not implemented")
}
func (UnimplementedFileCompressServer) mustEmbedUnimplementedFileCompressServer() {}

// UnsafeFileCompressServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FileCompressServer will
// result in compilation errors.
type UnsafeFileCompressServer interface {
	mustEmbedUnimplementedFileCompressServer()
}

func RegisterFileCompressServer(s grpc.ServiceRegistrar, srv FileCompressServer) {
	s.RegisterService(&FileCompress_ServiceDesc, srv)
}

func _FileCompress_CompressFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendFile)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileCompressServer).CompressFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FileCompress_CompressFile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileCompressServer).CompressFile(ctx, req.(*SendFile))
	}
	return interceptor(ctx, in, info, handler)
}

// FileCompress_ServiceDesc is the grpc.ServiceDesc for FileCompress service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FileCompress_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "compression.FileCompress",
	HandlerType: (*FileCompressServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CompressFile",
			Handler:    _FileCompress_CompressFile_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/compression/compression.proto",
}
