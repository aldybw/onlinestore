package product

import (
	"onlinestore/category"
	"time"
)

type Product struct {
	ID          int
	CategoryID  int
	Name        string
	Description string
	Price       int
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Category    category.Category
}