package shop_test

import (
	"context"
	"github.com/hikaru-shindo/fiber-playground/internal/shop"
	"math/rand"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func newTestProduct() shop.Product {
	return shop.Product{
		Id:          uuid.New(),
		Name:        uuid.NewString(),
		Description: uuid.NewString(),
		Price: shop.Price{
			Value:    rand.Int(),
			Currency: uuid.NewString(),
		},
	}
}

func TestInMemoryProductStore_Create(t *testing.T) {
	testProduct := newTestProduct()

	sut := shop.NewInMemoryProductStore()
	err := sut.Create(context.Background(), testProduct)
	result, _ := sut.FindAll(context.Background())

	assert.Nil(t, err)
	assert.Len(t, result, 1)
	assert.Equal(t, testProduct, result[0])
}

func TestInMemoryProductStore_Delete(t *testing.T) {
	testProduct := newTestProduct()

	tests := []struct {
		name             string
		products         []shop.Product
		expectedProducts []shop.Product
		expectedError    error
	}{
		{
			name:             "deletes product successfully",
			products:         []shop.Product{testProduct},
			expectedProducts: make([]shop.Product, 0),
			expectedError:    nil,
		},
		{
			name:             "returns error if product does not exist",
			products:         make([]shop.Product, 0),
			expectedProducts: make([]shop.Product, 0),
			expectedError:    shop.ErrProductDoesNotExist,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sut := shop.NewInMemoryProductStore()

			for _, product := range tt.products {
				_ = sut.Create(context.Background(), product)
			}

			err := sut.Delete(context.Background(), testProduct.Id)
			result, _ := sut.FindAll(context.Background())

			assert.Equal(t, tt.expectedError, err)
			assert.Equal(t, tt.expectedProducts, result)
		})
	}
}

func TestInMemoryProductStore_FindAll(t *testing.T) {
	tests := []struct {
		name     string
		products []shop.Product
	}{
		{
			name:     "returns all products",
			products: []shop.Product{newTestProduct(), newTestProduct(), newTestProduct()},
		}, {
			name:     "returns empty list if no product exists",
			products: make([]shop.Product, 0),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sut := shop.NewInMemoryProductStore()

			for _, product := range tt.products {
				_ = sut.Create(context.Background(), product)
			}

			result, err := sut.FindAll(context.Background())

			assert.Nil(t, err)
			assert.Len(t, result, len(tt.products))

			for _, product := range tt.products {
				assert.Contains(t, result, product)
			}
		})
	}
}

func TestInMemoryProductStore_FindById(t *testing.T) {
	testProduct := newTestProduct()

	tests := []struct {
		name          string
		id            uuid.UUID
		expectedError error
	}{
		{
			name:          "returns product by id",
			id:            testProduct.Id,
			expectedError: nil,
		}, {
			name:          "returns error if product does not exist",
			id:            uuid.New(),
			expectedError: shop.ErrProductDoesNotExist,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sut := shop.NewInMemoryProductStore()

			for _, product := range []shop.Product{newTestProduct(), newTestProduct(), testProduct, newTestProduct()} {
				_ = sut.Create(context.Background(), product)
			}

			result, err := sut.FindById(context.Background(), tt.id)

			if tt.expectedError != nil {
				assert.Equal(t, tt.expectedError, err)
				assert.Nil(t, result)
			} else {
				assert.Nil(t, err)
				assert.Equal(t, testProduct, *result)
			}
		})

	}
}

func TestInMemoryProductStore_Update(t *testing.T) {
	testProduct := newTestProduct()
	tests := []struct {
		name          string
		modify        func(*shop.Product)
		expectedError error
	}{
		{
			name: "updates product correctly",
			modify: func(product *shop.Product) {
				product.Description = uuid.NewString()
			},
			expectedError: nil,
		}, {
			name: "returns error if id changes or product does not exist",
			modify: func(product *shop.Product) {
				product.Id = uuid.New()
			},
			expectedError: shop.ErrProductDoesNotExist,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sut := shop.NewInMemoryProductStore()

			for _, product := range []shop.Product{newTestProduct(), testProduct, newTestProduct()} {
				_ = sut.Create(context.Background(), product)
			}

			updatedProduct := testProduct.Clone()
			tt.modify(&updatedProduct)

			err := sut.Update(context.Background(), updatedProduct)

			if tt.expectedError != nil {
				assert.Equal(t, tt.expectedError, err)
			} else {
				result, _ := sut.FindById(context.Background(), testProduct.Id)

				assert.Equal(t, updatedProduct, *result)
			}
		})
	}
}
