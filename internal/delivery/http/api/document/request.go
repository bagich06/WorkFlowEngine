package document

import "workflow_engine/internal/domain/entities"

type DocumentCreateRequest struct {
	Topic  string                  `json:"topic" validate:"required"`
	Status entities.DocumentStatus `json:"status" validate:"required"`
}
