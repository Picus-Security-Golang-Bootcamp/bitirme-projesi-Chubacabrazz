package cart

import (
	"github.com/Chubacabrazz/picus-storeApp/pkg/config"
	"github.com/Chubacabrazz/picus-storeApp/storage/product"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
	ID         string `gorm:"type:uuid;default:uuid_generate_v4()"`
	CustomerId string
	Products   []product.Product `gorm:"foreignKey:ID"`
}

func (Cart) TableName() string {
	//default table name
	return "Cart"
}

/* type Shopping_Session struct {
	gorm.Model
	ID     int `gorm:"unique"`
	UserId int
	Total  int
}

func (Shopping_Session) TableName() string {
	//default table name
	return "Shopping_Session"
} */

var cfg *config.Config

func Create(customer string) (*Cart, error) {
	if len(customer) == 0 {
		return nil, errors.New("CustomerId cannot be equal to zero!")
	}
	return &Cart{
		ID:         uuid.New().String(),
		CustomerId: customer,
		Products:   nil,
	}, nil
}

func (b *Cart) AddItem(sku string, price float64, quantity int) (*product.Product, error) {
	if quantity >= cfg.CartConfig.MaxAllowedQtyPerProduct {
		return nil, errors.Errorf("You can't add more this item to your basket. Maximum allowed item count is %d", cfg.CartConfig.MaxAllowedQtyPerProduct)
	}
	if (len(b.Products) + quantity) >= cfg.CartConfig.MaxAllowedForBasket {
		return nil, errors.Errorf("You can't add more item to your basket. Maximum allowed basket item count is %d", cfg.CartConfig.MaxAllowedForBasket)
	}
	_, item := b.SearchItemBySku(sku)
	if item != nil {
		return item, errors.New("Service: Product already added")
	}
	item = &product.Product{
		ID:       uuid.New().String(),
		SKU:      sku,
		Price:    price,
		Quantity: quantity,
	}

	b.Products = append(b.Products, *item)
	return item, nil
}

func (b *Cart) UpdateItem(itemId string, quantity int) (err error) {

	if index, item := b.SearchItem(itemId); index != -1 {

		if quantity >= cfg.CartConfig.MaxAllowedQtyPerProduct {
			return errors.Errorf("You can't add more item. Product count can be less then %d", cfg.CartConfig.MaxAllowedQtyPerProduct)
		}

		item.Quantity = quantity
	} else {
		return errors.Errorf("Product can not found. ItemId : %s", itemId)
	}

	return
}

func (b *Cart) RemoveItem(itemId string) (err error) {

	if index, _ := b.SearchItem(itemId); index != -1 {
		b.Products = append(b.Products[:index], b.Products[index+1:]...)
	} else {
		return errors.Errorf("Product can not found. ItemID : %s", itemId)
	}

	return
}
func (b *Cart) ValidateBasket() error {

	totalPrice := calculateBasketAmount(b)

	if totalPrice <= cfg.CartConfig.MinCartAmountForOrder {
		return errors.Errorf("Total basket amount must be greater then %f.2", cfg.CartConfig.MinCartAmountForOrder)
	}
	return nil
}

func calculateBasketAmount(b *Cart) (totalPrice float64) {
	for _, item := range b.Products {
		totalPrice += float64(item.Quantity) * item.Price
	}
	return
}
func (b *Cart) SearchItem(itemId string) (int, *product.Product) {

	for i, n := range b.Products {
		if n.ID == itemId {
			return i, &n
		}
	}
	return -1, nil
}
func (b *Cart) SearchItemBySku(sku string) (int, *product.Product) {

	for i, n := range b.Products {
		if n.SKU == sku {
			return i, &n
		}
	}
	return -1, nil
}
