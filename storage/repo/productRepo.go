package repo

import (
	"github.com/Chubacabrazz/picus-storeApp/storage/entity"
	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) Migration() {
	r.db.AutoMigrate(&entity.Product{})
}
