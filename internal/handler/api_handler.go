package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"log/slog"

	"maxnap/platform/internal/pkg/logger"

	"github.com/jmoiron/sqlx"
)

type Server struct {
	logger *logger.StructLogger
	db     *sqlx.DB
}

func New(
	log *logger.StructLogger,
	db *sqlx.DB,
) Server {
	return Server{
		logger: log,
		db:     db,
	}
}

func ParseBody[T interface{}](body io.ReadCloser, logger *logger.StructLogger) (*T, error) {
	var requestBody T

	byteBody, err := io.ReadAll(body)
	if err != nil {
		logger.WithError(err)
		logger.Error("cannot read request body")

		return nil, fmt.Errorf("cannot read request body: %w", err)
	}

	err = json.Unmarshal(byteBody, &requestBody)
	if err != nil {
		logger.WithError(err)
		logger.WithFields(slog.String("body", string(byteBody)))
		logger.Error("cannot unmarshal byte to struct")

		return nil, fmt.Errorf("cannot unmarshal byte to struct: %w", err)
	}

	return &requestBody, nil
}
