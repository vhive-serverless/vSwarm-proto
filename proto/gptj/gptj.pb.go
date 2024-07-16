// MIT License

// Copyright (c) 2021 EASE Lab

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

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.0
// source: proto/gptj/gptj.proto

package gptj

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GptJBenchmarkRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Regenerate string `protobuf:"bytes,1,opt,name=regenerate,proto3" json:"regenerate,omitempty"` // Calling this will rerun the benchmark if true, else just read from benchmark results on txt file
}

func (x *GptJBenchmarkRequest) Reset() {
	*x = GptJBenchmarkRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_gptj_gptj_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GptJBenchmarkRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GptJBenchmarkRequest) ProtoMessage() {}

func (x *GptJBenchmarkRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_gptj_gptj_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GptJBenchmarkRequest.ProtoReflect.Descriptor instead.
func (*GptJBenchmarkRequest) Descriptor() ([]byte, []int) {
	return file_proto_gptj_gptj_proto_rawDescGZIP(), []int{0}
}

func (x *GptJBenchmarkRequest) GetRegenerate() string {
	if x != nil {
		return x.Regenerate
	}
	return ""
}

type GptJBenchmarkReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LatencyInfo string `protobuf:"bytes,1,opt,name=latency_info,json=latencyInfo,proto3" json:"latency_info,omitempty"`
}

func (x *GptJBenchmarkReply) Reset() {
	*x = GptJBenchmarkReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_gptj_gptj_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GptJBenchmarkReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GptJBenchmarkReply) ProtoMessage() {}

func (x *GptJBenchmarkReply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_gptj_gptj_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GptJBenchmarkReply.ProtoReflect.Descriptor instead.
func (*GptJBenchmarkReply) Descriptor() ([]byte, []int) {
	return file_proto_gptj_gptj_proto_rawDescGZIP(), []int{1}
}

func (x *GptJBenchmarkReply) GetLatencyInfo() string {
	if x != nil {
		return x.LatencyInfo
	}
	return ""
}

var File_proto_gptj_gptj_proto protoreflect.FileDescriptor

var file_proto_gptj_gptj_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x70, 0x74, 0x6a, 0x2f, 0x67, 0x70, 0x74,
	0x6a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x67, 0x70, 0x74, 0x6a, 0x22, 0x36, 0x0a,
	0x14, 0x47, 0x70, 0x74, 0x4a, 0x42, 0x65, 0x6e, 0x63, 0x68, 0x6d, 0x61, 0x72, 0x6b, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x65, 0x72,
	0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x67, 0x65, 0x6e,
	0x65, 0x72, 0x61, 0x74, 0x65, 0x22, 0x37, 0x0a, 0x12, 0x47, 0x70, 0x74, 0x4a, 0x42, 0x65, 0x6e,
	0x63, 0x68, 0x6d, 0x61, 0x72, 0x6b, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x21, 0x0a, 0x0c, 0x6c,
	0x61, 0x74, 0x65, 0x6e, 0x63, 0x79, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x6c, 0x61, 0x74, 0x65, 0x6e, 0x63, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x32, 0x57,
	0x0a, 0x0d, 0x47, 0x70, 0x74, 0x4a, 0x42, 0x65, 0x6e, 0x63, 0x68, 0x6d, 0x61, 0x72, 0x6b, 0x12,
	0x46, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x42, 0x65, 0x6e, 0x63, 0x68, 0x6d, 0x61, 0x72, 0x6b, 0x12,
	0x1a, 0x2e, 0x67, 0x70, 0x74, 0x6a, 0x2e, 0x47, 0x70, 0x74, 0x4a, 0x42, 0x65, 0x6e, 0x63, 0x68,
	0x6d, 0x61, 0x72, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x67, 0x70,
	0x74, 0x6a, 0x2e, 0x47, 0x70, 0x74, 0x4a, 0x42, 0x65, 0x6e, 0x63, 0x68, 0x6d, 0x61, 0x72, 0x6b,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x35, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x76, 0x68, 0x69, 0x76, 0x65, 0x2d, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x6c, 0x65, 0x73, 0x73, 0x2f, 0x76, 0x53, 0x77, 0x61, 0x72, 0x6d, 0x2d, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x70, 0x74, 0x6a, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_gptj_gptj_proto_rawDescOnce sync.Once
	file_proto_gptj_gptj_proto_rawDescData = file_proto_gptj_gptj_proto_rawDesc
)

func file_proto_gptj_gptj_proto_rawDescGZIP() []byte {
	file_proto_gptj_gptj_proto_rawDescOnce.Do(func() {
		file_proto_gptj_gptj_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_gptj_gptj_proto_rawDescData)
	})
	return file_proto_gptj_gptj_proto_rawDescData
}

var file_proto_gptj_gptj_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_gptj_gptj_proto_goTypes = []interface{}{
	(*GptJBenchmarkRequest)(nil), // 0: gptj.GptJBenchmarkRequest
	(*GptJBenchmarkReply)(nil),   // 1: gptj.GptJBenchmarkReply
}
var file_proto_gptj_gptj_proto_depIdxs = []int32{
	0, // 0: gptj.GptJBenchmark.GetBenchmark:input_type -> gptj.GptJBenchmarkRequest
	1, // 1: gptj.GptJBenchmark.GetBenchmark:output_type -> gptj.GptJBenchmarkReply
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_gptj_gptj_proto_init() }
func file_proto_gptj_gptj_proto_init() {
	if File_proto_gptj_gptj_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_gptj_gptj_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GptJBenchmarkRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_gptj_gptj_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GptJBenchmarkReply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_gptj_gptj_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_gptj_gptj_proto_goTypes,
		DependencyIndexes: file_proto_gptj_gptj_proto_depIdxs,
		MessageInfos:      file_proto_gptj_gptj_proto_msgTypes,
	}.Build()
	File_proto_gptj_gptj_proto = out.File
	file_proto_gptj_gptj_proto_rawDesc = nil
	file_proto_gptj_gptj_proto_goTypes = nil
	file_proto_gptj_gptj_proto_depIdxs = nil
}
