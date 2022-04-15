package order

import (
	"gorm.io/gorm"
)

// Order Details Repo
type OrderDetailsRepository struct {
	db *gorm.DB
}

func NewOrderDetailsRepository(db *gorm.DB) *OrderDetailsRepository {
	return &OrderDetailsRepository{db: db}
}

func (r *OrderDetailsRepository) Migration() {
	r.db.AutoMigrate(&Order_details{})
}
