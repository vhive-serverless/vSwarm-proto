package grpcclient

import (
	"context"
	"fmt"
	"math/rand"

	pb "github.com/vhive-serverless/vSwarm-proto/proto/fibonacci"
)

type FibonacciGenerator struct {
	GeneratorBase
}

func (g *FibonacciGenerator) Next() Input {
	var pkt = g.defaultInput
	switch g.GeneratorBase.generator {
	case Unique:

	case Linear:
		pkt.Value = fmt.Sprintf("%d", g.Increment())

	case Random:
		fibNum := g.lowerBound + rand.Intn(g.upperBound-g.lowerBound)
		pkt.Value = fmt.Sprintf("%d", fibNum)
	}
	return pkt
}

func (c *FibonacciClient) GetGenerator() Generator {
	return new(FibonacciGenerator)
}

type FibonacciClient struct {
	ClientBase
	client pb.GreeterClient
}

func (c *FibonacciClient) Init(ctx context.Context, ip, port string) error {
	err := c.Connect(ctx, ip, port)
	if err != nil {
		return err
	}
	c.client = pb.NewGreeterClient(c.conn)
	return nil
}

func (c *FibonacciClient) Request(ctx context.Context, req Input) (string, error) {
	var fibonacciMessage = req.Value
	r, err := c.client.SayHello(ctx, &pb.HelloRequest{Name: fibonacciMessage})
	if err != nil {
		return "", err
	}
	return r.GetMessage(), nil
}
