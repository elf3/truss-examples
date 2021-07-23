package handlers

import (
	"context"
	"core/lib/etcd"
	pb "core/proto"
	"github.com/go-kit/kit/sd/etcdv3"
	"log"
)

// NewService returns a naïve, stateless implementation of Service.
func NewService() pb.EchoServer {
	options := etcdv3.ClientOptions{
		DialTimeout:   3,
		DialKeepAlive: 2,
	}
	Register, err := etcd.NewEtcdRegister(
		[]string{"127.0.0.1:2379"},
		options,
		"serverEcho",
		globalConfig.GRPCAddr,
	)
	if err != nil {
		log.Fatal(err)
	}
	Register.Register()
	return echoService{}
}

type echoService struct{}

func (s echoService) Echo(ctx context.Context, in *pb.EchoRequest) (*pb.EchoResponse, error) {
	var resp pb.EchoResponse
	return &resp, nil
}

func (s echoService) Louder(ctx context.Context, in *pb.LouderRequest) (*pb.EchoResponse, error) {
	var resp pb.EchoResponse
	return &resp, nil
}

func (s echoService) LouderGet(ctx context.Context, in *pb.LouderRequest) (*pb.EchoResponse, error) {
	var resp pb.EchoResponse
	return &resp, nil
}
