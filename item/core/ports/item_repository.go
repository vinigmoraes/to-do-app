package ports

import "to-do-app/item/core/domain"

type ItemRepository interface {

	Insert(item *domain.Item)
}

type UserRepositoryAdapter struct {
	items []domain.Item
}

func (u UserRepositoryAdapter) Insert(item domain.Item) {
	u.items = append(u.items, item)
}