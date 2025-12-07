package main

import (
	"log"
	"net"

	"github.com/paralaxrus/health-project/dbsvc/internal/gen/github.com/paralaxrus/health-project/dbsvc/proto"
	"github.com/paralaxrus/health-project/dbsvc/internal/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const grpcPort = ":5005"

func main() {
	listener, err := net.Listen("tcp", grpcPort)
	if err != nil {
		log.Fatalf("failed to listen: %v on port %v", err, grpcPort)
	}
	log.Printf("user service is listening on %v", grpcPort)

	service := service.NewUserService()
	if err = service.Open(); err != nil {
		log.Fatalf("%s", err.Error())
	}
	defer service.Close()

	grpcServer := grpc.NewServer()
	proto.RegisterUserServiceServer(grpcServer, service)
	reflection.Register(grpcServer)

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
