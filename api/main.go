package main

import (
	"github.com/hiro08gh/go-api-playground/api/handlers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.CORS())
	e.Use(middleware.Recover())

	e.GET("/users", handlers.GetUser)
	e.POST("/login", handlers.Login)
	e.POST("/proteced", handlers.Protected)
	e.Logger.Fatal(e.Start(":3001"))
}