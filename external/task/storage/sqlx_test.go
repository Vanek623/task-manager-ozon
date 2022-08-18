package storage

import (
	"context"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gitlab.ozon.dev/Vanek623/task-manager-system/external/task/models"
	"regexp"
	"testing"
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

		query := regexp.QuoteMeta(`DELETE FROM tasks WHERE id = $1 RETURNING 1`)

		rows := sqlmock.NewRows([]string{"id"}).AddRow("1")
		f.dbMock.ExpectQuery(query).
			WithArgs("1").
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

		query := regexp.QuoteMeta(`DELETE FROM tasks WHERE id = $1 RETURNING 1`)

		f.dbMock.ExpectQuery(query).
			WithArgs("1").
			WillReturnRows(nil)
		// act
		err := f.s.Delete(context.Background(), 1)

		// assert
		assert.Error(t, err)
	})

}

func TestSqlxDb_Update(t *testing.T) {
	f := setUp(t)
	defer f.tearDown()

}
