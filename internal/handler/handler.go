package handler

import "github.com/hikaru-shindo/fiber-playground/internal/store"

type Handler struct {
	validator    *Validator
	productStore store.ProductStore
}

func NewHandler(productStore store.ProductStore) *Handler {
	validator := NewValidator()

	return &Handler{
		validator:    validator,
		productStore: productStore,
	}
}
