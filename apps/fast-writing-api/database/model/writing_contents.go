package model

import (
	"github.com/google/uuid"
	"time"
)

type Contents struct {
	Id          int64     `gorm:"column:id"`
	UserId      uuid.UUID `gorm:"type:uuid;column:user_id"`
	Title       string    `gorm:"column:title"`
	LastUpdated time.Time `gorm:"column:last_updated"`
}

func (c Contents) TableName() string {
	return "writing_contents"
}
