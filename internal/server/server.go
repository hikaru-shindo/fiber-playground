package server

import (
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/healthcheck"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/gofiber/fiber/v3/middleware/requestid"
)

type FiberServer struct {
	*fiber.App
}

func New(idleTimeout time.Duration) *FiberServer {
	server := &FiberServer{
		App: fiber.New(fiber.Config{
			ServerHeader:  "fiber-playground",
			AppName:       "fiber-playground",
			CaseSensitive: true,
			IdleTimeout:   idleTimeout,
		}),
	}

	server.Use(requestid.New())
	server.Use(logger.New())
	server.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{"Origin", "Content-Type", "Accept", "Authorization"},
		AllowMethods: []string{"GET", "HEAD", "PUT", "PATCH", "POST", "DELETE"},
	}))

	server.Get(healthcheck.StartupEndpoint, healthcheck.New())
	server.Get(healthcheck.ReadinessEndpoint, healthcheck.New())
	server.Get(healthcheck.LivenessEndpoint, healthcheck.New())

	return server
}
