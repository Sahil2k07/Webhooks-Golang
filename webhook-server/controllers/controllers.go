package controllers

import (
	"github.com/Sahil2k07/Webhooks-Golang/webhook-server/services"
	"github.com/labstack/echo"
)

func InitControllers(e *echo.Echo) {
	e.POST("/api/payment", func(c echo.Context) error {
		ps := services.NewPaymentService(c)

		return ps.ProcessPayment()
	})
}
