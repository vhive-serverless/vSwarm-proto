package grpcclient

import (
	pb "github.com/ease-lab/vSwarm-proto/proto/auth"
	log "github.com/sirupsen/logrus"
)

type AuthClient struct {
	ClientBase
	client pb.AuthClient
}

func (c *AuthClient) Init(ip, port string) {
	c.Connect(ip, port)
	c.client = pb.NewAuthClient(c.conn)
}

func (c *AuthClient) Request(req string) string {
	r, err := c.client.ShowEncryption(c.ctx, &pb.PlainTextMessage{PlaintextMessage: req})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	return r.GetEncryptionInfo()
}
