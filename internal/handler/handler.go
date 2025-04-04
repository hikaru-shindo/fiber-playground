package handler

import "github.com/hikaru-shindo/fiber-playground/internal/store"

type Handler struct {
	productStore store.ProductStore
}

func NewHandler(productStore store.ProductStore) *Handler {
	return &Handler{
		productStore: productStore,
	}
}
