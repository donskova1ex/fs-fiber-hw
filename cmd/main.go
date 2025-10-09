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

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/template/html/v2"
	"github.com/samber/slog-fiber"
)

func main() {
	config.Init()
	cfg := config.NewConfig()

	customLog := logger.NewLogger(cfg.LogConfig)
	engine := html.New("./html", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})
	
	app.Use(slogfiber.New(customLog))
	app.Use(recover.New())

	pages.NewHandler(app, customLog)

	go func() {
		err := app.Listen(cfg.AppConfig.Port)
		if err != nil {
			customLog.Error(err.Error())
		}
	}()
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan
	customLog.Info("Shutting down signal received")

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	customLog.Info("Shutting down gracefully for 3 seconds")
	if err := app.ShutdownWithContext(ctx); err != nil {
		customLog.Error(err.Error())
	}
	customLog.Info("Server gracefully stopped")
}
