package shop

import (
	ctx "context"
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/google/uuid"
)

func (handler *Handler) getAllProducts(context *fiber.Ctx) error {
	products, err := handler.productStore.FindAll(ctx.Background())

	if err != nil {
		return err
	}

	return context.JSON(products)
}

func (handler *Handler) createProduct(context *fiber.Ctx) error {
	context.Accepts("application/json")
	product := new(Product)

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

func (handler *Handler) deleteProduct(context *fiber.Ctx) error {
	productId, err := uuid.Parse(context.Params("id"))

	if err != nil {
		return err
	}

	err = handler.productStore.Delete(ctx.Background(), productId)

	if err != nil && !errors.Is(err, ErrProductDoesNotExist) {
		return err
	}

	context.Status(fiber.StatusNoContent)

	return nil
}

func (handler *Handler) updateProduct(context *fiber.Ctx) error {
	context.Accepts("application/json")
	productId, err := uuid.Parse(context.Params("id"))
	updatedProduct := new(Product)

	if err != nil {
		return err
	}

	if err := context.BodyParser(updatedProduct); err != nil {
		return err
	}

	_, err = handler.productStore.FindById(ctx.Background(), productId)
	if err != nil {
		context.Status(fiber.StatusNotFound)
		return nil
	}

	err = handler.productStore.Update(ctx.Background(), *updatedProduct)
	if err != nil {
		return err
	}

	return context.JSON(updatedProduct)
}
