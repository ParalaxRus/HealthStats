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

	handler := service.NewUserServiceHandler()
	defer handler.Close()
	if err = handler.Open(); err != nil {
		log.Fatalf("%s", err.Error())
	}

	grpcServer := grpc.NewServer()
	proto.RegisterUserServiceServer(grpcServer, handler)
	reflection.Register(grpcServer)

	log.Printf("health state service is listening on %v", grpcPort)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
