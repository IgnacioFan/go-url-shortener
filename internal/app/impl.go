package app

import (
	"go-url-shortener/internal/handler"

	"github.com/gin-gonic/gin"
)

type ShortUrl struct {
	*gin.Engine
}

func InitShortUrl() *ShortUrl {
	service := &ShortUrl{
		Engine:   gin.Default(),
	}

	service.SetRouter()
	return service
}

func (s *ShortUrl) SetRouter() {
	// s.GET("/health", func(ctx *gin.Context) {
	// 	ctx.JSON(http.StatusOK, &handler.Response{
	// 		Message: "healthy",
	// 	})
	// })
	v1 := s.Group("v1")
	{
		v1.POST("/urls", handler.CreateURL)
		v1.GET("/urls/:name", handler.RedirectURL)
	}
	// s.DELETE("/urls/:code", s.ShortUrl.Delete)
}
