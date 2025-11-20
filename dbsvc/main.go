package main

import (
	"context"
	"log"
	"net"

	"github.com/paralaxrus/health-project/dbsvc/github.com/paralaxrus/health-project/dbsvc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const grpcPort = ":5005"

type server struct {
	proto.UnimplementedUserServiceServer
}

func (s *server) Register(ctx context.Context, req *proto.RegisterRequest) (*proto.RegisterResponse, error) {
	return nil, nil
}

func (s *server) Find(ctx context.Context, req *proto.FindRequest) (*proto.FindResponse, error) {
	return nil, nil
}

func main() {
	listener, err := net.Listen("tcp", grpcPort)
	if err != nil {
		log.Fatalf("failed to listen: %v on port %v", err, grpcPort)
	}

	grpcServer := grpc.NewServer()
	proto.RegisterUserServiceServer(grpcServer, &server{})
	reflection.Register(grpcServer)

	log.Printf("health state service is listening on %v", grpcPort)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
