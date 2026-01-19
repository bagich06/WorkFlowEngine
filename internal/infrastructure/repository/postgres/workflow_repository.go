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
	query := `SELECT id, entity_id, "group", roles_status, groups_status, status from workflows where id=$1;`

	wf := &workflow.Workflow{}
	err := repo.db.Pool.QueryRow(ctx, query, id).Scan(&wf.ID, &wf.EntityID, &wf.Group, &wf.RolesStatus, &wf.GroupsStatus, &wf.Status)
	if err != nil {
		return nil, err
	}

	return wf, nil
}

func (repo *WorkFlowRepository) GetByEntityID(ctx context.Context, entityID int64) (*workflow.Workflow, error) {
	query := `SELECT id, entity_id, "group", roles_status, groups_status, status from workflows where entity_id=$1;`

	wf := &workflow.Workflow{}
	err := repo.db.Pool.QueryRow(ctx, query, entityID).Scan(&wf.ID, &wf.EntityID, &wf.Group, &wf.RolesStatus, &wf.GroupsStatus, &wf.Status)
	if err != nil {
		return nil, err
	}

	return wf, nil
}

func (repo *WorkFlowRepository) Create(ctx context.Context, wf *workflow.Workflow) error {
	query := `INSERT into workflows (entity_id, "group", roles_status, groups_status, status) values ($1, $2, $3, $4, $5);`

	_, err := repo.db.Pool.Exec(ctx, query, wf.EntityID, wf.Group, wf.RolesStatus, wf.GroupsStatus, wf.Status)
	if err != nil {
		return err
	}

	return nil
}

func (repo *WorkFlowRepository) Save(ctx context.Context, wf *workflow.Workflow) error {
	query := `UPDATE workflows SET "group"=$1, roles_status=$2, groups_status=$3, status=$4 WHERE id=$5;`

	_, err := repo.db.Pool.Exec(ctx, query, wf.Group, wf.RolesStatus, wf.GroupsStatus, wf.Status, wf.ID)
	if err != nil {
		return err
	}

	return nil
}
