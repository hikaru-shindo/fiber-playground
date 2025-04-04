package handler

import (
	"github.com/google/uuid"
	"github.com/hikaru-shindo/fiber-playground/internal/data"
)

type productResponse struct {
	Id          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       struct {
		Value    int    `json:"value"`
		Currency string `json:"currency"`
	} `json:"price"`
}

type productListResponse struct {
	Products []*productResponse `json:"products"`
}

func newProductResponse(product data.Product) *productResponse {
	response := new(productResponse)

	response.Id = product.Id
	response.Name = product.Name
	response.Description = product.Description

	response.Price.Value = product.Price.Value
	response.Price.Currency = product.Price.Currency

	return response
}

func newProductListResponse(products ...data.Product) *productListResponse {
	response := new(productListResponse)
	response.Products = make([]*productResponse, len(products))

	for index, product := range products {
		response.Products[index] = newProductResponse(product)
	}

	return response
}
