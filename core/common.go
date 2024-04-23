package core

import (
	"net/http"
)

type AppError struct {
	message string
}

func (e AppError) Error() string {
	return e.message
}

type ApiHandler struct {
	Logger AppLogger
}

type apiFunc func(
	w http.ResponseWriter, r *http.Request,
) error

func (h *ApiHandler) AddRoute(
	method string,
	path string,
	fn apiFunc,
	handler *http.ServeMux,
) {
	handler.HandleFunc(method+" "+path, h.HandleFunc(fn))
	h.Logger.Info("Register route: " + method + " " + path)
}

func (h *ApiHandler) HandleFunc(
	f apiFunc,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			h.Logger.Error(err.Error())
		}
	}
}
