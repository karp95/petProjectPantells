package main

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
)

var db *gorm.DB

func initDb() {
	dsn := "host=localhost user=postgres password=my_pass dbname=postgres port=5432 sslmode=disable"
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database", err.Error())
	}

	err = db.AutoMigrate(&Task{})
	if err != nil {
		log.Fatal("Failed to migrate database", err.Error())
	}
}

type Task struct {
	ID     string `json:"id" gorm:"primaryKey"`
	Task   string `json:"task"`
	Status string `json:"status"`
}

type CreateTaskRequest struct {
	Task   string `json:"task" validate:"required"`
	Status string `json:"status" validate:"required"`
}

func main() {
	initDb()

	e := echo.New()
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	e.GET("/tasks", GetTask)
	e.POST("/tasks", AddTask)
	e.DELETE("/tasks/:id", DeleteTask)
	e.PATCH("/tasks/:id", PatchTask)

	err := e.Start("localhost:8080")
	if err != nil {
		log.Fatal("Failed to start server", err.Error())
	}
}

func GetTask(c echo.Context) error {
	var task []Task

	err := db.Find(&task).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": "Could not get task",
		})
	}
	return c.JSON(http.StatusOK, task)
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

	err = db.Create(&newTask).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": "Could not add task",
		})
	}

	return c.JSON(http.StatusCreated, newTask)
}

func DeleteTask(c echo.Context) error {
	urlId := c.Param("id")
	err := db.Where("id = ?", urlId).Delete(&Task{}).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": "Could not delete task",
		})
	}
	return echo.NewHTTPError(http.StatusNoContent)
}

func PatchTask(c echo.Context) error {
	urlId := c.Param("id")

	var TaskUpdate struct {
		Task   *string `json:"task"`
		Status *string `json:"status"`
	}

	if err := c.Bind(&TaskUpdate); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "invalid request body",
		})
	}
	err := db.Model(&Task{}).Where("id = ?", urlId).Updates(&TaskUpdate).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": "update failed",
		})
	}
	return echo.NewHTTPError(http.StatusOK, TaskUpdate)
}
