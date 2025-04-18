package main

import (
	"zad4/config"
	"zad4/routes"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	config.ConnectDB()
	routes.RegisterRoutes(e)

	e.Logger.Fatal(e.Start(":8080"))
}