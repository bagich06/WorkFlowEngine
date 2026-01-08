package document

import "workflow_engine/internal/domain/entities"

type DocumentCreateRequest struct {
	Topic  string                  `json:"topic"`
	Status entities.DocumentStatus `json:"status"`
}

type DocumentUpdateRequest struct {
	Status entities.DocumentStatus `json:"status"`
}
