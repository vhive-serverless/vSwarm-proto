package grpcclient

import (
	"context"
	"fmt"
	"math/rand"

	pb "github.com/vhive-serverless/vSwarm-proto/proto/helloworld"
	log "github.com/sirupsen/logrus"
)

type GamingAnywhereGenerator struct {
	GeneratorBase
}

func (g *GamingAnywhereGenerator) Next() Input {
	var pkt = g.defaultInput
	switch g.GeneratorBase.generator {
	case Unique:
		pkt.Value = "A unique message"
	case Linear:
		// g.count = g.count + 1
		// pkt.Value = fmt.Sprintf("%d", g.count)
		pkt.Value = "A linear message"
	case Random:
		pkt.Value = "A random message"
	}
	return pkt
}

func (c *GamingAnywhereClient) GetGenerator() Generator {
	return new(GamingAnywhereGenerator)
}

type GamingAnywhereClient struct {
	ClientBase
	client pb.GreeterClient
}

func (c *GamingAnywhereClient) Init(ctx context.Context, ip, port string) {
	c.Connect(ctx, ip, port)
	c.client = pb.NewGreeterClient(c.conn)
}

func (c *GamingAnywhereClient) Request(ctx context.Context, req Input) string {
	var message = req.Value
	r, err := c.client.SayHello(ctx, &pb.HelloRequest{Name: message})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	return r.GetMessage()
}
