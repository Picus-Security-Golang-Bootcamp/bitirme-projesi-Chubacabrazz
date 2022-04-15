package entity

import (
	"gorm.io/gorm"
)

type Order_details struct {
	gorm.Model
	ID         string `gorm:"type:uuid;default:uuid_generate_v4()"`
	User_Id    int
	Total      int
	IsCanceled bool
}

func (Order_details) TableName() string {
	//default table name
	return "Order_details"
}
