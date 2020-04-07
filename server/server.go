package server

import (
	"crypto/tls"
	assetfs "github.com/elazarl/go-bindata-assetfs"
	"github.com/golang/glog"
	"log"
	"net"
	"net/http"
	"path"
	"strings"

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

func Run() (err error) {
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

func preflightHandler(w http.ResponseWriter, r *http.Request) {
	headers := []string{"Content-Type", "Accept"}
	w.Header().Set("Access-Control-Allow-Headers", strings.Join(headers, ","))
	methods := []string{"GET", "HEAD", "POST", "PUT", "DELETE"}
	w.Header().Set("Access-Control-Allow-Methods", strings.Join(methods, ","))
	glog.Infof("preflight request for %s", r.URL.Path)
	return
}

func allowCORS(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if origin := r.Header.Get("Origin"); origin != "" {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			if r.Method == "OPTIONS" && r.Header.Get("Access-Control-Request-Method") != "" {
				preflightHandler(w, r)
				return
			}
		}
		h.ServeHTTP(w, r)
	})
}

func newServer(conn net.Listener) *http.Server {
	grpcServer := newGrpc(CertPemPath,CertKeyPath)
	gwmux, err := newGateway(CertPemPath,CertServerName)
	if err != nil {
		panic(err)
	}

	mux := http.NewServeMux()
	mux.Handle("/", gwmux)
	mux.HandleFunc("/swagger/", serveSwaggerFile)
	serveSwaggerUI(mux)

	return &http.Server{
		Addr:      EndPoint,
		Handler:   util.GrpcHandlerFunc(grpcServer,allowCORS(mux)),
		TLSConfig: tlsConfig,
	}
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
