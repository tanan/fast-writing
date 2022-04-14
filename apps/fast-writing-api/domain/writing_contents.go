package domain

import "time"

type ContentsId int64
type QuizId int64

type Contents struct {
	ContentsId  ContentsId
	Creator     string
	Title       string
	Description string
	Scope       string
	QuizList    []Quiz
	LastUpdated time.Time
}

type Quiz struct {
	Id         QuizId
	Question   string
	Answer     string
	ContentsId ContentsId
}
