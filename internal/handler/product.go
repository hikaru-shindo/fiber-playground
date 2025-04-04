package handler

import (
	ctx "context"
	"errors"

	"github.com/hikaru-shindo/fiber-playground/internal/data"
	"github.com/hikaru-shindo/fiber-playground/internal/store"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/google/uuid"
)

func (handler *Handler) Products(context *fiber.Ctx) error {
	products, err := handler.productStore.FindAll(ctx.Background())

	if err != nil {
		return err
	}

	return context.JSON(products)
}

func (handler *Handler) GetProduct(context *fiber.Ctx) error {
	productId, err := uuid.Parse(context.Params("id"))

	if err != nil {
		return err
	}

	product, err := handler.productStore.FindById(ctx.Background(), productId)

	if err != nil && !errors.Is(err, store.ErrProductDoesNotExist) {
		return err
	} else if errors.Is(err, store.ErrProductDoesNotExist) {
		return fiber.NewError(fiber.StatusNotFound)
	}

	return context.JSON(product)
}

func (handler *Handler) CreateProduct(context *fiber.Ctx) error {
	context.Accepts("application/json")
	product := new(data.Product)

	if err := context.BodyParser(product); err != nil {
		return err
	}

	product.Id = uuid.New()

	if err := handler.productStore.Create(ctx.Background(), *product); err != nil {
		return err
	}

	log.Infow("Product added", "productId", product.Id.String())

	context.Status(fiber.StatusCreated)

	return context.JSON(product)
}

func (handler *Handler) DeleteProduct(context *fiber.Ctx) error {
	productId, err := uuid.Parse(context.Params("id"))

	if err != nil {
		return err
	}

	err = handler.productStore.Delete(ctx.Background(), productId)

	if err != nil && !errors.Is(err, store.ErrProductDoesNotExist) {
		return err
	}

	context.Status(fiber.StatusNoContent)

	return nil
}

func (handler *Handler) UpdateProduct(context *fiber.Ctx) error {
	context.Accepts("application/json")
	productId, err := uuid.Parse(context.Params("id"))
	updatedProduct := new(data.Product)

	if err != nil {
		return err
	}

	if err := context.BodyParser(updatedProduct); err != nil {
		return err
	}

	_, err = handler.productStore.FindById(ctx.Background(), productId)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound)
	}

	err = handler.productStore.Update(ctx.Background(), *updatedProduct)
	if err != nil {
		return err
	}

	return context.JSON(updatedProduct)
}
