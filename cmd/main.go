package main

import (
	"fiber-hw/config"
	"fiber-hw/internal/pages"
	"fiber-hw/pkg/logger"

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

	err := app.Listen(":3000")
	if err != nil {
		panic(err)
	}
}
