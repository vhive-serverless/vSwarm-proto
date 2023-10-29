package grpcclient

import (
	"context"
	log "github.com/sirupsen/logrus"
	pb "github.com/vhive-serverless/vSwarm-proto/proto/gptj"
)


type GptJGenerator struct {
	GeneratorBase
}

func (g *GptJGenerator) Next() Input {
	var pkt = g.defaultInput
	switch g.GeneratorBase.generator {
	case Unique:
		pkt.Value = "false"
	case Linear:
		pkt.Value = "false"
	case Random:
		pkt.Value = "true"
	}
	return pkt
}

type GptJClient struct {
	ClientBase
	client pb.GptJBenchmarkClient
}

func (c *GptJClient) GetGenerator() Generator {
	return new(GptJGenerator)
}


func (c *GptJClient) Init(ctx context.Context, ip, port string) {
	log.Printf("Connect to: %s:%s\n", ip, port)
	c.Connect(ctx, ip, port)
	c.client = pb.NewGptJBenchmarkClient(c.conn)
}

func (c *GptJClient) Request(ctx context.Context, req Input) string {
	r, err := c.client.GetBenchmark(ctx, &pb.GptJBenchmarkRequest{Regenerate: req.Value})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	return r.GetLatencyInfo();
}
