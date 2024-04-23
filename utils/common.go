package utils

import (
	"net/http"

	"github.com/rs/zerolog"
)

type AppError struct {
	message string
}

func (e AppError) Error() string {
	return e.message
}

type AppLogger struct {
	Logger zerolog.Logger
}

func (l *AppLogger) Info(msg string) {
	l.Logger.Info().Msg(msg)
}

func (l *AppLogger) Error(msg string) {
	l.Logger.Error().Msg(msg)
}

type ApiHandler struct {
	Logger AppLogger
}

type apiFunc func(
	w http.ResponseWriter, r *http.Request,
) error

func (h *ApiHandler) HandleFunc(
	f apiFunc,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			h.Logger.Error(err.Error())
		}
	}
}
