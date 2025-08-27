module github.com/vhive-serverless/vSwarm-proto

go 1.23.0

replace (
	github.com/vhive-serverless/vSwarm-proto/proto/aes => ./proto/aes
	github.com/vhive-serverless/vSwarm-proto/proto/auth => ./proto/auth
	github.com/vhive-serverless/vSwarm-proto/proto/compression => ./proto/compression
	github.com/vhive-serverless/vSwarm-proto/proto/fibonacci => ./proto/fibonacci
	github.com/vhive-serverless/vSwarm-proto/proto/graph_bfs => ./proto/graph_bfs
	github.com/vhive-serverless/vSwarm-proto/proto/helloworld => ./proto/helloworld
	github.com/vhive-serverless/vSwarm-proto/proto/hipstershop => ./proto/hipstershop
	github.com/vhive-serverless/vSwarm-proto/proto/hotel_reserv => ./proto/hotel_reserv
	github.com/vhive-serverless/vSwarm-proto/proto/image_rotate => ./proto/image_rotate
	github.com/vhive-serverless/vSwarm-proto/proto/rnn_serving => ./proto/rnn_serving
	github.com/vhive-serverless/vSwarm-proto/proto/video_analytics_standalone => ./proto/video_analytics_standalone
	github.com/vhive-serverless/vSwarm-proto/proto/video_processing => ./proto/video_processing
)

require (
	github.com/sirupsen/logrus v1.9.3
	github.com/vhive-serverless/vSwarm/utils/tracing/go v0.0.0-20240715172541-1edab0214c9d
	google.golang.org/grpc v1.75.0
	google.golang.org/protobuf v1.36.6
)

require (
	github.com/go-logr/logr v1.4.3 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/openzipkin/zipkin-go v0.4.3 // indirect
	go.opentelemetry.io/auto/sdk v1.1.0 // indirect
	go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc v0.51.0 // indirect
	go.opentelemetry.io/otel v1.37.0 // indirect
	go.opentelemetry.io/otel/exporters/zipkin v1.26.0 // indirect
	go.opentelemetry.io/otel/metric v1.37.0 // indirect
	go.opentelemetry.io/otel/sdk v1.37.0 // indirect
	go.opentelemetry.io/otel/trace v1.37.0 // indirect
	golang.org/x/net v0.41.0 // indirect
	golang.org/x/sys v0.33.0 // indirect
	golang.org/x/text v0.26.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250707201910-8d1bb00bc6a7 // indirect
)
