package product

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	ID       string `gorm:"unique"`
	Name     string `gorm:"unique"`
	Desc     string
	SKU      string
	Price    float64
	Quantity int
}

func (Product) TableName() string {
	//default table name
	return "Product"
}
