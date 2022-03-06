package model

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	Id          uuid.UUID `gorm:"type:uuid;primary_key;column:id"`
	Name        string    `gorm:"column:name"`
	Email       string    `gorm:"column:email"`
	LastUpdated time.Time `gorm:"column:last_updated"`
}

func (u User) TableName() string {
	return "user"
}
