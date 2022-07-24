package product

import (
	"onlinestore/category"
	"time"

	"github.com/leekchan/accounting"
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

func (c Product) PriceFormatIDR() string {
	ac := accounting.Accounting{Symbol: "Rp", Precision: 2, Thousand: ".", Decimal: ","}
	return ac.FormatMoney(c.Price)
}

