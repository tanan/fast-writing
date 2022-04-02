package service

import "search-api/domain"

type SearchHandler interface {
	SaveContents(contents domain.Contents) (string, error)
	FindContents(keyword string, limit int32, offset int32) ([]domain.ContentsScore, error)
	Version()
}
