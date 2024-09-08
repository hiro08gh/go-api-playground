package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetUser(c echo.Context) error {
	return c.JSON(http.StatusOK, "Get User")
}