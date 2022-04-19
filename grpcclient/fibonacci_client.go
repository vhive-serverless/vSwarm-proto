package grpcclient

import (
	pb "github.com/ease-lab/vSwarm-proto/proto/fibonacci"
	log "github.com/sirupsen/logrus"
)

type FibonacciClient struct {
	ClientBase
	client pb.FibonacciClient
}

func (c *FibonacciClient) Init(ip, port string) {
	c.Connect(ip, port)
	c.client = pb.NewFibonacciClient(c.conn)
}

func (c *FibonacciClient) Request(req Input) string {
	var plainTextMessage string
	if req.generator == Unique {
		plainTextMessage = "A unique message"
	} else if req.generator == Linear {
		plainTextMessage = "A linear message"
	} else if req.generator == Random {
		plainTextMessage = "A random message"
	}
	r, err := c.client.ShowEncryption(c.ctx, &pb.PlainTextMessage{PlaintextMessage: plainTextMessage})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	return r.GetEncryptionInfo()
}
