package main

import (
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/hikaru-shindo/fiber-playground/internal/database"

	"github.com/hikaru-shindo/fiber-playground/internal/handler"
	"github.com/hikaru-shindo/fiber-playground/internal/server"
	"github.com/hikaru-shindo/fiber-playground/internal/store"

	_ "github.com/joho/godotenv/autoload"
)

const idleTimeout = time.Second * 5

func main() {
	server := server.New(idleTimeout)

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

	go func() {
		port, _ := strconv.Atoi(os.Getenv("SERVER_PORT"))
		err := server.Listen(fmt.Sprintf(":%d", port))
		if err != nil {
			panic(fmt.Sprintf("http server error: %s", err))
		}
	}()

	channel := make(chan os.Signal, 1)                    // Create channel to signify a signal being sent
	signal.Notify(channel, os.Interrupt, syscall.SIGTERM) // When an interrupt or termination signal is sent, notify the channel

	<-channel // This blocks the main thread until an interrupt is received
	fmt.Println("Gracefully shutting down...")
	_ = server.Shutdown()

	fmt.Println("Shutdown completed successfully.")
}
