package middleware

import (
	"expense-manager-backend/core"
	"fmt"
	"net/http"
	"time"
)

type wrappedWriter struct {
	http.ResponseWriter
	statusCode int
}

func (w *wrappedWriter) WriteHeader(statusCode int) {
	w.ResponseWriter.WriteHeader(statusCode)
	w.statusCode = statusCode
}

type MiddlewareLogger struct {
	*core.Logger
}

func NewMiddlewareLogger(logger core.Logger) MiddlewareLogger {
	return MiddlewareLogger{
		Logger: &logger,
	}
}

func (m *MiddlewareLogger) Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		wrapped := &wrappedWriter{
			ResponseWriter: w,
			statusCode:     http.StatusOK,
		}
		next.ServeHTTP(wrapped, r)

		m.Logger.Info(
			fmt.Sprintf(
				"%d %s %s %s",
				wrapped.statusCode, r.Method, r.URL.Path, time.Since(start),
			))

	})
}
