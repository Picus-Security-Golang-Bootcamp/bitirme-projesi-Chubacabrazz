package category

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	ID       string `gorm:"unique"`
	Name     string
	Desc     string
	IsActive bool
}

func (Category) TableName() string {
	//default table name
	return "Category"
}
