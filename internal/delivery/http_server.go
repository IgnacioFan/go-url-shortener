package delivery

import (
	"go-url-shortener/internal/delivery/handler"

	"github.com/gin-gonic/gin"
)

var shortUrlHandler = handler.NewShortUrlHandler()

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
