package postgres

import (
	"errors"
	"regexp"
	"testing"

	"github.com/go-playground/assert/v2"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestUrlCreate(t *testing.T) {
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
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("Failed to setup sqlmock: %s", err)
			}
			orm, err := gorm.Open(postgres.New(postgres.Config{DriverName: "postgres", Conn: db}), &gorm.Config{})
			if err != nil {
				t.Fatalf("Failed to use gorm to open DB connection: %s", err)
			}
			test.runSQL(mock)
			urlRepository := Url{DB: orm}
			id, err := urlRepository.Create(test.input)

			if err != nil {
				assert.Equal(t, errors.New("record not found"), err)
			} else {
				assert.Equal(t, id, uint64(1))
			}
		})
	}
}

func TestUrlFindBy(t *testing.T) {
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
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("Failed to setup sqlmock: %s", err)
			}
			orm, err := gorm.Open(postgres.New(postgres.Config{DriverName: "postgres", Conn: db}), &gorm.Config{})
			if err != nil {
				t.Fatalf("Failed to use gorm to open DB connection: %s", err)
			}
			test.runSQL(mock)
			urlRepository := Url{DB: orm}
			id, err := urlRepository.FindBy(test.input)

			if err != nil {
				assert.Equal(t, errors.New("record not found"), err)
			} else {
				assert.Equal(t, id, uint64(1))
			}
		})
	}
}
