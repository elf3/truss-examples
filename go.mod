module core

go 1.15

require (
	github.com/go-kit/kit v0.10.0
	github.com/gogo/protobuf v1.2.2-0.20190601103108-21df5aa0e680
	github.com/gorilla/mux v1.8.0
	github.com/metaverse/truss v0.3.1
	github.com/pkg/errors v0.8.1
	golang.org/x/net v0.0.0-20190813141303-74dc4d7220e7
	google.golang.org/grpc v1.38.0
)

replace (
	github.com/coreos/etcd => github.com/ozonru/etcd v3.3.20-grpc1.26-origmodule+incompatible
	google.golang.org/grpc => google.golang.org/grpc v1.26.0
)
