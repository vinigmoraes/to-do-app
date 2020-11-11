package ports

import "to-do-app/item/core/domain"

type ItemRepository interface {

	Insert(item *domain.Item)
}

type ItemRepositoryAdapter struct {
	items []*domain.Item
}

func (adapter ItemRepositoryAdapter) Insert(item *domain.Item) {
	adapter.items = append(adapter.items, item)
}