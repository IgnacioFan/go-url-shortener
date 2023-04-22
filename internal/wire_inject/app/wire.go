//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.
package app

import (
	"go-url-shortener/config"
	"go-url-shortener/internal/delivery/handler"
	"go-url-shortener/internal/repository"
	"go-url-shortener/internal/usecase/shorturl"
	"go-url-shortener/pkg/postgres"
	"go-url-shortener/pkg/redis"

	"github.com/google/wire"
)

func Initialize() (Application, error) {
	wire.Build(
		NewApplication,
		handler.NewShortUrlHandler,
		shorturl.NewShortUrl,
		repository.NewShortUrlRepo,
		postgres.InitPostgres,
		redis.InitClient,
		config.New,
	)
	return Application{}, nil
}
