package data

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Product struct {
	bun.BaseModel `bun:"table:product"`

	Id          uuid.UUID `gorm:"type:uuid,primarykey" bun:",type:uuid,pk"`
	Name        string
	Description string
	Price       struct {
		Value    int    `gorm:"column:price" bun:"price"`
		Currency string `gorm:"column:currency" bun:"currency"`
	} `gorm:"embedded" bun:"embed:"`
}

func (Product) TableName() string {
	return "product"
}

func (product Product) Clone() Product {
	clonedProduct := new(Product)
	*clonedProduct = product

	return *clonedProduct
}
