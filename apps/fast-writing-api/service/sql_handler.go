package service

import "fast-writing-api/domain"

type SQLHandler interface {
	FindUserById(id domain.UserId) (domain.User, error)
	FindContentsById(id domain.ContentsId) (domain.Contents, error)
}
