package shop

import (
	"github.com/gofiber/fiber/v2"
)

func (handler *Handler) Register(router fiber.Router) {
	router.Get("/product", handler.getAllProducts)
	router.Post("/product", handler.createProduct)
	router.Put("/product/:id", handler.updateProduct)
	router.Delete("/product/:id", handler.deleteProduct)
}
