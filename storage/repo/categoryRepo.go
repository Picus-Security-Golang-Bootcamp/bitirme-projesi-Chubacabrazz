package repo

import (
	"github.com/Chubacabrazz/picus-storeApp/storage/entity"
	"github.com/Chubacabrazz/picus-storeApp/storage/helper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var csvfile string = "category.csv"

type CategoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *CategoryRepository {
	return &CategoryRepository{db: db}
}

func (r *CategoryRepository) GetAll() (*[]entity.Category, error) {
	zap.L().Debug("repo.repo.getAll")

	var bs = &[]entity.Category{}
	if err := r.db.Find(&bs).Error; err != nil {
		zap.L().Error("failed to get categories", zap.Error(err))
		return nil, err
	}

	return bs, nil

}
func (r *CategoryRepository) Migration() {
	r.db.AutoMigrate(&entity.Category{})
}

// Func InsertData starts a concurrent csv reading operation and write them to database.
func (c *CategoryRepository) InsertData() {
	helper.ReadCategoryWithWorkerPool(csvfile)
	categories := []entity.Category{}
	for _, category := range helper.CategoryList {
		newItem := entity.Category{
			ID:       category.ID,
			Name:     category.Name,
			Desc:     category.Desc,
			IsActive: category.IsActive}
		categories = append(categories, newItem)
	}
	for _, eachCategory := range categories {
		//c.db.Unscoped().Where(entity.Category{Name: eachCategory.Name}).FirstOrCreate(&eachCategory)
		c.db.Unscoped().Create(&eachCategory)
	}

}
