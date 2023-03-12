package delivery

import (
	"go-url-shortener/internal/delivery/handler"
	"go-url-shortener/internal/repository/postgres"
	"go-url-shortener/internal/usecase"

	"github.com/gin-gonic/gin"
)

var (
	urlRepo         = &postgres.Url{DB: postgres.InitConn()}
	shortUrlUsecase = usecase.NewShortUrl(urlRepo)
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
