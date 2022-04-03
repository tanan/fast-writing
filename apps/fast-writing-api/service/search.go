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
		return &pb.ContentsScoreList{
			ContentsScore: nil,
		}, errors.New("cannot find contents Id list by title on search-api:" + err.Error())
	}
	return list, nil
}

func (s *SearchService) SaveSearchContents(ctx context.Context, req *models.Contents) (*pb.CreateSearchResponse, error) {
	res, err := s.Client.SaveSearchContents(context.Background(), req, grpc.MaxCallRecvMsgSize(10240))
	if err != nil {
		return &pb.CreateSearchResponse{
			Created: false,
			Message: "error",
		}, errors.New("cannot create contents on search-api:" + err.Error())
	}
	return res, nil
}
