module github.com/vhive-serverless/vSwarm-proto

go 1.18

replace (
	github.com/vhive-serverless/vSwarm-proto/proto/aes => ./proto/aes
	github.com/vhive-serverless/vSwarm-proto/proto/auth => ./proto/auth
	github.com/vhive-serverless/vSwarm-proto/proto/fibonacci => ./proto/fibonacci
	github.com/vhive-serverless/vSwarm-proto/proto/helloworld => ./proto/helloworld
	github.com/vhive-serverless/vSwarm-proto/proto/hipstershop => ./proto/hipstershop
	github.com/vhive-serverless/vSwarm-proto/proto/hotel_reserv => ./proto/hotel_reserv
	github.com/vhive-serverless/vSwarm-proto/proto/image_classification => ./proto/image_classification
)

require (
	github.com/sirupsen/logrus v1.9.3
	github.com/vhive-serverless/vSwarm/utils/tracing/go v0.0.0-20230802102142-dbfda39fc27c
	google.golang.org/grpc v1.58.2
	google.golang.org/protobuf v1.31.0
)

require (
	cloud.google.com/go/compute v1.23.0 // indirect
	github.com/containerd/containerd v1.7.3 // indirect
	github.com/go-logr/logr v1.2.4 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/openzipkin/zipkin-go v0.4.1 // indirect
	go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc v0.42.0 // indirect
	go.opentelemetry.io/otel v1.16.0 // indirect
	go.opentelemetry.io/otel/exporters/zipkin v1.8.0 // indirect
	go.opentelemetry.io/otel/metric v1.16.0 // indirect
	go.opentelemetry.io/otel/sdk v1.14.0 // indirect
	go.opentelemetry.io/otel/trace v1.16.0 // indirect
	golang.org/x/net v0.13.0 // indirect
	golang.org/x/sys v0.10.0 // indirect
	golang.org/x/text v0.11.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20230731193218-e0aa005b6bdf // indirect
)
