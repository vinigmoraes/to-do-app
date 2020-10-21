package request

import "to-do-app/item/core/ports"

type CreateItemRequest struct {
	Input    string `json:"input"`
	Category string `json:"category"`
}

func (request *CreateItemRequest) CreateItemRequestDTO() ports.CreateItemDTO {
	return ports.CreateItemDTO{Input: request.Input, Category: request.Category}
}
