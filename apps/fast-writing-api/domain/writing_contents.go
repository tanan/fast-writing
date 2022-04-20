package domain

import "time"

type ContentsId int64

type Contents struct {
	ContentsId  ContentsId
	Creator     string
	Title       string
	Description string
	Scope       string
	QuizList    []Quiz `json:"quizlist"`
	LastUpdated time.Time
}

type Quiz struct {
	Question string `json:"question"`
	Answer   string `json:"answer"`
	Order    string `json:"order"`
}
