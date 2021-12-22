package main

import (
	"context"
	"github.com/moby/moby/pkg/pubsub"
	"google.golang.org/grpc"
	"grpc/proto"
	"log"
	"net"
	"strings"
	"time"
)

type PubsubService struct {
	pub *pubsub.Publisher
	proto.UnimplementedPubsubServiceServer
}

func NewPubsubService() *PubsubService {
	return &PubsubService{
		pub: pubsub.NewPublisher(100*time.Millisecond, 10),
	}
}

func (p *PubsubService) Publish(ctx context.Context, arg *proto.String) (*proto.String, error) {
	p.pub.Publish(arg.GetValue())
	return &proto.String{}, nil
}

func (p *PubsubService) Subscribe(arg *proto.String, stream proto.PubsubService_SubscribeServer) error {
	ch := p.pub.SubscribeTopic(func(v interface{}) bool {
		if key, ok := v.(string); ok {
			if strings.HasPrefix(key, arg.GetValue()) {
				return true
			}
		}
		return false
	})

	for v := range ch {
		if err := stream.Send(&proto.String{Value: v.(string)}); err != nil {
			return err
		}
	}

	return nil
}

func main() {
	grpcServer := grpc.NewServer()
	proto.RegisterPubsubServiceServer(grpcServer, NewPubsubService())

	lis, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}
	grpcServer.Serve(lis)
}
