package domain

import "time"

type ContentsId int64

type Contents struct {
	Id       ContentsId `json:"contents_id"`
	Title    string     `json:"title"`
	Category string     `json:"category"`
	Username string     `json:"username"`
	Question string     `json:"question"`
	Answer   string     `json:"answer"`
}

type ContentsScore struct {
	Id          ContentsId
	Score       float32
	LastUpdated time.Time
}
