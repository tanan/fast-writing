package service

import (
	"context"
	"errors"
	"fast-writing-api/domain"
	"fast-writing-api/util"
	"fast-writing/pkg/pb/models"
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

func (s *UserService) UpdateUser(ctx context.Context, req *models.User) (*models.UserId, error) {
	userId, err := util.GetUserId(ctx)
	if err != nil {
		return &models.UserId{}, err
	}
	user := domain.User{
		Id:    domain.UserId(userId),
		Name:  req.Name,
		Email: req.Email,
	}
	id, err := s.SQLHandler.UpdateUser(user)
	if err != nil {
		return &models.UserId{}, err
	}
	return &models.UserId{Id: string(id)}, nil
}

func (s *UserService) CreateUser(ctx context.Context, req *models.User) (*models.UserId, error) {
	userId, err := util.GetUserId(ctx)
	if err != nil {
		return &models.UserId{}, err
	}
	user := domain.User{
		Id:    domain.UserId(userId),
		Name:  req.Name,
		Email: req.Email,
	}
	id, err := s.SQLHandler.CreateUser(user)
	if err != nil {
		return &models.UserId{}, err
	}
	return &models.UserId{Id: string(id)}, nil
}

func (s *UserService) GetUser(ctx context.Context, req *models.UserId) (*models.User, error) {
	user, err := s.SQLHandler.FindUserById(domain.UserId(req.Id))
	if err != nil {
		return nil, errors.New("cannot find user: " + err.Error())
	}
	return &models.User{
		Id:          &models.UserId{Id: string(user.Id)},
		Name:        user.Name,
		Email:       user.Email,
		LastUpdated: timestamppb.Now(),
	}, nil
}
