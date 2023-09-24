package app

import (
	"fmt"
	"go-url-shortener/internal/delivery/http"
)

type Application struct {
	HttpServer *http.HttpServer
}

func NewApplication(handler *http.ShortUrlHandler) Application {
	server := http.NewHttpServer(handler)

	return Application{
		HttpServer: server,
	}
}

func (app *Application) Start(port int) error {
	return app.HttpServer.Run(fmt.Sprintf(":%d", port))
}
