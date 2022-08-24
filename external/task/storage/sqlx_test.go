package storage

import (
	"context"
	"regexp"
	"testing"

	"github.com/google/uuid"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gitlab.ozon.dev/Vanek623/task-manager-system/external/task/models"
)

func TestSqlxDb_Add(t *testing.T) {
	t.Run("normal", func(t *testing.T) {
		// arrange
		f := setUp(t)
		defer f.tearDown()

		query := regexp.QuoteMeta(`INSERT INTO tasks (id, title, description) VALUES ($1, $2, $3) RETURNING id`)

		id := uuid.New()
		rows := sqlmock.NewRows([]string{"id"}).AddRow(id)
		f.dbMock.ExpectQuery(query).
			WithArgs(id, "test_t", "test_d").
			WillReturnRows(rows)
		// act
		res, err := f.s.Add(context.Background(), &models.Task{
			ID:          id,
			Title:       "test_t",
			Description: "test_d",
		})
		// assert
		require.NoError(t, err)
		assert.Equal(t, *res, id)
	})

	t.Run("fail", func(t *testing.T) {
		// arrange
		f := setUp(t)
		defer f.tearDown()

		query := regexp.QuoteMeta(`INSERT INTO tasks (id, title, description) VALUES ($1, $2, $3) RETURNING id`)

		id := uuid.New()
		f.dbMock.ExpectQuery(query).
			WithArgs(id, "test_t", "test_d").
			WillReturnError(errors.New(""))
		// act
		_, err := f.s.Add(context.Background(), &models.Task{
			ID:          id,
			Title:       "test_t",
			Description: "test_d",
		})
		// assert
		assert.Error(t, err)
	})
}

func TestSqlxDb_Delete(t *testing.T) {
	t.Run("normal", func(t *testing.T) {
		// arrange
		f := setUp(t)
		defer f.tearDown()

		query := regexp.QuoteMeta(`DELETE FROM tasks WHERE id = $1 RETURNING id`)

		id := uuid.New()
		rows := sqlmock.NewRows([]string{"id"}).AddRow(id)
		f.dbMock.ExpectQuery(query).
			WithArgs(id).
			WillReturnRows(rows)
		// act
		err := f.s.Delete(context.Background(), &id)

		// assert
		assert.NoError(t, err)
	})

	t.Run("no_task", func(t *testing.T) {
		// arrange
		f := setUp(t)
		defer f.tearDown()

		query := regexp.QuoteMeta(`DELETE FROM tasks WHERE id = $1 RETURNING id`)

		id := uuid.New()
		f.dbMock.ExpectQuery(query).
			WithArgs(id).
			WillReturnRows(sqlmock.NewRows([]string{"id"}))
		// act
		err := f.s.Delete(context.Background(), &id)

		// assert
		assert.Error(t, err)
	})
}

func TestSqlxDb_Update(t *testing.T) {
	t.Run("normal", func(t *testing.T) {
		f := setUp(t)
		defer f.tearDown()

		query := regexp.QuoteMeta(`UPDATE tasks SET title = $1, description = $2, edited = NOW() WHERE id = $3 RETURNING id`)

		id := uuid.New()
		f.dbMock.ExpectQuery(query).
			WithArgs("test_t", "test_d", id).
			WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(id))

		err := f.s.Update(context.Background(), &models.Task{ID: id, Title: "test_t", Description: "test_d"})

		assert.NoError(t, err)
	})

	t.Run("no_task", func(t *testing.T) {
		f := setUp(t)
		defer f.tearDown()

		query := regexp.QuoteMeta(`UPDATE tasks SET title = $1, description = $2, edited = NOW() WHERE id = $3 RETURNING id`)

		id := uuid.New()
		f.dbMock.ExpectQuery(query).
			WithArgs("test_t", "test_d", id).
			WillReturnRows(sqlmock.NewRows([]string{"id"}))

		err := f.s.Update(context.Background(), &models.Task{ID: id, Title: "test_t", Description: "test_d"})

		assert.Error(t, err)
	})

	t.Run("normal", func(t *testing.T) {
		f := setUp(t)
		defer f.tearDown()

		query := regexp.QuoteMeta(`UPDATE tasks SET title = $1, description = $2, edited = NOW() WHERE id = $3 RETURNING id`)

		id := uuid.New()
		f.dbMock.ExpectQuery(query).
			WithArgs("test_t", "test_d", id).
			WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(id))

		err := f.s.Update(context.Background(), &models.Task{ID: id, Title: "test_t", Description: "test_d"})

		assert.NoError(t, err)
	})
}
