package grpcclient

import (
	"context"
	"log"

	pb "github.com/vhive-serverless/vSwarm-proto/proto/parking"
)

type ParkingGenerator struct {
	GeneratorBase
}

func (g *ParkingGenerator) Next() Input {
	var pkt = g.defaultInput
	switch g.GeneratorBase.generator {
	case Unique:
		pkt.Value = "A unique message"
	case Linear:
		// g.count = g.count + 1
		// pkt.Value = fmt.Sprintf("%d", g.count)
		pkt.Value = "A linear message"
	case Random:
		pkt.Value = "Let's do spright parking"
	}
	return pkt
}

func (c *ParkingClient) GetGenerator() Generator {
	return new(ParkingGenerator)
}

type ParkingClient struct {
	ClientBase
	client pb.ParkingClient
}

func (c *ParkingClient) Init(ctx context.Context, ip, port string) error {
	log.Printf("Connect to: %s:%s\n", ip, port)
	err := c.Connect(ctx, ip, port)
	if err != nil {
		return err
	}
	c.client = pb.NewParkingClient(c.conn)
	return nil
}

func (c *ParkingClient) Request(ctx context.Context, req Input) (string, error) {

	r, err := c.client.DoParking(ctx, &pb.ParkingRequest{Name: req.Value})
	if err != nil {
		return "", err
	}
	return r.GetResult(), nil
}
