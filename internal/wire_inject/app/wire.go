//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.
package app

import (
	"go-url-shortener/internal/delivery/http"
	"go-url-shortener/internal/repository"
	"go-url-shortener/internal/usecase/shorturl"
	"go-url-shortener/pkg/postgres"
	"go-url-shortener/pkg/redis"

	"github.com/google/wire"
)

func Initialize() (Application, error) {
	wire.Build(
		NewApplication,
		http.NewShortUrlHandler,
		shorturl.NewShortUrlUsecase,
		repository.NewShortUrlRepo,
		postgres.NewPostgres,
		redis.InitClient,
	)
	return Application{}, nil
}
