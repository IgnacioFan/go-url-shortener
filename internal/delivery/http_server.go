package delivery

import (
	"go-url-shortener/internal/delivery/handler"
	"net/http"

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
	s.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, &handler.Response{
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
