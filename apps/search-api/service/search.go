package service

import (
	"context"
	"fast-writing/pkg/pb"
	"fast-writing/pkg/pb/models"
	"log"
)

type SearchService struct {

}

func NewSearchService() *SearchService {
	return &SearchService{}
}

func (s *SearchService) FindContentsIdListByTitle(ctx context.Context, params *models.TitleQueryParams) (*pb.ContentsScoreList, error) {
	log.Println("--> findContentsIdListByTitle")
	return nil, nil
}