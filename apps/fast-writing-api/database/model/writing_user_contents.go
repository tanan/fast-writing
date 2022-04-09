package model

import (
	"time"
)

type UserContents struct {
	Id          int64     `gorm:"column:id"`
	UserId      string    `gorm:"column:user_id"`
	Title       string    `gorm:"column:title"`
	LastUpdated time.Time `gorm:"column:last_updated"`
}

func (u UserContents) TableName() string {
	return "writing_user_contents"
}
