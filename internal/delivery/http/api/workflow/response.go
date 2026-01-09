package workflow

import "workflow_engine/internal/domain/entities"

type WorkflowResponse struct {
	Message string                  `json:"message"`
	Status  entities.DocumentStatus `json:"status"`
}

type StatusResponse struct {
	DocumentID int64                   `json:"document_id"`
	Status     entities.DocumentStatus `json:"status"`
}
