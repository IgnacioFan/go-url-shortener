package http

import (
	"go-url-shortener/internal/entity"
	"net/http"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

var (
	codeRegex = "^[a-zA-Z0-9]{1,7}$"
)

type ShortUrlHandler struct {
	ShortUrl entity.ShortUrlUsecase
}

type ShortUrlRequest struct {
	Url                 string `json:"url"`
	ExpirationlenInMins int    `json:"expiration_len_in_mins"`
}

type ShortUrlResponse struct {
	UrlID      string     `json:"short_url"`
	Expiration *time.Time `json:"expiration"`
}

func NewShortUrlHandler(shortUrl entity.ShortUrlUsecase) *ShortUrlHandler {
	return &ShortUrlHandler{ShortUrl: shortUrl}
}

func (h *ShortUrlHandler) Create(ctx *gin.Context) {
	req := &ShortUrlRequest{}
	if err := ctx.BindJSON(req); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	if !govalidator.IsURL(req.Url) {
		ctx.JSON(http.StatusBadRequest, invalidParams)
		return
	}
	url, err := h.ShortUrl.Create(req.Url)
	if err != nil {
		ctx.JSON(http.StatusNotFound, &ErrorResponse{ErrorMessage: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, &Response{
		Message: "Short URL created successfully",
		Data:    ShortUrlResponse{UrlID: url},
	})
}

func (h *ShortUrlHandler) Redirect(ctx *gin.Context) {
	url, ok := ctx.Params.Get("code")
	if !ok || !govalidator.Matches(url, codeRegex) {
		ctx.JSON(http.StatusBadRequest, invalidParams)
		return
	}
	originalURL, err := h.ShortUrl.Redirect(url)
	if err != nil {
		ctx.JSON(http.StatusNotFound, &ErrorResponse{ErrorMessage: err.Error()})
		return
	}
	ctx.Redirect(http.StatusFound, originalURL)
}

func (h *ShortUrlHandler) Delete(ctx *gin.Context) {
	code, ok := ctx.Params.Get("code")
	if !ok || !govalidator.Matches(code, codeRegex) {
		ctx.JSON(http.StatusBadRequest, invalidParams)
		return
	}
	if err := h.ShortUrl.Delete(code); err != nil {
		ctx.JSON(http.StatusNotFound, &ErrorResponse{ErrorMessage: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, &Response{
		Message: "URL deleted successfully.",
	})
}
