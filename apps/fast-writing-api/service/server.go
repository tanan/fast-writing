package service

import (
	"context"
	"errors"
	"fast-writing-api/domain"
	"fast-writing-api/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type UserService struct {
	SQLHandler SQLHandler
}

func NewUserService(s SQLHandler) *UserService {
	return &UserService{
		SQLHandler: s,
	}
}

func (s *UserService) GetUser(ctx context.Context, req *pb.UserId) (*pb.User, error) {
	user, err := s.SQLHandler.FindUserById(domain.UserId(req.Id))
	if err != nil {
		return nil, errors.New("cannot find user: " + err.Error())
	}
	return &pb.User{
		Id:          &pb.UserId{Id: string(user.Id)},
		Name:        user.Name,
		Email:       user.Email,
		LastUpdated: timestamppb.Now(),
	}, nil
}
