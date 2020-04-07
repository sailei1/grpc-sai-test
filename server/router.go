package server

import (
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	pb "grpc-sai-test/proto"
	"grpc-sai-test/server/service"
	"net/http"
	"golang.org/x/net/context"
)

func newGrpc(CertPemPath string,CertKeyPath string) *grpc.Server {
	creds, err := credentials.NewServerTLSFromFile(CertPemPath, CertKeyPath)
	if err != nil {
		panic(err)
	}

	opts := []grpc.ServerOption{
		grpc.Creds(creds),
	}

	server := grpc.NewServer(opts...)


	//
	pb.RegisterHelloWorldServer(server, service.NewHelloService())
	pb.RegisterTestServer(server, service.NewTestService())

	return server
}

func newGateway(CertPemPath string,CertServerName string) (http.Handler, error) {
	ctx := context.Background()
	dcreds, err := credentials.NewClientTLSFromFile(CertPemPath, CertServerName)
	if err != nil {
		return nil, err
	}
	dopts := []grpc.DialOption{grpc.WithTransportCredentials(dcreds)}

	gwmux := runtime.NewServeMux()
	//
	if err := pb.RegisterHelloWorldHandlerFromEndpoint(ctx, gwmux, EndPoint, dopts); err != nil {
		return nil, err
	}

	if err := pb.RegisterTestHandlerFromEndpoint(ctx, gwmux, EndPoint, dopts); err != nil {
		return nil, err
	}

	return gwmux, nil
}
