package main

import (
	"github.com/hiro08gh/go-api-sample/api/handlers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.CORS())

	e.GET("/users", handlers.GetUser)
	e.Logger.Fatal(e.Start(":3001"))
}