package store_test

import (
	"github.com/hikaru-shindo/fiber-playground/internal/data"
	"math/rand"

	"github.com/google/uuid"
)

// newTestProduct creates a new product with test data
func newTestProduct() data.Product {
	product := new(data.Product)

	product.Id = uuid.New()
	product.Name = uuid.NewString()
	product.Description = uuid.NewString()

	product.Price.Value = rand.Int()
	product.Price.Currency = uuid.NewString()

	return *product
}
