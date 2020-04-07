package service

import (
	"golang.org/x/net/context"
	pb "grpc-sai-test/proto"
)

type testService struct{}

func NewTestService() *testService {
	return &testService{}
}

func (h testService) SayTest(ctx context.Context, r *pb.TestRequest) (*pb.TestResponse, error) {
	return &pb.TestResponse{
		//Message: r.Referer,
		Message : "test",
	}, nil
}

