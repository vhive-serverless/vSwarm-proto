# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/auth/auth.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x15proto/auth/auth.proto\x12\x04\x61uth\"\x1c\n\x0cHelloRequest\x12\x0c\n\x04name\x18\x01 \x01(\t\"\x1d\n\nHelloReply\x12\x0f\n\x07message\x18\x01 \x01(\t2=\n\x07Greeter\x12\x32\n\x08SayHello\x12\x12.auth.HelloRequest\x1a\x10.auth.HelloReply\"\x00\x42\x35Z3github.com/vhive-serverless/vSwarm-proto/proto/authb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'proto.auth.auth_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'Z3github.com/vhive-serverless/vSwarm-proto/proto/auth'
  _globals['_HELLOREQUEST']._serialized_start=31
  _globals['_HELLOREQUEST']._serialized_end=59
  _globals['_HELLOREPLY']._serialized_start=61
  _globals['_HELLOREPLY']._serialized_end=90
  _globals['_GREETER']._serialized_start=92
  _globals['_GREETER']._serialized_end=153
# @@protoc_insertion_point(module_scope)
