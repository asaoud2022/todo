package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/asaoud2022/todo/app/repos"


	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

func RegisterRoutes(app *fiber.App, db *gorm.DB) {
	h := &handler{
		DB: db,
	}

	// Group is used for Routes with common prefix to define a new sub-router with optional middleware.
	routes := app.Group("/todos")

	// Route for Get all todos -> navigate to => http://127.0.0.1:3000/v1/todos/
	routes.Get("/", todos.GetAll)

	// Route for Get a todo -> navigate to => http://127.0.0.1:3000/v1/todos/<todo's id>
	routes.Get("/:id", todo.GetOne)

	// Route for Add a todo -> navigate to => http://127.0.0.1:3000/v1/todos/
	routes.Post("/", todos.AddTodo)

	// Route for Delete a todo -> navigate to => http://127.0.0.1:3000/v1/todos/<todo's id>
	routes.Delete("/:id", todos.DeleteTodo)

	// Route for Update a todo -> navigate to => http://127.0.0.1:3000/v1/todos/<todo's id>
	routes.Patch("/:id", todos.UpdateTodo)
}
