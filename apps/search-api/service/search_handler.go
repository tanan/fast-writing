package service

import "search-api/domain"

type SearchHandler interface {
	SaveContents(contents domain.Contents) (domain.ContentsId, error)
	FindContents(keyword string, limit int32, offset int32) ([]domain.ContentsScore, error)
	Version()
}
