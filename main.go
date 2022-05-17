package main

import (
	"log"

	"github.com/asaoud2022/todo/config"
	"github.com/asaoud2022/todo/database"
	"github.com/gofiber/fiber/v2"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	h := database.Init(&c)
	app := fiber.New()

	controllers.RegisterRoutes(app, h)

	app.Listen(c.Port)
}
