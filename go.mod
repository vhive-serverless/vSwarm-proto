module github.com/vhive-serverless/vSwarm-proto

go 1.21

toolchain go1.21.4

replace (
	github.com/vhive-serverless/vSwarm-proto/proto/aes => ./proto/aes
	github.com/vhive-serverless/vSwarm-proto/proto/auth => ./proto/auth
	github.com/vhive-serverless/vSwarm-proto/proto/compression => ./proto/compression
	github.com/vhive-serverless/vSwarm-proto/proto/fibonacci => ./proto/fibonacci
	github.com/vhive-serverless/vSwarm-proto/proto/helloworld => ./proto/helloworld
	github.com/vhive-serverless/vSwarm-proto/proto/hipstershop => ./proto/hipstershop
	github.com/vhive-serverless/vSwarm-proto/proto/hotel_reserv => ./proto/hotel_reserv
)

require (
	github.com/sirupsen/logrus v1.9.3
	github.com/vhive-serverless/vSwarm/utils/tracing/go v0.0.0-20240715172541-1edab0214c9d
	google.golang.org/grpc v1.65.0
	google.golang.org/protobuf v1.34.2
)

require (
	github.com/go-logr/logr v1.4.1 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/openzipkin/zipkin-go v0.4.3 // indirect
	go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc v0.51.0 // indirect
	go.opentelemetry.io/otel v1.26.0 // indirect
	go.opentelemetry.io/otel/exporters/zipkin v1.26.0 // indirect
	go.opentelemetry.io/otel/metric v1.26.0 // indirect
	go.opentelemetry.io/otel/sdk v1.26.0 // indirect
	go.opentelemetry.io/otel/trace v1.26.0 // indirect
	golang.org/x/net v0.25.0 // indirect
	golang.org/x/sys v0.20.0 // indirect
	golang.org/x/text v0.15.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240528184218-531527333157 // indirect
)
