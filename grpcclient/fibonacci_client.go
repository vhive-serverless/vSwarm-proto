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

func (c *FibonacciClient) Request(req string) string {
	r, err := c.client.ShowEncryption(c.ctx, &pb.PlainTextMessage{PlaintextMessage: req})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	return r.GetEncryptionInfo()
}
