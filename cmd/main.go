package main

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"strings"
)

type RequestBody struct {
	Task string `json:"task"`
}

var taskList = []RequestBody{
	{Task: "Start"},
}

func main() {
	e := echo.New()
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	e.GET("/task", GetTask)
	e.POST("/update", UpdateTask)

	err := e.Start("localhost:8080")
	if err != nil {
		return
	}
}

func GetTask(c echo.Context) error {
	var tasks []string
	for _, t := range taskList {
		tasks = append(tasks, t.Task)
	}
	return c.String(http.StatusOK, "Hello "+strings.Join(tasks, ", "))
}

func UpdateTask(c echo.Context) error {
	var req RequestBody
	if err := json.NewDecoder(c.Request().Body).Decode(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid request body",
		})
	}
	taskList = append(taskList, req)
	return c.JSON(http.StatusOK, map[string]string{
		"status": "task added",
	})
}
