package shop

import "github.com/google/uuid"

type Product struct {
	Id          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       Price     `json:"price"`
}

func (product Product) Clone() Product {
	clonedProduct := new(Product)
	*clonedProduct = product

	return *clonedProduct
}

type Price struct {
	Value    int    `json:"value"`
	Currency string `json:"currency"`
}
