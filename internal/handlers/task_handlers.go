package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"petProjetPantella/internal/taskservice"
)

type TaskHandler struct {
	service taskservice.TaskService
}

func NewTaskHandler(service taskservice.TaskService) *TaskHandler {
	return &TaskHandler{
		service: service,
	}
}

func (handler *TaskHandler) RegisterRoutes(e *echo.Echo) {
	e.GET("/tasks", handler.GetAll)
	e.POST("/tasks", handler.Create)
	e.DELETE("/tasks/:id", handler.Delete)
	e.PATCH("/tasks/:id", handler.Patch)
}

func (handler *TaskHandler) GetAll(c echo.Context) error {
	tasks, err := handler.service.GetTasks()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, tasks)
}

func (handler *TaskHandler) Create(c echo.Context) error {
	var req taskservice.CreateTaskRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "invalid request",
		})
	}
	task, err := handler.service.AddTask(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": "create failed",
		})
	}
	return c.JSON(http.StatusCreated, task)
}

func (handler *TaskHandler) Delete(c echo.Context) error {
	id := c.Param("id")
	if err := handler.service.DeleteTask(id); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": "delete failed",
		})
	}
	return c.NoContent(http.StatusNoContent)
}

func (handler *TaskHandler) Patch(c echo.Context) error {
	id := c.Param("id")
	var updates map[string]interface{}
	if err := c.Bind(&updates); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "invalid update body",
		})
	}
	task, err := handler.service.PatchTask(id, updates)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": "update failed",
		})
	}
	return c.JSON(http.StatusOK, task)
}
