package storage

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"gitlab.ozon.dev/Vanek623/task-manager-system/external/counters"
)

type taskStorageFixture struct {
	s      iTaskStorage
	db     *sqlx.DB
	dbMock sqlmock.Sqlmock
}

func setUp(t *testing.T) taskStorageFixture {
	var fixture taskStorageFixture

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}

	fixture.db = sqlx.NewDb(db, "sqlmock")

	fixture.dbMock = mock
	fixture.s = NewSqlx(fixture.db, counters.New("test"))

	return fixture
}

func (f *taskStorageFixture) tearDown() {
	_ = f.db.Close()
}
