package http

import (
	"log/slog"

	"github.com/labstack/echo/v4"
)

func NewLoggingEchoMiddlware(logger *slog.Logger) LoggingMiddleware {
	return &loggingEchoMiddlware{
		logger: logger,
	}
}

type loggingEchoMiddlware struct {
	logger *slog.Logger
}

func (m *loggingEchoMiddlware) LoggingMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {

		return nil
	}
}
