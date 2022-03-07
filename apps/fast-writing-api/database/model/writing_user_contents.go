package model

import (
	"github.com/google/uuid"
	"time"
)

type UserContents struct {
	Id          int64     `gorm:"column:id"`
	UserId      uuid.UUID `gorm:"type:uuid;column:id"`
	Title       string    `gorm:"column:title"`
	LastUpdated time.Time `gorm:"column:last_updated"`
}

func (u UserContents) TableName() string {
	return "writing_user_contents"
}
