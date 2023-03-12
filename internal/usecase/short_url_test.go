package usecase

import (
	"errors"
	"go-url-shortener/internal/mocks"
	"testing"

	"github.com/go-playground/assert/v2"
)

var (
	urlMock = new(mocks.UrlRepository)
	usecase = NewShortUrl(urlMock)
)

func TestShortUrlCreate(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		expectedRes string
		expectedErr error
	}{
		{
			"Create with URL",
			"https://example.com/foobar",
			"abc",
			nil,
		},
		{
			"Without URL",
			"",
			"",
			errors.New("Url is empty"),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			urlMock.On("Create", test.input).Return(uint64(1), nil)

			res, err := usecase.Create(test.input)
			if test.expectedErr != nil {
				assert.Equal(t, test.expectedRes, res)
				assert.Equal(t, test.expectedErr, err)
			} else {
				assert.Equal(t, test.expectedRes, res)
				assert.Equal(t, test.expectedErr, err)
			}
		})
	}
}

func TestShortUrlRedirect(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		expectedRes string
		expectedErr error
	}{
		{
			"Redirect with valid short URL",
			"valid",
			"https://example.com/foobar",
			nil,
		},
		{
			"Invalid short URL",
			"invalid",
			"",
			errors.New("Short URL not found"),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res, err := usecase.Redirect(test.input)
			if test.expectedErr != nil {
				assert.Equal(t, test.expectedRes, res)
				assert.Equal(t, test.expectedErr, err)
			} else {
				assert.Equal(t, test.expectedRes, res)
				assert.Equal(t, test.expectedErr, err)
			}
		})
	}
}
