package grpcclient

import (
	"context"

	pb "github.com/vhive-serverless/vSwarm-proto/proto/graph_bfs"
)

type GraphBFSGenerator struct {
	GeneratorBase
}

func (g *GraphBFSGenerator) Next() Input {
	var pkt = g.defaultInput
	return pkt
}

func (c *GraphBFSClient) GetGenerator() Generator {
	return new(GraphBFSGenerator)
}

type GraphBFSClient struct {
	ClientBase
	client pb.GraphBFSBenchmarkClient
}

func (c *GraphBFSClient) Init(ctx context.Context, ip, port string) error {
	err := c.Connect(ctx, ip, port)
	if err != nil {
		return err
	}
	c.client = pb.NewGraphBFSBenchmarkClient(c.conn)
	return nil
}

func (c *GraphBFSClient) Request(ctx context.Context, req Input) (string, error) {
	var graphBFSMessage = req.Value
	r, err := c.client.GetBfs(ctx, &pb.GraphBFSBenchmarkRequest{Name: graphBFSMessage})
	if err != nil {
		return "", err
	}
	return r.GetMessage(), nil
}
