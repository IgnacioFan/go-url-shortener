package app

import (
	"fmt"
	"go-url-shortener/config"
	"go-url-shortener/internal/delivery"
	"go-url-shortener/internal/delivery/handler"
)

type Application struct {
	HttpServer *delivery.HttpServer
	Config     *config.Config
}

func NewApplication(handler *handler.ShortUrlHandler, config *config.Config) Application {
	server := delivery.NewHttpServer(handler)

	return Application{
		HttpServer: server,
		Config:     config,
	}
}

func (app *Application) Start() error {
	port := app.Config.Http.Port
	return app.HttpServer.Run(fmt.Sprintf(":%d", port))
}
