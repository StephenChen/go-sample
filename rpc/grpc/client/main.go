package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"grpc/auth"
	"grpc/proto"
	"io"
	"io/ioutil"
	"log"
	"net"
	"time"
)

type RestServiceImpl struct {
	*proto.UnimplementedRestServiceServer
}

func (r *RestServiceImpl) Get(ctx context.Context, message *proto.String) (*proto.String, error) {
	return &proto.String{Value: "Get hi:" + message.Value + "#"}, nil
}

func (r *RestServiceImpl) Post(ctx context.Context, message *proto.String) (*proto.String, error) {
	return &proto.String{Value: "Post hi:" + message.Value + "@"}, nil
}

func main() {
	go func() {
		grpcServer := grpc.NewServer()
		proto.RegisterRestServiceServer(grpcServer, new(RestServiceImpl))
		lis, _ := net.Listen("tcp", ":5000")
		grpcServer.Serve(lis)
	}()
	//select {}

	certificate, err := tls.LoadX509KeyPair("client/client.crt", "client/client.key")
	if err != nil {
		log.Fatal(err)
	}

	certPool := x509.NewCertPool()
	ca, err := ioutil.ReadFile("server/ca.crt")
	if err != nil {
		log.Fatal(err)
	}
	if ok := certPool.AppendCertsFromPEM(ca); !ok {
		log.Fatal("failed to append ca certs")
	}

	creds := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{certificate},
		ServerName:   "server.io",
		RootCAs:      certPool,
	})

	auth := auth.Authentication{
		User:     "gopher",
		Password: "password",
	}

	conn, err := grpc.Dial(
		"localhost:1234",
		grpc.WithTransportCredentials(creds),
		grpc.WithPerRPCCredentials(&auth),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := proto.NewHelloServiceClient(conn)
	reply, err := client.Hello(context.Background(), &proto.String{Value: "hello"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reply.GetValue())
	stream, err := client.Channel(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		for {
			if err := stream.Send(&proto.String{Value: "hi"}); err != nil {
				log.Fatal(err)
			}
			time.Sleep(time.Second)
		}
	}()

	for {
		reply, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		fmt.Println(reply.GetValue())
	}
}
