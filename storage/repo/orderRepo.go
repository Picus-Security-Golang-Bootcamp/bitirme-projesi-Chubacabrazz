package repo

import (
	"github.com/Chubacabrazz/picus-storeApp/storage/entity"
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
	r.db.AutoMigrate(&entity.Order_details{})
}
