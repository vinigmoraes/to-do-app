package core

import (
	"github.com/google/uuid"
	"to-do-app/item/core/domain"
	"to-do-app/item/core/ports"
)

type ItemService struct {
	Repository ports.ItemRepository
}

func (s *ItemService) CreateItem(dto ports.CreateItemDTO) uuid.UUID {
	item := domain.NewItem(dto.Input, dto.Category)

	s.Repository.Insert(item)

	return item.ID
}