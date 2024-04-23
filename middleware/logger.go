package middleware

import (
	"expense-manager-backend/utils"
	"fmt"
	"net/http"
	"time"

	"github.com/rs/zerolog"
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
	utils.AppLogger
}

func (a *MiddlewareLogger) Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		wrapped := &wrappedWriter{
			ResponseWriter: w,
			statusCode:     http.StatusOK,
		}

		next.ServeHTTP(wrapped, r)
		zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

		a.Info(
			fmt.Sprintf(
				"%d %s %s %s",
				wrapped.statusCode, r.Method, r.URL.Path, time.Since(start),
			))

	})
}
