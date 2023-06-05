package test

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	h "go-url-shortener/internal/delivery/handler"
	"go-url-shortener/test/mocks"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var (
	shortUrlMock = new(mocks.ShortUrlUsecase)
	handler      = h.NewShortUrlHandler(shortUrlMock)
)

type Expected struct {
	Status int
	Body   string
}

func TestShortURLHandlerCreate(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tests := []struct {
		Name     string
		Input    map[string]string
		RunMock  func()
		Expected *Expected
	}{
		{
			"when success",
			map[string]string{
				"url": "https://example.com/foobar",
			},
			func() { shortUrlMock.On("Create", "https://example.com/foobar").Return("abc", nil) },
			&Expected{
				Status: http.StatusOK,
				Body:   `{"message":"Short URL created successfully","data":{"short_url":"abc","expiration":null}}`,
			},
		},
		{
			"when url is invalid",
			map[string]string{
				"url": "abc",
			},
			func() {},
			&Expected{
				Status: http.StatusBadRequest,
				Body:   `{"error_message":"Invalid params."}`,
			},
		},
	}
	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			w := httptest.NewRecorder()

			requestBody, _ := json.Marshal(test.Input)

			req, _ := http.NewRequest("POST", "/api/v1/urls", bytes.NewBuffer(requestBody))
			req.Header.Set("Content-Type", "application/json")

			c, _ := gin.CreateTestContext(w)
			c.Request = req

			test.RunMock()
			handler.Create(c)

			assert.Equal(t, test.Expected.Status, w.Code)
			assert.Equal(t, test.Expected.Body, w.Body.String())
		})
	}
}

func TestShortURLHandlerRedirect(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tests := []struct {
		Name     string
		Input    string
		RunMock  func()
		Expected *Expected
	}{
		{
			"when success",
			"abc",
			func() { shortUrlMock.On("Redirect", "abc").Return("https://example.com/foobar", nil) },
			&Expected{
				Status: http.StatusFound,
				Body:   "https://example.com/foobar",
			},
		},
		{
			"when url not found",
			"test",
			func() { shortUrlMock.On("Redirect", "test").Return("", errors.New("Short URL not found.")) },
			&Expected{
				Status: http.StatusNotFound,
				Body:   `{"error_message":"Short URL not found."}`,
			},
		},
		{
			"when the length of code exceeds 7",
			"Abcd1234",
			func() {},
			&Expected{
				Status: http.StatusBadRequest,
				Body:   `{"error_message":"Invalid params."}`,
			},
		},
		{
			"when code contians non-alphanumeric chars",
			"A!,_b",
			func() {},
			&Expected{
				Status: http.StatusBadRequest,
				Body:   `{"error_message":"Invalid params."}`,
			},
		},
	}
	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			w := httptest.NewRecorder()

			req, _ := http.NewRequest("GET", fmt.Sprintf("/%s", test.Input), nil)

			c, _ := gin.CreateTestContext(w)
			c.Request = req

			c.Params = append(c.Params, gin.Param{Key: "code", Value: test.Input})

			test.RunMock()
			handler.Redirect(c)

			assert.Equal(t, test.Expected.Status, w.Code)
			if test.Expected.Status == http.StatusFound {
				assert.Equal(t, test.Expected.Body, w.Header().Get("Location"))
			} else {
				assert.Equal(t, test.Expected.Body, w.Body.String())
			}
		})
	}
}

func TestShortURLHandlerDelete(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tests := []struct {
		Name     string
		Input    string
		RunMock  func()
		Expected *Expected
	}{
		{
			"when success",
			"abc",
			func() { shortUrlMock.On("Delete", "abc").Return(nil) },
			&Expected{
				Status: http.StatusOK,
				Body:   `{"message":"URL deleted successfully.","data":null}`,
			},
		},
		{
			"when url not found",
			"123",
			func() { shortUrlMock.On("Delete", "123").Return(errors.New("URL not found.")) },
			&Expected{
				Status: http.StatusNotFound,
				Body:   `{"error_message":"URL not found."}`,
			},
		},
		{
			"when the length of code exceeds 7",
			"abcdefgh",
			func() {},
			&Expected{
				Status: http.StatusBadRequest,
				Body:   `{"error_message":"Invalid params."}`,
			},
		},
	}
	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			w := httptest.NewRecorder()

			req, _ := http.NewRequest("DELETE", fmt.Sprintf("/api/v1/urls/%s", test.Input), nil)

			c, _ := gin.CreateTestContext(w)
			c.Request = req
			c.Params = append(c.Params, gin.Param{Key: "code", Value: test.Input})

			test.RunMock()
			handler.Delete(c)

			assert.Equal(t, test.Expected.Status, w.Code)
			assert.Equal(t, test.Expected.Body, w.Body.String())
		})
	}
}
