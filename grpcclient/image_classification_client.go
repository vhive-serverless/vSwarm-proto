package grpcclient

import (
	"context"
	"fmt"

	log "github.com/sirupsen/logrus"
	pb "github.com/vhive-serverless/vSwarm-proto/proto/image_classification"
)

type ImgClassificationGenerator struct {
	GeneratorBase
}

func (g *ImgClassificationGenerator) Next() Input {
	var pkt = g.defaultInput
	switch g.GeneratorBase.generator {
	case Unique:
		pkt.Value = "A unique message"
	case Linear:
		pkt.Value = "A linear message"
	case Random:
		pkt.Value = "A random message"
	}
	return pkt
}

func (c *ImgClassificationClient) GetGenerator() Generator {
	return new(ImgClassificationGenerator)
}

type ImgClassificationClient struct {
	ClientBase
	client pb.GreeterClient
}

func (c *ImgClassificationClient) Init(ctx context.Context, ip, port string) {
	c.Connect(ctx, ip, port)
	c.client = pb.NewGreeterClient(c.conn)
}

func (c *ImgClassificationClient) Request(ctx context.Context, req Input) string {
	var message = req.Value
	r, err := c.client.SayHello(ctx, &pb.HelloRequest{Name: message})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	msg := fmt.Sprintf("inference time in ns: max: %d; min: %d; mean: %d", r.GetMaxLatency(),r.GetMinLatency(),r.GetMeanLatency())
	return msg
}