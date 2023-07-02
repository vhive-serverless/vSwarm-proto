package grpcclient

import (
	"context"

	log "github.com/sirupsen/logrus"
	pb "github.com/vhive-serverless/vSwarm-proto/proto/bert"
)

// var inputs = []string{"allow", "deny", "unbertorized", "bla bla"}

type BertGenerator struct {
	GeneratorBase
}

func (g *BertGenerator) Next() Input {
	var pkt = g.defaultInput
	switch g.GeneratorBase.generator {
	case Unique:
		pkt.Value = "A unique message"
	case Linear:
		// g.count = g.count + 1
		// pkt.Value = fmt.Sprintf("%d", g.count)
		pkt.Value = "A linear message"
	case Random:
		pkt.Value = "random"
	}
	return pkt
}

func (c *BertClient) GetGenerator() Generator {
	return new(BertGenerator)
}

type BertClient struct {
	ClientBase
	client pb.GreeterClient
}

func (c *BertClient) Init(ctx context.Context, ip, port string) {
	c.Connect(ctx, ip, port)
	c.client = pb.NewGreeterClient(c.conn)
}

func (c *BertClient) Request(ctx context.Context, req Input) string {
	var bertMessage = req.Value
	r, err := c.client.SayHello(ctx, &pb.HelloRequest{Name: bertMessage})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	return r.GetMessage()
}
