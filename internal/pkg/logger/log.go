package logger

import (
	"log/slog"
)

type StructLogger struct {
	logger *slog.Logger
}

func New(logger *slog.Logger) *StructLogger {
	return &StructLogger{
		logger: logger,
	}
}

func (l *StructLogger) WithError(err error) *StructLogger {
	l.logger = l.logger.With(slog.String("error", err.Error()))

	return l
}

func (l *StructLogger) WithFields(fields ...slog.Attr) *StructLogger {
	for _, value := range fields {
		l.logger = l.logger.With(value)
	}

	return l
}

func (l *StructLogger) Error(msg string) {
	l.logger.Error(msg)
}
