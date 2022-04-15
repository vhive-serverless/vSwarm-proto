package grpcclient

import (
	pb "github.com/ease-lab/vSwarm-proto/proto/helloworld"
	log "github.com/sirupsen/logrus"
)

type HelloWorldClient struct {
	ClientBase
	client pb.GreeterClient
}

func (c *HelloWorldClient) Init(ip, port string) {
	c.Connect(ip, port)
	c.client = pb.NewGreeterClient(c.conn)
}

func (c *HelloWorldClient) Request(req Input) string {
	var helloWorldMessage string
	if req.generator == Unique {
		helloWorldMessage = "A unique message"
	} else if req.generator == Linear {
		helloWorldMessage = "A linear message"
	} else if req.generator == Random {
		helloWorldMessage = "A random message"
	}
	r, err := c.client.SayHello(c.ctx, &pb.HelloRequest{Name: helloWorldMessage})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	return r.GetMessage()
}
