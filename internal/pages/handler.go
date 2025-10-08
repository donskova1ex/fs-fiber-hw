package pages

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
)

type Handler struct {
	router fiber.Router
	logger *zerolog.Logger
}

func NewHandler(router fiber.Router, logger *zerolog.Logger) {
	handler := &Handler{
		router: router,
		logger: logger,
	}

	handler.router.Get("/", handler.homePage)
}

func (h *Handler) homePage(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}
