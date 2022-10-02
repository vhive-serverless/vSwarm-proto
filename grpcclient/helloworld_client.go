package grpcclient

import (
	"context"

	pb "github.com/vhive-serverless/vSwarm-proto/proto/helloworld"
	log "github.com/sirupsen/logrus"
)

// type HelloWorldGenerator struct {
// 	Generator
// }

// func (g *HelloWorldGenerator) Next() Input {
// 	g.count = g.count + 1
// 	var pkt Input
// 	pkt.value = fmt.Sprintf("%d", g.count)
// 	return pkt
// }

// type AesClient struct {
// 	ClientBase
// 	client pb.AesClient
// 	gen    HelloWorldGenerator
// }
// func (g *HelloWorldClient) Next() Input {
// 	g.count = g.count + 1
// 	var pkt Input
// 	pkt.value = fmt.Sprintf("%d", g.count)
// 	return pkt
// }

type HelloWorldGenerator struct {
	GeneratorBase
}

func (g *HelloWorldGenerator) Next() Input {
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

func (c *HelloWorldClient) GetGenerator() Generator {
	return new(HelloWorldGenerator)
}

type HelloWorldClient struct {
	ClientBase
	client pb.GreeterClient
}

func (c *HelloWorldClient) Init(ctx context.Context, ip, port string) {
	c.Connect(ctx, ip, port)
	c.client = pb.NewGreeterClient(c.conn)
}

func (c *HelloWorldClient) Request(ctx context.Context, req Input) string {
	r, err := c.client.SayHello(ctx, &pb.HelloRequest{Name: req.Value})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	return r.GetMessage()
}
