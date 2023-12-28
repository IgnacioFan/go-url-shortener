package integration

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/suite"
)

const (
	INSERT_SQL        = `insert into short_urls (url) values ('https://example.com/foobar');`
	RESTORE_TABLE_SQL = `truncate table short_urls;alter sequence short_urls_id_seq restart with 1;`
	REQUEST_HEADERS   = `application/json`
)

type ShortUrlTestSuite struct {
	TestSuite
}

func TestShortUrlTestSuite(t *testing.T) {
	suite.Run(t, new(ShortUrlTestSuite))
}

func (t *ShortUrlTestSuite) Seed() {
	t.TestDB.DB.Exec(INSERT_SQL)
}

func (t *ShortUrlTestSuite) RestoreTable() {
	t.TestDB.DB.Exec(RESTORE_TABLE_SQL)
}

func (t *ShortUrlTestSuite) TestCreate() {
	url := t.TestServer.URL
	defer t.RestoreTable()

	testCases := []struct {
		Desc    string
		ReqBody string
		ResBody string
	}{
		{
			"return a short code",
			`{"url": "https://example.com/foobar"}`,
			`{"message":"Short URL created successfully","data":{"short_url":"B","expiration":null}}`,
		},
		{
			"return the same short code when url is the same",
			`{"url": "https://example.com/foobar"}`,
			`{"message":"Short URL created successfully","data":{"short_url":"B","expiration":null}}`,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.Desc, func() {
			io := bytes.NewBuffer([]byte(tc.ReqBody))
			res, err := http.Post(
				fmt.Sprintf("%s/api/v1/urls", url),
				REQUEST_HEADERS,
				io,
			)
			t.Require().NoError(err)
			body, _ := ioutil.ReadAll(res.Body)
			defer res.Body.Close()

			t.Require().Equal(http.StatusOK, res.StatusCode)
			t.Require().Equal(tc.ResBody, string(body))
		})
	}
}

// func (t *ShortUrlTestSuite) TestRedirect() {
// 	url := t.TestServer.URL
// 	t.Seed()
// 	defer t.RestoreTable()

// 	testCases := []struct {
// 		Desc     string
// 		Input    string
// 		Expected string
// 	}{
// 		{
// 			"redirect to the original URL",
// 			"B",
// 			"https://example.com/foobar",
// 		},
// 	}
// 	for _, tc := range testCases {
// 		t.Run(tc.Desc, func() {
// 			client := http.Client{
// 				// use http.Client with checkRedirect to make Go NOT follow redirects
// 				CheckRedirect: func(req *http.Request, via []*http.Request) error {
// 					return http.ErrUseLastResponse
// 				}}
// 			res, err := client.Get(fmt.Sprintf("%s/%s", url, tc.Input))
// 			t.Require().NoError(err)
// 			defer res.Body.Close()

// 			t.Require().Equal(http.StatusFound, res.StatusCode)
// 			t.Require().Equal(tc.Expected, res.Header.Get("Location"))
// 		})
// 	}
// }

// func (t *ShortUrlTestSuite) TestDelete() {
// 	url := t.TestServer.URL
// 	t.Seed()
// 	defer t.RestoreTable()

// 	testCases := []struct {
// 		Desc    string
// 		Input   string
// 		ResBody string
// 	}{
// 		{
// 			"success",
// 			"B",
// 			`{"message":"URL deleted successfully.","data":null}`,
// 		},
// 	}
// 	for _, tc := range testCases {
// 		t.Run(tc.Desc, func() {
// 			client := &http.Client{}
// 			req, err := http.NewRequest(
// 				"DELETE",
// 				fmt.Sprintf("%s/api/v1/urls/%s", url, tc.Input),
// 				nil,
// 			)
// 			t.Require().NoError(err)
// 			res, err := client.Do(req)
// 			t.Require().NoError(err)
// 			body, _ := ioutil.ReadAll(res.Body)
// 			defer res.Body.Close()

// 			t.Require().Equal(http.StatusOK, res.StatusCode)
// 			t.Require().Equal(tc.ResBody, string(body))
// 		})
// 	}
// }
