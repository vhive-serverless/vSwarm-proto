package grpcclient

import (
	"context"

	pb "github.com/vhive-serverless/vSwarm-proto/proto/image_rotate"
)

type ImageRotateGenerator struct {
	GeneratorBase
}

func (g *ImageRotateGenerator) Next() Input {
	var pkt = g.defaultInput
	return pkt
}

func (c *ImageRotateClient) GetGenerator() Generator {
	return new(ImageRotateGenerator)
}

type ImageRotateClient struct {
	ClientBase
	client pb.ImageRotateClient
}

func (c *ImageRotateClient) Init(ctx context.Context, ip, port string) error {
	err := c.Connect(ctx, ip, port)
	if err != nil {
		return err
	}
	c.client = pb.NewImageRotateClient(c.conn)
	return nil
}

func (c *ImageRotateClient) Request(ctx context.Context, req Input) (string, error) {
	r, err := c.client.RotateImage(ctx, &pb.SendImage{SendImage: req.Value})
	if err != nil {
		return "", err
	}
	return r.GetRotatedImage(), nil
}