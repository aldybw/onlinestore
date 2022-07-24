package shoppingcart

type Service interface {
	CreateShoppingCart(input CreateShoppingCartInput, userID int) (ShoppingCart, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) CreateShoppingCart(input CreateShoppingCartInput, userID int) (ShoppingCart, error) {
	shoppingCart := ShoppingCart{}
	shoppingCart.UserID = userID
	shoppingCart.ProductID = input.ProductID

	newShoppingCart, err := s.repository.Save(shoppingCart)
	if err != nil {
		return newShoppingCart, err
	}

	return newShoppingCart, nil
}