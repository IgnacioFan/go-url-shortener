package integration

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func (s *TestSuite) TestPostUrls() {
	defer s.TearDown()
	reqBody := `{
		"url": "https://example.com/foobar"
	}`
	io := bytes.NewBuffer([]byte(reqBody))
	res, err := http.Post(
		fmt.Sprintf("%s/api/v1/urls", s.TestServer.URL),
		"application/json",
		io,
	)
	s.Require().NoError(err)
	body, err := ioutil.ReadAll(res.Body)
	s.Require().NoError(err)
	defer res.Body.Close()

	s.Require().Equal(http.StatusOK, res.StatusCode)
	s.Require().Equal(`{"message":"Short URL created successfully","data":{"short_url":"B","expiration":null}}`, string(body))
}

func (s *TestSuite) SetupShortenedURL() {
	sql := `insert into short_urls (url) values ('https://example.com/foobar');`
	s.TestDB.DB.Exec(sql)
	fmt.Println("set up shortened URL")
}

func (s *TestSuite) TearDown() {
	sql := `
		truncate table short_urls;
		alter sequence short_urls_id_seq restart with 1;
	`
	s.TestDB.DB.Exec(sql)
	fmt.Println("tear down tables")
}

func (s *TestSuite) TestRedirectCode() {
	s.SetupShortenedURL()
	defer s.TearDown()
	client := http.Client{
		// use http.Client with checkRedirect to make Go NOT follow redirects
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		}}
	res, err := client.Get(fmt.Sprintf("%s/B", s.TestServer.URL))
	s.Require().NoError(err)
	defer res.Body.Close()

	s.Require().Equal(http.StatusFound, res.StatusCode)
	s.Require().Equal("https://example.com/foobar", res.Header.Get("Location"))
}

func (s *TestSuite) TestDeleteCode() {
	s.SetupShortenedURL()
	defer s.TearDown()

	client := &http.Client{}
	req, err := http.NewRequest(
		"DELETE",
		fmt.Sprintf("%s/api/v1/urls/B", s.TestServer.URL),
		nil,
	)
	s.Require().NoError(err)
	res, err := client.Do(req)
	s.Require().NoError(err)
	body, err := ioutil.ReadAll(res.Body)
	s.Require().NoError(err)
	defer res.Body.Close()

	s.Require().Equal(http.StatusOK, res.StatusCode)
	s.Require().Equal(`{"message":"URL deleted successfully.","data":null}`, string(body))
}
