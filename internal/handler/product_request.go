package handler

import (
	"github.com/hikaru-shindo/fiber-playground/internal/data"

	"github.com/gofiber/fiber/v2"
)

type productCreateRequest struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description"`
	Price       struct {
		Value    int    `json:"value" validate:"required,min=0"`
		Currency string `json:"currency" validate:"required,iso4217"`
	} `json:"price"`
}

type productUpdateRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       struct {
		Value    int    `json:"value" validate:"min=0"`
		Currency string `json:"currency" validate:"iso4217"`
	} `json:"price"`
}

func (request *productCreateRequest) bind(context *fiber.Ctx, product *data.Product, validator *Validator) error {
	if err := context.BodyParser(request); err != nil {
		return err
	}

	if err := validator.Validate(request); err != nil {
		return err
	}

	product.Name = request.Name
	product.Description = request.Description
	product.Price.Value = request.Price.Value
	product.Price.Currency = request.Price.Currency

	return nil
}

func (request *productUpdateRequest) populate(product *data.Product) {
	request.Name = product.Name
	request.Description = product.Description

	request.Price.Value = product.Price.Value
	request.Price.Currency = product.Price.Currency
}

func (request *productUpdateRequest) bind(context *fiber.Ctx, product *data.Product, validator *Validator) error {
	if err := context.BodyParser(request); err != nil {
		return err
	}

	if err := validator.Validate(request); err != nil {
		return err
	}

	product.Name = request.Name
	product.Description = request.Description
	product.Price.Value = request.Price.Value
	product.Price.Currency = request.Price.Currency

	return nil
}
