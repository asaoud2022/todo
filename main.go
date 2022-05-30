package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"github.com/asaoud2022/todo/app/handlers"
	"github.com/asaoud2022/todo/app/repository"
	"github.com/asaoud2022/todo/config"
)

func main() {

	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	// Connect to database
	_, err1 := repository.Connect(&c)
	if err1 != nil {
		log.Fatal(err1)
	}

	// Create a new Fiber instance
	app := fiber.New()

	// Response with a hello message for calling root path
	app.Get("/", hello)

	// Use logger
	app.Use(logger.New())

	// Group User related APIs
	userGroup := app.Group("/user")

	userGroup.Get("/", handlers.GetAllUsers)
	userGroup.Get("/:id", handlers.GetSingleUser)
	userGroup.Post("/", handlers.AddNewUser)
	userGroup.Patch("/:id", handlers.UpdateUser)
	userGroup.Delete("/:id", handlers.DeleteUser)

	// Group Task related APIs
	taskGroup := app.Group("/task")

	taskGroup.Get("/", handlers.GetAllTasks)
	taskGroup.Get("/:id", handlers.GetSingleTask)
	taskGroup.Post("/", handlers.AddNewTask)
	taskGroup.Put("/:id", handlers.UpdateTask)
	taskGroup.Delete("/:id", handlers.DeleteTask)

	err = app.Listen(c.Port)
	if err != nil {
		log.Fatal(err)
	}
}

func hello(c *fiber.Ctx) error {
	return c.SendString("Hello, World 👋!")
}
