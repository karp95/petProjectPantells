package handlers

import (
	"context"
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
	"petProjetPantella/internal/taskservice"
	"petProjetPantella/internal/web/tasks"
)

type TaskHandler struct {
	service taskservice.TaskService
}

func (handler *TaskHandler) DeleteTasksId(ctx context.Context, request tasks.DeleteTasksIdRequestObject) (tasks.DeleteTasksIdResponseObject, error) {
	err := handler.service.DeleteTask(request.Id)
	if err != nil {
		if errors.Is(err, taskservice.ErrNotFound) {
			return tasks.DeleteTasksId404JSONResponse{
				Message: "Task not found",
			}, nil
		}
		return nil, err
	}

	return tasks.DeleteTasksId204Response{}, nil
}

func (handler *TaskHandler) PatchTasksId(ctx context.Context, request tasks.PatchTasksIdRequestObject) (tasks.PatchTasksIdResponseObject, error) {
	if request.Body == nil {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "empty body")
	}
	updates := make(map[string]interface{})
	updates["task"] = request.Body.Task
	updates["is_done"] = request.Body.IsDone
	if len(updates) == 0 {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "no fields to update")
	}
	updatedTask, err := handler.service.PatchTask(request.Id, updates)
	if err != nil {
		if errors.Is(err, taskservice.ErrNotFound) {
			return tasks.PatchTasksId404JSONResponse{
				Message: "Task not found",
			}, nil
		}
		return nil, err
	}
	resp := tasks.PatchTasksId200JSONResponse{
		Id:     &updatedTask.ID,
		Task:   &updatedTask.Task,
		IsDone: &updatedTask.IsDone,
	}
	return resp, nil
}

func (handler *TaskHandler) GetTasks(ctx context.Context, request tasks.GetTasksRequestObject) (tasks.GetTasksResponseObject, error) {
	allTasks, err := handler.service.GetTasks()
	if err != nil {
		return nil, err
	}
	response := tasks.GetTasks200JSONResponse{}
	for _, tsk := range allTasks {
		task := tasks.Task{
			Id:     &tsk.ID,
			Task:   &tsk.Task,
			IsDone: &tsk.IsDone,
		}
		response = append(response, task)
	}
	return response, nil
}

func (handler *TaskHandler) PostTasks(ctx context.Context, request tasks.PostTasksRequestObject) (tasks.PostTasksResponseObject, error) {
	if request.Body == nil {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "empty body")
	}
	created, err := handler.service.AddTask(taskservice.CreateTaskRequest{
		Task:   request.Body.Task,
		IsDone: request.Body.IsDone,
	})
	if err != nil {
		return nil, err
	}
	resp := tasks.PostTasks201JSONResponse{
		Id:     &created.ID,
		Task:   &created.Task,
		IsDone: &created.IsDone,
	}
	return resp, nil
}

func NewTaskHandler(service taskservice.TaskService) *TaskHandler {
	return &TaskHandler{
		service: service,
	}
}
