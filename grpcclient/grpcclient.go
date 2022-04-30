package grpcclient

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	tracing "github.com/ease-lab/vSwarm/utils/tracing/go"

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

// ------ gRPC Client interface ------
// Every client must implement this interface
type GrpcClient interface {
	Init(ip, port string)
	Request(req Input) string
	Close()
	GetGenerator() Generator
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

func getMethodPayload(req Input) (method int, payload string) {
	payload = req.Value
	// In case we have specified the exact request we want in the string extract the info
	str := strings.SplitN(req.Value, "|", 2)
	if len(str) == 2 {
		method, _ = strconv.Atoi(strings.ReplaceAll(str[0], " ", ""))
		payload = strings.ReplaceAll(str[1], " ", "")
	}
	return
}
