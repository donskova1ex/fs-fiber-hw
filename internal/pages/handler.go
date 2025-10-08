package pages

import (
	"log/slog"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	router fiber.Router
	logger *slog.Logger
}

func NewHandler(router fiber.Router, logger *slog.Logger) {
	handler := &Handler{
		router: router,
		logger: logger,
	}

	handler.router.Get("/", handler.homePage)
}

func (h *Handler) homePage(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}
