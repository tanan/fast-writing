package service

import "fast-writing-api/domain"

type SQLHandler interface {
	FindUserById(id domain.UserId) (domain.User, error)
	FindContentsById(id domain.ContentsId) (domain.Contents, error)
	FindContentsList(limit int32, offset int32) ([]*domain.Contents, error)
	CreateContents(contents domain.Contents, userId domain.UserId) (int, error)
	CreateQuiz(contentsId domain.ContentsId, quiz domain.Quiz) (int, error)
}
