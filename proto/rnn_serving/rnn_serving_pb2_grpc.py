# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

from proto.rnn_serving import rnn_serving_pb2 as proto_dot_rnn__serving_dot_rnn__serving__pb2


class RNNServingStub(object):
    """The greeting service definition.
    """

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.GenerateString = channel.unary_unary(
                '/rnn_serving.RNNServing/GenerateString',
                request_serializer=proto_dot_rnn__serving_dot_rnn__serving__pb2.SendLanguage.SerializeToString,
                response_deserializer=proto_dot_rnn__serving_dot_rnn__serving__pb2.GetString.FromString,
                )


class RNNServingServicer(object):
    """The greeting service definition.
    """

    def GenerateString(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_RNNServingServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'GenerateString': grpc.unary_unary_rpc_method_handler(
                    servicer.GenerateString,
                    request_deserializer=proto_dot_rnn__serving_dot_rnn__serving__pb2.SendLanguage.FromString,
                    response_serializer=proto_dot_rnn__serving_dot_rnn__serving__pb2.GetString.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'rnn_serving.RNNServing', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class RNNServing(object):
    """The greeting service definition.
    """

    @staticmethod
    def GenerateString(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/rnn_serving.RNNServing/GenerateString',
            proto_dot_rnn__serving_dot_rnn__serving__pb2.SendLanguage.SerializeToString,
            proto_dot_rnn__serving_dot_rnn__serving__pb2.GetString.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)
