package storage

import (
	"context"
	"regexp"
	"testing"

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

		query := regexp.QuoteMeta(`INSERT INTO tasks (title, description) VALUES ($1, $2) RETURNING id`)

		rows := sqlmock.NewRows([]string{"id"}).AddRow("1")
		f.dbMock.ExpectQuery(query).
			WithArgs("test_t", "test_d").
			WillReturnRows(rows)
		// act
		res, err := f.s.Add(context.Background(), &models.Task{
			Title:       "test_t",
			Description: "test_d",
		})
		// assert
		require.NoError(t, err)
		assert.Equal(t, res, uint64(1))
	})

	t.Run("fail", func(t *testing.T) {
		// arrange
		f := setUp(t)
		defer f.tearDown()

		query := regexp.QuoteMeta(`INSERT INTO tasks (title, description) VALUES ($1, $2) RETURNING id`)

		f.dbMock.ExpectQuery(query).
			WithArgs("test_t", "test_d").
			WillReturnError(errors.New(""))
		// act
		_, err := f.s.Add(context.Background(), &models.Task{
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

		rows := sqlmock.NewRows([]string{"id"}).AddRow(1)
		f.dbMock.ExpectQuery(query).
			WithArgs(1).
			WillReturnRows(rows)
		// act
		err := f.s.Delete(context.Background(), 1)

		// assert
		assert.NoError(t, err)
	})

	t.Run("no_task", func(t *testing.T) {
		// arrange
		f := setUp(t)
		defer f.tearDown()

		query := regexp.QuoteMeta(`DELETE FROM tasks WHERE id = $1 RETURNING id`)

		f.dbMock.ExpectQuery(query).
			WithArgs(1).
			WillReturnRows(sqlmock.NewRows([]string{"id"}))
		// act
		err := f.s.Delete(context.Background(), 1)

		// assert
		assert.Error(t, err)
	})
}

func TestSqlxDb_Update(t *testing.T) {
	t.Run("normal", func(t *testing.T) {
		f := setUp(t)
		defer f.tearDown()

		query := regexp.QuoteMeta(`UPDATE tasks SET title = $1, description = $2, edited = NOW() WHERE id = $3 RETURNING id`)

		f.dbMock.ExpectQuery(query).
			WithArgs("test_t", "test_d", 1).
			WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

		err := f.s.Update(context.Background(), &models.Task{ID: 1, Title: "test_t", Description: "test_d"})

		assert.NoError(t, err)
	})

	t.Run("no_task", func(t *testing.T) {
		f := setUp(t)
		defer f.tearDown()

		query := regexp.QuoteMeta(`UPDATE tasks SET title = $1, description = $2, edited = NOW() WHERE id = $3 RETURNING id`)

		f.dbMock.ExpectQuery(query).
			WithArgs("test_t", "test_d", 1).
			WillReturnRows(sqlmock.NewRows([]string{"id"}))

		err := f.s.Update(context.Background(), &models.Task{ID: 1, Title: "test_t", Description: "test_d"})

		assert.Error(t, err)
	})

	t.Run("normal", func(t *testing.T) {
		f := setUp(t)
		defer f.tearDown()

		query := regexp.QuoteMeta(`UPDATE tasks SET title = $1, description = $2, edited = NOW() WHERE id = $3 RETURNING id`)

		f.dbMock.ExpectQuery(query).
			WithArgs("test_t", "test_d", 1).
			WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

		err := f.s.Update(context.Background(), &models.Task{ID: 1, Title: "test_t", Description: "test_d"})

		assert.NoError(t, err)
	})
}
