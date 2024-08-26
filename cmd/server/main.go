package main

import (
	"log"
	"log/slog"
	"net/http"
	"os"

	"maxnap/platform/internal/generated/schema"
	"maxnap/platform/internal/handler"
	"maxnap/platform/internal/pkg/logger"

	"github.com/go-chi/chi/v5"
)

func main() {
	slogLogger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	structLogger := logger.New(slogLogger)
	server := handler.New(structLogger)

	r := chi.NewRouter()
	handler := schema.NewStrictHandler(server, nil)

	h := schema.HandlerFromMux(handler, r)

	s := &http.Server{
		Handler: h,
		Addr:    "0.0.0.0:8080",
	}

	log.Fatal(s.ListenAndServe())
}
