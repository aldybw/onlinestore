package shoppingcart

import (
	"onlinestore/product"
	"onlinestore/user"
)

type CreateShoppingCartInput struct {
	ProductID int `json:"product_id" binding:"required"`
	User      user.User
	Product   product.Product
}
