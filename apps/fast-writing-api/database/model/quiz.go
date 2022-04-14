package model

import (
	"time"
)

type Quiz struct {
	Id          int64     `gorm:"column:id"`
	ContentsId  int64     `gorm:"column:contents_id"`
	UserId      string    `gorm:"column:user_id"`
	Question    string    `gorm:"column:question"`
	Answer      string    `gorm:"column:answer"`
	LastUpdated time.Time `gorm:"column:last_updated"`
}

func (q Quiz) TableName() string {
	return "quiz"
}
