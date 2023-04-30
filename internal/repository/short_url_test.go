package repository

import (
	"database/sql"
	"regexp"
	"testing"

	pkg "go-url-shortener/pkg/postgres"

	"gorm.io/driver/postgres"

	"github.com/go-playground/assert/v2"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"gorm.io/gorm"
)

func MockDB(t *testing.T) (*sql.DB, *pkg.Postgres, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Unexpected error '%s' happened when opening a mock database", err)
	}
	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
	if err != nil {
		t.Errorf("Failed to open gorm db, got error: %v", err)
	}
	conn := &pkg.Postgres{DB: gormDB}
	return db, conn, mock
}

func TestShortUrlCreate(t *testing.T) {
	db, conn, mock := MockDB(t)
	defer db.Close()
	repository := NewShortUrlRepo(conn)

	tests := []struct {
		name     string
		input    string
		runSQL   func()
		expected struct {
			ID  uint64
			Err error
		}
	}{
		{
			"Return URL ID",
			"https://example.com/foobar",
			func() {
				mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "urls"`)).
					WillReturnError(gorm.ErrRecordNotFound)
				mock.ExpectBegin()
				mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "urls"`)).
					WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
				mock.ExpectCommit()
			},
			struct {
				ID  uint64
				Err error
			}{1, nil},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.runSQL()

			id, err := repository.Create(test.input)
			assert.Equal(t, test.expected.ID, id)
			assert.Equal(t, test.expected.Err, err)
		})
	}
}

func TestShortUrlFindBy(t *testing.T) {
	db, conn, mock := MockDB(t)
	defer db.Close()
	repository := NewShortUrlRepo(conn)

	tests := []struct {
		name     string
		input    string
		runSQL   func()
		expected struct {
			ID  uint64
			Err error
		}
	}{
		{
			"Return URL ID",
			"https://example.com/foobar",
			func() {
				mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "urls"`)).
					WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
			},
			struct {
				ID  uint64
				Err error
			}{1, nil},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.runSQL()

			id, err := repository.FindBy(test.input)
			assert.Equal(t, test.expected.ID, id)
			assert.Equal(t, test.expected.Err, err)
		})
	}
}

func TestShortUrlDelet(t *testing.T) {
	db, conn, mock := MockDB(t)
	defer db.Close()
	repository := NewShortUrlRepo(conn)

	tests := []struct {
		name     string
		input    uint64
		runSQL   func()
		expected error
	}{
		{
			"when delete successfully",
			1,
			func() {
				mock.ExpectBegin()
				mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM "urls"`)).WithArgs(1).WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectCommit()
			},
			nil,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.runSQL()

			err := repository.Delete(test.input)
			assert.Equal(t, test.expected, err)
		})
	}
}
