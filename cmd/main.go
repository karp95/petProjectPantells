package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"petProjetPantella/internal/db"
	"petProjetPantella/internal/handlers"
	"petProjetPantella/internal/taskservice"
)

func main() {
	database, err := db.InitDb()
	if err != nil {
		log.Fatal(err)
	}

	repo := taskservice.NewTaskRepository(database)
	service := taskservice.NewTaskService(repo)
	hand := handlers.NewTaskHandler(service)

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())

	hand.RegisterRoutes(e)

	if err := e.Start(":8080"); err != nil {
		log.Fatal("Server failed:", err)
	}
}
