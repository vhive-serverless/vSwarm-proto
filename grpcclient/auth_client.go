package grpcclient

import (
	"context"
	"math/rand"

	pb "github.com/vhive-serverless/vSwarm-proto/proto/auth"
)

var inputs = []string{"allow", "deny", "unauthorized", "bla bla"}

type AuthGenerator struct {
	GeneratorBase
}

func (g *AuthGenerator) Next() Input {
	var pkt = g.defaultInput
	switch g.GeneratorBase.generator {
	case Unique:
		pkt.Value = "A unique message"
	case Linear:
		// g.count = g.count + 1
		// pkt.Value = fmt.Sprintf("%d", g.count)
		pkt.Value = "A linear message"
	case Random:
		pkt.Value = inputs[rand.Intn(len(inputs))]
	}
	return pkt
}

func (c *AuthClient) GetGenerator() Generator {
	return new(AuthGenerator)
}

type AuthClient struct {
	ClientBase
	client pb.GreeterClient
}

func (c *AuthClient) Init(ctx context.Context, ip, port string) error {
	err := c.Connect(ctx, ip, port)
	if err != nil {
		return err
	}
	c.client = pb.NewGreeterClient(c.conn)
	return nil
}

func (c *AuthClient) Request(ctx context.Context, req Input) (string, error) {
	var authMessage = req.Value
	r, err := c.client.SayHello(ctx, &pb.HelloRequest{Name: authMessage})
	if err != nil {
		return "", err
	}
	return r.GetMessage(), nil
}
