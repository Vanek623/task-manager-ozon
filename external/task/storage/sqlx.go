package storage

import (
	"context"
	"github.com/jmoiron/sqlx"

	"gitlab.ozon.dev/Vanek623/task-manager-system/external/task/models"
)

type repository struct {
	db *sqlx.DB
}

func (p *repository) Add(ctx context.Context, t *models.Task) (uint64, error) {
	const query = "INSERT INTO tasks (title, description) VALUES ($1, $2) RETURNING id"

	row := p.db.QueryRow(query, t.Title, t.Description)

	var id uint64
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (p *repository) Delete(ctx context.Context, ID uint64) error {
	const query = "DELETE FROM tasks WHERE id = $1 RETURNING 1"

	res := p.db.QueryRow(query, ID)
	if res.Scan(&ID) != nil {
		return ErrTaskNotExist
	}

	return nil
}

func (p *repository) List(ctx context.Context, limit, offset uint64) ([]*models.Task, error) {
	const query = "SELECT * FROM tasks ORDER BY id LIMIT $1 OFFSET $2"

	var tasks []*models.Task

	if err := sqlx.Select(p.db, &tasks, query, limit, offset); err != nil {
		return nil, err
	}

	return tasks, nil
}

func (p *repository) Update(ctx context.Context, t *models.Task) error {
	const query = "UPDATE tasks SET title = $1, description = $2, edited = NOW() WHERE id = $3 RETURNING 1"

	res := p.db.QueryRow(query, t.Title, t.Description, t.ID)
	var tmp int
	if res.Scan(&tmp) != nil {
		return ErrTaskNotExist
	}

	return nil
}

func (p *repository) Get(ctx context.Context, ID uint64) (*models.Task, error) {
	const query = "SELECT * FROM tasks WHERE id = $1"

	var task models.Task
	if err := sqlx.Get(p.db, &task, query, ID); err != nil {
		return nil, err
	}

	return &task, nil
}
