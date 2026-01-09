package postgres

import (
	"context"
	"workflow_engine/internal/domain/entities"
)

type DocumentRepository struct {
	db *DB
}

func NewDocumentRepository(db *DB) *DocumentRepository {
	return &DocumentRepository{db: db}
}

func (repo *DocumentRepository) Create(ctx context.Context, document *entities.Document) (*entities.Document, error) {
	query := `INSERT INTO documents (topic, status) VALUES ($1, $2) RETURNING id`

	err := repo.db.Pool.QueryRow(ctx, query, document.Topic, document.Status).Scan(&document.ID)
	if err != nil {
		return nil, err
	}

	return document, nil
}

func (repo *DocumentRepository) GetByID(ctx context.Context, id int64) (*entities.Document, error) {
	query := `SELECT id, topic, status FROM documents WHERE id = $1`

	document := &entities.Document{}
	err := repo.db.Pool.QueryRow(ctx, query, id).Scan(&document.ID, &document.Topic, &document.Status)
	if err != nil {
		return nil, err
	}

	return document, nil
}

func (repo *DocumentRepository) UpdateStatus(ctx context.Context, newStatus entities.DocumentStatus, id int64) error {
	query := `UPDATE documents SET status = $1 WHERE id = $2`

	_, err := repo.db.Pool.Exec(ctx, query, newStatus, id)
	if err != nil {
		return err
	}

	return nil
}

func (repo *DocumentRepository) GetStatusByID(ctx context.Context, id int64) (entities.DocumentStatus, error) {
	query := `SELECT status FROM documents WHERE id = $1`

	var status entities.DocumentStatus

	err := repo.db.Pool.QueryRow(ctx, query, id).Scan(&status)
	if err != nil {
		return "", err
	}

	return status, nil
}
