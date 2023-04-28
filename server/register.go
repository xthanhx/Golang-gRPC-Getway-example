package server

import (
	"context"
	helloworldpb "github.com/myuser/myrepo/proto/helloworld"
	"google.golang.org/grpc"
	"log"
)

type server struct {
	helloworldpb.UnimplementedGreeterServer
}

func NewServer() *server {
	return &server{}
}

func (s *server) SayHello(ctx context.Context, req *helloworldpb.HelloRequest) (*helloworldpb.HelloReply, error) {
	log.Println(req.Name)
	return &helloworldpb.HelloReply{Message: req.Name + " world"}, nil
}

func (s *server) Lala(ctx context.Context, req *helloworldpb.HelloRequest) (*helloworldpb.HelloReply, error) {
	return &helloworldpb.HelloReply{Message: req.Name + " world"}, nil
}

func Register(s *grpc.Server) {
	helloworldpb.RegisterGreeterServer(s, &server{})
}
