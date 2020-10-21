package response

import "github.com/google/uuid"

type CreateItemResponse struct {
	Id uuid.UUID `json:"id"`
}