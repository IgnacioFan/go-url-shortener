package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"go-url-shortener/internal/mocks"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
)

var (
	shortUrlMock = new(mocks.ShortUrlUsecase)
	handler      = NewShortUrlHandler(shortUrlMock)
)

type Expected struct {
	Status int
	Body   string
}

func TestShortURLHandlerCreate(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tests := []struct {
		Name        string
		Input       map[string]string
		MockRes     string
		ExpectedRes string
		ExpectedErr error
	}{
		{
			"Create with valid params",
			map[string]string{
				"url": "https://example.com/foobar",
			},
			"abc",
			`{"message":"Short URL created successfully","data":{"short_url":"abc","expiration":null}}`,
			nil,
		},
		{
			"Create with invalid params",
			map[string]string{},
			"",
			`{"error_message":"Url is empty"}`,
			errors.New("Url is empty"),
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

			shortUrlMock.On("Create", test.Input["url"]).Return(test.MockRes, test.ExpectedErr)
			handler.Create(c)

			if test.ExpectedErr != nil {
				assert.Equal(t, http.StatusNotFound, w.Code)
				assert.Equal(t, test.ExpectedRes, w.Body.String())
			} else {
				assert.Equal(t, http.StatusOK, w.Code)
				assert.Equal(t, test.ExpectedRes, w.Body.String())
			}
		})
	}
}

func TestShortURLHandlerRedirect(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tests := []struct {
		Name        string
		Input       string
		MockRes     string
		ExpectedRes string
		ExpectedErr error
	}{
		{
			"Redirect with valid params",
			"valid",
			"https://example.com/foobar",
			"https://example.com/foobar",
			nil,
		},
		{
			"Redirect with invalid params",
			"invalid",
			"",
			`{"error_message":"Short URL not found"}`,
			errors.New("Short URL not found"),
		},
	}
	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			w := httptest.NewRecorder()

			req, _ := http.NewRequest("GET", fmt.Sprintf("/%s", test.Input), nil)

			c, _ := gin.CreateTestContext(w)
			c.Request = req

			c.Params = append(c.Params, gin.Param{Key: "url", Value: test.Input})

			shortUrlMock.On("Redirect", test.Input).Return(test.ExpectedRes, test.ExpectedErr)
			handler.Redirect(c)

			if test.ExpectedErr != nil {
				assert.Equal(t, http.StatusNotFound, w.Code)
				assert.Equal(t, test.ExpectedRes, w.Body.String())
			} else {
				assert.Equal(t, http.StatusFound, w.Code)
				assert.Equal(t, test.ExpectedRes, w.Header().Get("Location"))
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
			"when code exceeds max length",
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
