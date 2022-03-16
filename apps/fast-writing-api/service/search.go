package service

import (
	"context"
	"errors"
	"fast-writing/pkg/pb"
	"fast-writing/pkg/pb/models"
	"google.golang.org/grpc"
)

type SearchService struct {
	SQLHandler SQLHandler
	Client     pb.SearchServiceClient
}

func NewSearchService(sqlHandler SQLHandler, serviceClient pb.SearchServiceClient) *SearchService {
	return &SearchService{
		SQLHandler: sqlHandler,
		Client:     serviceClient,
	}
}

func (s *SearchService) FindContentsIdListByTitle(ctx context.Context, req *models.TitleQueryParams) (*pb.ContentsScoreList, error) {
	list, err := s.Client.FindContentsIdListByTitle(context.Background(), req, grpc.MaxCallRecvMsgSize(10240))
	if err != nil {
		return nil, errors.New("cannot find contents Id list by title on search-api:" + err.Error())
	}
	return list, nil
}
