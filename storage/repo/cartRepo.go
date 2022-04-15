package repo

import (
	"context"

	"github.com/Chubacabrazz/picus-storeApp/storage/entity"
	"github.com/Chubacabrazz/picus-storeApp/storage/models"
	"gorm.io/gorm"
)

// Repository encapsulates the logic to access basket from the data source.
type Repository interface {
	// Get returns the basket with the specified basket Id.
	Get(ctx context.Context, id string) *models.Cart
	// GetByCustomerId returns the basket with the specified customer Id.
	GetByCustomerId(ctx context.Context, customerId string) *models.Cart
	// Create saves a new basket in the storage.
	Create(ctx context.Context, basket *models.Cart) error
	// Update updates the basket with given Is in the storage.
	Update(ctx context.Context, basket models.Cart) error
	// Delete removes the basket with given Is from the storage.
	Delete(ctx context.Context, basket models.Cart) error
}

// cartRepository Repo
type cartRepository struct {
	db *gorm.DB
}

func NewCartRepository(db *gorm.DB) *cartRepository {
	return &cartRepository{db: db}
}

func (r *cartRepository) Migration() {
	r.db.AutoMigrate(&entity.Cart{})
}

// Shopping Session Repo
/* type Session struct {
	db *gorm.DB
}

func NewSession(db *gorm.DB) *Session {
	return &Session{db: db}
}

func (r *Session) Migration() {
	r.db.AutoMigrate(&entity.Shopping_Session{})
} */
