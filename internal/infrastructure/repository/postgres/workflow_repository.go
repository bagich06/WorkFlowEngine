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
	query := `SELECT id, entity_id, step, status from workflows where id=$1;`

	wf := &workflow.Workflow{}
	err := repo.db.Pool.QueryRow(ctx, query, id).Scan(&wf.ID, &wf.EntityID, &wf.Step, &wf.Status)
	if err != nil {
		return nil, err
	}

	return wf, nil
}

func (repo *WorkFlowRepository) GetByEntityID(ctx context.Context, entityID int64) (*workflow.Workflow, error) {
	query := `SELECT id, entity_id, step, status from workflows where entity_id=$1;`

	wf := &workflow.Workflow{}
	err := repo.db.Pool.QueryRow(ctx, query, entityID).Scan(&wf.ID, &wf.EntityID, &wf.Step, &wf.Status)
	if err != nil {
		return nil, err
	}

	return wf, nil
}

func (repo *WorkFlowRepository) Create(ctx context.Context, wf *workflow.Workflow) error {
	query := `INSERT into workflows (entity_id, step, status) values ($1, $2, $3);`

	_, err := repo.db.Pool.Exec(ctx, query, wf.EntityID, wf.Step, wf.Status)
	if err != nil {
		return err
	}

	return nil
}

func (repo *WorkFlowRepository) Save(ctx context.Context, wf *workflow.Workflow) error {
	query := `UPDATE workflows SET step=$1, status=$2 WHERE id=$3;`

	_, err := repo.db.Pool.Exec(ctx, query, wf.Step, wf.Status, wf.ID)
	if err != nil {
		return err
	}

	return nil
}
