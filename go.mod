module github.com/ease-lab/vSwarm-proto

go 1.18

replace (
	github.com/ease-lab/vSwarm-proto/proto/aes => ./proto/aes
	github.com/ease-lab/vSwarm-proto/proto/auth => ./proto/auth
	github.com/ease-lab/vSwarm-proto/proto/fibonacci => ./proto/fibonacci
	github.com/ease-lab/vSwarm-proto/proto/helloworld => ./proto/helloworld
	github.com/ease-lab/vSwarm-proto/proto/hipstershop => ./proto/hipstershop
	github.com/ease-lab/vSwarm-proto/proto/hotel_reserv => ./proto/hotel_reserv
	github.com/ease-lab/vSwarm/utils/tracing/go => github.com/ease-lab/vhive/utils/tracing/go v0.0.0-20220315183234-1c8a70fc7019
)

require (
	github.com/ease-lab/vSwarm/utils/tracing/go v0.0.0-00010101000000-000000000000
	github.com/sirupsen/logrus v1.8.1
	google.golang.org/grpc v1.45.0
	google.golang.org/protobuf v1.28.0
)

require (
	github.com/containerd/containerd v1.5.10 // indirect
	github.com/openzipkin/zipkin-go v0.2.5 // indirect
	go.opentelemetry.io/contrib v0.20.0 // indirect
	go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc v0.20.0 // indirect
	go.opentelemetry.io/otel v0.20.0 // indirect
	go.opentelemetry.io/otel/exporters/trace/zipkin v0.20.0 // indirect
	go.opentelemetry.io/otel/metric v0.20.0 // indirect
	go.opentelemetry.io/otel/sdk v0.20.0 // indirect
	go.opentelemetry.io/otel/trace v0.20.0 // indirect
)

require (
	github.com/golang/protobuf v1.5.2 // indirect
	golang.org/x/net v0.0.0-20210226172049-e18ecbb05110 // indirect
	golang.org/x/sys v0.0.0-20210426230700-d19ff857e887 // indirect
	golang.org/x/text v0.3.4 // indirect
	google.golang.org/genproto v0.0.0-20201110150050-8816d57aaa9a // indirect
)
