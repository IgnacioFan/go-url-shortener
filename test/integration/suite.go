package integration

import (
	"go-url-shortener/internal/wire_inject/app"
	"go-url-shortener/pkg/postgres"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/suite"
)

func init() {
	gin.SetMode(gin.TestMode)
	godotenv.Load()
}

type TestSuite struct {
	suite.Suite
	TestServer *httptest.Server
	TestDB     *postgres.Postgres
}

func (t *TestSuite) SetupSuite() {
	db, err := postgres.NewPostgres()
	t.Require().NoError(err)

	if err = db.NewMigrate(); err != nil {
		t.Require().NoError(err)
	}

	app, err := app.Initialize()
	t.Require().NoError(err)

	t.TestDB = db
	t.TestServer = httptest.NewServer(app.HttpServer.Engine)
}
