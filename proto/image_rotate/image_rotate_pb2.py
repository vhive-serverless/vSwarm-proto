# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/image_rotate/image_rotate.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n%proto/image_rotate/image_rotate.proto\x12\x0cimage_rotate\"\x19\n\tSendImage\x12\x0c\n\x04name\x18\x01 \x01(\t\"\"\n\x0fGetRotatedImage\x12\x0f\n\x07message\x18\x01 \x01(\t2V\n\x0bImageRotate\x12G\n\x0bRotateImage\x12\x17.image_rotate.SendImage\x1a\x1d.image_rotate.GetRotatedImage\"\x00\x42=Z;github.com/vhive-serverless/vSwarm-proto/proto/image_rotateb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'proto.image_rotate.image_rotate_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'Z;github.com/vhive-serverless/vSwarm-proto/proto/image_rotate'
  _globals['_SENDIMAGE']._serialized_start=55
  _globals['_SENDIMAGE']._serialized_end=80
  _globals['_GETROTATEDIMAGE']._serialized_start=82
  _globals['_GETROTATEDIMAGE']._serialized_end=116
  _globals['_IMAGEROTATE']._serialized_start=118
  _globals['_IMAGEROTATE']._serialized_end=204
# @@protoc_insertion_point(module_scope)
