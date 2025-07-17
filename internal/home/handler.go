package home

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

type HomeHandler struct {
	router fiber.Router
}

// v1/
// v1/api/
// v1/api/error

func NewHandler(router fiber.Router) {
	h := &HomeHandler{
		router: router,
	}
	v1 := h.router.Group("/v1")
	api := v1.Group("/api")
	api.Get("/", h.home)
	api.Get("/error", h.error)
}

func (h *HomeHandler) home(c *fiber.Ctx) error {
	return c.SendString("Hello world")
}

func (h *HomeHandler) error(c *fiber.Ctx) error {
	log.Info("Info")
	return fiber.NewError(fiber.StatusBadRequest, "Limit Params undefined")
}

// panic
