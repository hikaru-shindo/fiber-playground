package main

import (
	"context"
	"fmt"
	"github.com/hikaru-shindo/fiber-playground/internal/database"
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/hikaru-shindo/fiber-playground/internal/handler"
	"github.com/hikaru-shindo/fiber-playground/internal/server"
	"github.com/hikaru-shindo/fiber-playground/internal/store"

	_ "github.com/joho/godotenv/autoload"
)

func gracefulShutdown(fiberServer *server.FiberServer, done chan bool) {
	// Create context that listens for the interrupt signal from the OS.
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// Listen for the interrupt signal.
	<-ctx.Done()

	log.Println("shutting down gracefully, press Ctrl+C again to force")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := fiberServer.ShutdownWithContext(ctx); err != nil {
		log.Printf("Server forced to shutdown with error: %v", err)
	}

	log.Println("Server exiting")

	// Notify the main goroutine that the shutdown is complete
	done <- true
}

func main() {

	server := server.New()

	db, err := database.NewGormSqliteDatabase(os.Getenv("DATABASE_FILE"))
	if err != nil {
		panic(fmt.Sprintf("storage error: %s", err))
	}

	if err := database.GormMigrate(db); err != nil {
		panic(fmt.Sprintf("migration error: %s", err))
	}

	productStore := store.NewGormProductStore(db)

	handler := handler.NewHandler(productStore)
	handler.Register(server)

	// Create a done channel to signal when the shutdown is complete
	done := make(chan bool, 1)

	go func() {
		port, _ := strconv.Atoi(os.Getenv("SERVER_PORT"))
		err := server.Listen(fmt.Sprintf(":%d", port))
		if err != nil {
			panic(fmt.Sprintf("http server error: %s", err))
		}
	}()

	// Run graceful shutdown in a separate goroutine
	go gracefulShutdown(server, done)

	// Wait for the graceful shutdown to complete
	<-done
	log.Println("Graceful shutdown complete.")
}
