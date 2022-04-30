package grpcclient

import (
	"fmt"
	"math/rand"

	pb "github.com/ease-lab/vSwarm-proto/proto/fibonacci"
	log "github.com/sirupsen/logrus"
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
		fibNum := rand.Intn(1)
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

func (c *FibonacciClient) Init(ip, port string) {
	c.Connect(ip, port)
	c.client = pb.NewGreeterClient(c.conn)
}

func (c *FibonacciClient) Request(req Input) string {
	var fibonacciMessage = req.Value
	r, err := c.client.SayHello(c.ctx, &pb.HelloRequest{Name: fibonacciMessage})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	return r.GetMessage()
}
