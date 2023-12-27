package handler_test

import (
	"bytes"
	"encoding/json"
	"go-url-shortener/internal/handler"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type Expected struct {
	Status int
	Body   string
}

func TestCreateURL(t *testing.T) {
	tests := []struct {
		Name     string
		Input    map[string]string
		// RunMock  func()
		Expected *Expected
	}{
		{
			"when success",
			map[string]string{
				"long_url": "https://example.com/foobar",
			},
			// func() { shortUrlMock.On("Create", "https://example.com/foobar").Return("abc", nil) },
			&Expected{
				Status: http.StatusOK,
				Body:   `{"data":"B"}`,
			},
		},
		{
			"when long_url is invalid",
			map[string]string{
				"long_url": "abc",
			},
			// func() {},
			&Expected{
				Status: http.StatusBadRequest,
				Body:   `{"error":"invalid params"}`,
			},
		},
	}
	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			w := httptest.NewRecorder()
			reqBody, _ := json.Marshal(test.Input)

			ctx := CreateGinContext("POST", "/v1/urls", reqBody, w)
			// test.RunMock()
			handler.CreateURL(ctx)
			
			assert.Equal(t, test.Expected.Status, w.Code)
			assert.Equal(t, test.Expected.Body, w.Body.String())
		})
	}
}

func TestRedirectURL(t *testing.T) {
	tests := []struct {
		Name     string
		Input    string
		// RunMock  func()
		Expected *Expected
	}{
		{
			"when success",
			"abc",
			// func() { shortUrlMock.On("Redirect", "abc").Return("https://example.com/foobar", nil) },
			&Expected{
				Status: http.StatusFound,
				Body:   "https://example.com/foobar",
			},
		},
		// {
		// 	"when url not found",
		// 	"test",
		// 	// func() { shortUrlMock.On("Redirect", "test").Return("", errors.New("Short URL not found.")) },
		// 	&Expected{
		// 		Status: http.StatusNotFound,
		// 		Body:   `{"error_message":"Short URL not found."}`,
		// 	},
		// },
		{
			"when the length of name exceeds 6",
			"Abcd1234",
			// func() {},
			&Expected{
				Status: http.StatusBadRequest,
				Body:   `{"error":"invalid params"}`,
			},
		},
		{
			"when name includes non-alphanumeric chars",
			"A!,_b",
			// func() {},
			&Expected{
				Status: http.StatusBadRequest,
				Body:   `{"error":"invalid params"}`,
			},
		},
	}
	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			w := httptest.NewRecorder()

			ctx := CreateGinContext("GET", "/v1/urls", nil, w)
			ctx.Params = append(ctx.Params, gin.Param{Key: "name", Value: test.Input})
			// test.RunMock()
			handler.RedirectURL(ctx)
			
			assert.Equal(t, test.Expected.Status, w.Code)
			if test.Expected.Status == http.StatusFound {
				assert.Equal(t, test.Expected.Body, w.Header().Get("Location"))
			} else {
				assert.Equal(t, test.Expected.Body, w.Body.String())
			}
		})
	}
}

// func TestShortURLHandlerDelete(t *testing.T) {
// 	gin.SetMode(gin.TestMode)

// 	tests := []struct {
// 		Name     string
// 		Input    string
// 		RunMock  func()
// 		Expected *Expected
// 	}{
// 		{
// 			"when success",
// 			"abc",
// 			func() { shortUrlMock.On("Delete", "abc").Return(nil) },
// 			&Expected{
// 				Status: http.StatusOK,
// 				Body:   `{"message":"URL deleted successfully.","data":null}`,
// 			},
// 		},
// 		{
// 			"when url not found",
// 			"123",
// 			func() { shortUrlMock.On("Delete", "123").Return(errors.New("URL not found.")) },
// 			&Expected{
// 				Status: http.StatusNotFound,
// 				Body:   `{"error_message":"URL not found."}`,
// 			},
// 		},
// 		{
// 			"when the length of code exceeds 7",
// 			"abcdefgh",
// 			func() {},
// 			&Expected{
// 				Status: http.StatusBadRequest,
// 				Body:   `{"error_message":"Invalid params."}`,
// 			},
// 		},
// 	}
// 	for _, test := range tests {
// 		t.Run(test.Name, func(t *testing.T) {
// 			w := httptest.NewRecorder()

// 			req, _ := http.NewRequest("DELETE", fmt.Sprintf("/api/v1/urls/%s", test.Input), nil)

// 			c, _ := gin.CreateTestContext(w)
// 			c.Request = req
// 			c.Params = append(c.Params, gin.Param{Key: "code", Value: test.Input})

// 			test.RunMock()
// 			handler.Delete(c)

// 			assert.Equal(t, test.Expected.Status, w.Code)
// 			assert.Equal(t, test.Expected.Body, w.Body.String())
// 		})
// 	}
// }

func CreateGinContext(action, path string, input []byte,  w *httptest.ResponseRecorder) *gin.Context {
	gin.SetMode(gin.TestMode)

	var body io.Reader
	if input != nil {
		body = bytes.NewBuffer(input)
	}
	req, _ := http.NewRequest(action, path, body)
	req.Header.Set("Content-Type", "application/json")

	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = req

	return ctx
}
