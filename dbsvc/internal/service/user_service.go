package service

import (
	"context"
	"fmt"

	"github.com/paralaxrus/health-project/dbsvc/github.com/paralaxrus/health-project/dbsvc/proto"
	"github.com/paralaxrus/health-project/dbsvc/internal/storage"
	"github.com/paralaxrus/health-project/dbsvc/internal/storage/model"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type userServiceHandler struct {
	proto.UnimplementedUserServiceServer

	store *storage.UserDataSource
}

func (s *userServiceHandler) Open() error {
	return s.store.Connect()
}

func (s *userServiceHandler) Close() {
	s.store.Disconnect()
}

func (s *userServiceHandler) Register(ctx context.Context, req *proto.RegisterRequest) (*proto.RegisterResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("nil request")
	}
	id, err := s.store.CreateUser(ctx, req.GetName(), req.GetEmail(), req.GetPassword())
	if err != nil {
		return &proto.RegisterResponse{ErrMsg: &proto.ErrorMsg{Description: err.Error()}}, err
	}

	return &proto.RegisterResponse{Id: id}, nil
}

func (s *userServiceHandler) Find(ctx context.Context, req *proto.FindRequest) (*proto.FindResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("nil request")
	}
	user, err := s.store.FindUser(ctx, toUserIndex(req))
	if err != nil {
		return &proto.FindResponse{ErrMsg: &proto.ErrorMsg{Description: err.Error()}}, err
	}

	return &proto.FindResponse{User: toProtoUser(user)}, nil
}

func NewUserServiceHandler() *userServiceHandler {
	return &userServiceHandler{store: storage.NewUserDataSource()}
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
