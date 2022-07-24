package product

type ProductFormatter struct {
	ID          int    `json:"id"`
	CategoryID  int    `json:"category_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
}

func FormatProduct(product Product) ProductFormatter {
	productFormatter := ProductFormatter{}
	productFormatter.ID = product.ID
	productFormatter.CategoryID = product.CategoryID
	productFormatter.Name = product.Name
	productFormatter.Description = product.Description
	productFormatter.Price = product.Price

	return productFormatter
}

func FormatProducts(products []Product) []ProductFormatter {
	productsFormatter := []ProductFormatter{}

	for _, product := range products {
		productFormatter := FormatProduct(product)
		productsFormatter = append(productsFormatter, productFormatter)
	}

	return productsFormatter
}