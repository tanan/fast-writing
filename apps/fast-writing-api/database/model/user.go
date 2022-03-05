package model

type User struct {
	Id          string `gorm:"column:id"`
	Name        string `gorm:"column:name"`
	Email       string `gorm:"column:email"`
	LastUpdated string `gorm:"column:last_updated"`
}
