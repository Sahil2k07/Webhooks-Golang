package controllers

import (
	"github.com/Sahil2k07/Webhooks-Golang/main-server/services"
	"github.com/labstack/echo"
)

func InitControllers(e *echo.Echo) {
	e.PUT("/api/account/upgrade/webhook", func(c echo.Context) error {
		as := services.NewAccountService(c)

		return as.UpgradeUserAccount()
	})

	e.POST("api/account/upgrade", func(c echo.Context) error {
		as := services.NewAccountService(c)

		return as.MakeUserPayment()
	})
}
