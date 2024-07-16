package grpcclient

import (
	"context"

	pb "github.com/vhive-serverless/vSwarm-proto/proto/compression"
)

type FileCompressGenerator struct {
	GeneratorBase
}

func (g *FileCompressGenerator) Next() Input {
	var pkt = g.defaultInput
	return pkt
}

func (c *FileCompressClient) GetGenerator() Generator {
	return new(FileCompressGenerator)
}

type FileCompressClient struct {
	ClientBase
	client pb.FileCompressClient
}

func (c *FileCompressClient) Init(ctx context.Context, ip, port string) error {
	err := c.Connect(ctx, ip, port)
	if err != nil {
		return err
	}
	c.client = pb.NewFileCompressClient(c.conn)
	return nil
}

func (c *FileCompressClient) Request(ctx context.Context, req Input) (string, error) {
	r, err := c.client.CompressFile(ctx, &pb.SendFile{Name: req.Value})
	if err != nil {
		return "", err
	}
	return r.GetMessage(), nil
}