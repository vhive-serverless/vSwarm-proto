// MIT License
//
// Copyright (c) 2022 EASE lab
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
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
// source: proto/aes/aes.proto

package aes

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
	Aes_ShowEncryption_FullMethodName = "/aes.Aes/ShowEncryption"
)

// AesClient is the client API for Aes service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AesClient interface {
	// Performs encryption on received message and returns information about it
	ShowEncryption(ctx context.Context, in *PlainTextMessage, opts ...grpc.CallOption) (*ReturnEncryptionInfo, error)
}

type aesClient struct {
	cc grpc.ClientConnInterface
}

func NewAesClient(cc grpc.ClientConnInterface) AesClient {
	return &aesClient{cc}
}

func (c *aesClient) ShowEncryption(ctx context.Context, in *PlainTextMessage, opts ...grpc.CallOption) (*ReturnEncryptionInfo, error) {
	out := new(ReturnEncryptionInfo)
	err := c.cc.Invoke(ctx, Aes_ShowEncryption_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AesServer is the server API for Aes service.
// All implementations must embed UnimplementedAesServer
// for forward compatibility
type AesServer interface {
	// Performs encryption on received message and returns information about it
	ShowEncryption(context.Context, *PlainTextMessage) (*ReturnEncryptionInfo, error)
	mustEmbedUnimplementedAesServer()
}

// UnimplementedAesServer must be embedded to have forward compatible implementations.
type UnimplementedAesServer struct {
}

func (UnimplementedAesServer) ShowEncryption(context.Context, *PlainTextMessage) (*ReturnEncryptionInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ShowEncryption not implemented")
}
func (UnimplementedAesServer) mustEmbedUnimplementedAesServer() {}

// UnsafeAesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AesServer will
// result in compilation errors.
type UnsafeAesServer interface {
	mustEmbedUnimplementedAesServer()
}

func RegisterAesServer(s grpc.ServiceRegistrar, srv AesServer) {
	s.RegisterService(&Aes_ServiceDesc, srv)
}

func _Aes_ShowEncryption_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PlainTextMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AesServer).ShowEncryption(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Aes_ShowEncryption_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AesServer).ShowEncryption(ctx, req.(*PlainTextMessage))
	}
	return interceptor(ctx, in, info, handler)
}

// Aes_ServiceDesc is the grpc.ServiceDesc for Aes service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Aes_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "aes.Aes",
	HandlerType: (*AesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ShowEncryption",
			Handler:    _Aes_ShowEncryption_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/aes/aes.proto",
}
