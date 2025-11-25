package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/paralaxrus/health-project/dbsvc/github.com/paralaxrus/health-project/dbsvc/proto"
	"github.com/paralaxrus/health-project/dbsvc/internal/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const healthDatabaseSvcPort = "5005"

type userInput struct {
	name  *string
	email *string
	pass  *string
}

func main() {
	time.Sleep(time.Second)

	input := userInput{}
	input.name = flag.String("name", "test", "user name")
	input.email = flag.String("email", "test@gmail.com", "user email")
	input.pass = flag.String("password", "testpass", "user password")
	flag.Parse()

	address := fmt.Sprintf("%s:%s", host(), healthDatabaseSvcPort)
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := proto.NewUserServiceClient(conn)

	resp, err := client.Register(context.Background(),
		&proto.RegisterRequest{Name: *input.name, Email: *input.email, Password: *input.pass})
	if err != nil {
		log.Fatal(err)
	}

	respAsStr, err := utils.ToPrettyString(resp)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Response: %s", respAsStr)

	log.Printf("Health dbsvc client completed")
}

func host() string {
	host := os.Getenv("GRPC_SERVER_ADDR")
	if len(host) == 0 {
		return "localhost"
	}
	return host
}
