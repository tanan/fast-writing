package service

import (
	"context"
	"fast-writing/pkg/pb"
	"fast-writing/pkg/pb/models"
	"search-api/database"
	"time"

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

func (s *SearchService) FindContentsIdListByTitle(ctx context.Context, params *models.TitleQueryParams) (*pb.ContentsScoreList, error) {
	s.FindContents(params.Title, params.Params.Limit, params.Params.Offset)
	return &pb.ContentsScoreList{
		ContentsScore: []*pb.ContentsScore{
			{
				ContentsId: &models.ContentsId{
					Id: 1,
				},
				Score:       0,
				LastUpdated: timestamppb.New(time.Now()),
			},
		},
	}, nil
}
