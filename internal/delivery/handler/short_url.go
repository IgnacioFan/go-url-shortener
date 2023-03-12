package handler

import (
	"go-url-shortener/internal/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ShortUrlHandler struct {
	ShortUrl *usecase.ShortUrl
}

type ShortUrlRequest struct {
	Url                 string `json:"url"`
	ExpirationlenInMins int    `json:"expiration_len_in_mins"`
}

func NewShortUrlHandler(usecase *usecase.ShortUrl) *ShortUrlHandler {
	return &ShortUrlHandler{
		ShortUrl: usecase,
	}
}

func (h *ShortUrlHandler) Create(ctx *gin.Context) {
	request := &ShortUrlRequest{}
	if err := ctx.BindJSON(request); err != nil {
		ctx.String(http.StatusNotFound, "Invalid params.")
		return
	}

	url, err := h.ShortUrl.Create(request.Url)
	if err != nil {
		ctx.JSON(http.StatusNotFound, nil)
		return
	}
	ctx.JSON(http.StatusOK, url)
}

func (h *ShortUrlHandler) Redirect(ctx *gin.Context) {
	url, ok := ctx.Params.Get("url")
	if !ok {
		ctx.JSON(http.StatusNotFound, nil)
		return
	}
	originalURL, err := h.ShortUrl.Redirect(url)
	if err != nil {
		ctx.JSON(http.StatusNotFound, nil)
		return
	}
	ctx.Redirect(http.StatusFound, originalURL)
}
