package main

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"

	pb "grpc-sai-test/proto"
)

func main() {
	//生成证书时的hostname需要满足hostname的格式，改为grpc.sai
	creds, err := credentials.NewClientTLSFromFile("./conf/certs/server.pem", "grpc.sai")
	if err != nil {
		log.Println("Failed to create TLS credentials %v", err)
	}
	conn, err := grpc.Dial(":50052", grpc.WithTransportCredentials(creds))
	defer conn.Close()

	if err != nil {
		log.Println(err)
	}

	c := pb.NewTestClient(conn)
	context := context.Background()
	body := &pb.TestRequest{
		Referer: "Grpc",
	}

	r, err := c.SayTest(context, body)
	if err != nil {
		log.Println(err)
	}

	log.Println(r.Message)
}
