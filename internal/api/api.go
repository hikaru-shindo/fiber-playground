package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/google/uuid"
)

type Product struct {
	Id          uuid.UUID `json:"id"`
    Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       Price     `json:"price"`
}

type Price struct {
	Value    int    `json:"value"`
	Currency string `json:"currency"`
}

var products []Product = make([]Product, 0)

func RegisterRoutes(router fiber.Router) {
	router.Get("/test", func(context *fiber.Ctx) error {
		return context.SendString("worked")
	}).Name(".test")

	router.Get("/json", func(context *fiber.Ctx) error {
		return context.JSON(fiber.Map{
			"foo": "bar",
		})
	}).Name(".json")

	router.Get("/product", func(context *fiber.Ctx) error {
		return context.JSON(products)
	}).Name(".product")

	router.Post("/product", func(context *fiber.Ctx) error {
        context.Accepts("application/json")
        product := new(Product)

        if err := context.BodyParser(product); err != nil {
            return err
        }

        product.Id = uuid.New()

        products = append(products, *product)

        log.Infow("Product added", "productId", product.Id.String())

        context.Status(fiber.StatusCreated)

        return context.JSON(product)
	})

    router.Delete("/product/:id", func(context *fiber.Ctx) error {
        productId, err := uuid.Parse(context.Params("id"))

        if err != nil {
            return err
        }

        for index, product := range products {
            if product.Id == productId {
                products = append(products[:index], products[index+1:]...)
            }
        }

        context.Status(fiber.StatusNoContent)

        return nil
    })

    router.Put("product/:id", func(context *fiber.Ctx) error {
        context.Accepts("application/json")
        productId, err := uuid.Parse(context.Params("id"))
        updatedProduct := new(Product)

        if err != nil {
            return err
        }

        if err := context.BodyParser(updatedProduct); err != nil {
            return err
        }

        for index, product := range products {
            if product.Id == productId {
                updatedProduct.Id = product.Id

                products[index] = *updatedProduct

                return context.JSON(updatedProduct)
            }
        }

        context.Status(fiber.StatusNotFound)
        return nil
    })
}
