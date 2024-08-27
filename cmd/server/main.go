package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"

	conf "maxnap/platform/internal/config"
	"maxnap/platform/internal/generated/schema"
	"maxnap/platform/internal/handler"
	"maxnap/platform/internal/pkg/logger"
	"maxnap/platform/internal/pkg/pg_client"

	"github.com/go-chi/chi/v5"
)

func main() {
	ctx := context.Background()

	slogLogger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	structLogger := logger.New(slogLogger)

	cfg, err := conf.NewConfigServer(ctx, "config.toml")
	if err != nil {
		panic(fmt.Errorf("cannot prepare config: %w", err))
	}

	db, err := pg_client.New(*conf.NewPostgresConfig(cfg.PostgresDatabase))
	if err != nil {
		panic(fmt.Errorf("database connection error: %w", err))
	}

	// TODO: There is must be Service, not DB connection
	server := handler.New(structLogger, db)

	r := chi.NewRouter()
	handler := schema.NewStrictHandler(server, nil)

	h := schema.HandlerFromMux(handler, r)

	s := &http.Server{
		Handler: h,
		Addr:    "0.0.0.0:8080",
	}

	log.Fatal(s.ListenAndServe())
}
