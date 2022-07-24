package shoppingcart

import (
	"onlinestore/product"
	"onlinestore/user"
	"time"
)

type ShoppingCart struct {
	ID        int
	UserID    int
	ProductID int
	CreatedAt time.Time
	UpdatedAt time.Time
	User      user.User
	Product   product.Product
}