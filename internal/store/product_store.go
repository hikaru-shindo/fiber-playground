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
