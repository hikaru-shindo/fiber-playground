package store

import (
	"context"
	"fmt"
	"github.com/hikaru-shindo/fiber-playground/internal/data"

	"github.com/google/uuid"
)

type ProductStore interface {
	FindAll(context.Context) ([]data.Product, error)
	FindById(context.Context, uuid.UUID) (*data.Product, error)
	Create(context.Context, data.Product) error
	Update(context.Context, data.Product) error
	Delete(context.Context, uuid.UUID) error
}

var ErrProductDoesNotExist = fmt.Errorf("product does not exist")

type inMemoryProductStore struct {
	products []data.Product
}

// NewInMemoryProductStore creates a simple in memory product store which will not be thread safe.
func NewInMemoryProductStore() ProductStore {
	return &inMemoryProductStore{
		products: make([]data.Product, 0),
	}
}

func (store *inMemoryProductStore) FindAll(_ context.Context) ([]data.Product, error) {
	return store.products, nil
}

func (store *inMemoryProductStore) FindById(_ context.Context, id uuid.UUID) (*data.Product, error) {
	for _, product := range store.products {
		if product.Id == id {
			productCopy := product.Clone()
			return &productCopy, nil
		}
	}

	return nil, ErrProductDoesNotExist
}

func (store *inMemoryProductStore) Create(_ context.Context, product data.Product) error {
	store.products = append(store.products, product)

	return nil
}

func (store *inMemoryProductStore) Update(_ context.Context, product data.Product) error {
	for index, existingProduct := range store.products {
		if existingProduct.Id == product.Id {
			store.products[index] = product
			return nil
		}
	}

	return ErrProductDoesNotExist
}

func (store *inMemoryProductStore) Delete(_ context.Context, id uuid.UUID) error {
	for index, product := range store.products {
		if product.Id == id {
			store.products = append(store.products[:index], store.products[index+1:]...)
			return nil
		}
	}

	return ErrProductDoesNotExist
}
