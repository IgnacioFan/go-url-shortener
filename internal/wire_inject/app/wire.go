//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.
package app

import (
	"go-url-shortener/config"
	"go-url-shortener/internal/delivery"

	"github.com/google/wire"
)

func Initialize() (Application, error) {
	wire.Build(
		NewApplication,
		delivery.NewHttpServer,
		config.New,
	)
	return Application{}, nil
}
