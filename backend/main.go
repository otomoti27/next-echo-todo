package main

import (
	"echo-api/db"

	"github.com/labstack/echo/v4"
)

func main() {
	db.Init()

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(200, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":8080"))
}
