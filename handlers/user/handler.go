package user

import (
	"expense-manager-backend/core"
	"net/http"
)

type UserHandler struct {
	*core.ApiHandler
}

func (h *UserHandler) RegisterRoutes() http.Handler {
	handler := http.NewServeMux()
	h.AddRoute("GET", "/users", h.CreateUser, handler)
	return handler
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) error {
	h.Logger.Info("Create user")

	w.Write([]byte("Create user"))

	return nil
}

func GetAllUser(w http.ResponseWriter, r *http.Request) error {
	return nil

}

func GetUser(w http.ResponseWriter, r *http.Request) error {
	return nil

}

func UpdateUser(w http.ResponseWriter, r *http.Request) error {
	return nil

}

func DeleteUser(w http.ResponseWriter, r *http.Request) error {
	return nil

}
