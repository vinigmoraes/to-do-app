package application

import (
	"github.com/labstack/echo"
	"net/http"
	"to-do-app/item/application/request"
	"to-do-app/item/application/response"
	"to-do-app/item/core"
)

type ItemController struct {
	service core.ItemService
}

type Items []request.CreateItemRequest

func (c *ItemController) CreateItem(ctx echo.Context) error {
	item := request.CreateItemRequest{}
	ctx.Bind(&item)

	itemId := c.service.CreateItem(item.CreateItemRequestDTO())

	return ctx.JSON(http.StatusCreated, response.CreateItemResponse{Id: itemId})
}
