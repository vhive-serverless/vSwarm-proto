# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/bert/bert.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x15proto/bert/bert.proto\x12\x04\x62\x65rt\"\x1c\n\x0cHelloRequest\x12\x0c\n\x04name\x18\x01 \x01(\t\"L\n\nHelloReply\x12\x13\n\x0bmin_latency\x18\x01 \x01(\x03\x12\x13\n\x0bmax_latency\x18\x02 \x01(\x03\x12\x14\n\x0cmean_latency\x18\x03 \x01(\x03\x32=\n\x07Greeter\x12\x32\n\x08SayHello\x12\x12.bert.HelloRequest\x1a\x10.bert.HelloReply\"\x00\x42\x35Z3github.com/vhive-serverless/vSwarm-proto/proto/bertb\x06proto3')



_HELLOREQUEST = DESCRIPTOR.message_types_by_name['HelloRequest']
_HELLOREPLY = DESCRIPTOR.message_types_by_name['HelloReply']
HelloRequest = _reflection.GeneratedProtocolMessageType('HelloRequest', (_message.Message,), {
  'DESCRIPTOR' : _HELLOREQUEST,
  '__module__' : 'proto.bert.bert_pb2'
  # @@protoc_insertion_point(class_scope:bert.HelloRequest)
  })
_sym_db.RegisterMessage(HelloRequest)

HelloReply = _reflection.GeneratedProtocolMessageType('HelloReply', (_message.Message,), {
  'DESCRIPTOR' : _HELLOREPLY,
  '__module__' : 'proto.bert.bert_pb2'
  # @@protoc_insertion_point(class_scope:bert.HelloReply)
  })
_sym_db.RegisterMessage(HelloReply)

_GREETER = DESCRIPTOR.services_by_name['Greeter']
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'Z3github.com/vhive-serverless/vSwarm-proto/proto/bert'
  _HELLOREQUEST._serialized_start=31
  _HELLOREQUEST._serialized_end=59
  _HELLOREPLY._serialized_start=61
  _HELLOREPLY._serialized_end=137
  _GREETER._serialized_start=139
  _GREETER._serialized_end=200
# @@protoc_insertion_point(module_scope)
