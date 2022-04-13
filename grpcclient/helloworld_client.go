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

func (c *HelloWorldClient) Request(req string) string {
	r, err := c.client.SayHello(c.ctx, &pb.HelloRequest{Name: req})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	return r.GetMessage()
}
