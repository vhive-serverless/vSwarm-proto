package grpcclient

import (
	"context"

	pb "github.com/vhive-serverless/vSwarm-proto/proto/compression"
)

type CompressionGenerator struct {
	GeneratorBase
}

func (g *CompressionGenerator) Next() Input {
	var pkt = g.defaultInput
	return pkt
}

func (c *CompressionClient) GetGenerator() Generator {
	return new(CompressionGenerator)
}

type CompressionClient struct {
	ClientBase
	client pb.CompressionClient
}

func (c *CompressionClient) Init(ctx context.Context, ip, port string) error {
	err := c.Connect(ctx, ip, port)
	if err != nil {
		return err
	}
	c.client = pb.NewCompressionClient(c.conn)
	return nil
}

func (c *CompressionClient) Request(ctx context.Context, req Input) (string, error) {
	r, err := c.client.FileCompress(ctx, &pb.SendImage{Name: req.Value})
	if err != nil {
		return "", err
	}
	return r.GetMessage(), nil
}