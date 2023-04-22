package repository

import (
	"errors"
	"regexp"
	"testing"

	_pkg "go-url-shortener/pkg/postgres"

	"github.com/go-playground/assert/v2"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func MockDB() (*_pkg.Postgres, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		panic(err)
	}
	orm, err := gorm.Open(postgres.New(postgres.Config{DriverName: "postgres", Conn: db}), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	conn := &_pkg.Postgres{
		DB: orm,
	}
	return conn, mock
}

func TestShortUrlCreate(t *testing.T) {
	db, mock := MockDB()

	tests := []struct {
		name   string
		input  string
		runSQL func(mock sqlmock.Sqlmock)
	}{
		{
			name:  "Return URL ID",
			input: "https://example.com/foobar",
			runSQL: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "urls"`)).
					WillReturnError(gorm.ErrRecordNotFound)
				mock.ExpectBegin()
				mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "urls"`)).
					WillReturnRows(sqlmock.NewRows([]string{"id"}).
						AddRow(1))
				mock.ExpectCommit()
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.runSQL(mock)
			repository := NewShortUrlRepo(db)
			id, err := repository.Create(test.input)

			if err != nil {
				assert.Equal(t, errors.New("record not found"), err)
			} else {
				assert.Equal(t, id, uint64(1))
			}
		})
	}
}

func TestShortUrlFindBy(t *testing.T) {
	db, mock := MockDB()

	tests := []struct {
		name   string
		input  string
		runSQL func(mock sqlmock.Sqlmock)
	}{
		{
			name:  "Return URL ID",
			input: "https://example.com/foobar",
			runSQL: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "urls"`)).
					WillReturnRows(sqlmock.NewRows([]string{"id"}).
						AddRow(1))
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.runSQL(mock)
			repository := NewShortUrlRepo(db)
			id, err := repository.FindBy(test.input)

			if err != nil {
				assert.Equal(t, errors.New("record not found"), err)
			} else {
				assert.Equal(t, id, uint64(1))
			}
		})
	}
}
