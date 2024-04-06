# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

from proto.video_processing import video_processing_pb2 as proto_dot_video__processing_dot_video__processing__pb2


class VideoProcessingStub(object):
    """The greeting service definition.
    """

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.ConvertToGrayscale = channel.unary_unary(
                '/video_processing.VideoProcessing/ConvertToGrayscale',
                request_serializer=proto_dot_video__processing_dot_video__processing__pb2.SendVideo.SerializeToString,
                response_deserializer=proto_dot_video__processing_dot_video__processing__pb2.GetGrayscaleVideo.FromString,
                )


class VideoProcessingServicer(object):
    """The greeting service definition.
    """

    def ConvertToGrayscale(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_VideoProcessingServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'ConvertToGrayscale': grpc.unary_unary_rpc_method_handler(
                    servicer.ConvertToGrayscale,
                    request_deserializer=proto_dot_video__processing_dot_video__processing__pb2.SendVideo.FromString,
                    response_serializer=proto_dot_video__processing_dot_video__processing__pb2.GetGrayscaleVideo.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'video_processing.VideoProcessing', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class VideoProcessing(object):
    """The greeting service definition.
    """

    @staticmethod
    def ConvertToGrayscale(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/video_processing.VideoProcessing/ConvertToGrayscale',
            proto_dot_video__processing_dot_video__processing__pb2.SendVideo.SerializeToString,
            proto_dot_video__processing_dot_video__processing__pb2.GetGrayscaleVideo.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)
