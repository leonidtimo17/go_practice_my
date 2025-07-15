package main

import (
	"github.com/gofiber/fiber/v2"
	"go_practice/internal/home"
)

func main() {
	app := fiber.New()

	home.NewHandler(app)

	app.Listen("localhost:3000")
}
