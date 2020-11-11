package application

import (
	"net/http"
	"to-do-app/item/core"
	"to-do-app/item/core/ports"

	"github.com/labstack/echo"
)

func CreateItemsRoutes(e *echo.Echo) {
	repository := ports.ItemRepositoryAdapter{}
	service := core.ItemService{Repository: repository}
	controller := ItemController{service}

	e.GET("/items", func(c echo.Context) error {
		return c.String(http.StatusOK, "Olá, World!")
	})

	e.POST("/items", func(c echo.Context) error {
		return controller.CreateItem(c)
	})
}
