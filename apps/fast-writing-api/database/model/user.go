package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type User struct {
	Id          MysqlUUID `gorm:"primary_key;default:(UUID_TO_BIN(UUID()));"`
	Name        string    `gorm:"column:name"`
	Email       string    `gorm:"column:email"`
	LastUpdated time.Time `gorm:"column:last_updated"`
}

func (u User) TableName() string {
	return "user"
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	id, err := uuid.NewRandom()
	u.Id = MysqlUUID(id)
	return err
}
