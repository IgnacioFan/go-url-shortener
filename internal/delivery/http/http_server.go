package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HttpServer struct {
	*gin.Engine
	ShortUrl *ShortUrlHandler
}

func NewHttpServer(shortUrl *ShortUrlHandler) *HttpServer {
	server := &HttpServer{
		Engine:   gin.Default(),
		ShortUrl: shortUrl,
	}

	server.SetRouter()
	return server
}

func (s *HttpServer) SetRouter() {
	s.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, &Response{
			Message: "healthy",
		})
	})

	v1 := s.Group("api/v1")
	{
		v1.POST("/urls", s.ShortUrl.Create)
		v1.DELETE("/urls/:code", s.ShortUrl.Delete)
	}
	s.GET("/:code", s.ShortUrl.Redirect)
}
