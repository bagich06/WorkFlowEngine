package document

import "workflow_engine/internal/domain/entities"

type DocumentResponse struct {
	Topic  string                  `json:"topic"`
	Status entities.DocumentStatus `json:"status"`
}

type UpdateDocumentResponse struct {
	RespStr string                  `json:"response_string"`
	Status  entities.DocumentStatus `json:"status"`
}
