package storage

import (
	"context"

	"github.com/google/uuid"
	"gitlab.ozon.dev/Vanek623/task-manager-system/external/counters"

	"gitlab.ozon.dev/Vanek623/task-manager-system/external/task/models"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4/pgxpool"
)

type postgres struct {
	pool *pgxpool.Pool
	cs   *counters.Counters
}

func (p *postgres) Add(ctx context.Context, t *models.Task) error {
	const query = "INSERT INTO tasks (id, title, description) VALUES ($1, $2, $3) RETURNING id"

	p.cs.Inc(counters.Outbound)
	row := p.pool.QueryRow(ctx, query, t.ID, t.Title, t.Description)

	var id uuid.UUID
	if err := row.Scan(&id); err != nil {
		return err
	}

	return nil
}

func (p *postgres) Delete(ctx context.Context, ID *uuid.UUID) error {
	const query = "DELETE FROM tasks WHERE id = $1 RETURNING id"

	p.cs.Inc(counters.Outbound)
	res := p.pool.QueryRow(ctx, query, *ID)
	if res.Scan(ID) != nil {
		return ErrTaskNotExist
	}

	return nil
}

func (p *postgres) List(ctx context.Context, limit, offset uint64) ([]*models.Task, error) {
	const query = "SELECT * FROM tasks ORDER BY id LIMIT $1 OFFSET $2"

	p.cs.Inc(counters.Outbound)
	var tasks []*models.Task
	if err := pgxscan.Select(ctx, p.pool, &tasks, query, limit, offset); err != nil {
		return nil, err
	}

	return tasks, nil
}

func (p *postgres) Update(ctx context.Context, t *models.Task) error {
	const query = "UPDATE tasks SET title = $1, description = $2, edited = now() WHERE id = $3 RETURNING id"

	p.cs.Inc(counters.Outbound)
	res := p.pool.QueryRow(ctx, query, t.Title, t.Description, t.ID)
	var tmp uuid.UUID
	if res.Scan(&tmp) != nil {
		return ErrTaskNotExist
	}

	return nil
}

func (p *postgres) Get(ctx context.Context, ID *uuid.UUID) (*models.Task, error) {
	const query = "SELECT * FROM tasks WHERE id = $1"

	p.cs.Inc(counters.Outbound)
	var task models.Task
	if err := pgxscan.Get(ctx, p.pool, &task, query, *ID); err != nil {
		return nil, err
	}

	return &task, nil
}
