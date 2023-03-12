package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"go-url-shortener/internal/usecase"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
)

var (
	response     string
	shortUrlMock = usecase.NewShortUrl()
	handler      = NewShortUrlHandler(shortUrlMock)
)

func TestCreate(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tests := []struct {
		Name        string
		Input       map[string]string
		ExpectedRes string
		ExpectedErr error
	}{
		{
			"Create with valid params",
			map[string]string{
				"url": "https://example.com/foobar",
			},
			"abc",
			nil,
		},
		{
			"Create with invalid params",
			map[string]string{},
			"",
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

			handler.Create(c)

			if test.ExpectedErr != nil {
				assert.Equal(t, http.StatusNotFound, w.Code)
			} else {
				assert.Equal(t, http.StatusOK, w.Code)
				response = ""
				json.Unmarshal(w.Body.Bytes(), &response)
				assert.Equal(t, test.ExpectedRes, response)
			}
		})
	}
}

func TestRedirect(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tests := []struct {
		Name        string
		Input       string
		ExpectedRes string
		ExpectedErr error
	}{
		{
			"Redirect with valid params",
			"valid",
			"https://example.com/foobar",
			nil,
		},
		{
			"Redirect with invalid params",
			"invalid",
			"",
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

			handler.Redirect(c)

			if test.ExpectedErr != nil {
				assert.Equal(t, http.StatusNotFound, w.Code)
			} else {
				assert.Equal(t, http.StatusFound, w.Code)
				assert.Equal(t, test.ExpectedRes, w.Header().Get("Location"))
			}
		})
	}
}
