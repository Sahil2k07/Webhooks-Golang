package controllers

import (
	"log"
	"net/http"

	"github.com/labstack/echo"
)

func InitControllers(e *echo.Echo) {
	e.PUT("/api/account/upgrade", func(c echo.Context) error {
		log.Println("Received webhook call")

		return c.String(http.StatusOK, "successfull")
	})
}
