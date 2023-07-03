module github.com/vhive-serverless/vSwarm-proto

go 1.16

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
	cloud.google.com/go/bigquery v1.19.0 // indirect
	cloud.google.com/go/compute/metadata v0.2.0 // indirect
	cloud.google.com/go/datastore v1.5.0 // indirect
	cloud.google.com/go/firestore v1.5.0 // indirect
	cloud.google.com/go/logging v1.5.0-jsonlog-preview // indirect
	cloud.google.com/go/pubsub v1.12.0 // indirect
	cloud.google.com/go/pubsublite v0.10.2 // indirect
	cloud.google.com/go/spanner v1.21.0 // indirect
	cloud.google.com/go/storage v1.16.0 // indirect
	github.com/AdaLogics/go-fuzz-headers v0.0.0-20210715213245-6c3934b029d8 // indirect
	github.com/Microsoft/hcsshim v0.9.0 // indirect
	github.com/Shopify/sarama v1.30.0 // indirect
	github.com/census-instrumentation/opencensus-proto v0.3.0 // indirect
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/checkpoint-restore/go-criu/v5 v5.1.0 // indirect
	github.com/cilium/ebpf v0.7.0 // indirect
	github.com/cncf/udpa/go v0.0.0-20220112060539-c52dc94e7fbe // indirect
	github.com/cncf/xds/go v0.0.0-20230105202645-06c439db220b // indirect
	github.com/containerd/cgroups v1.0.3 // indirect
	github.com/containerd/continuity v0.2.2 // indirect
	github.com/containerd/go-cni v1.1.6 // indirect
	github.com/containerd/imgcrypt v1.1.2 // indirect
	github.com/containernetworking/plugins v1.1.1 // indirect
	github.com/containers/ocicrypt v1.1.3 // indirect
	github.com/cyphar/filepath-securejoin v0.2.3 // indirect
	github.com/envoyproxy/protoc-gen-validate v0.6.1 // indirect
	github.com/godbus/dbus/v5 v5.0.6 // indirect
	github.com/golang/glog v1.0.0 // indirect
	github.com/google/go-cmp v0.5.9 // indirect
	github.com/google/pprof v0.0.0-20210720184732-4bb14d4b1be1 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/googleapis/enterprise-certificate-proxy v0.2.0 // indirect
	github.com/googleapis/go-type-adapters v1.0.0 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.5.0 // indirect
	github.com/iancoleman/strcase v0.2.0 // indirect
	github.com/intel/goresctrl v0.1.0 // indirect
	github.com/j-keck/arping v1.0.2 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/lyft/protoc-gen-star v0.6.1 // indirect
	github.com/moby/sys/mountinfo v0.5.0 // indirect
	github.com/moby/sys/signal v0.6.0 // indirect
	github.com/moby/sys/symlink v0.2.0 // indirect
	github.com/onsi/ginkgo v1.16.5 // indirect
	github.com/opencontainers/image-spec v1.0.3-0.20211202183452-c5a74bcca799 // indirect
	github.com/opencontainers/runc v1.0.3 // indirect
	github.com/opencontainers/selinux v1.10.1 // indirect
	github.com/pelletier/go-toml v1.9.3 // indirect
	github.com/prometheus/client_golang v1.11.1 // indirect
	github.com/prometheus/common v0.30.0 // indirect
	github.com/prometheus/procfs v0.7.3 // indirect
	github.com/rabbitmq/amqp091-go v1.1.0 // indirect
	github.com/seccomp/libseccomp-golang v0.9.2-0.20210429002308-3879420cc921 // indirect
	github.com/sirupsen/logrus v1.9.2
	github.com/spf13/afero v1.9.2 // indirect
	go.opencensus.io v0.24.0 // indirect
	go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc v0.31.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/internal/retry v1.3.0 // indirect
	go.opentelemetry.io/otel/exporters/zipkin v1.1.0 // indirect
	go.opentelemetry.io/otel/sdk v1.8.0 // indirect
	go.opentelemetry.io/proto/otlp v0.9.0 // indirect
	golang.org/x/net v0.5.0 // indirect
	golang.org/x/time v0.1.0 // indirect
	golang.org/x/tools v0.3.0 // indirect
	golang.org/x/xerrors v0.0.0-20220907171357-04be3eba64a2 // indirect
	google.golang.org/api v0.50.0 // indirect
	google.golang.org/grpc v1.46.0-dev
	google.golang.org/protobuf v1.26.0
	k8s.io/apiserver v0.22.5 // indirect
	k8s.io/cri-api v0.23.0-alpha.4 // indirect
	k8s.io/klog/v2 v2.30.0 // indirect
	k8s.io/utils v0.0.0-20210930125809-cb0fa318a74b // indirect
)
