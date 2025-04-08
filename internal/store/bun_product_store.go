package store

import (
	"context"
	"errors"

	"github.com/hikaru-shindo/fiber-playground/internal/data"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type bunProductStore struct {
	db *bun.DB
}

func NewBunProductStore(db *bun.DB) ProductStore {
	return &bunProductStore{
		db: db,
	}
}

func (store *bunProductStore) FindAll(ctx context.Context) ([]data.Product, error) {
	var products []data.Product

	store.db.WithContext(ctx).Find(&products)

	return products, nil
}

func (store *bunProductStore) FindById(ctx context.Context, id uuid.UUID) (*data.Product, error) {
	var product data.Product

	err := store.db.WithContext(ctx).Where(&id).First(&product).Error

	if err != nil {
		if errors.Is(err, bun.ErrRecordNotFound) {
			return nil, ErrProductDoesNotExist
		}

		return nil, err
	}

	return &product, nil
}

func (store *bunProductStore) Create(ctx context.Context, product data.Product) error {
	return store.db.WithContext(ctx).Transaction(func(tx *bun.DB) error {
		if err := tx.Create(&product).Error; err != nil {
			return err
		}

		return nil
	})
}

func (store *bunProductStore) Update(ctx context.Context, product data.Product) error {
	return store.db.WithContext(ctx).Transaction(func(tx *bun.DB) error {
		result := tx.Model(&product).Updates(&product)

		if result.Error != nil {
			return result.Error
		}

		if result.RowsAffected != 1 {
			return ErrProductDoesNotExist
		}

		return nil
	})
}

func (store *bunProductStore) Delete(ctx context.Context, id uuid.UUID) error {
	return store.db.WithContext(ctx).Transaction(func(tx *bun.DB) error {
		result := tx.Delete(&data.Product{}, id)

		if result.Error != nil {
			return result.Error
		}

		if result.RowsAffected != 1 {
			return ErrProductDoesNotExist
		}

		return nil
	})
}
