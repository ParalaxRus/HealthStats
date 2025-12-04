package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/paralaxrus/health-project/dbsvc/internal/gen/github.com/paralaxrus/health-project/dbsvc/proto"
	"github.com/paralaxrus/health-project/dbsvc/internal/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type userInput struct {
	name  *string
	email *string
	pass  *string
}

func (i *userInput) String() string {
	return fmt.Sprintf("Name: %s Email: %s Pass: %s", *i.name, *i.email, *i.pass)
}

func main() {
	time.Sleep(time.Second)

	input := userInput{}
	input.name = flag.String("name", "test", "user name")
	input.email = flag.String("email", "test@gmail.com", "user email")
	input.pass = flag.String("password", "testpass", "user password")
	flag.Parse()

	log.Println(input.String())

	host := getHealthStorageHost()
	log.Println(host)
	conn, err := grpc.NewClient(host, grpc.WithTransportCredentials(insecure.NewCredentials()))
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

func getHealthStorageHost() string {
	host := os.Getenv("GRPC_SERVER_ADDR")
	if len(host) == 0 {
		return "localhost:5005"
	}
	return host
}
