package core

import (
	"github.com/google/uuid"
	"to-do-app/item/core/domain"
	"to-do-app/item/core/ports"
)

type ItemService struct {
	repository ports.ItemRepository
}

func (s ItemService) CreateItem(dto ports.CreateItemDTO) uuid.UUID {
	item := domain.NewItem(dto.Input, dto.Category)

	s.repository.Insert(item)

	return item.ID
}