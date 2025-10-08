package main

import (
	"context"
	"fiber-hw/config"
	"fiber-hw/internal/pages"
	"fiber-hw/pkg/logger"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/contrib/fiberzerolog"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	config.Init()
	cfg := config.NewConfig()

	customLog := logger.NewLogger(cfg.LogConfig)

	app := fiber.New(fiber.Config{})
	app.Use(fiberzerolog.New(fiberzerolog.Config{
		Logger: customLog,
	}))
	app.Use(recover.New())

	pages.NewHandler(app, customLog)

	go func() {
		err := app.Listen(cfg.AppConfig.Port)
		if err != nil {
			customLog.Println("Error starting server")
		}
	}()
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan
	customLog.Println("Shutting down signal received")

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	customLog.Println("Shutting down gracefully for 3 seconds")
	if err := app.ShutdownWithContext(ctx); err != nil {
		customLog.Println("Error shutting down app")
	}
	customLog.Println("Server gracefully stopped")
}
