package service

import (
	"context"
	"errors"
	"fast-writing-api/domain"
	"fast-writing-api/util"
	"fast-writing/pkg/pb"
	"fast-writing/pkg/pb/models"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"strconv"
)

type WritingService struct {
	SQLHandler SQLHandler
	Client     pb.SearchServiceClient
}

func NewWritingService(s SQLHandler, serviceClient pb.SearchServiceClient) *WritingService {
	return &WritingService{
		SQLHandler: s,
		Client:     serviceClient,
	}
}

func (s *WritingService) GetContents(ctx context.Context, req *models.ContentsId) (*models.Contents, error) {
	contents, err := s.SQLHandler.FindContentsById(domain.ContentsId(req.Id))
	if err != nil {
		return nil, errors.New("cannot find user: " + err.Error())
	}
	return s.toContents(contents), nil
}

func (s *WritingService) GetContentsList(ctx context.Context, req *models.ContentsQueryParams) (*models.ContentsList, error) {
	contentsList, err := s.SQLHandler.FindContentsList(req.GetParams().GetLimit(), req.GetParams().GetOffset())
	if err != nil {
		log.Printf("cannot find contents list: " + err.Error())
		return &models.ContentsList{}, err
	}
	return &models.ContentsList{
		ContentsList: s.toContentsList(contentsList),
	}, nil
}

func (s *WritingService) GetContentsListByTags(ctx context.Context, req *models.TagQueryParams) (*models.ContentsList, error) {
	contentsList, err := s.SQLHandler.FindContentsListByTags(req.GetTags(), req.GetParams().GetLimit(), req.GetParams().GetOffset())
	if err != nil {
		log.Printf("cannot find contents list: " + err.Error())
		return &models.ContentsList{}, err
	}
	return &models.ContentsList{
		ContentsList: s.toContentsList(contentsList),
	}, nil
}

func (s *WritingService) GetUserContentsList(ctx context.Context, req *models.UserContentsQueryParams) (*models.ContentsList, error) {
	userId, err := util.GetUserId(ctx)
	if err != nil {
		return &models.ContentsList{}, err
	}
	contentsList, err := s.SQLHandler.FindContentsListByUserId(domain.UserId(userId), req.GetParams().GetLimit(), req.GetParams().GetOffset())
	if err != nil {
		log.Printf("cannot find contents list: " + err.Error())
		return &models.ContentsList{}, err
	}
	return &models.ContentsList{
		ContentsList: s.toContentsList(contentsList),
	}, nil
}

func (s *WritingService) CreateUserContents(ctx context.Context, req *pb.CreateContentsRequest) (*pb.CreateContentsResponse, error) {
	contents := domain.Contents{
		Title:       req.Contents.Title,
		Description: req.Contents.Description,
		Scope:       req.Contents.Scope,
		Tags:        req.Contents.Tags,
	}
	if req.Contents.Id != nil {
		contents.ContentsId = domain.ContentsId(req.Contents.Id.Id)
		contents.QuizList = s.decodeQuizList(req.Contents.QuizList)
	}
	contents, err := s.SQLHandler.CreateContents(contents, domain.UserId(req.UserId.Id))
	if err != nil {
		return &pb.CreateContentsResponse{
			Created: false,
			Message: "failed to create contents",
		}, err
	}
	return &pb.CreateContentsResponse{
		Created:  true,
		Message:  "success",
		Contents: s.toContents(contents),
	}, nil
}

func (s *WritingService) DeleteUserContents(ctx context.Context, req *pb.DeleteContentsRequest) (*pb.DeleteResponse, error) {
	userId, err := util.GetUserId(ctx)
	if err != nil {
		return &pb.DeleteResponse{
			Deleted: false,
			Message: "failed to delete quiz",
		}, err
	}
	count, err := s.SQLHandler.DeleteContents(domain.UserId(userId), domain.ContentsId(req.ContentsId.Id))
	if count == 0 || err != nil {
		return &pb.DeleteResponse{
			Deleted: false,
			Message: "failed to delete contents",
		}, err
	}
	return &pb.DeleteResponse{
		Deleted: true,
		Message: "success",
	}, nil
}

func (s *WritingService) encodeQuizList(l []domain.Quiz) []*models.Quiz {
	var quizList []*models.Quiz
	for _, v := range l {
		order, _ := strconv.Atoi(v.Order)
		quizList = append(quizList, &models.Quiz{
			Question: v.Question,
			Answer:   v.Answer,
			Order:    int64(order),
		})
	}
	return quizList
}

func (s *WritingService) decodeQuizList(quiz []*models.Quiz) []domain.Quiz {
	var list []domain.Quiz
	for _, v := range quiz {
		list = append(list, domain.Quiz{
			Question: v.Question,
			Answer:   v.Answer,
			Order:    strconv.Itoa(int(v.Order)),
		})
	}
	return list
}

func (s *WritingService) toContentsList(contentsList []*domain.Contents) []*models.Contents {
	var m []*models.Contents
	for _, v := range contentsList {
		m = append(m, s.toContents(*v))
	}
	return m
}

func (s *WritingService) toContents(contents domain.Contents) *models.Contents {
	return &models.Contents{
		Id:          &models.ContentsId{Id: int64(contents.ContentsId)},
		Title:       contents.Title,
		Description: contents.Description,
		Creator:     contents.Creator,
		Scope:       contents.Scope,
		Tags:        contents.Tags,
		QuizList:    s.encodeQuizList(contents.QuizList),
		LastUpdated: timestamppb.New(contents.LastUpdated),
	}
}
