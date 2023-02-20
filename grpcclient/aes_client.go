package grpcclient

import (
	"context"
	"math/rand"

	log "github.com/sirupsen/logrus"
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
	l := g.lowerBound + rand.Intn(g.upperBound-g.lowerBound)
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

func (c *AesClient) Init(ctx context.Context, ip, port string) {
	log.Printf("Connect to: %s:%s\n", ip, port)
	c.Connect(ctx, ip, port)
	c.client = pb.NewAesClient(c.conn)
}

func (c *AesClient) Request(ctx context.Context, req Input) string {
	r, err := c.client.ShowEncryption(ctx, &pb.PlainTextMessage{PlaintextMessage: req.Value})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	return r.GetEncryptionInfo()
}
