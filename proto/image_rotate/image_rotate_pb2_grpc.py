# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

from proto.image_rotate import image_rotate_pb2 as proto_dot_image__rotate_dot_image__rotate__pb2


class ImageRotateStub(object):
    """The greeting service definition.
    """

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.RotateImage = channel.unary_unary(
                '/image_rotate.ImageRotate/RotateImage',
                request_serializer=proto_dot_image__rotate_dot_image__rotate__pb2.SendImage.SerializeToString,
                response_deserializer=proto_dot_image__rotate_dot_image__rotate__pb2.GetRotatedImage.FromString,
                )


class ImageRotateServicer(object):
    """The greeting service definition.
    """

    def RotateImage(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_ImageRotateServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'RotateImage': grpc.unary_unary_rpc_method_handler(
                    servicer.RotateImage,
                    request_deserializer=proto_dot_image__rotate_dot_image__rotate__pb2.SendImage.FromString,
                    response_serializer=proto_dot_image__rotate_dot_image__rotate__pb2.GetRotatedImage.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'image_rotate.ImageRotate', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class ImageRotate(object):
    """The greeting service definition.
    """

    @staticmethod
    def RotateImage(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/image_rotate.ImageRotate/RotateImage',
            proto_dot_image__rotate_dot_image__rotate__pb2.SendImage.SerializeToString,
            proto_dot_image__rotate_dot_image__rotate__pb2.GetRotatedImage.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)
