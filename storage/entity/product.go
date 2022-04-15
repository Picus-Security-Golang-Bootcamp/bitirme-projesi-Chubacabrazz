package entity

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	ID       int    `gorm:"unique"`
	Name     string `gorm:"unique"`
	Desc     string
	SKU      string
	Price    int
	Quantity int
}

func (Product) TableName() string {
	//default table name
	return "Product"
}
