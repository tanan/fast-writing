package service

import (
	"context"
	"errors"
	"fast-writing-api/domain"
	"fast-writing/pkg/pb"
	"fast-writing/pkg/pb/models"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
)

type WritingService struct {
	SQLHandler SQLHandler
	Client     pb.SearchServiceClient
}

func NewWritingService(s SQLHandler, serviceClient pb.SearchServiceClient) *WritingService {
	return &WritingService{
		SQLHandler: s,
		Client:     serviceClient,
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
		Description: contents.Description,
		Creator:     contents.Creator,
		QuizList:    s.encodeQuizList(contents.QuizList),
		LastUpdated: timestamppb.New(contents.LastUpdated),
	}, nil
}

func (s *WritingService) encodeQuizList(l []domain.Quiz) []*models.Quiz {
	var quizList []*models.Quiz
	for _, v := range l {
		quizList = append(quizList, &models.Quiz{
			Id: &models.QuizId{
				Id: int64(v.Id),
			},
			Question: v.Question,
			Answer:   v.Answer,
		})
	}
	return quizList
}

func (s *WritingService) decodeQuizList(quiz []*models.Quiz, contentsId *models.ContentsId) []domain.Quiz {
	var list []domain.Quiz
	for _, v := range quiz {
		list = append(list, domain.Quiz{
			Question:   v.Question,
			Answer:     v.Answer,
			ContentsId: domain.ContentsId(contentsId.Id),
		})
	}
	return list
}

func (s *WritingService) GetContentsList(ctx context.Context, req *models.ContentsQueryParams) (*models.ContentsList, error) {
	contentsList, err := s.SQLHandler.FindContentsList(req.GetParams().GetLimit(), req.GetParams().GetOffset())
	if err != nil {
		log.Printf("cannot find contents list: " + err.Error())
		return &models.ContentsList{}, nil
	}
	var m []*models.Contents
	for _, v := range contentsList {
		m = append(m, &models.Contents{
			Id:          &models.ContentsId{Id: int64(v.ContentsId)},
			Title:       v.Title,
			Description: v.Description,
			Creator:     v.Creator,
			Scope:       v.Scope,
			LastUpdated: timestamppb.New(v.LastUpdated),
		})
	}
	return &models.ContentsList{
		ContentsList: m,
	}, nil
}

func (s *WritingService) GetUserContentsList(ctx context.Context, req *models.UserContentsQueryParams) (*models.ContentsList, error) {
	contentsList, err := s.SQLHandler.FindContentsListByUserId(domain.UserId(req.GetId().Id), req.GetParams().GetLimit(), req.GetParams().GetOffset())
	if err != nil {
		log.Printf("cannot find contents list: " + err.Error())
		return &models.ContentsList{}, nil
	}
	var m []*models.Contents
	for _, v := range contentsList {
		m = append(m, &models.Contents{
			Id:          &models.ContentsId{Id: int64(v.ContentsId)},
			Title:       v.Title,
			Description: v.Description,
			Creator:     v.Creator,
			Scope:       v.Scope,
			LastUpdated: timestamppb.New(v.LastUpdated),
		})
	}
	return &models.ContentsList{
		ContentsList: m,
	}, nil
}

func (s *WritingService) CreateUserContents(ctx context.Context, req *pb.CreateContentsRequest) (*pb.CreateContentsResponse, error) {
	contents := domain.Contents{
		Title:       req.Contents.Title,
		Description: req.Contents.Description,
		Scope:       req.Contents.Scope,
	}
	if req.Contents.Id != nil {
		contents.ContentsId = domain.ContentsId(req.Contents.Id.Id)
		contents.QuizList = s.decodeQuizList(req.Contents.QuizList, req.Contents.Id)
	}
	contents, err := s.SQLHandler.CreateContents(contents, domain.UserId(req.UserId.Id))
	if err != nil {
		return &pb.CreateContentsResponse{
			Created: false,
			Message: "failed to create contents",
		}, err
	}
	//_, err = s.Client.SaveSearchContents(context.Background(), req.Contents, grpc.MaxCallRecvMsgSize(10240))
	//if err != nil {
	//	return &pb.CreateContentsResponse{
	//		Created: false,
	//		Message: "failed to create contents",
	//	}, errors.New("failed to create contents on search-api:" + err.Error())
	//}
	return &pb.CreateContentsResponse{
		Created: true,
		Message: "success",
		Contents: &models.Contents{
			Id: &models.ContentsId{
				Id: int64(contents.ContentsId),
			},
			Title:       contents.Title,
			Description: contents.Description,
			Creator:     contents.Creator,
			Scope:       contents.Scope,
			QuizList:    s.encodeQuizList(contents.QuizList),
			LastUpdated: timestamppb.New(contents.LastUpdated),
		},
	}, nil
}

func (s *WritingService) CreateUserQuiz(ctx context.Context, req *pb.CreateQuizRequest) (*pb.CreateQuizResponse, error) {
	quiz := domain.Quiz{
		Question:   req.Quiz.Question,
		Answer:     req.Quiz.Answer,
		ContentsId: domain.ContentsId(req.ContentsId.Id),
	}
	if req.Quiz.Id != nil {
		quiz.Id = domain.QuizId(req.Quiz.Id.Id)
	}
	quizId, err := s.SQLHandler.CreateQuiz(quiz)
	if err != nil {
		return &pb.CreateQuizResponse{
			Created: false,
			Message: "failed to create contents",
		}, err
	}
	return &pb.CreateQuizResponse{
		Created: true,
		Message: "success",
		QuizId: &models.QuizId{
			Id: quizId,
		},
	}, nil
}

func (s *WritingService) DeleteUserContents(ctx context.Context, req *pb.DeleteContentsRequest) (*pb.DeleteResponse, error) {
	userId := ctx.Value("user_id")
	if userId == nil {
		return &pb.DeleteResponse{
			Deleted: false,
			Message: "failed to delete quiz",
		}, errors.New("failed to get user_id. please check authentication")
	}
	count, err := s.SQLHandler.DeleteContents(domain.UserId(userId.(string)), domain.ContentsId(req.ContentsId.Id))
	if count == 0 || err != nil {
		return &pb.DeleteResponse{
			Deleted: false,
			Message: "failed to delete contents",
		}, err
	}
	return &pb.DeleteResponse{
		Deleted: true,
		Message: "success",
	}, nil
}
func (s *WritingService) DeleteUserQuiz(ctx context.Context, req *pb.DeleteQuizRequest) (*pb.DeleteResponse, error) {
	userId := ctx.Value("user_id")
	if userId == nil {
		return &pb.DeleteResponse{
			Deleted: false,
			Message: "failed to delete quiz",
		}, errors.New("failed to get user_id. please check authentication")
	}
	count, err := s.SQLHandler.DeleteQuiz(domain.UserId(userId.(string)), domain.ContentsId(req.ContentsId.Id), domain.QuizId(req.QuizId.Id))
	if count == 0 || err != nil {
		return &pb.DeleteResponse{
			Deleted: false,
			Message: "failed to delete quiz",
		}, err
	}
	return &pb.DeleteResponse{
		Deleted: true,
		Message: "success",
	}, nil
}
