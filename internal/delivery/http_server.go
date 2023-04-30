package delivery

import (
	"go-url-shortener/internal/delivery/handler"

	"github.com/gin-gonic/gin"
)

type HttpServer struct {
	*gin.Engine
	ShortUrl *handler.ShortUrlHandler
}

func NewHttpServer(shortUrl *handler.ShortUrlHandler) *HttpServer {
	server := &HttpServer{
		Engine:   gin.Default(),
		ShortUrl: shortUrl,
	}

	server.SetRouter()
	return server
}

func (s *HttpServer) SetRouter() {
	v1 := s.Group("api/v1")
	{
		v1.POST("/urls", s.ShortUrl.Create)
		v1.DELETE("/urls/:code", s.ShortUrl.Delete)
	}
	s.GET("/:code", s.ShortUrl.Redirect)
}
