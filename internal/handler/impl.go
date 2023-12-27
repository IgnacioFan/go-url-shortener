package handler

import (
	"go-url-shortener/internal/service/url"
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
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

const (
	SHROT_URL_REGEX = "^[a-zA-Z0-9]{1,7}$"
)

var (
	invalidParams = &ErrorResponse{Error: "invalid params"}
)

func CreateURL(ctx *gin.Context) {
	req := &URLRequest{}
	if err := ctx.BindJSON(req); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	if !govalidator.IsURL(req.LongUrl) {
		ctx.JSON(http.StatusBadRequest, invalidParams)
		return
	}
	name, err := url.GenerateShortURL(req.LongUrl)
	if err != nil {
		ctx.JSON(http.StatusNotFound, &ErrorResponse{Error: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, &Response{
		Data:    name,
	})
}

func RedirectURL(ctx *gin.Context) {
	shortUrl, ok := ctx.Params.Get("name")
	if !ok || !govalidator.Matches(shortUrl, SHROT_URL_REGEX) {
		ctx.JSON(http.StatusBadRequest, invalidParams)
		return
	}
	originalUrl, err := url.OriginalURL(shortUrl)
	if err != nil {
		ctx.JSON(http.StatusNotFound, &ErrorResponse{Error: err.Error()})
		return
	}
	ctx.Redirect(http.StatusFound, originalUrl)
}

func DeleteURL(ctx *gin.Context)  {
	shortUrl, ok := ctx.Params.Get("name")
	if !ok || !govalidator.Matches(shortUrl, SHROT_URL_REGEX) {
		ctx.JSON(http.StatusBadRequest, invalidParams)
		return
	}
	// if err := h.ShortUrl.Delete(code); err != nil {
	// 	ctx.JSON(http.StatusNotFound, &ErrorResponse{Error: err.Error()})
	// 	return
	// }
	ctx.JSON(http.StatusOK, &Response{
		Message: "deleted the URL",
	})
}
