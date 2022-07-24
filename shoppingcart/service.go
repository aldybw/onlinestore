package shoppingcart

type Service interface {
	GetShoppingCarts(userID int) ([]ShoppingCart, error)
	CreateShoppingCart(input CreateShoppingCartInput, userID int) (ShoppingCart, error)
	DeleteShoppingCart(input GetShoppingCartDetailInput) error
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetShoppingCarts(userID int) ([]ShoppingCart, error) {
	shoppingCart, err := s.repository.FindByUserID(userID)
	if err != nil {
		return shoppingCart, err
	}

	return shoppingCart, nil
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

func (s *service) DeleteShoppingCart(input GetShoppingCartDetailInput) error {
	err := s.repository.Remove(input.ID)
	if err != nil {
		return err
	}

	return nil
}