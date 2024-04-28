package grpcclient

import (
	"context"
	"math/rand"

	pb "github.com/vhive-serverless/vSwarm-proto/proto/rnn_serving"
)

type RNNServingGenerator struct {
	GeneratorBase
}

func (g *RNNServingGenerator) Next() Input {

	countries := []string{
        "French", "Czech", "Dutch", "Polish", "Scottish", "Chinese", 
        "English", "Italian", "Portuguese", "Japanese", "German", 
        "Russian", "Korean", "Arabic", "Greek", "Vietnamese", "Spanish", "Irish",
    }

	var pkt = g.defaultInput
	switch g.GeneratorBase.generator {
	case Unique:
		pkt.Value = "French"
	default:
		randomCountry := countries[rand.Intn(len(countries))]
		pkt.Value = randomCountry
	}
	
	return pkt
}

func (c *RNNServingClient) GetGenerator() Generator {
	return new(RNNServingGenerator)
}

type RNNServingClient struct {
	ClientBase
	client pb.RNNServingClient
}

func (c *RNNServingClient) Init(ctx context.Context, ip, port string) error {
	err := c.Connect(ctx, ip, port)
	if err != nil {
		return err
	}
	c.client = pb.NewRNNServingClient(c.conn)
	return nil
}

func (c *RNNServingClient) Request(ctx context.Context, req Input) (string, error) {

	var numSamples int
	if g.lowerBound == g.upperBound {
		numSamples = g.lowerBound
	} else {
		numSamples = g.lowerBound + rand.Intn(g.upperBound-g.lowerBound)
	}

	r, err := c.client.GenerateString(ctx, &pb.SendLanguage{Language: float32(req.Value), NumSamples: float32(numSamples)})
	if err != nil {
		return "", err
	}
	return r.GetMessage(), nil
}