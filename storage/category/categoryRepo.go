package category

import (
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

func (r *CategoryRepository) GetAll() (*[]Category, error) {
	zap.L().Debug("repo.repo.getAll")

	var bs = &[]Category{}
	if err := r.db.Find(&bs).Error; err != nil {
		zap.L().Error("failed to get categories", zap.Error(err))
		return nil, err
	}

	return bs, nil

}
func (r *CategoryRepository) Migration() {
	r.db.AutoMigrate(&Category{})
}

// Func InsertData starts a concurrent csv reading operation and write them to database.
func (c *CategoryRepository) InsertData() {
	ReadCategoryWithWorkerPool(csvfile)
	categories := []Category{}
	for _, category := range CategoryList {
		newItem := Category{
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
