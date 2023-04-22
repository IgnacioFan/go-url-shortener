package shorturl

import (
	"errors"
	"go-url-shortener/internal/mocks"
	"testing"

	"github.com/go-playground/assert/v2"
)

var (
	encodedUrl  = "SlC"
	originalUrl = "https://example.com/foobar"
	urlRepo     = new(mocks.UrlRepository)
	client      = new(mocks.RedisClient)
	shortUrl    = NewShortUrlUsecase(urlRepo, client)
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
			originalUrl,
			"SlC",
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
			urlRepo.On("Create", test.input).Return(uint64(10000), nil)

			res, err := shortUrl.Create(test.input)
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
		mockFunc    func(client *mocks.RedisClient, repo *mocks.UrlRepository)
		expectedRes string
		expectedErr error
	}{
		{
			"Invalid short URL",
			"abcdefgh",
			func(client *mocks.RedisClient, repo *mocks.UrlRepository) {},
			"",
			errors.New("Short URL not found"),
		},
		{
			"With non-alphanumeric characters",
			"AB]C",
			func(client *mocks.RedisClient, repo *mocks.UrlRepository) {},
			"",
			errors.New("Invalid character: ]"),
		},
		{
			"When url is cached, redirect with valid short URL",
			"SlC",
			func(client *mocks.RedisClient, repo *mocks.UrlRepository) {
				client.On("Get", "SlC").Return(originalUrl, nil)
			},
			originalUrl,
			nil,
		},
		{
			"When entry doesn't exist, ReadThruCache",
			"ABC",
			func(client *mocks.RedisClient, repo *mocks.UrlRepository) {
				client.On("Get", "ABC").Return("", errors.New("No entry"))
				repo.On("Find", uint64(7750)).Return(originalUrl, nil)
				client.On("Set", "ABC", originalUrl).Return(nil)
			},
			originalUrl,
			nil,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.mockFunc(client, urlRepo)

			res, err := shortUrl.Redirect(test.input)
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
