package grpcclient

import (
	pb "github.com/ease-lab/vSwarm-proto/proto/auth"
	log "github.com/sirupsen/logrus"
)

type AuthClient struct {
	ClientBase
	client pb.GreeterClient
}

func (c *AuthClient) Init(ip, port string) {
	c.Connect(ip, port)
	c.client = pb.NewGreeterClient(c.conn)
}

func (c *AuthClient) Request(req Input) string {
	var authMessage string
	if req.generator == Unique {
		authMessage = "A unique message"
	} else if req.generator == Linear {
		authMessage = "A linear message"
	} else if req.generator == Random {
		authMessage = "A random message"
	} else {
		log.WithFields(
			log.Fields{
				"event": "Send Request to benchmark server",
				"key":   "Invalid GeneratorType",
			}).Fatal("Failed to determine generator.")
	}

	r, err := c.client.SayHello(c.ctx, &pb.HelloRequest{Name: authMessage})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	return r.GetMessage()
}
