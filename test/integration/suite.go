package integration

import (
	"go-url-shortener/internal/adpater/postgres"
	"go-url-shortener/internal/app/rest"
	"go-url-shortener/internal/service/url_service"

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
	TestServer *rest.Impl
	TestDB     *postgres.Postgres
}

func (t *TestSuite) SetupSuite() {
	db, err := postgres.NewPostgres()
	t.Require().NoError(err)
	t.TestDB = db

	if err = db.NewMigrate(); err != nil {
		t.Require().NoError(err)
	}

	service, err := url_service.NewUrlService()   
	t.Require().NoError(err)

	t.TestServer = &rest.Impl{
		Engine: gin.Default(),
		Url: service,
	}
	t.TestServer.SetRouter()
}
