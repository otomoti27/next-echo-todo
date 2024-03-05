package main

import (
	"echo-api/db"
	"echo-api/domain"
	"echo-api/internal/repository"
	"echo-api/internal/rest"
	"echo-api/internal/rest/middleware"
	"echo-api/service/user"

	"github.com/labstack/echo/v4"
)

func main() {
	dbConn := db.Init()
	dbConn.AutoMigrate(&domain.User{})

	e := echo.New()
	e.Use(middleware.CORS())
	e.Use(middleware.CSRF())

	userRepo := repository.NewUserRepository(dbConn)
	userService := user.NewService(userRepo)
	rest.NewUserHandler(e, userService)

	e.Logger.Fatal(e.Start(":8080"))
}
