package service

import (
	"context"
	"fast-writing/pkg/pb"
	"fast-writing/pkg/pb/models"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type SearchService struct {
}

func NewSearchService() *SearchService {
	return &SearchService{}
}

func (s *SearchService) FindContentsIdListByTitle(ctx context.Context, params *models.TitleQueryParams) (*pb.ContentsScoreList, error) {
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
