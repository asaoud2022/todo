package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/asaoud2022/todo/app/models"
	"github.com/asaoud2022/todo/app/repository"
	"github.com/gofiber/fiber/v2"
)

var taskRepository repository.TaskRepository

func init() {
	taskRepository = repository.NewTaskRepository()
}

// GetAllCompanies gets all repository information
func GetAllTasks(c *fiber.Ctx) error {
	tasks := taskRepository.FindAll()

	resp := models.Response{
		Code:    http.StatusOK,
		Body:    tasks,
		Title:   "GetAllTasks",
		Message: "All Tasks",
	}

	return c.Status(resp.Code).JSON(resp)
}

// GetSingleTask Gets single task information
func GetSingleTask(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 0)

	if err != nil {
		errorResp := models.Response{
			Code:    http.StatusNotAcceptable,
			Body:    err.Error(),
			Title:   "NotAcceptable",
			Message: "Error in getting task information",
		}

		return c.Status(errorResp.Code).JSON(errorResp)
	}

	task, err := taskRepository.FindByID(uint64(id))
	if err != nil {
		errorResp := models.Response{
			Code:    http.StatusNotFound,
			Body:    err.Error(),
			Title:   "NotFound",
			Message: "Error in getting task information",
		}

		return c.Status(errorResp.Code).JSON(errorResp)
	}

	if task == nil {
		errorResp := models.Response{
			Code:    http.StatusNotFound,
			Body:    fmt.Sprintf("task with id %d could not be found", id),
			Title:   "NotFound",
			Message: "Error in finding task",
		}

		return c.Status(errorResp.Code).JSON(errorResp)
	}

	resp := models.Response{
		Code:    http.StatusOK,
		Body:    task,
		Title:   "OK",
		Message: "Task information",
	}
	return c.Status(resp.Code).JSON(resp)

}

// AddNewTask adds new task
func AddNewTask(c *fiber.Ctx) error {
	task := &models.Task{}

	err := c.BodyParser(task)

	if err != nil {
		errorResp := models.Response{
			Code:    http.StatusNotAcceptable,
			Body:    err.Error(),
			Title:   "Error",
			Message: "Error in parsing task body information",
		}

		return c.Status(errorResp.Code).JSON(errorResp)
	}

	id, err := taskRepository.Save(*task)
	if err != nil {
		errorResp := models.Response{
			Code:    http.StatusInternalServerError,
			Body:    err.Error(),
			Title:   "InternalServerError",
			Message: "Error in adding new task",
		}

		return c.Status(errorResp.Code).JSON(errorResp)
	}

	task, err = taskRepository.FindByID(id)
	if err != nil {
		errorResp := models.Response{
			Code:    http.StatusInternalServerError,
			Body:    err.Error(),
			Title:   "InternalServerError",
			Message: "Error in finding newly added task",
		}

		return c.Status(errorResp.Code).JSON(errorResp)
	}
	if task == nil {
		errorResp := models.Response{
			Code:    http.StatusNotFound,
			Body:    fmt.Sprintf("task with id %d could not be found", id),
			Title:   "NotFound",
			Message: "Error in finding task",
		}

		return c.Status(errorResp.Code).JSON(errorResp)
	}

	resp := models.Response{
		Code:    http.StatusOK,
		Body:    task,
		Title:   "OK",
		Message: "new task added successfully",
	}
	return c.Status(resp.Code).JSON(resp)

}

// UpdateTask updates a task by task id
func UpdateTask(c *fiber.Ctx) error {
	task := &models.Task{}

	err := c.BodyParser(task)
	if err != nil {
		errorResp := models.Response{
			Code:    http.StatusNotAcceptable,
			Body:    err.Error(),
			Title:   "NotAcceptable",
			Message: "Error in parsing task body information",
		}

		return c.Status(errorResp.Code).JSON(errorResp)
	}

	id, err := strconv.ParseUint(c.Params("id"), 10, 0)
	if err != nil {
		errorResp := models.Response{
			Code:    http.StatusNotAcceptable,
			Body:    err.Error(),
			Title:   "NotAcceptable",
			Message: "Error in parsing task ID. (it should be an integer)",
		}

		return c.Status(errorResp.Code).JSON(errorResp)
	}

	updatingTask, err := taskRepository.FindByID(uint64(id))
	if err != nil {
		errorResp := models.Response{
			Code:    http.StatusNotFound,
			Body:    err.Error(),
			Title:   "NotFound",
			Message: "Error in getting task information",
		}

		return c.Status(errorResp.Code).JSON(errorResp)
	}

	if updatingTask == nil {
		errorResp := models.Response{
			Code:    http.StatusNotFound,
			Body:    fmt.Sprintf("task with id %d could not be found", id),
			Title:   "NotFound",
			Message: "Error in finding task",
		}

		return c.Status(errorResp.Code).JSON(errorResp)
	}

	task.ID = uint64(id)

	err = taskRepository.Update(*task)
	if err != nil {
		errorResp := models.Response{
			Code:    http.StatusInternalServerError,
			Body:    err.Error(),
			Title:   "InternalServerError",
			Message: "Error in updating task information",
		}

		return c.Status(errorResp.Code).JSON(errorResp)
	}

	task, err = taskRepository.FindByID(uint64(id))
	if err != nil {
		errorResp := models.Response{
			Code:    http.StatusInternalServerError,
			Body:    err.Error(),
			Title:   "InternalServerError",
			Message: "Error in finding newly updated task",
		}

		return c.Status(errorResp.Code).JSON(errorResp)
	}

	if task == nil {
		errorResp := models.Response{
			Code:    http.StatusNotFound,
			Body:    fmt.Sprintf("task with id %d could not be found", id),
			Title:   "NotFound",
			Message: "Error in finding task",
		}

		return c.Status(errorResp.Code).JSON(errorResp)
	}

	resp := models.Response{
		Code:    http.StatusOK,
		Body:    task,
		Title:   "UpdateTask",
		Message: "task updated successfully",
	}
	return c.Status(resp.Code).JSON(resp)
}

// DeleteTask deletes the task from db
func DeleteTask(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 0)

	if err != nil {
		errorResp := models.Response{
			Code:    http.StatusNotAcceptable,
			Body:    err.Error(),
			Title:   "Error",
			Message: "Error in getting task information",
		}

		return c.Status(errorResp.Code).JSON(errorResp)
	}

	task, err := taskRepository.FindByID(uint64(id))
	if err != nil {
		errorResp := models.Response{
			Code:    http.StatusInternalServerError,
			Body:    err.Error(),
			Title:   "InternalServerError",
			Message: "Error in finding task",
		}

		return c.Status(errorResp.Code).JSON(errorResp)
	}

	if task == nil {
		errorResp := models.Response{
			Code:    http.StatusNotFound,
			Body:    fmt.Sprintf("task with id %d could not be found", id),
			Title:   "NotFound",
			Message: "Error in finding task",
		}

		return c.Status(errorResp.Code).JSON(errorResp)
	}

	err = taskRepository.Delete(*task)
	if err != nil {
		errorResp := models.Response{
			Code:    http.StatusNotAcceptable,
			Body:    err.Error(),
			Title:   "NotAcceptable",
			Message: "Error in deleting task object",
		}

		return c.Status(errorResp.Code).JSON(errorResp)
	}

	resp := models.Response{
		Code:    http.StatusOK,
		Body:    "task deleted successfully",
		Title:   "OK",
		Message: "task deleted successfully",
	}
	return c.Status(resp.Code).JSON(resp)
}
