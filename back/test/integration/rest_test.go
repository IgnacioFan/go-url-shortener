package integration

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/suite"
)

const (
  INSERT_SQL        = `INSERT INTO urls (long_url, short_url) VALUES ('https://example.com/foobar', 'B');`
  RESTORE_TABLE_SQL = `truncate table urls;alter sequence urls_id_seq restart with 1;`
  REQUEST_HEADERS   = `application/json`
)

type RestTestSuite struct {
  TestSuite
}

func TestShortUrlTestSuite(t *testing.T) {
  suite.Run(t, new(RestTestSuite))
}

func (t *RestTestSuite) Seed() {
  t.TestDB.DB.Exec(INSERT_SQL)
}

func (t *RestTestSuite) RestoreTable() {
  t.TestDB.DB.Exec(RESTORE_TABLE_SQL)
}

func (t *RestTestSuite) TestHealth() {
  req, err := http.NewRequest("GET", "/v1/health", nil) 
  t.Require().NoError(err)
  w := httptest.NewRecorder()
  t.TestServer.ServeHTTP(w, req)
  
	t.Require().Equal(http.StatusOK, w.Code)
	t.Require().Equal(`{"message":"healthy"}`, w.Body.String())
}

func (t *RestTestSuite) TestCreate() {
  defer t.RestoreTable()

  testCases := []struct {
    Desc    string
    Input string
    Output string
  }{
    {
      "when success",
      `{"long_url": "https://example.com/foobar"}`,
      `{"data":"B"}`,
    },
    // {
    // 	"return the same short code when url is the same",
    // 	`{"url": "https://example.com/foobar"}`,
    // 	`{"data":{"short_url":"B","expiration":null}}`,
    // },
  }
  for _, test := range testCases {
    t.Run(test.Desc, func() {      
      payload := bytes.NewBuffer([]byte(test.Input))
      req, err := http.NewRequest("POST", "/v1/urls", payload) 
      t.Require().NoError(err)
      w := httptest.NewRecorder()
      t.TestServer.ServeHTTP(w, req)
    
      t.Require().Equal(http.StatusOK, w.Code)
      t.Require().NotEmpty(w.Body.String())
    })
  }
}

func (t *RestTestSuite) TestRedirect() {
	t.Seed()
	defer t.RestoreTable()

	testCases := []struct {
		Desc     string
		Input    string
		Output string
	}{
		{
			"redirect to the original URL",
			"B",
			"https://example.com/foobar",
		},
	}
	for _, test := range testCases {
		t.Run(test.Desc, func() {
      req, err := http.NewRequest("GET", fmt.Sprintf("/v1/urls/%s", test.Input), nil) 
      t.Require().NoError(err)
      w := httptest.NewRecorder()
      t.TestServer.ServeHTTP(w, req)

			t.Require().Equal(http.StatusFound, w.Code)
			t.Require().Equal(test.Output, w.Header().Get("Location"))
		})
	}
}

