package handler

import (
	"go-url-shortener/internal/usecase"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type ShortUrlHandler struct {
	ShortUrl usecase.ShortUrl
}

type ShortUrlResponse struct {
	UrlID      string     `json:"short_url"`
	Expiration *time.Time `json:"expiration"`
}

func NewShortUrlHandler(shortUrl usecase.ShortUrl) *ShortUrlHandler {
	return &ShortUrlHandler{ShortUrl: shortUrl}
}

func (h *ShortUrlHandler) Create(ctx *gin.Context) {
	request := &ShortUrlRequest{}
	if err := ctx.BindJSON(request); err != nil {
		ctx.JSON(http.StatusNotFound, invalidParams)
		return
	}

	url, err := h.ShortUrl.Create(request.Url)
	if err != nil {
		ctx.JSON(http.StatusNotFound, &ErrorResponse{ErrorMessage: err.Error()})
		return
	}

	response := &Response{
		Message: "Short URL created successfully",
		Data:    ShortUrlResponse{UrlID: url},
	}
	ctx.JSON(http.StatusOK, response)
}

func (h *ShortUrlHandler) Redirect(ctx *gin.Context) {
	url, ok := ctx.Params.Get("url")
	if !ok {
		ctx.JSON(http.StatusNotFound, invalidParams)
		return
	}
	originalURL, err := h.ShortUrl.Redirect(url)
	if err != nil {
		ctx.JSON(http.StatusNotFound, &ErrorResponse{ErrorMessage: err.Error()})
		return
	}
	ctx.Redirect(http.StatusFound, originalURL)
}
