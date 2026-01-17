package postgres

import (
	"context"
	"workflow_engine/internal/domain/entities/workflow"
)

type WorkFlowRepository struct {
	db *DB
}

func NewWorkFlowRepository(db *DB) *WorkFlowRepository {
	return &WorkFlowRepository{db: db}
}

func (repo *WorkFlowRepository) GetByID(ctx context.Context, id int64) (*workflow.Workflow, error) {
	query := `SELECT entity_id, step, status from workflow where id=$1;`

	workflow := &workflow.Workflow{}
	err := repo.db.Pool.QueryRow(ctx, query, id).Scan(workflow.EntityID, workflow.Step, workflow.Status)
	if err != nil {
		return nil, err
	}

	return workflow, nil
}

func (repo *WorkFlowRepository) Save(ctx context.Context, wf *workflow.Workflow) error {
	query := `INSERT into workflow (entity_id, step, status) values ($1, $2, $3);`

	_, err := repo.db.Pool.Exec(ctx, query, wf.Step, wf.Status)
	if err != nil {
		return err
	}

	return nil
}
