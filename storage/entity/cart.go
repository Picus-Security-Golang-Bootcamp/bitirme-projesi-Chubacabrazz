package entity

import (
	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
	ID         string `gorm:"type:uuid;default:uuid_generate_v4()"`
	CustomerId int
	Products   []Product `gorm:"foreignKey:ID"`
}

func (Cart) TableName() string {
	//default table name
	return "Cart"
}

/* type Shopping_Session struct {
	gorm.Model
	ID     int `gorm:"unique"`
	UserId int
	Total  int
}

func (Shopping_Session) TableName() string {
	//default table name
	return "Shopping_Session"
} */
