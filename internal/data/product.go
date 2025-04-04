package data

import (
	"github.com/google/uuid"
)

type Product struct {
	Id          uuid.UUID `gorm:"type:uuid,primarykey"`
	Name        string
	Description string
	Price       struct {
		Value    int    `gorm:"column:price"`
		Currency string `gorm:"column:currency"`
	} `gorm:"embedded"`
}

func (Product) TableName() string {
	return "product"
}

func (product Product) Clone() Product {
	clonedProduct := new(Product)
	*clonedProduct = product

	return *clonedProduct
}
