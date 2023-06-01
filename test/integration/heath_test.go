package integration

import (
	"fmt"
	"go-url-shortener/internal/wire_inject/app"
	"go-url-shortener/pkg/postgres"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/suite"
)

func init() {
	godotenv.Load()
}

type ServiceTestSuite struct {
	suite.Suite
	TestServer *httptest.Server
}

func TestServiceTestSuite(t *testing.T) {
	suite.Run(t, new(ServiceTestSuite))
}

func (suite *ServiceTestSuite) SetupSuite() {
	db, err := postgres.NewPostgres()
	suite.Require().NoError(err)

	if err = db.NewMigrate(); err != nil {
		suite.Require().NoError(err)
	}

	app, err := app.Initialize()
	suite.Require().NoError(err)

	suite.TestServer = httptest.NewServer(
		app.HttpServer.Engine,
	)
}

func (suite *ServiceTestSuite) SetupTest() {
	// TBD
}

func (suite *ServiceTestSuite) TestGetHeath() {
	resp, err := http.Get(fmt.Sprintf("%s/health", suite.TestServer.URL))

	suite.Require().NoError(err)
	suite.Require().Equal(http.StatusOK, resp.StatusCode)
}
