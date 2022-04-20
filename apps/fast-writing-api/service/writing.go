package service

import (
	"context"
	"errors"
	"fast-writing-api/domain"
	"fast-writing/pkg/pb"
	"fast-writing/pkg/pb/models"
	"google.golang.org/grpc/metadata"
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

func (s *WritingService) getUserId(ctx context.Context) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	userId := md.Get("user-id")
	if !ok || len(userId) == 0 {
		return "", errors.New("failed to get user-id on metadata")
	}
	return userId[0], nil
}

func (s *WritingService) GetContents(ctx context.Context, req *models.ContentsId) (*models.Contents, error) {
	contents, err := s.SQLHandler.FindContentsById(domain.ContentsId(req.Id))
	if err != nil {
		return nil, errors.New("cannot find user: " + err.Error())
	}
	return &models.Contents{
		Id:          &models.ContentsId{Id: int64(contents.ContentsId)},
		Title:       contents.Title,
		Description: contents.Description,
		Creator:     contents.Creator,
		QuizList:    s.encodeQuizList(contents.QuizList),
		LastUpdated: timestamppb.New(contents.LastUpdated),
	}, nil
}

func (s *WritingService) GetContentsList(ctx context.Context, req *models.ContentsQueryParams) (*models.ContentsList, error) {
	contentsList, err := s.SQLHandler.FindContentsList(req.GetParams().GetLimit(), req.GetParams().GetOffset())
	if err != nil {
		log.Printf("cannot find contents list: " + err.Error())
		return &models.ContentsList{}, err
	}
	var m []*models.Contents
	for _, v := range contentsList {
		m = append(m, &models.Contents{
			Id:          &models.ContentsId{Id: int64(v.ContentsId)},
			Title:       v.Title,
			Description: v.Description,
			Creator:     v.Creator,
			Scope:       v.Scope,
			LastUpdated: timestamppb.New(v.LastUpdated),
		})
	}
	return &models.ContentsList{
		ContentsList: m,
	}, nil
}

func (s *WritingService) GetUserContentsList(ctx context.Context, req *models.UserContentsQueryParams) (*models.ContentsList, error) {
	userId, err := s.getUserId(ctx)
	if err != nil {
		return &models.ContentsList{}, err
	}
	contentsList, err := s.SQLHandler.FindContentsListByUserId(domain.UserId(userId), req.GetParams().GetLimit(), req.GetParams().GetOffset())
	if err != nil {
		log.Printf("cannot find contents list: " + err.Error())
		return &models.ContentsList{}, err
	}
	var m []*models.Contents
	for _, v := range contentsList {
		m = append(m, &models.Contents{
			Id:          &models.ContentsId{Id: int64(v.ContentsId)},
			Title:       v.Title,
			Description: v.Description,
			Creator:     v.Creator,
			Scope:       v.Scope,
			LastUpdated: timestamppb.New(v.LastUpdated),
		})
	}
	return &models.ContentsList{
		ContentsList: m,
	}, nil
}

func (s *WritingService) CreateUserContents(ctx context.Context, req *pb.CreateContentsRequest) (*pb.CreateContentsResponse, error) {
	contents := domain.Contents{
		Title:       req.Contents.Title,
		Description: req.Contents.Description,
		Scope:       req.Contents.Scope,
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
		Created: true,
		Message: "success",
		Contents: &models.Contents{
			Id: &models.ContentsId{
				Id: int64(contents.ContentsId),
			},
			Title:       contents.Title,
			Description: contents.Description,
			Creator:     contents.Creator,
			Scope:       contents.Scope,
			QuizList:    s.encodeQuizList(contents.QuizList),
			LastUpdated: timestamppb.New(contents.LastUpdated),
		},
	}, nil
}

func (s *WritingService) DeleteUserContents(ctx context.Context, req *pb.DeleteContentsRequest) (*pb.DeleteResponse, error) {
	userId, err := s.getUserId(ctx)
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
