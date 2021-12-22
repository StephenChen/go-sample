package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"grpc/proto"
	"io/ioutil"
	"log"
	"net"
	"net/http"
)

func main() {
	go func() {
		ctx := context.Background()
		ctx, cancel := context.WithCancel(ctx)
		defer cancel()

		mux := runtime.NewServeMux()

		err := proto.RegisterRestServiceHandlerFromEndpoint(
			ctx, mux, "localhost:5000",
			[]grpc.DialOption{grpc.WithInsecure()},
		)
		if err != nil {
			log.Fatal(err)
		}

		http.ListenAndServe(":8080", mux)
	}()

	certificate, err := tls.LoadX509KeyPair("server/server.crt", "server/server.key")
	if err != nil {
		log.Fatal(err)
	}

	certPool := x509.NewCertPool()
	ca, err := ioutil.ReadFile("server/ca.crt")
	if err != nil {
		log.Fatal(err)
	}
	if ok := certPool.AppendCertsFromPEM(ca); !ok {
		log.Fatal("failed to append certs")
	}

	creds := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{certificate},
		ClientAuth:   tls.RequireAndVerifyClientCert, // NOTE: this is optional!
		ClientCAs:    certPool,
	})
	var filter = func(ctx context.Context,
		req interface{}, info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (resp interface{}, err error) {
		log.Println("filter:", info)

		defer func() {
			if r := recover(); r != nil {
				err = fmt.Errorf("panic: %v", r)
			}
		}()

		return handler(ctx, req)
	}
	grpcServer := grpc.NewServer(
		grpc.Creds(creds),
		grpc.UnaryInterceptor(filter),
	)
	proto.RegisterHelloServiceServer(grpcServer, new(HelloServiceImpl))

	lis, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}
	grpcServer.Serve(lis)
}
