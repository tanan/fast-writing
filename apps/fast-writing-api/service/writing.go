package service

import (
	"context"
	"errors"
	"fast-writing-api/domain"
	"fast-writing-api/pb"
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

func (s *WritingService) GetContents(ctx context.Context, req *pb.ContentsId) (*pb.Contents, error) {
	contents, err := s.SQLHandler.FindContentsById(domain.ContentsId(req.Id))
	if err != nil {
		return nil, errors.New("cannot find user: " + err.Error())
	}
	return &pb.Contents{
		Id:          &pb.ContentsId{Id: int64(contents.ContentsId)},
		Title:       contents.Title,
		Creator:     contents.Creator,
		QuizList:    s.toQuizList(contents.QuizList),
		LastUpdated: timestamppb.New(contents.LastUpdated),
	}, nil
}

func (s *WritingService) toQuizList(l []domain.Quiz) []*pb.Quiz {
	var quizList []*pb.Quiz
	for _, v := range l {
		quizList = append(quizList, &pb.Quiz{
			Id:       v.Id,
			Question: v.Question,
			Answer:   v.Answer,
		})
	}
	return quizList
}

func (s *WritingService) GetContentsList(ctx context.Context, req *pb.ContentsQueryParams) (*pb.ContentsList, error) {
	return nil, nil
}

func (s *WritingService) GetUserContentsList(ctx context.Context, req *pb.UserContentsQueryParams) (*pb.ContentsList, error) {
	return nil, nil
}
