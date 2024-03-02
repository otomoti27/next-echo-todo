package main

import (
	"echo-api/db"
	"echo-api/domain"

	"github.com/labstack/echo/v4"
)

func main() {
	dbConn := db.Init()
	dbConn.AutoMigrate(&domain.User{})

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(200, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":8080"))
}
