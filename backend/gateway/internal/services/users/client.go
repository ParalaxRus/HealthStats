package users

import (
	"context"
	"errors"
	"gateway/internal/gen/github.com/paralaxrus/health-project/dbsvc/proto"
	"log"
	"os"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type User struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`

	// Move to a child struct
	RegisteredAt time.Time
	LastLoginAt  time.Time
}

func NewUser(name string, email string, password string) *User {
	return &User{Name: name, Email: email, Password: password}
}

type UserService interface {
	Create(ctx context.Context, user User) (int64, error)
	Find(ctx context.Context, email string) (*User, error)
	List(ctx context.Context, pageToken string) ([]*User, error)

	Close()
}

type userService struct {
	connection *grpc.ClientConn
	client     proto.UserServiceClient
}

func NewUserServiceClient() (UserService, error) {
	host := getUserServiceHost()
	log.Println(host)
	conn, err := grpc.NewClient(host, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	client := proto.NewUserServiceClient(conn)
	return &userService{connection: conn, client: client}, nil
}

func (u *userService) Create(ctx context.Context, user User) (int64, error) {
	req := proto.RegisterRequest{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}
	resp, err := u.client.Register(ctx, &req)
	if err != nil {
		return -1, err
	}
	return resp.Id, nil
}

func (u *userService) Find(ctx context.Context, email string) (*User, error) {
	req := proto.FindRequest{
		Kind: &proto.FindRequest_Email{Email: email},
	}
	resp, err := u.client.Find(ctx, &req)
	if err != nil {
		return nil, err
	}
	return toUser(resp.GetUser())
}

func (u *userService) List(ctx context.Context, pageToken string) ([]*User, error) {
	return []*User{}, nil
}

func (u *userService) Close() {
	if u.connection != nil {
		u.connection.Close()
	}
}

func getUserServiceHost() string {
	host := os.Getenv("USER_SERVICE_HOST")
	if len(host) == 0 {
		return "localhost:5005"
	}
	return host
}

func toUser(respUser *proto.User) (*User, error) {
	if respUser == nil {
		return nil, errors.New("response returned nil user")
	}

	return &User{
		Name:         respUser.GetName(),
		Email:        respUser.GetEmail(),
		Password:     respUser.GetPassword(),
		RegisteredAt: respUser.GetRegisteredAt().AsTime(),
		LastLoginAt:  respUser.GetLastLoginAt().AsTime(),
	}, nil
}
