package document

import "workflow_engine/internal/domain/entities"

type DocumentResponse struct {
	ID     int64                   `json:"id"`
	Topic  string                  `json:"topic"`
	Status entities.DocumentStatus `json:"status"`
}
