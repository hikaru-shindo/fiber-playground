package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hikaru-shindo/fiber-playground/internal/api"
)

func (server *FiberServer) RegisterFiberRoutes() {
	server.Route("/api", func(router fiber.Router) {
		api.RegisterRoutes(router)
	}, "api")
}
