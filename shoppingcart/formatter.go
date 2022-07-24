package shoppingcart

type ShoppingCartFormatter struct {
	ID        int    `json:"id"`
	UserID    int    `json:"user_id"`
	ProductID int    `json:"product_id"`
	Name      string `json:"name"`
}

func FormatShoppingCart(shoppingCart ShoppingCart, userID int) ShoppingCartFormatter {
	ShoppingCartFormatter := ShoppingCartFormatter{}
	ShoppingCartFormatter.ID = shoppingCart.ID
	ShoppingCartFormatter.UserID = userID
	ShoppingCartFormatter.ProductID = shoppingCart.ProductID
	ShoppingCartFormatter.Name = shoppingCart.Product.Name

	return ShoppingCartFormatter
}

func FormatShoppingCarts(shoppingCarts []ShoppingCart, userID int) []ShoppingCartFormatter {
	shoppingCartsFormatter := []ShoppingCartFormatter{}

	for _, shoppingCart := range shoppingCarts {
		ShoppingCartFormatter := FormatShoppingCart(shoppingCart, userID)
		shoppingCartsFormatter = append(shoppingCartsFormatter, ShoppingCartFormatter)
	}

	return shoppingCartsFormatter
}