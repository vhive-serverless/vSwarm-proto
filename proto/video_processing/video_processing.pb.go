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

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.0
// source: proto/video_processing/video_processing.proto

package video_processing

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

// The request message containing the user's name.
type SendVideo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *SendVideo) Reset() {
	*x = SendVideo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_video_processing_video_processing_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendVideo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendVideo) ProtoMessage() {}

func (x *SendVideo) ProtoReflect() protoreflect.Message {
	mi := &file_proto_video_processing_video_processing_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendVideo.ProtoReflect.Descriptor instead.
func (*SendVideo) Descriptor() ([]byte, []int) {
	return file_proto_video_processing_video_processing_proto_rawDescGZIP(), []int{0}
}

func (x *SendVideo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// The response message containing the greetings
type GetGrayscaleVideo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *GetGrayscaleVideo) Reset() {
	*x = GetGrayscaleVideo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_video_processing_video_processing_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGrayscaleVideo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGrayscaleVideo) ProtoMessage() {}

func (x *GetGrayscaleVideo) ProtoReflect() protoreflect.Message {
	mi := &file_proto_video_processing_video_processing_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGrayscaleVideo.ProtoReflect.Descriptor instead.
func (*GetGrayscaleVideo) Descriptor() ([]byte, []int) {
	return file_proto_video_processing_video_processing_proto_rawDescGZIP(), []int{1}
}

func (x *GetGrayscaleVideo) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_proto_video_processing_video_processing_proto protoreflect.FileDescriptor

var file_proto_video_processing_video_processing_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f, 0x70, 0x72,
	0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f, 0x70,
	0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x10, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e,
	0x67, 0x22, 0x1f, 0x0a, 0x09, 0x53, 0x65, 0x6e, 0x64, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x22, 0x2d, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x47, 0x72, 0x61, 0x79, 0x73, 0x63, 0x61,
	0x6c, 0x65, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x32, 0x6b, 0x0a, 0x0f, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73,
	0x73, 0x69, 0x6e, 0x67, 0x12, 0x58, 0x0a, 0x12, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x54,
	0x6f, 0x47, 0x72, 0x61, 0x79, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x12, 0x1b, 0x2e, 0x76, 0x69, 0x64,
	0x65, 0x6f, 0x5f, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x2e, 0x53, 0x65,
	0x6e, 0x64, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x1a, 0x23, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f,
	0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x2e, 0x47, 0x65, 0x74, 0x47, 0x72,
	0x61, 0x79, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x22, 0x00, 0x42, 0x41,
	0x5a, 0x3f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x76, 0x68, 0x69,
	0x76, 0x65, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x6c, 0x65, 0x73, 0x73, 0x2f, 0x76, 0x53,
	0x77, 0x61, 0x72, 0x6d, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e,
	0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_video_processing_video_processing_proto_rawDescOnce sync.Once
	file_proto_video_processing_video_processing_proto_rawDescData = file_proto_video_processing_video_processing_proto_rawDesc
)

func file_proto_video_processing_video_processing_proto_rawDescGZIP() []byte {
	file_proto_video_processing_video_processing_proto_rawDescOnce.Do(func() {
		file_proto_video_processing_video_processing_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_video_processing_video_processing_proto_rawDescData)
	})
	return file_proto_video_processing_video_processing_proto_rawDescData
}

var file_proto_video_processing_video_processing_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_video_processing_video_processing_proto_goTypes = []interface{}{
	(*SendVideo)(nil),         // 0: video_processing.SendVideo
	(*GetGrayscaleVideo)(nil), // 1: video_processing.GetGrayscaleVideo
}
var file_proto_video_processing_video_processing_proto_depIdxs = []int32{
	0, // 0: video_processing.VideoProcessing.ConvertToGrayscale:input_type -> video_processing.SendVideo
	1, // 1: video_processing.VideoProcessing.ConvertToGrayscale:output_type -> video_processing.GetGrayscaleVideo
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_video_processing_video_processing_proto_init() }
func file_proto_video_processing_video_processing_proto_init() {
	if File_proto_video_processing_video_processing_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_video_processing_video_processing_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendVideo); i {
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
		file_proto_video_processing_video_processing_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetGrayscaleVideo); i {
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
			RawDescriptor: file_proto_video_processing_video_processing_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_video_processing_video_processing_proto_goTypes,
		DependencyIndexes: file_proto_video_processing_video_processing_proto_depIdxs,
		MessageInfos:      file_proto_video_processing_video_processing_proto_msgTypes,
	}.Build()
	File_proto_video_processing_video_processing_proto = out.File
	file_proto_video_processing_video_processing_proto_rawDesc = nil
	file_proto_video_processing_video_processing_proto_goTypes = nil
	file_proto_video_processing_video_processing_proto_depIdxs = nil
}