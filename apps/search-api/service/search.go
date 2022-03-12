package service

import (
	"context"
	"fast-writing/pkg/pb"
	"fast-writing/pkg/pb/models"
)

type SearchService struct {

}

func NewSearchService() *SearchService {
	return &SearchService{}
}

func (s *SearchService) FindContentsIdListByTitle(context.Context, *models.TitleQueryParams) (*pb.ContentsScoreList, error) {
	return nil, nil
}