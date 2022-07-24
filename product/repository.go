package product

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]Product, error)
	FindByCategoryID(productID int) ([]Product, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]Product, error) {
	var products []Product

	err := r.db.Find(&products).Error
	if err != nil {
		return products, err
	}

	return products, nil
}

func (r *repository) FindByCategoryID(productID int) ([]Product, error) {
	var products []Product

	err := r.db.Where("category_id = ?", productID).Find(&products).Error
	if err != nil {
		return products, err
	}

	return products, nil
}