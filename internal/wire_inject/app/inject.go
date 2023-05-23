package app

import (
	"fmt"
	"go-url-shortener/internal/delivery"
	"go-url-shortener/internal/delivery/handler"
)

type Application struct {
	HttpServer *delivery.HttpServer
}

func NewApplication(handler *handler.ShortUrlHandler) Application {
	server := delivery.NewHttpServer(handler)

	return Application{
		HttpServer: server,
	}
}

func (app *Application) Start(port int) error {
	return app.HttpServer.Run(fmt.Sprintf(":%d", port))
}
