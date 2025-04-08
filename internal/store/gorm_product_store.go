package store

import (
	"context"
	"errors"

	"github.com/hikaru-shindo/fiber-playground/internal/data"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type gormProductStore struct {
	db *gorm.DB
}

func NewGormProductStore(db *gorm.DB) ProductStore {
	return &gormProductStore{
		db: db,
	}
}

func (store *gormProductStore) FindAll(ctx context.Context) ([]data.Product, error) {
	var products []data.Product

	store.db.WithContext(ctx).Find(&products)

	return products, nil
}

func (store *gormProductStore) FindById(ctx context.Context, id uuid.UUID) (*data.Product, error) {
	var product data.Product

	err := store.db.WithContext(ctx).Where(&id).First(&product).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrProductDoesNotExist
		}

		return nil, err
	}

	return &product, nil
}

func (store *gormProductStore) Create(ctx context.Context, product data.Product) error {
	return store.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&product).Error; err != nil {
			return err
		}

		return nil
	})
}

func (store *gormProductStore) Update(ctx context.Context, product data.Product) error {
	return store.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
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

func (store *gormProductStore) Delete(ctx context.Context, id uuid.UUID) error {
	return store.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
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
