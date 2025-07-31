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
		return
	}
}

func GetTask(c echo.Context) error {
	return c.JSON(http.StatusOK, taskList)
}

func AddTask(c echo.Context) error {
	var req Task
	if err := json.NewDecoder(c.Request().Body).Decode(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid request body",
		})
	}
	req.ID = uuid.NewString()
	taskList = append(taskList, req)
	return c.JSON(http.StatusOK, req)
}

func DeleteTask(c echo.Context) error {
	id := c.Param("id")

	for i, t := range taskList {
		if t.ID == id {
			taskList = append(taskList[:i], taskList[i+1:]...)
			return c.NoContent(http.StatusNoContent)
		}
	}

	return c.JSON(http.StatusNotFound, map[string]string{
		"error": "task not found",
	})
}

func PatchTask(c echo.Context) error {
	id := c.Param("id")

	var updateData struct {
		ID     *string `json:"id"`
		Task   *string `json:"task"`
		Status *string `json:"status"`
	}

	if err := c.Bind(&updateData); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid request body",
		})
	}
	for i, t := range taskList {
		if t.ID == id {
			if updateData.ID != nil {
				taskList[i].ID = *updateData.ID
			}
			if updateData.Task != nil {
				taskList[i].Task = *updateData.Task
			}
			if updateData.Status != nil {
				taskList[i].Status = *updateData.Status
			}

			return c.JSON(http.StatusOK, taskList[i])
		}
	}
	return c.JSON(http.StatusBadRequest, map[string]string{
		"error": "task not found",
	})
}
