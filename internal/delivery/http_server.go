package delivery

import (
	"go-url-shortener/internal/delivery/handler"
	"go-url-shortener/internal/repository/urlrepo"
	"go-url-shortener/internal/usecase/shorturl"
	"go-url-shortener/pkg/postgres"
	"go-url-shortener/pkg/redis"

	"github.com/gin-gonic/gin"
)

var (
	urlRepo         = urlrepo.NewUrlRepository(postgres.InitConn())
	urlClient       = redis.NewUrlClient()
	shortUrlUsecase = shorturl.NewShortUrl(urlRepo, urlClient)
	shortUrlHandler = handler.NewShortUrlHandler(shortUrlUsecase)
)

type HttpServer struct {
	*gin.Engine
}

func NewHttpServer() *HttpServer {
	server := &HttpServer{
		Engine: gin.Default(),
	}

	server.SetRouter()
	return server
}

func (s *HttpServer) SetRouter() {
	v1 := s.Group("api/v1")
	{
		v1.POST("/urls", shortUrlHandler.Create)
	}
	s.GET("/:url", shortUrlHandler.Redirect)
}
