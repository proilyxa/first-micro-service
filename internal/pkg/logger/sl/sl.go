package sl

import (
	"context"
	"github.com/go-chi/chi/v5/middleware"
	"golang.org/x/exp/slog"
	"net/http"
)

func Err(err error) slog.Attr {
	return slog.Attr{
		Key:   "error",
		Value: slog.StringValue(err.Error()),
	}
}

type ctxLogger struct{}

func ContextWithLogger(log *slog.Logger) func(next http.Handler) http.Handler {
	return middleware.WithValue(ctxLogger{}, log)
}

func LoggerFromContext(ctx context.Context) *slog.Logger {
	return ctx.Value(ctxLogger{}).(*slog.Logger)
}
