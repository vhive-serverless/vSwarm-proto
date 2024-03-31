# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

import gptj_pb2 as gptj__pb2


class GptJBenchmarkStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.GetBenchmark = channel.unary_unary(
                '/gptj.GptJBenchmark/GetBenchmark',
                request_serializer=gptj__pb2.GptJBenchmarkRequest.SerializeToString,
                response_deserializer=gptj__pb2.GptJBenchmarkReply.FromString,
                )


class GptJBenchmarkServicer(object):
    """Missing associated documentation comment in .proto file."""

    def GetBenchmark(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_GptJBenchmarkServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'GetBenchmark': grpc.unary_unary_rpc_method_handler(
                    servicer.GetBenchmark,
                    request_deserializer=gptj__pb2.GptJBenchmarkRequest.FromString,
                    response_serializer=gptj__pb2.GptJBenchmarkReply.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'gptj.GptJBenchmark', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class GptJBenchmark(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def GetBenchmark(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/gptj.GptJBenchmark/GetBenchmark',
            gptj__pb2.GptJBenchmarkRequest.SerializeToString,
            gptj__pb2.GptJBenchmarkReply.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)
