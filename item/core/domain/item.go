package domain

import (
	"time"

	"github.com/google/uuid"
)

type Item struct {
	ID          uuid.UUID
	Input       string
	Category    string
	completed   bool
	createdAt   time.Time
	completedAt time.Time
}

func NewItem(input string, category string) *Item {
	return &Item{
		ID:          uuid.New(),
		Input:       input,
		Category:    category,
		completed:   false,
		createdAt:   time.Now(),
	}
}
