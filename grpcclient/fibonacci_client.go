package grpcclient

import (
	"math/rand"
	"strconv"

	pb "github.com/ease-lab/vSwarm-proto/proto/fibonacci"
	log "github.com/sirupsen/logrus"
)

type FibonacciClient struct {
	ClientBase
	client pb.GreeterClient
}

func (c *FibonacciClient) Init(ip, port string) {
	c.Connect(ip, port)
	c.client = pb.NewGreeterClient(c.conn)
}

func (c *FibonacciClient) Request(req Input) string {
	var fibonacciMessage string
	if req.generator == Unique {
		fibonacciMessage = req.value
	} else if req.generator == Linear {
		value64, err := strconv.ParseInt(req.value, 10, 0)
		value := int(value64)
		if err != nil {
			log.Fatalf("Could not parse value for fibonacci")
		}
		fibonacciMessage = strconv.Itoa(value + req.count*2) // Arbitrary linear factor of 2
	} else if req.generator == Random {
		fibNum := req.lowerBound + rand.Intn(req.upperBound-req.lowerBound)
		fibonacciMessage = strconv.Itoa(fibNum)
	} else {
		log.WithFields(
			log.Fields{
				"event": "Send Request to benchmark server",
				"key":   "Invalid GeneratorType",
			}).Fatal("Failed to determine generator.")
	}

	r, err := c.client.SayHello(c.ctx, &pb.HelloRequest{Name: fibonacciMessage})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	return r.GetMessage()
}
