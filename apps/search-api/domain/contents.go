package domain

import "time"

type ContentsId int64

type Contents struct {
	Id       ContentsId
	Title    string
	Category string
	Username string
	Quiz     string
	Answer   string
}

type ContentsScore struct {
	Id          ContentsId
	Score       float32
	LastUpdated time.Time
}
