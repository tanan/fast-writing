package service

import (
	"context"
	"errors"
	"fast-writing-api/domain"
	"fast-writing/pkg/pb/models"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type WritingService struct {
	SQLHandler SQLHandler
}

func NewWritingService(s SQLHandler) *WritingService {
	return &WritingService{
		SQLHandler: s,
	}
}

func (s *WritingService) GetContents(ctx context.Context, req *models.ContentsId) (*models.Contents, error) {
	contents, err := s.SQLHandler.FindContentsById(domain.ContentsId(req.Id))
	if err != nil {
		return nil, errors.New("cannot find user: " + err.Error())
	}
	return &models.Contents{
		Id:          &models.ContentsId{Id: int64(contents.ContentsId)},
		Title:       contents.Title,
		Creator:     contents.Creator,
		QuizList:    s.toQuizList(contents.QuizList),
		LastUpdated: timestamppb.New(contents.LastUpdated),
	}, nil
}

func (s *WritingService) toQuizList(l []domain.Quiz) []*models.Quiz {
	var quizList []*models.Quiz
	for _, v := range l {
		quizList = append(quizList, &models.Quiz{
			Id:       v.Id,
			Question: v.Question,
			Answer:   v.Answer,
		})
	}
	return quizList
}

func (s *WritingService) GetContentsList(ctx context.Context, req *models.ContentsQueryParams) (*models.ContentsList, error) {
	return nil, nil
}

func (s *WritingService) GetUserContentsList(ctx context.Context, req *models.UserContentsQueryParams) (*models.ContentsList, error) {
	return nil, nil
}
