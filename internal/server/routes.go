package server

import (
	"github.com/hikaru-shindo/fiber-playground/internal/shop"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

func (server *FiberServer) RegisterDefaultRoutes(productStore shop.ProductStore) {
	server.Use(requestid.New())
	server.Use(healthcheck.New())
	server.Use(logger.New())
	server.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
		AllowMethods: "GET, HEAD, PUT, PATCH, POST, DELETE",
	}))

	server.Route("/shop", func(router fiber.Router) {
		handler := shop.NewHandler(productStore)
		handler.Register(router)
	}, "shop")
}
