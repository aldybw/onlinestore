package product

type Service interface {
	GetProducts(categoryID int) ([]Product, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetProducts(categoryID int) ([]Product, error) {
	if categoryID != 0 {
		products, err := s.repository.FindByCategoryID(categoryID)
		if err != nil {
			return products, err
		}

		return products, nil
	}

	products, err := s.repository.FindAll()
	if err != nil {
		return products, err
	}

	return products, nil
}