package service

import "fast-writing-api/domain"

type SQLHandler interface {
	FindUserById(id domain.UserId) (domain.User, error)
	FindContentsById(id domain.ContentsId) (domain.Contents, error)
	FindContentsList(limit int32, offset int32) ([]*domain.Contents, error)
	CreateContents(contents domain.Contents, userId domain.UserId) (int64, error)
	CreateQuiz(quiz domain.Quiz) (int64, error)
}
