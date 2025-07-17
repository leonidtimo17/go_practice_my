package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"go_practice/config"
	"go_practice/internal/home"
	"log"
)

func main() {
	config.Init()
	dbConf := config.NewDatabaseConfig()
	log.Println(dbConf)
	app := fiber.New()
	app.Use(recover.New())

	home.NewHandler(app)

	app.Listen("localhost:3000")
}
