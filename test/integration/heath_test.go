package integration

import (
	"fmt"
	"go-url-shortener/internal/wire_inject/app"
	"go-url-shortener/pkg/postgres"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/suite"
)

func init() {
	godotenv.Load()
}

type TestSuite struct {
	suite.Suite
	TestServer *httptest.Server
	TestDB     *postgres.Postgres
}

func TestTestSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}

func (s *TestSuite) SetupSuite() {
	db, err := postgres.NewPostgres()
	s.Require().NoError(err)

	if err = db.NewMigrate(); err != nil {
		s.Require().NoError(err)
	}

	app, err := app.Initialize()
	s.Require().NoError(err)

	s.TestDB = db
	s.TestServer = httptest.NewServer(app.HttpServer.Engine)
}

func (s *TestSuite) SetupTest() {
	// TBD
}

func (s *TestSuite) TestGetHeath() {
	res, err := http.Get(fmt.Sprintf("%s/health", s.TestServer.URL))
	s.Require().NoError(err)
	body, err := ioutil.ReadAll(res.Body)
	s.Require().NoError(err)
	defer res.Body.Close()

	s.Require().Equal(http.StatusOK, res.StatusCode)
	s.Require().Equal(`{"message":"heathy","data":null}`, string(body))
}
