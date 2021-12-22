package main

import (
	"context"
	"google.golang.org/grpc"
	"grpc/proto"
	"log"
)

func main() {
	conn, err := grpc.Dial("localhost:1234", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := proto.NewPubsubServiceClient(conn)

	_, err = client.Publish(context.Background(), &proto.String{Value: "golang: hello go"})
	if err != nil {
		log.Fatal(err)
	}
	_, err = client.Publish(context.Background(), &proto.String{Value: "docker: hello docker"})
	if err != nil {
		log.Fatal(err)
	}
}
