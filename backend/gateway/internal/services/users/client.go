package users

import (
	"gateway/internal/gen/github.com/paralaxrus/health-project/dbsvc/proto"
	"log"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type UsersClient struct {
	connection *grpc.ClientConn
	client     proto.UserServiceClient
}

func (u *UsersClient) Close() {
	if u.connection != nil {
		u.connection.Close()
	}
}

func NewUsersClient() (*UsersClient, error) {
	host := getUsersServiceHost()
	log.Println(host)
	conn, err := grpc.NewClient(host, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	client := proto.NewUserServiceClient(conn)
	return &UsersClient{connection: conn, client: client}, nil
}

func getUsersServiceHost() string {
	host := os.Getenv("GRPC_SERVER_ADDR")
	if len(host) == 0 {
		return "localhost:5005"
	}
	return host
}
