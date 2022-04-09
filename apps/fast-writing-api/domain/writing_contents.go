package domain

import "time"

type ContentsId int64

type Contents struct {
	ContentsId  ContentsId
	Creator     string
	Title       string
	Scope       string
	QuizList    []Quiz
	LastUpdated time.Time
}

type Quiz struct {
	Id         int64
	Question   string
	Answer     string
	ContentsId ContentsId
}
