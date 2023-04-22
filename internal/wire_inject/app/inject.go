package app

import (
	"fmt"
	"go-url-shortener/config"
	"go-url-shortener/internal/delivery"
)

type Application struct {
	HttpServer *delivery.HttpServer
	Config     *config.Config
}

func NewApplication(server *delivery.HttpServer, config *config.Config) Application {
	return Application{
		HttpServer: server,
		Config:     config,
	}
}

func (app *Application) Start() error {
	port := app.Config.Http.Port
	return app.HttpServer.Run(fmt.Sprintf(":%d", port))
}
