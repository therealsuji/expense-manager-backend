package core

import (
	"expense-manager-backend/utils"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type ApiHandler struct {
	Logger    AppLogger
	Validator *validator.Validate
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
			if e, ok := err.(utils.ApiError); ok {
				utils.WriteError(w, e, e.Code)
			}

			if e, ok := err.(utils.AppError); ok {
				utils.WriteError(w, e, http.StatusInternalServerError)
			}
		}
	}
}
