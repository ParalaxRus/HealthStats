package service

import (
	"context"
	"fmt"
	"log"

	"github.com/paralaxrus/health-project/dbsvc/github.com/paralaxrus/health-project/dbsvc/proto"
	"github.com/paralaxrus/health-project/dbsvc/internal/storage"
	"github.com/paralaxrus/health-project/dbsvc/internal/storage/model"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type userServiceHandler struct {
	proto.UnimplementedUserServiceServer

	dataSource *storage.UserDataSource
}

func (s *userServiceHandler) Open() error {
	return s.dataSource.Connect()
}

func (s *userServiceHandler) Close() {
	s.dataSource.Disconnect()
}

func (s *userServiceHandler) Register(ctx context.Context, req *proto.RegisterRequest) (*proto.RegisterResponse, error) {
	log.Println("register grpc API is called")

	if req == nil {
		return nil, fmt.Errorf("nil request")
	}
	id, err := s.dataSource.CreateUser(ctx, req.GetName(), req.GetEmail(), req.GetPassword())
	if err != nil {
		return &proto.RegisterResponse{ErrMsg: &proto.ErrorMsg{Description: err.Error()}}, err
	}

	return &proto.RegisterResponse{Id: int64(id)}, nil
}

func (s *userServiceHandler) Find(ctx context.Context, req *proto.FindRequest) (*proto.FindResponse, error) {
	log.Println("find grpc API is called")
	if req == nil {
		return nil, fmt.Errorf("nil request")
	}
	user, err := s.dataSource.FindUser(ctx, toUserIndex(req))
	if err != nil {
		return &proto.FindResponse{ErrMsg: &proto.ErrorMsg{Description: err.Error()}}, err
	}

	return &proto.FindResponse{User: toProtoUser(user)}, nil
}

func NewUserServiceHandler() *userServiceHandler {
	return &userServiceHandler{dataSource: storage.NewUserDataSource()}
}

func toUserIndex(req *proto.FindRequest) storage.Index {
	return storage.NewIndex(req.GetId(), req.GetEmail())
}

func toProtoUser(user *model.User) *proto.User {
	return &proto.User{Name: user.Name,
		Email:        user.Email,
		Password:     user.Password,
		RegisteredAt: timestamppb.New(user.Created),
		LastLoginAt:  timestamppb.New(user.Created)}
}
