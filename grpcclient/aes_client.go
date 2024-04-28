package grpcclient

import (
	"context"
	"math/rand"

	pb "github.com/vhive-serverless/vSwarm-proto/proto/aes"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

type AesGenerator struct {
	GeneratorBase
}

func (g *AesGenerator) Next() Input {
	var pkt = g.defaultInput
	switch g.GeneratorBase.generator {
	case Unique:
		pkt.Value = "A unique message"
	case Linear:
		// g.count = g.count + 1
		// pkt.Value = fmt.Sprintf("%d", g.count)
		pkt.Value = "A linear message"
	case Random:
		pkt.Value = g.randSeq()
	}
	return pkt
}

func (g *AesGenerator) randSeq() string {
	// Make random length
	var l int
	if g.lowerBound == g.upperBound {
		l = g.lowerBound
	} else {
		l = g.lowerBound + rand.Intn(g.upperBound-g.lowerBound)
	}

	b := make([]rune, l)
	// make the string
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func (c *AesClient) GetGenerator() Generator {
	return new(AesGenerator)
}

type AesClient struct {
	ClientBase
	client pb.AesClient
}

func (c *AesClient) Init(ctx context.Context, ip, port string) error {
	err := c.Connect(ctx, ip, port)
	if err != nil {
		return err
	}
	c.client = pb.NewAesClient(c.conn)
	return nil
}

func (c *AesClient) Request(ctx context.Context, req Input) (string, error) {
	r, err := c.client.ShowEncryption(ctx, &pb.PlainTextMessage{PlaintextMessage: req.Value})
	if err != nil {
		return "", err
	}
	return r.GetEncryptionInfo(), nil
}
