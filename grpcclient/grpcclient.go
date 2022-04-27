package grpcclient

import (
	"context"
	"fmt"

	tracing "github.com/ease-lab/vSwarm/utils/tracing/go"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// Base functionality ==============================================

type GeneratorType int64

const (
	Unique GeneratorType = iota
	Linear
	Random
)

type Input struct {
	generator  GeneratorType
	lowerBound int
	upperBound int
	value      string
	count      int
	method     string
}

func (s *Input) SetGenerator(gt GeneratorType) {
	s.generator = gt
}
func (s *Input) SetLowerBound(lb int) {
	s.lowerBound = lb
}
func (s *Input) SetUpperBound(ub int) {
	s.upperBound = ub
}
func (s *Input) SetValue(value string) {
	s.value = value
}
func (s *Input) SetMethod(method string) {
	s.method = method
}

// ------ gRPC Client interface ------
// Every client must implement this interface
type GrpcClient interface {
	Init(ip, port string)
	Request(req Input) string
	Close()
}

// The base of the client
type ClientBase struct {
	ip   string
	port string
	ctx  context.Context
	conn *grpc.ClientConn
}

func (c *ClientBase) Connect(ip, port string) {
	c.ip = ip
	c.port = port
	// Connect to the given address
	address := fmt.Sprintf("%s:%s", c.ip, c.port)
	log.Debug("Connect to ", address)
	var conn *grpc.ClientConn
	var err error
	if tracing.IsTracingEnabled() {
		log.Debug("Tracing is enabled.")
		conn, err = tracing.DialGRPCWithUnaryInterceptor(address, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	} else {
		log.Debug("Tracing is disabled.")
		conn, err = grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	}
	if err != nil {
		log.WithFields(
			log.Fields{
				"event": "Connecting to benchmark server",
				"key":   err,
			}).Fatal("Failed to connect.")
	}
	c.conn = conn

	// Create a new context.
	// The context is used all the time while the connection is established
	ctx, _ := context.WithCancel(context.Background())
	c.ctx = ctx
}

func (c *ClientBase) Close() {
	c.conn.Close()
}
