package main

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
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

	e.GET("/tasks", GetTask)
	e.POST("/tasks", UpdateTask)

	err := e.Start("localhost:8080")
	if err != nil {
		return
	}
}

func GetTask(c echo.Context) error {
	return c.JSON(http.StatusOK, taskList)
}

func UpdateTask(c echo.Context) error {
	var req RequestBody
	if err := json.NewDecoder(c.Request().Body).Decode(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid request body",
		})
	}
	taskList = append(taskList, req)
	return c.JSON(http.StatusOK, req)
}
