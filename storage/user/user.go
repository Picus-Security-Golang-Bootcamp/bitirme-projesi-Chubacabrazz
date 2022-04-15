package user

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID         int `gorm:"unique"`
	Password   string
	Email      string `gorm:"unique"`
	First_name string
	Last_name  string
	IsAdmin    bool
}

func (User) TableName() string {
	//default table name
	return "User"
}
