package main

import (
	"github.com/Sahil2k07/Webhooks-Golang/webhook-server/controllers"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	controllers.InitControllers(e)

	e.Logger.Fatal(e.Start(":4000"))
}
