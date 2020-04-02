package server

import (
	"crypto/tls"
	assetfs "github.com/elazarl/go-bindata-assetfs"
	"log"
	"net"
	"net/http"
	"path"
	"strings"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	//"github.com/elazarl/go-bindata-assetfs"

	pb "grpc-sai-test/proto"
	"grpc-sai-test/swagger"
	"grpc-sai-test/util"
)

var (
	ServerPort     string
	CertServerName string
	CertPemPath    string
	CertKeyPath    string
	SwaggerDir     string
	EndPoint       string

	tlsConfig *tls.Config
)

func Serve() (err error) {
	EndPoint = ":" + ServerPort
	tlsConfig = util.GetTLSConfig(CertPemPath, CertKeyPath)

	conn, err := net.Listen("tcp", EndPoint)
	if err != nil {
		log.Printf("TCP Listen err:%v\n", err)
	}

	srv := newServer(conn)

	log.Printf("gRPC and https listen on: %s\n", ServerPort)

	if err = srv.Serve(util.NewTLSListener(conn, tlsConfig)); err != nil {
		log.Printf("ListenAndServe: %v\n", err)
	}

	return err
}

func newServer(conn net.Listener) *http.Server {
	grpcServer := newGrpc()
	gwmux, err := newGateway()
	if err != nil {
		panic(err)
	}

	mux := http.NewServeMux()
	mux.Handle("/", gwmux)
	mux.HandleFunc("/swagger/", serveSwaggerFile)
	serveSwaggerUI(mux)

	return &http.Server{
		Addr:      EndPoint,
		Handler:   util.GrpcHandlerFunc(grpcServer, mux),
		TLSConfig: tlsConfig,
	}
}

func newGrpc() *grpc.Server {
	creds, err := credentials.NewServerTLSFromFile(CertPemPath, CertKeyPath)
	if err != nil {
		panic(err)
	}

	opts := []grpc.ServerOption{
		grpc.Creds(creds),
	}

	server := grpc.NewServer(opts...)

	pb.RegisterHelloWorldServer(server, NewHelloService())

	return server
}

func newGateway() (http.Handler, error) {
	ctx := context.Background()
	dcreds, err := credentials.NewClientTLSFromFile(CertPemPath, CertServerName)
	if err != nil {
		return nil, err
	}
	dopts := []grpc.DialOption{grpc.WithTransportCredentials(dcreds)}

	gwmux := runtime.NewServeMux()
	if err := pb.RegisterHelloWorldHandlerFromEndpoint(ctx, gwmux, EndPoint, dopts); err != nil {
		return nil, err
	}

	return gwmux, nil
}

func serveSwaggerFile(w http.ResponseWriter, r *http.Request) {
	if !strings.HasSuffix(r.URL.Path, "swagger.json") {
		log.Printf("Not Found: %s", r.URL.Path)
		http.NotFound(w, r)
		return
	}

	p := strings.TrimPrefix(r.URL.Path, "/swagger/")
	p = path.Join(SwaggerDir, p)

	log.Printf("Serving swagger-file: %s", p)

	http.ServeFile(w, r, p)
}

func serveSwaggerUI(mux *http.ServeMux) {
	fileServer := http.FileServer(&assetfs.AssetFS{
		Asset:    swagger.Asset,
		AssetDir: swagger.AssetDir,
		Prefix:   "third_party/swagger-ui",
	})
	prefix := "/swagger-ui/"
	mux.Handle(prefix, http.StripPrefix(prefix, fileServer))
}

//var (
//	ServerPort string
//	CertName string
//	CertPemPath string
//	CertKeyPath string
//	EndPoint string
//)
//
//func Serve() (err error){
//	EndPoint = ":" + ServerPort
//	conn, err := net.Listen("tcp", EndPoint)
//	if err != nil {
//		log.Printf("TCP Listen err:%v\n", err)
//	}
//
//	tlsConfig := util.GetTLSConfig(CertPemPath, CertKeyPath)
//	srv := createInternalServer(conn, tlsConfig)
//
//	log.Printf("gRPC and https listen on: %s\n", ServerPort)
//
//	if err = srv.Serve(tls.NewListener(conn, tlsConfig)); err != nil {
//		log.Printf("ListenAndServe: %v\n", err)
//	}
//
//	return err
//}
//
//func createInternalServer(conn net.Listener, tlsConfig *tls.Config) (*http.Server) {
//	var opts []grpc.ServerOption
//
//	// grpc server
//	creds, err := credentials.NewServerTLSFromFile(CertPemPath, CertKeyPath)
//	if err != nil {
//		log.Printf("Failed to create server TLS credentials %v", err)
//	}
//
//	opts = append(opts, grpc.Creds(creds))
//	grpcServer := grpc.NewServer(opts...)
//
//	// register grpc pb
//	pb.RegisterHelloWorldServer(grpcServer, NewHelloService())
//
//	// gw server
//	ctx := context.Background()
//	dcreds, err := credentials.NewClientTLSFromFile(CertPemPath, CertName)
//	if err != nil {
//		log.Printf("Failed to create client TLS credentials %v", err)
//	}
//	dopts := []grpc.DialOption{grpc.WithTransportCredentials(dcreds)}
//	gwmux := runtime.NewServeMux()
//
//	// register grpc-gateway pb
//	if err := pb.RegisterHelloWorldHandlerFromEndpoint(ctx, gwmux, EndPoint, dopts); err != nil {
//		log.Printf("Failed to register gw server: %v\n", err)
//	}
//
//	// http服务
//	mux := http.NewServeMux()
//	mux.Handle("/", gwmux)
//
//	return &http.Server{
//		Addr:      EndPoint,
//		Handler:   util.GrpcHandlerFunc(grpcServer, mux),
//		TLSConfig: tlsConfig,
//	}
//}
