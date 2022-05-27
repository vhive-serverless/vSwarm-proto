package grpcclient

import (
	"context"

	pb "github.com/ease-lab/vSwarm-proto/proto/aes"
	log "github.com/sirupsen/logrus"
)

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
		pkt.Value = "A random message"
	}
	return pkt
}

func (c *AesClient) GetGenerator() Generator {
	return new(AesGenerator)
}

// func (g *AesGenerator) Next() Input {
// 	var pkt Input
// 	fmt.Println(g.count)
// 	return pkt
// }

type AesClient struct {
	ClientBase
	client pb.AesClient
}

func (c *AesClient) Init(ctx context.Context, ip, port string) {
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
