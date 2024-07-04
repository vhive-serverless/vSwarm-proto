package grpcclient

import (
	"context"
	"fmt"

	tracing "github.com/vhive-serverless/vSwarm/utils/tracing/go"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// Base functionality ==============================================

type Input struct {
	Value  string
	Method string
}

type GeneratorType int64

const (
	Unique GeneratorType = iota
	Linear
	Random
)

type Generator interface {
	SetGenerator(gt GeneratorType)
	SetLowerBound(lb int)
	SetUpperBound(ub int)
	SetValue(value string)
	SetMethod(method string)
	Next() Input
}

type GeneratorBase struct {
	generator    GeneratorType
	lowerBound   int
	upperBound   int
	count        int
	defaultInput Input
}

func (s *GeneratorBase) SetGenerator(gt GeneratorType) {
	s.generator = gt
}
func (s *GeneratorBase) SetLowerBound(lb int) {
	s.lowerBound = lb
}
func (s *GeneratorBase) SetUpperBound(ub int) {
	s.upperBound = ub
}
func (s *GeneratorBase) SetValue(value string) {
	s.defaultInput.Value = value
}
func (s *GeneratorBase) SetMethod(method string) {
	s.defaultInput.Method = method
}

func StringToGeneratorType(gs string) GeneratorType {
	switch gs {
	case "unique":
		return Unique
	case "linear":
		return Linear
	case "random":
		return Random
	default:
		log.Error("Unknown generator type. Allowed values are: unique, linear, random.")
		return Unique
	}
}

func (s *GeneratorBase) Increment() int {
	s.count += 1
	if s.count > s.upperBound {
		s.count = s.lowerBound
	}
	return s.count
}
func (s *GeneratorBase) Decrement() int {
	s.count -= 1
	if s.count < s.lowerBound {
		s.count = s.upperBound
	}
	return s.count
}

// ------ gRPC Client interface ------
// Every client must implement this interface
type GrpcClient interface {
	Init(ctx context.Context, ip, port string) error
	Request(ctx context.Context, req Input) (string, error)
	Close()
	GetGenerator() Generator
}

// The base of the client
type ClientBase struct {
	ip   string
	port string
	conn *grpc.ClientConn
}

func (c *ClientBase) Connect(ctx context.Context, ip, port string) error {
	c.ip = ip
	c.port = port
	// Connect to the given address
	address := fmt.Sprintf("%s:%s", c.ip, c.port)
	log.Println("Connect to ", address)
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
			})
		return err
	}
	c.conn = conn
	return nil
}

func (c *ClientBase) Close() {
	c.conn.Close()
}
