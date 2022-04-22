module github.com/ease-lab/vSwarm-proto

go 1.16

replace (
	github.com/ease-lab/vSwarm-proto/proto/aes => ./proto/aes
	github.com/ease-lab/vSwarm-proto/proto/auth => ./proto/auth
	github.com/ease-lab/vSwarm-proto/proto/fibonacci => ./proto/fibonacci
	github.com/ease-lab/vSwarm-proto/proto/helloworld => ./proto/helloworld
	github.com/ease-lab/vSwarm-proto/proto/hipstershop => ./proto/hipstershop
	github.com/ease-lab/vSwarm-proto/proto/hotel_reserv => ./proto/hotel_reserv
)

require (
	github.com/ease-lab/vSwarm/utils/tracing/go v0.0.0-20220421134056-989317c4ecb5
	github.com/sirupsen/logrus v1.8.1
	google.golang.org/grpc v1.45.0
	google.golang.org/protobuf v1.28.0
)
