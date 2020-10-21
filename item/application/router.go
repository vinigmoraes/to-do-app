package application

import (
	"net/http"

	"github.com/labstack/echo"
)

func CreateItemsRoutes(e *echo.Echo) {
	controller := ItemController{}

	e.GET("/items", func(c echo.Context) error {
		return c.String(http.StatusOK, "Olá, World!")
	})

	e.POST("/items", func(c echo.Context) error {
		return controller.CreateItem(c)
	})
}
