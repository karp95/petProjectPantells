package main

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

type Task struct {
	ID     string `json:"id"`
	Task   string `json:"task"`
	Status string `json:"status"`
}

type CreateTaskRequest struct {
	Task   string `json:"task" validate:"required"`
	Status string `json:"status" validate:"required"`
}

var taskList = []Task{
	{Task: "Start", Status: "Running", ID: uuid.NewString()},
	{Task: "Stop", Status: "Done", ID: uuid.NewString()},
}

func main() {
	e := echo.New()
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())
	e.GET("/tasks", GetTask)
	e.POST("/tasks", AddTask)
	e.DELETE("/tasks/:id", DeleteTask)
	e.PATCH("/tasks/:id", PatchTask)

	err := e.Start("localhost:8080")
	if err != nil {
		panic("failed to start server")
	}
}

func GetTask(c echo.Context) error {
	return c.JSON(http.StatusOK, taskList)
}

func AddTask(c echo.Context) error {
	var task CreateTaskRequest

	decoder := json.NewDecoder(c.Request().Body)
	decoder.DisallowUnknownFields()

	err := decoder.Decode(&task)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "Invalid request body: " + err.Error(),
		})
	}

	newTask := Task{
		ID:     uuid.NewString(),
		Task:   task.Task,
		Status: task.Status,
	}
	taskList = append(taskList, newTask)

	return c.JSON(http.StatusCreated, newTask)
}

func DeleteTask(c echo.Context) error {
	urlId := c.Param("id")

	for i, task := range taskList {
		if task.ID == urlId {
			taskList = append(taskList[:i], taskList[i+1:]...)
			return c.NoContent(http.StatusNoContent)
		}
	}
	return echo.NewHTTPError(http.StatusNotFound, "Task not found")
}

func PatchTask(c echo.Context) error {
	UrlId := c.Param("id")

	var TaskUpdate struct {
		Task   *string `json:"task"`
		Status *string `json:"status"`
	}

	if err := c.Bind(&TaskUpdate); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "invalid request body",
		})
	}
	for i, task := range taskList {
		if task.ID == UrlId {
			if TaskUpdate.Task != nil {
				taskList[i].Task = *TaskUpdate.Task
			}
			if TaskUpdate.Status != nil {
				taskList[i].Status = *TaskUpdate.Status
			}

			return c.JSON(http.StatusOK, taskList[i])
		}
	}
	return echo.NewHTTPError(http.StatusNotFound, "Task not found")
}
