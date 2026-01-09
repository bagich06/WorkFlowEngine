package postgres

import (
	"context"
	"errors"
	"fmt"
	"workflow_engine/internal/domain/entities"

	"github.com/jackc/pgx/v4"
)

type UserRepository struct {
	db *DB
}

func NewUserRepository(db *DB) *UserRepository {
	return &UserRepository{db: db}
}

func (repo *UserRepository) Create(ctx context.Context, user *entities.User) (*entities.User, error) {
	query := `
		INSERT INTO users (first_name, last_name, phone, password, role)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id
	`

	err := repo.db.Pool.QueryRow(ctx, query, user.FirstName, user.LastName, user.Phone, user.Password, user.Role).Scan(&user.ID)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (repo *UserRepository) GetByPhone(ctx context.Context, phone string) (*entities.User, error) {
	query := `
		SELECT id, first_name, last_name, phone, password, role 
		FROM users 
		WHERE phone = $1
	`

	user := &entities.User{}
	err := repo.db.Pool.QueryRow(ctx, query, phone).Scan(
		&user.ID, &user.FirstName, &user.LastName, &user.Phone, &user.Password, &user.Role,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil // Пользователь не найден — это не ошибка
		}
		return nil, fmt.Errorf("failed to get user by phone: %w", err)
	}

	return user, nil
}

func (repo *UserRepository) GetByID(ctx context.Context, id int64) (*entities.User, error) {
	query := `
		SELECT id, first_name, last_name, phone, password, role 
		FROM users 
		WHERE id = $1
	`

	user := &entities.User{}
	err := repo.db.Pool.QueryRow(ctx, query, id).Scan(
		&user.ID, &user.FirstName, &user.LastName, &user.Phone, &user.Password, &user.Role,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errors.New("user not found")
		}
		return nil, fmt.Errorf("failed to get user by id: %w", err)
	}

	return user, nil
}

func (repo *UserRepository) GetByRole(ctx context.Context, role string) ([]*entities.User, error) {
	query := `
		SELECT id, first_name, last_name, phone, password, role 
		FROM users 
		WHERE role = $1
	`

	rows, err := repo.db.Pool.Query(ctx, query, role)
	if err != nil {
		return nil, fmt.Errorf("failed to get users by role: %w", err)
	}
	defer rows.Close()

	var users []*entities.User
	for rows.Next() {
		user := &entities.User{}
		err := rows.Scan(
			&user.ID, &user.FirstName, &user.LastName, &user.Phone, &user.Password, &user.Role,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan user: %w", err)
		}
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over users: %w", err)
	}

	return users, nil
}

func (repo *UserRepository) GetAll(ctx context.Context) ([]*entities.User, error) {
	query := `
		SELECT id, first_name, last_name, phone, password, role 
		FROM users
	`

	rows, err := repo.db.Pool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to get all users: %w", err)
	}
	defer rows.Close()

	var users []*entities.User
	for rows.Next() {
		user := &entities.User{}
		err := rows.Scan(
			&user.ID, &user.FirstName, &user.LastName, &user.Phone, &user.Password, &user.Role,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan user: %w", err)
		}
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over users: %w", err)
	}

	return users, nil
}

func (repo *UserRepository) Delete(ctx context.Context, id int64) error {
	query := `DELETE FROM users WHERE id = $1`

	commandTag, err := repo.db.Pool.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}

	if commandTag.RowsAffected() == 0 {
		return errors.New("user not found")
	}

	return nil
}
