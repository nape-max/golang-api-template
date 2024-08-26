package main

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"

	conf "maxnap/platform/internal/config"
	"maxnap/platform/internal/generated/schema"
	"maxnap/platform/internal/handler"
	"maxnap/platform/internal/pkg/logger"

	"github.com/BurntSushi/toml"
	"github.com/go-chi/chi/v5"
)

func main() {
	slogLogger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	structLogger := logger.New(slogLogger)

	mydir, err := os.Getwd()
	if err != nil {
		panic(fmt.Errorf("cannot receive current dir: %w", err))
	}

	_, err = os.Stat(mydir + "/config.toml")
	if err != nil {
		panic(fmt.Errorf("cannot receive stat of config file: %w", err))
	}

	var cfg conf.ConfigServer
	_, err = toml.DecodeFile(mydir+"/config.toml", &cfg)
	if err != nil {
		panic(fmt.Errorf("cannot decode config to struct: %w", err))
	}

	fmt.Println(cfg)

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
