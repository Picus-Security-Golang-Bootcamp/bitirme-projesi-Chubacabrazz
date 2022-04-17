package cart

import (
	"context"

	"gorm.io/gorm"
)

type Repository interface {
	// Get returns the basket with the specified basket Id.
	Get(ctx context.Context, id string) *Cart
	// GetByCustomerId returns the basket with the specified customer Id.
	GetByCustomerId(ctx context.Context, customerId string) *Cart
	// Create saves a new basket in the storage.
	Create(ctx context.Context, basket *Cart) error
	// Update updates the basket with given Is in the storage.
	Update(ctx context.Context, basket Cart) error
	// Delete removes the basket with given Is from the storage.
	Delete(ctx context.Context, basket Cart) error
}

// cartRepository Repo
type cartRepository struct {
	db *gorm.DB
}

func NewCartRepository(db *gorm.DB) *cartRepository {
	return &cartRepository{db: db}
}

func (r *cartRepository) Migration() {
	r.db.AutoMigrate(&Cart{})
}
