package service

import (
	"context"
	"fast-writing-api/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type UserService struct {

}


func NewUserService() *UserService {
	return &UserService{}
}

func (s *UserService) GetUser(ctx context.Context, req *pb.UserId) (*pb.User, error) {
	return &pb.User{
		Id:          &pb.UserId{Id: "test-id"},
		Name:        "testname",
		Email:       "testname@example.com",
		LastUpdated: timestamppb.Now(),
	}, nil
}