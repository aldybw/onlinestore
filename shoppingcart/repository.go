package shoppingcart

import "gorm.io/gorm"

type Repository interface {
	FindByUserID(userID int) ([]ShoppingCart, error)
	Save(shoppingCart ShoppingCart) (ShoppingCart, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindByUserID(userID int) ([]ShoppingCart, error) {
	var shoppingCarts []ShoppingCart

	err := r.db.Where("user_id = ?", userID).Find(&shoppingCarts).Error
	if err != nil {
		return shoppingCarts, err
	}

	return shoppingCarts, nil
}

func (r *repository) Save(shoppingCart ShoppingCart) (ShoppingCart, error) {
	err := r.db.Create(&shoppingCart).Error
	if err != nil {
		return shoppingCart, err
	}

	return shoppingCart, nil
}