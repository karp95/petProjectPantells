package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"petProjetPantella/internal/database"
	"petProjetPantella/internal/handlers"
	"petProjetPantella/internal/taskservice"
	"petProjetPantella/internal/web/tasks"
)

func main() {
	database.InitDb()
	database.Db.AutoMigrate(&taskservice.Task{})

	repo := taskservice.NewTaskRepository(database.Db)
	service := taskservice.NewTaskService(repo)
	hand := handlers.NewTaskHandler(service)

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	strictHandler := tasks.NewStrictHandler(hand, nil) // тут будет ошибка
	tasks.RegisterHandlers(e, strictHandler)

	if err := e.Start(":8080"); err != nil {
		log.Fatal("Server failed:", err)
	}
}
