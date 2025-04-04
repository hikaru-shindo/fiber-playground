package handler

import (
	"github.com/gofiber/fiber/v2"
)

func (handler *Handler) Register(router fiber.Router) {
	product := router.Group("/product")
	product.Get("", handler.Products)
	product.Post("", handler.CreateProduct)
	product.Get("/:id", handler.GetProduct)
	product.Put("/:id", handler.UpdateProduct)
	product.Delete("/:id", handler.DeleteProduct)
}
