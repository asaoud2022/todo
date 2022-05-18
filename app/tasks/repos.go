package tasks

import (
	"github.com/asaoud2022/todo/app/models"
	"github.com/google/uuid"	
	"github.com/gofiber/fiber/v2"
)

// @ func GetAll -> function that fetches a single all todos (Get all todos)
// @param c *fiber.Ctx -- fiber context
func (h handler) GetAll(c *fiber.Ctx) error {
	var todos []models.Todo

	if result := h.DB.Find(&todos); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusOK).JSON(&todos)
}

// @ func GetOne -> function that fetches a single todo (Get single todo)
// @param c *fiber.Ctx -- fiber context
func (h handler) GetOne(c *fiber.Ctx) error {
	id := c.Params("id")

	var myTodo models.Todo

	if result := h.DB.First(&myTodo, id); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusOK).JSON(&myTodo)
}

type AddTodoRequestBody struct {
	Name string `json:"name"`
}

// @func AddTodo -> function that stores a new data (Create new todo)
// @param c *fiber.Ctx -- fiber context
func (h handler) AddTodo(c *fiber.Ctx) error {
	body := AddTodoRequestBody{}

	// parse body, attach to AddProductRequestBody struct
	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	// Get the json struct that is required to send
	id := uuid.New()
	newTodo := models.Todo{
		Id:        int(id.ID()),
		Name:      body.Name,
		Completed: false,
	}

	// insert new db entry
	if result := h.DB.Create(&newTodo); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(&newTodo)
}

// @func DeleteTodo -> a function that deletes the data (Delete todo)
// @param c *fiber.Ctx -- fiber context
func (h handler) DeleteTodo(c *fiber.Ctx) error {
	id := c.Params("id")

	var deletedTodo models.Todo

	if result := h.DB.First(&deletedTodo, id); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	// delete product from db
	h.DB.Delete(&deletedTodo)

	return c.SendStatus(fiber.StatusOK)
}

type UpdateTodoRequestBody struct {
	Name string `json:"name"`
}

// @func UpdateTodo -> a function that ulters a todo data (Update todo)
// @param c *fiber.Ctx -- fiber context
func (h handler) UpdateTodo(c *fiber.Ctx) error {
	id := c.Params("id")
	body := UpdateTodoRequestBody{}

	// getting request's body
	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var updateTodo models.Todo

	if result := h.DB.First(&updateTodo, id); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	updateTodo.Name = body.Name
	// save product
	h.DB.Save(&updateTodo)

	return c.Status(fiber.StatusOK).JSON(&updateTodo)
}
