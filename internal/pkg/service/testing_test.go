package service

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"testing"
)

type taskStorageFixture struct {
	service iStorage
	db      *sqlx.DB
	dbMock  sqlmock.Sqlmock
}

func setUp(t *testing.T) taskStorageFixture {
	t.Parallel()

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}

	f := taskStorageFixture{}
	f.db = sqlx.NewDb(db, "sqlmock")
	f.dbMock = mock
	f.service = New()

	return f
}
