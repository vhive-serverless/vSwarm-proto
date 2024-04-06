package grpcclient

import (
	"context"

	pb "github.com/vhive-serverless/vSwarm-proto/proto/video_processing"
)

type VideoProcessingGenerator struct {
	GeneratorBase
}

func (g *VideoProcessingGenerator) Next() Input {
	var pkt = g.defaultInput
	return pkt
}

func (c *VideoProcessingClient) GetGenerator() Generator {
	return new(VideoProcessingGenerator)
}

type VideoProcessingClient struct {
	ClientBase
	client pb.VideoProcessingClient
}

func (c *VideoProcessingClient) Init(ctx context.Context, ip, port string) error {
	err := c.Connect(ctx, ip, port)
	if err != nil {
		return err
	}
	c.client = pb.NewVideoProcessingClient(c.conn)
	return nil
}

func (c *VideoProcessingClient) Request(ctx context.Context, req Input) (string, error) {
	r, err := c.client.ConvertToGrayscale(ctx, &pb.SendVideo{Name: req.Value})
	if err != nil {
		return "", err
	}
	return r.GetMessage(), nil
}