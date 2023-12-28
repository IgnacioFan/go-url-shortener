package rest

import (
	"go-url-shortener/internal/service/url"
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

const (
	SHROT_URL_REGEX = "^[a-zA-Z0-9]{1,7}$"
)

var (
	invalidParams = &ErrorResponse{Error: "invalid params"}
)

type Response struct {
	Message string `json:"message,omitempty"`
	Data    interface{} `json:"data"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

type URLRequest struct {
	LongUrl                 string `json:"long_url"`
}

type Impl struct {
	*gin.Engine
	Url url.UrlService
}

func InitShortUrl(url url.UrlService) *Impl {
	service := &Impl{
		Engine:   gin.Default(),
		Url: url, 
	}

	service.SetRouter()
	return service
}

func (i *Impl) SetRouter() {
	// s.GET("/health", func(ctx *gin.Context) {
	// 	ctx.JSON(http.StatusOK, &handler.Response{
	// 		Message: "healthy",
	// 	})
	// })
	v1 := i.Group("v1")
	{
		v1.POST("/urls", i.CreateURL)
		v1.GET("/urls/:name", i.RedirectURL)
	}
}

func (i *Impl) CreateURL(ctx *gin.Context) {
	req := &URLRequest{}
	if err := ctx.BindJSON(req); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	if !govalidator.IsURL(req.LongUrl) {
		ctx.JSON(http.StatusBadRequest, invalidParams)
		return
	}
	name, err := i.Url.GenerateShortURL(req.LongUrl)
	if err != nil {
		ctx.JSON(http.StatusNotFound, &ErrorResponse{Error: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, &Response{
		Data:    name,
	})
}

func (i *Impl) RedirectURL(ctx *gin.Context) {
	shortUrl, ok := ctx.Params.Get("name")
	if !ok || !govalidator.Matches(shortUrl, SHROT_URL_REGEX) {
		ctx.JSON(http.StatusBadRequest, invalidParams)
		return
	}
	originalUrl, err := i.Url.OriginalURL(shortUrl)
	if err != nil {
		ctx.JSON(http.StatusNotFound, &ErrorResponse{Error: err.Error()})
		return
	}
	ctx.Redirect(http.StatusFound, originalUrl)
}
