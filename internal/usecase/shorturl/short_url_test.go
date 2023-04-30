package shorturl

import (
	"errors"
	"go-url-shortener/internal/mocks"
	"testing"

	"github.com/go-playground/assert/v2"
)

type Expected struct {
	Res interface{}
	Err error
}

func TestShortUrlCreate(t *testing.T) {
	repo, client := new(mocks.ShortUrlRepository), new(mocks.RedisClient)
	usecase := NewShortUrlUsecase(repo, client)

	tests := []struct {
		name     string
		input    string
		runMock  func()
		expected *Expected
	}{
		{
			"when success",
			"https://example.com/foobar",
			func() {
				repo.On("Create", "https://example.com/foobar").Return(uint64(10000), nil)
			},
			&Expected{
				Res: "SlC",
				Err: nil,
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.runMock()

			res, err := usecase.Create(test.input)
			assert.Equal(t, test.expected.Res, res)
			assert.Equal(t, test.expected.Err, err)
		})
	}
}

func TestShortUrlRedirect(t *testing.T) {
	repo, client := new(mocks.ShortUrlRepository), new(mocks.RedisClient)
	usecase := NewShortUrlUsecase(repo, client)

	tests := []struct {
		name     string
		input    string
		runMock  func()
		expected *Expected
	}{
		{
			"when code contains non-alphanumeric chars",
			"AB]C",
			func() {},
			&Expected{
				Res: "",
				Err: errors.New("Invalid character: ]"),
			},
		},
		{
			"when code is cached, return the cache entry",
			"SlC",
			func() {
				client.On("Get", "SlC").Return("https://example.com/foobar", nil)
			},
			&Expected{
				Res: "https://example.com/foobar",
				Err: nil,
			},
		},
		{
			"when code is cached and url is empty string",
			"B",
			func() {
				client.On("Get", "B").Return("", nil)
			},
			&Expected{
				Res: "",
				Err: errors.New("URL not found."),
			},
		},
		{
			"when code exists, find and cache",
			"ABC",
			func() {
				client.On("Get", "ABC").Return("", errors.New("No entry"))
				repo.On("Find", uint64(7750)).Return("https://example.com/foobar", nil)
				client.On("Set", "ABC", "https://example.com/foobar").Return(nil)
			},
			&Expected{
				Res: "https://example.com/foobar",
				Err: nil,
			},
		},
		{
			"when code isn't found, find and cache",
			"abc",
			func() {
				client.On("Get", "abc").Return("", errors.New("No entry"))
				repo.On("Find", uint64(109332)).Return("", errors.New("URL not found."))
				client.On("Set", "abc", "").Return(nil)
			},
			&Expected{
				Res: "",
				Err: errors.New("URL not found."),
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.runMock()

			res, err := usecase.Redirect(test.input)
			assert.Equal(t, test.expected.Res, res)
			assert.Equal(t, test.expected.Err, err)
		})
	}
}

func TestShortUrlDelete(t *testing.T) {
	repo, client := new(mocks.ShortUrlRepository), new(mocks.RedisClient)
	usecase := NewShortUrlUsecase(repo, client)

	tests := []struct {
		name     string
		input    string
		runMock  func()
		expected error
	}{
		{
			"when successs",
			"SlC",
			func() {
				repo.On("Delete", uint64(10000)).Return(nil)
				client.On("Del", "SlC").Return(nil)
			},
			nil,
		},
		{
			"when URL not found",
			"ABC",
			func() {
				repo.On("Delete", uint64(7750)).Return(errors.New("URL not found."))
			},
			errors.New("URL not found."),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.runMock()

			err := usecase.Delete(test.input)
			assert.Equal(t, test.expected, err)
		})
	}
}
