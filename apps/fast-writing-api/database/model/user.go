package model

import (
	"time"
)

type User struct {
	Id          string    `gorm:"column:id;primary_key"`
	Name        string    `gorm:"column:name"`
	Email       string    `gorm:"column:email"`
	LastUpdated time.Time `gorm:"column:last_updated"`
}

func (u User) TableName() string {
	return "user"
}
