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

var task = RequestBody{
	Task: "Изначальное",
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

func GetTask(e echo.Context) error {
	return e.JSON(http.StatusOK, "Hello "+task.Task)
}

func UpdateTask(e echo.Context) error {
	var req RequestBody
	if err := json.NewDecoder(e.Request().Body).Decode(&req); err != nil {
		return e.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid request body",
		})
	}
	task = req
	return e.JSON(http.StatusOK, map[string]string{
		"status": "task updated",
	})

}
