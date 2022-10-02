module github.com/vhive-serverless/vSwarm-proto

go 1.16

replace (
	github.com/vhive-serverless/vSwarm-proto/proto/aes => ./proto/aes
	github.com/vhive-serverless/vSwarm-proto/proto/auth => ./proto/auth
	github.com/vhive-serverless/vSwarm-proto/proto/fibonacci => ./proto/fibonacci
	github.com/vhive-serverless/vSwarm-proto/proto/helloworld => ./proto/helloworld
	github.com/vhive-serverless/vSwarm-proto/proto/hipstershop => ./proto/hipstershop
	github.com/vhive-serverless/vSwarm-proto/proto/hotel_reserv => ./proto/hotel_reserv
)

require (
	github.com/ease-lab/vSwarm/utils/tracing/go v0.0.0-20220427112636-f8f3fc171804
	github.com/sirupsen/logrus v1.8.1
	google.golang.org/grpc v1.47.0
	google.golang.org/protobuf v1.28.0
)
