package grpcclient

import (
	pb "github.com/ease-lab/vSwarm-proto/proto/aes"
	log "github.com/sirupsen/logrus"
)

type AesClient struct {
	ClientBase
	client pb.AesClient
}

func (c *AesClient) Init(ip, port string) {
	c.Connect(ip, port)
	c.client = pb.NewAesClient(c.conn)
}

func (c *AesClient) Request(req string) string {
	r, err := c.client.ShowEncryption(c.ctx, &pb.PlainTextMessage{PlaintextMessage: req})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	return r.GetEncryptionInfo()
}
