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

func (s *SearchService) SaveSearchContents(ctx context.Context, params *models.ContentsQueryParams) (*pb.ContentsScoreList, error) {
	s.SaveContents(domain.Contents{})
	return nil, nil
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
