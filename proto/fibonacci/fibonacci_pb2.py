# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/fibonacci/fibonacci.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x1fproto/fibonacci/fibonacci.proto\x12\tfibonacci\"\x1c\n\x0cHelloRequest\x12\x0c\n\x04name\x18\x01 \x01(\t\"\x1d\n\nHelloReply\x12\x0f\n\x07message\x18\x01 \x01(\t2G\n\x07Greeter\x12<\n\x08SayHello\x12\x17.fibonacci.HelloRequest\x1a\x15.fibonacci.HelloReply\"\x00\x42:Z8github.com/vhive-serverless/vSwarm-proto/proto/fibonaccib\x06proto3')



_HELLOREQUEST = DESCRIPTOR.message_types_by_name['HelloRequest']
_HELLOREPLY = DESCRIPTOR.message_types_by_name['HelloReply']
HelloRequest = _reflection.GeneratedProtocolMessageType('HelloRequest', (_message.Message,), {
  'DESCRIPTOR' : _HELLOREQUEST,
  '__module__' : 'proto.fibonacci.fibonacci_pb2'
  # @@protoc_insertion_point(class_scope:fibonacci.HelloRequest)
  })
_sym_db.RegisterMessage(HelloRequest)

HelloReply = _reflection.GeneratedProtocolMessageType('HelloReply', (_message.Message,), {
  'DESCRIPTOR' : _HELLOREPLY,
  '__module__' : 'proto.fibonacci.fibonacci_pb2'
  # @@protoc_insertion_point(class_scope:fibonacci.HelloReply)
  })
_sym_db.RegisterMessage(HelloReply)

_GREETER = DESCRIPTOR.services_by_name['Greeter']
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'Z8github.com/vhive-serverless/vSwarm-proto/proto/fibonacci'
  _HELLOREQUEST._serialized_start=46
  _HELLOREQUEST._serialized_end=74
  _HELLOREPLY._serialized_start=76
  _HELLOREPLY._serialized_end=105
  _GREETER._serialized_start=107
  _GREETER._serialized_end=178
# @@protoc_insertion_point(module_scope)
