package service

import (
	"context"
	"fast-writing/pkg/pb"
	"fast-writing/pkg/pb/models"
	"search-api/database"
	"search-api/domain"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type SearchService struct {
	SearchHandler
}

func NewSearchService() *SearchService {
	return &SearchService{
		database.NewElasticsearch(),
	}
}

func (s *SearchService) SaveSearchContents(ctx context.Context, params *models.Contents) (*pb.CreateSearchResponse, error) {
	var quiz string
	var answer string
	for i, v := range params.QuizList {
		if i == 0 {
			quiz = v.Question
			answer = v.Answer
			continue
		}
		quiz += " " + v.Question
		answer += " " + v.Answer
	}
	_, err := s.SaveContents(domain.Contents{
		Id:       domain.ContentsId(params.Id.Id),
		Title:    params.Title,
		Username: params.Creator,
		Quiz:     quiz,
		Answer:   answer,
	})
	if err != nil {
		return &pb.CreateSearchResponse{}, err
	}
	return &pb.CreateSearchResponse{
		Created: true,
		Message: "success",
	}, nil
}

func (s *SearchService) FindContentsIdListByTitle(ctx context.Context, params *models.TitleQueryParams) (*pb.ContentsScoreList, error) {
	contents, err := s.FindContents(params.Title, params.Params.Limit, params.Params.Offset)
	if err != nil {
		return &pb.ContentsScoreList{}, err
	}
	return s.toContentsScoreList(contents), nil
}

func (s *SearchService) toContentsScoreList(contents []domain.ContentsScore) *pb.ContentsScoreList {
	var list []*pb.ContentsScore
	for _, v := range contents {
		list = append(list, &pb.ContentsScore{
			ContentsId: &models.ContentsId{
				Id: int64(v.Id),
			},
			//Score:       v.Score,
			LastUpdated: timestamppb.New(v.LastUpdated),
		})
	}
	return &pb.ContentsScoreList{
		ContentsScore: list,
	}
}
