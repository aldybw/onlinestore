package shoppingcart

import "gorm.io/gorm"

type Repository interface {
	Save(shoppingCart ShoppingCart) (ShoppingCart, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(shoppingCart ShoppingCart) (ShoppingCart, error) {
	err := r.db.Create(&shoppingCart).Error
	if err != nil {
		return shoppingCart, err
	}

	return shoppingCart, nil
}