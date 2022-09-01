package storage

import (
	"context"

	"github.com/google/uuid"
	"gitlab.ozon.dev/Vanek623/task-manager-system/external/counters"

	"github.com/jmoiron/sqlx"

	"gitlab.ozon.dev/Vanek623/task-manager-system/external/task/models"
)

type sqlxDb struct {
	db *sqlx.DB
	cs *counters.Counters
}

func (p *sqlxDb) Add(_ context.Context, t *models.Task) error {
	const query = "INSERT INTO tasks (id, title, description) VALUES ($1, $2, $3) RETURNING id"

	p.cs.Inc(counters.Outbound)
	row := p.db.QueryRow(query, t.ID, t.Title, t.Description)

	var id uuid.UUID
	if err := row.Scan(&id); err != nil {
		return err
	}

	return nil
}

func (p *sqlxDb) Delete(_ context.Context, ID *uuid.UUID) error {
	const query = "DELETE FROM tasks WHERE id = $1 RETURNING id"

	p.cs.Inc(counters.Outbound)
	res := p.db.QueryRow(query, *ID)
	if res.Scan(ID) != nil {
		return ErrTaskNotExist
	}

	return nil
}

func (p *sqlxDb) List(_ context.Context, limit, offset uint64) ([]*models.Task, error) {
	const query = "SELECT * FROM tasks ORDER BY id LIMIT $1 OFFSET $2"

	var tasks []*models.Task

	p.cs.Inc(counters.Outbound)
	if err := sqlx.Select(p.db, &tasks, query, limit, offset); err != nil {
		return nil, err
	}

	return tasks, nil
}

func (p *sqlxDb) Update(_ context.Context, t *models.Task) error {
	const query = "UPDATE tasks SET title = $1, description = $2, edited = NOW() WHERE id = $3 RETURNING id"

	p.cs.Inc(counters.Outbound)
	res := p.db.QueryRow(query, t.Title, t.Description, t.ID)
	var tmp uuid.UUID
	if res.Scan(&tmp) != nil {
		return ErrTaskNotExist
	}

	return nil
}

func (p *sqlxDb) Get(_ context.Context, ID *uuid.UUID) (*models.Task, error) {
	const query = "SELECT * FROM tasks WHERE id = $1"

	p.cs.Inc(counters.Outbound)
	var task models.Task
	if err := sqlx.Get(p.db, &task, query, *ID); err != nil {
		return nil, err
	}

	return &task, nil
}
