package main

import (
	"github.com/labstack/echo"
	"to-do-app/item/application"
)

func main() {
	server := echo.New()

	application.CreateItemsRoutes(server)

	server.Logger.Fatal(server.Start(":8080"))
}
