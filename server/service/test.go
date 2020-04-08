package service

import (
	"golang.org/x/net/context"
	pb "grpc-sai-test/proto"
	"grpc-sai-test/server/service/models"
	"grpc-sai-test/store/log"
)

type testService struct{}

func NewTestService() *testService {
	return &testService{}
}

func (h testService) SayTest(ctx context.Context, r *pb.TestRequest) (*pb.TestResponse, error) {

	uid, err := models.TUser{}.Create()

	if err != nil {
		log.Logger.Println(err)
		return &pb.TestResponse{
			//Message: r.Referer,
			Message: "failure",
		}, err
	}

	return &pb.TestResponse{
		//Message: r.Referer,
		Message: uid,
	}, nil
}
