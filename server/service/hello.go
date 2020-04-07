package service

import (
	"golang.org/x/net/context"
	pb "grpc-sai-test/proto"
)

type helloService struct{}

func NewHelloService() *helloService {
	return &helloService{}
}

func (h helloService) SayHelloWorld(ctx context.Context, r *pb.HelloWorldRequest) (*pb.HelloWorldResponse, error) {
	return &pb.HelloWorldResponse{
		Message: r.Referer,
		//Message : "test",
	}, nil
}

