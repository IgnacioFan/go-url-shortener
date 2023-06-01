package integration

import (
	"fmt"
	"go-url-shortener/internal/wire_inject/app"
	"go-url-shortener/pkg/postgres"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
}

func Test_get_health(t *testing.T) {
	db, err := postgres.NewPostgres()
	if err != nil {
		t.Fatalf("Test DB connection error: %s", err)
	}
	if err = db.NewMigrate(); err != nil {
		t.Fatalf("Test DB migration error: %s", err)
	}
	app, err := app.Initialize()
	if err != nil {
		t.Fatalf("App initialization error: %s", err)
	}
	ts := httptest.NewServer(
		app.HttpServer.Engine,
	)

	t.Run("returns ok when is heathy", func(t *testing.T) {
		resp, err := http.Get(fmt.Sprintf("%s/health", ts.URL))

		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		assert.Equal(t, 200, resp.StatusCode)
	})
}
